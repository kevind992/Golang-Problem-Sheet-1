//Author: Kevin Delassus
//Date: 29/09/17
//Write a function that merges two sorted lists into a new sorted list,
// e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
//Source: https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

package main

import (
	"fmt"
	"sort"
)

func main() {

	
	//Calling append merges 2 lists together
	s := append([]int{1, 4, 3, 7, 50}, []int{5, 2, 6, 8, 11}...)

	//sort.Ints() is sorting the conbined List
	sort.Ints(s);

	//Printing the result which is the 2 lists combined
	fmt.Println("Here is the sorted Lists ", s )
}