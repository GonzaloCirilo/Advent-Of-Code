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
		fmt.Sscanf(line, "%d-%d,%d-%d", &l1, &r1, &l2, &r2)

		if (l1 >= l2 && r1 <= r2) || (l2 >= l1 && r2 <= r1) {
			ans++
		}

	}
	fmt.Println(ans)
}
