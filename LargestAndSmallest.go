//Coder: Kevin Delassus
//Date: 03/10/2017
//Write a function that returns the largest and smallest elements in a list.
//Sources: https://gist.github.com/pyk/10519339

package main

import "fmt"

func main() {
  var n, smallest, biggest int

  array := []int{
    47,96,3,68,
    57,1,63,70,
    37,34,83,27,
    19,100, 8,50,
  }

  fmt.Println("This Program Displays the Largest and Smallest Numbers in a List")
  fmt.Println("=================================================================")
  
  //For loop checking for biggest number
  for _,v:=range array {
    if v > n {
      n = v
      biggest = n
    } else {}
  }

  fmt.Println("The Biggest number is: ", biggest)//Displaying Biggest Number

  //For loop checking for smallest number
  for _,v:=range array {

    if v > n {
    } else {
      n = v
      smallest = n
    }
  }
  fmt.Println("The Smallest number is: ", smallest) //Displaying Smallest Number
}
