package clusterOption


import (

    "fmt"
    "time"
    "os"
    "os/user"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
    "sr-controller/cluster/importCluster"

)



func ImportCluster(clusterName string, metaFile string) {


    var infoMess string

    osUser, _ := user.Current()
    module.GSRCtlRoot = os.Getenv("SRCTLROOT")
    if module.GSRCtlRoot == "" {
        module.GSRCtlRoot = fmt.Sprintf("%s/.starrocks-controller", osUser.HomeDir)
    }


    // check cluster exist
    if !checkStatus.CheckClusterName(clusterName) {
        infoMess = fmt.Sprintf("Error in importing the cluster. The cluster exist, pls change another namne. [ClusterName = %s, MetaFile = %s]", clusterName, metaFile)
	utl.Log("ERROR", infoMess)
	os.Exit(1)
    }

    // check the metaFile exists
    _, err := os.Stat(metaFile)
    if err != nil {
        // the metafile doesn't exist
	infoMess = fmt.Sprintf("Error in importing cluster. The MetaFile doesn't exist. [ClusterName = %s, MetaFile = %s]", clusterName, metaFile)
	utl.Log("ERROR", infoMess)
	os.Exit(1)
    }

    module.InitConf(clusterName, metaFile)
    //fmt.Println(module.GYamlConf.FeServers)
    feEntryId, err := checkStatus.GetFeEntry(-1)
    module.SetFeEntry(feEntryId)
    if err != nil {
        infoMess = fmt.Sprintf("Error in get FE Entry ID when import cluter info.")
        utl.Log("ERROR", infoMess)
    }


    module.GYamlConf.ClusterInfo.User = module.GYamlConf.Global.User
    module.GYamlConf.ClusterInfo.CreateDate = time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05")
    module.GYamlConf.ClusterInfo.MetaPath = module.GWriteBackMetaPath
    module.GYamlConf.ClusterInfo.PrivateKey = module.GSshKeyRsa


    importCluster.GetFeConf()
    importCluster.GetBeConf()

    module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)

}
