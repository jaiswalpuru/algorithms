package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomain_visit_count(str []string) []string {
	domain := make(map[string]int)
	n := len(str)
	for i := 0; i < n; i++ {
		val := strings.Split(str[i], " ")
		num, _ := strconv.Atoi(val[0])
		d := strings.Split(val[1], ".")
		temp := ""
		for j := len(d) - 1; j >= 0; j-- {
			temp = d[j] + temp
			domain[temp] += num
			temp = "." + temp
		}
	}
	res := []string{}
	for k, v := range domain {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}

func main() {
	fmt.Println(subdomain_visit_count([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
