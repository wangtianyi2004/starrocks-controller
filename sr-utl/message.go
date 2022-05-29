package utl

import(
    "fmt"
    "time"
    "os"
//    "io"
//    "io/ioutil"
    "bufio"
)

/*

LogLevel 
- DEBUG     10
- INFO      20
- WARN      30
- ERROR     40

*/
//var GLOGLEVEL string = "DEBUG"
var GLOGLEVEL string = "DEBUG"

func Log(logLevel string, mess string) {


    dt := string(time.Now().Format("20060102-150405"))

    debugFile := "debug.log"

    if GLOGLEVEL == "DEBUG" {

	/*
	file, err := os.OpenFile(debugFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
            fmt.Printf("Error in open debug.log")
	}
	defer file.Close()
	logMess := fmt.Sprintf("[%s %s] %s", dt, logLevel, mess)
	_, _ = io.WriteString(file, logMess)
        */
	logMess := fmt.Sprintf("[%s %s] %s", dt, logLevel, mess)
        file, _ := os.OpenFile(debugFile, os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	write := bufio.NewWriter(file)
        write.WriteString(logMess)
	write.Flush()
    }

    // logLevel: DEBUG INFO WARN ERROR

    switch logLevel {
        case "DEBUG":
            // output: DEBUG INFO WARN ERROR
            //fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
	    debugFile = "debug.log"
        case "INFO":
	    if logLevel != "DEBUG" {
	        fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
	    }
	case "WARN":
	    if logLevel != "DEBUG" || logLevel != "INFO" {
	        fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
	    }
	case "ERROR":
	    if logLevel != "DEBUG" || logLevel != "INFO" || logLevel != "WARN" {
                fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
            }
	default:
	    fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
    }
    //fmt.Printf("%s   %8s %15s  %s\n", dt, logLevel, process, mess)
    //fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)

}


