package prepareOption

import(
    "fmt"
    "os"
    "strings"
    "strconv"
    "sr-controller/module"
    "sr-controller/sr-utl"
)


// check dir: 
//  1. SRCTLROOT: tmp & download & log - nothing need to precheck
//  2. FE Deploy Dir
//  3. FE port
//  4. BE Deploy Dir 
//  5. BE port


func PreCheckSR() {

    var infoMess string

    feDirCheck, fePortCheck := preCheckFe()
    beDirCheck, bePortCheck := preCheckBe()

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

func preCheckFe() (string, string){

    var tmpSshHost string
    var tmpSshPort int
    var tmpDetectDir string
    var tmpDetectPort int
    var tmpUser string = module.GYamlConf.Global.User
    var tmpKeyRsa string = "/root/.ssh/id_rsa"
    var preCheckDirOutput string = ""
    var preCheckPortOutput string = ""

    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        tmpSshHost = module.GYamlConf.FeServers[i].Host
	tmpSshPort = module.GYamlConf.FeServers[i].SshPort

	// check FE deploy folder
	tmpDetectDir = module.GYamlConf.FeServers[i].DeployDir
	preCheckDirOutput =  preCheckDirOutput + dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE deploy folder")
	// check meta folder 
	tmpDetectDir = module.GYamlConf.FeServers[i].MetaDir
	preCheckDirOutput = preCheckDirOutput + dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "FE meta folder")

	// check FE HttpPort
        tmpDetectPort = module.GYamlConf.FeServers[i].HttpPort
        preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE Http Port")
        // check FE RpcPort
	tmpDetectPort = module.GYamlConf.FeServers[i].RpcPort
	preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE RPC Port")
	// check FE EditLogPort
	tmpDetectPort = module.GYamlConf.FeServers[i].EditLogPort
        preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "FE Edit Log Port")

    }

    return preCheckDirOutput, preCheckPortOutput

}


func preCheckBe() (string, string){

    var tmpSshHost string
    var tmpSshPort int
    var tmpDetectDir string
    var tmpDetectPort int
    var tmpUser string = module.GYamlConf.Global.User
    var tmpKeyRsa string = "/root/.ssh/id_rsa"
    var preCheckDirOutput string = ""
    var preCheckPortOutput string = ""

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {
        tmpSshHost = module.GYamlConf.BeServers[i].Host
        tmpSshPort = module.GYamlConf.BeServers[i].SshPort

        // check BE deploy folder
        tmpDetectDir = module.GYamlConf.BeServers[i].DeployDir
        preCheckDirOutput =  preCheckDirOutput + dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "BE deploy folder")
        // check BE storage folder
        tmpDetectDir = module.GYamlConf.BeServers[i].StorageDir
        preCheckDirOutput = preCheckDirOutput + dirExist(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectDir, "BE storage folder")

        // check BePort
        tmpDetectPort = module.GYamlConf.BeServers[i].BePort
        preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Port")
        // check BE WebServerPort
        tmpDetectPort = module.GYamlConf.BeServers[i].WebServerPort
        preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Web Server Port")
        // check HeartbeatServicePort
        tmpDetectPort = module.GYamlConf.BeServers[i].HeartbeatServicePort
        preCheckPortOutput = preCheckPortOutput + portUsed(tmpUser, tmpKeyRsa, tmpSshHost, tmpSshPort, tmpDetectPort, "BE Heartbeat Service Port")


    }

    return preCheckDirOutput, preCheckPortOutput

}




func dirExist(user string, keyRsa string, sshHost string, sshPort int, dirName string, logStr string) string {

    var infoMess string

    cmd := "ls -l " + dirName

    // SshRun(user string, keyFile string, host string, port int, command string) (outPut []byte, err error)

    output, _ := utl.SshRun(user, keyRsa, sshHost, sshPort, cmd)
    if strings.Contains(string(output), "total") {
	infoMess = fmt.Sprintf("Detect the %-20s exist [Host = %-20s, Dir = %-30s]\n", logStr, sshHost, dirName)
	utl.Log("DEBUG", infoMess)
	return infoMess
    }

    return ""
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
        return infoMess
    }

    return ""

}


func TestDirExist() {
/*
    user := "root"
    keyRsa := "/root/.ssh/id_rsa"
    sshHost := "192.168.230.41"
    sshPort := 22
    dirName := "/opt/starrocks/aaa"
    res := dirExist(user, keyRsa, sshHost, sshPort, dirName)
    fmt.Printf("dir res = %v\n", res)



    res = portUsed(user, keyRsa, sshHost, sshPort,22)
    fmt.Printf("port res = %v\n", res)
*/
    fmt.Println("Hello")
}


