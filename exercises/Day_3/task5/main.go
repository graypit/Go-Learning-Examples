package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func check() {
	if len(os.Args) < 2 {
		fmt.Printf("You have to input people names as arguments\n")
		os.Exit(1)
	}
}

func main() {
	check()
	clearTerminal()
	getNames()
}

func getNames() {
	fmt.Printf("There are %v people in classrom!\n", len(os.Args)-1)
	fmt.Printf("Hello %v!\n", os.Args[1])
	if len(os.Args) >= 3 {
		var s, sep string
		for i := 2; i < len(os.Args); i++ {
			s += sep + os.Args[i] + " and"
			sep = " "
		}
		i := strings.LastIndex(s, " and")
		excludingLast := s[:i] + strings.Replace(s[i:], " and", "", 1)
		fmt.Println("Hi " + excludingLast + "!")
	}
}

func clearTerminal() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
