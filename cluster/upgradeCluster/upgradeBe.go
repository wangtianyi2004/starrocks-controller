package upgradeCluster

import (

    "fmt"
    "time"
    "strings"
    //"errors"
    "stargo/module"
    "stargo/sr-utl"
    "stargo/cluster/stopCluster"
    "stargo/cluster/startCluster"
    "stargo/cluster/checkStatus"
    //"stargo/cluster/prepareOption"

)


func UpgradeBeCluster() { //(err error){

    var infoMess       string
    var err            error
    var feEntryId      int


    feEntryId, err = checkStatus.GetFeEntry(-1)
    if err != nil ||  feEntryId == -1 {
        //infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        //utl.Log("WARN", infoMess)
        module.SetFeEntry(0)
    } else {
        module.SetFeEntry(feEntryId)
    }


    for i := 0; i < len(module.GYamlConf.BeServers); i++ {
        infoMess = fmt.Sprintf("Starting upgrade BE node. [beId = %d]", i)
        utl.Log("OUTPUT", infoMess)
        UpgradeBeNode(i)
    }

}

 
func UpgradeBeNode(beId int) {
    // step 1. backup be lib
    // step 2. upload new be lib
    // step 3. stop be node
    // step 4. start be node

    var infoMess                   string
    var user                       string
    var sourceDir                  string
    var targetDir                  string
    var sshHost                    string
    var sshPort                    int
    var beDeployDir                string
    var beHeartBeatServicePort     int
    var keyRsa                     string
    var beStat                     map[string]string
    var err                        error


    user = module.GYamlConf.Global.User
    keyRsa = module.GSshKeyRsa
    sshHost = module.GYamlConf.BeServers[beId].Host
    sshPort = module.GYamlConf.BeServers[beId].SshPort
    beDeployDir = module.GYamlConf.BeServers[beId].DeployDir
    beHeartBeatServicePort = module.GYamlConf.BeServers[beId].HeartbeatServicePort




    // step1. backup be lib
    sourceDir = fmt.Sprintf("%s/lib", beDeployDir)
    targetDir = fmt.Sprintf("%s/lib.bak-%s", beDeployDir, time.Now().Format("20060102150405"))

    err = utl.RenameDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in rename dir when backup be lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("ERROR", infoMess)
    } else {
        infoMess = fmt.Sprintf("upgrade be node - backup be lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("INFO", infoMess)
    }


    // step2. upload new be lib
    sourceDir = fmt.Sprintf("%s/StarRocks-%s/be/lib", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
    // sourceDir = fmt.Sprintf("%s/download/StarRocks-%s/be/lib", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
    targetDir = fmt.Sprintf("%s/lib", beDeployDir)
    utl.UploadDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    infoMess = fmt.Sprintf("upgrade be node - upload new be lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
    utl.Log("INFO", infoMess)



    // step3. stop be node
    err = stopCluster.StopBeNode(user, keyRsa, sshHost, sshPort, beDeployDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in stop be node when upgrade be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
        utl.Log("ERROR", infoMess)
    } else {
        infoMess = fmt.Sprintf("upgrade be node - stop be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
        utl.Log("INFO", infoMess)
    }

    // step4. start be node

    for j := 0; j < 3; j++ {
        startCluster.StartBeNode(user, keyRsa, sshHost, sshPort, beHeartBeatServicePort, beDeployDir)
        infoMess = fmt.Sprintf("upgrade be node - start be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
        utl.Log("INFO", infoMess)

        beStat, err = checkStatus.CheckBeStatus(beId)
        if beStat["Alive"] == "true" && strings.Contains(beStat["Version"], strings.Replace(module.GSRVersion, "v", "", -1)) {
            break
        }
        time.Sleep(10 * time.Second)
    }

    if err != nil {
        infoMess = fmt.Sprintf("Error in get the Be status [beId = %d, error = %v]", beId, err)
        utl.Log("DEBUG", infoMess)
    } else if beStat["Alive"] == "false" {
        infoMess = fmt.Sprintf("The BE node upgrade failed. The BE node doesn't work. [beId = %d]\n", beId)
        utl.Log("ERROR", infoMess)
    } else if ! strings.Contains(beStat["Version"], strings.Replace(module.GSRVersion, "v", "", -1)) {
        infoMess = fmt.Sprintf("The BE node upgrade failed.  [beId = %d, targetVersion = %s, currentVersion = v%s]", beId, module.GSRVersion, beStat["Version"])
        utl.Log("ERROR", infoMess)
    } else {
        infoMess = fmt.Sprintf("The Be node upgrade successfully. [beId = %d, currentVersion = v%s]", beId, beStat["Version"])
        utl.Log("OUTPUT", infoMess)
    }

}




