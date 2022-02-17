package main

import (
	"fmt"
)

func main(){
	//points := [3]int{7,10,5}
	//points := [...]int{7,10,5}
	var points [3]int

	points[0] = 1

	fmt.Printf("%v, %T\n", points, points)
	fmt.Println(len(points))

	/*
	a := [3]int{1, 2, 3}
	//b := a
	b := &a // giong con tro
	b[2] = 10
	fmt.Println(a)
	fmt.Println(b)
	*/

	//array slice <cap phat dong cho mang>
	x := []int{1, 2, 3}
	y := x
	y[2] = 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(cap(x)) //capacity

	//thao tac mang
	//cat mang
	a := []int {1, 2, 3, 4, 5, 6, 7}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:7]
	fmt.Printf("a %v, %v, %v\n", a, len(a), cap(a))
	fmt.Printf("b %v, %v, %v\n", b, len(b), cap(b))
	fmt.Printf("c %v, %v, %v\n", c, len(c), cap(c))
	fmt.Printf("d %v, %v, %v\n", d, len(d), cap(d))
	fmt.Printf("e %v, %v, %v\n", e, len(e), cap(e))

	//make
	//v := make([]int, 10, 10) // 10: len; 20: cap
	v := make([]int, 10, 11) // 10: len; 20:
	v = append(v, 1)
	fmt.Printf("v %v, %v, %v\n", v, len(v), cap(v))
	//v = append(v, 1, 2, 3)
	//fmt.Printf("v %v, %v, %v\n", v, len(v), cap(v)) // ????

	//append slice cho trc vao 1 slice khac - bat buoc them ... vao cuoi
	v = append(v, []int{1, 2, 3}...)
	fmt.Printf("v %v, %v, %v\n", v, len(v), cap(v))

	//ngan xep
	m := []int{1, 2, 3, 4, 5, 6}
	//n := m[:5]
	fmt.Println(m)
	n := append(m[:2], m[3:]...)
	fmt.Println(m)
	fmt.Println(n)

}
