package importCluster

import (

    "fmt"
    "sr-controller/module"
    "sr-controller/sr-utl"
    "sr-controller/cluster/checkStatus"
    "io/ioutil"
    "net/http"
    "regexp"
    "os"
    "strings"
    "strconv"

)


func GetFeConf() {

    var infoMess      string
    var feHttpUrl     string


    for i := 0; i < len(module.GYamlConf.FeServers); i++ {

        feStat, err := checkStatus.CheckFeStatus(i)
        module.GYamlConf.FeServers[i].HttpPort, _ = strconv.Atoi(feStat["HttpPort"])
	module.GYamlConf.FeServers[i].RpcPort, _ = strconv.Atoi(feStat["RpcPort"])
	module.GYamlConf.FeServers[i].EditLogPort, _ = strconv.Atoi(feStat["EditLogPort"])
	module.GSRVersion = "v" + strings.Split(feStat["Version"], "-")[0]
	rootPasswd := ""

	feHttpUrl = fmt.Sprintf("http://root:%s@%s:%d/variable", rootPasswd, module.GYamlConf.FeServers[i].Host, module.GYamlConf.FeServers[i].HttpPort)
        res, err := http.Get(feHttpUrl)
        defer res.Body.Close()
        if err != nil {
            infoMess = fmt.Sprintf("Error in create http get request when get FE conf. [feHttpUrl = %s, error = %v]", feHttpUrl, err)
            utl.Log("ERROR", infoMess)
            os.Exit(1)
        }

        robots, err := ioutil.ReadAll(res.Body)
        if err != nil {
            infoMess = fmt.Sprintf("Error in read body.[error = %v]", err)
            utl.Log("ERROR", infoMess)
            os.Exit(1)
        }

        //fmt.Println(string(robots))
        // get priority_networks
        r, _ := regexp.Compile("priority_networks=.*")
        module.GYamlConf.FeServers[i].PriorityNetworks = strings.Replace(r.FindString(string(robots)), "priority_networks=", "", -1)

        // get MetaDir
        r, _ = regexp.Compile("meta_dir=.*")
        module.GYamlConf.FeServers[i].MetaDir = strings.Replace(r.FindString(string(robots)), "meta_dir=", "", -1)

        // get LogDir
        r, _ = regexp.Compile("sys_log_dir=.*")
        module.GYamlConf.FeServers[i].LogDir = strings.Replace(r.FindString(string(robots)), "sys_log_dir=", "", -1)
    }

}
