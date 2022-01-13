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

    // modify JAVA_OPS for fe.conf
    modFile = "/root/.starrocks-controller/playground/fe/conf/fe.conf"
    utl.ModifyConfig(modFile, "-Xmx8192m", "-Xmx512m")
    fmt.Printf("Modify JAVA_OPS for fe.con, \"-Xmx8192m\" => \"-Xmx512m\"\n")

    // modify priority_networks for fe.conf
    modFile = "/root/.starrocks-controller/playground/fe/conf/fe.conf"
    utl.ModifyConfig(modFile, "# priority_networks = 10.10.10.0/24;192.168.0.0/16", "# priority_networks = 10.10.10.0/24;192.168.0.0/16\npriority_networks = 127.0.0.1")
    fmt.Printf("Mofify priority_networks for fe.conf, append priority_networks = 127.0.0.1\n")

    // modify JAVA_HOME for start_fe.sh
    modFile = "/root/.starrocks-controller/playground/fe/bin/start_fe.sh"
    utl.ModifyConfig(modFile, "# java", "# java\nJAVA_HOME=/root/.starrocks-controller/playground/jdk1.8.0\n")
    fmt.Printf("Modify JAVA_HOME for start_fe.sh, append JAVA_HOME=/root/.starrocks-controller/playground/jdk1.8.0\n")

}

func RunFEProcess() {

    // mkdir /root/.starrocks-controller/playground/fe/meta
    metaDir := "/root/.starrocks-controller/playground/fe/meta"
    _, err := os.Stat(metaDir)
    if err == nil { 
        fmt.Printf("Detect meta folder %s exists, delete it\n", metaDir)
	err = os.RemoveAll(metaDir)
    }

    err = os.Mkdir(metaDir, 0666)
    if err != nil { panic(err) }

    // run start_fe.sh
    execCMD := "/root/.starrocks-controller/playground/fe/bin/start_fe.sh --daemon"
    _ = utl.RunShellScript(execCMD)

    time.Sleep(time.Duration(15) * time.Second)

}


func CheckFEStatus() {

    execCMD := "mysql -uroot -h127.0.0.1 -P9030 -e 'show frontends\\G' | grep Alive"
    for i := 0; i < 5; i++ {
	res := utl.RunShellScript(execCMD)
        if strings.Contains(res, "true") {
            fmt.Println("fe start successfully.")
            break
        }
	time.Sleep(time.Duration(5) * time.Second)
    }

}
