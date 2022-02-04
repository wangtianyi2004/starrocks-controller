package preparePkg

import (
    "sr-controller/sr-utl"
    "strings"
    "fmt"
    "os"
)

func PreCheck() {

    errFlg := 0

    // check mysql client exists
    execCMD := "which mysql"
    res := utl.RunShellScript(execCMD)
    if !strings.Contains(res, "mysql") {
        fmt.Println("[precheck] Detect MySQL client not exists.")
	errFlg = 1
    }

    // check StarRocks ports
    portArr := [8] string{"0.0.0.0:8060", "0.0.0.0:8040", "127.0.0.1:9010", ":::9050", ":::9020", ":::8030", ":::9060", ":::9030"}
    for _, portStr := range portArr {
	if utl.PortUsed(portStr) {
	    fmt.Println("[precheck] Detect the port " + portStr + " used.")
	    errFlg = 1
	}
    }

    // check folder exist
    playgroundFolder := "/root/.starrocks-controller/playground"
    _, err := os.Stat(playgroundFolder)
    if err == nil {
        fmt.Printf("[precheck] Detect the folder %s exists, please romove it first using [ rm -rf %s ]", playgroundFolder, playgroundFolder)
	errFlg = 1
    }

    starRocksManagerFolder := "/root/.starrocks-controller"
    _, err = os.Stat(starRocksManagerFolder)
    if err != nil {
	fmt.Println("[precheck] mkdir /root/.starrocks-controller")
        os.Mkdir(starRocksManagerFolder, 0666)
    }

    starRocksDownload := "/root/.starrocks-controller/download"
    _, err = os.Stat(starRocksDownload)
    if err != nil {
        fmt.Println("[precheck] mkdir /root/.starrocks-controller/download")
        os.Mkdir(starRocksDownload, 0666)
    }


    if errFlg == 0 {
        os.Mkdir(playgroundFolder, 0666)
        os.Mkdir(playgroundFolder + "/download", 0666)
    } else {
        os.Exit(1)
    }


}
