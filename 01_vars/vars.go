/*
vars

__author__ = "chiara"
__version__ = "01.01"
__date__ = "2020-06-19"
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// variable declaration & assignment init value
	var age int = 16

	/* if is not assigned any value,
	automatically initialized with the zero value of the variable's type

	if the variable is declared using the following syntax
	var age = 16
	it's possible to omit the type
	*/

	fmt.Println("my age is", age)

	//declaring multiple variables

	//same type
	var width, height int
	fmt.Println("width is", width, "height is", height)

	width = 100
	height = 50
	fmt.Println("new width is", width, "new height is", height)

	// different types
	var (
		name  = "Ahri"
		year  = 1970
		alive bool
	)
	fmt.Println("name: ", name, "born: ", year, "alive: ", alive)

	// short hand declaration

	count := 10
	fmt.Println("count = ", count)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
}
