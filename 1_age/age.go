/*
age

__author__ = "chiara@sabaini.com"
__version__ = "01.01"
__date__ = "2020-06-19"
*/

package main

import "fmt"

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
}
