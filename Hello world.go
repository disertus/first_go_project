package main

import (
	"fmt"
	"math"
	"time"
)

//declaring functions outside of the main function
func add(x int, y int) int { //explainer: types are declared after the variable. output type outside arguments
	return x + y
}

func subtract(x, y int) int { //if both arguments are of same type, one can put the type at the end without repeating
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum float32) (x, y float32) { //if the result is a float, arguments should be floats too
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

var c, python, java bool //declares a list of variables
//x := 15 //assignment with 'warlus' opeartor is impossible outside a function, should begin with 'var'

func main() {
	fmt.Println("Hello, world")                          //Function accepts only double quotes "..." for strings
	fmt.Println("The time is", time.Now())               //string concatenation example
	fmt.Printf("Now I have %g problems\n", math.Sqrt(7)) //formatted string example
	fmt.Println(add(150, 450))                           //example of outer function execution with two positional arguments
	fmt.Println(subtract(1000, 933))

	a, b := swap("hello", "world") //variables contain the respective values returned by swap()
	fmt.Println(a, b)

	fmt.Println(split(19))

	var i int
	fmt.Println(i, c, python, java) // c, python and java = bool variables, default for those is False. i = int, default is '0'
}
