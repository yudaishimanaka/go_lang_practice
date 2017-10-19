package main

import (
    "fmt"
    "math"
)

const eps = 1e-9 // 差があまりにも小さい場合ループ回数を減らす際に使用する

func Sqrt(x float64) float64 {
    z := 1.0
    p := z
    i := 0
    var diff float64 = 0
    for {
        // ニュートン方による平方根の近似
        z = z - ((z * z -x) / (2 * z))

        // 差があまりにも小さい場合はループを抜ける
        diff = math.Abs(z -p)
        if math.Abs(z - p) < eps {
            break
        }
        p = z
        i++
        // ループ回数と差の表示
        fmt.Printf("loop:%v %v\n", i, diff)
    }
    return z
}

func main(){
    fmt.Println("=====result=====")
    fmt.Println(math.Sqrt(2))
    fmt.Println("=====result=====")
    fmt.Println(Sqrt(2))
}
