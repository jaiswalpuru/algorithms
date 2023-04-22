package main

import (
	"fmt"
	"strings"
)

type HtmlParser interface{ GetUrls(url string) []string }

/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */

func web_crawl(start_url string, html_parser HtmlParser) []string {
	visited := make(map[string]struct{})
	q := []string{start_url}
	dom := strings.Split(start_url, "/")[2]
	visited[start_url] = struct{}{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, val := range html_parser.GetUrls(curr) {
			temp := strings.Split(val, "/")[2]
			if dom == temp {
				if _, ok := visited[val]; !ok {
					visited[val] = struct{}{}
					q = append(q, val)
				}
			}
		}
	}

	res := []string{}
	for k := range visited {
		res = append(res, k)
	}
	return res
}

func main() {
	fmt.Println("This is webcrawler")
}
