package utl

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
)


func ModifyConfig(fileName string, sourceStr string, targetStr string) {

    input, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    output := bytes.Replace(input, []byte(sourceStr), []byte(targetStr), -1)

    if err = ioutil.WriteFile(fileName, output, 0666); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

