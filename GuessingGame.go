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

var entry [10]int

func main() {

	var i int
	var ranNum int
	count := 0
	check := true

	
	
	ranNum = rand.Intn(100)

	fmt.Print(ranNum)
   
	fmt.Println("Please Enter a Single Number between 1 & 100: ")
	fmt.Println("---------------------")

	for i != ranNum{

		fmt.Print("-> ")
		fmt.Scan(&i)
		
		
		if i > ranNum{

			check = CheckArray(i,count)

			if check == true{
				fmt.Println("You are to High!")
				
				//fmt.Println(entry[count])
				count++
			}
		}
		if i < ranNum{

			 check = CheckArray(i,count)

			if check == true{
				fmt.Println("You are to Low!")
				
				//fmt.Println(entry[count])
				count++
			}
			
		}
		if i == ranNum{
			break
		}
	}
	fmt.Println("Correct You Win")
	CountTimes(count)

	
}
func CheckArray(i int, guess int){

	entry[i] = guess

	for x := 0; x < count; x++{
		if entry[i] == guess{
			return false
		}else {return true}
	}
}
func CountTimes(c int){
	if c == 0{
		fmt.Println("Woow you guessed correct on the first go!! WELL DONE!!")
	}else{
		fmt.Println("You Guessed ", c," times!!")
	}
	
}