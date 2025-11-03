package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var c, python, java bool // Variables without initializers
var i, j int = 1, 2      // Variables with initializers

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

// Go's basic types are

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// Type: bool Value: false
	// Type: uint64 Value: 18446744073709551615
	// Type: complex128 Value: (2+3i)
}

// Variables declared without an explicit initial value are given their zero value.
// The zero value is:
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
func ZeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Type conversions
// The expression T(v) converts the value v to the type T.

// Some numeric conversions:

// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// Or, put more simply:

// i := 42
// f := float64(i)
// u := uint(f)
// Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

func TypeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// Type inference
// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

// When the right hand side of the declaration is typed, the new variable is of that same type:

// var i int
// j := i // j is an int
// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128
// Try changing the initial value of v in the example code and observe how its type is affected.

func TypeInference() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

// Constants
// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.

const Pi = 3.14

func Constants() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// Numeric Constants
// Numeric constants are high-precision values.

// An untyped constant takes the type needed by its context.

// Try printing needInt(Big) too.

// (An int can store at maximum a 64-bit integer, and sometimes less.)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func NumericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	a, b := swap("hello", "world") // multi return
	fmt.Println(a, b)
	fmt.Println(split(17)) // named return
	var i int              // Variables without initializers
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!" // Variables with initializers
	fmt.Println(i, j, c, python, java)
	k := 3 // Short variable declarations
	fmt.Println(k)
	types() // Go's basic types
	ZeroValues()
	TypeConversions()
	TypeInference()
	Constants()
	NumericConstants()
}
