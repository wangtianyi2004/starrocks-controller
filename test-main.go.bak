package main

import(
    "fmt"
    "os"
    "sr-controller-v1/module"
)

func main() {

    fmt.Println("This is test main ===> ")
    fmt.Println("args[0]: ", os.Args[0])
    fmt.Println("args[1]: ", os.Args[1])

    controllerModule := os.Args[1]

    switch controllerModule {
        case "playground":
            fmt.Println("This is PLAYGROUND module.")
        case "cluster":
	    fmt.Println("This is CLUSTER module.")
	case "test":
            fmt.Println("This is TEST module.")
	    yamlFileName := os.Args[2]
            module.TestParseYamlConfig(yamlFileName)
        default:
            fmt.Println("******** UNKNOWN ***********")
    }
}
