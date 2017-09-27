//Coder: Kevin Delassus
//Date: 27-09-2017
//FizzBuzz - Question 9

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64, z float64) float64{

	var z_next float64 = z - (((z*z) - x) / (2 * z))

	return z_next 
}

func main() {

	x := 100.0
	var z float64
	

	for z = 1.0; z != sqrt(x , z); z = sqrt(x , z){

		fmt.Println("Newtons Solutions are: ",z)
	}

	fmt.Println("")

	fmt.Println("Solution using the SquareRoot Function: ",math.Sqrt(x))
}
