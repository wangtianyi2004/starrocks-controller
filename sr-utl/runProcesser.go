package utl

import (
    "os/exec"
    "fmt"
)

// Run local shell command
func RunShellScript(scriptName string) string {
    var errmess string
    cmd := exec.Command("/bin/bash", "-c", scriptName)
    res, err := cmd.Output()
    if err != nil {
        errmess = fmt.Sprint("Error in run command [ %s ]", scriptName)
	Log("ERROR", errmess)
        panic(err)
    }
    //fmt.Println(string(res))
    return string(res)
}

// Run ssh shell command


