package module


import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "fmt"
    "os"
    "os/user"
    "time"
    "strings"
    "sr-controller/sr-utl"
)



const NULLSTR = ""

var GClusterName           string
var GYamlConf              *ConfStruct
var GYamlConfAppend        *ConfStruct
var GSshKeyRsa             string
var GSRCtlRoot             string
var GWriteBackMetaPath     string
var GJdbcUser              string
var GJdbcPasswd            string
var GJdbcDb                string
var GFeEntryHost           string
var GFeEntryPort           int


type ConfStruct struct {

    ClusterInfo struct {
        User                  string                 `yaml:"user"`
	CreateDate            string                 `yaml:"create_date"`
	MetaPath              string                 `yaml:"meta_path"`
        PrivateKey            string                 `yaml:"private_key"`
    }

    Global struct {
        User                  string                 `yaml:"user"`
        SshPort               int                    `yaml:"ssh_port"`
    } `yaml:"global"`

    ServerConfig struct {
        Fe                    map[string]string      `yaml:"fe"`
        Be                    map[string]string      `yaml:"be"`
    } `yaml:"server_configs"`

    FeServers []struct {
        Host                  string                 `yaml:"host"`
        SshPort               int                    `yaml:"ssh_port"`
        HttpPort              int                    `yaml:"http_port"`
        RpcPort               int                    `yaml:"rpc_port"`
        QueryPort             int                    `yaml:"query_port"`
        EditLogPort           int                    `yaml:"edit_log_port"`
        DeployDir             string                 `yaml:"deploy_dir"`
        MetaDir               string                 `yaml:"meta_dir"`
        LogDir                string                 `yaml:"log_dir"`
        PriorityNetworks      string                 `yaml:"priority_networks"`
        Config                map[string]string      `yaml:"config"`
    }  `yaml:"fe_servers"`

    BeServers []struct {
        Host                 string                   `yaml:"host"`
        SshPort              int                      `yaml:"ssh_port"`
        BePort               int                      `yaml:"be_port"`
        WebServerPort        int                      `yaml:"web_server_port"`
        HeartbeatServicePort int                      `yaml:"heartbeat_service_port"`
        BrpcPort             int                      `yaml:"brpc_port"`
        DeployDir            string                   `yaml:"deploy_dir"`
        StorageDir           string                   `yaml:"storage_dir"`
        LogDir               string                   `yaml:"log_dir"`
	PriorityNetworks     string                   `yaml:"priority_networks"`
        Config               map[string]string        `yaml:"configs`
    } `yaml:"be_servers"`

    PrometheusServer struct {
        Host                  string                  `yaml:"host"`
        SshPort               int                     `yaml:"ssh_port"`
        HttpPort              int                     `yaml:"http_port"`
        DeployDir             string                  `yaml:"deploy_dir"`
        DataDir               string                  `yaml:"data_dir"`
        LogDir                string                  `yaml:"log_dir"`
    } `yaml:"prometheus_servers"`

    GrafanaServer struct {
        Host                  string                  `yaml:"host"`
        SshPort               int                     `yaml:"ssh_port"`
        HttpPort              int                     `yaml:"http_port"`
        DeployDir             string                  `yaml:"deploy_dir"`
    } `yaml:"grafana_servers"`

    AlertManagerServer struct {
        Host                  string                   `yaml:"host"`
        SshPort               int                      `yaml:"ssh_port"`
        WebPort               int                      `yaml:"web_port"`
        ClusterPort           int                      `yaml:"cluster_port"`
        DeployDir             string                   `yaml:"deploy_dir"`
        DataDir               string                   `yaml:"data_dir"`
        LogDir                string                   `yaml:"log_dir"`
    } `yaml:"alertmanager_servers"`
}


func (cc *ConfStruct) GetConf(fileName string) *ConfStruct {

    yamlFile, err := ioutil.ReadFile(fileName)
    if err != nil { panic(err) }

    err = yaml.Unmarshal(yamlFile, cc)
    if err != nil { panic(err) }

    return cc
}




func InitConf(clusterName string, fileName string) {


    var confS ConfStruct

    // get home dir & ssh auth key
    osUser, _ := user.Current()
    GSshKeyRsa = fmt.Sprintf("%s/.ssh/id_rsa", osUser.HomeDir)

    // get sr-ctl root dir
    GSRCtlRoot = os.Getenv("SRCTLROOT")
    if GSRCtlRoot == "" {
        GSRCtlRoot = fmt.Sprintf("%s/.starrocks-controller", osUser.HomeDir)
    }

    // get the write back meta path
    GClusterName = clusterName
    GWriteBackMetaPath = fmt.Sprintf("%s/cluster/%s", GSRCtlRoot, GClusterName)

    // get the FE jdbc connection parameters
    GJdbcUser = "root"
    GJdbcPasswd = ""
    GJdbcDb = ""

    // parse config yaml file
    if fileName == "" {
        GYamlConf = confS.GetConf(GWriteBackMetaPath + "/meta.yaml")
    } else {
        GYamlConf = confS.GetConf(fileName)
    }

}



