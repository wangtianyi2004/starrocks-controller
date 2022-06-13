package main

import (

    "fmt"
    "os"
    "flag"
    "sr-controller/sr-utl"
    "sr-controller/playground"
    "sr-controller/cluster/clusterOption"
    //"sr-controller/cluster/checkStatus"
    // "sr-controller/module"
    //"sr-controller/cluster/prepareOption"

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
	    playground.RunPlayground()
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

               case "downgrade":
                    clusterName = os.Args[3]
                    clusterVersion = os.Args[4]
                    infoMess = fmt.Sprintf("Downgrade cluster. [ClusterName = %s, TargetVersion = %s]", clusterName, clusterVersion)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.Downgrade(clusterName, clusterVersion)

               case "scale-out":
                    clusterName = os.Args[3]
                    //clusterVersion = os.Args[4]
                    metaFile = os.Args[4]
                    infoMess = fmt.Sprintf("Scale out cluster. [ClusterName = %s]", clusterName)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.ScaleOut(clusterName, metaFile)

               case "scale-in":
                    clusterName = os.Args[3]
                    firstArgWithDash = 1
                    for i := 1; i < len(os.Args); i++ {
                        firstArgWithDash = i
                        if len(os.Args[i]) > 0 && os.Args[i][0] == '-' {
                            break
                        }
                    }
                    flag.StringVar(&node, "node", "", "The Node ID. Use display command to check the node id.")
                    flag.CommandLine.Parse(os.Args[firstArgWithDash:])
                    infoMess = fmt.Sprintf("Scale in cluster [clusterName = %s, nodeId = %s]", clusterName, node)
                    utl.Log("OUTPUT", infoMess)
                    clusterOption.ScaleIn(clusterName, node)

               case "import":
                   clusterName = os.Args[3]
		   metaFile = os.Args[4]
                   infoMess = fmt.Sprintf("Import the cluster [clusterName = %s, metaFile = %s]", clusterName, metaFile)
		   utl.Log("OUTPUT", infoMess)
		   clusterOption.ImportCluster(clusterName, metaFile)
               case "test":
		   utl.Log("OUTPUT", "TEST >>>>>>>>>")
		   // checkStatus.TestFeStatus()
		   //prepareOption.TestPreCheck()
		   //prepareOption.PreCheckSR()
                   //playground.DeployPlayground()
               default:
                    infoMess = fmt.Sprintf("ERROR, sr-ctl-cluster don't support %s option", command)
                    utl.Log("ERROR", infoMess)
            } // end of switch command, end of case cluster
	default:
            fmt.Printf("ERROR component input.\n")
    } // end of switch component

}
