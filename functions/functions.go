package functions

import "fmt"

func FunctionsTest() {
	testMultipleReturns()
}

func testMultipleReturns() {
	fmt.Println("--- Testing multiple returns ---")
	var result1, result2 = multipleReturn()
	fmt.Println("First return: ", result1, " second return: ", result2)
}

func multipleReturn() (string, int) {
	return "hello", 0
}
