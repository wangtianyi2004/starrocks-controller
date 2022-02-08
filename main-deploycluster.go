package main

import (
    "fmt"
    "sr-controller/module"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/modifyConfig"
    "sr-controller/cluster/startFe"
    "sr-controller/cluster/startBe"
)

func main() {

    f := "./sr-c1.yaml"
    module.InitConf(f)
    fmt.Println("############################## Build by 王大可 ##############################")
    fmt.Println("################################ 广告位招租 #################################")
    fmt.Println("################################ 广告位招租 #################################")
    fmt.Println("################################ 广告位招租 #################################")
    fmt.Println("################################ 广告位招租 #################################")
    fmt.Println("############################## Build by 王大可 ##############################")
    prepareOption.PreCheckSR()
    
    prepareOption.CreateDir()
    prepareOption.PrepareSRPkg()
    prepareOption.DistributeSrDir()
    modifyConfig.ModifyClusterConfig()
    //fmt.Println("############################################# START FE CLUSTER #############################################")
    //fmt.Println("############################################# START FE CLUSTER #############################################")
    startFe.StartFeCluster()
    //fmt.Println("############################################# START BE CLUSTER #############################################")
    //fmt.Println("############################################# START BE CLUSTER #############################################")
    startBe.StartBeCluster()
    
}


