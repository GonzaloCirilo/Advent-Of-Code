package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getPixelvalue(pos, Xpos int) rune {
	if math.Abs(float64(Xpos)-float64(pos)) <= 1 {
		return '#'
	} else {
		return '.'
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cycles := make([]int, 500)
	crt := make([]rune, 240)
	X := 1
	cycles[0] = X
	currentCycle := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		var cmd string
		pixelPos := currentCycle - 1
		if line[0] == 'a' {
			var amount int
			fmt.Sscanf(line, "%s %d", &cmd, &amount)
			cycles[currentCycle] = X
			crt[pixelPos] = getPixelvalue(pixelPos%40, X)
			cycles[currentCycle+1] = X
			crt[pixelPos+1] = getPixelvalue((pixelPos+1)%40, X)
			X += amount
			currentCycle += 2
		} else {
			cycles[currentCycle] = X
			crt[pixelPos] = getPixelvalue((pixelPos)%40, X)
			currentCycle++
		}

	}
	for i, c := range crt {
		if i%40 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Print(string(c))
	}
}
