package checkStatus

import(
    "fmt"
    "strings"
    "strconv"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "database/sql"
)

type BeStatusStruct struct{

    BackendId                int
    Cluster                  string
    IP                       string
    HeartbeatServicePort     int
    BePort                   int
    HttpPort                 int
    BrpcPort                 int
    LastStartTime            sql.NullString
    LastHeartbeat            sql.NullString
    Alive                    bool
    SystemDecommissioned     bool
    ClusterDecommissioned    bool
    TabletNum                int
    DataUsedCapacity         string
    AvailCapacity            string
    TotalCapacity            sql.NullString
    UsedPct                  string
    MaxDiskUsedPct           string
    ErrMsg                   sql.NullString
    Version                  sql.NullString
    Status                   sql.NullString
    DataTotalCapacity        sql.NullString
    DataUsedPct              sql.NullString

}


var GBeStatArr []BeStatusStruct

func CheckBeStatus(beId int, user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int) (beStat BeStatusStruct, err error) {

    var infoMess string
    var tmpBeStat BeStatusStruct
    // check port stat by [netstat -nltp | grep 9050]
    cmd := fmt.Sprintf("netstat -nltp | grep ':%d '", heartbeatServicePort)
    output, err := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)

    if strings.Contains(string(output), ":" + strconv.Itoa(heartbeatServicePort)) {
        infoMess = fmt.Sprintf("Check the be heartbeat service port %s:%d run successfully", sshHost, heartbeatServicePort)
        utl.Log("INFO", infoMess)
    }

    // check be status by jdbc (from the master fe node)
    //RunSQL(userName string, password string, ip string, port int, dbName string, sqlStat string)(rows *sql.Rows, err error)
    feMasterUserName := "root"
    feMasterPassword := ""
    feMasterIP := module.GYamlConf.FeServers[0].Host
    feMasterQueryPort := module.GYamlConf.FeServers[0].QueryPort
    feMasterDbName := ""
    sqlStat := "show backends"
    rows, err := utl.RunSQL(feMasterUserName, feMasterPassword, feMasterIP, feMasterQueryPort, feMasterDbName, sqlStat)

    if err != nil{
        infoMess = fmt.Sprintf(`Error in run sql when check be status:
              feUserName = %s
              fePassword = %s
              feIP = %s
              queryPort = %d
              dbName = %s
              sqlStat = %s]
              error = %v`, feMasterUserName, feMasterPassword, feMasterIP, feMasterQueryPort, feMasterDbName, sqlStat, err)
        utl.Log("ERROR", infoMess)
        return beStat, err
    }

    for rows.Next(){
        err = rows.Scan(  &tmpBeStat.BackendId,
                          &tmpBeStat.Cluster,
                          &tmpBeStat.IP,
                          &tmpBeStat.HeartbeatServicePort,
                          &tmpBeStat.BePort,
                          &tmpBeStat.HttpPort,
                          &tmpBeStat.BrpcPort,
                          &tmpBeStat.LastStartTime,
                          &tmpBeStat.LastHeartbeat,
                          &tmpBeStat.Alive,
                          &tmpBeStat.SystemDecommissioned,
                          &tmpBeStat.ClusterDecommissioned,
                          &tmpBeStat.TabletNum,
                          &tmpBeStat.DataUsedCapacity,
                          &tmpBeStat.AvailCapacity,
                          &tmpBeStat.TotalCapacity,
                          &tmpBeStat.UsedPct,
                          &tmpBeStat.MaxDiskUsedPct,
                          &tmpBeStat.ErrMsg,
                          &tmpBeStat.Version,
                          &tmpBeStat.Status,
                          &tmpBeStat.DataTotalCapacity,
                          &tmpBeStat.DataUsedPct)
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
            return beStat, err
        }

        if beStat.IP == sshHost && beStat.HeartbeatServicePort == heartbeatServicePort {
            beStat = tmpBeStat
            GBeStatArr[beId] = beStat
            return beStat, nil
        }
    }
    return beStat,nil
}

