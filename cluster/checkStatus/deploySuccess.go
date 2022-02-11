package checkStatus

import(
    "fmt"
    "sr-controller/sr-utl"
)

func DeploySuccess() {

    var infoMess string

    infoMess = fmt.Sprintf("恭喜你，这么多的 bug 还能部署成功。")
    utl.Log("OUTPUT", infoMess)
    infoMess = fmt.Sprintf("四百多个异常捕获都没有捕获到你的芳心。RESPECT")
    utl.Log("OUTPUT", infoMess)
    
}
