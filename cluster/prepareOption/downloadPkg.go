package prepareOption


import (
    "stargo/sr-utl"
    "stargo/module"
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)



func PrepareSRPkg() {

    var infoMess string

    infoMess = "Download StarRocks package & jdk ..."
    utl.Log("OUTPUT", infoMess)
    DownloadSRPkg()

    infoMess = "Decompress StarRocks pakcage & jdk ..."
    utl.Log("OUTPUT", infoMess)
    DecompressSRPkg()
}


func GetDownloadUrl(srVersion string)  (downloadUrl string) {
    
    var   infoMess  string
    var   repoUrl   string

    module.GetRepo()

    // deal with 
    if strings.Contains(module.GRepo.Repo, "file://") {
        downloadUrl = ""
        return downloadUrl
    }

    repoUrl = module.GRepo.Repo + "/packageVersion.list"
    res, err := http.Get(repoUrl)
    defer res.Body.Close()
    if err != nil { 
        infoMess = fmt.Sprintf("Error in create http get request when download the repo list. [error = %v]", err)
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }

    robots, err := ioutil.ReadAll(res.Body)
    if err != nil{ 
        infoMess = fmt.Sprintf("Error in read body.[error = %v]", err)
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }


    versionList := strings.Split(string(robots), "\n")
    vLabel := fmt.Sprintf("[%s]", srVersion)

    for i := 0; i < len(versionList); i++ {
        if strings.Contains(versionList[i], vLabel) {
            downloadUrl = strings.Replace(versionList[i], vLabel, "", -1)
            downloadUrl = strings.Replace(downloadUrl, " ", "", -1)
            break
        } else {
            downloadUrl = ""
        }
    }

    if downloadUrl == "" {
        infoMess = fmt.Sprintf("Error in get version %s package, pls check it again. [DownloadUrl = %s]", srVersion, downloadUrl)
        utl.Log("ERROR", infoMess)
        os.Exit(1)
    }
    return downloadUrl

}

func DownloadSRPkg() {


    // download sr & jdk union package
    // "http://cdn-thirdparty.starrocks.com/starrocks-2.0.1-quickstart.tar.gz?Expires=10282764349&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=kXpA4RHT3sg4Lz9vyRJtbnPdmqM%3D"
    // "http://cdn-thirdparty.starrocks.com/starrocks-2.1.3-quickstart.tar.gz?Expires=2511847820&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=izihf34yKm7ppk5DENn8jEO2vuw%3D"
    // fmt.Sprintf("http://192.168.88.89:9000/starrocks-quick-start/starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))

    var infoMess string
    pkgUrl := GetDownloadUrl(module.GSRVersion)

    // downloadPath := module.GSRCtlRoot + "/download"
    if pkgUrl != "" {
        downloadFile := fmt.Sprintf("starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
        utl.DownloadFile(pkgUrl, module.GDownloadPath, downloadFile)
        infoMess = fmt.Sprintf("Download done.")
        utl.Log("OUTPUT", infoMess)
    }
}


func DecompressSRPkg() {


    var tarFileName string
    var destFilePath string
    var infoMess string

    // Decompress SR & JDK union pakcage
    tarFileName = fmt.Sprintf("%s/starrocks-%s-quickstart.tar.gz", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
    //tarFileName = module.GSRCtlRoot + fmt.Sprintf("/download/starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    // destFilePath = module.GSRCtlRoot + "/download"
    destFilePath = module.GDownloadPath
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

    // Decompress StarRocks Package
    tarFileName = fmt.Sprintf("%s/StarRocks-%s.tar.gz", module.GDownloadPath, strings.Replace(module.GSRVersion, "v", "", -1))
    // tarFileName = module.GSRCtlRoot + fmt.Sprintf("/download/StarRocks-%s.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    // destFilePath = module.GSRCtlRoot + "/download"
    destFilePath = module.GDownloadPath
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

    // Decompress JDK Package
    tarFileName = module.GDownloadPath + "/jdk-8u301-linux-x64.tar.gz"
    // tarFileName = module.GSRCtlRoot + "/download/jdk-8u301-linux-x64.tar.gz"
    destFilePath = module.GDownloadPath
    // destFilePath = module.GSRCtlRoot + "/download"
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

}



