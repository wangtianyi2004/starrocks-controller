package startBe

import(
    "fmt"
)

func TestStartBe() {

    // startBeNode(user string, keyRsa string, sshHost string, sshPort int, heartbeatServicePort int, beDeployDir string) (err error)
    user := "root"
    keyRsa := "/root/.ssh/id_rsa"
    sshHost := "192.168.230.41"
    sshPort := 22
    heartbeatServicePort := 9050
    beDeployDir := "/opt/starrocks/be"

    err := startBeNode(user, keyRsa, sshHost, sshPort, heartbeatServicePort, beDeployDir)
    if err != nil {
        fmt.Printf("ERROR >>>>>>>>>>>>>>>>>>>>>>>> %v\n", err)
    }
}
