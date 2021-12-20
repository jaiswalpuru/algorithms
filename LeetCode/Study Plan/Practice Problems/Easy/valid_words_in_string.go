/*
A sentence consists of lowercase letters ('a' to 'z'), digits ('0' to '9'), hyphens ('-'), punctuation marks
('!', '.', and ','), and spaces (' ') only. Each sentence can be broken down into one or more tokens separated by
one or more spaces ' '.

A token is a valid word if all three of the following are true:

It only contains lowercase letters, hyphens, and/or punctuation (no digits).
There is at most one hyphen '-'. If present, it must be surrounded by lowercase characters ("a-b" is valid, but
"-ab" and "ab-" are not valid).
There is at most one punctuation mark. If present, it must be at the end of the token ("ab,", "cd!", and "."
are valid, but "a!b" and "c.," are not valid).
Examples of valid words include "a-b.", "afad", "ba-c", "a!", and "!".
Given a string sentence, return the number of valid words in sentence.
*/

package main

import (
	"fmt"
	"strings"
)

func count_valid(str string) int {
	s := strings.Split(str, " ")
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != "" {
			word := s[i]
			m := len(word)
			mp := make(map[byte]int)
			cnt++
			for j := 0; j < m; j++ {
				if word[j] >= '0' && word[j] <= '9' {
					cnt--
					break
				} else {
					if word[j] == '-' {
						mp[word[j]]++
						if j == 0 || j == m-1 || mp[word[j]] > 1 || (j != m-1 && !(word[j+1] >= 'a' && word[j+1] <= 'z')) {
							cnt--
							break
						}
					} else if word[j] == '!' || word[j] == '.' || word[j] == ',' {
						mp[word[j]]++
						if j != m-1 || mp[word[j]] > 1 {
							cnt--
							break
						}
					}
				}
			}
		}
	}
	return cnt
}

func main() {
	// fmt.Println(count_valid("a-b-c"))
	// fmt.Println(count_valid("cat and  dog"))
	// fmt.Println(count_valid("!this  1-s b8d!"))
	// fmt.Println(count_valid("alice and  bob are playing stone-game10"))
	fmt.Println(count_valid(" 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex uy5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1  !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49, ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   lh w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq    inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4fo9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif "))
}
