package playground

import (
    "stargo/sr-utl"
    "stargo/module"
    "fmt"
    "time"
    "strings"
    "os"
)

func ModifyBEConfig() {

    var modFile string
    var srcConfig string
    var tarConfig string
    var infoMess string

    // modify priority_networks for be.conf
    modFile = module.GSRCtlRoot + "/playground/be/conf/be.conf"
    srcConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16"
    tarConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16\npriority_networks = 127.0.0.1/32"
    err := utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modifing BE config [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Modify BE config [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("DEBUG", infoMess)
}


func AddBENode() {

    var infoMess        string
    // add be node in fe
    addExecCMD := "mysql -uroot -P9030 -h127.0.0.1 -e 'alter system add backend \"127.0.0.1:9050\"'"
    _, err := utl.RunShellScript(addExecCMD)
    if err != nil{
        infoMess = fmt.Sprintf("Error in running cmd, cmd = %s, err = %v", addExecCMD, err)
	utl.Log("ERROR", infoMess)
    }

    time.Sleep(time.Duration(5) * time.Second)
    checkExecCMD := "mysql -uroot -P9030 -h127.0.0.1 -e 'show backends \\G'"
    res, err := utl.RunShellScript(checkExecCMD)
    if err != nil{
        infoMess = fmt.Sprintf("Error in running cmd.[cmd = %s, err = %v]", checkExecCMD, err)
	utl.Log("ERROR", infoMess)
    }

    if strings.Contains(res, "127.0.0.1") {
        utl.Log("OUTPUT", "BE node 127.0.0.1 added successfully.")
    }

}


func RunBEProcess() {

    var infoMess     string
    // mkdir /root/.stargo/playground/fe/meta
    storageDir := module.GSRCtlRoot + "/playground/be/storage"
    _, err := os.Stat(storageDir)
    if err == nil {
        infoMess = fmt.Sprintf("Detect meta folder %s exists, delete it\n", storageDir)
	utl.Log("ERROR", infoMess)
        err = os.RemoveAll(storageDir)
    }

    err = os.Mkdir(storageDir, 0751)
    if err != nil { panic(err) }

    // run start_fe.sh
    execCMD := module.GSRCtlRoot + "/playground/be/bin/start_be.sh --daemon"
    _, err = utl.RunShellScript(execCMD)
    if err != nil {
	infoMess = fmt.Sprintf("Error in running be process, cmd = %s, err = %v", execCMD, err)
        utl.Log("ERROR", infoMess )
    }

    time.Sleep(time.Duration(15) * time.Second)

}

func CheckBEStatus() bool {

    var res    string
    execCMD := "mysql -uroot -h127.0.0.1 -P9030 -e 'show backends\\G' | grep Alive"
    for i := 0; i < 5; i++ {
        res, _ = utl.RunShellScript(execCMD)
        if strings.Contains(res, "true") {
            utl.Log("OUTPUT", "BE start successfully.")
            return true
        }
        time.Sleep(time.Duration(5) * time.Second)
    }

    if !strings.Contains(res, "true") {
        utl.Log("ERROR", "BE start failed.")
    }

    return false


}

