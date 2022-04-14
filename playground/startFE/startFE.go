package startFE

import(
    "sr-controller/sr-utl"
    "fmt"
    "os"
    "time"
    "strings"
)

func ModifyFEConfig() {

    var modFile string
    var srcConfig string
    var tarConfig string
    var infoMess string
    // modify JAVA_OPS for fe.conf
    modFile = "/root/.starrocks-controller/playground/fe/conf/fe.conf"
    srcConfig = "-Xmx8192m"
    tarConfig = "-Xmx512m"
    err := utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
	panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("INFO", infoMess)

    // modify priority_networks for fe.conf
    modFile = "/root/.starrocks-controller/playground/fe/conf/fe.conf"
    srcConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16"
    tarConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16\npriority_networks = 127.0.0.1"
    err = utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
	utl.Log("ERROR", infoMess)
	panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("INFO", infoMess)

    // modify JAVA_HOME for start_fe.sh
    modFile = "/root/.starrocks-controller/playground/fe/bin/start_fe.sh"
    srcConfig = "# java"
    tarConfig = "# java\nJAVA_HOME=/root/.starrocks-controller/playground/jdk1.8.0\n"
    err = utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("INFO", infoMess)

}

func RunFEProcess() {

    // mkdir /root/.starrocks-controller/playground/fe/meta
    metaDir := "/root/.starrocks-controller/playground/fe/meta"
    _, err := os.Stat(metaDir)
    if err == nil { 
        fmt.Printf("Detect meta folder %s exists, delete it\n", metaDir)
	err = os.RemoveAll(metaDir)
    }

    err = os.Mkdir(metaDir, 0751)
    if err != nil { panic(err) }

    // run start_fe.sh
    execCMD := "/root/.starrocks-controller/playground/fe/bin/start_fe.sh --daemon"
    _, err = utl.RunShellScript(execCMD)
    if err != nil {
        fmt.Println("Error in running cmd, cmd = %s, err = %v\n", execCMD, err)
    }

    time.Sleep(time.Duration(15) * time.Second)

}


func CheckFEStatus() {

    execCMD := "mysql -uroot -h127.0.0.1 -P9030 -e 'show frontends\\G' | grep Alive"
    for i := 0; i < 5; i++ {
	res, _:= utl.RunShellScript(execCMD)

        if strings.Contains(res, "true") {
            fmt.Println("fe start successfully.")
            break
        }
	time.Sleep(time.Duration(5) * time.Second)
    }

    fmt.Println("DEBUG >>>>>>>>>>>>>>>>> checkFEStatus end")

}
