package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := 0
	points := map[string]int{}
	points["X"] = 0
	points["Y"] = 3
	points["Z"] = 6

	outcomePoints := map[string]int{}
	outcomePoints["AX"] = 3
	outcomePoints["AY"] = 1
	outcomePoints["AZ"] = 2
	outcomePoints["BX"] = 1
	outcomePoints["BY"] = 2
	outcomePoints["BZ"] = 3
	outcomePoints["CX"] = 2
	outcomePoints["CY"] = 3
	outcomePoints["CZ"] = 1
	for scanner.Scan() {
		var a, b string
		fmt.Sscanf(scanner.Text(), "%s %s", &a, &b)
		ans += points[b] + outcomePoints[a+b]

	}
	fmt.Println(ans)
}
