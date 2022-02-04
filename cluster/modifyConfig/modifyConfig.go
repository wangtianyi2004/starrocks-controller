package modifyConfig 
import (
    "fmt"
    "sr-controller/module"
    "sr-controller/sr-utl"
)

func ModifyClusterConfig() {

    var infoMess string
    var tmpConfigKey string
    var tmpConfigValue string
    var tmpUser string = module.GYamlConf.Global.User
    var tmpKeyFile string = "/root/.ssh/id_rsa"
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
	tmpConfigKey = "priority_networks"
	tmpConfigValue = module.GYamlConf.FeServers[i].PriorityNetworks
	appendConfig(tmpFeTargetConfFile, tmpConfigKey, tmpConfigValue)


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
        tmpConfigKey = "priority_networks"
        tmpConfigValue = module.GYamlConf.BeServers[i].PriorityNetworks
        appendConfig(tmpBeTargetConfFile, tmpConfigKey, tmpConfigValue)


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
    utl.Log("INFO", infoMess)

}


