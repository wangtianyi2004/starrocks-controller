package prepareOption

import(
    "fmt"
    "os"
    "strings"
    "strconv"
    "path"
    "regexp"
    "sr-controller/module"
    "sr-controller/sr-utl"
)


// check dir: 
//  SRCTLROOT: tmp & download & log - nothing need to precheck
//  FE ssh auth
//  FE user & sudo privileges
//  FE Deploy Dir
//  FE port
//  BE Deploy Dir 
//  BE port

const  CHECKPASS      string = "PASS"
const  CHECKFAILED    string = "FAILED"


type FePreCheckStruct struct {

    SshAuth                     string                   // check ssh auth
    UserSudo                    string                   // check user sudo privilege
    MetaDir                     string                   // check meta dir
    DeployDir                   string                   // check deploy dir
    HttpPortRes                 string                   // check http port used
    RpcPortRes                  string                   // check rpc port used
    QueryPortRes                string                   // check query port used
    EditLogPortRes              string                   // check edit log port used

}

type BePreCheckStruct struct {

    SshAuth                     string                   // check ssh auth
    UserSudo                    string                   // check user sudo privilege
    storageDir                  string                   // check storageDir
    DeployDir                   string                   // check deploy dir
    WebServerPort               string                   // check web server port
    HeartbeatServicePort        string                   // check heartbeat service port
    BrpcPort                    string                   // check brpc port 
    BePort                      string                   // check be port

}




func PreCheckSR () {

    //var preCheckRes bool
    var infoMess    string

    tmpMinus := []byte("---------------------------------------")
    fePreCheckStat := preCheckFe()
    bePreCheckStat := preCheckBe()

    infoMess = fmt.Sprintf("PreCheck FE:\n")
    infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", "IP", "ssh auth", "meta dir", "deploy dir", "http port", "rpc port", "query port", "edit log port")
    infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", tmpMinus[:20], tmpMinus[:15], tmpMinus[:25], tmpMinus[:25], tmpMinus[:15], tmpMinus[:15], tmpMinus[:15], tmpMinus[:15])
    for ip, _ := range fePreCheckStat {
        infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", ip,
	                                                                      fePreCheckStat[ip].SshAuth,
									      fePreCheckStat[ip].MetaDir,
									      fePreCheckStat[ip].DeployDir,
									      fePreCheckStat[ip].HttpPortRes,
									      fePreCheckStat[ip].RpcPortRes,
									      fePreCheckStat[ip].QueryPortRes,
									      fePreCheckStat[ip].EditLogPortRes)
    }
    infoMess = infoMess + fmt.Sprintf("\n")
    infoMess = infoMess + fmt.Sprintf("PreCheck BE:\n")

    infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", "IP", "ssh auth", "storage dir", "deploy dir", "webSer port", "heartbeat port", "brpc port", "be port")
    infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", tmpMinus[:20], tmpMinus[:15], tmpMinus[:25], tmpMinus[:25], tmpMinus[:15], tmpMinus[:15], tmpMinus[:15], tmpMinus[:15])
    for ip, _ := range bePreCheckStat {
        infoMess = infoMess + fmt.Sprintf("%-20s  %-15s  %-25s  %-25s  %-15s  %-15s  %-15s  %-15s\n", ip,
	                                                                bePreCheckStat[ip].SshAuth,
									bePreCheckStat[ip].storageDir,
									bePreCheckStat[ip].DeployDir,
									bePreCheckStat[ip].WebServerPort,
									bePreCheckStat[ip].HeartbeatServicePort,
									bePreCheckStat[ip].BrpcPort,
									bePreCheckStat[ip].BePort)
    }
    infoMess = "PRE CHECK DEPLOY ENV:\n" + infoMess + fmt.Sprintf("\n")
    utl.Log("OUTPUT", infoMess)
    //fmt.Println(fePreCheckStat)
    //fmt.Println(bePreCheckStat)
    if strings.Contains(infoMess, CHECKFAILED) {
	infoMess = "PreCheck failed."
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    } else {
        infoMess = "PreCheck successfully. RESPECT"
	utl.Log("OUTPUT", infoMess)
    }

}
/*
func PreCheckSR() {

    var infoMess string

    //feDirCheck, fePortCheck := preCheckFe()
    //beDirCheck, bePortCheck := preCheckBe()

    if feDirCheck != "" {
        feDirCheck = "precheck FE dir:\n" + feDirCheck
    }

    if fePortCheck != "" {
        fePortCheck = "precheck FE port:\n" + fePortCheck
    }

    if beDirCheck != "" {
        beDirCheck = "precheck BE dir:\n" + beDirCheck
    }

    if bePortCheck != "" {
        bePortCheck = "precheck BE port:\n" + bePortCheck
    }

    tmpOutput := feDirCheck + fePortCheck + beDirCheck + bePortCheck
    if tmpOutput == "" {
        tmpOutput = "Precheck successfully. Respect."
    }

    infoMess = fmt.Sprintf("PreCheck result: %s ", tmpOutput)
    utl.Log("INFO", infoMess)


    if tmpOutput != "Precheck successfully. Respect." {
        infoMess = "Please remove the deploy folder and kill the FE/BE process."
	os.Exit(1)
    }
}
*/



