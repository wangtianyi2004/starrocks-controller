package clusterOption


import(
    "stargo/cluster/destroyCluster"
    "stargo/module"
    "stargo/sr-utl"
    "os"
    "stargo/cluster/checkStatus"
)

func Destroy(clusterName string) {

    var infoMess string
    module.InitConf(clusterName, "")


    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }

    Stop(clusterName, module.NULLSTR, module.NULLSTR)
    destroyCluster.DestroyCluster(clusterName)
}

