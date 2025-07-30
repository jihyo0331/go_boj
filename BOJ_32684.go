// BOJ 32684 장기

package main

import "fmt"

type pieces struct{
    zol int
    sa int
    sang int
    ma int
    po int
    cha int
    gung int
    sum float64
}

func main() {
    var cho pieces
    var han pieces
    fmt.Scanf("%d %d %d %d %d %d %d", &cho.cha, &cho.po, &cho.ma, &cho.sang, &cho.sa, &cho.zol)
    fmt.Scanf("%d %d %d %d %d %d %d", &han.cha, &han.po, &han.ma, &han.sang, &han.sa, &han.zol)
    cho.sum = float64((cho.zol * 2) + (cho.sa * 3) + (cho.sang * 3) + (cho.ma * 5) + (cho.po * 7) + (cho.cha * 13))
    han.sum = float64((han.zol * 2) + (han.sa * 3) + (han.sang * 3) + (han.ma * 5) + (han.po * 7) + (han.cha * 13)) + 1.5

    if cho.sum > han.sum{
        fmt.Println("cocjr0208")
    } else{
        fmt.Println("ekwoo")
    }
}