package utl

import (
    "bytes"
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    "regexp"
)


func ModifyConfig(fileName string, sourceStr string, targetStr string) (err error){

    var infoMess string
    _, err = os.Stat(fileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modifing configuration, the configuration file doesn't exist [fileName = %s, sourceStr = %s, targetStr = %s]", fileName, sourceStr, targetStr)
        Log("ERROR", infoMess)
        return err
    }

    input, err := ioutil.ReadFile(fileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modifing configuration, cannot read file [fileName = %s, sourceStr = %s, targetStr = %s]", fileName, sourceStr, targetStr)
	Log("ERROR", infoMess)
	return err
    }

    output := bytes.Replace(input, []byte(sourceStr), []byte(targetStr), -1)

    err = ioutil.WriteFile(fileName, output, 0644)
    if err != nil {
        infoMess = fmt.Sprintf("Error in modifing configuration, cannot read file [fileName = %s, output = %s]", fileName, output)
        Log("ERROR", infoMess)
	return err
    }

    return nil
}
/*
func AppendConfig(fileName string, configStr string) (err error){

    var infoMess string
    _, err = os.Stat(fileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, the configuration file doesn't exist [fileName = %s, configStr = %s]", fileName, configStr)
        Log("ERROR", infoMess)
        return err
    }

    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, cannot open file [fileName = %s, configStr = %s]", fileName, configStr)
	Log("ERROR", infoMess)
	return err
    }
    defer file.Close()

    // if the config exists

    
    _, err = file.WriteString(configStr + "\n")
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, cannot write configStr [fileName = %s, configStr = %s]", fileName, configStr)
	Log("ERROR", infoMess)
	return err
    }

    return err


} 
*/


func AppendConfig(fileName string, configKey string, configValue string) (err error){

    var infoMess string
    _, err = os.Stat(fileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, the configuration file doesn't exist [fileName = %s, configkey = %s, configValue = %s]", fileName, configKey, configValue)
        Log("ERROR", infoMess)
        return err
    }


    configFile, err := ioutil.ReadFile(fileName)
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, read configuration file failed [fileName = %s, configKey = %s, configValue = %s]", fileName, configKey, configValue)
        Log("ERROR", infoMess)
        return err
    }

    lines := strings.Split(string(configFile), "\n")

    for i, line := range lines {
            pattern := fmt.Sprintf("^%s.*", configKey)
            match, _ := regexp.MatchString(pattern, line)
            if match {
	        infoMess := fmt.Sprintf("Comment the default value [fileName = %s, configKey = %s, configValue = %s]", fileName, configKey, configValue)
		Log("INFO", infoMess)
		lines[i] = "# " + lines[i] + "\t\t\t## comment by sr-controller"
	    }
    }

    configStr := fmt.Sprintf("\n%s = %s\n", configKey, configValue)
    output := strings.Join(lines, "\n")
    output = output + configStr

    err = ioutil.WriteFile(fileName, []byte(output), 0644)
    if err != nil {
        infoMess = fmt.Sprintf("Error in appending configuration, write the result to config file failed [fileName = %s, configStr= %s]", fileName, configStr)
	Log("ERROR", infoMess)
    }

    infoMess = fmt.Sprintf("Append configuration [fileName = %s, configStr= %s]", fileName, strings.Replace(configStr, "\n", "", -1))
    Log("DEBUG", infoMess)

    return nil


}

