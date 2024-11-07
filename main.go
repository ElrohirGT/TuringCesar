package main

import "fmt"

func main() {
	fmt.Println("Please write your turing machine...")
	machine, input := ParseMachine()

	fmt.Printf("INPUT: %v\n", input)

	fmt.Printf("MACHINE: %v\n", machine)

	fmt.Println("SIMULATING MACHINE...")
	result := machine.Run(input)

	if result.Accepted {
		fmt.Println("The input is ACCEPTED!")
	} else {
		fmt.Println("The input is REJECTED!")
	}

	fmt.Println("It took", result.Iterations, "iterations to finish...")

}
