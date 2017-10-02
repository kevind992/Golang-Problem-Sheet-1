//Coder: Kevin Delassus
//Date: 26/09/17
//Discription: Write a function that calculates the sum of the digits of the factorial of a number.
// n! means n x (n − 1) ... x 3 x 2 x 1. 
//For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, 
//and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
//Then find the sum of the digits in the number 100!.


package main

import (
	"fmt"
	"math/big"
)

func main() {

	
	fmt.Println(factorial(10)) //100 can be changed to any number
	sum := big.NewInt(0)
	n := new(big.Int) //to store the factorial
	fact := (factorial(100))

	//func (z *Int) DivMod(x,y,m*Int)(*Int, *Int)
	//sets z to quotient x div y and m to modulus x mod y and returns
	//the pair (z,m) where y is not 0.
	
	for fact.BitLen() > 0 {
		_, n := fact.DivMod(fact, new(big.Int).SetUint64(uint64(10)), n)
		sum = sum.Add(sum, n)
	}
	//_, acts like a place holder. BitLen breaks the result into individual digits
	fmt.Println(sum)
}
func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))

}
