/*
convert

__author__ = "chiara"
__version__ = "01.01"
__date__ = "2020-06-19"
*/

package main

import "fmt"

func main() {
	i := 1
	j := 2.8

	sum := i + int(j) // int(j) -> converts j from float to int
	fmt.Println(sum)

	k := 10
	var l float64 = float64(k) // doesn't work without explicit conversion
	fmt.Println("l", l)
}
