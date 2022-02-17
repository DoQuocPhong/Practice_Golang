package main

import (
	"fmt"
)

func main(){
	m := map[string]int{
		"A": 1,
		"B": 2,
	}

	if number, isExist := m["A"]; isExist == true {
		fmt.Println(number)
	}
	//
	number := 90
	guess := 95
	if guess < 10 {
		fmt.Println("nho hon 10")
	}
	if guess > 100 {
		fmt.Println("lon hon 100")
	}
	if guess < number {
		fmt.Println("nho hon so co san")
	}
	if guess > number {
		fmt.Println("lon hon so co san")
	}

	// || vs &&
	if guess < 10 || guess > 100 {
		fmt.Println("nho hon 10 hoac lon hon 100")
	}

	//if else nhu cac ngon ngu khac


	//switch case
	exam := 90
	switch exam {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("Default")
	}
}
