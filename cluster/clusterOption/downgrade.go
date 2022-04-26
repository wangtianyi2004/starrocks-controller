package clusterOption

import(

    "fmt"
    "os"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/prepareOption"
    "sr-controller/cluster/downgradeCluster"
)

func Downgrade(clusterName string, clusterVersion string) {

    var infoMess           string
    //var err                error
    if !(clusterVersion == "v2.0.1" || clusterVersion == "v2.1.3") {
    //if clusterVersion != "v2.0.1" || clusterVersion != "v2.1.3" {
        fmt.Println("Only support v2.0.1 & v2.1.3 version")
        os.Exit(1)
    } 




    module.InitConf(clusterName, "")
    module.SetGlobalVar(clusterVersion)    
 
    oldVersion := module.GYamlConf.ClusterInfo.Version
    newVersion := clusterVersion
    if !(oldVersion > newVersion) {
        infoMess = fmt.Sprintf("OldVersion = %s  NewVersion = %s, the NewVersion is not higher than OldVersion", oldVersion, newVersion)
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    } else {
        infoMess = fmt.Sprintf("Downgrade StarRocks Cluster %s, from version %s to version %s", clusterName, oldVersion, newVersion)
        utl.Log("OUTPUT", infoMess)
    }
    
    fmt.Println("DEBUG >>>>>>>>>>> metafile", module.GYamlConf.ClusterInfo.MetaPath + "/meta.yaml")
    prepareOption.PrepareSRPkg() 
    downgradeCluster.DowngradeBeCluster()
    downgradeCluster.DowngradeFeCluster()

    module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)

}
