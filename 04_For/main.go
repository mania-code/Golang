package main

import "fmt"

// for is only construct in go for looping.
func main() {

	i := 1
	// while loop

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop

	// for {
	// 	fmt.Println(i)
	// }

	// classic for loop

	// for i := 0; i <= 3; i++ {
	// 	// break

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// go-varion 1.22 we have new concept of range
	// for i := range 11 {
	// 	fmt.Println(i)
	// }

}
