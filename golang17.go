package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }else{
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func main(){
    // 関数の実行結果を関数ごとに出力している
    fmt.Println("=============result==============")
    fmt.Println(pow(3, 2, 10))
    fmt.Println(pow(3, 3, 20))

    // 先に関数を全て実行し終えてから出力する
    fmt.Println("=============result==============")
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
