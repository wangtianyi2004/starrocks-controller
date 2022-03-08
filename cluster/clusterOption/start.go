package clusterOption

import(
    "sr-controller/cluster/startCluster"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
    "os"
    "fmt"
)


func Start(clusterName string, nodeId string, role string) {

    var infoMess         string
    //var tmpNodeType      string
    var tmpNodeHost      string
    var tmpNodePort      int
    var tmpUser          string
    var tmpKeyRsa        string
    var tmpSshPort       int
    var tmpDeployDir     string
    //var tmpNodeInd       int

    module.InitConf(clusterName, "")

    tmpUser = module.GYamlConf.Global.User
    tmpKeyRsa = module.GSshKeyRsa


    // start all cluster: sr-ctl-cluster start sr-c1 
    // start 1 node: sr-ctl-cluster start sr-c1 -N 192.168.88.33:9010
    // start all FE node: sr-ctl-cluster start sr-c1 -R FE
    // start all BE node: sr-ctl-cluster start sr-c1 -R BE

    // start all cluster: sr-ctl-cluster start sr-c1


    //  -----------------------------------------------------------------------------------------
    //  |  case id   |  node id   |  role     |  option                                         |
    //  -----------------------------------------------------------------------------------------
    //  |  1         | null       |  null     |  start all cluster                              |
    //  |  2         | null       |  !null    |  start FE or BE cluster                         |
    //  |  3         | !null      |  null     |  start the FE/BE node (BE only)                 |
    //  |  4         | !null      |  !null    |  error                                          |
    //  -----------------------------------------------------------------------------------------
    // fmt.Printf("DEBUG >>>>>>>>>>>>>>>>> nodeId = %s, role = %s\n", nodeId, role)
    if nodeId == module.NULLSTR && role == module.NULLSTR {
        // case id 1: - start all cluster: sr-ctl-cluster start sr-c1
        startCluster.StartFeCluster()
        startCluster.StartBeCluster()
    } // end of case 1

    if nodeId == module.NULLSTR && role != module.NULLSTR {
        // case id 2: start FE or BE cluster
	if role == "FE" {
	    infoMess = "Starting FE cluster ...."
	    utl.Log("INFO", infoMess)
	    startCluster.StartFeCluster()
	} else if role == "BE" {
	    startCluster.StartBeCluster()
            infoMess = "Starting BE cluster ..."
            utl.Log("INFO", infoMess)
	} else {
	    infoMess = fmt.Sprintf("Error in get Node type. Please check the nodeId. You can use 'sr-ctl-cluster display %s ' to check the node id.[NodeId = %s]", clusterName, nodeId)
	    utl.Log("ERROR", infoMess)
	}
    } // end of case 2

    if nodeId != module.NULLSTR && role == module.NULLSTR {
        // case id 3: start the FE/BE node
	// get the node type
	tmpNodeType, i := checkStatus.GetNodeType(nodeId)
	if tmpNodeType == "FE" {
	    infoMess = "Please use -R FE to start all the FE node."
	    utl.Log("ERROR", infoMess)
	    os.Exit(1)
            //tmpNodeHost = module.GYamlConf.FeServers[i].Host
            //tmpSshPort = module.GYamlConf.FeServers[i].SshPort
            //tmpNodePort = module.GYamlConf.FeServers[i].EditLogPort
            //tmpDeployDir = module.GYamlConf.FeServers[i].DeployDir
            //startCluster.StartFeNode(tmpUser, tmpKeyRsa, tmpNodeHost, tmpSshPort, tmpNodePort, tmpDeployDir)
	} else if tmpNodeType == "BE" {
            tmpNodeHost = module.GYamlConf.BeServers[i].Host
            tmpSshPort = module.GYamlConf.BeServers[i].SshPort
            tmpNodePort = module.GYamlConf.BeServers[i].HeartbeatServicePort
            tmpDeployDir = module.GYamlConf.BeServers[i].DeployDir
            infoMess = fmt.Sprintf("Start BE node. [BeHost = %s, HeartbeatServicePort = %d]", tmpNodeHost, tmpNodePort)
            utl.Log("INFO", infoMess)
            startCluster.StartBeNode(tmpUser, tmpKeyRsa, tmpNodeHost, tmpSshPort, tmpNodePort, tmpDeployDir)
	} else {
	    infoMess = fmt.Sprintf("Error in get Node type. Please check the nodeId. You can use 'sr-ctl-cluster display %s ' to check the node id.[NodeId = %s]", clusterName, nodeId)
	    utl.Log("ERROR", infoMess)
	}
    }// end of case 3

    if nodeId != module.NULLSTR && role != module.NULLSTR {
        infoMess = "Detect both -N & -R option."
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


