package utl

import (
    "os/exec"
)


func PortUsed(portStr string) bool {

    output, _ := exec.Command("/bin/bash", "-c", "netstat -nltp | grep " + portStr).CombinedOutput()

    if len(output) > 0 {
        return true
    } else {
        return false
    }
}

