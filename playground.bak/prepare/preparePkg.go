package preparePkg

import (
    "sr-controller/sr-utl"
    "fmt"
    "os"
//    "os/exec"
//    "strings"
)

func DownloadSRPkg() {

    /*
    // pkgUrl := "https://cdn-release.starrocks.com/StarRocks-2.0.0-GA.tar.gz?Expires=1641905572&OSSAccessKeyId=LTAI4GE6QnJq4iytf6SawLgb&Signature=9oDxXL8a10WRGdCRtBBjkkyS3YQ%3D"
    // download sr package
    pkgUrl := "http://10.10.10.20:9000/starrocks/StarRocks-2.0.0-GA.tar.gz"
    downloadPath := "/root/.starrocks-controller/download"
    downloadFile := "StarRocks-2.0.0-GA.tar.gz"

    utl.DownloadFile(pkgUrl, downloadPath, downloadFile)

    // download jdk package 
    pkgUrl = "http://10.10.10.20:9000/starrocks/jdk-8u301-linux-x64.tar.gz"
    downloadPath = "/root/.starrocks-controller/download"
    downloadFile = "jdk-8u301-linux-x64.tar.gz"
    utl.DownloadFile(pkgUrl, downloadPath, downloadFile)
    */

    // download sr & jdk union package 
    // http://cdn-thirdparty.starrocks.com/starrocks-2.0.1-quickstart.tar.gz?Expires=10282764349&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=kXpA4RHT3sg4Lz9vyRJtbnPdmqM%3D
    pkgUrl := "http://cdn-thirdparty.starrocks.com/starrocks-2.0.1-quickstart.tar.gz?Expires=10282764349&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=kXpA4RHT3sg4Lz9vyRJtbnPdmqM%3D"
    downloadPath := "/root/.starrocks-controller/download"
    downloadFile := "starrocks-2.0.1-quickstart.tar.gz"
    utl.DownloadFile(pkgUrl, downloadPath, downloadFile)

}


func DecompressSRPkg() {


    var tarFileName string
    var destFilePath string

    // Decompress SR & JDK union pakcage
    tarFileName = "/root/.starrocks-controller/download/starrocks-2.0.1-quickstart.tar.gz"
    destFilePath = "/root/.starrocks-controller/download"
    utl.UnTargz(tarFileName, destFilePath)
    fmt.Printf("The tar file %s has been decompressed under %s\n", tarFileName, destFilePath)

    // Decompress StarRocks Package
    tarFileName = "/root/.starrocks-controller/download/StarRocks-2.0.1.tar.gz"
    destFilePath = "/root/.starrocks-controller/download"
    utl.UnTargz(tarFileName, destFilePath)
    fmt.Printf("The tar file %s has been decompressed under %s\n", tarFileName, destFilePath)

    // Decompress JDK Package
    tarFileName = "/root/.starrocks-controller/download/jdk-8u301-linux-x64.tar.gz"
    destFilePath = "/root/.starrocks-controller/download"
    utl.UnTargz(tarFileName, destFilePath)
    fmt.Printf("The tar file %s has been decompressed under %s\n", tarFileName, destFilePath)

}


func DeployPkg() {

    var sourceDir string
    var targetDir string
    var err error

    // deploy jdk folder
    sourceDir = "/root/.starrocks-controller/download/jdk1.8.0_301"
    targetDir = "/root/.starrocks-controller/playground/jdk1.8.0"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }
    fmt.Printf("mv %s to %s\n", sourceDir, targetDir)

    // deploy fe folder
    sourceDir = "/root/.starrocks-controller/download/StarRocks-2.0.1/fe"
    targetDir = "/root/.starrocks-controller/playground/fe"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }
    fmt.Printf("mv %s to %s\n", sourceDir, targetDir)

    // deploy be folder
    sourceDir = "/root/.starrocks-controller/download/StarRocks-2.0.1/be"
    targetDir = "/root/.starrocks-controller/playground/be"
    err = os.Rename(sourceDir, targetDir)
    if err != nil { panic(err) }
    fmt.Printf("mv %s to %s\n", sourceDir, targetDir)

}
/*
func RunShellScript(scriptName string) string {
    cmd := exec.Command("/bin/bash", "-c", scriptName)
    res, err := cmd.Output()
    if err != nil { panic(err) }
    fmt.Println(string(res))
    return string(res)
}
*/


