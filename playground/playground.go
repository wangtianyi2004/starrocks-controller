package playground

import (

    "fmt"
    "sr-controller/sr-utl"

)



func StartFePlayground() bool {

    ModifyFEConfig()
    RunFEProcess()
    res := CheckFEStatus()
    return res

}


func StartBePlayground() bool {

    ModifyBEConfig()
    AddBENode()
    RunBEProcess()
    res := CheckBEStatus()
    return res

}


func RunPlayground() {


    var infoMess       string
    InitPlaygroundConf()

    PrecheckPlayground()
    PreparePlaygroundDir()
    feSuccess := StartFePlayground()
    beSuccess := StartBePlayground()

    if feSuccess && beSuccess {
	infoMess = fmt.Sprintf("Playground run successfully. Please use bellowing command to connect StarRocks playground cluster:\nmysql -uroot -P9030 -h127.0.0.1")
        utl.Log("OUTPUT", infoMess)
    }

}
