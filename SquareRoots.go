//Question 9
//Author: Kevin Delassus
//Date: 27-09-2017
//Write a function to calculate the square root of a number using Newton's method.
//Newton's method is to approximate the square root of a number x by picking a starting 
//point z and then repeating the following operation.
//z_next = z - ((z*z - x) / (2 * z))

package main

import (
	//Importing the print and math classes
	"fmt"
	"math"
)

//Function calculating the square root using Newtons formula
func sqrt(x float64, z float64) float64{

	var z_next float64 = z - (((z*z) - x) / (2 * z))
	
	//Returning the answer
	return z_next 
}

func main() {
	//Setting Variables
	x := 100.0
	var z float64
	
	//For loop calculating the answers and displaying them to the console
	for z = 1.0; z != sqrt(x , z); z = sqrt(x , z){

		fmt.Println("Newtons Solutions are: ",z)
	}

	fmt.Println("")
	//Displaying the Square root using the SQRT() function throught .Math
	fmt.Println("Solution using the SquareRoot Function: ",math.Sqrt(x))
}
