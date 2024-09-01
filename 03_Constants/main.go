package main

import "fmt"

// you can declear const and variable outside function.

const age int32 = 32

var lang string = "go lang"

// We can't declear variable using :- outside fucntion

func main() {
	const name = "golang"
	// you can't create const like :=
	fmt.Println(name)

	// printing function outside function
	fmt.Println(age)
	fmt.Println(lang)

	// assign multiple consts
	const (
		port     = 8080
		hostname = "webserver"
	)

	fmt.Println(port, hostname)
}
