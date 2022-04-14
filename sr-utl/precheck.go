package utl

import (
    "os/exec"
    "fmt"
)


func PortUsed(portStr string) bool {

    output, _ := exec.Command("/bin/bash", "-c", "netstat -na | grep " + portStr).CombinedOutput()
    fmt.Println("DEBUG >>>>>>>>>>>>>>>>>>>>>", string(output))

    if len(output) > 0 {
        return true
    } else {
        return false
    }
}

