package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type void struct{}

var member void

func updatePosition(tx, ty, hx, hy int) (int, int) {
	//fmt.Println("prev Up", tx, ty, hx, hy)
	if tx == hx || ty == hy {
		dx := []int{1, 0, 0, -1}
		dy := []int{0, -1, 1, 0}
		if math.Abs(float64(hx)-float64(tx)) >= 2 || math.Abs(float64(hy)-float64(ty)) >= 2 {
			for i := range dx {
				auxX, auxY := tx+dx[i], ty+dy[i]
				if math.Abs(float64(hx)-float64(auxX)) < 2 && math.Abs(float64(hy)-float64(auxY)) < 2 {
					return auxX, auxY
				}
			}
		}
	} else {
		dx := []int{1, 1, -1, -1}
		dy := []int{1, -1, -1, 1}
		if math.Abs(float64(hx)-float64(tx)) >= 2 || math.Abs(float64(hy)-float64(ty)) >= 2 {
			for i := range dx {
				auxX, auxY := tx+dx[i], ty+dy[i]
				if math.Abs(float64(hx)-float64(auxX)) < 2 && math.Abs(float64(hy)-float64(auxY)) < 2 {
					return auxX, auxY
				}
			}
		}
	}
	return tx, ty
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var hx, hy, tx, ty int
	bx := make([]int, 8)
	by := make([]int, 8)

	visited := make(map[string]interface{})
	visited["0,0"] = member
	for scanner.Scan() {
		line := scanner.Text()
		var dir string
		var amount int
		fmt.Sscanf(line, "%s %d", &dir, &amount)
		for i := 1; i <= amount; i++ {
			if dir == "R" {
				hx++
			} else if dir == "L" {
				hx--
			} else if dir == "U" {
				hy--
			} else if dir == "D" {
				hy++
			}
			bx[0], by[0] = updatePosition(bx[0], by[0], hx, hy)
			for j := 1; j <= 7; j++ {
				bx[j], by[j] = updatePosition(bx[j], by[j], bx[j-1], by[j-1])
			}
			tx, ty = updatePosition(tx, ty, bx[len(bx)-1], by[len(by)-1])
			visited[strconv.Itoa(tx)+","+strconv.Itoa(ty)] = member
		}

	}

	fmt.Println(len(visited))
}
