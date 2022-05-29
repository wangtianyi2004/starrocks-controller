package clusterOption

import(

    "fmt"
    "os"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/upgradeCluster"
)

func Upgrade(clusterName string, clusterVersion string) {

    var infoMess           string
    //var err                error


    module.InitConf(clusterName, "")
    module.SetGlobalVar("GSRVersion", clusterVersion)    

    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }
 
    oldVersion := module.GYamlConf.ClusterInfo.Version
    newVersion := clusterVersion
    if !(oldVersion < newVersion) {
        infoMess = fmt.Sprintf("OldVersion = %s  NewVersion = %s, the NewVersion is not higher than OldVersion", oldVersion, newVersion)
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    } else {
        infoMess = fmt.Sprintf("Upgrade StarRocks Cluster %s, from version %s to version %s", clusterName, oldVersion, newVersion)
        utl.Log("OUTPUT", infoMess)
    }
    
    prepareOption.PrepareSRPkg() 
    upgradeCluster.UpgradeBeCluster()
    upgradeCluster.UpgradeFeCluster()

    module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)

}
