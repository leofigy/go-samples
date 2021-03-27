package main

import (
	"container/heap"
	"fmt"
	"log"
)

func main() {
	data := []int{2, 4545, 93, 84, -12, 4, -10}
	log.Println(data)
	min, err := getKGreatest(data, 2)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("minimum ", min)
}

func getKGreatest(array []int, k int) (int, error) {
	if k <= 1 {
		return 0, fmt.Errorf("Invalid k arg, must be greater than 1")
	}

	if len(array) <= k {
		return 0, fmt.Errorf("The k must be smaller or equal to array size")
	}

	localHeap := make(IntHeap, 0, k)
	heap.Init(&localHeap)

	for i, item := range array {

		// initialization
		if i < k {
			heap.Push(&localHeap, item)
			log.Println(localHeap)
			continue
		}

		// item minimum item
		min := heap.Pop(&localHeap)
		log.Println()
		newMax := min
		if item >= min.(int) {
			newMax = item
		}
		heap.Push(&localHeap, newMax)

	}

	log.Println(localHeap)

	return (heap.Pop(&localHeap)).(int), nil

}

type IntHeap []int

// implement interface len,less, swap

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
