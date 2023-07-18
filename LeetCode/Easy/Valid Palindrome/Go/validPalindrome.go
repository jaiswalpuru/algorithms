func isPalindrome(s string) bool {
	n := len(s)
	filter := ""
	for i := 0; i < n; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			filter += strings.ToLower(string(s[i]))
		}
		if s[i] >= '0' && s[i] <= '9' {
			filter += string(s[i])
		}
	}

	//reverse the string
	reverse_string := ""
	n = len(filter)
	for i := n - 1; i >= 0; i-- {
		reverse_string += string(filter[i])
	}
	return filter == reverse_string
}
