//Given an array of integers return the new array such that each element at index i of new array is the product
//of all the numbers in the original array
//eg  arr = []int{1,2,3,4,5 } O/P->[120,60,40,30,24]
//This problem can be solved by division easily -> calculate the product of all the numbers in the array and
//then for each index store -> π(arr)/arr[i]

//This problem can also be solved using division method
// π(xi) where 1≤x≤n
//Compute the product of all the elements in the array, for each element replace arr[i] with prod/arr[i]

package main

import "fmt"

func prefixProduct(arr []int) []int {
	prefixProd := []int{}
	for i := 0; i < len(arr); i++ {
		if len(prefixProd) != 0 {
			prefixProd = append(prefixProd, prefixProd[len(prefixProd)-1]*arr[i])
		} else {
			prefixProd = append(prefixProd, arr[i])
		}
	}
	return prefixProd
}

func suffixProduct(arr []int) []int {
	suffixProd := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		if len(suffixProd) != 0 {
			suffixProd = append(suffixProd, suffixProd[len(suffixProd)-1]*arr[i])
		} else {
			suffixProd = append(suffixProd, arr[i])
		}
	}
	return suffixProd
}

func reverse(arr *[]int) {
	j := len(*arr) - 1
	for i := 0; i <= len(*arr)/2; i++ {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		j--
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	prefProd := prefixProduct(arr)
	suffProd := suffixProduct(arr)
	reverse(&suffProd)
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result = append(result, suffProd[i+1])
		} else if i == len(arr)-1 {
			result = append(result, prefProd[i-1])
		} else {
			result = append(result, prefProd[i-1]*suffProd[i+1])
		}
	}
	fmt.Println(result)
}
