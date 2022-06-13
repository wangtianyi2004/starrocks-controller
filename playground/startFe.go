package playground

import (

    "sr-controller/sr-utl"
    "sr-controller/module"
    "fmt"
    "os"
    "time"
    "strings"

)


func ModifyFEConfig() {

    var modFile           string
    var srcConfig         string
    var tarConfig         string
    var infoMess          string

    // modify JAVA_OPS for fe.conf
    modFile = module.GSRCtlRoot + "/playground/fe/conf/fe.conf"
    srcConfig = "-Xmx8192m"
    tarConfig = "-Xmx512m"
    err := utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("DEBUG", infoMess)

    // modify priority_networks for fe.conf
    modFile = module.GSRCtlRoot + "/playground/fe/conf/fe.conf"
    srcConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16"
    tarConfig = "# priority_networks = 10.10.10.0/24;192.168.0.0/16\npriority_networks = 127.0.0.1/32"
    err = utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("DEBUG", infoMess)

    // modify JAVA_HOME for start_fe.sh
    modFile = module.GSRCtlRoot + "/playground/fe/bin/start_fe.sh"
    srcConfig = "# java"
    tarConfig = fmt.Sprintf("# java\nJAVA_HOME=%s/playground/jdk1.8.0\n", module.GSRCtlRoot)
    err = utl.ModifyConfig(modFile, srcConfig, tarConfig)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Modify FE configuration [modFile = %s, srcConfig = %s, tarConfig = %s]", modFile, srcConfig, tarConfig)
    utl.Log("DEBUG", infoMess)

}




func RunFEProcess() {

    var infoMess string
    // mkdir /root/.starrocks-controller/playground/fe/meta
    metaDir := module.GSRCtlRoot + "/playground/fe/meta"
    _, err := os.Stat(metaDir)
    if err == nil {
        fmt.Printf("Detect meta folder %s exists, delete it\n", metaDir)
        err = os.RemoveAll(metaDir)
    }

    err = os.Mkdir(metaDir, 0751)
    if err != nil { panic(err) }

    // run start_fe.sh
    execCMD := module.GSRCtlRoot + "/playground/fe/bin/start_fe.sh --daemon"
    _, err = utl.RunShellScript(execCMD)
    if err != nil {
        infoMess = fmt.Sprintf("Error in running cmd, cmd = %s, err = %v\n", execCMD, err)
	utl.Log("ERROR", infoMess)

    }

    time.Sleep(time.Duration(15) * time.Second)

}


func CheckFEStatus() bool {

    execCMD := "mysql -uroot -h127.0.0.1 -P9030 -e 'show frontends\\G' | grep Alive"
    for i := 0; i < 5; i++ {
        res, _:= utl.RunShellScript(execCMD)

        if strings.Contains(res, "true") {
            utl.Log("OUTPUT", "fe start successfully.")
            return true
        }
        time.Sleep(time.Duration(5) * time.Second)
    }

    return false

}


