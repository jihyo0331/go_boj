// BOJ 5357 Dedupe
package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)

    for i := 0; i < n; i++ {
        var s string
        fmt.Scanf("%s", &s)

        if len(s) == 0 {
            fmt.Println()
            continue
        }

        result := []byte{s[0]}
        for j := 1; j < len(s); j++ {
            if s[j] != s[j-1] {
                result = append(result, s[j])
            }
        }
        fmt.Println(string(result))
    }
}