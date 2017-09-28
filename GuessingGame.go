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
	count := 0
	
	ranNum = rand.Intn(10)

	fmt.Print(ranNum)
   
	fmt.Println("Please Enter a Single Number between 1 & 100: ")
	fmt.Println("---------------------")

	for i != ranNum{

		fmt.Print("-> ")
		fmt.Scan(&i)
		
		if i > ranNum{
			fmt.Println("You are to High!")
			count++
		}
		if i < ranNum{

			fmt.Println("You are to Low!")
			count++
		}
		if i == ranNum{
			break
		}
	}
	fmt.Println("Correct You Win")
	CountTimes(count)

	
}

func CountTimes(c int){
	if c == 0{
		fmt.Println("Woow you guessed on the first go!! WELL DONE!!")
	}else{
		fmt.Println("You Guessed ", c," times!!")
	}
	
}