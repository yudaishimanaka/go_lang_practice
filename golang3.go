package main

import "fmt"

// 戻り値の変数名をしてすることによってnaked returnにする
func sum(x, y int) (sum int) {
        sum = x + y
        return
        }

func main(){
        fmt.Println(sum(5, 5))
        }
