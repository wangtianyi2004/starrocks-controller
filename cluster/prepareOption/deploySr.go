package prepareOption

import(
    "fmt"
    "strings"
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

    //module.WriteBackMeta(module.GYamlConf, module.GWriteBackMetaPath)

}

func DistributeFeDir() {

    var infoMess string
    // scp -r -P 22 -i rsaKey sourceDir root@nd1:targetDir
    // distribute FE folder
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        sshUser := module.GYamlConf.Global.User
        rsaKey := module.GSshKeyRsa
        sshPort := module.GYamlConf.FeServers[i].SshPort
        sshHost := module.GYamlConf.FeServers[i].Host

        //utl.UploadDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string)
        // upload fe dir
        feSourceDir := fmt.Sprintf("%s/StarRocks-%s/fe", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
        // feSourceDir := fmt.Sprintf("%s/download/StarRocks-%s/fe", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
        feTargetDir := module.GYamlConf.FeServers[i].DeployDir
        utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, feSourceDir, feTargetDir)
        infoMess = fmt.Sprintf("Upload dir feSourceDir = [%s] to feTargetDir = [%s] on FeHost = [%s]", feSourceDir, feTargetDir, sshHost)
        utl.Log("INFO", infoMess)

        // upload jdk dir
        jdkSourceDir := fmt.Sprintf("%s/jdk1.8.0_301", module.GDownloadPath)
        // jdkSourceDir := fmt.Sprintf("%s/download/jdk1.8.0_301", module.GSRCtlRoot)
        jdkTargetDir := fmt.Sprintf("%s/jdk", module.GYamlConf.FeServers[i].DeployDir)
        utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, jdkSourceDir, jdkTargetDir)
        infoMess = fmt.Sprintf("Upload dir JDKSourceDir = [%s] to JDKTargetDir = [%s] on FeHost = [%s]", jdkSourceDir, jdkTargetDir, sshHost)
        utl.Log("INFO", infoMess)


        // modify JAVA_HOME
        startFeFilePath := fmt.Sprintf("%s/bin/start_fe.sh", module.GYamlConf.FeServers[i].DeployDir)
        jdkPath := fmt.Sprintf("%s/jdk", module.GYamlConf.FeServers[i].DeployDir)
        modifyJavaHome(sshUser, rsaKey, sshHost, sshPort, startFeFilePath, jdkPath)
        infoMess = fmt.Sprintf("Modify JAVA_HOME: host = [%s], filePath = [%s]", sshHost, startFeFilePath)
        utl.Log("INFO", infoMess)

    }

}




func DistributeBeDir() {

    var infoMess string
    // scp -r -P 22 -i rsaKey sourceDir root@nd1:targetDir
    // distribute FE folder
    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

	sshUser := module.GYamlConf.Global.User
	rsaKey := module.GSshKeyRsa
	sshPort := module.GYamlConf.BeServers[i].SshPort
	sshHost := module.GYamlConf.BeServers[i].Host
        beSourceDir := fmt.Sprintf("%s/StarRocks-%s/be", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
	// beSourceDir := fmt.Sprintf("%s/download/StarRocks-%s/be", module.GSRCtlRoot, strings.Replace(module.GSRVersion, "v", "", -1))
	beTargetDir := module.GYamlConf.BeServers[i].DeployDir

	//utl.UploadDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string)
	utl.UploadDir(sshUser, rsaKey, sshHost, sshPort, beSourceDir, beTargetDir)
	infoMess = fmt.Sprintf("Upload dir BeSourceDir = [%s] to BeTargetDir = [%s] on BeHost = [%s]", beSourceDir, beTargetDir, sshHost)
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
        infoMess = fmt.Sprintf("Error in modify JAVA_HOME. [FeHost = %s, cmd = %s, Error = %v]", host, cmd, err)
        utl.Log("ERROR", infoMess)
        panic(err)
    }


}

