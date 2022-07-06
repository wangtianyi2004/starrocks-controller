package clusterOption

import (
    "fmt"
    "time"
//    "strings"
    "os"
    "stargo/sr-utl"
    "stargo/module"
    "stargo/cluster/checkStatus"
    "stargo/cluster/stopCluster"
)


func ScaleIn(clusterName string, nodeId string) {

    var feEntryId       int
    var err             error
    var tmpNodeType                   string
    var nid                           int 
    var dropCmd                       string
    var infoMess                      string
    var user                          string
    var keyRsa                        string
    var sshHost                       string
    var sshPort                       int
    var beHeartbeatServicePort        int
    var feDeployDir                   string
    var beDeployDir                   string


    module.InitConf(clusterName, "")


    if checkStatus.CheckClusterName(clusterName) {
        infoMess = "Don't find the Cluster " + clusterName 
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }

    tmpNodeType, nid = checkStatus.GetNodeType(nodeId)
    feEntryId, err = checkStatus.GetFeEntry(nid)

    if err != nil ||  feEntryId == -1 {
        //infoMess = "All FE nodes are down, please start FE node and display the cluster status again."
        //utl.Log("WARN", infoMess)
        module.SetFeEntry(0)
    } else {
        module.SetFeEntry(feEntryId)
    }


    sqlIp := module.GFeEntryHost
    sqlPort := module.GFeEntryQueryPort       
    sqlUserName := "root"
    sqlPassword := ""
    sqlDbName := ""
    user = module.GYamlConf.Global.User
    keyRsa = module.GSshKeyRsa

    if tmpNodeType == "FE" {
        
        // stop fe node first         
        sshHost = module.GYamlConf.FeServers[nid].Host
        sshPort = module.GYamlConf.FeServers[nid].SshPort
        feDeployDir = module.GYamlConf.FeServers[nid].DeployDir
        err = stopCluster.StopFeNode(user, keyRsa, sshHost, sshPort, feDeployDir)
        if err != nil {
            infoMess = fmt.Sprintf("Error in stop FE node. [nodeId = %s, error = %v]", nodeId, err)
            utl.Log("ERROR", infoMess)
            os.Exit(1)
        }

        time.Sleep(time.Duration(10) * time.Second)
        // drop BE node
        dropCmd = fmt.Sprintf("ALTER SYSTEM DROP FOLLOWER '%s'", nodeId)
        _, err := utl.RunSQL(sqlUserName, sqlPassword, sqlIp, sqlPort, sqlDbName, dropCmd)

        if err != nil {
            infoMess = fmt.Sprintf("Error in scale in FE node. [clusterName = %s, nodeId = %s, error = %s]", clusterName, nodeId, err)
            utl.Log("ERROR", infoMess) 
        }


        // remove FE dir: data, deploy, log
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.FeServers[nid].LogDir)
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.FeServers[nid].MetaDir)
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.FeServers[nid].DeployDir)


        if nid != len(module.GYamlConf.FeServers) - 1 {
            module.GYamlConf.FeServers = append(module.GYamlConf.FeServers[:nid], module.GYamlConf.FeServers[nid+1:]...)
        } else {
            module.GYamlConf.FeServers = module.GYamlConf.FeServers[:nid]
        }
        module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)

        infoMess = fmt.Sprintf("Scale in FE node successfully. [clusterName = %s, nodeId = %s]", clusterName, nodeId)
        utl.Log("OUTPUT", infoMess)

    } else if tmpNodeType == "BE" {
        // stop BE node first
        sshHost = module.GYamlConf.BeServers[nid].Host
        sshPort = module.GYamlConf.BeServers[nid].SshPort
        sshHost = module.GYamlConf.BeServers[nid].Host
        beDeployDir = module.GYamlConf.BeServers[nid].DeployDir
        beHeartbeatServicePort = module.GYamlConf.BeServers[nid].HeartbeatServicePort
        err = stopCluster.StopBeNode(user, keyRsa, sshHost, sshPort, beDeployDir)

        if err != nil {
             infoMess = fmt.Sprintf("Error in stop BE node. [nodeId = %s, error = %v]", nodeId, err)
             utl.Log("ERROR", infoMess)
             os.Exit(1)
        }

        time.Sleep(time.Duration(10) * time.Second)

        // drop BE node
        dropCmd = fmt.Sprintf("ALTER SYSTEM DROP BACKEND '%s:%d'", sshHost, beHeartbeatServicePort)

        _, err := utl.RunSQL(sqlUserName, sqlPassword, sqlIp, sqlPort, sqlDbName, dropCmd)

        if err != nil {
            infoMess = fmt.Sprintf("Error in scale in BE node. [clusterName = %s, nodeId = %s, error = %s]", clusterName, nodeId, err)
            utl.Log("ERROR", infoMess)
        }



        // remove dir: data, deploy, log
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.BeServers[nid].LogDir)
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.BeServers[nid].StorageDir)
        utl.SshRun(user, keyRsa, sshHost, sshPort, "rm -rf " + module.GYamlConf.BeServers[nid].DeployDir)

        if nid != len(module.GYamlConf.BeServers) - 1 {
            module.GYamlConf.BeServers = append(module.GYamlConf.BeServers[:nid], module.GYamlConf.BeServers[nid+1:]...)
        } else {
            module.GYamlConf.BeServers = module.GYamlConf.BeServers[:nid]
        }
        // module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)

        module.WriteBackMeta(module.GYamlConf, module.GYamlConf.ClusterInfo.MetaPath)
        infoMess = fmt.Sprintf("Scale in BE node successfully. [clusterName = %s, nodeId = %s]", clusterName, nodeId)
        utl.Log("OUTPUT", infoMess)

    } else {
        infoMess = fmt.Sprintf("Error in get Node type. Please check the nodeId. You can use 'sr-ctl-cluster display %s ' to check the node id.[NodeId = %s]", clusterName, nodeId)
        utl.Log("ERROR", infoMess)
    }

}

