package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseEncodedMessage(input string) (int, string) {
	if len(input) < 3 {
		panic("La entrada debe tener al menos 3 caracteres: una llave (2 dígitos) y un mensaje.")
	}
	key, err := strconv.Atoi(input[:2])
	if err != nil {
		panic("La llave debe ser un número válido entre 0 y 25.")
	}
	message := input[3:]
	return key, message
}

func DecryptWithMachine(message string, machine TuringMachineJson, key int) string {
	// Agregar el punto al final de la cinta
	message += "."

	tape := strings.Split(message, "") // Convertimos el mensaje en una cinta (array de caracteres)
	head := 0                          // Posición inicial del cabezal
	state := machine.InitialState      // Estado inicial

	// Filtrar transiciones para la clave actual (shift/key)
	filteredTransitions := filterTransitionsByShift(machine.Transitions, key)

	// Iterar sobre la cinta hasta que el estado sea de aceptación o se produzca un error
	for head < len(tape) && head >= 0 {
		currentSymbol := tape[head]
		transitionFound := false

		// Buscar una transición válida para el estado y símbolo actuales
		for _, transition := range filteredTransitions {
			// Verificar si la transición coincide con el estado y símbolo actuales
			if transition["state"] == state && transition["input"] == currentSymbol {
				fmt.Printf("Transición encontrada: state=%s, input=%s, write=%s, next_state=%s, move=%s\n",
					state, currentSymbol, transition["write"], transition["next_state"], transition["move"])

				// Aplicar transición
				tape[head] = transition["write"].(string)
				state = transition["next_state"].(string)

				// Mover el cabezal
				if transition["move"] == "R" {
					head++
				} else if transition["move"] == "L" {
					head--
				}

				transitionFound = true
				break
			}
		}

		// Si no se encontró una transición, mostrar error detallado
		if !transitionFound {
			panic(fmt.Sprintf("Transición no encontrada para state=%s, input=%s, head=%d",
				state, currentSymbol, head))
		}
	}

	// Verificar si el estado final es un estado de aceptación
	for _, acceptState := range machine.AcceptStates {
		if state == acceptState {
			// Retornar la cinta descifrada sin el punto final
			return strings.TrimSuffix(strings.Join(tape, ""), ".")
		}
	}

	panic("La máquina terminó en un estado no aceptado.")
}

// Filtrar transiciones basadas en el shift (clave)
func filterTransitionsByShift(transitions []map[string]interface{}, shift int) []map[string]interface{} {
	var filtered []map[string]interface{}
	for _, transition := range transitions {
		if transition["shift"] == float64(shift) { // JSON decodifica números como float64
			filtered = append(filtered, transition)
		}
	}
	return filtered
}
