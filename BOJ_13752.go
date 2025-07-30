// BOG 13752 히스토그램

package main

import "fmt" 

func main() {
	var n int
    hi := []int{}
    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++ {
        var temp int
        fmt.Scanf("%d", &temp)
        hi = append(hi, temp)
    }
    for i := 0; i< len(hi); i++ {
        for j := 0; j < hi[i]; j++ {
            fmt.Print("=")
        }
        fmt.Println()
    }
}
