package startFe

import(
    "fmt"
    "golang.org/x/crypto/ssh"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
)

func TestStartFeCluster() {


    fmt.Println("[TEST] This is test func TestStartFeCluster")
    var sshHost string
    var sshPort int
    var feQueryPort int
    var editLogPort int
    var feDeployDir string
    var i int
    var feStat checkStatus.FeStatusStruct

    /*
    i = 0
    sshHost = module.GYamlConf.FeServers[i].Host
    sshPort = module.GYamlConf.FeServers[i].SshPort
    feQueryPort = module.GYamlConf.FeServers[i].QueryPort
    feStat = checkStatus.CheckFeStatus(i)
    */

    i = 1
    sshHost = module.GYamlConf.FeServers[i].Host
    sshPort = module.GYamlConf.FeServers[i].SshPort
    feQueryPort = module.GYamlConf.FeServers[i].QueryPort
    editLogPort = module.GYamlConf.FeServers[i].EditLogPort
    feDeployDir = module.GYamlConf.FeServers[i].DeployDir
    feStat, _ = checkStatus.CheckFeStatus(i)
    fmt.Printf("sshHost = %s, sshPort = %d, feQueryPort = %d, editLogPort = %d, feDeployDir = %s\n feStat = %v\n", sshHost, sshPort, feQueryPort, editLogPort, feDeployDir, feStat)
    //startFeNode(user, keyRsa, sshHost, sshPort, editLogPort, feDeployDir)
 
}

func TestSsh() {

    user := "root"
    keyFile := "/root/.ssh/id_rsa"
    host := "192.168.230.41"
    port := 22
    cmd := ""
    var output []byte
    var err error
    // func SshRun(user string, keyFile string, host string, port int, command string) (outPut []byte, err error)
/*
    cmd = "whoami"
    output, err = utl.SshRun(user, keyFile, host, port, cmd)
    fmt.Printf("cmd = [%s], output = [%s], error = [%v]\n", cmd, string(output), err)
    if err != nil {
        v, ok := err.(*ssh.ExitError)
	if ok {
	    fmt.Println("TEST ", v.Msg())
	}
    }
*/
    cmd = "ls2"
    output, err = utl.SshRun(user, keyFile, host, port, cmd)
    fmt.Printf("[TEST] cmd = [%s], output = [%s], error = [%v]\n", cmd, string(output), err)
    if err != nil {
	fmt.Println("[TEST] HERE XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if v, ok := err.(*ssh.ExitError); ok {
            fmt.Println(v.Msg())
        }
    }



}
