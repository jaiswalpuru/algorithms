package main

import ("fmt"; "strconv")

func reconstruct_original_digits_from_english(s string) string {
	alpha_map := make(map[byte]int)

	n := len(s)
	for i:=0; i<n; i++{
		alpha_map[s[i]]++
	}

	numbers := make([]int, 10)
	//even
	numbers[0] = alpha_map['z']
	numbers[2] = alpha_map['w']
	numbers[4] = alpha_map['u']
	numbers[6] = alpha_map['x']
	numbers[8] = alpha_map['g']

	//odd
	numbers[3] = alpha_map['h'] - numbers[8]
	numbers[5] = alpha_map['f'] - numbers[4]
	numbers[7] = alpha_map['s'] - numbers[6]
	numbers[9] = alpha_map['i'] - numbers[5] - numbers[6] - numbers[8]
	numbers[1] = alpha_map['n'] - numbers[7] - (numbers[9]*2) //2 because nine has 2 n's

	res := ""
	for i:=0; i<10; i++{
		for j:=0; j<numbers[i]; j++{
			res += strconv.Itoa(i)
		}
	}
	return res

}

func main() {
	fmt.Println(reconstruct_original_digits_from_english("owoztneoer"))
}