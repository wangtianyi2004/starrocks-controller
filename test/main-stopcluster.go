package main

import(
    "sr-controller/cluster/stopCluster"
    "sr-controller/module"
)

func main() {

    metaFile := "/tmp/c1-meta.yaml"
    clusterName := "test"
    module.InitConf(clusterName, metaFile)
 
    stopCluster.StopFeCluster(clusterName)
    stopCluster.StopBeCluster(clusterName)
}
