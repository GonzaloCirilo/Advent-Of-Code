package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stacks := make([]list.List, 9)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}
		if line[1] == '1' {
			continue
		} else if line[0] == 'm' {
			var amount, from, to int
			fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
			from--
			to--
			aux := list.List{}
			for i := 0; i < amount; i++ {
				ele := stacks[from].Back()
				aux.PushFront(ele.Value)
				stacks[from].Remove(ele)
			}
			stacks[to].PushBackList(&aux)
		} else {
			for i := 0; i+1+i*3 < len(line); i++ {
				ind := i + 1 + i*3
				if line[ind] != ' ' {
					stacks[i].PushFront(line[ind])
				}
			}
		}
	}
	for _, s := range stacks {
		if s.Back() != nil {
			fmt.Printf("%c", s.Back().Value)
		}

	}
}
