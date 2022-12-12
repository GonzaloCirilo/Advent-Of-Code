package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation string
	OpAmount  int
	DivTest   int
	TrueTest  int
	FalseTest int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := make([]Monkey, 0)
	monkeyCount := make([]int, 0)
	diver := 3
	for scanner.Scan() {
		line := scanner.Text()

		scanner.Scan()
		line = scanner.Text()
		aux := make([]int, 0)
		for _, t := range strings.Split(line[18:], ",") {
			t = strings.TrimSpace(t)
			n, _ := strconv.Atoi(t)
			aux = append(aux, n)
		}

		scanner.Scan()
		line = scanner.Text()
		var op, s string
		var amount, divTest, trueTest, falseTest int
		fmt.Sscanf(line[23:], "%s %s", &op, &s)

		amount, e := strconv.Atoi(s)
		if e != nil {
			amount = -1
		}

		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line[21:], "%d", &divTest)

		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line[29:], "%d", &trueTest)

		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line[30:], "%d", &falseTest)

		scanner.Scan()

		monkeys = append(monkeys, Monkey{Items: aux, Operation: op, OpAmount: amount, DivTest: divTest, TrueTest: trueTest, FalseTest: falseTest})
		monkeyCount = append(monkeyCount, 0)
	}

	loop := 20

	for l := 0; l < loop; l++ {
		for j := range monkeys {
			for i := range monkeys[j].Items {
				if monkeys[j].Operation == "*" {
					if monkeys[j].OpAmount == -1 {
						monkeys[j].Items[i] = monkeys[j].Items[i] * monkeys[j].Items[i]
					} else {
						monkeys[j].Items[i] = monkeys[j].Items[i] * monkeys[j].OpAmount
					}
				} else {
					if monkeys[j].OpAmount == -1 {
						monkeys[j].Items[i] = monkeys[j].Items[i] + monkeys[j].Items[i]
					} else {
						monkeys[j].Items[i] = monkeys[j].Items[i] + monkeys[j].OpAmount
					}
				}

				monkeys[j].Items[i] = int(math.Floor(float64(monkeys[j].Items[i]) / float64(diver)))

				if monkeys[j].Items[i]%monkeys[j].DivTest == 0 {
					monkeys[monkeys[j].TrueTest].Items = append(monkeys[monkeys[j].TrueTest].Items, monkeys[j].Items[i])
				} else {
					monkeys[monkeys[j].FalseTest].Items = append(monkeys[monkeys[j].FalseTest].Items, monkeys[j].Items[i])
				}
				monkeyCount[j]++
			}
			monkeys[j].Items = make([]int, 0)
			//fmt.Println("++++", monkeys)
		}
		//fmt.Println(monkeys)
	}

	sort.Slice(monkeyCount, func(i, j int) bool {
		return monkeyCount[i] < monkeyCount[j]
	})
	fmt.Println(monkeyCount[len(monkeyCount)-1] * monkeyCount[len(monkeyCount)-2])
}
