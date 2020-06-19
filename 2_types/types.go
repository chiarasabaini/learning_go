/*
types

__author__ = "chiara@sabaini.com"
__version__ = "01.01"
__date__ = "2020-06-19"
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b // and
	fmt.Println("c:", c)

	d := a || b // or
	fmt.Println("d:", d)

	// signed integers (negatives and positives)
	var f int8 = -1  // 8 bit
	var g int16 = -2 // 16 bit
	var h int32 = -3 // 32 bit
	var i int64 = -4 // 64 int
	var j int = -5   // 32 or 64 bit, depends on the system, use this tho represent integers unless there is a need to use a specific sized integer

	fmt.Println(f, g, h, i, j)

	// %T is the format specifier to print the type, %d is used to print the size
	fmt.Printf("type of f is %T, size of f is %d", f, unsafe.Sizeof(a))   //type and size of f
	fmt.Printf("\ntype of g is %T, size of g is %d", g, unsafe.Sizeof(b)) //type and size of g

	// unsigned integers (only positives)
	var k uint8 = 1  // 8 bit
	var l uint16 = 2 // 16 bit
	var m uint32 = 3 // 32 bit
	var n uint64 = 4 // 64 bit
	var o uint = 5   // 32 or 64 bit, depends on the system

	fmt.Println(k, l, m, n, o)

	// floating points
	var p float32 = 4.20 // 32 bit
	var q float64 = 6.90 // 64 bit, default

	ftm.Println(p, q)

	// complex
	/*
		builtin function complex is used to construct a complex number
		with real and imaginary parts

		func complex(r, i FloatType) ComplexType

		complex64: complex numbers which have float32 real and imaginary parts
		complex128: complex numbers with float64 real and imaginary parts
	*/

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// string -> collection of bytes
	var str string = "abcd"
	fmt.Println(str)

	/* others
	byte alias of uint8
	rune alias of int32
	*/
}
