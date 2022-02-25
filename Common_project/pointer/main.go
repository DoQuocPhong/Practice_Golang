package main

import (
	"fmt"
)

func applyPointer(pointer *int) {
	*pointer = 777
}

func main() {
	a := 100

	var pointer *int // nil neu zero value

	pointer = &a
	fmt.Println(a)

	*pointer = 999 // tuong duong a := 999

	fmt.Println(a)

	fmt.Println(pointer)

	// b :=123
	// p2 := new(int) // tra ra gia tri bo nho neu zero value

	// p2 = &b

	// fmt.Println(p2)

	//pointer -> array

	array := [3]int{1, 2, 3}
	var demoPointer *[3]int

	demoPointer = &array

	fmt.Println(demoPointer)

	c := 30
	var pointer2 *int = &c

	applyPointer(pointer2)

	fmt.Println(c)
}
