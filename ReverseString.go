//Question 10
//Author: Kevin Delassus
//Date: 02/10/17
//Write a function to reverse a string in Go.	
//Adapted from: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main

import(
	"fmt"
)

func main(){
	
	//String which needs to be reversed
	string := "Kevin"

	//Calling Funstion and Displaying Answer
	fmt.Println(Reverse(string))

}

//Function which reverces the string
func Reverse(s string) string {

    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}