/*
Implementa un árbol de segmentos en Go que admita dos operaciones: actualizar un valor en un índice específico del arreglo subyacente y calcular la suma de un rango de valores en el arreglo. Proporciona ejemplos de uso para demostrar la eficiencia de tu implementación.
*/

package main

import "fmt"

type SegmentTree struct {
	data     []int
	tree     []int
	nodeSize int
}

func NewSegmentTree(data []int) *SegmentTree {
	nodeSize := len(data)
	tree := make([]int, nodeSize*2)
	st := &SegmentTree{data, tree, nodeSize}
	st.build()
	return st
}

func (st *SegmentTree) build() {
	// Inicializamos los nodos
	for i := st.nodeSize; i < st.nodeSize*2; i++ {
		st.tree[i] = st.data[i-st.nodeSize]
	}
	// Construimos el árbol calculando los padres
	for i := st.nodeSize - 1; i > 0; i-- {
		st.tree[i] = st.tree[i*2] + st.tree[i*2+1]
	}
}

func (st *SegmentTree) Query(left, right int) int {
	res := 0
	left += st.nodeSize
	right += st.nodeSize
	for left < right {
		if left%2 == 1 {
			res += st.tree[left]
			left++
		}
		if right%2 == 1 {
			right--
			res += st.tree[right]
		}
		left /= 2
		right /= 2
	}
	return res
}

func (st *SegmentTree) update(index, value int) {
	index += st.nodeSize
	st.tree[index] = value
	for index > 1 {
		index /= 2
		st.tree[index] = st.tree[index*2] + st.tree[index*2+1]
	}
}

func main() {
	date := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	st := NewSegmentTree(date)

	fmt.Println("Suma de rango (1,5):", st.Query(1, 5))
}
