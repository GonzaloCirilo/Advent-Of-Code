package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func reverse(arr []int) []int {
	aux := make([]int, len(arr))
	copy(aux, arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		aux[i], aux[j] = aux[j], aux[i]
	}
	return aux
}

func countTrees(arr []int, h int, c chan int) {
	//fmt.Println(arr, h)
	count := 0
	for _, e := range arr {
		count++
		if e >= h {
			c <- count
			return
		}
	}
	c <- count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		aux := make([]int, 0)
		for _, t := range line {
			aux = append(aux, int(t)-'0')
		}
		grid = append(grid, aux)
	}

	rotGrid := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		aux := make([]int, 0)
		for j := 0; j < len(grid[i]); j++ {
			aux = append(aux, grid[j][i])
		}
		rotGrid = append(rotGrid, aux)
	}
	ans := 0

	var wg, ansWg sync.WaitGroup
	wg.Add((len(grid) - 2) * (len(grid[0]) - 2))
	ansCh := make(chan int)
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			go func(i, j int) {
				defer wg.Done()
				c := make(chan int, 4)

				go countTrees(reverse(grid[i][:j]), grid[i][j], c)    // left
				go countTrees(grid[i][j+1:], grid[i][j], c)           // right
				go countTrees(reverse(rotGrid[j][:i]), grid[i][j], c) // top
				go countTrees(rotGrid[j][i+1:], grid[i][j], c)        // bottom
				a1, a2, a3, a4 := <-c, <-c, <-c, <-c
				score := a1 * a2 * a3 * a4
				ansCh <- score
			}(i, j)
		}
	}
	go func() {
		ansWg.Add(1)
		defer ansWg.Done()
		for x := range ansCh {
			if x > ans {
				ans = x
			}
		}
	}()
	wg.Wait()
	close(ansCh)
	ansWg.Wait()
	fmt.Println(ans)
}
