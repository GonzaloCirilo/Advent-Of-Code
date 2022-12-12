package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

type Monkey struct {
	Items     []int64
	Operation string
	OpAmount  int64
	DivTest   int64
	TrueTest  int64
	FalseTest int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := make([]Monkey, 0)
	monkeyCount := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()

		scanner.Scan()
		line = scanner.Text()
		aux := make([]int64, 0)
		for _, t := range strings.Split(line[18:], ",") {
			t = strings.TrimSpace(t)
			n, _ := strconv.Atoi(t)
			ne := int64(n)
			aux = append(aux, ne)
		}

		scanner.Scan()
		line = scanner.Text()
		var op, s string
		var amount, divTest, trueTest, falseTest int64
		fmt.Sscanf(line[23:], "%s %s", &op, &s)

		a, e := strconv.Atoi(s)
		amount = int64(a)
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

	loop := 10000

	mod := int64(1)
	// need to multiply all div tests values to do not get wrong values when doing module to worry levels
	for _, e := range monkeys {
		mod *= e.DivTest
	}

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

				monkeys[j].Items[i] = monkeys[j].Items[i] % mod

				if monkeys[j].Items[i]%monkeys[j].DivTest == 0 {
					monkeys[monkeys[j].TrueTest].Items = append(monkeys[monkeys[j].TrueTest].Items, monkeys[j].Items[i])
				} else {
					monkeys[monkeys[j].FalseTest].Items = append(monkeys[monkeys[j].FalseTest].Items, monkeys[j].Items[i])
				}
				monkeyCount[j]++
			}
			monkeys[j].Items = make([]int64, 0)
			//fmt.Print64ln("++++", monkeys)
		}
		//fmt.Print64ln(monkeys)
	}

	sort.Slice(monkeyCount, func(i, j int) bool {
		return monkeyCount[i] < monkeyCount[j]
	})
	fmt.Println(monkeyCount)
	fmt.Println(monkeyCount[len(monkeyCount)-1] * monkeyCount[len(monkeyCount)-2])
}
