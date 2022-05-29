package utl

import (
    "os/exec"
)


func PortUsed(portStr string) bool {

    output, _ := exec.Command("/bin/bash", "-c", "netstat -na | grep " + portStr + " | grep -v ESTABLISHED").CombinedOutput()

    if len(output) > 0 {
        return true
    } else {
        return false
    }
}

