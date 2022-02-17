//Khai bao bien
//constant

package main

import "fmt"

func main(){
	var i int
	i = 42
	var j int = 54
	k := 11
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v, %T", k, k)

	//constant
	const a int16 = 42
	fmt.Println(a)


	//const iota
	const(
		//errorDefault = iota
		//_ = iota
		_ = iota + 5
		x
		y
		z
		t
	)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", t, t)


}



