package main

import (
	"fmt"
	"strconv"
)

var (
	actorName string = "Cool Actor"
	companion string = "Cool Companion"
	season    int    = 3
	episode   int    = 10
)

func main() {
	fmt.Println(actorName)

	//Conversion
	var i int = 42
	var j float32
	j = float32(i)
	fmt.Println(j)

	var x string
	x = strconv.Itoa(i)
	fmt.Println(x)
	constants()
}

func primitives() {
	// Boolean, Int, Float, Complex, String, Rune

	// Booleans
	var n bool
	x := 1 == 1
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", x, x)

	// Numeric
	// Integer (int8, int16, int32, int64, uint8, uint16, uint32)
	var m uint = 42
	fmt.Printf("%v, %T\n", m, m)

	//Binary ops
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	//Bit shifting
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0

	// float
	z := 3.14
	z = 13.7e72
	z = 2.1e14
	fmt.Printf("%v, %T\n", z, z)

	// complex
	var t complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(t), real(t))
	fmt.Printf("%v, %T\n", imag(t), imag(t))

	// string
	s := "this is a string"
	d := []byte(s)
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	fmt.Printf("%v, %T\n", d, d)

	// rune
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}

func constants() {
	const myConst int = 42
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const (
		isAdmin = 1 << iota
		isHQ
		canSeeFinancials
		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNA
		canSeeSA
	)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Can see Africa? %v\n", canSeeAfrica&roles == canSeeAfrica)
}
