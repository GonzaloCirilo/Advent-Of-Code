package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 200)

	for i := range grid {
		grid[i] = make([]int, 600)
	}
	for scanner.Scan() {
		line := scanner.Text()
		coord := strings.Split(line, " -> ")
		points := make([]Point, 0)
		var x, y int
		for _, v := range coord {
			fmt.Sscanf(v, "%d,%d", &x, &y)
			points = append(points, Point{X: x, Y: y})
		}

		for i := 1; i < len(points); i++ {
			if points[i].X == points[i-1].X {
				lb, ub := Min(points[i-1].Y, points[i].Y), Max(points[i-1].Y, points[i].Y)
				for j := lb; j <= ub; j++ {
					grid[j][points[i].X] = 1
				}
			} else if points[i].Y == points[i-1].Y {
				lb, ub := Min(points[i-1].X, points[i].X), Max(points[i-1].X, points[i].X)
				for j := lb; j <= ub; j++ {
					grid[points[i].Y][j] = 1
				}
			}

		}
	}
	ans := 0
	dx := []int{0, -1, 1}
	dy := []int{1, 1, 1}
	done := false
	for {
		sPoint := Point{X: 500, Y: 0}
		if done {
			break
		}
		for {
			if sPoint.Y+1 >= len(grid) { // fall to infinity
				done = true
				break
			}
			moved := false
			for i := range dx {
				newPoint := Point{X: sPoint.X + dx[i], Y: sPoint.Y + dy[i]}
				if grid[newPoint.Y][newPoint.X] != 0 {
					continue
				} else {
					sPoint = newPoint
					moved = true
					break
				}
			}
			if moved == false {
				grid[sPoint.Y][sPoint.X] = 2
				ans++
				break
			}
		}
	}

	fmt.Println(ans)
}
