package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Shape struct {
	Coord  Point
	Points []Point
}

func (s *Shape) moveHorizontally(amount int, grid [][]int) {
	aux := Shape{Coord: s.Coord, Points: s.Points}
	minX, maxX := 9, 0
	aux.Coord.X += amount
	for i := range aux.Points {
		x, y := aux.Points[i].X+aux.Coord.X, aux.Points[i].Y+aux.Coord.Y
		if minX > x {
			minX = x
		}
		if maxX < x {
			maxX = x
		}
		if x < 0 || x > 6 {
			return
		}
		if grid[y][x] > 0 {
			return
		}
	}
	if minX >= 0 && maxX <= 6 {
		*s = aux
	}
}

func (s *Shape) moveDown(grid [][]int) bool {
	aux := Shape{Coord: s.Coord, Points: s.Points}
	aux.Coord.Y -= 1
	min := 99999999
	for i := range aux.Points {
		x, y := aux.Coord.X+aux.Points[i].X, aux.Coord.Y+aux.Points[i].Y
		if aux.Coord.Y+aux.Points[i].Y < min {
			min = aux.Coord.Y + aux.Points[i].Y
		}
		if y < 0 {
			return false
		}
		if grid[y][x] > 0 {
			return false
		}
	}
	*s = aux
	return true

}

func (s *Shape) initCoord(maxSize int) {
	minX, minY := 9, 9
	for _, p := range s.Points {
		if minX > p.X {
			minX = p.X
		}
		if minY > p.Y {
			minY = p.Y
		}
	}
	minY *= -1
	s.Coord = Point{X: minX + 2, Y: minY + maxSize + 3}
}

func (s Shape) getHighestPeak() {

}

func Print(grid [][]int) {
	for _, v := range grid {
		fmt.Println(v)
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 3500)
	for i, _ := range grid {
		grid[i] = make([]int, 7)
	}
	var cmdLine string
	for scanner.Scan() {
		cmdLine = scanner.Text()
	}
	rocks := []Shape{
		Shape{
			Points: []Point{
				Point{X: 0, Y: 0},
				Point{X: 1, Y: 0},
				Point{X: 2, Y: 0},
				Point{X: 3, Y: 0},
			},
		},
		Shape{
			Points: []Point{
				Point{X: 0, Y: 0},
				Point{X: 1, Y: 0},
				Point{X: 2, Y: 0},
				Point{X: 1, Y: 1},
				Point{X: 1, Y: -1},
			},
		},
		Shape{
			Points: []Point{
				Point{X: 0, Y: 0},
				Point{X: 1, Y: 0},
				Point{X: 2, Y: 0},
				Point{X: 2, Y: 1},
				Point{X: 2, Y: 2},
			},
		},
		Shape{
			Points: []Point{
				Point{X: 0, Y: 0},
				Point{X: 0, Y: 1},
				Point{X: 0, Y: 2},
				Point{X: 0, Y: 3},
			},
		},
		Shape{
			Points: []Point{
				Point{X: 0, Y: 0},
				Point{X: 0, Y: 1},
				Point{X: 1, Y: 0},
				Point{X: 1, Y: 1},
			},
		},
	}

	cmdSize := len(cmdLine)
	rocksSize := len(rocks)
	cmdIndex := 0
	rockIndex := 0
	rockCount := 0
	rockCanMove := false
	towerSize := 1
	sh := Shape{}
	extra := ((1000000000000 - 195) / 1740) * 2724
	mod := ((1000000000000 - 195) % 1740)

	grid[0][1], grid[0][2] = 1, 1

	for rockCount < mod {
		rockCanMove = true
		sh = rocks[rockIndex]
		sh.initCoord(towerSize)
		//fmt.Println(sh.Coord)
		for rockCanMove {
			if cmdLine[cmdIndex] == '<' {
				sh.moveHorizontally(-1, grid)
			} else {
				sh.moveHorizontally(1, grid)
			}
			//fmt.Println(sh.Coord)
			cmdIndex = (cmdIndex + 1) % cmdSize
			hasMoved := sh.moveDown(grid)
			//fmt.Println(sh.Coord)

			if !hasMoved {
				x, y := 0, 0
				top := 0
				for _, p := range sh.Points {
					x = sh.Coord.X + p.X
					y = sh.Coord.Y + p.Y
					if y > top {
						top = y
					}
					grid[y][x] = rockCount + 1
				}
				if towerSize < top+1 {
					//fmt.Println(top + 1 - towerSize)
					towerSize = top + 1
				} else {
					//fmt.Println(0)
				}
				//fmt.Println(towerSize)
				rockCanMove = false
				break
			}
		}
		//Print(grid)
		rockIndex = (rockIndex + 1) % rocksSize
		rockCount++
	}
	//Print(grid)
	// towersize is the height of the final pattern at the end,
	fmt.Println(towerSize)
	// 295 is the height offset from there upwards it stars repeating the patterns every 2724 objects
	fmt.Println(towerSize + extra + 295 - 1)

}
