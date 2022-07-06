package startCluster

import (
    "fmt"
    "time"
    "stargo/sr-utl"
    "stargo/module"
    "stargo/cluster/checkStatus"
)



func StartFeCluster() {

    var infoMess string
    //var err error
    //var feStat checkStatus.FeStatusStruct

    // start Fe node one by one
    var tmpUser string
    var tmpKeyRsa string
    var tmpSshHost string
    var tmpSshPort int
    var tmpEditLogPort int
    //var tmpQueryPort int
    var tmpFeDeployDir string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
    // for i := 0; i < 1; i++ { ## debug leader node

        tmpSshHost = module.GYamlConf.FeServers[i].Host
        tmpSshPort = module.GYamlConf.FeServers[i].SshPort
        tmpEditLogPort = module.GYamlConf.FeServers[i].EditLogPort
        //tmpQueryPort = module.GYamlConf.FeServers[i].QueryPort
        tmpFeDeployDir = module.GYamlConf.FeServers[i].DeployDir

        infoMess = fmt.Sprintf("Starting FE node [FeHost = %s, EditLogPort = %d]", tmpSshHost, tmpEditLogPort)
        utl.Log("INFO", infoMess)

        // startFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error)
        _ = StartFeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpEditLogPort, tmpFeDeployDir)
	for j := 0; j < 3; j++ {
            portStat, _ := checkStatus.CheckFePortStatus(i)
	    if portStat {
                break
                //time.Sleep(10 * time.Second)
            } else {
                time.Sleep(10 * time.Second)
	    }
        }
    }
}






func StartFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error) {


    var infoMess string
    //var isMasterFe bool
    var startFeCmd string

    // check master node
    startFeCmd = fmt.Sprintf("%s/bin/start_fe.sh --daemon", feDeployDir)
    infoMess = fmt.Sprintf("Run starting FE process [host = %s, editLogPort = %d]", sshHost, editLogPort)
    utl.Log("DEBUG", infoMess)
    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, startFeCmd)

    if err != nil {
        infoMess = fmt.Sprintf("Waiting for starting FE node [FeHost = %s]", sshHost)
        utl.Log("DEBUG", infoMess)
        return err
    }
    return nil

}

