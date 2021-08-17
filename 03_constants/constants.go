/*
constants

__author__ = "chiara"
__version__ = "01.01"
__date__ = "2020-06-20"
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// constants
	/*cannot be reassigned again to any other value,
	hence the program will not work
	and it will fail with compilation error: cannot assign to
	*/

	const a = 55
	// a = 89 -> reassignment not allowed

	/*value of a constant should be known at compile time,
	hence it cannot be assigned to a value returned by a function call
	since the function call takes place at run time.
	*/
	var a = math.Sqrt(4)   // allowed
	const b = math.Sqrt(4) // not allowed

	// str const

	/*they are untyped (don't have any type)
	 */
	const hello = "Hello World"
	fmt.Printf("type %T value %v", hello, hello)

	// that's how you can create a typed const
	const typedhello string = "Hello World"

	// bool const
	/*same rules for string constants apply to booleans
	 */
	const trueConst = true
	type myBool bool                  // myBool is an alias for bool
	var defaultBool = trueConst       // allowed
	var customBool myBool = trueConst // allowed
	defaultBool = customBool          // not allowed

	// num const
	/*include integers, floats and complex constants.
	default type of these kind of constants
	can be thought of as being generated on the fly
	depending on the context.
	var intVar int = a requires a to be int so it becomes an int constant.
	var complex64Var complex64 = a requires a to be a complex number and hence it becomes a complex constant
	*/
	const i = 5
	var intVar int = i
	var int32Var int32 = i
	var float64Var float64 = i
	var complex64Var complex64 = i
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// numeric expressions
	var g = 5.9 / 8
	fmt.Printf("a's type %T value %v", g, g)
}
