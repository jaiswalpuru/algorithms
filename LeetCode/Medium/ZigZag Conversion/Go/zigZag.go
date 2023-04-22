package main

import (
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	size := len(s)
	sections := int(math.Ceil(float64(size) / float64(2*numRows-2)))
	numCols := sections * (numRows - 1)

	zigZagTable := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		zigZagTable[i] = make([]byte, numCols)
	}

	currRow, currCol := 0, 0
	index := 0

	for index < size {
		//down
		for currRow < numRows && index < size {
			zigZagTable[currRow][currCol] = s[index]
			index++
			currRow++
		}

		currRow -= 2
		currCol += 1

		// up diagonally
		for currRow > 0 && currCol < numCols && index < size {
			zigZagTable[currRow][currCol] = s[index]
			index++
			currRow--
			currCol++
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if zigZagTable[i][j] != '\u0000' {
				res += string(zigZagTable[i][j])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
