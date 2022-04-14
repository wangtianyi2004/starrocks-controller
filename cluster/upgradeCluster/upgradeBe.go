package upgradeCluster

import (

    "fmt"
    "time"
    "strings"
    //"errors"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/stopCluster"
    "sr-controller/cluster/startCluster"
    "sr-controller/cluster/checkStatus"
    //"sr-controller/cluster/prepareOption"

)


func UpgradeBeCluster() { //(err error){

    var infoMess       string
    var err            error
    var beStat         checkStatus.BeStatusStruct
    var feEntryId      int


    feEntryId, err = checkStatus.GetFeEntry()
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
        err = UpgradeBeNode(i)
        if err != nil {
            infoMess = fmt.Sprintf("Error in upgrade be node. [nodeid = %d]", i)
            utl.Log("ERROR", infoMess)
        }

        beStat, err = checkStatus.CheckBeStatus(i)


        for j := 0; j < 3; j++ {
            infoMess = fmt.Sprintf("The %d time to check be status: %v", j, beStat.Alive)
            utl.Log("DEBUG", infoMess)
            if beStat.Alive {
                break
            } else {
                infoMess = fmt.Sprintf("The BE node doesn't work, wait for 10s and check the status again. [beId = %d]\n", i)
                utl.Log("DEBUG", infoMess)
                time.Sleep(10 * time.Second)
                beStat, err = checkStatus.CheckBeStatus(i)
            }
        }


        if err != nil {
            infoMess = fmt.Sprintf("Error in get the Be status [beId = %d, error = %v]", i, err)
            utl.Log("DEBUG", infoMess)
            //return err
        }
        if !beStat.Alive {
            infoMess = fmt.Sprintf("The BE node upgrade failed. The BE node doesn't work. [beId = %d]\n", i)
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else if ! strings.Contains(beStat.Version.String, strings.Replace(module.GSRVersion, "v", "", -1)) {
            infoMess = fmt.Sprintf("The BE node upgrade failed.  [beId = %d, targetVersion = %s, currentVersion = v%s]", i, module.GSRVersion, beStat.Version.String)
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else {
            infoMess = fmt.Sprintf("The Be node upgrade successfully. [beId = %d, currentVersion = v%s]", i, beStat.Version.String)
            utl.Log("OUTPUT", infoMess)
        }
    }

    //return nil
    

}

 
func UpgradeBeNode(beId int) (err error) {
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
        return err
    } else {
        infoMess = fmt.Sprintf("upgrade be node - backup be lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("INFO", infoMess)
    }


    // step2. upload new be lib
    sourceDir = fmt.Sprintf("%s/download/StarRocks-%s/be/lib", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
    targetDir = fmt.Sprintf("%s/lib", beDeployDir)
    utl.UploadDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    infoMess = fmt.Sprintf("upgrade be node - upload new be lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
    utl.Log("INFO", infoMess)



    // step3. stop be node
    err = stopCluster.StopBeNode(user, keyRsa, sshHost, sshPort, beDeployDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in stop be node when upgrade be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
        utl.Log("ERROR", infoMess)
        return err   
    } else {
        infoMess = fmt.Sprintf("upgrade be node - stop be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
        utl.Log("INFO", infoMess)
    }

    // step4. start be node
    startCluster.StartBeNode(user, keyRsa, sshHost, sshPort, beHeartBeatServicePort, beDeployDir)
    infoMess = fmt.Sprintf("upgrade be node - start be node. [host = %s, beDeployDir = %s]", sshHost, beDeployDir)
    utl.Log("INFO", infoMess)
        

    return nil

}




func TestUpgradeBe() {

    module.InitConf("sr-c1", "")
    module.SetGlobalVar("v2.1.3")
    //utl.RenameDir("starrocks", "/home/sr-dev/.ssh/id_rsa", "192.168.88.83", 22, "/tmp/aaa", "/tmp/bbb")
    err := UpgradeBeNode(0)
    if err != nil {
        panic(err)
    }

}
