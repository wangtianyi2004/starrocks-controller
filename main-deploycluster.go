package main

import (
    "fmt"
    "sr-controller/module"
    //"sr-controller/cluster/prepareOption"
    //"sr-controller/cluster/modifyConfig"
    //"sr-controller/cluster/startFe"
    "sr-controller/cluster/startBe"
)

func main() {

    f := "./sr-c1.yaml"
    module.InitConf(f)
    fmt.Println("############### Build by 王大可 ##############################")
    fmt.Println("############### Build by 王大可 ##############################")
    //prepareOption.CreateFeDir()
    //prepareOption.CreateBeDir()
    //prepareOption.DownloadSRPkg()
    //prepareOption.DecompressSRPkg()
    //prepareOption.DistributeFeDir()
    //prepareOption.DistributeBeDir()
    //modifyConfig.ModifyClusterConfig()
    //fmt.Println("############################################# START FE CLUSTER #############################################")
    //fmt.Println("############################################# START FE CLUSTER #############################################")
    //startFe.StartFeCluster()
    //fmt.Println("############################################# START BE CLUSTER #############################################")
    //fmt.Println("############################################# START BE CLUSTER #############################################")
    startBe.StartBeCluster()
    //startBe.TestStartBe()
}


