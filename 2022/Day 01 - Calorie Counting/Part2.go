package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// logic inverted to have a max heap
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	current := 0
	h := &IntHeap{}
	heap.Init(h)
	for scanner.Scan() {
		text := scanner.Text()
		e, _ := strconv.Atoi(text)
		if len(text) <= 0 {
			heap.Push(h, current)
			current = 0
		} else {
			current += e
		}
	}
	ans := 0
	for i := 0; i < 3; i++ {
		ans += heap.Pop(h).(int)
	}

	fmt.Println(ans)
}
