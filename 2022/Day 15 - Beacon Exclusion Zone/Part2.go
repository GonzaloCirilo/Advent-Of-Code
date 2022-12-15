package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	X, Y, DistToB int
}

func ManhattanDist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sensorCoord := make([]Point, 0)
	limit := 4000000
	for scanner.Scan() {
		line := scanner.Text()
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensorCoord = append(sensorCoord, Point{X: sx, Y: sy, DistToB: ManhattanDist(sx, sy, bx, by)})

	}
	sort.Slice(sensorCoord, func(i, j int) bool {
		return sensorCoord[i].X < sensorCoord[j].X
	})
	for target := 0; target <= limit; target++ {
		x := 0
		ind := 0
		for {
			if ind >= len(sensorCoord) || x > limit {
				//fmt.Println("++", ind, x)
				break
			}
			dif := int(math.Abs(float64(target - sensorCoord[ind].Y)))
			if dif > sensorCoord[ind].DistToB {
				ind++
				continue
			}

			right := (sensorCoord[ind].X + (sensorCoord[ind].DistToB - dif))
			left := (sensorCoord[ind].X - (sensorCoord[ind].DistToB - dif))

			//fmt.Println("bounds", l, r, ind, dif, sensorCoord[ind].DistToB)
			if x < left {
				ind++
				continue
			}
			if x <= right {
				x = right + 1
			}
			ind++

		}
		if x <= limit {
			fmt.Println(x*4000000 + target /*, x, target*/)
			break
		}
	}
}
