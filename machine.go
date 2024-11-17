package main

import (
	"fmt"
	"strconv"
)

type HeadMovement int

const (
	Left HeadMovement = iota
	Right
)

func (ss HeadMovement) String() string {
	switch ss {
	case Left:
		return "LEFT"
	case Right:
		return "RIGHT"
	default:
		return "INVALID STATE RECEIVED!"
	}
}

type TransitionInput struct {
	state     string
	character rune
}

type TransitionOutput struct {
	state        string
	character    rune
	headMovement HeadMovement
}

type TuringHead struct {
	state    string
	position int
}

type TuringMachine struct {
	tape []rune
	head TuringHead

	transitions map[TransitionInput]TransitionOutput
	accepted    []string
	rejected    []string
}

type SimulationResult struct {
	Accepted   bool
	Iterations int
}

func (machine TuringMachine) Run(input string) SimulationResult {
	machine.tape = []rune{}
	for _, char := range input {
		machine.tape = append(machine.tape, char)
	}

	isAccepted := true
	machine.head.position = 0

	maxIterations := 10000
	i := 0
machineLoop:
	for i = 0; i < maxIterations; i++ {
		machine.PrintState()

		tranInput := TransitionInput{
			state:     machine.head.state,
			character: machine.tape[machine.head.position],
		}

		var output TransitionOutput
		if tranInput.state == " " {
			output = TransitionOutput{
				state:        machine.head.state,
				character:    ' ',
				headMovement: Right,
			}
		} else {
			fmt.Println("Trying to find transition with state:", tranInput.state, "and char:", strconv.QuoteRune(tranInput.character))

			machineOutput, found := machine.transitions[tranInput]
			if !found {
				panic("No transition found for this input!")
			}

			output = machineOutput
		}

		fmt.Println("OUTPUT:", output.state, strconv.QuoteRune(output.character), output.headMovement.String())

		machine.head.state = output.state
		machine.tape[machine.head.position] = output.character

		if output.headMovement == Left {
			machine.head.position -= 1

			if machine.head.position == -1 {
				machine.tape = append([]rune{'_'}, machine.tape...)
				machine.head.position = 0
			}
		} else {
			machine.head.position += 1

			if machine.head.position == len(machine.tape) {
				machine.tape = append(machine.tape, '_')
			}
		}

		for _, state := range machine.accepted {
			if state == machine.head.state {
				isAccepted = true
				break machineLoop
			}
		}

		for _, state := range machine.rejected {
			if state == machine.head.state {
				isAccepted = false
				break machineLoop
			}
		}
		fmt.Println("---")
	}
	fmt.Println("===")

	fmt.Println("FINAL MACHINE STATE:")
	machine.PrintState()

	return SimulationResult{
		Accepted:   isAccepted,
		Iterations: i,
	}
}

func (machine TuringMachine) PrintState() {
	fmt.Print("TAPE:")
	for idx, char := range machine.tape {
		if machine.head.position == idx {
			fmt.Printf(" [%s: %s]", strconv.QuoteRune(char), machine.head.state)
		} else {
			fmt.Printf(" %s", strconv.QuoteRune(char))
		}
	}

	fmt.Println()
}
