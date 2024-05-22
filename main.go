package main

import "fmt"

func main() {
	var num1, num2 int = printHelloWorld()
	fmt.Printf("The number 1 is %v and the number 2 is %v", num1, num2)
}

func printHelloWorld() (int, int) {
	fmt.Println("Hello, World!")
	return 1, 2
}