func AppendConf(fileName string) {
    var confS ConfStruct
    GYamlConfAppend = confS.GetConf(fileName)
}




func WriteBackMeta(cc *ConfStruct, metaFilePath string) {

    var infoMess       string
    var metaFileName   string
    // check the metaFile exist, if the file doesn't exist, create a new one. 
    metaFileName = metaFilePath + "/meta.yaml"
    _ = os.MkdirAll(metaFilePath, 0666)
    _, err := os.Create(metaFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in create the meta file [fileName = %s]", metaFileName)
        utl.Log("ERROR", infoMess)
    }

    metaF, err := os.OpenFile(metaFileName, os.O_RDWR, 0644)
    if err != nil {
        infoMess = fmt.Sprintf("Error in opening write-back meta file [fileName = %s]", metaFileName)
	utl.Log("ERROR", infoMess)
        clusterNameArr := strings.Split(metaFilePath, "/")
	clusterName := clusterNameArr[len(clusterNameArr)-1]
	infoMess = fmt.Sprintf(`You can shoot the trouble as bellowing step:
	        1. check the meta file status [fileName = %s]
		2. check the cluster name you input [clusterName = %s]
		3. check the os env $SRCTLROOT, if you don't set this env variable, please check the ~/.starrocks-controller folder
	`, metaFileName, clusterName)
	// panic(err)
    }
    defer metaF.Close()



    // write back cluster info
    cc.ClusterInfo.User = GYamlConf.Global.User
    cc.ClusterInfo.CreateDate = time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05")
    cc.ClusterInfo.MetaPath = GWriteBackMetaPath
    cc.ClusterInfo.PrivateKey = GSshKeyRsa

    yamlStr, err := yaml.Marshal(cc)
    if err != nil {
        infoMess = fmt.Sprintf("Error in marshalling yaml structure.")
	utl.Log("ERROR", infoMess)
    }

    _, err = metaF.WriteString(string(yamlStr))
    if err != nil {
        infoMess = fmt.Sprintf("Error in writing back to meta file [fileName = %s]", metaFileName)
	utl.Log("ERROR", infoMess)
    }

}


func SetFeEntry(host string, port int) {
    GFeEntryHost = host
    GFeEntryPort = port
}



