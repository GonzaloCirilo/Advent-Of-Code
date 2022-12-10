package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	max := 0
	current := 0
	for scanner.Scan() {
		text := scanner.Text()
		e, _ := strconv.Atoi(text)
		if len(text) <= 0 {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += e
		}
	}

	if max < current {
		max = current
		current = 0
	}

	fmt.Println(max)
}
