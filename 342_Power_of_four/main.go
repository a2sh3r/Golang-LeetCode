package main

import (
	"fmt"
	"strconv"
)

// n&(n-1) проверяет на четность, а маска в контексте проверяет биты, в которых может быть четверка
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	if n&(n-1) != 0 {
		return false
	}
	return (n & 0b01010101010101010101010101010101) != 0
}

func main() {
	fmt.Printf("%s", strconv.FormatBool(isPowerOfFour(17)))
}
