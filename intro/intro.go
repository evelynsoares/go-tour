package main // every program is made of packages

import (
	"fmt"
	"math/rand" // the package name is the same as the last element of the import path -> rand
	"math" // import again?
	"time"
	"math/cmplx"
)

const E = 2.6

// type ommited because i declared with an initializer
var cpp, python, golang, solidity = true, false, "no", "yes"

// the type comes after the variable name, after () comes the return type
// type of the first can be omited if its the same type
func add(x, y int) int {
	return x + y
}

func convertTypes(x, y float64, z uint) (int, int, int) {
	var a int = int(x * y) // did 23.6*0.9=21.24 then converted the result to 21 
	b := int(z-1000)
	c := int(math.Sqrt(float64(z))) // 3.3333 converted to int
	return a, b, c
}

func swap(x, y string) (string, string) {
	return y, x
}

// only use naked returns in short functions for better readability
func returnNaked(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe 	bool = false
	MaxInt 	uint64 = 1<<64 - 1 // biggest possible value of a uint64: 1<<64 = 2^64
	z 		complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big  = 1 << 100 //shifts the bit 1 left of 100 0's
	Small = Big >> 99 // shifts Big right of 99 0's  
)

func needInt(x int) int {
	return x*10+1
}

func needFloat(x float64) float64 {
	return x*0.1
}

func main() {

	var i, j int = 1, 2
	// short assigned state that can be used in place of a var declaration with implicit type
	// not available outside a function
	k := 3 
	fmt.Println(i, cpp, python, golang, solidity) // will print variables of different types
	

	fmt.Println(returnNaked(10))
	fmt.Println("My fav number is", rand.Intn(10))
	fmt.Println(math.Pi) // a name is exported if it begins with a capital letter
	fmt.Println("The time is", time.Now())
	fmt.Printf("Now you have %g problems. -> has to be done with Printf\n", math.Sqrt(7))

	fmt.Println(add(20, 22))

	a, b := swap("first", "second")
	fmt.Println(a, b)
	fmt.Println(swap("swap", "this"))
	
	fmt.Println(i, j, k)

	cs, javascript, typescript := "love", "dislike", "like"
	fmt.Println(cs, javascript, typescript)

	// %t and %T display different things-> for z: %t = %!t(complex128=(2+3i)) and %T = complex128
	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)

	fmt.Println(convertTypes(23.6, 0.9, 10))

	fmt.Printf("%g is a constant\n", E)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}