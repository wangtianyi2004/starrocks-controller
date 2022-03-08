package destroyCluster

import(
    "fmt"
    "os"
    "sr-controller/sr-utl"
    "sr-controller/module"
)

func DestroyCluster(clusterName string) {

    rmFeDir(clusterName)
    rmBeDir(clusterName)
    rmMeta(clusterName)

}


func rmFeDir(clusterName string) {

    var infoMess                string
    var tmpFeDeployDir          string
    var tmpFeMetaDir            string
    var tmpUser                 string
    var tmpKeyRsa               string
    var tmpFeHost               string
    var tmpFeSshPort            int
    var tmpRemoveDeployCmd      string
    var tmpRemoveMetaCmd     string



    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        tmpFeDeployDir = module.GYamlConf.FeServers[i].DeployDir
	tmpUser = module.GYamlConf.Global.User
	tmpKeyRsa = module.GSshKeyRsa
	tmpFeMetaDir = module.GYamlConf.FeServers[i].MetaDir
        tmpFeHost = module.GYamlConf.FeServers[i].Host
	tmpFeSshPort = module.GYamlConf.FeServers[i].SshPort
        tmpRemoveDeployCmd = "rm -rf " + tmpFeDeployDir
	tmpRemoveMetaCmd = "rm -rf " + tmpFeMetaDir

	// remove deploy dir
	infoMess = fmt.Sprintf("Waiting for remove FE deploy dir. [FeHost = %s, DeployDir = %s]", tmpFeHost, tmpRemoveDeployCmd)
	utl.Log("INFO", infoMess)
	_, _ = utl.SshRun(tmpUser, tmpKeyRsa, tmpFeHost, tmpFeSshPort, tmpRemoveDeployCmd)

	infoMess = fmt.Sprintf("Waiting for remove FE meta dir. [FeHost = %s, MetaDir = %s]", tmpFeHost, tmpRemoveMetaCmd)
	utl.Log("INFO", infoMess)
	_, _ = utl.SshRun(tmpUser, tmpKeyRsa, tmpFeHost, tmpFeSshPort, tmpRemoveMetaCmd)

        infoMess = fmt.Sprintf("Fe node removed. [FeHost = %s]", tmpFeHost)
	utl.Log("OUTPUT", infoMess)
    }
}

func rmBeDir(clusterName string) {

    var infoMess                string
    var tmpBeDeployDir          string
    var tmpBeStorageDir         string
    var tmpUser                 string
    var tmpKeyRsa               string
    var tmpBeHost               string
    var tmpBeSshPort            int
    var tmpRemoveDeployCmd      string
    var tmpRemoveStorageCmd     string


    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        tmpBeDeployDir = module.GYamlConf.BeServers[i].DeployDir
        tmpUser = module.GYamlConf.Global.User
        tmpKeyRsa = module.GSshKeyRsa
        tmpBeStorageDir = module.GYamlConf.BeServers[i].StorageDir
        tmpBeHost = module.GYamlConf.BeServers[i].Host
        tmpBeSshPort = module.GYamlConf.BeServers[i].SshPort
        tmpRemoveDeployCmd = "rm -rf " + tmpBeDeployDir
        tmpRemoveStorageCmd = "rm -rf " + tmpBeStorageDir

        // remove deploy dir
        infoMess = fmt.Sprintf("Waiting for remove BE deploy dir. [BeHost = %s, DeployDir = %s]", tmpBeHost, tmpRemoveDeployCmd)
        utl.Log("INFO", infoMess)
        _, _ = utl.SshRun(tmpUser, tmpKeyRsa, tmpBeHost, tmpBeSshPort, tmpRemoveDeployCmd)

        infoMess = fmt.Sprintf("Waiting for remove BE storage dir. [BeHost = %s, StorageDir = %s]", tmpBeHost, tmpRemoveStorageCmd)
        utl.Log("INFO", infoMess)
        _, _ = utl.SshRun(tmpUser, tmpKeyRsa, tmpBeHost, tmpBeSshPort, tmpRemoveStorageCmd)

        infoMess = fmt.Sprintf("Be node removed. [BeHost = %s]", tmpBeHost)
        utl.Log("OUTPUT", infoMess)
    }
}


func rmMeta(clusterName string) {

    var infoMess string
    var metaFileDir string

    metaFileDir = fmt.Sprintf("%s/cluster/%s", module.GSRCtlRoot, clusterName)
    err := os.RemoveAll(metaFileDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in remove the meta dir. [Dir = %s]", metaFileDir)
	utl.Log("ERROR", infoMess)
    } else {
        infoMess = fmt.Sprintf("Meta Dir removed. [Dir = %s]", metaFileDir)
	utl.Log("OUTPUT", infoMess)
    }

}
