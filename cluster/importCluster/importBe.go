package importCluster

import (

    "fmt"
    "stargo/module"
    "stargo/sr-utl"
    "stargo/cluster/checkStatus"
    "io/ioutil"
    "net/http"
    "regexp"
    "os"
    "strings"
    "strconv"

)


func GetBeConf() {

    var infoMess      string
    var beHttpUrl     string

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {

        beStat, err := checkStatus.CheckBeStatus(i)
        module.GYamlConf.BeServers[i].WebServerPort, _ = strconv.Atoi(beStat["HttpPort"])
	module.GYamlConf.BeServers[i].BrpcPort, _ = strconv.Atoi(beStat["BrpcPort"])
	module.GYamlConf.BeServers[i].BePort, _ = strconv.Atoi(beStat["BePort"])

	rootPasswd := ""

	beHttpUrl = fmt.Sprintf("http://root:%s@%s:%d/varz", rootPasswd, module.GYamlConf.BeServers[i].Host, module.GYamlConf.BeServers[i].WebServerPort)
        res, err := http.Get(beHttpUrl)
        defer res.Body.Close()
        if err != nil {
            infoMess = fmt.Sprintf("Error in create http get request when get BE conf. [beHttpUrl = %s, error = %v]", beHttpUrl, err)
            utl.Log("ERROR", infoMess)
            os.Exit(1)
        }

        robots, err := ioutil.ReadAll(res.Body)
        if err != nil{
            infoMess = fmt.Sprintf("Error in read body.[error = %v]", err)
            utl.Log("ERROR", infoMess)
            os.Exit(1)
        }

        //fmt.Println(string(robots))
        // get priority_networks
        r, _ := regexp.Compile("priority_networks=.*")
        module.GYamlConf.BeServers[i].PriorityNetworks = strings.Replace(r.FindString(string(robots)), "priority_networks=", "", -1)

        // get MetaDir
        r, _ = regexp.Compile("storage_root_path=.*")
        module.GYamlConf.BeServers[i].StorageDir = strings.Replace(r.FindString(string(robots)), "storage_root_path=", "", -1)

        // get LogDir
        r, _ = regexp.Compile("sys_log_dir=.*")
        module.GYamlConf.BeServers[i].LogDir = strings.Replace(r.FindString(string(robots)), "sys_log_dir=", "", -1)


    }

}
