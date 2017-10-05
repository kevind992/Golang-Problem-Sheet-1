//Question 4
//Author: Kevin Delassus
//Date: 04/10/17
//Write a function that calculates the sum of the digits of the factorial of a number.
//n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
//and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//Then find the sum of the digits in the number 100!.
//Adapted From: https://github.com/chirag04/golang-project_euler/blob/master/p020.go


package main

import (

	"fmt"
)
 
func main(){

	//Setting Variables
	var answer int

	//Change this value to get a different answer
	check := 100
	
	//Running calFactorial funtion
    answer = calFactorial(check)

	//Printing out answer
	fmt.Print("The Factorial sum of ",check ," is: ", answer)
}

func calFactorial(num int) int {
	
	//Setting variables
    sum := 0;
    
    nums := [200] int{};
    nums[0] = 1;
	
	//double for loop calculating the factorial sum and putting the individual digits into a slice
    for i := 2; i <= num; i++ {
        
    	for x := 0; x < len(nums); x++ {
            
            nums[x] *= i;
            //
    		if x > 0 && nums[x - 1] > 9 {
    			nums[x] += int(nums[x - 1] / 10);
    			nums[x - 1] %= 10;
    		}
    	}
	}
	
	//for adding all the numbers together to get final answer
    for i := 0; i < len(nums); i++ {
    	sum += nums[i];
	}
	//Returning the answer
    return sum;
}

