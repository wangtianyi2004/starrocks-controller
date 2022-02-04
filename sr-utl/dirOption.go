package utl

import (
    "os"
    "io"
    "fmt"
)


func MkDir(dirPath string) {

    // check dir exists
    mess := ""
    dir, _ := os.Stat(dirPath)
    if dir == nil {
        // dir doesn't exist, create new one
        e := os.MkdirAll(dirPath, 751)
        if e != nil {
            mess = "Error in create folder [" + dirPath + "]"
            Log("ERROR", mess)
            panic(e)
        }
    } else {
        mess = "Detect the folder [" + dirPath + "] exists"
        Log("INFO", mess)
    }
}

func CopyFile(sourceFileName string, targetFileName string) (fileByte int64, err error) {

    var infoMess string

    sourceFileStat, err := os.Stat(sourceFileName)

    if err != nil {
	infoMess = fmt.Sprintf("Error in copy file, the source file doesn't exist [sourceFile = %, targetFile = %s]", sourceFileName, targetFileName)
	Log("ERROR", infoMess)
        return 0, err
    }

    if !sourceFileStat.Mode().IsRegular() {
        infoMess = fmt.Sprintf("Error in copy file, the source file isn't a regular file [sourceFile = %s, targetFile = %s]", sourceFileName, targetFileName)
	Log("ERROR", infoMess)
	return 0, err
    }

    src, err := os.Open(sourceFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in copy file, the source file cannot be opened [sourceFile = %s, targetFile = %s]", sourceFileName, targetFileName)
	Log("ERROR", infoMess)
	return 0, err
    }
    defer src.Close()

    dest, err := os.Create(targetFileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in copy file, the target file cannot be created [sourceFile = %s, targetFile = %s]", sourceFileName, targetFileName)
        Log("ERROR", infoMess)
        return 0, err
    }
    defer dest.Close()

    fileByte, err = io.Copy(dest, src)
    if err != nil {
        infoMess = fmt.Sprintf("Error in copy file [sourceFile = %s, targetFile = %s]", sourceFileName, targetFileName)
        Log("ERROR", infoMess)
	return 0, err
    }

    return fileByte, err

}
