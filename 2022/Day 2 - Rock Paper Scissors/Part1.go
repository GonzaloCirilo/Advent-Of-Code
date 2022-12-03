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
	points["X"] = 1
	points["Y"] = 2
	points["Z"] = 3

	outcomePoints := map[string]int{}
	outcomePoints["AX"] = 3
	outcomePoints["AY"] = 6
	outcomePoints["AZ"] = 0
	outcomePoints["BX"] = 0
	outcomePoints["BY"] = 3
	outcomePoints["BZ"] = 6
	outcomePoints["CX"] = 6
	outcomePoints["CY"] = 0
	outcomePoints["CZ"] = 3
	for scanner.Scan() {
		var a, b string
		fmt.Sscanf(scanner.Text(), "%s %s", &a, &b)
		ans += points[b] + outcomePoints[a+b]

	}
	fmt.Println(ans)
}
