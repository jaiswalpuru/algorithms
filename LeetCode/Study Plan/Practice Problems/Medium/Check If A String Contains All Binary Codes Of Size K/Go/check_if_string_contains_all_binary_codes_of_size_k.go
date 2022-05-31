package main

import "fmt"
import "math"

func check_string_contains_all_bcodes_size_k(s string, k int) bool {
    set := make(map[string]string)

    for i:=0; i<=len(s)-k; i++ {
        set[s[i:i+k]] = s[i:i+k]
    }
    
    return float64(len(set)) == math.Pow(2,float64(k))
}

func main() {
    fmt.Println(check_string_contains_all_bcodes_size_k("00110110", 2))
}
