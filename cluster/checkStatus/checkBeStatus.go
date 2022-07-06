package checkStatus

import(
    "fmt"
    "strings"
    "strconv"
    "stargo/sr-utl"
    "stargo/module"
//    "database/sql"
)


// 哪个大聪明在 2.2 版本改了 show backends
// 我真是谢谢你，艹

/*
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
*/


//var GBeStatArr []BeStatusStruct

func CheckBePortStatus(beId int) (checkPortRes bool, err error) {

    var infoMess string

    tmpUser := module.GYamlConf.Global.User
    tmpKeyRsa := module.GSshKeyRsa
    tmpBeHost := module.GYamlConf.BeServers[beId].Host
    tmpSshPort := module.GYamlConf.BeServers[beId].SshPort
    tmpHeartbeatServicePort := module.GYamlConf.BeServers[beId].HeartbeatServicePort
    checkCMD := fmt.Sprintf("netstat -an | grep ':%d ' | grep -v ESTABLISHED", tmpHeartbeatServicePort)

    output, err := utl.SshRun(tmpUser, tmpKeyRsa, tmpBeHost, tmpSshPort, checkCMD)

    if err != nil {
        infoMess = fmt.Sprintf("Error in run cmd when check BE port status [BeHost = %s, error = %v]", tmpBeHost, err)
        utl.Log("DEBUG", infoMess)
        return false, err
    }

    if strings.Contains(string(output), ":" + strconv.Itoa(tmpHeartbeatServicePort)) {
        infoMess = fmt.Sprintf("Check the BE query port %s:%d run successfully", tmpBeHost, tmpHeartbeatServicePort)
        utl.Log("DEBUG", infoMess)
        return true, nil
    }

    return false, err
}

/*
func GetBeStatJDBC(beId int) (beStat BeStatusStruct, err error) {

    var infoMess string
    var tmpBeStat BeStatusStruct
    //GJdbcUser = "root"
    //GJdbcPasswd = ""
    //GJdbcDb = ""
    queryCMD := "show backends"
    tmpBeHost := module.GYamlConf.BeServers[beId].Host
    tmpHeartbeatServicePort := module.GYamlConf.BeServers[beId].HeartbeatServicePort
    rows, err := utl.RunSQL(module.GJdbcUser, module.GJdbcPasswd, module.GFeEntryHost, module.GFeEntryQueryPort, module.GJdbcDb, queryCMD)
    if err != nil{
        infoMess = fmt.Sprintf("Error in run sql when check BE status: [BeHost = %s, error = %v]", tmpBeHost, err)
        utl.Log("DEBUG", infoMess)
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
            infoMess = fmt.Sprintf("Error in scan sql result [BeHost = %s, error = %v]", tmpBeHost, err)
            utl.Log("DEBUG", infoMess)
            return beStat, err
        }

        if string(tmpBeStat.IP) == tmpBeHost && tmpBeStat.HeartbeatServicePort == tmpHeartbeatServicePort {
            beStat = tmpBeStat
            //GFeStatusArr[feId] = feStat
            return beStat, nil
        }
    }

    return beStat, err
}

*/


func GetBeStatJDBC(beId int) (beStatus map[string]string, err error) {

    var infoMess                   string
    //var tmpBeStat                  map[string]interface{}
    var queryCMD                   string
    var tmpBeHost                  string
    var tmpHeartbeatServicePort    int


    queryCMD = "show backends"
    tmpBeHost = module.GYamlConf.BeServers[beId].Host
    tmpHeartbeatServicePort = module.GYamlConf.BeServers[beId].HeartbeatServicePort


    rows, err := utl.RunSQL(module.GJdbcUser, module.GJdbcPasswd, module.GFeEntryHost, module.GFeEntryQueryPort, module.GJdbcDb, queryCMD)


    if err != nil{
        infoMess = fmt.Sprintf("Error in run sql when check BE status: [BeHost = %s, error = %v]", tmpBeHost, err)
        utl.Log("DEBUG", infoMess)
        return beStatus, err
    }


    columns, _ := rows.Columns()
    columnLength := len(columns)
    cache := make([]interface{}, columnLength)

    for index, _ := range cache {
        var tmpVal interface{}
        cache[index] = &tmpVal
    }


    for rows.Next(){
        err = rows.Scan(cache...)

        if err != nil {
            infoMess = fmt.Sprintf("Error in scan sql result [BeHost = %s, error = %v]", tmpBeHost, err)
            utl.Log("DEBUG", infoMess)
            return beStatus, err
        }

        /*
        if string(tmpBeStat.IP) == tmpBeHost && tmpBeStat.HeartbeatServicePort == tmpHeartbeatServicePort {
            beStat = tmpBeStat
            //GFeStatusArr[feId] = feStat
            return beStat, nil
        }
        */

        beStatus = make(map[string]string)
        for i, data := range cache {
            beStatus[columns[i]] = fmt.Sprintf("%s", *data.(*interface{}))
        }


	hertbeatPort, _ := strconv.Atoi(beStatus["HeartbeatPort"])
	if beStatus["IP"] == tmpBeHost && hertbeatPort == tmpHeartbeatServicePort {
            return beStatus, err
	}
	//statList = append(statList, item)
    }




    return beStatus, err

}


func CheckBeStatus(beId int) (beStat map[string]string, err error) {

    var bePortRun   bool
    bePortRun, err = CheckBePortStatus(beId)

    if bePortRun {
        beStat, err = GetBeStatJDBC(beId)
    }

    return beStat, err
}


func TestBeStatus() {

    module.InitConf("sr-c1", "")
    feEntryId, _ := GetFeEntry(-1)
    module.SetFeEntry(feEntryId)

    aaa, _ := CheckBeStatus(0)
    fmt.Println(aaa)
}
