package stopCluster

import (
    "fmt"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "sr-controller/cluster/checkStatus"
)

// func startFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error) {

func StopBeNode(user string, keyRsa string, sshHost string, sshPort int, beDeployDir string) (err error){

    var infoMess string
    var stopBeCmd string

    // /opt/starrocks/be/bin/stop_be.sh
    stopBeCmd = fmt.Sprintf("%s/bin/stop_be.sh", beDeployDir)
 
    infoMess = fmt.Sprintf("Waiting for stoping BE node [BeHost = %s]", sshHost)
    utl.Log("INFO", infoMess)
    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, stopBeCmd)
    if err != nil {
        infoMess = fmt.Sprintf("Stop BE failed [BeHost = %s, error = %v]", sshHost, err)
        utl.Log("DEBUG", infoMess)
        return err
    }
    return nil

}

func StopBeCluster(clusterName string) {

    var infoMess                   string
    var err                        error
    var beStat                     map[string]string


    // Stop BE node one by one
    var tmpUser                    string
    var tmpKeyRsa                  string
    var tmpSshHost                 string
    var tmpSshPort                 int
    var tmpBeDeployDir             string
    var tmpHeartbeatServicePort    int
    //var beStatusList             string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    infoMess = "Stop cluster " + clusterName
    utl.Log("OUTPUT", infoMess)
    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        tmpSshHost = module.GYamlConf.BeServers[i].Host
        tmpSshPort = module.GYamlConf.BeServers[i].SshPort
        tmpBeDeployDir = module.GYamlConf.BeServers[i].DeployDir
        tmpHeartbeatServicePort = module.GYamlConf.BeServers[i].HeartbeatServicePort
        // func StopFeNode(user string, keyRsa string, sshHost string, sshPort int, feDeployDir string) (err error)
        err = StopBeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpBeDeployDir)
        if err != nil {
            infoMess = fmt.Sprintf("Error in stoping BE node [BeHost = %s, HeartbeatServicePort = %d, error = %v]", tmpSshHost, tmpHeartbeatServicePort, err)
            utl.Log("DEBUG", infoMess)
        }

        beStat, err = checkStatus.CheckBeStatus(i)

        if err != nil {
            infoMess = fmt.Sprintf("Error in get the Be status [BeHost = %s, HeartbeatServicePort = %d, error = %v]", tmpSshHost, tmpHeartbeatServicePort, err)
            utl.Log("DEBUG", infoMess)
        }
        if beStat["Alive"] == "false" {
            infoMess = fmt.Sprintf("The BE node stop succefully [BeHost = %s, HeartbeatServicePort = %d]", tmpSshHost, tmpHeartbeatServicePort)
            utl.Log("INFO", infoMess)
        } else {
            infoMess = fmt.Sprintf("The BE node stop failed [BeHost = %s, HeartbeatServicePort = %d]", tmpSshHost, tmpHeartbeatServicePort)
            utl.Log("ERROR", infoMess)
        }
    }

}
