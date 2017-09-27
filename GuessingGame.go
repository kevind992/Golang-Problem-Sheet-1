//Author: Kevin Delassus
//Date: 27/09/17
//Source: https://adriann.github.io/programming_problems.html


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please Enter a Single Number: ")
	fmt.Println("---------------------")
	fmt.Print("-> ")
	number, _ := reader.ReadString('\n')

	fmt.Println(number)
}