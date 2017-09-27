//Coder: Kevin Delassus
//Date: 27-09-2017
//FizzBuzz - Question o

package main

import (
	"fmt"
	"math"
)

func calcutate(x float64, z float64) float64{
	z_next := z - (((z*z) - x) / (2 * z))

	return z_next
}

func main() {

	z := 20
	var x float64


	for x = 1.0, x != calcutate( x , z ), z = calcutate( x , z ){

		fmt.Println("Current Guess" + x)
	}

	fmt.Println("")

	fmt.Println("Result using the .SquareRoot" + z.getSQuareRoot())
}
