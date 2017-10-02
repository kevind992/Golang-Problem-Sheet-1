//Coder: Kevin Delassus
//Date: 02/10/17
//Source: https://adriann.github.io/programming_problems.html
//		: http://www.golangpro.com/2016/01/check-string-palindrome-go.html

package main

import (
 "fmt"
 "strings"
)

func main() {

 var x string

 //User is entering a String to see if its a Palindrome
 fmt.Println("Enter string:")
 fmt.Scanf("%s\n", &x)
 x = strings.ToLower(x)
 fmt.Println(isP(x))
}

//Function to test if the entered string is a Palindrome
func isP(s string) string {
 mid := len(s) / 2
 last := len(s) - 1

 //Loop going through the letters
 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "This is not a Palimdrome."
  }
 }
 return "You have entered a Palindrome"
}