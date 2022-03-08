package main

import(
    "sr-controller/cluster/startCluster"
    "sr-controller/module"
)

func main() {

    clusterName := "test"
    metaFile := "/tmp/c1-meta.yaml"
    module.InitConf(clusterName, metaFile)


    startCluster.StartFeCluster()
    startCluster.StartBeCluster()

}
