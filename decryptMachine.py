import json

# Define los elementos de la máquina
letters = [chr(i) for i in range(ord("A"), ord("Z") + 1)]  # Alfabeto A-Z
states = (
    ["q0"] + [f"q{letter}" for letter in letters] + ["accept"]
)  # Estado inicial, estados para letras, y nuevo estado de aceptación
input_alphabet = letters + [" ", "."]  # Alfabeto de entrada incluye espacio y punto
tape_alphabet = input_alphabet  # Mismo alfabeto para la cinta
initial_state = "q0"
accept_states = ["accept"]  # Nuevo estado de aceptación

# Generar todas las transiciones posibles para cada shift (0-25)
transitions = []

for shift in range(0, 26):  # Shift desde 0 hasta 25
    for letter in input_alphabet:
        if letter == " ":
            # Los espacios mantienen el estado actual
            transitions.append(
                {
                    "state": "q0",
                    "input": letter,
                    "next_state": "q0",  # Mantenerse en el estado actual
                    "write": letter,  # Mantener el espacio
                    "move": "R",
                    "shift": shift,
                }
            )
        elif letter == ".":
            # El punto lleva directamente al estado de aceptación
            transitions.append(
                {
                    "state": "q0",
                    "input": ".",
                    "next_state": "accept",
                    "write": ".",
                    "move": "R",
                    "shift": shift,
                }
            )
        else:
            # Procesar letras del alfabeto, cada letra tiene su propio estado
            current_index = letters.index(letter)
            shifted_index = (current_index - shift) % 26  # Solo 26 letras (A-Z)
            transitions.append(
                {
                    "state": "q0",
                    "input": letter,
                    "next_state": f"q{letter}",  # Ir al estado específico de la letra
                    "write": letters[shifted_index],  # Aplicar descifrado
                    "move": "R",
                    "shift": shift,
                }
            )

# Generar transiciones para todos los estados `q{letra}`
for shift in range(0, 26):
    for current_letter in letters:
        for next_letter in input_alphabet:
            if next_letter == " ":
                # Si se encuentra un espacio, permanecer en el estado actual
                transitions.append(
                    {
                        "state": f"q{current_letter}",
                        "input": next_letter,
                        "next_state": f"q{current_letter}",  # Mantenerse en el estado actual
                        "write": next_letter,
                        "move": "R",
                        "shift": shift,
                    }
                )
            elif next_letter == ".":
                # Si se encuentra un punto, ir directamente a aceptación
                transitions.append(
                    {
                        "state": f"q{current_letter}",
                        "input": next_letter,
                        "next_state": "accept",
                        "write": next_letter,
                        "move": "R",
                        "shift": shift,
                    }
                )
            else:
                # Continuar procesando desde el estado actual
                current_index = letters.index(next_letter)
                shifted_index = (current_index - shift) % 26
                transitions.append(
                    {
                        "state": f"q{current_letter}",
                        "input": next_letter,
                        "next_state": f"q{next_letter}",
                        "write": letters[shifted_index],
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
