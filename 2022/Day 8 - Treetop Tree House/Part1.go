package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxValue(arr []int) int {
	temp := -1
	for _, e := range arr {
		if e > temp {
			temp = e
		}
	}
	return temp
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
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			left := maxValue(grid[i][:j])
			right := maxValue(grid[i][j+1:])
			top := maxValue(rotGrid[j][:i])
			bottom := maxValue(rotGrid[j][i+1:])
			//fmt.Println(grid[i][j] > left, grid[i][j] > right, grid[i][j] > top, grid[i][j] > bottom)
			//fmt.Println(left, right, top, bottom)
			//fmt.Println("===")
			if grid[i][j] > left || grid[i][j] > right || grid[i][j] > top || grid[i][j] > bottom {
				ans++
			}
		}
	}

	ans += (len(grid)*2 + len(grid[0])*2) - 4

	fmt.Println(ans)
}
