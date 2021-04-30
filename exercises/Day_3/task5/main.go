package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	clearTerminal()
	fmt.Print("Hi there! Please enter three names.\n")
	getNames()
}

func getNames() {
	fmt.Print("1.Enter the first person's name: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("2.Enter the second person's name: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print("3.Enter the third person's name: ")
	var third string
	fmt.Scanln(&third)
	clearTerminal()
	fmt.Print("There are 3 people in classroom!")
	fmt.Printf("\nHello %v!\nHi %v and %v!\n", first, second, third)
}

func clearTerminal() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
