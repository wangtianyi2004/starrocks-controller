package clusterOption

import (

    "stargo/module"
    "stargo/cluster/prepareOption"

)


// sr-ctl-cluster deploy sr-c1 v2.0.1 /tmp/sr-c1.yaml

func TestOpt() {

    clusterName := "test-sr"
    metaFile := "sr-c1.yaml"
    module.InitConf(clusterName, metaFile)
    prepareOption.PreCheckSR()
    
}

