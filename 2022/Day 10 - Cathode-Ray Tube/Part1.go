package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cycles := make([]int, 500)
	X := 1
	cycles[0] = X
	currentCycle := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		var cmd string
		if line[0] == 'a' {
			var amount int
			fmt.Sscanf(line, "%s %d", &cmd, &amount)
			cycles[currentCycle] = X
			cycles[currentCycle+1] = X
			X += amount
			//fmt.Println(X, amount)
			currentCycle += 2
		} else {
			cycles[currentCycle] = X
			currentCycle++
		}

	}
	fmt.Println(cycles[20]*20 + cycles[60]*60 + cycles[100]*100 + cycles[140]*140 + cycles[180]*180 + cycles[220]*220)
}
