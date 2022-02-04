package module


import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "fmt"
    "os"
)

var GYamlConf   *ConfStruct
var GSshKeyRsa   string
var GSRCtlRoot   string
type ConfStruct struct {
    Global struct {
        User        string         `yaml:"user"`
        SshPort     int            `yaml:ssh_port`
    } `yaml:"global"`

    ServerConfig struct {
        Fe            map[string]string `yaml:"fe"`
        Be            map[string]string `yaml:"be"`
    } `yaml:"server_configs"`

    FeServers []struct {
        Host                  string                 `yaml:"host"`
        SshPort               int                    `yaml:"ssh_port"`
        HttpPort              int                    `yaml:"http_port"`
        RpcPort               int                    `yaml:"rpc_port`
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

func InitConf(fileName string) {
    var confS ConfStruct 
    GYamlConf = confS.GetConf(fileName)
    GSshKeyRsa = "/root/.ssh/id_rsa"

    
    GSRCtlRoot = os.Getenv("SRCTLROOT")
    if GSRCtlRoot == "" {
        if GYamlConf.Global.User == "root" {
            GSRCtlRoot = "/root/.starrocks-controller"
        } else {
            GSRCtlRoot = fmt.Sprintf("/home/%s/.starrocks-controller", GYamlConf.Global.User) 
        }
    }

    
}

func TestParseYamlConfig(fileName string) {


    var confS ConfStruct
    yamlConf := confS.GetConf(fileName)

    // Print configuration
    fmt.Println(">>>>>>>>", yamlConf)
    fmt.Println("######################### GLOBAL #########################")
    fmt.Println("Global -> User: ", yamlConf.Global.User)
    fmt.Println("Global -> ssh_port: ", yamlConf.Global.SshPort)
    fmt.Println("######################### SERVER CONFIG #########################")
    fmt.Println("ServerConfig -> FE -> sys_log_level: ", yamlConf.ServerConfig.Fe["sys_log_level"])
    fmt.Println("ServerConfig -> FE -> fe_sys_log_1: ", yamlConf.ServerConfig.Fe["fe_sys_log_1"])
    fmt.Println("ServerConfig -> BE -> sys_log_level: ", yamlConf.ServerConfig.Be["sys_log_level"])
    fmt.Println("ServerConfig -> BE -> sys_log_level: ", yamlConf.ServerConfig.Be["be_sys_log_2"])
    fmt.Println("######################### FE SERVER #########################")
    for i := 0; i < 3; i++ {
        fmt.Printf("FeServer -> [%d] -> host:                             %s\n",     i, yamlConf.FeServers[i].Host)
        fmt.Printf("FeServer -> [%d] -> ssh_port:                         %s\n",     i, yamlConf.FeServers[i].SshPort)
        fmt.Printf("FeServer -> [%d] -> http_port:                        %s\n",     i, yamlConf.FeServers[i].HttpPort)
        fmt.Printf("FeServer -> [%d] -> rpc_port:                         %s\n",     i, yamlConf.FeServers[i].RpcPort)
        fmt.Printf("FeServer -> [%d] -> query_port:                       %s\n",     i, yamlConf.FeServers[i].QueryPort)
        fmt.Printf("FeServer -> [%d] -> edit_log_port:                    %s\n",     i, yamlConf.FeServers[i].EditLogPort)
        fmt.Printf("FeServer -> [%d] -> deploy_dir:                       %s\n",     i, yamlConf.FeServers[i].DeployDir)
        fmt.Printf("FeServer -> [%d] -> meta_dir:                         %s\n",     i, yamlConf.FeServers[i].MetaDir)
        fmt.Printf("FeServer -> [%d] -> log_dir:                          %s\n",     i, yamlConf.FeServers[i].LogDir)
        fmt.Printf("FeServer -> [%d] -> priority_networks:                %s\n",     i, yamlConf.FeServers[i].PriorityNetworks)
        fmt.Printf("FeServer -> [%d] -> config -> sys_log_level:          %s\n",     i, yamlConf.FeServers[i].Config["sys_log_level"])
        fmt.Printf("FeServer -> [%d] -> config -> sys_log_delete_age:     %s\n",     i, yamlConf.FeServers[i].Config["sys_log_delete_age"])
    }

    fmt.Println("######################### BE SERVER #########################")
    for i := 0; i < 3; i++ {
        fmt.Printf("BeServer -> [%d] -> host:                             %s\n",     i, yamlConf.BeServers[i].Host)
        fmt.Printf("BeServer -> [%d] -> ssh_port:                         %s\n",     i, yamlConf.BeServers[i].SshPort)
        fmt.Printf("BeServer -> [%d] -> be_port:                          %s\n",     i, yamlConf.BeServers[i].BePort)
        fmt.Printf("BeServer -> [%d] -> webserver_port:                   %s\n",     i, yamlConf.BeServers[i].WebServerPort)
        fmt.Printf("BeServer -> [%d] -> heartbeat_service_port:           %s\n",     i, yamlConf.BeServers[i].HeartbeatServicePort)
        fmt.Printf("BeServer -> [%d] -> deploy_dir:                       %s\n",     i, yamlConf.BeServers[i].DeployDir)
        fmt.Printf("BeServer -> [%d] -> storage_dir:                      %s\n",     i, yamlConf.BeServers[i].StorageDir)
        fmt.Printf("BeServer -> [%d] -> log_dir:                          %s\n",     i, yamlConf.BeServers[i].LogDir)
        fmt.Printf("BeServer -> [%d] -> config -> sys_log_level:          %s\n",     i, yamlConf.BeServers[i].Config["create_tablet_worker_count"])
        fmt.Printf("BeServer -> [%d] -> config -> sys_log_delete_age:     %s\n",     i, yamlConf.BeServers[i].Config["sys_log_delete_age"])
    }

    fmt.Println("######################### PROMETHEUS SERVER #########################")
    fmt.Println("PrometheusServer -> host: ",      yamlConf.PrometheusServer.Host)
    fmt.Println("PrometheusServer -> ssh_port: ",      yamlConf.PrometheusServer.SshPort)
    fmt.Println("PrometheusServer -> http_port: ",      yamlConf.PrometheusServer.HttpPort)
    fmt.Println("PrometheusServer -> deploy_dir: ",      yamlConf.PrometheusServer.DeployDir)
    fmt.Println("PrometheusServer -> data_dir: ",      yamlConf.PrometheusServer.DataDir)
    fmt.Println("PrometheusServer -> log_dir: ",      yamlConf.PrometheusServer.LogDir)

    fmt.Println("######################### GRAFANA SERVER #########################")
    fmt.Println("GrafanaServer -> host: ",      yamlConf.GrafanaServer.Host)
    fmt.Println("GrafanaServer -> ssh_port: ",      yamlConf.GrafanaServer.SshPort)
    fmt.Println("GrafanaServer -> http_port: ",      yamlConf.GrafanaServer.HttpPort)
    fmt.Println("GrafanaServer -> deploy_dir: ",      yamlConf.GrafanaServer.DeployDir)

    fmt.Println("######################### ALERTMANAGER SERVER #########################")
    fmt.Println("AlertManagerServer -> host: ",      yamlConf.AlertManagerServer.Host)
    fmt.Println("AlertManagerServer -> ssh_port: ",      yamlConf.AlertManagerServer.SshPort)
    fmt.Println("AlertManagerServer -> web_port: ",      yamlConf.AlertManagerServer.WebPort)
    fmt.Println("AlertManagerServer -> cluster_port: ",      yamlConf.AlertManagerServer.ClusterPort)
    fmt.Println("AlertManagerServer -> deploy_dir: ",      yamlConf.AlertManagerServer.DeployDir)
    fmt.Println("AlertManagerServer -> data_dir: ",      yamlConf.AlertManagerServer.DataDir)
    fmt.Println("AlertManagerServer -> log_dir: ",      yamlConf.AlertManagerServer.LogDir)


}
