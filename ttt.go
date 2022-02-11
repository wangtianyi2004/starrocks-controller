package main

import(
    "fmt"
    "reflect"
)

func main() {

    var aaa int
    for i := 0; i < 3; i++ {
	aaa = 20 - 3 * i
        fmt.Printf("i = %d, Type(i) = %v\naaa = %v, Type(aaa) = %v\n", i, reflect.TypeOf(i), aaa, reflect.TypeOf(aaa))
    }

}
