package clusterOption


import(
    "stargo/cluster/displayCluster"
    "stargo/cluster/checkStatus"
    "stargo/module"
    "stargo/sr-utl"
    "os"
)

func Display(clusterName string) {

    var infoMess string
    module.InitConf(clusterName, "")

    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }

    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }


    clusterStatus.ClusterStat(clusterName)
}

