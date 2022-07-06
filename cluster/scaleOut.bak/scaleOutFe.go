package scaleOutCluster

import (
    "fmt"
    "time"
    "stargo/sr-utl"
    "stargo/module"
    "stargo/cluster/startCluster"
    "stargo/cluster/checkStatus"
)



func ScaleOutFeCluster() {

    var infoMess            string
    var err                 error
    var feStat              checkStatus.FeStatusStruct

    // start Fe node one by one
    var tmpUser             string
    var tmpKeyRsa           string
    var tmpSshHost          string
    var tmpSshPort          int
    var tmpEditLogPort      int
    var tmpQueryPort        int
    var tmpFeDeployDir      string
    var feStatusList        string

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    // GYamlConfAppend is for scale-out.yaml   (node6)
    // GYamlConf is for sr-c1/meta.yaml        (node3,4,5)


    for i := 0; i < len(module.GYamlConfAppend.FeServers); i++ {

        tmpSshHost = module.GYamlConfAppend.FeServers[i].Host
        tmpSshPort = module.GYamlConfAppend.FeServers[i].SshPort
        tmpEditLogPort = module.GYamlConfAppend.FeServers[i].EditLogPort
        tmpQueryPort = module.GYamlConfAppend.FeServers[i].QueryPort
        tmpFeDeployDir = module.GYamlConfAppend.FeServers[i].DeployDir

        //infoMess = fmt.Sprintf("Starting FE node [FeHost = %s, FeEditLogPort = %d]", tmpSshHost, tmpEditLogPort)
        //utl.Log("INFO", infoMess)

        for startTimeInd := 0; startTimeInd < 3; startTimeInd++ {
            // initFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error)
            infoMess = fmt.Sprintf("The %d time to start [%s]", (startTimeInd + 1), tmpSshHost)
            utl.Log("DEBUG", infoMess)
            err = startCluster.InitFeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpEditLogPort, tmpFeDeployDir)
            startWaitTime := time.Duration(20 - startTimeInd * 5)
            time.Sleep(startWaitTime * time.Second)

            feStat, err = checkStatus.CheckFeStatus(i)

            if err != nil {
                infoMess = fmt.Sprintf("Error in get the fe status [FeHost = %s, error = %v]", tmpSshHost, err)
                utl.Log("DEBUG", infoMess)
            }
            if feStat.FeAlive {
                infoMess = fmt.Sprintf("The FE node start succefully [host = %s, queryPort = %d]", tmpSshHost, tmpQueryPort)
                utl.Log("INFO", infoMess)
                break
            } else {
                infoMess = fmt.Sprintf("The FE node doesn't start, wait for 10s [FeHost = %s, FeQueryPort = %d, error = %v]", tmpSshHost, tmpQueryPort, err)
                utl.Log("WARN", infoMess)
            }
        } // FOR-END: 3 time to restart FE node

        if !feStat.FeAlive {
            infoMess = fmt.Sprintf("The FE node start failed [host = %s, queryPort = %d, error = %v]", tmpSshHost, tmpQueryPort, err)
            utl.Log("ERROR", infoMess)
        }
        feStatusList = feStatusList + "                                        " + fmt.Sprintf("feHost = %-20sfeQueryPort = %d     feStatus = true\n", tmpSshHost, tmpQueryPort)
    } // FOR-END: list all FE node

    feStatusList = "List all FE status:\n" + feStatusList
    utl.Log("INFO", feStatusList)

}
