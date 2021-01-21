package main

import "fmt"

func main() {
    x := "Sean"
    ret := ""
    for i := len(x)-1; i >= 0; i-- {
        ret += string(x[i])
    }
    fmt.Println(ret)
}
