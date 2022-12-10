package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(arr []int) []int {
	aux := make([]int, len(arr))
	copy(aux, arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		aux[i], aux[j] = aux[j], aux[i]
	}
	return aux
}

func countTrees(arr []int, h int) int {
	//fmt.Println(arr, h)
	count := 0
	for _, e := range arr {
		count++
		if e >= h {
			return count
		}
	}
	return count
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
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			left := countTrees(reverse(grid[i][:j]), grid[i][j])
			right := countTrees(grid[i][j+1:], grid[i][j])
			top := countTrees(reverse(rotGrid[j][:i]), grid[i][j])
			bottom := countTrees(rotGrid[j][i+1:], grid[i][j])
			//fmt.Println(grid[i][j] > left, grid[i][j] > right, grid[i][j] > top, grid[i][j] > bottom)
			//fmt.Println(left, right, top, bottom)
			//fmt.Println("===")
			score := left * right * top * bottom
			if score > ans {
				ans = score
			}
		}
	}

	fmt.Println(ans)
}
