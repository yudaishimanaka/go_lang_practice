package main

import (
    "fmt"
)

const Pi = 3.14

func main(){
    const World = "世界" // 定数
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true // 定数
    fmt.Println("Go rules?", Truth)
}
