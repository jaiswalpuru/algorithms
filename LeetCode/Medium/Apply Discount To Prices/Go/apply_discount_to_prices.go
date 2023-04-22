package main

import "fmt"
import "strings"
import "strconv"

func apply_discount_to_prices(sentence string, discount int) string {
	str := strings.Split(sentence, " ")
	d := float64(discount)
	res := []string{}
	for i := 0; i < len(str); i++ {
		curr := str[i]
		if curr[0] == '$' && curr[len(curr)-1] != '$' {
			f := false
			for j := 1; j < len(curr); j++ {
				if curr[j] == '$' || !(curr[j] >= '0' && curr[j] <= '9') {
					f = true
					break
				}
			}
			if f {
				res = append(res, curr)
			} else {
				curr = curr[1:]
				f, _ := strconv.ParseFloat(curr, 64)
				num := f - (f*d)/100.0
				val := fmt.Sprintf("%.2f", num)
				res = append(res, "$"+val)
			}
		} else {
			res = append(res, curr)
		}
	}
	s := strings.Join(res, " ")
	return s
}

func main() {
	fmt.Println(apply_discount_to_prices("there are $1 $2 and 5$ candies in the shop", 50))
}
