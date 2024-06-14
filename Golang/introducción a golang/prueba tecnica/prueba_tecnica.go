package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name       string
	Identifier int
}

func main() {
	//lista de jugadores actuales
	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}
	//registro de nuevos jugadores
	newPlayer := []string{"new player 1", "new player 2", "new player 3"}
	//eliminar los 3 primeros jugadores
	players = players[3:]
	//agregar
	for _, name := range newPlayer {
		players = append(players, name)
	}

	//mostrar la lista

	fmt.Println("Lista completa: ")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
	}

}
func sanitizeName(name string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1
	}, name)
}
