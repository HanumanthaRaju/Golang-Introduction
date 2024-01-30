package main

import "fmt"

func main() {
	car, cost := "Audi", 50000000
	fmt.Println(car, cost)

	//car, cost := "BMW", 10000000  //Error
	car, year := "BMW", 10000000 //atleast one variable should be new
	_ = year                     // to make the unused variable mute

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	_, _ = i, j //to mute if un used

	//Swap two variables using assignemnt
	j, i = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3 //Expression in a short declaration
	fmt.Println(sum)

}
