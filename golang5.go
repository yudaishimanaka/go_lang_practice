package main

import "fmt"

// 型宣言をしない場合
var i, j int = 1, 2

// 型宣言をしない場合
func main(){
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
