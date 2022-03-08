package stopFe

import (
    "fmt"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "sr-controller/cluster/checkStatus"
)

// func startFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error) {

func StopFeNode(user string, keyRsa string, sshHost string, sshPort int, feDeployDir string) (err error){

    var infoMess string
    var stopFeCmd string

    // /opt/starrocks/fe/bin/stop_fe.sh
    stopFeCmd = fmt.Sprintf("%s/bin/stop_fe.sh", feDeployDir)

    infoMess = fmt.Sprintf("Waiting for stoping FE node [FeHost = %s]", sshHost)
    utl.Log("INFO", infoMess)
    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, stopFeCmd)
    if err != nil {
        infoMess = fmt.Sprintf("Stop FE failed [FeHost = %s, error = %v]", sshHost, err)
        utl.Log("INFO", infoMess)
        return err
    }
    return nil

}

func StopFeCluster(clusterName string) {

    var infoMess string
    var err error
    var feStat checkStatus.FeStatusStruct

    // start Fe node one by one
    var tmpUser string
    var tmpKeyRsa string
    var tmpSshHost string
    var tmpSshPort int
    var tmpFeDeployDir string
    var tmpFeQueryPort int
    //var feStatusList string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    infoMess = "Stop cluster " + clusterName
    utl.Log("OUTPUT", infoMess)
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        tmpSshHost = module.GYamlConf.FeServers[i].Host
        tmpSshPort = module.GYamlConf.FeServers[i].SshPort
        tmpFeQueryPort = module.GYamlConf.FeServers[i].QueryPort
        tmpFeDeployDir = module.GYamlConf.FeServers[i].DeployDir

        // func StopFeNode(user string, keyRsa string, sshHost string, sshPort int, feDeployDir string) (err error)
        err = StopFeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpFeDeployDir)
        if err != nil {
            infoMess = fmt.Sprintf("Error in stoing FE node [FeHost = %s]", tmpSshHost)
        }

        feStat, err = checkStatus.CheckFeStatus(i)

        if err != nil {
            infoMess = fmt.Sprintf("Error in get the fe status [FeHost = %s, error = %v]", tmpSshHost, err)
            utl.Log("DEBUG", infoMess)
        }
        if !feStat.FeAlive {
            infoMess = fmt.Sprintf("The FE node stop succefully [host = %s, queryPort = %d]", tmpSshHost, tmpFeQueryPort)
            utl.Log("OUTPUT", infoMess)
        } else {
            infoMess = fmt.Sprintf("The FE node stop failed [host = %s, queryPort = %d]", tmpSshHost, tmpFeQueryPort)
            utl.Log("ERROR", infoMess)
        }
    }

}
