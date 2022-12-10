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
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		c := make(chan []int, 3)
		go getUniqueAppearances(line1, c)
		go getUniqueAppearances(line2, c)
		go getUniqueAppearances(line3, c)
		appearL1, appearL2, appearL3 := <-c, <-c, <-c
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
