package main

import "fmt"

func add(a int, b int) int {
	var out = a + b
	return out
}

func calc(a int, b int) (int, int) {
	var out1 = a + b
	var out2 = a - b
	return out1, out2
}
func main() {
	num1 := 5
	num2 := 4

	result := add(num1, num2)
	result1, result2 := calc(num1, num2)
	fmt.Print(result)
	fmt.Print(result1, result2)
}