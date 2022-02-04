package startFe

import (
    "fmt"
    "strings"
    //"strconv"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "database/sql"
)

type FeStatusStruct struct{

    FeName              string
    FeIp                string
    FeEditLogPort       int
    FeHttpPort          int
    FeQueryPort         int
    FeRpcPort           int
    FeRole              string
    FeIsMaster          bool
    FeClusterId         int
    FeJoin              bool
    FeAlive             bool
    FeReplayedJournalId int
    FeLastHeartbeat     sql.NullString
    FeIsHelper          bool
    FeErrMsg            string
    FeStartTime         string
    FeVersion           sql.NullString

}


func CheckFeStatus(user string, keyRsa string, sshHost string, sshPort int, feQueryPort int) (feStat FeStatusStruct, err error) {

    var infoMess string
    // check port stat by [netstat -nltp | grep 9030]

    cmd := fmt.Sprintf("netstat -nltp | grep ':%d '", feQueryPort)
    output, err := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)
    /*
    if err != nil {
        infoMess = fmt.Sprintf("Error in run cmd when check fe status [user = %s, keyRsa = %s, sshHost = %s, sshPort = %d, cmd = %s, error = %v]", user, keyRsa, sshHost, sshPort, cmd, err)
	utl.Log("ERROR", infoMess)
	return false
    }
    */
    if strings.Contains(string(output), ": " + string(feQueryPort)) {
        infoMess = fmt.Sprintf("Check the fe query port %s:%d run successfully", sshHost, feQueryPort)
	utl.Log("INFO", infoMess)
    }

    // check fe status by jdbc (from the master fe node)
    //RunSQL(userName string, password string, ip string, port int, dbName string, sqlStat string)(rows *sql.Rows, err error)
    feMasterUserName := "root"
    feMasterPassword := ""
    feMasterIP := module.GYamlConf.FeServers[0].Host
    feMasterQueryPort := module.GYamlConf.FeServers[0].QueryPort
    feMasterDbName := ""
    sqlStat := "show frontends"
    rows, err := utl.RunSQL(feMasterUserName, feMasterPassword, feMasterIP, feMasterQueryPort, feMasterDbName, sqlStat)
    if err != nil{
        infoMess = fmt.Sprintf(`Error in run sql when check fe status:
	      feUserName = %s
	      fePassword = %s
	      feIP = %s
	      queryPort = %d
	      dbName = %s
	      sqlStat = %s]
	      error = %v`,
          feMasterUserName, feMasterPassword, feMasterIP, feMasterQueryPort, feMasterDbName, sqlStat, err)
        utl.Log("ERROR", infoMess)
	return feStat, err
    }

    for rows.Next(){
	err = rows.Scan(  &feStat.FeName,
                          &feStat.FeIp,
                          &feStat.FeEditLogPort,
                          &feStat.FeHttpPort,
                          &feStat.FeQueryPort,
                          &feStat.FeRpcPort,
                          &feStat.FeRole,
                          &feStat.FeIsMaster,
                          &feStat.FeClusterId,
                          &feStat.FeJoin,
                          &feStat.FeAlive,
                          &feStat.FeReplayedJournalId,
                          &feStat.FeLastHeartbeat,
                          &feStat.FeIsHelper,
                          &feStat.FeErrMsg,
                          &feStat.FeStartTime,
                          &feStat.FeVersion)
        if err != nil {
	    infoMess = fmt.Sprintf(`Error in scan sql result:
                         feUserName = %s
                         fePassword = %s
                         feIP = %s
                         queryPort = %d
                         dbName = %s
                         sqlStat = %s]
			 error = %v`,
                     feMasterUserName, feMasterPassword, feMasterIP, feMasterQueryPort, feMasterDbName, sqlStat, err)
            utl.Log("ERROR", infoMess)
	    return feStat, err
	}

        if string(feStat.FeIp) == sshHost && feStat.FeQueryPort == feQueryPort {
	    return feStat, nil
	}

    }

    return feStat, err 

}
