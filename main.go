package main

import (

    "fmt"
    "os"
    "flag"
    "sr-controller/sr-utl"
    "sr-controller/module"
    "sr-controller/cluster/clusterOption"
    "sr-controller/cluster/upgradeCluster"

)
func main() {

    // sr-ctl-cluster deploy    sr-c1   v2.0.1   /tmp/sr-c1.yaml
    // sr-ctl-cluster start     sr-c1
    // sr-ctl-cluster stop      sr-c1
    // sr-ctl-cluster display   sr-c1

    // sr-ctl playground v2.0.1


    var component              string
    var command                string
    var clusterName            string
    var clusterVersion         string
    var metaFile               string
    var infoMess               string
    var node                   string
    var role                   string
    var firstArgWithDash       int

    component = os.Args[1]
    //command = os.Args[2]


    switch component {

        case "playground":
            fmt.Println("Playground component is developping .......................")
        case "cluster":
	    command = os.Args[2]
            switch command {

                case "deploy":
                    clusterName = os.Args[3]
                    clusterVersion = os.Args[4]
                    metaFile = os.Args[5]
                    infoMess = fmt.Sprintf("Deploy cluster [clusterName = %s, clusterVersion = %s, metaFile = %s]\n", clusterName, clusterVersion, metaFile)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.Deploy(clusterName, clusterVersion, metaFile)

                case "start":
                    clusterName = os.Args[3]
                    infoMess = fmt.Sprintf("Start cluster [clusterName = %s]", clusterName)
                    utl.Log("OUTPUT", infoMess)
                    firstArgWithDash = 1
                    for i := 1; i < len(os.Args); i++ {
                        firstArgWithDash = i
                        if len(os.Args[i]) > 0 && os.Args[i][0] == '-' {
                            break
                        }
                    }
                    flag.StringVar(&node, "node", "", "The Node ID. Use display command to check the node id.")
                    flag.StringVar(&role, "role", "", "The start component type. You can input FE or BE.")
                    flag.CommandLine.Parse(os.Args[firstArgWithDash:])
                    //fmt.Printf("DEBUG start option: [role = %s, node = %s]\n", role, node)
                    clusterOption.Start(clusterName, node, role)

               case "stop":
                    clusterName = os.Args[3]
                    infoMess = fmt.Sprintf("Stop cluster [clusterName = %s]", clusterName)
                    utl.Log("OUTPUT", infoMess)
                    firstArgWithDash = 1
                    for i := 1; i < len(os.Args); i++ {
                        firstArgWithDash = i
                        if len(os.Args[i]) > 0 && os.Args[i][0] == '-' {
                            break
                        }
                    }
                    flag.StringVar(&node, "node", "", "The Node ID. Use display command to check the node id.")
                    flag.StringVar(&role, "role", "", "The start component type. You can input FE or BE.")
                    flag.CommandLine.Parse(os.Args[firstArgWithDash:])
                    //fmt.Printf("DEBUG stop option: [role = %s, node = %s]\n", role, node)
                    clusterOption.Stop(clusterName, node, role)

               case "display":
                    clusterName = os.Args[3]
                    infoMess = fmt.Sprintf("Display cluster [clusterName = %s]", clusterName)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.Display(clusterName)

               case "list":
                    infoMess = fmt.Sprintf("List all clusters")
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.List()

               case "destroy":
                    clusterName = os.Args[3]
                    infoMess = fmt.Sprintf("Destroy cluster. [ClusterName = %s]", clusterName)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.Destroy(clusterName)
               case "upgrade":
                    clusterName = os.Args[3]
                    clusterVersion = os.Args[4]
                    infoMess = fmt.Sprintf("Upgrade cluster. [ClusterName = %s, TargetVersion = %s]", clusterName, clusterVersion)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.Upgrade(clusterName, clusterVersion)

               case "test":
                    //clusterOption.Upgrade("sr-c1", "v2.1.3")
                    //utl.RenameDir("starrocks", "/home/sr-dev/.ssh/id_rsa", "192.168.88.83", 22, "/tmp/aaa", "/tmp/bbb")
                    //upgradeCluster.TestUpgradeBe()
                    module.InitConf("sr-c1", "")
                    module.SetGlobalVar("v2.1.3")
                    upgradeCluster.UpgradeFeCluster()
                    
               default:
                    infoMess = fmt.Sprintf("ERROR, sr-ctl-cluster don't support %s option", command)
                    utl.Log("ERROR", infoMess)
            } // end of switch command, end of case cluster

	default:
            fmt.Printf("ERROR component input.\n")
    } // end of switch component

}
