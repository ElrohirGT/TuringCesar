package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type TuringMachineJson struct {
	States        []string                 `json:"states"`
	InputAlphabet []string                 `json:"input_alphabet"`
	TapeAlphabet  []string                 `json:"tape_alphabet"`
	InitialState  string                   `json:"initial_state"`
	AcceptStates  []string                 `json:"accept_states"`
	Transitions   []map[string]interface{} `json:"transitions"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- MENÚ ---")
		fmt.Println("1. Encriptar")
		fmt.Println("2. Descifrar")
		fmt.Println("3. Salir")
		fmt.Print("Selecciona una opción: ")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, selecciona una opción válida.")
			continue
		}

		switch option {
		case 1:
			fmt.Print("Ingresa el mensaje para encriptar: ")
			message, _ := reader.ReadString('\n')
			message = strings.ToLower(strings.TrimSpace(message))

			filePath := "machines/generatedEncrypt.txt"
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error al leer el archivo en la ruta %s: %v", filePath, err)
				continue
			}
			machine := ParseMachine(file)
			machine.Run(message)

		case 2:
			jsonPath := strings.TrimSpace("machines/decryptMachine.json")

			fmt.Print("Ingresa la cadena para descifrar: ")
			encodedMessage, _ := reader.ReadString('\n')
			encodedMessage = strings.ToUpper(strings.TrimSpace(encodedMessage))

			machine := parseJSONMachine(jsonPath)
			key, message := ParseEncodedMessage(encodedMessage)
			decoded := DecryptCesar(message, key, machine.InputAlphabet)

			fmt.Printf("Mensaje descifrado: %s\n", decoded)
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func parseJSONMachine(filePath string) TuringMachineJson {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error al leer el archivo JSON en la ruta '%s': %v\n", filePath, err)
		os.Exit(1) // Finaliza el programa si no puede abrir el archivo.
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var machine TuringMachineJson
	if err := json.Unmarshal(byteValue, &machine); err != nil {
		panic(fmt.Sprintf("Error al parsear el archivo JSON: %v", err))
	}
	return machine
}
