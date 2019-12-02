package main

import (
	"fmt"
	"strings"
	"strconv"
)

func getFirst(codes []int) int {
	for i := 0; i < len(codes) - 4; i += 4 { 
		target := codes[i+3]
		a := codes[i+1]
		b := codes[i+2]
		opcode := codes[i]

		if opcode == 1 {
			codes[target] = codes[a] + codes[b]
		}else{
			if opcode == 2 {
				codes[target] = codes[a] * codes[b]
			}else{
				if opcode == 99{
					break;
				}
			}
		}
	} 
	return codes[0]
}

func simulate(codes []int) int{
	aux := make([]int, len(codes))
	copy(aux, codes)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++{
			aux[1] = i;
			aux[2] = j;
			copy(codes, aux)
			if(getFirst(codes) == 19690720){
				return i * 100 + j
			}
			
		}
	}
	return 0
}

func main(){

	s := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,6,19,23,1,13,23,27,1,6,27,31,1,31,10,35,1,35,6,39,1,39,13,43,2,10,43,47,1,47,6,51,2,6,51,55,1,5,55,59,2,13,59,63,2,63,9,67,1,5,67,71,2,13,71,75,1,75,5,79,1,10,79,83,2,6,83,87,2,13,87,91,1,9,91,95,1,9,95,99,2,99,9,103,1,5,103,107,2,9,107,111,1,5,111,115,1,115,2,119,1,9,119,0,99,2,0,14,0"
	//s := "1,9,10,3,2,3,11,0,99,30,40,50"
	var codes []int
	n := strings.Split(s, ",")

	for _, numberS := range n {
		number, _ := strconv.Atoi(numberS)
		codes = append(codes, number)	
	}

	fmt.Println(simulate(codes))
	
}