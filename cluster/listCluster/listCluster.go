package listCluster

import (
    "fmt"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "io/ioutil"
)

func ListCluster() {

    // module.GSRCtlRoot
    var infoMess string
    var clusterName string
    var metaFile string
    metaPath := fmt.Sprintf("%s/cluster", module.GSRCtlRoot)

    dir, err := ioutil.ReadDir(metaPath)
    if err != nil {
        infoMess = fmt.Sprintf("Error in read dir [DirPath = %s]", metaPath)
        utl.Log("ERROR", infoMess)
    }

    tmpMinus := []byte("----------------------------------------------------------------------------------------")
    fmt.Printf("%-15s  %-10s  %-10s  %-25s  %-60s  %-50s\n", "ClusterName", "Version", "User", "CreateDate", "MetaPath", "PrivateKey")
    fmt.Printf("%-15s  %-10s  %-10s  %-25s  %-60s  %-50s\n", tmpMinus[:15], tmpMinus[:10], tmpMinus[:10], tmpMinus[:25], tmpMinus[:60], tmpMinus[:50])

    for _, info := range dir {
        clusterName = info.Name()
        metaFile = fmt.Sprintf("%s/cluster/%s/meta.yaml", module.GSRCtlRoot, clusterName)

        module.InitConf(clusterName, metaFile)
        fmt.Printf("%-15s  %-10s  %-10s  %-25s  %-60s  %-50s\n", clusterName,
                                                          module.GYamlConf.ClusterInfo.Version,
                                                          module.GYamlConf.ClusterInfo.User, 
                                                          module.GYamlConf.ClusterInfo.CreateDate, 
                                                          module.GYamlConf.ClusterInfo.MetaPath, 
                                                          module.GYamlConf.ClusterInfo.PrivateKey)
    }

}
