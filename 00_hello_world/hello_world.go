/*
hello_world

__author__ = "chiara@sabaini.com"
__version__ = "01.01"
__date__ = "2020-06-19"
*/

/*
Every go file must start with the package name statement.
Packages are used to provide code compartmentalization and reusability.
*/
package main

import "fmt" // stdio in c++

func main() {
	// prints without adding a new line (needs \n)
	fmt.Printf("hello, world")
	fmt.Printf("hello, world\n")

	// print adding a new line
	fmt.Print("hello, world")
	fmt.Println("hello, world")
}
