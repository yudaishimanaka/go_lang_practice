package main

import (
    "fmt"
)

func main (){
    i := 0
    for ; i < 1000; {
        fmt.Printf("%v\n", i)
        i++
    }
}
