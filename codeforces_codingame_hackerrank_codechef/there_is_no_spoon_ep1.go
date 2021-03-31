package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Don't let the machines win. You are humanity's last hope...
 **/

type Pair struct {
	x, y, dir int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: the number of cells on the X axis
	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	// height: the number of cells on the Y axis
	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)
	arr := make([][]rune, height)
	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		tmp := []rune(strings.Join(strings.Split(line, " "), "")) // to avoid unused error // width characters, each either 0 or .
		arr[i] = make([]rune, len(tmp))
		for j := 0; j < len(tmp); j++ {
			arr[i][j] = tmp[j]
		}
	}

	graph := make([][]Pair, height*width)
	p := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			graph[p] = append(graph[p], Pair{x: i, y: j, dir: -1})
			if arr[i][j] != '.' {
				//right
				for k := j + 1; k < width; k++ {
					if arr[i][k] != '.' {
						graph[p] = append(graph[p], Pair{x: k, y: i, dir: 0})
						break
					}
				}

				//botton
				for q := i + 1; q < height; q++ {
					if arr[q][j] != '.' {
						graph[p] = append(graph[p], Pair{x: j, y: q, dir: 1})
						break
					}
				}
			}
			p++
		}
	}

	for i := 0; i < len(graph); i++ {
		k := graph[i][0].x
		j := graph[i][0].y
		//if current visiting node is not '.'
		f := true
		if len(graph[i]) > 1 {
			x1, x2, y1, y2 := 0, 0, 0, 0
			if graph[i][1].dir == 0 { //right
				x1, y1 = graph[i][1].x, graph[i][1].y
				if len(graph[i]) > 2 {
					x2, y2 = graph[i][2].x, graph[i][2].y
				} else {
					x2, y2 = -1, -1
				}
			} else { //down
				x2, y2 = graph[i][1].x, graph[i][1].y
				if len(graph[i]) > 2 {
					x1, y1 = graph[i][2].x, graph[i][2].y
				} else {
					x1, y1 = -1, -1
				}
			}
			if arr[j][k] == '.' {
				f = false
			}
			if f == true {
				fmt.Println(j, k, x1, y1, x2, y2)
			}
		} else {
			if arr[j][k] == '.' {
				f = false
			}
			if f == true {
				fmt.Println(j, k, -1, -1, -1, -1)
			}
		}
	}
}
