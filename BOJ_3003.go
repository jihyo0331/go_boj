// BOJ 3003 킹, 퀸, 룩, 비숍, 나이트, 폰

package main

import "fmt"

func main(){
    var pieces = [6]int{1, 1, 2, 2, 2, 8}
    var n = [6]int{}
    var r = [6]int{}
    fmt.Scanf("%d %d %d %d %d %d", &n[0], &n[1], &n[2], &n[3], &n[4], &n[5])
    for i := 0; i < 6; i++ {
        if pieces[i] > n[i]{
            r[i] = pieces[i] - n[i]
        }else if pieces[i] < n[i]{
            r[i] = -n[i] + pieces[i]
        }
    }
    fmt.Println(r[0], r[1], r[2], r[3], r[4], r[5])
}