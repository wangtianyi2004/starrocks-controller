package startCluster

import (
    "fmt"
    "time"
    "strings"
    "stargo/sr-utl"
    "stargo/module"
    "stargo/cluster/checkStatus"
)


func InitFeCluster(yamlConf *module.ConfStruct) {

    var infoMess string
    var err error
    var feStat map[string]string

    // start Fe node one by one
    var tmpUser string
    var tmpKeyRsa string
    var tmpSshHost string
    var tmpSshPort int
    var tmpEditLogPort int
    var tmpQueryPort int
    var tmpFeDeployDir string
    var feStatusList string
    var feEntryId  int
    tmpUser = yamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa


    // get FE entry
    feEntryId, err = checkStatus.GetFeEntry(-1)

    if err != nil ||  feEntryId == -1 {
        //infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        //utl.Log("WARN", infoMess)
        module.SetFeEntry(0)
    } else {
        module.SetFeEntry(feEntryId)
    }

    for i := 0; i < len(yamlConf.FeServers); i++ {
    // for i := 0; i < 1; i++ { ## debug leader node

        tmpSshHost = yamlConf.FeServers[i].Host
        tmpSshPort = yamlConf.FeServers[i].SshPort
        tmpEditLogPort = yamlConf.FeServers[i].EditLogPort
        tmpQueryPort = yamlConf.FeServers[i].QueryPort
        tmpFeDeployDir = yamlConf.FeServers[i].DeployDir

        //infoMess = fmt.Sprintf("Starting FE node [FeHost = %s, FeEditLogPort = %d]", tmpSshHost, tmpEditLogPort)
        //utl.Log("INFO", infoMess)

        for startTimeInd := 0; startTimeInd < 3; startTimeInd++ {
            // initFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error)
            infoMess = fmt.Sprintf("The %d time to start [%s]", (startTimeInd + 1), tmpSshHost)
            utl.Log("DEBUG", infoMess)
            err = InitFeNode(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpEditLogPort, tmpFeDeployDir)
            startWaitTime := time.Duration(20 - startTimeInd * 5)
            time.Sleep(startWaitTime * time.Second)

            feStat, err = checkStatus.CheckFeStatus(i)

            if err != nil {
                infoMess = fmt.Sprintf("Error in get the fe status [FeHost = %s, error = %v]", tmpSshHost, err)
                utl.Log("DEBUG", infoMess)
            }
            if feStat["Alive"] == "true" {
                infoMess = fmt.Sprintf("The FE node start succefully [host = %s, queryPort = %d]", tmpSshHost, tmpQueryPort)
                utl.Log("INFO", infoMess)
                break
            } else {
                infoMess = fmt.Sprintf("The FE node doesn't start, wait for 10s [FeHost = %s, FeQueryPort = %d, error = %v]", tmpSshHost, tmpQueryPort, err)
                utl.Log("WARN", infoMess)
            }
        } // FOR-END: 3 time to restart FE node

        if feStat["Alive"] == "false" {
            infoMess = fmt.Sprintf("The FE node start failed [host = %s, queryPort = %d, error = %v]", tmpSshHost, tmpQueryPort, err)
            utl.Log("ERROR", infoMess)
        }
        feStatusList = feStatusList + "                                        " + fmt.Sprintf("feHost = %-20sfeQueryPort = %d     feStatus = true\n", tmpSshHost, tmpQueryPort)
    } // FOR-END: list all FE node

    feStatusList = "List all FE status:\n" + feStatusList
    utl.Log("INFO", feStatusList)

}




func InitFeNode(user string, keyRsa string, sshHost string, sshPort int, editLogPort int, feDeployDir string) (err error) {


    var infoMess string
    //var isMasterFe bool
    var startFeCmd string
    // check master node
    if sshHost == module.GYamlConf.FeServers[0].Host && editLogPort == module.GYamlConf.FeServers[0].EditLogPort {
        //isMasterFe = true
        infoMess = fmt.Sprintf("Starting leader FE node [host = %s, editLogPort = %d]", module.GYamlConf.FeServers[0].Host, module.GYamlConf.FeServers[0].EditLogPort)
        utl.Log("INFO", infoMess)
        startFeCmd = fmt.Sprintf("%s/bin/start_fe.sh --daemon", feDeployDir)
        // time.Sleep(30 * time.Second)
    } else {
        infoMess = fmt.Sprintf("Starting follower FE node [host = %s, editLogPort = %d]", sshHost, editLogPort)
        utl.Log("INFO", infoMess)
        
        startFeCmd = fmt.Sprintf("%s/bin/start_fe.sh --helper %s:%d --daemon", feDeployDir, module.GFeEntryHost, module.GFeEntryEditLogPort)

        // if the start node is follower node, ALTER SYSTEM ADD FOLLOWER "host:editLogPort";
        // func RunSQL(userName string, password string, ip string, port int, dbName string, sqlStat string) (rows *sql.Rows, err error)

        sqlUserName := "root"
        sqlPassword := ""
        sqlIp := module.GFeEntryHost
        sqlPort := module.GFeEntryQueryPort
        sqlDbName := ""
        addFollowerSql := fmt.Sprintf("ALTER SYSTEM ADD FOLLOWER \"%s:%d\"", sshHost, editLogPort)
        _, err := utl.RunSQL(sqlUserName, sqlPassword, sqlIp, sqlPort, sqlDbName, addFollowerSql)
        if err != nil  {
            if strings.Contains(err.Error(), "frontend already exists name") {
            } else {
                infoMess = fmt.Sprintf("Error in add follower fe node [FeHost = %s, Error = %v", sqlIp, err)
                utl.Log("ERROR", infoMess)
                return err
            }
        }

    }


    // run feDeploy/bin/start_fe.sh --daemon --helper hsot:edit_log_port

    _, err = utl.SshRun(user, keyRsa, sshHost, sshPort, startFeCmd)
    if err != nil {
        infoMess = fmt.Sprintf("Waiting for starting FE node [FeHost = %s]", sshHost)
        utl.Log("DEBUG", infoMess)
        return err
    }
    return nil
}




