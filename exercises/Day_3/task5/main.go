package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	clearTerminal()
	fmt.Print("Hi there! Please enter the 3 names of your co-workers!\n")
	getNames()
}

func getNames() {
	fmt.Print("Enter the first person's name:")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter the second person's name:")
	var second string
	fmt.Scanln(&second)
	fmt.Print("Enter the third person's name:")
	var third string
	fmt.Scanln(&third)
	clearTerminal()
	fmt.Print("There are 3 people in classroom!")
	fmt.Printf("\nHello %v!\nHi %v and %v!", first, second, third)
}

func clearTerminal() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
