package clusterOption

import(
    "sr-controller/cluster/stopCluster"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
    "fmt"
    "os"
)


func Stop(clusterName string, nodeId string, role string) {

    var infoMess         string
    //var tmpNodeType      string
    var tmpNodeHost      string
    var tmpUser          string
    var tmpKeyRsa        string
    var tmpSshPort       int
    var tmpDeployDir     string
    //var tmpNodeInd       int

    module.InitConf(clusterName, "")

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa

    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }

    // stop all cluster: sr-ctl-cluster stop sr-c1 
    // stop 1 node: sr-ctl-cluster stop sr-c1 --node 192.168.88.33:9010
    // stop all FE node: sr-ctl-cluster stop sr-c1 --role FE
    // stop all BE node: sr-ctl-cluster stop sr-c1 --role BE



    //  -----------------------------------------------------------------------------------------
    //  |  case id   |  node id   |  role     |  option                                         |
    //  -----------------------------------------------------------------------------------------
    //  |  1         | null       |  null     |  stop  all cluster                              |
    //  |  2         | null       |  !null    |  stop  FE or BE cluster                         |
    //  |  3         | !null      |  null     |  stop  the FE/BE node (BE only)                 |
    //  |  4         | !null      |  !null    |  error                                          |
    //  -----------------------------------------------------------------------------------------
    if nodeId == module.NULLSTR && role == module.NULLSTR {
        // case id 1: - stop all cluster: sr-ctl-cluster stop sr-c1
        stopCluster.StopFeCluster(clusterName)
        stopCluster.StopBeCluster(clusterName)
    } // end of case 1

    if nodeId == module.NULLSTR && role != module.NULLSTR {
        // case id 2: stop FE or BE cluster
	if role == "FE" {
	    infoMess = "Stopping FE cluster ...."
	    utl.Log("INFO", infoMess)
	    stopCluster.StopFeCluster(clusterName)
	} else if role == "BE" {
            infoMess = "Stopping BE cluster ..."
            utl.Log("INFO", infoMess)
            stopCluster.StopBeCluster(clusterName)
	} else {
	    infoMess = fmt.Sprintf("Error in get Node type. Please check the nodeId. You can use 'sr-ctl-cluster display %s ' to check the node id.[NodeId = %s]", clusterName, nodeId)
	    utl.Log("ERROR", infoMess)
	}
    } // end of case 2

    if nodeId != module.NULLSTR && role == module.NULLSTR {
        // case id 3: stop the FE/BE node
	// get the node type
	tmpNodeType, i := checkStatus.GetNodeType(nodeId)
	if tmpNodeType == "FE" {
            tmpNodeHost = module.GYamlConf.FeServers[i].Host
            tmpSshPort = module.GYamlConf.FeServers[i].SshPort
            tmpDeployDir = module.GYamlConf.FeServers[i].DeployDir
	    // func StopFeNode(user string, keyRsa string, sshHost string, sshPort int, feDeployDir string) (err error) 
	    // func StopBeNode(user string, keyRsa string, sshHost string, sshPort int, beDeployDir string) (err error)
	    infoMess = fmt.Sprintf("Stopping FE node. [BeHost = %s]", tmpNodeHost)
            utl.Log("INFO", infoMess)
            stopCluster.StopFeNode(tmpUser, tmpKeyRsa, tmpNodeHost, tmpSshPort, tmpDeployDir)
	} else if tmpNodeType == "BE" {
            tmpNodeHost = module.GYamlConf.BeServers[i].Host
            tmpSshPort = module.GYamlConf.BeServers[i].SshPort
            tmpDeployDir = module.GYamlConf.BeServers[i].DeployDir
            infoMess = fmt.Sprintf("Stopping BE node. [BeHost = %s]", tmpNodeHost)
            utl.Log("INFO", infoMess)
            stopCluster.StopBeNode(tmpUser, tmpKeyRsa, tmpNodeHost, tmpSshPort, tmpDeployDir)
	} else {
	    infoMess = fmt.Sprintf("Error in get Node type. Please check the nodeId. You can use 'sr-ctl-cluster display %s ' to check the node id.[NodeId = %s]", clusterName, nodeId)
	    utl.Log("ERROR", infoMess)
	}
    }// end of case 3

    if nodeId != module.NULLSTR && role != module.NULLSTR {
        infoMess = "Detect both --node & --role option."
	utl.Log("ERROR", infoMess)
    } // end of case 4

}


/*
func getNodeType(nodeId string) (nodeType string, nodeInd int) {

    // FEID: module.GYamlConf.FeServers[i].EditLogPort, module.GYamlConf.FeServers[i].QueryPort
    // BEID: module.GYamlConf.BeServers[i].Host, module.GYamlConf.BeServers[i].BePort

    tmpNodeId := strings.Split(nodeId, ":")
    fmt.Println("DEBUG>>>>>>>>>>>>>>>>>>>>>>>>>", tmpNodeId)

    // check FE
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        if tmpNodeId[0] == module.GYamlConf.FeServers[i].Host && 
	   tmpNodeId[1] == strconv.Itoa(module.GYamlConf.FeServers[i].EditLogPort) {
            nodeType = "FE"
	    //ip = module.GYamlConf.FeServers[i].Host
	    //port = tmpNodeId[1] == module.GYamlConf.FeServers[i].EditLogPort
	    nodeInd = i
	    break
	}
    }

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        if tmpNodeId[0] == module.GYamlConf.BeServers[i].Host &&
	   tmpNodeId[1] == strconv.Itoa(module.GYamlConf.BeServers[i].BePort) {
	    nodeType = "BE"
	    //ip = module.GYamlConf.BeServers[i].Host
	    //port = module.GYamlConf.BeServers[i].BePort
	    nodeInd = i
	    break
	}
    }

    return nodeType, nodeInd

}

*/