func TestParseYamlConfig(fileName string) {


    var confS ConfStruct
    yamlConf := confS.GetConf(fileName)

    // Print configuration
    fmt.Println("[TEST] >>>>>>>>", yamlConf)
    fmt.Println("[TEST] ######################### GLOBAL #########################")
    fmt.Println("[TEST] Global -> User: %s\n", yamlConf.Global.User)
    fmt.Println("[TEST] Global -> ssh_port: %s\n", yamlConf.Global.SshPort)
    fmt.Println("[TEST] ######################### SERVER CONFIG #########################")
    fmt.Println("[TEST] ServerConfig -> FE -> sys_log_level: ", yamlConf.ServerConfig.Fe["sys_log_level"])
    fmt.Println("[TEST] ServerConfig -> FE -> fe_sys_log_1: ", yamlConf.ServerConfig.Fe["fe_sys_log_1"])
    fmt.Println("[TEST] ServerConfig -> BE -> sys_log_level: ", yamlConf.ServerConfig.Be["sys_log_level"])
    fmt.Println("[TEST] ServerConfig -> BE -> sys_log_level: ", yamlConf.ServerConfig.Be["be_sys_log_2"])
    fmt.Println("[TEST] ######################### FE SERVER #########################")
    for i := 0; i < 3; i++ {
        fmt.Printf("[TEST] FeServer -> [%d] -> host:                             %s\n",     i, yamlConf.FeServers[i].Host)
        fmt.Printf("[TEST] FeServer -> [%d] -> ssh_port:                         %s\n",     i, yamlConf.FeServers[i].SshPort)
        fmt.Printf("[TEST] FeServer -> [%d] -> http_port:                        %s\n",     i, yamlConf.FeServers[i].HttpPort)
        fmt.Printf("[TEST] FeServer -> [%d] -> rpc_port:                         %s\n",     i, yamlConf.FeServers[i].RpcPort)
        fmt.Printf("[TEST] FeServer -> [%d] -> query_port:                       %s\n",     i, yamlConf.FeServers[i].QueryPort)
        fmt.Printf("[TEST] FeServer -> [%d] -> edit_log_port:                    %s\n",     i, yamlConf.FeServers[i].EditLogPort)
        fmt.Printf("[TEST] FeServer -> [%d] -> deploy_dir:                       %s\n",     i, yamlConf.FeServers[i].DeployDir)
        fmt.Printf("[TEST] FeServer -> [%d] -> meta_dir:                         %s\n",     i, yamlConf.FeServers[i].MetaDir)
        fmt.Printf("[TEST] FeServer -> [%d] -> log_dir:                          %s\n",     i, yamlConf.FeServers[i].LogDir)
        fmt.Printf("[TEST] FeServer -> [%d] -> priority_networks:                %s\n",     i, yamlConf.FeServers[i].PriorityNetworks)
        fmt.Printf("[TEST] FeServer -> [%d] -> config -> sys_log_level:          %s\n",     i, yamlConf.FeServers[i].Config["sys_log_level"])
        fmt.Printf("[TEST] FeServer -> [%d] -> config -> sys_log_delete_age:     %s\n",     i, yamlConf.FeServers[i].Config["sys_log_delete_age"])
    }

    fmt.Println("[TEST] ######################### BE SERVER #########################")
    for i := 0; i < 3; i++ {
        fmt.Printf("[TEST] BeServer -> [%d] -> host:                             %s\n",     i, yamlConf.BeServers[i].Host)
        fmt.Printf("[TEST] BeServer -> [%d] -> ssh_port:                         %s\n",     i, yamlConf.BeServers[i].SshPort)
        fmt.Printf("[TEST] BeServer -> [%d] -> be_port:                          %s\n",     i, yamlConf.BeServers[i].BePort)
        fmt.Printf("[TEST] BeServer -> [%d] -> webserver_port:                   %s\n",     i, yamlConf.BeServers[i].WebServerPort)
        fmt.Printf("[TEST] BeServer -> [%d] -> heartbeat_service_port:           %s\n",     i, yamlConf.BeServers[i].HeartbeatServicePort)
        fmt.Printf("[TEST] BeServer -> [%d] -> deploy_dir:                       %s\n",     i, yamlConf.BeServers[i].DeployDir)
        fmt.Printf("[TEST] BeServer -> [%d] -> storage_dir:                      %s\n",     i, yamlConf.BeServers[i].StorageDir)
        fmt.Printf("[TEST] BeServer -> [%d] -> log_dir:                          %s\n",     i, yamlConf.BeServers[i].LogDir)
        fmt.Printf("[TEST] BeServer -> [%d] -> config -> sys_log_level:          %s\n",     i, yamlConf.BeServers[i].Config["create_tablet_worker_count"])
        fmt.Printf("[TEST] BeServer -> [%d] -> config -> sys_log_delete_age:     %s\n",     i, yamlConf.BeServers[i].Config["sys_log_delete_age"])
    }

    fmt.Println("[TEST] ######################### PROMETHEUS SERVER #########################")
    fmt.Println("[TEST] PrometheusServer -> host: ",      yamlConf.PrometheusServer.Host)
    fmt.Println("[TEST] PrometheusServer -> ssh_port: ",      yamlConf.PrometheusServer.SshPort)
    fmt.Println("[TEST] PrometheusServer -> http_port: ",      yamlConf.PrometheusServer.HttpPort)
    fmt.Println("[TEST] PrometheusServer -> deploy_dir: ",      yamlConf.PrometheusServer.DeployDir)
    fmt.Println("[TEST] PrometheusServer -> data_dir: ",      yamlConf.PrometheusServer.DataDir)
    fmt.Println("[TEST] PrometheusServer -> log_dir: ",      yamlConf.PrometheusServer.LogDir)

    fmt.Println("[TEST] ######################### GRAFANA SERVER #########################")
    fmt.Println("[TEST] GrafanaServer -> host: ",      yamlConf.GrafanaServer.Host)
    fmt.Println("[TEST] GrafanaServer -> ssh_port: ",      yamlConf.GrafanaServer.SshPort)
    fmt.Println("[TEST] GrafanaServer -> http_port: ",      yamlConf.GrafanaServer.HttpPort)
    fmt.Println("[TEST] GrafanaServer -> deploy_dir: ",      yamlConf.GrafanaServer.DeployDir)

    fmt.Println("[TEST] ######################### ALERTMANAGER SERVER #########################")
    fmt.Println("[TEST] AlertManagerServer -> host: ",      yamlConf.AlertManagerServer.Host)
    fmt.Println("[TEST] AlertManagerServer -> ssh_port: ",      yamlConf.AlertManagerServer.SshPort)
    fmt.Println("[TEST] AlertManagerServer -> web_port: ",      yamlConf.AlertManagerServer.WebPort)
    fmt.Println("[TEST] AlertManagerServer -> cluster_port: ",      yamlConf.AlertManagerServer.ClusterPort)
    fmt.Println("[TEST] AlertManagerServer -> deploy_dir: ",      yamlConf.AlertManagerServer.DeployDir)
    fmt.Println("[TEST] AlertManagerServer -> data_dir: ",      yamlConf.AlertManagerServer.DataDir)
    fmt.Println("[TEST] AlertManagerServer -> log_dir: ",      yamlConf.AlertManagerServer.LogDir)


}
