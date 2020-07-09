package main 

import "fmt"



func f() *int{
    var v int = 1
    return &v
}


func main() {

    pointer := f()
    fmt.Printf("pointer:%x, *pointer:%d\n", pointer, *pointer)
}
