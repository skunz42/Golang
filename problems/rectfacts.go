package main

import "fmt"
import "os"
import "strconv"

func calcArea(x, y int) int {
    return x*y
}

func calcPerim(x, y int) int {
    return 2*x + 2*y
}

func main() {
    x, err1 := strconv.Atoi(os.Args[1])
    y, err2 := strconv.Atoi(os.Args[2])
    if err1 == nil && err2 == nil {
        fmt.Println("Area:", calcArea(x,y))
        fmt.Println("Perimeter:", calcPerim(x,y))
    }
}
