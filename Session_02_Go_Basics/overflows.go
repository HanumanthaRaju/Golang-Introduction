package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ //overflow, x is 0
	fmt.Println(x)

	// a := int8(255 + 1) //Overflow at compile time

	var b int8 = 127
	fmt.Println("%d\n", b+1)

	b = -128
	b--
	fmt.Println("%d\n", b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f = f * 1.2 //f overflows to infinite
	fmt.Println(f)

	// const i int8 = 300 //compile time error

}
