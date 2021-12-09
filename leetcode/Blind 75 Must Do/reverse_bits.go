/*
Reverse bits of a given 32 bits unsigned integer.

Note:
Note that in some languages such as Java, there is no unsigned integer type. In this case, both
input and output will be given as a signed integer type. They should not affect your implementation,
as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore,
in Example 2 above, the input represents the signed integer -3 and the output represents
the signed integer -1073741825.

Follow up:
If this function is called many times, how would you optimize it?

Example 1:
Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596,
so return 964176192 which its binary representation is 00111001011110000010100101000000.

Example 2:
Input: n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293,
so return 3221225471 which its binary representation is 10111111111111111111111111111111.
*/

package main

import "fmt"

/*
This is the simplest method in this we swap bit by bit. For bit present at index i, after reverse should be present
at 31-i position

Iterate through the bits, from right to left
n = n>>1. To retreive the right most bit -> n&1
For each bit we reverse it to the correct position.
n = (n&1)<<power
*/
func reverse_bits_method_one(n uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for n != 0 {
		ret += (n & 1) << power
		n = n >> 1
		power--
	}
	return ret
}

/*
Method 2:
Resource : LC
To reverse bits for a byte, one could apply the same algorithm as we show in the above approach.
Here we would like to show a different algorithm which is solely based on the arithmetic and bit
operations without resorting to any loop statement, as following:

def reverseByte(byte):
    return (byte * 0x0202020202 & 0x010884422010) % 1023
*/
func reverse_byte(b uint32, cache map[uint32]uint64) uint64 {
	value, ok := cache[b]
	if ok {
		return value
	}
	value = (uint64(b) * 0x0202020202 & 0x010884422010) % 1023
	cache[b] = value
	return value
}

func reverse_bits(num uint32) uint32 {
	ret := uint64(0)
	power := uint64(24)
	var cache = map[uint32]uint64{}

	for num != 0 {
		ret += reverse_byte(num&0xff, cache) << power
		num = num >> 8
		power -= 8
	}
	return uint32(ret)
}
func main() {
	fmt.Println(reverse_bits_method_one(uint32(162)))
	fmt.Println(reverse_bits(uint32(162)))
}
