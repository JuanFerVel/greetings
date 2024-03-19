# Saludos en Go

Este páguqete proporciona una forma simple de obtener saludos personalizados

## Instalación
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/JuanFerVel/greetings
```

## Uso

Aqui tienes un ejemplo de como usarlo:
```go
package main

import (
	"fmt"
	"log"

	"github.com/juanfervel/greetings"
)

func main() {
	log.SetPrefix("greetings(): ")
	log.SetFlags(0)

	names := []string{
		"Juan",
		"Felipe",
		"Jose",
	}

	message, err := greetings.Hellos(names)

	// message, err := greetings.Hello("Luis")

	if err != nil {
		log.Fatal(err)
	}

	for index, value := range message {
		fmt.Printf("%v: ", index)
		fmt.Println(value)
	}
}
```
