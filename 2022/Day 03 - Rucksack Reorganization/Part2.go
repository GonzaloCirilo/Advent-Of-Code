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
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		appearL1 := getUniqueAppearances(line1)
		appearL2 := getUniqueAppearances(line2)
		appearL3 := getUniqueAppearances(line3)
		repeat := -1
		for i, a := range appearL1 {
			if a+appearL2[i]+appearL3[i] > 2 {
				repeat = i
			}
		}
		ans += repeat + 1
	}
	fmt.Println(ans)
}
