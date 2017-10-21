package main

import (
    "fmt"
)

func aika(){
    fmt.Println("Aika!")
}

func kobayashi(){
    fmt.Println("Kobayashi")
}

func main(){
    defer aika()

    kobayashi()
}
