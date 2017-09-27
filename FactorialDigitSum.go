//Coder: Kevin Delassus
//Date: 26/09/17
//Discription: Write a function that calculates the sum of the digits of the factorial of a number.
// n! means n x (n − 1) ... x 3 x 2 x 1. 
//For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, 
//and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
//Then find the sum of the digits in the number 100!.


package main

import (
	"fmt"
)

func main() {

	amount := 10
	total := 0
	sum := 0

	for i := amount; i > 0; i-- {

		sum = amount - 1
		total = amount * sum 
	}

	total = amount * sum

	fmt.Println(total)
}
