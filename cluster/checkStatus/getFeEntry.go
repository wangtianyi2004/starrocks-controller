package checkStatus

import(
    "fmt"
    "errors"
    "strings"
    "strconv"
    "sr-controller/module"
    "sr-controller/sr-utl"
)

//func GetFeEntry() (feHost string, queryPort int, err error){
func GetFeEntry() (feEntryId int, err error) {
    // feHost, queryPort, err := checkStatus.GetFeEntry()
    // get a usable FE host & query port for checking FE/BE status by [show frontends] & [show backends] command

    var infoMess string

    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        tmpSshHost := module.GYamlConf.FeServers[i].Host
	tmpSshPort := module.GYamlConf.FeServers[i].SshPort
	tmpQueryPort := module.GYamlConf.FeServers[i].QueryPort
        tmpUser := module.GYamlConf.Global.User
	tmpKeyRsa := module.GSshKeyRsa
	// check port stat by [netstat -nltp | grep 9030]
	cmd := fmt.Sprintf("netstat -nltp | grep ':%d '", tmpQueryPort)
	output, err := utl.SshRun(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, cmd)
	if err != nil {
            infoMess = fmt.Sprintf("Error in get FE entry, checking query port failed. [FeHost = %s, QueryPort = %d]", tmpSshHost, tmpQueryPort)
            utl.Log("DEBUG", infoMess)
	}

	if strings.Contains(string(output), ":" + strconv.Itoa(tmpQueryPort)) {
            infoMess = fmt.Sprintf("Get a useable FE entry. [FeID = %d, FeHost = %s, QueryPort = %d]", i, tmpSshHost, tmpQueryPort)
            utl.Log("DEBUG", infoMess)
            return i, nil
	}
    }

    err = errors.New("There is no useable FE entry.")
    return -1, err

}
