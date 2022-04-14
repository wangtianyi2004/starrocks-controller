package prepareOption


import (
    "sr-controller/sr-utl"
    "sr-controller/module"
    "fmt"
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

func DownloadSRPkg() {


    // download sr & jdk union package
    // http://cdn-thirdparty.starrocks.com/starrocks-2.0.1-quickstart.tar.gz?Expires=10282764349&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=kXpA4RHT3sg4Lz9vyRJtbnPdmqM%3D
    var infoMess string
    pkgUrl := fmt.Sprintf("http://192.168.88.89:9000/starrocks-quick-start/starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    downloadPath := module.GSRCtlRoot + "/download"
    downloadFile := fmt.Sprintf("starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    utl.DownloadFile(pkgUrl, downloadPath, downloadFile)
    infoMess = fmt.Sprintf("Download done.")
    utl.Log("OUTPUT", infoMess)
}


func DecompressSRPkg() {


    var tarFileName string
    var destFilePath string
    var infoMess string

    // Decompress SR & JDK union pakcage
    tarFileName = module.GSRCtlRoot + fmt.Sprintf("/download/starrocks-%s-quickstart.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    destFilePath = module.GSRCtlRoot + "/download"
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

    // Decompress StarRocks Package
    tarFileName = module.GSRCtlRoot + fmt.Sprintf("/download/StarRocks-%s.tar.gz", strings.Replace(module.GSRVersion, "v", "", -1))
    destFilePath = module.GSRCtlRoot + "/download"
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

    // Decompress JDK Package
    tarFileName = module.GSRCtlRoot + "/download/jdk-8u301-linux-x64.tar.gz"
    destFilePath = module.GSRCtlRoot + "/download"
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

}



