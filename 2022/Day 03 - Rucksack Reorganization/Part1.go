package main

import (
	"bufio"
	"fmt"
	"os"
)

// returns array of unique appearances of each character
func getUniqueAppearances(line string) []int {
	appearances := make([]int, 52)
	for _, token := range line {
		ind := int(token) - int('a')
		if ind < 0 {
			ind = int(token) - int('A') + 26
		}
		appearances[ind] = 1
	}
	return appearances
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		left := line[:len(line)/2]
		right := line[len(line)/2:]
		appearLeft := getUniqueAppearances(left)
		appearRight := getUniqueAppearances(right)
		repeat := -1
		for i, a := range appearLeft {
			if a+appearRight[i] > 1 {
				repeat = i
			}
		}
		ans += repeat + 1
	}
	fmt.Println(ans)
}
