package clusterOption


import(
    "sr-controller/cluster/displayCluster"
    "sr-controller/module"
)

func Display(clusterName string) {


    module.InitConf(clusterName, "")
    clusterStatus.ClusterStat(clusterName)
}

