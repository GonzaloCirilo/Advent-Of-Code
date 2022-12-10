package main

import (
	"bufio"
	"fmt"
	"os"
)

// returns array of unique appearances of each character
func getUniqueAppearances(line string, c chan []int) {
	appearances := make([]int, 52)
	for _, token := range line {
		ind := int(token) - int('a')
		if ind < 0 {
			ind = int(token) - int('A') + 26
		}
		appearances[ind] = 1
	}
	c <- appearances
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := 0
	c := make(chan []int)
	for scanner.Scan() {
		line := scanner.Text()
		left := line[:len(line)/2]
		right := line[len(line)/2:]
		go getUniqueAppearances(left, c)
		go getUniqueAppearances(right, c)
		seg1, seg2 := <-c, <-c
		repeat := -1
		for i, a := range seg1 {
			if a+seg2[i] > 1 {
				repeat = i
			}
		}
		ans += repeat + 1
	}
	fmt.Println(ans)
}
