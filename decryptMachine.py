import json

# Define los elementos de la máquina
states = ["q0", "halt"]
input_alphabet = [chr(i) for i in range(ord("A"), ord("Z") + 1)] + [" "]
tape_alphabet = input_alphabet
initial_state = "q0"
accept_states = ["halt"]

# Generar todas las transiciones posibles para cada shift (1-25)
transitions = []

for shift in range(0, 26):  # Shift desde 1 hasta 25
    for letter in input_alphabet:
        if letter == " ":
            # Espacios no detienen la máquina; permanecen en q0
            transitions.append(
                {
                    "state": "q0",
                    "input": letter,
                    "next_state": "q0",  # Continuar en q0
                    "write": letter,  # Mantener el espacio
                    "move": "R",
                    "shift": shift,
                }
            )
        else:
            # Procesar letras del alfabeto
            current_index = input_alphabet.index(letter)
            shifted_index = (current_index - shift) % 26  # Solo 26 letras (A-Z)
            transitions.append(
                {
                    "state": "q0",
                    "input": letter,
                    "next_state": "q0",  # Continuar procesando en q0
                    "write": input_alphabet[shifted_index],  # Aplicar descifrado
                    "move": "R",
                    "shift": shift,
                }
            )

# Crear la estructura JSON
machine = {
    "states": states,
    "input_alphabet": input_alphabet,
    "tape_alphabet": tape_alphabet,
    "initial_state": initial_state,
    "accept_states": accept_states,
    "transitions": transitions,
}

# Guardar como archivo JSON
output_file = "machines/decryptMachine.json"
with open(output_file, "w") as json_file:
    json.dump(machine, json_file, indent=4)

print(f"Archivo {output_file} generado correctamente.")
