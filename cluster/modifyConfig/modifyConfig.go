package modifyConfig 
import (
    "fmt"
    "strconv"
    "sr-controller/module"
    "sr-controller/sr-utl"
)




func ModifyTest() {

    /*
    var configMap map[string] string
    configMap = make(map[string] string)

    i := 0
    configMap["priority_networks"] = module.GYamlConf.FeServers[i].PriorityNetworks
    configMap["http_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].HttpPort)
    configMap["rpc_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].RpcPort)
    configMap["edit_log_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].EditLogPort)
    fmt.Printf("%s:%d\n", "http_port", module.GYamlConf.FeServers[i].HttpPort)
    fmt.Printf("%s:%d\n", "rpc_port", module.GYamlConf.FeServers[i].RpcPort)
    fmt.Println("[TEST] #########################################################################")
    for key, val := range configMap {
	    fmt.Printf("key:%s\tvalue:%s\n", key, val)
    }
    */
    i := 0
    for k, v := range module.GYamlConf.FeServers[i].Config {
        fmt.Printf("%s:%s\n", k, v)
    }
    //fmt.Println("[TEST]", module.GYamlConf.FeServers[i].Config)

}


func ModifyClusterConfig() {

    var infoMess string
    //var tmpConfigKey string
    //var tmpConfigValue string
    var tmpUser string = module.GYamlConf.Global.User
    var tmpKeyFile string = module.GSshKeyRsa
    var configMap map[string] string

    infoMess = "Modify configuration for FE nodes & BE nodes ..."
    utl.Log("OUTPUT", infoMess)
    // modify FE config
    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        // copy fe config file
	tmpFeHost := module.GYamlConf.FeServers[i].Host
	tmpFeQueryPort := module.GYamlConf.FeServers[i].QueryPort
	tmpFeSourceConfFile := fmt.Sprintf("%s/download/StarRocks-2.0.1/fe/conf/fe.conf", module.GSRCtlRoot)
	tmpFeTargetConfFile := fmt.Sprintf("%s/tmp/fe.conf-%s-%d", module.GSRCtlRoot, tmpFeHost, tmpFeQueryPort)

	err := copyConfigFile(tmpFeSourceConfFile, tmpFeTargetConfFile)
	if err != nil {
	    infoMess = fmt.Sprintf("Error in modifing fe cluster configuration. Copy configuration file failed [sourceFile = %s, targetFile = %s]", tmpFeSourceConfFile, tmpFeTargetConfFile)
	    utl.Log("ERROR", infoMess)
	    panic(err)
	}

	// append new config into tmp configuration file

	// add network priority config
	//tmpConfigKey = "priority_networks"
	//tmpConfigValue = module.GYamlConf.FeServers[i].PriorityNetworks
	//appendConfig(tmpFeTargetConfFile, tmpConfigKey, tmpConfigValue)

        configMap = make(map[string] string)
        configMap["priority_networks"] = module.GYamlConf.FeServers[i].PriorityNetworks
        configMap["http_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].HttpPort)
        configMap["rpc_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].RpcPort)
        configMap["edit_log_port"] = strconv.Itoa(module.GYamlConf.FeServers[i].EditLogPort)
        for k, v := range configMap {
	    if v != "0" {
                appendConfig(tmpFeTargetConfFile, k, v)
            }
	}

        for k, v := range module.GYamlConf.FeServers[i].Config {
            appendConfig(tmpFeTargetConfFile, k, v)
        }


        // distribute tmp fe configuration file
	tmpUser := module.GYamlConf.Global.User
	tmpFeSshPort := module.GYamlConf.FeServers[i].SshPort
        tmpTargetFeConfPath := module.GYamlConf.FeServers[i].DeployDir + "/conf/fe.conf"
        utl.UploadFile(tmpUser, tmpKeyFile, tmpFeHost, tmpFeSshPort, tmpFeTargetConfFile, tmpTargetFeConfPath)


    }

    // modify BE config
    for i := 0; i < len(module.GYamlConf.BeServers); i++ {
        // copy BE config file
        tmpBeHost := module.GYamlConf.BeServers[i].Host
        tmpBeHeartbeatServicePort := module.GYamlConf.BeServers[i].HeartbeatServicePort
	tmpBeSourceConfFile := fmt.Sprintf("%s/download/StarRocks-2.0.1/be/conf/be.conf", module.GSRCtlRoot)
        tmpBeTargetConfFile := fmt.Sprintf("%s/tmp/be.conf-%s-%d", module.GSRCtlRoot, tmpBeHost, tmpBeHeartbeatServicePort)

        err := copyConfigFile(tmpBeSourceConfFile, tmpBeTargetConfFile)
        if err != nil {
            infoMess = fmt.Sprintf("Error in modifing BE cluster configuration. Copy configuration file failed [sourceFile = %s, targetFile = %s]", tmpBeSourceConfFile, tmpBeTargetConfFile)
            utl.Log("ERROR", infoMess)
            panic(err)
        }

        // append new config into tmp configuration file
        // add network priority config
        //tmpConfigKey = "priority_networks"
        //tmpConfigValue = module.GYamlConf.BeServers[i].PriorityNetworks
        //appendConfig(tmpBeTargetConfFile, tmpConfigKey, tmpConfigValue)
        configMap = make(map[string] string)
        configMap["priority_networks"] = module.GYamlConf.BeServers[i].PriorityNetworks
        configMap["be_port"] = strconv.Itoa(module.GYamlConf.BeServers[i].BePort)
        configMap["webserver_port"] = strconv.Itoa(module.GYamlConf.BeServers[i].WebServerPort)
        configMap["heartbeat_service_port"] = strconv.Itoa(module.GYamlConf.BeServers[i].HeartbeatServicePort)
	configMap["brpc_port"] = strconv.Itoa(module.GYamlConf.BeServers[i].BrpcPort)
        for k, v := range configMap {
            if v != "0" {
                appendConfig(tmpBeTargetConfFile, k, v)
            }
        }

        for k, v := range module.GYamlConf.BeServers[i].Config {
            appendConfig(tmpBeTargetConfFile, k, v)
        }


        // distribute tmp fe configuration file
        tmpBeSshPort := module.GYamlConf.BeServers[i].SshPort
        tmpTargetBeConfPath := module.GYamlConf.BeServers[i].DeployDir + "/conf/be.conf"
        utl.UploadFile(tmpUser, tmpKeyFile, tmpBeHost, tmpBeSshPort, tmpBeTargetConfFile, tmpTargetBeConfPath)

    }


}

func copyConfigFile(sourceFile string, targetFile string) (err error){

    var infoMess string

    fileByte, err := utl.CopyFile(sourceFile, targetFile)
    if err != nil || fileByte == 0 {
        infoMess = fmt.Sprintf("Error in copy fe config file, [sourceFile = %s, targetFile = %s, fileByte = %d, error = %v]", sourceFile, targetFile, fileByte, err)
	utl.Log("ERROR", infoMess)
	return err
    }

    return nil

}

func appendConfig(configFile string, configKey string, configValue string) {

    var infoMess string

    err := utl.AppendConfig(configFile, configKey, configValue)
    if err != nil {
        infoMess = fmt.Sprintf("Error in append new configuration to tmp config file [configFile = %s, configKey = %s, configValue = %s]", configFile, configKey, configValue)
        utl.Log("ERROR", infoMess)
        panic(err)
    }
    infoMess = fmt.Sprintf("Append new configuration to tmp config file [configFile = %s, configKey = %s, configValue = %s]", configFile, configKey, configValue)
    utl.Log("DEBUG", infoMess)

}


