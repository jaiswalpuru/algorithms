package gostlcode

import "strconv"

func ReverseRune(a *[]rune) string {
	for i, j := 0, len(*a)-1; i < len(*a)/2; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}

	str := ""

	for i := 0; i < len(*a); i++ {
		str += strconv.Itoa(int((*a)[i]))
	}

	return str
}

//ConvertStringToArrInt -> converts strings to array of integers
//return two arrays of numbers
func AddTwoLargeNum(num1, num2 string) string {
	//swapping if length of num1 is greater than num2
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	str := []rune{}
	n1 := []rune(num1)
	n2 := []rune(num2)
	d := len(n2) - len(n1)
	carry := 0

	for i := len(n1) - 1; i >= 0; i-- {
		sum := int(n1[i]-'0') + int(n2[i+d]-'0') + carry
		str = append(str, rune(sum%10))
		carry = sum / 10
	}
	for i := len(n2) - len(n1) - 1; i >= 0; i-- {
		sum := int(n2[i]-'0') + carry
		str = append(str, rune(sum%10))
		carry = sum / 10
	}

	//add remainning carry
	if carry != 0 {
		str = append(str, rune(carry))
	}
	return ReverseRune(&str)
}
