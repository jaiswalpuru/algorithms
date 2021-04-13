package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	mp := map[string]int{
		"2":  0,
		"3":  1,
		"4":  2,
		"5":  3,
		"6":  4,
		"7":  5,
		"8":  6,
		"9":  7,
		"10": 8,
		"J":  9,
		"Q":  10,
		"K":  11,
		"A":  12,
	}

	// n: the number of cards for player 1
	var n int
	fmt.Scan(&n)
	p1 := make([]int, n)
	for i := 0; i < n; i++ {
		var cardp1 string
		fmt.Scan(&cardp1)
		p1[i] = mp[string([]rune(cardp1)[0])]
	}

	// m: the number of cards for player 2
	var m int
	fmt.Scan(&m)
	p2 := make([]int, m)

	for i := 0; i < m; i++ {
		var cardp2 string
		fmt.Scan(&cardp2)
		p2[i] = mp[string([]rune(cardp2)[0])]
	}
	fmt.Fprintln(os.Stderr, p1, p2)
	rounds := 1
	temp1, temp2 := []int{}, []int{}
	for {
		tc1, tc2 := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		temp1 = append(temp1, tc1)
		temp2 = append(temp2, tc2)
		if tc1 != tc2 {
			if tc1 > tc2 {
				p1 = append(p1, temp1...)
				p1 = append(p1, temp2...)
			} else {
				p2 = append(p2, temp1...)
				p2 = append(p2, temp2...)
			}
			temp1, temp2 = []int{}, []int{}
			if len(p1) == 0 {
				fmt.Println("2", rounds)
				break
			}
			if len(p2) == 0 {
				fmt.Println("1", rounds)
				break
			}
			rounds++
		} else if len(p1) < 3 || len(p2) < 3 {
			fmt.Println("PAT")
			break
		} else {
			temp1 = append(temp1, []int{p1[0], p1[1], p1[2]}...)
			temp2 = append(temp2, []int{p2[0], p2[1], p2[2]}...)
			p1 = p1[3:]
			p2 = p2[3:]
		}
	}
}
