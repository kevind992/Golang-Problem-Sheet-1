//Author: Kevin Delassus
//Date: 27/09/17
//Source: https://adriann.github.io/programming_problems.html


package main

import (
	//"bufio"
	"fmt"
	//"os"
	"math/rand"
)

func main() {

	var i int
	var ranNum int
	
	ranNum = rand.Intn(10)

	fmt.Print(ranNum)
   
	fmt.Println("Please Enter a Single Number between 1 & 10: ")
	fmt.Println("---------------------")
	fmt.Print("-> ")

	
    fmt.Scan(&i)
    //fmt.Println("read number", i, "from stdin")

	if i > 10{
		fmt.Println("number is greater then 10: ", i)
	}else{
		if ranNum == i{
			fmt.Println("Correct Answer!!!")
		}else{
			fmt.Println("Wrong Answer")
		}

	}

	
}