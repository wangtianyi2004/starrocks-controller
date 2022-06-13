package playground

import (
    "fmt"
    "os"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "strconv"
    "strings"
)


func precheckPortUsed() bool {

    // FE Port: 8030 9020 9030 9010
    // BE Port: 8040 8060 9050 9060 
    var portNoUsed    bool = true

    portArr := [8] string{":8060", ":8040", ":9010", ":9050", ":9020", ":8030", ":9060", ":9030"}
    for _, portStr := range portArr {
        if utl.PortUsed(portStr) {
            fmt.Println("Detect the port " + portStr + " used. Please stop it first.")
            portNoUsed = false
        }
    }

    return portNoUsed
}

func precheckOpenFiles() bool {

    var infoMess     string
    var execCMD      string
    var cmdRes       string
    var err          error

    execCMD = "ulimit -n"
    cmdRes, err = utl.RunShellScript(execCMD)
    if err != nil {
        infoMess = fmt.Sprintf("Failed to run command. [cmd = %s]", execCMD)
	utl.Log("ERROR", infoMess)
    }

    fileCount, err := strconv.Atoi(strings.Replace(cmdRes, "\n", "", -1))
    if err != nil {
        infoMess = fmt.Sprintf("Failed to convert string to int.[res = %s]", fileCount)
	utl.Log("ERROR", infoMess)
    }


    if fileCount >= 65535 {
        return true
    } else {
	infoMess = fmt.Sprintf("Error in check the open file count. Please use the command [ulimit -n] to check the openfile count and make sure more than 65535.")
	utl.Log("ERROR", infoMess)
        return false
    }
}


func playgroundDirExist() bool {

    var infoMess          string
    var playgroundDir     string

    playgroundDir = module.GSRCtlRoot + "/playground"
    _, err := os.Stat(playgroundDir)

    if err == nil {
        infoMess = fmt.Sprintf("Detect the playground dir exists. Please delete first. [playground dir = %s]", playgroundDir)
	utl.Log("ERROR", infoMess)
        return false
    }

    // dir exists
    return true

}


func PrecheckPlayground() {

    var playgroundDir     bool
    var openFileCount     bool
    var portNoUsed        bool
    // var infoMess          string


    playgroundDir = playgroundDirExist()
    openFileCount = precheckOpenFiles()
    portNoUsed = precheckPortUsed()
    if !playgroundDir || !openFileCount || !portNoUsed {
        os.Exit(1)
    }

}
