package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func main() {
	// var name string = "hyeongseok"
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("hyeongseok")

	fmt.Println(totalLength, upperName)
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}
