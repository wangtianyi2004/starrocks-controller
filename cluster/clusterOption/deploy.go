package clusterOption

import (

    "fmt"
    "os"
    "sr-controller/module"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/modifyConfig"
    "sr-controller/cluster/startCluster"

)


// sr-ctl-cluster deploy sr-c1 v2.0.1 /tmp/sr-c1.yaml

func Deploy(clusterName string, clusterVersion string, metaFile string) {


    if !(clusterVersion == "v2.0.1" || clusterVersion == "v2.1.3") {
    //if clusterVersion != "v2.0.1" || clusterVersion != "v2.1.3" {
        fmt.Println("Only support v2.0.1 & v2.1.3 version")
        os.Exit(1)
    } 
      


    module.InitConf(clusterName, metaFile)
    module.SetGlobalVar(clusterVersion)

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

