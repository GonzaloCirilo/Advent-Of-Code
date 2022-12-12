package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func maxValue(arr []int, c chan int) {
	temp := -1
	for _, e := range arr {
		if e > temp {
			temp = e
		}
	}
	c <- temp
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
		//fmt.Println(rotGrid[i])
	}
	ans := 0
	var wg, ansWg sync.WaitGroup
	wg.Add((len(grid) - 2) * (len(grid[0]) - 2))
	ansch := make(chan int)
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			go func(i, j int) {
				defer wg.Done()
				c := make(chan int, 4)

				go maxValue(grid[i][:j], c)      //left
				go maxValue(grid[i][j+1:], c)    // right
				go maxValue(rotGrid[j][:i], c)   // top
				go maxValue(rotGrid[j][i+1:], c) // bottom
				a1, a2, a3, a4 := <-c, <-c, <-c, <-c
				if grid[i][j] > a1 || grid[i][j] > a2 || grid[i][j] > a3 || grid[i][j] > a4 {
					ansch <- 1
				}
			}(i, j)
		}
	}
	go func() {
		ansWg.Add(1)
		defer ansWg.Done()
		for x := range ansch {
			ans += x
		}
	}()

	wg.Wait()
	close(ansch)
	ans += (len(grid)*2 + len(grid[0])*2) - 4
	ansWg.Wait()
	fmt.Println(ans)
}
