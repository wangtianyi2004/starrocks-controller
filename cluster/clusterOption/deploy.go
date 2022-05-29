package clusterOption

import (

    "fmt"
    "sr-controller/module"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/modifyConfig"
    "sr-controller/cluster/startCluster"

)


// sr-ctl-cluster deploy sr-c1 v2.0.1 /tmp/sr-c1.yaml

func Deploy(clusterName string, clusterVersion string, metaFile string) {


    module.InitConf(clusterName, metaFile)
    module.SetGlobalVar("GSRVersion", clusterVersion)

    prepareOption.PreCheckSR()  
    prepareOption.CreateDir()
    prepareOption.PrepareSRPkg()
    prepareOption.DistributeSrDir()
    module.WriteBackMeta(module.GYamlConf, module.GWriteBackMetaPath)

    //### recover.sh ##############################")
    modifyConfig.ModifyClusterConfig()

    fmt.Println("############################################# START FE CLUSTER #############################################")
    fmt.Println("############################################# START FE CLUSTER #############################################")

    startCluster.InitFeCluster(module.GYamlConf)
    fmt.Println("############################################# START BE CLUSTER #############################################")
    fmt.Println("############################################# START BE CLUSTER #############################################")
    startCluster.InitBeCluster(module.GYamlConf)
    //checkStatus.DeploySuccess()


    
}

