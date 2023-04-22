package main

import "fmt"

func calculate_the_amt_paid_in_taxes(arr [][]int, income int) float64 {
	if income == 0 {
		return 0
	}
	tax := float64(0)
	if income > arr[0][0] {
		tax += float64(arr[0][0]*arr[0][1]) / 100
		income -= arr[0][0]
	} else {
		tax += float64(income*arr[0][1]) / 100
		income = 0
	}

	for i := 1; i < len(arr); i++ {
		if income == 0 {
			break
		}

		if income > (arr[i][0] - arr[i-1][0]) {
			tax += float64((arr[i][0]-arr[i-1][0])*arr[i][1]) / 100
			income -= arr[i][0] - arr[i-1][0]
		} else {
			tax += float64(income*arr[i][1]) / 100
			income = 0
		}
	}
	return tax
}

func main() {
	fmt.Println(calculate_the_amt_paid_in_taxes([][]int{
		{3, 50}, {7, 10}, {12, 25},
	}, 10))
}
