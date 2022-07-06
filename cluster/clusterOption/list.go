package clusterOption

import (
    "os"
    "os/user"
    "fmt"
    "stargo/cluster/listCluster"
    "stargo/module"
)

func List() {

    // get sr-ctl root dir
    osUser, _ := user.Current()
    module.GSRCtlRoot = os.Getenv("SRCTLROOT")
    if module.GSRCtlRoot == "" {
        module.GSRCtlRoot = fmt.Sprintf("%s/.stargo", osUser.HomeDir)
    }

    listCluster.ListCluster()
}
