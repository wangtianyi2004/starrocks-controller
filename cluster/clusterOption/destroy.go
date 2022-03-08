package clusterOption


import(
    "sr-controller/cluster/destroyCluster"
    "sr-controller/module"
)

func Destroy(clusterName string) {


    module.InitConf(clusterName, "")

    destroyCluster.DestroyCluster(clusterName)
}

