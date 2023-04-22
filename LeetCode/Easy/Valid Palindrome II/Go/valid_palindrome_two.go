package main

func valid_palindrome_two(s string) bool {
	i, j := 0, len(s)-1
	cnt1, cnt2 := 0, 0
	for i<=j {
		if s[i] == s[j] {
			i++
			j--
		}else{
			i++
			cnt1++
		}
	}

	i,j := 0, len(s)-1
	for i<=j {
		if s[i] == s[j] {
			i++
			j--
		}else {
			j--
			cnt2++
		}
	}


	if cnt1 <= 1 || cnt2 <= 1{
		return true
	}
	return false
}

func main() {
	fmt.Println(valid_palindrome_two("aba"))
}