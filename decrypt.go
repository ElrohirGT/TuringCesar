package main

import (
	"strconv"
	"strings"
)

func ParseEncodedMessage(input string) (int, string) {
	if len(input) < 3 {
		panic("La entrada debe tener al menos 3 caracteres: una llave (2 dígitos) y un mensaje.")
	}
	key, err := strconv.Atoi(input[:2]) // Los primeros dos dígitos son la llave.
	if err != nil {
		panic("La llave debe ser un número válido entre 0 y 26.")
	}
	message := input[3:] // El resto es el mensaje cifrado.
	return key, message
}

func DecryptCesar(message string, key int, alphabet []string) string {
	n := len(alphabet)
	alphabetMap := make(map[string]int)
	reverseMap := make(map[int]string)

	for i, char := range alphabet {
		alphabetMap[char] = i
		reverseMap[i] = char
	}

	var decoded strings.Builder
	for _, char := range message {
		if char == ' ' {
			decoded.WriteString(" ")
			continue
		}
		pos := (alphabetMap[string(char)] - key + n) % n
		decoded.WriteString(reverseMap[pos])
	}
	return decoded.String()
}
