package prepareOption

import(
    "fmt"
    "sr-controller/module"
    "sr-controller/sr-utl"
)


func DistributeSrDir() {

    var infoMess string

    infoMess = "Distribute FE Dir ..."
    utl.Log("OUTPUT", infoMess)
    DistributeFeDir()

    infoMess = "Distribute BE Dir ..."
    utl.Log("OUTPUT", infoMess)
    DistributeBeDir()

}

func DistributeFeDir() {

    var infoMess string
    // scp -r -P 22 -i rsaKey sourceDir root@nd1:targetDir
    // distribute FE folder
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        sshUser := module.GYamlConf.Global.User
        rsaKey := "/root/.ssh/id_rsa"
        sshPort := module.GYamlConf.FeServers[i].SshPort
        sshHost := module.GYamlConf.FeServers[i].Host

/*
        fmt.Println("sshPort: ", sshPort)
        fmt.Println("rsaKey: ", rsaKey)
        fmt.Println("user: ", sshUser)
        fmt.Println("sshHost: ", sshHost)
        fmt.Println("feSourceDir: ", feSourceDir)
        fmt.Println("feTargetDir", feTargetDir)
        fmt.Println("jdkSourceDir: ", jdkSourceDir)
        fmt.Println("jdkTargetDir", jdkTargetDir)
*/


        //utl.UploadDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string)
        // upload fe dir
        feSourceDir := fmt.Sprintf("%s/download/StarRocks-2.0.1/fe", module.GSRCtlRoot)
        feTargetDir := module.GYamlConf.FeServers[i].DeployDir
        utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, feSourceDir, feTargetDir)
        infoMess = fmt.Sprintf("Upload dir [%s] to [%s], user = %s, host = %s, port = %d, keyRsa = %s", feSourceDir, feTargetDir, sshUser, sshHost, sshPort, rsaKey)
        utl.Log("INFO", infoMess)

        // upload jdk dir
        jdkSourceDir := fmt.Sprintf("%s/download/jdk1.8.0_301", module.GSRCtlRoot)
        jdkTargetDir := fmt.Sprintf("%s/jdk", module.GYamlConf.FeServers[i].DeployDir)
        utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, jdkSourceDir, jdkTargetDir)
        infoMess = fmt.Sprintf("Upload dir [%s] to [%s], user = %s, host = %s, port = %d, keyRsa = %s", jdkSourceDir, jdkTargetDir, sshUser, sshHost, sshPort, rsaKey)
        utl.Log("INFO", infoMess)


        // modify JAVA_HOME
        startFeFilePath := fmt.Sprintf("%s/bin/start_fe.sh", module.GYamlConf.FeServers[i].DeployDir)
        jdkPath := fmt.Sprintf("%s/jdk", module.GYamlConf.FeServers[i].DeployDir)
        modifyJavaHome(sshUser, rsaKey, sshHost, sshPort, startFeFilePath, jdkPath)
        infoMess = fmt.Sprintf("Modify JAVA_HOME: sshUser = %s, rsaKey = %s, host = %s, port = %d, filePath = %s", sshUser, rsaKey, sshHost, sshPort, startFeFilePath)
        utl.Log("INFO", infoMess)

    }

}




func DistributeBeDir() {

    var infoMess string
    // scp -r -P 22 -i rsaKey sourceDir root@nd1:targetDir
    // distribute FE folder
    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

	sshUser := module.GYamlConf.Global.User
	rsaKey := "/root/.ssh/id_rsa"
	sshPort := module.GYamlConf.BeServers[i].SshPort
	sshHost := module.GYamlConf.BeServers[i].Host
	beSourceDir := fmt.Sprintf("%s/download/StarRocks-2.0.1/be", module.GSRCtlRoot)
	beTargetDir := module.GYamlConf.BeServers[i].DeployDir

/*
	fmt.Println("sshPort: ", sshPort)
	fmt.Println("rsaKey: ", rsaKey)
	fmt.Println("user: ", sshUser)
	fmt.Println("sshHost: ", sshHost)
	fmt.Println("feSourceDir: ", feSourceDir)
	fmt.Println("feTargetDir", feTargetDir)
	fmt.Println("jdkSourceDir: ", jdkSourceDir)
        fmt.Println("jdkTargetDir", jdkTargetDir)
*/
	//utl.UploadDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string)
	utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, beSourceDir, beTargetDir)
	infoMess = fmt.Sprintf("Upload dir [%s] to [%s], user = %s, host = %s, port = %d, keyRsa = %s", beSourceDir, beTargetDir, sshUser, sshHost, sshPort, rsaKey)
	utl.Log("INFO", infoMess)

    }

}


func modifyJavaHome(sshUser string, rsaKey string, host string, sshPort int, startFeFilePath string, jdkFilePath string) {

    var infoMess string
    var cmd string
    var err error

    // filePath = module.GYamlConf.FeServers[i].DeployDir
    // sed -i 's$# java$# java\nJAVA_HOME=module.GYamlConf.FeServers[i].DeployDir/fe/jdk1.8.0\n$g' filePath
    cmd = fmt.Sprintf("sed -i 's$# java$# java\\nJAVA_HOME=%s\\n$g' %s", jdkFilePath, startFeFilePath)

    _, err = utl.SshRun(sshUser, rsaKey, host, sshPort, cmd)
    if err != nil {
        infoMess = fmt.Sprintf(`Error in modify JAVA_HOME:
                       sshUser = %s
                       rsaKey = %s
                       host = %s
                       port = %d
                       cmd = %s`,
                  sshUser, rsaKey, host, sshPort, cmd)
        utl.Log("ERROR", infoMess)
        panic(err)
    }


}

