package main

import (
	"fmt"
	"math/rand"
)

func condition_if_1() {
	n := rand.Intn(10)
	fmt.Println("Number is ", n)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's a good number: ", n)
	}
}

func condition_if_2() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Number is ", n)
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("Number is ", n)
		fmt.Println("That's too big")
	} else {
		fmt.Println("Number is ", n)
		fmt.Println("That's a good number: ", n)
	}
}

// func condition_err() {
// 	if n := rand.Intn(10); n == 0 {
// 		fmt.Println("That's too low")
// 	} else if n > 5 {
// 		fmt.Println("That's too big")
// 	} else {
// 		fmt.Println("That's a good number: ", n)
// 	}
// 	fmt.Println("Number is ", n) // undeclared name: n
// }

func main() {
	fmt.Print("test1: ")
	condition_if_1()
	fmt.Print("test2: ")
	condition_if_2()
}
