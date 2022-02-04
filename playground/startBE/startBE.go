package startBE

import (
    "sr-controller/sr-utl"
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
    modFile = "/root/.starrocks-controller/playground/be/conf/be.conf"
    srcConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16"
    tarConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16\npriority_networks = 127.0.0.1"
    err := utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modifing BE config [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
	utl.Log("ERROR", infoMess)
	panic(err)
    }
    infoMess = fmt.Sprintf("Modify BE config [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("INFO", infoMess)
}


func AddBENode() {

    // add be node in fe
    addExecCMD := "mysql -uroot -P9030 -h127.0.0.1 -e 'alter system add backend \"127.0.0.1:9050\"'"
    _ = utl.RunShellScript(addExecCMD)

    time.Sleep(time.Duration(5) * time.Second)
    checkExecCMD := "mysql -uroot -P9030 -h127.0.0.1 -e 'show backends \\G'"
    res := utl.RunShellScript(checkExecCMD)
    if strings.Contains(res, "127.0.0.1") {
        fmt.Println("BE node 127.0.0.1 added successfully.")
    }

}


func RunBEProcess() {

    // mkdir /root/.starrocks-controller/playground/fe/meta
    storageDir := "/root/.starrocks-controller/playground/be/storage"
    _, err := os.Stat(storageDir)
    if err == nil {
        fmt.Printf("Detect meta folder %s exists, delete it\n", storageDir)
        err = os.RemoveAll(storageDir)
    }

    err = os.Mkdir(storageDir, 0666)
    if err != nil { panic(err) }

    // run start_fe.sh
    execCMD := "/root/.starrocks-controller/playground/be/bin/start_be.sh --daemon"
    _ = utl.RunShellScript(execCMD)

    time.Sleep(time.Duration(15) * time.Second)

}

func CheckBEStatus() {

    execCMD := "mysql -uroot -h127.0.0.1 -P9030 -e 'show backends\\G' | grep Alive"
    for i := 0; i < 5; i++ {
        res := utl.RunShellScript(execCMD)
        if strings.Contains(res, "true") {
            fmt.Println("be start successfully.")
            break
        }
        time.Sleep(time.Duration(5) * time.Second)
    }

}
