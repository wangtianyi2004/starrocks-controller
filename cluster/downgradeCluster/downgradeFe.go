package downgradeCluster

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


func DowngradeFeCluster() { //(err error){

    var infoMess       string
    var err            error
    var feStat         map[string]string
    var feEntryId      int


    feEntryId, err = checkStatus.GetFeEntry(-1)
    if err != nil ||  feEntryId == -1 {
        //infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        //utl.Log("WARN", infoMess)
        module.SetFeEntry(0)
    } else {
        module.SetFeEntry(feEntryId)
    }


    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        infoMess = fmt.Sprintf("Starting downgrade FE node. [feId = %d]", i)
        utl.Log("OUTPUT", infoMess)
        err = DowngradeFeNode(i)
        if err != nil {
            infoMess = fmt.Sprintf("Error in downgrade FE node. [nodeid = %d]", i)
            utl.Log("ERROR", infoMess)
        }

        feStat, err = checkStatus.CheckFeStatus(i)


        for j := 0; j < 3; j++ {
            infoMess = fmt.Sprintf("The %d time to check FE status: %v", j, feStat["Alive"])
            utl.Log("DEBUG", infoMess)
            if feStat["Alive"] == "true"{
                break
            } else {
                infoMess = fmt.Sprintf("The FE node doesn't work, wait for 10s and check the status again. [feId = %d]\n", i)
                utl.Log("DEBUG", infoMess)
                time.Sleep(10 * time.Second)
                feStat, err = checkStatus.CheckFeStatus(i)
            }
        }


        if err != nil {
            infoMess = fmt.Sprintf("Error in get the FE status [feId = %d, error = %v]", i, err)
            utl.Log("DEBUG", infoMess)
            //return err
        }
        if feStat["Alive"] == "false" {
            infoMess = fmt.Sprintf("The FE node downgrade failed. The FE node doesn't work. [feId = %d]\n", i)
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else if ! strings.Contains(feStat["FeVersion"], strings.Replace(module.GSRVersion, "v", "", -1)) {
            infoMess = fmt.Sprintf("The FE node downgrade failed.  [feId = %d, targetVersion = %s, currentVersion = v%s]", i, module.GSRVersion, feStat["FeVersion"])
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else {
            infoMess = fmt.Sprintf("The Fe node downgrade successfully. [feId = %d, currentVersion = v%s]", i, feStat["FeVersion"])
            utl.Log("OUTPUT", infoMess)
        }
    }

    //return nil


}


func DowngradeFeNode(feId int) (err error) {
    // step 1. backup fe lib
    // step 2. download new fe lib
    // step 3. stop fe node
    // step 4. start fe node

    var infoMess                   string
    var user                       string
    var sourceDir                  string
    var targetDir                  string
    var sshHost                    string
    var sshPort                    int
    var feDeployDir                string
    var feEditLogPort              int
    var keyRsa                     string


    user = module.GYamlConf.Global.User
    keyRsa = module.GSshKeyRsa
    sshHost = module.GYamlConf.FeServers[feId].Host
    sshPort = module.GYamlConf.FeServers[feId].SshPort
    feDeployDir = module.GYamlConf.FeServers[feId].DeployDir
    feEditLogPort = module.GYamlConf.FeServers[feId].EditLogPort




    // step1. backup fe lib
    sourceDir = fmt.Sprintf("%s/lib", feDeployDir)
    targetDir = fmt.Sprintf("%s/lib.bak-%s", feDeployDir, time.Now().Format("20060102150405"))

    err = utl.RenameDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in rename dir when backup FE lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("ERROR", infoMess)
        return err
    } else {
        infoMess = fmt.Sprintf("downgrade FE node - backup FE lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("INFO", infoMess)
    }


    // step2. download new FE lib
    sourceDir = fmt.Sprintf("%s/StarRocks-%s/fe/lib", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
    // sourceDir = fmt.Sprintf("%s/download/StarRocks-%s/fe/lib", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
    targetDir = fmt.Sprintf("%s/lib", feDeployDir)
    utl.UploadDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    infoMess = fmt.Sprintf("downgrade FE node - download new FE lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
    utl.Log("INFO", infoMess)



    // step3. stop FE node
    err = stopCluster.StopFeNode(user, keyRsa, sshHost, sshPort, feDeployDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in stop FE node when downgrade FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
        utl.Log("ERROR", infoMess)
        return err   
    } else {
        infoMess = fmt.Sprintf("downgrade FE node - stop FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
        utl.Log("INFO", infoMess)
    }

    // step4. start FE node
    startCluster.StartFeNode(user, keyRsa, sshHost, sshPort, feEditLogPort, feDeployDir)
    infoMess = fmt.Sprintf("downgrade FE node - start FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
    utl.Log("INFO", infoMess)
        

    return nil

}




