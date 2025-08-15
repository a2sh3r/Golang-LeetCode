package main

import "fmt"

func generate(numRows int) [][]int {
	switch numRows {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	}

	res := [][]int{{1}, {1, 1}}

	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1

		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, temp)
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
