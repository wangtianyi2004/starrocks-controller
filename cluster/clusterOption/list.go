package clusterOption

import (
    "os"
    "os/user"
    "fmt"
    "sr-controller/cluster/listCluster"
    "sr-controller/module"
)

func List() {

    // get sr-ctl root dir
    osUser, _ := user.Current()
    module.GSRCtlRoot = os.Getenv("SRCTLROOT")
    if module.GSRCtlRoot == "" {
        module.GSRCtlRoot = fmt.Sprintf("%s/.starrocks-controller", osUser.HomeDir)
    }

    listCluster.ListCluster()
}
