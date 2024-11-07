package main

import (
	"fmt"
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
	fmt.Printf("INITIAL TAPE: %v", machine.tape)

	maxIterations := 10000
	i := 0
machineLoop:
	for i = 0; i < maxIterations; i++ {
		machine.PrintState()

		tranInput := TransitionInput{
			state:     machine.head.state,
			character: machine.tape[machine.head.position],
		}

		fmt.Printf("Trying to find transition with input: %v", tranInput)
		output, found := machine.transitions[tranInput]

		if !found {
			panic("No transition found for this input!")
		}

		fmt.Printf("OUTPUT: %v", output)

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
	}

	return SimulationResult{
		Accepted:   isAccepted,
		Iterations: i,
	}
}

func (machine TuringMachine) PrintState() {
	fmt.Print("TAPE:")
	for idx, char := range machine.tape {
		if machine.head.position == idx {
			fmt.Printf(" [%v: %v]", char, machine.head.state)
		} else {
			fmt.Printf(" %v", char)
		}
	}
}
