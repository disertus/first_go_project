package main

import (
	"fmt"
	"time"
	"math"
)

//declaring functions outside of the main function
func add(x int, y int) int {  //explainer: types are declared after the variable. output type outside arguments
	return x + y
}

func substract(x, y int) int{ //if both arguments are of same type, one can put the type at the end without repeating
	return x - y
}

func main() {
	fmt.Println("Hello, world")
	fmt.Println("The time is", time.Now()) //string concatenation example
	fmt.Printf("Now I have %g problems\n", math.Sqrt(7))  //formatted string example
	fmt.Println(add(150, 450)) //example of outer function execution with two positional arguments
	fmt.Println(substract(1000, 933))
}