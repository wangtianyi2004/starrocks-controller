package utl

import (
    "os/exec"
    "fmt"
)

// Run local shell command
func RunShellScript(scriptName string) (string, error) {
    var errmess string
    cmd := exec.Command("/bin/bash", "-c", scriptName)
    res, err := cmd.Output()
    if err != nil {
        errmess = fmt.Sprintf("Error in run command [ %s ], err = %v", scriptName, err)
	Log("ERROR", errmess)
        //panic(err)
        return "", err
    }


    return string(res), nil
}

// Run ssh shell command


