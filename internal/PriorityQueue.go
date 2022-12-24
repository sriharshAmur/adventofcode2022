package internal

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	Queue []T
}

func (pq *PriorityQueue[T]) Append(e T) {
	length := len(pq.Queue)
	index := 0
	for index < length {
		current := pq.Queue[index]
		if current >= e {
			break
		}
		index++
	}
	if index == length {
		pq.Queue = append(pq.Queue, e)
	} else {
		pq.Queue = append(pq.Queue[:index+1], pq.Queue[index:]...)
		pq.Queue[index] = e
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	e := pq.Queue[0]
	pq.Queue = pq.Queue[1:]
	return e
}

func (pq *PriorityQueue[T]) Print() {
	fmt.Println(pq.Queue)
}
