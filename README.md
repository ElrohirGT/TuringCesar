# Proyecto de Máquina de Turing para Encriptación y Desencriptación

Este proyecto implementa una Máquina de Turing que permite realizar operaciones de encriptación y desencriptación utilizando el cifrado César. El programa es interactivo y funciona mediante un menú que guía al usuario en cada paso del proceso.

## Requisitos

- **Go**: Necesitas tener instalado el lenguaje Go en tu máquina para ejecutar el proyecto.

## Cómo ejecutar el proyecto

1. Clona o descarga este repositorio.

2. Navega a la carpeta del proyecto en tu terminal.

3. Ejecuta el comando:

bash

Copy code

`go run .`

4. Aparecerá un menú interactivo con las siguientes opciones:

- **1. Encriptar**
- **2. Descifrar**
- **3. Salir**

## Opciones del Menú

### **1. Encriptar**

Para encriptar un mensaje, debes proporcionar:

- Una clave (2 dígitos) seguida directamente del mensaje.
- El mensaje puede incluir espacios.

#### Ejemplo de entrada:

```bash
04HOLA MUNDO
```

**Clave:** `04`

**Cadena:** `HOLA MUNDO`

La salida será el mensaje encriptado.

---

### **2. Descifrar**

Para descifrar un mensaje, debes proporcionar:

- Una clave (2 dígitos) seguida de un espacio, y luego el mensaje encriptado.
- El mensaje puede incluir espacios.

#### Ejemplo de entrada:

```bash
04 HOLA MUNDO
```

- **Clave:** `04`
- **Mensaje:** `HOLA MUNDO`

La salida será el mensaje descifrado.

---

### **3. Salir**

Selecciona esta opción para salir del programa.

## Reglas del Programa

1. **Formato de entrada:**

- **Para encriptar:** La clave y el mensaje deben ir juntos sin espacios entre ellos.
- **Para desencriptar:** Debe haber un espacio entre la clave y el mensaje.

2. **Espacios:**

- Los espacios en el mensaje son preservados durante los procesos de encriptación y desencriptación.

3. **Punto (`.`):**

- El programa agrega automáticamente un punto (.) al final del mensaje, indicando el final de la cinta en la Máquina de Turing. Este punto no aparece en el resultado final.

4. **Clave:**

- La clave debe ser un número entre `00` y `25`. Cualquier número fuera de este rango generará un error.

5. **Estado de aceptación:**

- La Máquina de Turing procesará cada carácter hasta llegar al final del mensaje (indicado por el punto agregado). El estado de aceptación asegura que el mensaje fue procesado completamente.

6. **Errores:**

- Si el programa no encuentra una transición válida en la Máquina de Turing, mostrará un error detallado sobre el estado, símbolo y posición en la cinta donde ocurrió el problema.

## Archivos Importantes

- **`machines/decryptMachine.json`:** Contiene la definición de la Máquina de Turing utilizada para descifrar mensajes.
- **`machines/generatedEncrypt.txt`:** Puede contener datos para la operación de encriptación.

## Ejemplo de Uso

### Encriptar

```bash
--- MENÚ ---
1. Encriptar
2. Descifrar
3. Salir
Selecciona una opción: 1
Ingresa el mensaje para encriptar: 04HOLA MUNDO
El mensaje cifrado es: lspe qyrhs
```

### Descifrar

```bash
--- MENÚ ---
1. Encriptar
2. Descifrar
3. Salir
Selecciona una opción: 2
Ingresa la cadena para descifrar: 04 lspe qyrhs
Usando la clave: 4
Mensaje descifrado: HOLA MUNDO

```
