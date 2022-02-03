package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926
const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)
	const (
		Zero Digit = iota //basic iota expression generating values starting with 1.
		One
		Two
		Three
		Four
	)
	// The above uses the constant IOTA GENERATOR, which yeilds the same code as below.
	// const (
	// 	Zero = 0
	// 	One = 1
	// 	Two = 2
	// 	Three = 3
	// 	Four = 4
	// 	)
	fmt.Println(One)
	fmt.Println(Two)
	// IOTA generator with missing values.
	// the underscore allows to skip values.

	// a more complex iota which starts with 1 and applies bitwise left movement
	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
