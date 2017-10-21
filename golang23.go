package main

import (
    "fmt"
)

func main(){
    fmt.Println("counting")

    // deferされたfmt.Println(i)はLIFOの順で出力される
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
