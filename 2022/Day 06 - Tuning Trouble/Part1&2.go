package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(distinctChars int, line string) {
	for i := distinctChars - 1; i < len(line); i++ {
		appear := make([]int, 26)
		for j := 0; j < distinctChars; j++ {
			ind := i - j
			appear[line[ind]-'a']++
		}
		found := true
		for _, a := range appear {
			if a > 1 {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i + 1)
			return
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		solve(4, line)
		solve(14, line)
	}
}
