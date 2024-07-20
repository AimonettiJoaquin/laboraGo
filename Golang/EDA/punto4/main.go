/*
Ejercicio 4: Estructura de datos para el juego del ahorcado

Diseña una estructura de datos en Go para representar el estado de un juego del ahorcado. La estructura debe permitir almacenar la palabra a adivinar, el progreso de las letras adivinadas y el número de intentos restantes. Proporciona funciones para actualizar el progreso del juego y determinar si el jugador ha ganado o perdido.

*/

package main

import (
	"fmt"
	"strings"
)

type Ahorcado struct {
	palabra    string
	letrasProg []string
	nroIntento int
}

func NewAhorcado(palabra string, nroIntento int) *Ahorcado {
	intento := make([]string, len(palabra))
	for i := range palabra {
		intento[i] = "_"
	}

	return &Ahorcado{palabra: palabra, letrasProg: intento, nroIntento: nroIntento}

}

func (a *Ahorcado) Intento(letra string) {
	//existe := false

	if strings.Contains(a.palabra, letra) {
		for i, char := range a.palabra {
			if string(char) == letra {
				a.letrasProg[i] = string(char)
			}
		}
		fmt.Println("Correcto")
		fmt.Println(a.letrasProg)
	} else {
		a.nroIntento--
		fmt.Println("Lo siento,", letra, "no se encuentra.")
		fmt.Println("te quedan", a.nroIntento, "intentos")
	}
}

func (a *Ahorcado) IntentoPalabra(palabra string) bool {
	return palabra == a.palabra
}

func (a *Ahorcado) ganar() bool {
	return a.palabra == strings.Join(a.letrasProg, "")
}

func main() {

	a := NewAhorcado("playa", 6)
	fmt.Println(a.letrasProg)
	for {

		var intento string
		fmt.Scanln(&intento)

		if len(intento) > 1 && a.IntentoPalabra(intento) {
			fmt.Println("Ganaste, bien ahi wacho")
			break
		} else if len(intento) != 1 {
			fmt.Println("Lo siento,", intento, "no es la palabra.")
			a.nroIntento--
			fmt.Println("te quedan:", a.nroIntento)
		}

		if len(intento) == 1 {
			a.Intento(intento)
		}
		if a.nroIntento <= 0 {
			fmt.Println("Perdiste, sorry bro :c")
			break
		} else if a.ganar() {
			fmt.Println("Ganaste, bien ahi wacho")
			break
		}
	}
}
