package main

import "fmt"

func main() {
	//fmt.Println("Hello")
	//wont compile
	var x int = 5
	var y int = 10
	var sum int = x + y
	//emit type
	z := 15
	sum = x + z
	fmt.Println(sum)
	//if else statements
	if x < 6 {
		fmt.Println("Small number")
	} else {
		fmt.Println(" big number")

	}
	//arrays
	var a [5]int

	a[2] = 7
	fmt.Println(a)
	//slices
	s := []int{}
	s = append(s, 13)
	fmt.Println(s)
	// maps

	verticies := make(map[string]int)

	verticies["triangle"] = 2
	verticies["circle"] = 3
	fmt.Println(verticies)
	delete(verticies, "circle")
	fmt.Println(verticies)
	// loops
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//while loop
	b := 0
	for b < 5 {
		fmt.Println(b)
		b++
	}

	//iterate thru array or slice

	for key, value := range verticies {
		fmt.Println("key", key, "value", value)
	}
	//val := sum(2, 3)
	//fmt.Println(val)

}
func sum(x int, y int) int {
	return x + y

}
