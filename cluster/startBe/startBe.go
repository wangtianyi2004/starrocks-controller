
package startBe

import(
    "fmt"
    "time"
    "sr-controller/sr-utl"
    "sr-controller/module"
)

func StartBeCluster() {

    var infoMess string
    var err error
    var beStat BeStatusStruct

    // start Fe node one by one
    var tmpUser string
    var tmpKeyRsa string
    var tmpSshHost string
    var tmpSshPort int
    var tmpHeartbeatServicePort int
    var tmpBeDeployDir string
    var beStatusList string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = "/root/.ssh/id_rsa"

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        tmpSshHost = module.GYamlConf.BeServers[i].Host
        tmpSshPort = module.GYamlConf.BeServers[i].SshPort
        tmpHeartbeatServicePort = module.GYamlConf.BeServers[i].HeartbeatServicePort
        tmpBeDeployDir = module.GYamlConf.BeServers[i].DeployDir

	infoMess = fmt.Sprintf("Starting BE node [BeHost = %s HeartbeatServicePort = %d]", tmpSshHost, tmpHeartbeatServicePort)
        utl.Log("INFO", infoMess)

	for startTimeInd := 0; startTimeInd < 3; startTimeInd++ {

	    infoMess = fmt.Sprintf("The %d time to start [%s]",(startTimeInd + 1), tmpSshHost)
            utl.Log("DEBUG", infoMess)
	    // startBeNode(user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int, beDeployDir string) (err error)
	    err = startBeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpHeartbeatServicePort, tmpBeDeployDir)

	    startWaitTime := time.Duration(20 - startTimeInd * 5)
	    time.Sleep(startWaitTime  * time.Second)

            beStat, _ = CheckBeStatus(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpHeartbeatServicePort)
            if beStat.Alive {
                infoMess = fmt.Sprintf("The BE node start succefully [host = %s, heartbeatServicePort = %d]", tmpSshHost, tmpHeartbeatServicePort)
                utl.Log("INFO", infoMess)
                break
            } else {
                infoMess = fmt.Sprintf("The BE node doesn't start, wait for 10s [BeHost = %s, HeartbeatServicePort = %d, error = %v]", tmpSshHost, tmpHeartbeatServicePort, err)
                utl.Log("WARN", infoMess)
            }
        } // FOR-END: 3 time to restart BE node

	if !beStat.Alive {
             infoMess = fmt.Sprintf("The BE node start failed [BeHost = %s, HeartbeatServicePort = %d, error = %v]", tmpSshHost, tmpHeartbeatServicePort, err)
        }

	beStatusList = beStatusList + "                                        " + fmt.Sprintf("beHost = %-20sbeHeartbeatServicePort = %d\tbeStatus = %v\n", tmpSshHost, tmpHeartbeatServicePort, beStat.Alive)
    }
    beStatusList = "List all BE status:\n" + beStatusList
    utl.Log("OUTPUT", beStatusList)
}

func startBeNode(user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int, beDeployDir string) (err error) {

    var infoMess string


    addBeSQL := fmt.Sprintf("alter system add backend \"%s:%d\"", sshHost, heartbeatServicePort)
    addBeCMD := fmt.Sprintf("%s/bin/start_be.sh --daemon", beDeployDir)

    //infoMess = fmt.Sprintf("Starting BE node [host = %s, heartbeatServicePort = %d]", sshHost, heartbeatServicePort)
    //utl.Log("INFO", infoMess)

    // alter system add backend "sshHost:heartbeatServicePort"
    sqlUserName := "root"
    sqlPassword := ""
    sqlIp := module.GYamlConf.FeServers[0].Host
    sqlPort := module.GYamlConf.FeServers[0].QueryPort
    sqlDbName := ""

    _, err = utl.RunSQL(sqlUserName, sqlPassword, sqlIp, sqlPort, sqlDbName, addBeSQL)
    if err != nil {
        infoMess = fmt.Sprintf(`Error in add follower BE node, [
                                        sqlUserName = %s
                                        sqlPassword = %s
                                        sqlIP = %s
                                        sqlPort = %d
                                        sqlDBName = %s
                                        addFollowerSQL =%s
                                        errMess = %v]`, sqlUserName, sqlPassword, sqlIp, sqlPort, sqlDbName, addBeSQL, err)
        utl.Log("ERROR", infoMess)
        return err
    }

    // run beDeploy/bin/start_be.sh --daemon 
    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, addBeCMD)
    if err != nil {
        infoMess = fmt.Sprintf(`Waiting for startMastertFeNode:
                                        user = %s
                                        keyRsa = %s
                                        sshHost = %s
                                        sshPort = %d
                                        beDeployDir = %s`,
                user, keyRsa, sshHost, sshPort, beDeployDir)
        utl.Log("WARN", infoMess)
        return err
    }

    utl.Log("INFO", "广告招租 ****************************")
    utl.Log("INFO", "充值，跳过广告 **********************")
    // time.Sleep(5 * time.Second)
    return nil

}
