package prepareOption

import(
    "sr-controller/sr-utl"
    "sr-controller/module"
    "fmt"
)


func CreateFeDir() {

    var infoMess string
    var errMess  string
    var cmd string
    var err error
    //var outPut []byte

    sshUser := module.GYamlConf.Global.User
    sshKeyRsaFile := module.GSshKeyRsa

    for i := 0; i < len(module.GYamlConf.FeServers); i++ {
        sshHost := module.GYamlConf.FeServers[i].Host
        sshPort := module.GYamlConf.FeServers[i].SshPort

        // create DEPLOY dir for FE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.FeServers[i].DeployDir)
        infoMess = fmt.Sprintf("Create DEPLOY Folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating DEPLOY folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }

        // create META dir for FE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.FeServers[i].MetaDir)
        infoMess = fmt.Sprintf("Create META Folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd) 
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating META folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }
        
        if module.GYamlConf.FeServers[i].DeployDir + "/meta" != module.GYamlConf.FeServers[i].MetaDir {
            cmd = fmt.Sprintf("ln -s %s %s", module.GYamlConf.FeServers[i].MetaDir, module.GYamlConf.FeServers[i].DeployDir + "/meta")
            infoMess = fmt.Sprintf("Detect MetaDir isn't under DeployDir, Create the soft link, CMD %s", cmd)
            utl.Log("WARN", infoMess)
            _, err := utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
            if err != nil {
                errMess = fmt.Sprintf("Error in create soft link for MetaDir, CMD %s", cmd)
                utl.Log("ERROR", errMess)
                panic(err)
            }
        }

        // create LOG dir for FE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.FeServers[i].LogDir)
        infoMess = fmt.Sprintf("Create LOG Folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating LOG folder for FE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }
    } 
}




func CreateBeDir() {
    var infoMess string
    var errMess  string
    var cmd string
    var err error
    //var outPut []byte

    sshUser := module.GYamlConf.Global.User
    sshKeyRsaFile := module.GSshKeyRsa

    for i := 0; i < len(module.GYamlConf.BeServers); i++ {
        sshHost := module.GYamlConf.BeServers[i].Host
        sshPort := module.GYamlConf.BeServers[i].SshPort

        // create DEPLOY dir for BE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.BeServers[i].DeployDir)
        infoMess = fmt.Sprintf("Create DEPLOY Folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating DEPLOY folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }

        // create STORAGE dir for BE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.BeServers[i].StorageDir)
        infoMess = fmt.Sprintf("Create Storage Folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating STORAGE folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }

        if module.GYamlConf.BeServers[i].DeployDir + "/storage" != module.GYamlConf.BeServers[i].StorageDir {
            cmd = fmt.Sprintf("ln -s %s %s", module.GYamlConf.BeServers[i].StorageDir, module.GYamlConf.BeServers[i].DeployDir + "/storage")
            infoMess = fmt.Sprintf("Detect StorageDir isn't under DeployDir, Create the soft link, CMD %s", cmd)
            utl.Log("WARN", infoMess)
            _, err := utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
            if err != nil {
                errMess = fmt.Sprintf("Error in create soft link for StorageDir, CMD %s", cmd)
                utl.Log("ERROR", errMess)
                panic(err)
            }
        }

        // create LOG dir for BE nodes
        cmd = fmt.Sprintf("mkdir -p %s", module.GYamlConf.BeServers[i].LogDir)
        infoMess = fmt.Sprintf("Create LOG Folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
        utl.Log("INFO", infoMess)

        _, err = utl.SshRun(sshUser, sshKeyRsaFile, sshHost, sshPort, cmd)
        if err != nil {
            errMess = fmt.Sprintf("ERROR in creating LOG folder for BE node: %s@%s:%d \"%s\"", sshUser, sshHost, sshPort, cmd)
            utl.Log("ERROR", errMess)
            panic(err)
        }

    }

}
