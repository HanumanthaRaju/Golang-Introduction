package main

import "fmt"

func main() {
	//var age int = 30
	var age = 30
	fmt.Println("Age:", age)

	var name = "Tejas"
	fmt.Println("Your Name is:", name)

	var name1 = "Networks"
	fmt.Println("Your Name is:", name1) //Error, if you comment this

	var name2 = "Bangalore"
	_ = name2 //For unused variable

	//name3 := "Tejas Networks" //you cannot have a unused variable

	s := "Learning golang" //You cannnot use := for already decalred variable
	fmt.Println(s)

	//name2 := "Tejas Networks Limited" //no new variables on left side of :=

}
