package main

import "fmt"

func main() {
	fmt.Println("Please write your turing machine...")
	machine, input := ParseMachine()

	fmt.Printf("INPUT: %v\n", input)

	fmt.Printf("MACHINE: %v\n", machine)
}
