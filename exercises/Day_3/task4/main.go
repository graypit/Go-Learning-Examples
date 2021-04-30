package main

import "fmt"

func main() {
	a := int8(127)  // don't change this line
	b := int16(513) // don't change this line

	fmt.Println(int8(b) + a)
	fmt.Println(b + int16(a))
}
