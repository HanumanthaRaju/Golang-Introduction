package main

import (
	"fmt"
	"runtime"
)

func f1() {
	fmt.Println("f1 (go routine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println(("f1 (go routine) exection stopped"))
}

func f2() {
	fmt.Println("f2 (go routine) execution started")
	for i := 5; i < 6; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println(("f2 (go routine) exection stopped"))
}
func main() {
	fmt.Println("No. of CPUs:", runtime.NumCPU())             // => No. of CPUs:
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1
	fmt.Println("OS:", runtime.GOOS)                          // => OS: windows
	fmt.Println("Arch:", runtime.GOARCH)                      // => Arch: amd64
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))         // => GOMAXPROCS: 8

	go f1()
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())

	f2()

	//time.sleep(time.second * 2)
	fmt.Println("main execution stopped")
}
