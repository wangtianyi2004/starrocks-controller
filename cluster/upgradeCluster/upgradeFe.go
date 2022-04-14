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


func UpgradeFeCluster() { //(err error){

    var infoMess       string
    var err            error
    var feStat         checkStatus.FeStatusStruct
    var feEntryId      int


    feEntryId, err = checkStatus.GetFeEntry()
    if err != nil ||  feEntryId == -1 {
        //infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        //utl.Log("WARN", infoMess)
        module.SetFeEntry(0)
    } else {
        module.SetFeEntry(feEntryId)
    }


    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        infoMess = fmt.Sprintf("Starting upgrade BE node. [feId = %d]", i)
        utl.Log("OUTPUT", infoMess)
        err = UpgradeFeNode(i)
        if err != nil {
            infoMess = fmt.Sprintf("Error in upgrade FE node. [nodeid = %d]", i)
            utl.Log("ERROR", infoMess)
        }

        feStat, err = checkStatus.CheckFeStatus(i)


        for j := 0; j < 3; j++ {
            infoMess = fmt.Sprintf("The %d time to check FE status: %v", j, feStat.FeAlive)
            utl.Log("DEBUG", infoMess)
            if feStat.FeAlive {
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
        if !feStat.FeAlive {
            infoMess = fmt.Sprintf("The BE node upgrade failed. The BE node doesn't work. [feId = %d]\n", i)
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else if ! strings.Contains(feStat.FeVersion.String, strings.Replace(module.GSRVersion, "v", "", -1)) {
            infoMess = fmt.Sprintf("The BE node upgrade failed.  [feId = %d, targetVersion = %s, currentVersion = v%s]", i, module.GSRVersion, feStat.FeVersion.String)
            utl.Log("ERROR", infoMess)
            //return errors.New(infoMess)
        } else {
            infoMess = fmt.Sprintf("The Fe node upgrade successfully. [feId = %d, currentVersion = v%s]", i, feStat.FeVersion.String)
            utl.Log("OUTPUT", infoMess)
        }
    }

    //return nil
    

}

 
func UpgradeFeNode(feId int) (err error) {
    // step 1. backup fe lib
    // step 2. upload new fe lib
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
        infoMess = fmt.Sprintf("upgrade FE node - backup FE lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
        utl.Log("INFO", infoMess)
    }


    // step2. upload new FE lib
    sourceDir = fmt.Sprintf("%s/download/StarRocks-%s/fe/lib", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
    targetDir = fmt.Sprintf("%s/lib", feDeployDir)
    utl.UploadDir(user, keyRsa, sshHost, sshPort, sourceDir, targetDir)
    infoMess = fmt.Sprintf("upgrade FE node - upload new FE lib. [host = %s, sourceDir = %s, targetDir = %s]", sshHost, sourceDir, targetDir)
    utl.Log("INFO", infoMess)



    // step3. stop FE node
    err = stopCluster.StopFeNode(user, keyRsa, sshHost, sshPort, feDeployDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in stop FE node when upgrade FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
        utl.Log("ERROR", infoMess)
        return err   
    } else {
        infoMess = fmt.Sprintf("upgrade FE node - stop FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
        utl.Log("INFO", infoMess)
    }

    // step4. start FE node
    startCluster.StartFeNode(user, keyRsa, sshHost, sshPort, feEditLogPort, feDeployDir)
    infoMess = fmt.Sprintf("upgrade FE node - start FE node. [host = %s, feDeployDir = %s]", sshHost, feDeployDir)
    utl.Log("INFO", infoMess)
        

    return nil

}




func TestUpgradeFe() {

    module.InitConf("sr-c1", "")
    module.SetGlobalVar("v2.1.3")
    //utl.RenameDir("starrocks", "/home/sr-dev/.ssh/id_rsa", "192.168.88.83", 22, "/tmp/aaa", "/tmp/bbb")
    err := UpgradeFeNode(0)
    if err != nil {
        panic(err)
    }

}