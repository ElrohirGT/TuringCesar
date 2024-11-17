import json

# Define los elementos de la máquina
states = ["q0", "q1", "q2", "q3"]
input_alphabet = [chr(i) for i in range(ord("A"), ord("Z") + 1)] + [" "]
tape_alphabet = input_alphabet
initial_state = "q0"
accept_states = ["q2"]

# Genera todas las transiciones posibles
transitions = []

for state in states:
    for symbol in input_alphabet:
        if state == "q0":
            next_state = "q1"  # Estado de transición
        elif state == "q1" and symbol == " ":
            next_state = "q2"  # Estado de aceptación
        else:
            next_state = "q3"  # Estado genérico

        # Movimiento genérico
        move = "R" if state != "q3" else "L"

        transitions.append(
            {
                "state": state,
                "input": symbol,
                "next_state": next_state,
                "write": symbol,
                "move": move,
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
with open("machines/decryptMachine.json", "w") as json_file:
    json.dump(machine, json_file, indent=4)

print("Archivo decryptMachine.json generado correctamente.")
