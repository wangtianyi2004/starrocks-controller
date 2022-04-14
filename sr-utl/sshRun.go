package utl

import (
    "fmt"
    "os"
    "path"
    "golang.org/x/crypto/ssh"
    "github.com/pkg/sftp"
    "io/ioutil"
)




func NewConfig(keyFile string, user string)(config *ssh.ClientConfig, err error) {

    var errmess string

    key, err := ioutil.ReadFile(keyFile)
    if err != nil {
        errmess = fmt.Sprint("unable to read private key: %v", err)
        Log("ERROR", errmess)
        return nil, err
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        errmess = fmt.Sprint("unable to parse private key: %v", err)
        Log("ERROR", errmess)
        return nil, err
    }

    config = &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    return config, nil

}

func sshRun(config *ssh.ClientConfig, host string, port int, command string) (outPut []byte, err error) {

    var errmess string
    var infoMess string
    client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
    if err != nil {
        errmess = fmt.Sprint("unable to connect: %s error %v", host, err)
        Log("ERROR", errmess)
        return nil, err
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil{
        errmess = fmt.Sprint("ssh new session error %v", err)
        Log("ERROR", errmess)
        return nil, err
    }
    defer session.Close()

    outPut, err = session.CombinedOutput(command)
    if err != nil {
        if v, ok := err.(*ssh.ExitError); ok {
            errmess = v.Msg()
        }
    }

    infoMess = fmt.Sprintf(`ssh run: 
                                        host = %s
                                        cmd = %s
                                        error = %v
                                        result = %v`, host, command, errmess, string(outPut))
    Log("DEBUG", infoMess)

    return outPut, err
}

func SshRun(user string, keyFile string, host string, port int, command string) (outPut []byte, err error) {

    var infoMess string

    sshConfig, err := NewConfig(keyFile, user)
    if err != nil {
        infoMess = fmt.Sprintf("Failed to get the ssh config when run command [user = %s, host = %s, port = %d, cmd = %s], error = %v", user, host, port, command, err)
	Log("DEBUG", infoMess)
    }

    output, err := sshRun(sshConfig, host, port, command)
    if err != nil {
        infoMess = fmt.Sprintf("Failed to run command. [host = %s, cmd = %s, error = %v]", host, command, err)
        Log("DEBUG", infoMess)
    }
    return output, err

}

func sftpConnect(config *ssh.ClientConfig, host string, port int) (sfpClient *sftp.Client, err error) {

    var infoMess string
    addr := fmt.Sprintf("%s:%d", host, port)

    sshClient, err := ssh.Dial("tcp", addr, config)
    if err != nil {
        infoMess = fmt.Sprintf("Error in dail %s, %s", addr, config)
	Log("ERROR", infoMess)
	return nil, err
    }

    sftpClient, err := sftp.NewClient(sshClient)
    if err != nil {
        infoMess = fmt.Sprintf("Error in get sftp client")
	Log("ERROR", infoMess)
	return nil, err
    }

    return sftpClient, nil

}


func uploadFile(sftpClient *sftp.Client, localFileName string, remoteFileName string) (err error) {

    var infoMess string

    srcFile, err := os.Open(localFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in open file %s", localFileName)
	Log("ERROR", infoMess)
	return err
    }
    defer srcFile.Close()

    dstFile, err := sftpClient.Create(remoteFileName)
    if err != nil {
	infoMess = fmt.Sprintf("sftpClient.Create error : %s, error = %v", remoteFileName, err)
        Log("ERROR", infoMess)
	return err
    }
    defer dstFile.Close()

    ff, err := ioutil.ReadAll(srcFile)
    if err != nil {
	infoMess = fmt.Sprintf("ReadAll error : %s", localFileName)
        Log("ERROR", infoMess)
	return err
    }

    dstFile.Write(ff)
    //infoMess = localFileName + " copy file to remote server finished!"
    //Log("DEBUG", infoMess)
    // Chmod remoteFile
    fileStat, err := os.Stat(localFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in get file stat when upload file: [sourceFile = %s  targetFile = %s]", localFileName, remoteFileName)
	Log("ERROR", infoMess)
	return err
    }

    err = sftpClient.Chmod(remoteFileName, fileStat.Mode())
    if err != nil {
        infoMess = fmt.Sprintf("Error in chmod file stat when upload file: [sourceFile = %s  targetFile = %s]", localFileName, remoteFileName)
	Log("ERROR", infoMess)
	return err
    }
    //infoMess = fmt.Sprintf("chmod file [%s] to %s", remoteFileName, fileStat.Mode())
    //Log("DEBUG", infoMess)
    //Log("INFO", infoMess)
    return err
}


func uploadDirectory(sftpClient *sftp.Client, localPath string, remotePath string) (err error) {

    var infoMess string

    localFiles, err := ioutil.ReadDir(localPath)
    if err != nil {
        infoMess = "Read dir list fail."
	Log("ERROR", infoMess)
	return err
    }

    for _, backupDir := range localFiles {

	localFilePath := path.Join(localPath, backupDir.Name())
        remoteFilePath := path.Join(remotePath, backupDir.Name())

	if backupDir.IsDir() {
            sftpClient.Mkdir(remoteFilePath)
            err = uploadDirectory(sftpClient, localFilePath, remoteFilePath)
	    if err != nil {
	        infoMess = fmt.Sprintf("Error in upload dir %s\t%s\t%s", sftpClient, localFilePath, remoteFilePath)
		Log("ERROR", infoMess)
		return err
	    }
        } else {
	    localFileName := path.Join(localPath, backupDir.Name())
	    remoteFileName := path.Join(remotePath, backupDir.Name())
            err = uploadFile(sftpClient, localFileName, remoteFileName)
	    if err != nil {
	        infoMess = fmt.Sprintf("Error in upload file %s\t%s\t%s", sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
		Log("ERROR", infoMess)
		return err
	    }
        }

    }

    //infoMess = localPath + " copy directory to remote server finished!"
    //Log("INFO", infoMess)
    return err
}



func UploadFile(user string, keyFile string, host string, port int, sourceFile string, targetFile string) {

    var infoMess string

    sshConfig, err := NewConfig(keyFile, user)
    if err != nil {
        infoMess = fmt.Sprintf("Error in upload file, fail to get ssh config [keyfile = %s, user = %s]", keyFile, user)
	Log("ERROR", infoMess)
	panic(err)
    }

    sftpClient, err := sftpConnect(sshConfig, host, port)
    if err != nil {
        infoMess = fmt.Sprintf("Error in upload file, fail to get sftp client [keyfile = %s, user = %s, host = %s, port = %d]", keyFile, user, host, port)
        Log("ERROR", infoMess)
	panic(err)
    }

    err = uploadFile(sftpClient, sourceFile, targetFile)
    if err != nil {
        infoMess = fmt.Sprintf("Error in upload file [user = %s, keyFile = %s, host = %s, port = %d, sourceFile = %s, targetFile = %s]", user, keyFile, host, port, sourceFile, targetFile)
	Log("ERROR", infoMess)
	panic(err)
    }

}



func UploadDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string) {

    var infoMess string
    var err error
    // check the folder exist
    cmd := fmt.Sprintf("ls %s", targetDir)
    _, err = SshRun(user, keyFile, host, port, cmd)
    if err != nil {
	infoMess = fmt.Sprintf("The target dir [%s] doesn't exist on [%s:%d], create a new one", targetDir, host, port)
	Log("DEBUG", infoMess)
	cmd = fmt.Sprintf("mkdir -p %s", targetDir)
        _, err := SshRun(user, keyFile, host, port, cmd)
        if err != nil {
            infoMess = fmt.Sprintf("Error in create folder [%s] on [%s:%d]", targetDir, host, port)
	    Log("ERROR", infoMess)
            panic(err)
        }
        infoMess = fmt.Sprintf("Create folder [%s] on [%s:%d]", targetDir, host, port)
        Log("DEBUG", infoMess)
    }

    sshConfig, err := NewConfig(keyFile, user)
    if err != nil {
        infoMess = fmt.Sprintf(`Error in upload dir, failed to get the ssh config :user = %s, 
                                        keyFile = %s
                                        host = %s
                                        port = %d
                                        sourceDir = %s
                                        targetDir = %s
                                        error = %v`, user, keyFile, host, port, sourceDir, targetDir, err)
    }
    sftpClient, err := sftpConnect(sshConfig, host, port)
    if err != nil {
        infoMess = fmt.Sprintf(`Error in upload dir[sftp client]: user = %s
                                        keyFile = %s
                                        host = %s
                                        port = %d
                                        sourceDir = %s
                                        targetDir = %s
                                        error = %v`, user, keyFile, host, port, sourceDir, targetDir, err)
        Log("ERROR", infoMess)
	panic(err)
    }

    err = uploadDirectory(sftpClient, sourceDir, targetDir)
    if err != nil {
        infoMess = fmt.Sprintf(`Error in upload dir[upload dir]: user = %s
                                        keyFile = %s
                                        host = %s
                                        port = %d
                                        sourceDir = %s
                                        targetDir = %s
                                        error = %v`, user, keyFile, host, port, sourceDir, targetDir, err)
        Log("ERROR", infoMess)
	panic(err)
    }

}

func RenameDir(user string, keyFile string, host string, port int, sourceDir string, targetDir string) (err error){

    var infoMess string

    cmd := fmt.Sprintf("ls %s", sourceDir)
    _, err = SshRun(user, keyFile, host, port, cmd)

    if err != nil {
        infoMess = fmt.Sprintf("The source dir [%s] doesn't exist on [%s:%d], create a new one", sourceDir, host, port)
        Log("ERROR", infoMess)
        return err 
    }

    sshConfig, err := NewConfig(keyFile, user)
    if err != nil {
        infoMess = fmt.Sprintf("Error in rename dir, failed to get the ssh config. [host = %s, sourceDir = %s, targetDir = %s, err = %v]", host, sourceDir, targetDir, err)
        Log("ERROR", infoMess)
        return err
    }

    sftpClient, err := sftpConnect(sshConfig, host, port)
    if err != nil {
        infoMess = fmt.Sprintf("Error in rename dir when create sftp client.[host = %s, sourceDir = %s, targetDir = %s, error = %v", host, sourceDir, targetDir, err)
        Log("ERROR", infoMess)
        return err
    }
    err = sftpClient.Rename(sourceDir, targetDir)
    if err != nil {
        infoMess = fmt.Sprintf("Error in rename dir.[host = %s, sourceDir = %s, targetDir = %s, error = %v", host, sourceDir, targetDir, err)
        Log("ERROR", infoMess)
        return err
    }

    return nil

}

func TestUploadDir() {


    // check targetDir exist
    output, err := SshRun("root", "/root/.ssh/id_rsa",  "192.168.230.41", 22, "ls /opt/starrocks/fe/jdk")
    fmt.Printf("[TEST] The result of [ls /opt/starrocks/fe/jdk] on 192.168.230.41:22 ---- output = %s, error = %v\n", output, err)

    if err != nil {
	fmt.Println("[TEST] XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
        fmt.Println("[TEST] The target dir [/opt/starrocks/fe/jdk] doesn't exist on [192.168.230.41].")
        _, err := SshRun("root", "/root/.ssh/id_rsa", "192.168.230.41", 22, "mkdir -p /opt/starrocks/fe/jdk")

	if err != nil {
	    fmt.Println("[TEST] Error in create folder [/opt/starrocks/fe/jdk] on [192.168.230.41]")
	    panic(err)
	}

    }
/*
    sftpClient, err := sftpConnect(sshConfig, "192.168.230.41", 22)
    if err != nil { panic(err) }
    uploadDirectory(sftpClient, "/tmp/aaaDir", "/opt/soft/tmp")
*/
}


