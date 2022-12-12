package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Node struct {
	X int
	Y int
	V int
}

type IntHeap []Node

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i].V < h[j].V }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
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
	grid := make([][]int, 0)
	ans := make([][]int, 0)
	var sourceX, sourceY, targetX, targetY, cont int
	for scanner.Scan() {
		line := scanner.Text()
		aux, auxAns := make([]int, 0), make([]int, 0)
		for i, t := range line {
			if t == 'S' {
				aux = append(aux, 0)
				sourceX = i
				sourceY = cont
			} else if t == 'E' {
				aux = append(aux, int('z')-'a')
				targetY = cont
				targetX = i
			} else {
				aux = append(aux, int(t)-'a')
			}
			auxAns = append(auxAns, 100000000)
		}
		grid = append(grid, aux)
		ans = append(ans, auxAns)
		cont++
	}
	h := &IntHeap{}
	heap.Init(h)
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	heap.Push(h, Node{X: sourceX, Y: sourceY, V: 0})
	ans[sourceY][sourceX] = 0
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		for i := range dx {
			newX, newY := n.X+dx[i], n.Y+dy[i]
			if newX < 0 || newY < 0 || newY >= len(grid) || newX >= len(grid[0]) {
				continue
			}

			if grid[newY][newX]-grid[n.Y][n.X] <= 1 {
				if n.V+1 < ans[newY][newX] {
					ans[newY][newX] = n.V + 1
					newNode := Node{X: newX, Y: newY, V: n.V + 1}
					heap.Push(h, newNode)
				}
			}

		}
	}

	fmt.Println(ans[targetY][targetX])
}
