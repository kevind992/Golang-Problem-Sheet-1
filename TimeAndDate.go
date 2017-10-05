//Question 2
//Author: Kevin Delassus
//Date: 21/09/2017
//Write a program that prints the current time and date to the console.
//Adapted from https://tour.golang.org/welcome/4

package main

import (
	//Importing the format and time package	
	"fmt"
	"time"
)

func main() {
	
	//Printing the time and date to the console
	fmt.Println("The time is", time.Now(),"\n")	
}