func preCheckFe() (fePreCheckRes map[string] *FePreCheckStruct) {

    var tmpSshHost               string
    var tmpSshPort               int
    var tmpDetectDir             string
    var tmpDetectPort            int
    var tmpUser                  string
    var tmpKeyRsa                string


    fePreCheckRes = make(map[string] *FePreCheckStruct)
    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa


    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        var tmpCheckRes FePreCheckStruct
        tmpSshHost = module.GYamlConf.FeServers[i].Host
	tmpSshPort = module.GYamlConf.FeServers[i].SshPort

	// check ssh auth
	tmpCheckRes.SshAuth = sshAuth(tmpSshHost, tmpSshPort)

        // check sudo privilege
        tmpCheckRes.UserSudo = sudoPriv(tmpSshHost, tmpSshPort, tmpUser)

	// check FE deploy folder
	tmpDetectDir = module.GYamlConf.FeServers[i].DeployDir
	tmpCheckRes.DeployDir = dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE deploy folder")
	//tmpCheckRes.DeployDir = dirPriv(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE deploy folder")


	// check meta folder 
	tmpDetectDir = module.GYamlConf.FeServers[i].MetaDir
	tmpCheckRes.MetaDir = dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE meta folder")
	//tmpCheckRes.MetaDir = dirPriv(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE meta folder")

	// check FE HttpPort
        tmpDetectPort = module.GYamlConf.FeServers[i].HttpPort
	tmpCheckRes.HttpPortRes = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE Http Port")

        // check FE RpcPort
	tmpDetectPort = module.GYamlConf.FeServers[i].RpcPort
        tmpCheckRes.RpcPortRes = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE RPC Port")

	// check FE EditLogPort
	tmpDetectPort = module.GYamlConf.FeServers[i].EditLogPort
	tmpCheckRes.EditLogPortRes = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE Edit Log Port")

	// check FE QueryPort
	tmpDetectPort = module.GYamlConf.FeServers[i].QueryPort
	tmpCheckRes.QueryPortRes =  portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE Query Port")

        fePreCheckRes[tmpSshHost] = &tmpCheckRes
    }

    return fePreCheckRes


}


func preCheckBe() (bePreCheckRes map[string] *BePreCheckStruct) {

    var tmpSshHost            string
    var tmpSshPort            int
    var tmpDetectDir          string
    var tmpDetectPort         int
    var tmpUser               string
    var tmpKeyRsa             string

    bePreCheckRes = make(map[string] *BePreCheckStruct)
    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa



    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        var tmpCheckRes BePreCheckStruct
        tmpSshHost = module.GYamlConf.BeServers[i].Host
        tmpSshPort = module.GYamlConf.BeServers[i].SshPort

	// check ssh auth
        tmpCheckRes.SshAuth = sshAuth(tmpSshHost, tmpSshPort)

	// Check BE deploy user exist
        tmpCheckRes.UserSudo = sudoPriv(tmpSshHost, tmpSshPort, tmpUser)

        // check BE deploy folder
        tmpDetectDir = module.GYamlConf.BeServers[i].DeployDir
        tmpCheckRes.DeployDir = dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "BE deploy folder")

        // check BE storage folder
        tmpDetectDir = module.GYamlConf.BeServers[i].StorageDir
        tmpCheckRes.storageDir = dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "BE storage folder")

        // check BePort
        tmpDetectPort = module.GYamlConf.BeServers[i].BePort
        tmpCheckRes.BePort = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Port")

        // check BE WebServerPort
        tmpDetectPort = module.GYamlConf.BeServers[i].WebServerPort
        tmpCheckRes.WebServerPort = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Web Server Port")

        // check HeartbeatServicePort
        tmpDetectPort = module.GYamlConf.BeServers[i].HeartbeatServicePort
        tmpCheckRes.HeartbeatServicePort = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Heartbeat Service Port")

        // check BE brpc port 
	tmpDetectPort = module.GYamlConf.BeServers[i].BrpcPort
        tmpCheckRes.BrpcPort = portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE brpc Port")

        bePreCheckRes[tmpSshHost] = &tmpCheckRes

    }

    return bePreCheckRes

}


