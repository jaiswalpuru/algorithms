/*
Given a string queryIP, return "IPv4" if IP is a valid IPv4 address, "IPv6" if IP is a valid IPv6 address or "Neither" if
IP is not a correct IP of any type.

A valid IPv4 address is an IP in the form "x1.x2.x3.x4" where 0 <= xi <= 255 and xi cannot contain leading zeros.
For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses but "192.168.01.1", while "192.168.1.00" and
"192.168@1.1" are invalid IPv4 addresses.

A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:

1 <= xi.length <= 4
xi is a hexadecimal string which may contain digits, lower-case English letter ('a' to 'f') and upper-case English
letters ('A' to 'F').
Leading zeros are allowed in xi.
For example, "2001:0db8:85a3:0000:0000:8a2e:0370:7334" and "2001:db8:85a3:0:0:8A2E:0370:7334" are valid IPv6 addresses,
while "2001:0db8:85a3::8A2E:037j:7334" and "02001:0db8:85a3:0000:0000:8a2e:0370:7334" are invalid IPv6 addresses.
*/

package main

import (
	"fmt"
	"strings"
)

func validate_IP(str string) string {

	var list []string
	if strings.Contains(str, ".") {
		list = strings.Split(str, ".")
	} else {
		list = strings.Split(str, ":")
	}
	if len(list) != 4 && len(list) != 8 {
		return "Neither"
	}

	if len(list) == 4 {
		for i := 0; i < len(list); i++ {
			temp := list[i]
			m := len(temp)
			if m > 3 || m < 0 || (m > 1 && temp[0] == '0') {
				return "Neither"
			}
			num := 0
			for j := 0; j < m; j++ {
				if !(temp[j] >= '0' && temp[j] <= '9') {
					return "Neither"
				}
				num = (num * 10) + int(temp[j]-'0')
			}
			if num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		for i := 0; i < len(list); i++ {
			temp := list[i]
			m := len(temp)
			if m < 0 || m > 4 {
				return "Neither"
			}
			for j := 0; j < m; j++ {
				if !(temp[j] >= '0' && temp[j] <= '9') && !(temp[j] >= 'A' && temp[j] <= 'F') && !(temp[j] >= 'a' && temp[j] <= 'f') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}

}

func main() {
	fmt.Println(validate_IP("172.16.254.1"))
	fmt.Println(validate_IP("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}
