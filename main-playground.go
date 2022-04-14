package main

import (
    "fmt"
    "sr-controller/playground/prepare"
    "sr-controller/playground/startFE"
    "sr-controller/playground/startBE"
)

func main() {

    preparePkg.PreCheck()
    fmt.Println("DEBUG >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Precheck end")
    preparePkg.DownloadSRPkg()
    fmt.Println("DEBUG >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> download end")
    preparePkg.DecompressSRPkg()
    fmt.Println("DEBUG >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> decompress end")
    preparePkg.DeployPkg()
    fmt.Println("DEBUG >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> deploypkg end")

    startFE.ModifyFEConfig()
    startFE.RunFEProcess()
    startFE.CheckFEStatus()

    startBE.ModifyBEConfig()
    startBE.AddBENode()
    startBE.RunBEProcess()
    startBE.CheckBEStatus()

}

