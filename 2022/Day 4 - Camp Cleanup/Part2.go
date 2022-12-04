package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		var l1, r1, l2, r2 int
		fmt.Sscanf(line, "%d-%d,%d-%d\n", &l1, &r1, &l2, &r2)
		assig := make([]int, 121)
		for i := l1; i <= r1; i++ {
			assig[i]++
		}

		for i := l2; i <= r2; i++ {
			assig[i]++
		}

		for i := 0; i < 121; i++ {
			if assig[i] > 1 {
				ans++
				break
			}
		}

	}
	fmt.Println(ans)
}
