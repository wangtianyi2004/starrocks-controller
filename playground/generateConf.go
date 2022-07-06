package playground

import (

    "stargo/module"
    "time"
    "fmt"
    _ "embed"
    "os/user"
    "os"

)



func InitPlaygroundConf(){

    var tmp module.ConfStruct


    osUser, _ := user.Current()
    module.GSshKeyRsa = fmt.Sprintf("%s/.ssh/id_rsa", osUser.HomeDir)
    module.GSRCtlRoot = os.Getenv("SRCTLROOT")
    if module.GSRCtlRoot == "" {
        module.GSRCtlRoot = fmt.Sprintf("%s/.stargo", osUser.HomeDir)
    }
    tmpDeployDir := fmt.Sprintf("%s/playground", module.GSRCtlRoot)

    // ClusterInfo
    tmp.ClusterInfo.User = osUser.Username
    tmp.ClusterInfo.Version = "2.2.0"
    tmp.ClusterInfo.CreateDate = time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05")
    tmp.ClusterInfo.MetaPath = fmt.Sprintf("%s/cluster/sr-playground", module.GSRCtlRoot)
    tmp.ClusterInfo.PrivateKey = module.GSshKeyRsa


    tmp.Global.User = osUser.Username
    tmp.Global.SshPort = 22

    tmp.FeServers = append(tmp.FeServers,
        struct {
            Host                      string                 `yaml:"host"`
            SshPort                   int                    `yaml:"ssh_port"`
            HttpPort                  int                    `yaml:"http_port"`
            RpcPort                   int                    `yaml:"rpc_port"`
            QueryPort                 int                    `yaml:"query_port"`
            EditLogPort               int                    `yaml:"edit_log_port"`
            DeployDir                 string                 `yaml:"deploy_dir"`
            MetaDir                   string                 `yaml:"meta_dir"`
            LogDir                    string                 `yaml:"log_dir"`
            PriorityNetworks          string                 `yaml:"priority_networks"`
            Config                    map[string]string      `yaml:"config"`
        } {
            Host:                     "127.0.0.1",
            SshPort:                  22,
            HttpPort:                 8030,
            RpcPort:                  9020,
            QueryPort:                9030,
            EditLogPort:              9010,
            DeployDir:                tmpDeployDir + "/fe",
            MetaDir:                  tmpDeployDir + "/fe/meta",
            LogDir:                   tmpDeployDir + "/fe/log",
            PriorityNetworks:         "127.0.0.1/32",
            Config:                   nil,
        })

    tmp.BeServers = append(tmp.BeServers,
        struct {
            Host                     string                   `yaml:"host"`
            SshPort                  int                      `yaml:"ssh_port"`
            BePort                   int                      `yaml:"be_port"`
            WebServerPort            int                      `yaml:"webserver_port"`
            HeartbeatServicePort     int                      `yaml:"heartbeat_service_port"`
            BrpcPort                 int                      `yaml:"brpc_port"`
            DeployDir                string                   `yaml:"deploy_dir"`
            StorageDir               string                   `yaml:"storage_dir"`
            LogDir                   string                   `yaml:"log_dir"`
            PriorityNetworks         string                   `yaml:"priority_networks"`
            Config                   map[string]string        `yaml:"configs`
        } {
            Host:                    "127.0.0.1",
            SshPort:                 22,
            BePort:                  9060,
            WebServerPort:           8040,
            HeartbeatServicePort:    9050,
	    BrpcPort:                8060,
            DeployDir:               tmpDeployDir + "/be",
	    StorageDir:              tmpDeployDir + "/be/storage",
	    LogDir:                  tmpDeployDir + "/be/log",
	    PriorityNetworks:        "127.0.0.1/32",
	    Config:                  nil,
        })


    module.GYamlConf = &tmp
    module.GSRVersion = "v" + module.GYamlConf.ClusterInfo.Version
}





