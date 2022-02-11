package utl

import(
    "fmt"
    "time"
)

/*

LogLevel 
- DEBUG     10
- INFO      20
- WARN      30
- ERROR     40

*/
var GLOGLEVEL string = "INFO"


func Log(logLevel string, mess string) {

    // logLevel: DEBUG INFO WARN ERROR
    dt := string(time.Now().Format("20060102-150405"))
/*
    var infoMess string

    if logLevel == "ERROR" {
	infoMess = "你遇到异常的样子，都那么美 。。。"
        fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, infoMess)
	infoMess = "是 BUG 吗？不要怕，程序员就要多写点 BUG，要和测试一同进步。"

    }
*/

    switch GLOGLEVEL {
        case "DEBUG":
            // output: DEBUG INFO WARN ERROR
            fmt.Printf("[\x1b[47;30m%s\x1b[0m\x1b[43;30m%8s\x1b[0m] %s\n", dt, logLevel, mess)
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


