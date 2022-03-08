package clusterOption

import (

    "fmt"
    "os"
    "sr-controller/module"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/modifyConfig"
    "sr-controller/cluster/startCluster"
//    "sr-controller/cluster/checkStatus"

)


// sr-ctl-cluster deploy sr-c1 v2.0.1 /tmp/sr-c1.yaml

func Deploy(clusterName string, clusterVersion string, metaFile string) {



    if clusterVersion != "v2.0.1" {
        fmt.Println("Only support v2.0.1.")
        os.Exit(1)
    }


    module.InitConf(clusterName, metaFile)

    //fmt.Println("################################ 广告位招租 #################################")
    //fmt.Println("################################ 广告位招租 #################################")
    //fmt.Println("################################ 广告位招租 #################################")
    //fmt.Println("################################ 广告位招租 #################################")

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
    //checkStatus.DeploySuccess()
    
}

