package main

import ("fmt"; "strconv"; "sort")

func  to_string(a,b,c,d int) string { 
	return strconv.Itoa(a) + "->" + strconv.Itoa(b) + "->" + strconv.Itoa(c) + "->" + strconv.Itoa(d) 
}

//---------------------worst solution one can ever come up with-------------------------
func four_sum(arr []int, target int) [][]int {
	n := len(arr)
	sort.Ints(arr)
	res := [][]int{}
	hash_map := make(map[string]bool)
	for i:=0; i<n; i++{
		for j:=i+1; j<n; j++{
			for k:=j+1 ;k<n; k++{
				for l:=k+1; l<n; l++{
					if hash_map[to_string(arr[i],arr[j],arr[k],arr[l])] {
						continue
					}
					if arr[i]+arr[j]+arr[k]+arr[l] == target {
						res = append(res, []int{arr[i],arr[j],arr[k],arr[l]})
						hash_map[to_string(arr[i], arr[j], arr[k], arr[l])] = true
					}
				}
			}
		}
	}
	return res
}
//---------------------worst solution one can ever come up with-------------------------


//------------------------------efficient approach------------------------------ (reference : LC)
func two_sum_2p(arr []int, target, start int) [][]int {
	i, j := start, len(arr)-1
	res := [][]int{}
	for i<j {
		sum := arr[i]+arr[j]
		if sum < target || i > start && arr[i] == arr[i-1] {
			i++
		}else if sum > target || j < len(arr)-1 && arr[j] == arr[j+1] {
			j--
		}else {
			res = append(res, []int{arr[i], arr[j]})
			i++
			j--
		}
	}
	return res
}


func four_sum_eff(arr []int, target int) [][]int {
	sort.Ints(arr)
	return k_sum(arr, target, 0, 4) //this is generic function for k sum here we are doing 4 sum hence the end is 4 and start is zero
}

func k_sum(arr []int, target int, start, end int) [][]int{
	res := [][]int{}
	if start == len(arr) {
		return res
	}

	avg_val := target / end

	if avg_val < arr[start] || avg_val > arr[len(arr)-1] {
		return res
	}

	if end == 2 {
		return two_sum_2p(arr, target, start)
	}

	for i:=start; i<len(arr);i++{
		if i==start || arr[i-1]!=arr[i] {
			for _, v := range k_sum(arr, target-arr[i], i+1, end-1) {
				res = append(res, []int{arr[i]})
				res[len(res)-1] = append(res[len(res)-1], v...)
			}
		}
	}
	return res
}

//------------------------------efficient approach------------------------------

func main() {
	fmt.Println(four_sum_eff([]int{1,0,-1,0,-2,2}, 0))
}



