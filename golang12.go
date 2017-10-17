package main

import (
    "fmt"
)

func main(){
    var sum int = 0
    for i := 0; i < 10; i++ {
        sum += 1
    }
    // for var i int = 0;みたいな書き方はできない
    fmt.Println(sum)
}
