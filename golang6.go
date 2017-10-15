package main

import "fmt"

func main(){
    var i, j int = 1, 2 // 一般的な変数宣言
    var k = 3 // 初期化子の型を使用した変数宣言
    c, python, java := true, false, "no!" // varステートメントを省略した暗黙的な宣言

    fmt.Println(i, j, k, c, python, java)
}
