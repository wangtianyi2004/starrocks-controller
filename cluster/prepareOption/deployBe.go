package prepareOption

import(
    "fmt"
    "sr-controller/module"
    "sr-controller/sr-utl"
)


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
