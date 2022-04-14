package clusterOption

import (

    "fmt"
    //"os"
    "sr-controller/module"
//    "sr-controller/cluster/prepareOption"
//    "sr-controller/cluster/modifyConfig"
//    "sr-controller/cluster/startCluster"

)

func ScaleOut(clusterName string, scaleMetaFile string) {

    var clusterVersion string 

    // Get the cluster version 
    module.AppendConf(clusterName)
    clusterVersion = module.GYamlConfAppend.ClusterInfo.Version

    fmt.Println("DEBUG >>>>> ", clusterVersion)
    
    module.InitConf(clusterName, scaleMetaFile)
/*    
    prepareOption.PreCheckSR()
    prepareOption.CreateDir()
    prepareOption.PrepareSRPkg()
    prepareOption.DistributeSrDir()
*/
    tmpYamlConf := module.GYamlConfAppend
    module.GYamlConfAppend = module.GYamlConf
    module.GYamlConf = tmpYamlConf

    tmpYamlConf.FeServers = append(module.GYamlConf.FeServers, module.GYamlConfAppend.FeServers[0:]...)
    tmpYamlConf.BeServers = append(module.GYamlConf.BeServers, module.GYamlConfAppend.BeServers[0:]...)
    fmt.Println("DEBUG >>>>>>>>> meta = ", module.GYamlConf.ClusterInfo.MetaPath)
    module.WriteBackMeta(tmpYamlConf, module.GYamlConf.ClusterInfo.MetaPath + "/meta.yaml")
 
    fmt.Println("DEBUG >>> GYamlConfAppend.FeServers", module.GYamlConfAppend.FeServers)
    fmt.Println("################################################")
    fmt.Println("DEBUG >>> GYamlConf.FeServers", module.GYamlConf.FeServers)
    fmt.Println("################################################")
    fmt.Println("DEBUG >>> tmpYamlConf.FeServers", tmpYamlConf.FeServers)
/*
    modifyConfig.ModifyClusterConfig()
    fmt.Println("############################################# SCALE OUT FE CLUSTER #############################################")
    fmt.Println("############################################# SCALE OUT FE CLUSTER #############################################")
    startCluster.InitFeCluster(module.GYamlConfAppend)
    fmt.Println("############################################# START BE CLUSTER #############################################")
    fmt.Println("############################################# START BE CLUSTER #############################################")
    startCluster.InitBeCluster(module.GYamlConfAppend)
*/
}
