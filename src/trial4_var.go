package main

import "fmt"

/*
Basic Type
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var i, j int = 1, 2

// Outside the func, "var" should be declared!!
//t =3 		// error

func main() {
	// type declaration can be omitted if initialized,
	// type declaration requires "var"
	var m, n = 1, 2
	var str, bo = "me", true

	//t = 3 	// undefined error
	// short assignemnt. "var" can be ommitted
	c := 3

	fmt.Println(i, j, m, n, str, bo, c)
	//fmt.Println(t)

	// Types
	const mxInt uint64 = 1<<64 - 1
	fmt.Printf("Type: %T	Value: %v", mxInt, mxInt)
}
