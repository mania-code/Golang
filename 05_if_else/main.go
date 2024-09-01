package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }

	//else-if

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermission = false

	// if role == "admin" || hasPermission {
	// 	fmt.Println("Yes")
	// }
	// if role == "admin" && hasPermission {
	// 	fmt.Println("Yes")
	// }

	// we can declar a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age > 12 {
		fmt.Println("Person is teenager", age)
	}

	// go does not have a ternary operator, you have to use normal if else

}
