package main

import (
	"fmt"
)

type Cola struct {
	items []interface{}
}

func (q *Cola) Enqueue(elemento interface{}) {
	q.items = append(q.items, elemento)
}

func (q *Cola) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	elemento := q.items[0]
	q.items = q.items[1:]
	return elemento
}

func (q *Cola) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Cola) Size() int {
	return len(q.items)
}

func (q *Cola) Empty() bool {
	return len(q.items) == 0
}

func main() {
	cola := Cola{}

	cola.Enqueue(1)
	cola.Enqueue(2)
	cola.Enqueue(3)

	fmt.Println("Tamaño de la cola:", cola.Size())
	fmt.Println("Frente de la cola:", cola.Front())

	fmt.Println("Desencolando elementos:")
	for !cola.Empty() {
		fmt.Println(cola.Dequeue())
	}
}
