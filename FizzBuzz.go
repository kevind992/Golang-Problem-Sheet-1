//Coder: Kevin Delassus
//Date: 23-09-2017
//FizzBuzz - Question 3

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
