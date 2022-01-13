package utl

import (

    "io"
    "net/http"
    "os"
    "fmt"
    "strconv"

)

func DebugTest() {

    fmt.Println("This is DebugTest")
}

func IsFileExist(absFileName string, fileSize int64) bool {

    info, err := os.Stat(absFileName)

    if os.IsNotExist(err) {
        fmt.Println(info)
        return false
    }

    if fileSize == info.Size() {
        fmt.Println("The package has already exist.", info.Name(), info.Size(), info.ModTime())
        return true
    }

    del := os.Remove(absFileName)
    if del != nil {
        fmt.Println(del)
    }

    return false
}

func DownloadFile(fileUrl string, localPath string, fileName string) {


    tmpFileName := localPath + "/" + fileName + ".download"
    absFileName := localPath + "/" + fileName

    client := new(http.Client)
    resp, err := client.Get(fileUrl)
    if err != nil { panic(err) }
    fileSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
    if err != nil { fmt.Println(err) }
    fmt.Println("file size is ", fileSize)

    if IsFileExist(absFileName, fileSize) {
        fmt.Println("file exists.")
    }

    tmpFile, err := os.Create(tmpFileName)
    if err != nil { panic(err)}
    defer tmpFile.Close()


    if resp.Body == nil {
        fmt.Println("The download file Body is null.")
    }
    io.Copy(tmpFile, resp.Body)

    if err == nil {
        err = os.Rename(tmpFileName, absFileName)
    }
    fmt.Printf("The file %s [%d] download successfully", fileName, fileSize)

}

