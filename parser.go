package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseMachine() (TuringMachine, string) {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	acceptedStates := strings.Split(scanner.Text(), " ")

	scanner.Scan()
	rejectedStates := strings.Split(scanner.Text(), " ")

	scanner.Scan()
	initialState := scanner.Text()

	scanner.Scan()
	input := scanner.Text()

	transitions := map[TransitionInput]TransitionOutput{}
	for scanner.Scan() {
		transitionParts := strings.Split(scanner.Text(), "->")
		fmt.Printf("Transition: %v -> %v\n", transitionParts[0], transitionParts[1])

		transInputs := strings.Split(strings.TrimSpace(transitionParts[0]), " ")
		inputs := TransitionInput{
			state:     transInputs[0],
			character: []rune(transInputs[1])[0],
		}

		transOutputs := strings.Split(strings.TrimSpace(transitionParts[1]), " ")
		movement := Left
		if transOutputs[2] == "R" {
			movement = Right
		} else if transOutputs[2] == "L" {
			movement = Left
		} else {
			panic(fmt.Sprintf("Invalid head movement supplied! It must be `R` or `L` (%v)", transOutputs))
		}
		outputs := TransitionOutput{
			state:        transOutputs[0],
			character:    []rune(transOutputs[1])[0],
			headMovement: movement,
		}

		transitions[inputs] = outputs
	}

	return TuringMachine{
		accepted:    acceptedStates,
		rejected:    rejectedStates,
		transitions: transitions,
		tape:        []rune{},
		head:        TuringHead{state: initialState, position: 0},
	}, input
}
