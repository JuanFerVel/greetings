package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo  para la persona especificada.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	//Devuelve el saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola %v! ¡Bienvenido!",
		"¡Que hay de nuevo %v!",
		"Que bueno verte de nuevo %v!",
		"Buen día %v, encantado de conocerte.",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}
