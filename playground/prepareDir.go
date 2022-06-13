package playground

import (

    "fmt"
    "os"
    "sr-controller/cluster/prepareOption"
    "sr-controller/module"
)



func PreparePlaygroundDir() {

    // mkdir sr ctl dir
    prepareOption.CreateiSrCtlDir()
    _ = os.MkdirAll(module.GSRCtlRoot+"/playground", 0751)
    // download & decompress sr package
    prepareOption.PrepareSRPkg()
    DistributePlaygroundBinary()

}




func DistributePlaygroundBinary() {

    var sourceDir    string
    var targetDir    string
    var err          error

    // module.GDownloadPath is the folder /home/sr-dev/.starrocks-controller/download

    // deploy jdk folder
    sourceDir = module.GDownloadPath + "/jdk1.8.0_301"
    targetDir = module.GSRCtlRoot + "/playground/jdk1.8.0"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }

    // deploy fe folder
    sourceDir = fmt.Sprintf("%s/StarRocks-%s/fe", module.GDownloadPath, module.GYamlConf.ClusterInfo.Version)
    targetDir = module.GSRCtlRoot + "/playground/fe"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }

    // deploy be folder
    sourceDir = fmt.Sprintf("%s/StarRocks-%s/be", module.GDownloadPath, module.GYamlConf.ClusterInfo.Version)
    targetDir = module.GSRCtlRoot + "/playground/be"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }

}




