package main

import (
	"fmt"

	"github.com/bujosa/phpass"
)

func main() {
	password := "anything"

	// Crear un nuevo hasher de phpass
	hasher := phpass.New(phpass.NewConfig())

	// Generar el hash de la contraseña
	hashedPassword, err := hasher.Hash([]byte(password))
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Hashed password:", string(hashedPassword))

	// Example of check password
	// Check if the password matches the hash
	match := hasher.Check([]byte(password), hashedPassword)
	fmt.Print("Password match: ", match)
}
