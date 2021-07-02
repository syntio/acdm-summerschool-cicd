package main

import (
	"fmt"
	"syntio.com/dinosstuph/functions"
)

func main() {
	fmt.Println("Hello, world!")

	fmt.Print("Enter the first number: ")
	var x int
	fmt.Scanf("%d\n", &x)
	fmt.Println()

	fmt.Print("Enter the second number: ")
	var y int
	fmt.Scanf("%d\n", &y)
	fmt.Println()

	fmt.Print("Sum of ", x, " and ", y, " is: ", functions.Sum(x, y))
}