func dirPriv(user string, keyRsa string, sshHost string, sshPort int, dirName string, logStr string) string {

    var infoMess    string
    var cmd         string
    var dirBase     string
    var res         string


    dirBase = path.Dir(dirName)
    // dirBase = dirName
    cmd = fmt.Sprintf("ls -al %s | grep 'd.* .$'", dirBase)
    output, _ := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)
    reg := regexp.MustCompile("\\s+")   
    dirStatArr := reg.Split(string(output), -1)
    
    //fmt.Printf("DEBUG >>>>>>> cmd = %s, output = %s, priv = %s, user = %s\n", cmd, string(output), dirStatArr[0], dirStatArr[2] )

    if strings.Contains(dirStatArr[0], "drwx") && dirStatArr[2] == user {
        infoMess = fmt.Sprintf("Detect the %-20s don't have create folder privileges [Host = %-20s, Dir = %-30s]\n", logStr, sshHost, dirName)
	utl.Log("DEBUG", infoMess)
	return CHECKPASS
    } else {
	if dirBase != "/" {
            res = dirPriv(user, keyRsa, sshHost, sshPort, dirBase, logStr)
	    if res == CHECKPASS {
	        return CHECKPASS
	    }
	} else {
	    return "Priv failed"
	}
    }

    return "Priv failed"
}

func dirExist(user string, keyRsa string, sshHost string, sshPort int, dirName string, logStr string) string {

    var infoMess     string
    var cmd          string
    //var dirBase      string
    var res          string = CHECKPASS
    var resExist     string = CHECKPASS
    var resPrivs     string = CHECKPASS

    // check dir exist
    cmd = "ls -l " + dirName
    // SshRun(user string, keyFile string, host string, port int, command string) (outPut []byte, err error)
    output, _ := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)
    if strings.Contains(string(output), "total") {
	infoMess = fmt.Sprintf("Detect the %-20s exist [Host = %-20s, Dir = %-30s]\n", logStr, sshHost, dirName)
	utl.Log("DEBUG", infoMess)
	//return CHECKFAILED
	resExist = "Dir exist"
    }
    // check dir privs
    resPrivs = dirPriv(user, keyRsa, sshHost, sshPort, dirName, logStr)

    if resPrivs == CHECKPASS && resExist == CHECKPASS {
        res = CHECKPASS
    } else if resPrivs == CHECKPASS && resExist != CHECKPASS {
	res = fmt.Sprintf("%s: %s", CHECKFAILED, resExist)
    } else if resPrivs != CHECKPASS && resExist == CHECKPASS {
        res = fmt.Sprintf("%s: %s", CHECKFAILED, resPrivs)
    } else if resPrivs != CHECKPASS && resExist != CHECKPASS {
        res = fmt.Sprintf("%s: %s/%s", CHECKFAILED, resExist, resPrivs)
    }

    return res
    //return CHECKPASS
}



func portUsed(user string, keyRsa string, sshHost string, sshPort int, detectPort int, logStr string) string{

    var infoMess string

    cmd := fmt.Sprintf("netstat -nltp | grep ':%d '", detectPort)

    // SshRun(user string, keyFile string, host string, port int, command string) (outPut []byte, err error)

    output, _ := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)
    if strings.Contains(string(output), ":" + strconv.Itoa(detectPort)) {
    //fmt.Printf("DEBUG >>>>>>>>>>>>>", ":" + strconv.Itoa(detectPort))
        infoMess = fmt.Sprintf("Detect the %s used [Host = %s, Port = %d]\n",logStr, sshHost, detectPort)
	utl.Log("DEBUG", infoMess)
        return CHECKFAILED
    }

    return CHECKPASS

}


func sudoPriv(sshHost string, sshPort int, userName string) string {

    // check FE server user exist
    var infoMess string
    var cmd      string
    // check user sudo privilege
    cmd = "sudo date"
    keyRsa := module.GSshKeyRsa
    output, _ := utl.SshRun(userName, keyRsa, sshHost, sshPort, cmd)

    if strings.Contains(string(output), "202") {
        infoMess = fmt.Sprintf("Detect user has sudo privilege. [Host = %s, User = %s]", sshHost, userName)
	utl.Log("DEBUG", infoMess)
	return CHECKPASS
    } else {
        infoMess = fmt.Sprintf("Detect user doesn't have sudo privilege. [Host = %s, User = %s]", sshHost, userName)
	utl.Log("DEBUG", infoMess)
	return CHECKFAILED
    }

    return CHECKFAILED
}



func sshAuth(sshHost string, sshPort int) string {

    // check ssh auth
    var infoMess string

    keyRsa := module.GSshKeyRsa
    sshUser := module.GYamlConf.Global.User

    output, _ := utl.SshRun(sshUser, keyRsa, sshHost, sshPort, "date")

    if strings.Contains(string(output), "202") {
        // detect the result has the year 202X, return PASS
        infoMess = fmt.Sprintf("SSH auth check successfully, [host = %s]", sshHost)
	utl.Log("DEBUG", infoMess)
	return CHECKPASS
    }

    infoMess = fmt.Sprintf("SSH auth check failed, [host = %s]", sshHost)
    utl.Log("DEBUG", infoMess)
    return CHECKFAILED
}
