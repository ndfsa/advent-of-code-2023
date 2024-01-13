package util

type StateHeap[T any] []State[T]

func (b *StateHeap[T]) Len() int {
	return len(*b)
}

func (b *StateHeap[T]) Less(i, j int) bool {
	return (*b)[i].Cost < (*b)[j].Cost
}

func (b *StateHeap[T]) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *StateHeap[T]) Push(element any) {
	*b = append(*b, element.(State[T]))
}

func (b *StateHeap[T]) Pop() any {
	length := len(*b)
	res := (*b)[length-1]
	*b = (*b)[:length-1]
	return res
}

type State[T any] struct {
	Vertex T
	Cost  int
}
