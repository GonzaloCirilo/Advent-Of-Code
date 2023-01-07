package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type N struct {
	X, Ind int
}

func findFirst(arr []N, e N) int {
	for i, n := range arr {
		if n.Ind == e.Ind && n.X == e.X {
			return i
		}
	}
	return -1
}

func appendAt(arr []N, ind int, x N) []N {
	aux := make([]N, 0)
	for i, n := range arr {
		if i == ind {
			aux = append(aux, x)
		}
		aux = append(aux, n)
	}
	return aux
}

func removeAt(arr []N, ind int) []N {
	aux := make([]N, 0)
	for i, n := range arr {
		if i == ind {
			continue
		}
		aux = append(aux, n)
	}
	return aux
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	initial, aux := make([]N, 0), make([]N, 0)
	comp := N{}
	cont := 0
	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.Atoi(line)
		fmt.Println(x * 811589153)
		auxX := N{X: x * 811589153, Ind: cont}
		initial = append(initial, auxX)
		aux = append(aux, auxX)
		fmt.Println(auxX)
		if x == 0 {
			comp = auxX
		}
		cont++
	}
	for m := 0; m < 10; m++ {
		for _, n := range initial {
			ind := findFirst(aux, n)
			if ind == -1 || n.X == 0 {
				continue
			}
			newInd := (ind + n.X + (len(initial)-1)*8200000000) % (len(initial) - 1)
			if newInd > 5000 || newInd < 0 {
				fmt.Println("---", newInd)
			}
			aux = removeAt(aux, ind)
			aux = appendAt(aux, int(newInd), n)
			if m == 0 {
				fmt.Println(len(aux))

			}
		}
	}

	offset := findFirst(aux, comp)
	fmt.Println(aux[(1000+offset)%len(aux)].X, len(aux), aux[(3000+offset)%len(aux)].X)
	ans := aux[(1000+offset)%len(aux)].X + aux[(2000+offset)%len(aux)].X + aux[(3000+offset)%len(aux)].X
	fmt.Println(ans)
}
