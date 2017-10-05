//Question 3
//Coder: Kevin Delassus
//Date: 23-09-2017
//FizzBuzz - Write a program that prints the numbers from 1 to 100,
//each on a new line, to the console, except for the following conditions.
//For multiples of three print Fizz instead of the number,
//and for multiples of five print Buzz. For numbers that are multiples
//of both three and five print FizzBuzz.

package main

import (
	"fmt"
)

func main() {

	//Creating two Strings "Fizz" & "Buzz"
	fizz := "Fizz"
	buzz := "Buzz"
	
	//For loop going from 0 -> 100
	for i := 0; i < 100; i++ {
		
		
		switch{
		//Checking is i is a multible of 3 or 5
		case i % 15 == 0:
			
			fmt.Println(fizz + " " + buzz)
		
		//Checking if i is a multible of 3	
		case i % 3 == 0:
		
			fmt.Println(buzz)
		//Checking if i is a multible of 5	
		case i % 5 == 0:
		
			fmt.Println(fizz)
		//If i isnt a multible of 5 or 3, the number is displayed	
		default:
			fmt.Println(i)	
		}	
		
	}
	fmt.Println(" ")
	fmt.Println("Fizz Buzz Complete")

}
