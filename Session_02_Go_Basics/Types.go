package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	//Stastically typed language

	//a = b  We cannot assign float to int
	//The type of variable cannot be changed

	a = int(b)
	fmt.Println(a, b)

	var x int
	//x = "5" //compile time error
	_ = x

	//There is no un initalized variables
	//For numeric it is 0, string empty, pointer null

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
