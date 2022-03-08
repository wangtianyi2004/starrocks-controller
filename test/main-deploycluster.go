package main

import (

    "fmt"
    "sr-controller/module"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/modifyConfig"
    "sr-controller/cluster/startCluster"
//    "sr-controller/cluster/checkStatus"

)

func main() {

    metaFile := "./sr-c1.yaml"
    clusterName := "test"
    module.InitConf(clusterName, metaFile)
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

    //### recover.sh ##############################")
    modifyConfig.ModifyClusterConfig()
    fmt.Println("############################################# START FE CLUSTER #############################################")
    fmt.Println("############################################# START FE CLUSTER #############################################")
    startCluster.InitFeCluster()
    fmt.Println("############################################# START BE CLUSTER #############################################")
    fmt.Println("############################################# START BE CLUSTER #############################################")
    startCluster.InitBeCluster()
//    checkStatus.DeploySuccess()

}


