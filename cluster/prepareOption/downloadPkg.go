package prepareOption


import (
    "sr-controller/sr-utl"
    "sr-controller/module"
    "fmt"
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
    pkgUrl := "http://cdn-thirdparty.starrocks.com/starrocks-2.0.1-quickstart.tar.gz?Expires=10282764349&OSSAccessKeyId=LTAI4GFYjbX9e7QmFnAAvkt8&Signature=kXpA4RHT3sg4Lz9vyRJtbnPdmqM%3D"
    downloadPath := module.GSRCtlRoot + "/download"
    downloadFile := "starrocks-2.0.1-quickstart.tar.gz"
    utl.DownloadFile(pkgUrl, downloadPath, downloadFile)
    infoMess = fmt.Sprintf("这么大的文件，这么慢的网，费劲吧啦下载好了。跨过千山万水，来爱你，比心。")
    utl.Log("OUTPUT", infoMess)
}


func DecompressSRPkg() {


    var tarFileName string
    var destFilePath string
    var infoMess string

    // Decompress SR & JDK union pakcage
    tarFileName = module.GSRCtlRoot + "/download/starrocks-2.0.1-quickstart.tar.gz"
    destFilePath = module.GSRCtlRoot + "/download"
    utl.UnTargz(tarFileName, destFilePath)
    infoMess = fmt.Sprintf("The tar file %s has been decompressed under %s", tarFileName, destFilePath)
    utl.Log("INFO", infoMess)

    // Decompress StarRocks Package
    tarFileName = module.GSRCtlRoot + "/download/StarRocks-2.0.1.tar.gz"
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



