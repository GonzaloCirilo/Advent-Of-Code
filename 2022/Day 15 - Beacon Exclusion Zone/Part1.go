package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	X, Y int
}

func ManhattanDist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	manhattan := make([]int, 0)
	sensorCoord, beaconCoord := make([]Point, 0), make(map[int]int, 0)
	blocked := make(map[int]int, 0)
	target := 2000000

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		manhattan = append(manhattan, ManhattanDist(sx, sy, bx, by))
		sensorCoord = append(sensorCoord, Point{X: sx, Y: sy})
		if by == target {
			beaconCoord[bx] = 1
		}

	}
	for i, p := range sensorCoord {
		dif := int(math.Abs(float64(target - p.Y)))
		for j := 0; j <= (manhattan[i] - dif); j++ {
			blocked[p.X+j] = 1
			blocked[p.X-j] = 1
		}
	}
	fmt.Println(len(blocked) - len(beaconCoord))
}
