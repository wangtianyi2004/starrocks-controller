package utl

import (
    "os/exec"
    "fmt"
)


func RunShellScript(scriptName string) string {
    cmd := exec.Command("/bin/bash", "-c", scriptName)
    res, err := cmd.Output()
    if err != nil {
	fmt.Println("Error in run command ", scriptName)
        panic(err)
    }
    fmt.Println(string(res))
    return string(res)
}

