package main

import ("fmt"; "strconv")

func string_compression(arr []byte) int {
	i,j := 0,0
	n := len(arr)
	for i<n {
		cnt :=0
		arr[j] = arr[i]
		for i<n && arr[i] == arr[j] {
			cnt++
			i++
		}

		if cnt > 1{
			v := []byte(strconv.Itoa(cnt))
			for k:=0;k<len(v);k++{
				j++
				arr[j] = v[k]
			}
		}
		j++
	}
	return j
}

func main() {
	fmt.Println(string_compression([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}
