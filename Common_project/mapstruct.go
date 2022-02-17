package main

import (
	"fmt"
	"reflect"
)

/*
type Student struct {
	number 		int
	name 		string
	isMale 		bool
	subjects 	[]string
}
*/
type People struct {
	name string
	age int
}

type Student struct {
	People
	number int `Maximum cant over 100`
	subject [] string
}

func main(){
	studentNameAgeMap := map[string]int {
		"A" : 1,
		"B" : 2,
		"C" : 3,
	}
	fmt.Println(studentNameAgeMap["A"])

	copyMap := studentNameAgeMap
	fmt.Println(copyMap)

	x := studentNameAgeMap["B"]
	fmt.Println(x)

	//struct
	/*
	studentA := Student{
		number: 	1,
		name: 		"A",
		isMale:		true,
		subjects:	[]string {
			"Math",
			"English",
			"IT",
		},
	}
	*/

	studentA := Student{}
	studentA.name = "A"
	studentA.number = 1
	//studentA.isMale = true
	//studentA.subjects = []string{"Math", "English", "IT"}


	t := reflect.TypeOf(Student{})
	result, _ := t.FieldByName("number")
	fmt.Println(result)


	fmt.Println(studentA)






}