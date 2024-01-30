/////////////////////////////////
// The Typical Structure of a Go Application
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
// fmt is a standard library package used to print to standard output (console)
import "fmt"

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

	// Local Scoped Variables and Constants Declarations, calling functions etc
	var age int = 17
	var today string = "Monday"

	// Println() function prints out a line to stdout
	// It belongs to package fmt
	fmt.Println("Today is", today) // => Today is Monday
	fmt.Println("You age is", age) // => You age is 17

}
