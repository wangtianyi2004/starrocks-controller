
package startCluster

import(
    "fmt"
    "time"
//    "errors"
    "stargo/sr-utl"
    "stargo/module"
    "stargo/cluster/checkStatus"
)




func StartBeCluster() {


    // start Be node one by one
    var infoMess string
    var tmpUser string
    var tmpKeyRsa string
    var tmpSshHost string
    var tmpSshPort int
    var tmpHeartbeatServicePort int
    //var tmpQueryPort int
    var tmpBeDeployDir string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {
    // for i := 0; i < 1; i++ { ## debug leader node

        tmpSshHost = module.GYamlConf.BeServers[i].Host
        tmpSshPort = module.GYamlConf.BeServers[i].SshPort
        tmpHeartbeatServicePort = module.GYamlConf.BeServers[i].HeartbeatServicePort
        tmpBeDeployDir = module.GYamlConf.BeServers[i].DeployDir

        infoMess = fmt.Sprintf("Starting BE node [BeHost = %s, HeartbeatServicePort = %d]", tmpSshHost, tmpHeartbeatServicePort)
        utl.Log("INFO", infoMess)

        _ = StartBeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpHeartbeatServicePort, tmpBeDeployDir)
        for j := 0; j < 3; j++ {
            portStat, _ := checkStatus.CheckBePortStatus(i)
            if portStat {
                break
                //time.Sleep(10 * time.Second)
            } else {
                time.Sleep(10 * time.Second)
            }
        }
    }
}


func StartBeNode(user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int, beDeployDir string) (err error) {

    var infoMess string


    startBeCMD := fmt.Sprintf("%s/bin/start_be.sh --daemon", beDeployDir)

    infoMess = fmt.Sprintf("Starting BE node [host = %s, heartbeatServicePort = %d]", sshHost, heartbeatServicePort)
    utl.Log("DEBUG", infoMess)


    // run beDeploy/bin/start_be.sh --daemon 
    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, startBeCMD)
    if err != nil {
        infoMess = fmt.Sprintf("Waiting for start BE node.[BeHost = %s, Error =  %v", sshHost, err)
        utl.Log("DEBUG", infoMess)
        return err
    }

    // time.Sleep(5 * time.Second)
    return nil

}
