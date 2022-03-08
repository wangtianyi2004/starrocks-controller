package clusterStatus

import(
    "fmt"
//    "errors"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
)

func ClusterStat(clusterName string) {

    var infoMess string
    fmt.Printf("clusterName = %s\n", clusterName)
    //metaFile := "/tmp/c1-meta.yaml"
    //module.InitConf(metaFile)

    var tmpID                  string
    var tmpRole                string
    var tmpHost                string
    var tmpPort                string
    var tmpStat                string
    var tmpDataDir             string
    var tmpDeployDir           string

    var noFeEntry              bool


    // Get FE entry
    feEntryHost, feEntryQueryPort, err := checkStatus.GetFeEntry()
    if err != nil || feEntryHost == "" || feEntryQueryPort == 0 {
        infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        utl.Log("WARN", infoMess)
        noFeEntry = true 
    } else {
        module.SetFeEntry(feEntryHost, feEntryQueryPort)
    }

    tmpMinus := []byte("------------------------------------------------------------------------------------------------------")
    fmt.Printf("%-26s  %-6s  %-20s  %-15s  %-10s  %-50s  %-50s\n", "ID", "ROLE", "HOST", "PORT", "STAT", "DATADIR", "DEPLOYDIR")
    fmt.Printf("%-26s  %-6s  %-20s  %-15s  %-10s  %-50s  %-50s\n", tmpMinus[:26], tmpMinus[:6], tmpMinus[:20], tmpMinus[:15], tmpMinus[:10], tmpMinus[:50], tmpMinus[:50])
    // Get FE status
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        tmpID = fmt.Sprintf("%s:%d", module.GYamlConf.FeServers[i].Host, module.GYamlConf.FeServers[i].EditLogPort)
        tmpRole = "FE"
        tmpHost = module.GYamlConf.FeServers[i].Host
        tmpPort = fmt.Sprintf("%d/%d", module.GYamlConf.FeServers[i].EditLogPort, module.GYamlConf.FeServers[i].QueryPort)
        tmpDataDir = module.GYamlConf.FeServers[i].MetaDir
        tmpDeployDir = module.GYamlConf.FeServers[i].DeployDir


	if !noFeEntry {

            // If we can get a FE entry(more than one FE node is running), we can use [show frontends] command by JDBC)
            // CheckFeStatus(feId int, user string, keyRsa string, sshHost string, sshPort int, feQueryPort int) (feStat FeStatusStruct, err error)
	    feStatStruct, err := checkStatus.CheckFeStatus(i)
	    if err != nil {
                infoMess = fmt.Sprintf("Error in checking FE status [FeHost = %s, error = %v]", tmpHost, err)
	        utl.Log("DEBUG", infoMess)
	    }

            if feStatStruct.FeAlive {
                tmpStat = "UP"
	    } else {
                tmpStat = "DOWN"
	    }

        } else {

            // If we cannot get a FE entry, it means no FE node is running, so we don't need to check FE status using [show frontends] command
            tmpStat = "DOWN"
	}
	// get the Dir output string
	// if the dir output string is too long, only display the tail 43 chars
	// for example, the tmpDeployDir is "/opt/starrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrocks/fe",
	// it will show the tmpDeployDir "... rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrocks/fe"

	fmt.Printf("%-26s  %-6s  %-20s  %-15s  %-10s  %-50s  %-50s\n", tmpID, tmpRole, tmpHost, tmpPort, tmpStat, tmpDeployDir, tmpDataDir)

    }

    // check BE status
    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

	tmpID = fmt.Sprintf("%s:%d", module.GYamlConf.BeServers[i].Host, module.GYamlConf.BeServers[i].BePort)
        tmpRole = "BE"
        tmpHost = module.GYamlConf.BeServers[i].Host
        tmpPort = fmt.Sprintf("%d/%d", module.GYamlConf.BeServers[i].BePort, module.GYamlConf.BeServers[i].HeartbeatServicePort)
        tmpDataDir = module.GYamlConf.BeServers[i].StorageDir
        tmpDeployDir = module.GYamlConf.BeServers[i].DeployDir


        if !noFeEntry {
            // If we can get a FE entry(more than one FE node is running), we can use [show backends] command by JDBC
            // CheckBeStatus(beId int, user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int) (beStat BeStatusStruct, err error)
            beStatStruct, err := checkStatus.CheckBeStatus(i)
	    if err != nil {
                infoMess = fmt.Sprintf("Error in checking BE status [BeHost = %s, error = %v]", tmpHost, err)
                utl.Log("DEBUG", infoMess)
	    }

	    if beStatStruct.Alive {
                tmpStat = "UP"
            } else {
                tmpStat = "DOWN"
	    }
            //fmt.Printf("id = %s\t role = %s\t host = %s\t tmpPort = %s\t tmpStat = %s\t tmpDataDir = %s\t tmpDeployDir = %s\n", tmpID, tmpRole, tmpHost, tmpPort, tmpStat, tmpDataDir, tmpDeployDir)
        } else {

            // If we cannot get a FE entry, it means no FE node is running, so we don't need to check BE status using [show frontends] command, the BE status is "WAITING FE"
            bePortRun, _ := checkStatus.CheckBePortStatus(i)
            if bePortRun {
                tmpStat = "WAITING FE"
            } else {
                tmpStat = "DOWN"
            }
	}
	fmt.Printf("%-26s  %-6s  %-20s  %-15s  %-10s  %-50s  %-50s\n", tmpID, tmpRole, tmpHost, tmpPort, tmpStat, tmpDeployDir, tmpDataDir)
    }

}
