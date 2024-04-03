package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Simular la entrada del usuario
	input := "test_user\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))

	// Captura la salida estándar
	var output bytes.Buffer
	stdout := &output

	// Ejecutar la función con el nuevo lector
	// os.Stdin = reader

	// Ejecutar la función
	fmt.Fprint(stdout, "Ingrese su nombre de usuario: ")
	username, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		t.Errorf("Error al leer el nombre de usuario: %v", err)
	}
	username = strings.TrimSpace(username)

	// Comprobar si la salida coincide con lo esperado
	expectedOutput := "Ingrese su nombre de usuario: "
	if output.String() != expectedOutput {
		t.Errorf("La salida esperada es %q pero se obtuvo %q", expectedOutput, output.String())
	}

	// Comprobar si el nombre de usuario coincide con lo esperado
	expectedUsername := "test_user"
	if username != expectedUsername {
		t.Errorf("El nombre de usuario esperado es %q pero se obtuvo %q", expectedUsername, username)
	}
}
func TestGetPassowrd(t *testing.T) {
	// Simular la entrada del usuario
	input := "test_password\n"
	reader := bufio.NewReader(bytes.NewBufferString(input))

	// Captura la salida estándar
	var output bytes.Buffer
	stdout := &output

	// Ejecutar la función con el nuevo lector
	// os.Stdin = reader

	// Ejecutar la función
	fmt.Fprint(stdout, "Ingrese su password de usuario: ")
	passord, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		t.Errorf("Error al leer el passord de usuario: %v", err)
	}
	passord = strings.TrimSpace(passord)

	// Comprobar si la salida coincide con lo esperado
	expectedOutput := "Ingrese su password de usuario: "
	if output.String() != expectedOutput {
		t.Errorf("La salida esperada es %q pero se obtuvo %q", expectedOutput, output.String())
	}

	// Comprobar si el nombre de usuario coincide con lo esperado
	expectedUsername := "test_password"
	if passord != expectedUsername {
		t.Errorf("El nombre de usuario esperado es %q pero se obtuvo %q", expectedUsername, passord)
	}
}
