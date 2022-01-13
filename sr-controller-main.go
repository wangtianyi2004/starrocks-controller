package main

import (
    "fmt"
    "sr-controller/playground/prepare"
    "sr-controller/playground/startFE"
    "sr-controller/playground/startBE"
)

func main() {

    fmt.Println("This is sr-controller-main")
    preparePkg.PreCheck()
    preparePkg.DownloadSRPkg()
    preparePkg.DecompressSRPkg()
    preparePkg.DeployPkg()

    startFE.ModifyFEConfig()
    startFE.RunFEProcess()
    startFE.CheckFEStatus()

    startBE.ModifyBEConfig()
    startBE.AddBENode()
    startBE.RunBEProcess()
    startBE.CheckBEStatus()

}

