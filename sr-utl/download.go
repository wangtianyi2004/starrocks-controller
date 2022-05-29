package utl

import (

    "io"
    "net/http"
    "os"
    "fmt"
    "strconv"

)


func IsFileExist(absFileName string, fileSize int64) bool {

    var infoMess string
    info, err := os.Stat(absFileName)

    if os.IsNotExist(err) {
        infoMess = fmt.Sprintf("Detect file %s doesn't exist.", absFileName)
        Log("DEBUG", infoMess)
        return false
    }

    if fileSize == info.Size() {
        infoMess = fmt.Sprintf("The package has already exist [fileName = %v, fileSize = %v, fileModTime = %v]", info.Name(), info.Size(), info.ModTime())
        Log("INFO", infoMess)
        return true
    }

    del := os.Remove(absFileName)
    if del != nil {
        infoMess = fmt.Sprintf("Delete file %s", absFileName)
        Log("WARN", infoMess)
    }

    return false
}


func DownloadFile(fileUrl string, localPath string, fileName string) {

    var infoMess string
    tmpFileName := localPath + "/" + fileName + ".download"
    absFileName := localPath + "/" + fileName

    client := new(http.Client)
    resp, err := client.Get(fileUrl)
    if err != nil { 
        infoMess = fmt.Sprintf("Error in get the response for %s", fileUrl)
        Log("ERROR", infoMess)
        panic(err) 
    }

    fileSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
    if err != nil { 
        infoMess = fmt.Sprintf("Error in parsing the size for file %s", fileUrl)
        Log("ERROR", infoMess)
    }

    if IsFileExist(absFileName, fileSize) {
        // the file exist, it doesn't need to download the one
	return
    }

    tmpFile, err := os.Create(tmpFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in create the tmp file %s", tmpFileName)
        Log("ERROR", infoMess)
        panic(err)
    }
    defer tmpFile.Close()


    if resp.Body == nil {
        Log("ERROR", "The download file Body is null.")
	Log("ERROR", infoMess)
	panic(err)
    }

    io.Copy(tmpFile, resp.Body)
    info, err := os.Stat(tmpFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Cannot get the tmp file stat [fileName = %s]", tmpFileName)
	Log("ERROR", infoMess)
	panic(err)
    }

    if info.Size() != fileSize {
        infoMess = fmt.Sprintf("Error in download, pls check your network connection.")
	Log("ERROR", infoMess)
	panic(err)
    }

    if err == nil {
        err = os.Rename(tmpFileName, absFileName)
    }

    infoMess = fmt.Sprintf("The file %s [%d] download successfully", fileName, fileSize)
    Log("INFO", infoMess)

}

