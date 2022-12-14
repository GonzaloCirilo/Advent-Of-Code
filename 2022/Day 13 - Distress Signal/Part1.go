package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Elements []Node
	Value    *int
}

func (n Node) ToString() string {
	aux := "{"
	if n.Value == nil {
		for _, n2 := range n.Elements {
			aux += n2.ToString()
		}
	} else {
		aux += strconv.Itoa(*n.Value) + " "
	}
	aux += "}"
	return aux
}

func compare(nodeA, nodeB Node) int {

	if nodeA.Value != nil && nodeB.Value != nil {
		if *nodeA.Value < *nodeB.Value {
			return 1
		} else if *nodeA.Value == *nodeB.Value {
			return 0
		} else {
			return -1
		}
	}

	if nodeA.Value == nil && nodeB.Value != nil {
		return compare(nodeA, Node{Elements: []Node{nodeB}})
	}

	if nodeA.Value != nil && nodeB.Value == nil {
		return compare(Node{Elements: []Node{nodeA}}, nodeB)
	}

	for i := range nodeA.Elements {
		if i < len(nodeB.Elements) {
			res := compare(nodeA.Elements[i], nodeB.Elements[i])
			if res == 1 {
				return 1
			} else if res == -1 {
				return -1
			}
		}
	}
	fmt.Println(len(nodeA.Elements), len(nodeB.Elements))
	if len(nodeA.Elements) < len(nodeB.Elements) {
		return 1
	} else if len(nodeA.Elements) == len(nodeB.Elements) {
		return 0
	} else {
		return -1
	}
}

func Parse(s string) Node {
	//fmt.Println(s)
	x, e := strconv.Atoi(s)
	if e == nil {
		return Node{Value: &x}
	} else {
		capture := false
		open := 0
		aux := ""
		node := Node{Elements: []Node{}}
		for i := 0; i < len(s); i++ {
			if capture {
				aux += string(s[i])
			} else {
				_, e := strconv.Atoi(string(s[i]))
				ind := strings.Index(s[i:], ",")
				if ind == -1 {
					ind = len(s)
				} else {
					ind += i
				}
				if e == nil {
					node.Elements = append(node.Elements, Parse(s[i:ind]))
					i = ind
					continue
				}
			}
			if s[i] == '[' {
				capture = true
				open++
			} else if s[i] == ']' {
				if capture && open == 1 {
					capture = false
					node.Elements = append(node.Elements, Parse(aux[:len(aux)-1]))
					aux = ""
				}
				open--
			}

		}
		return node
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//treeA := make([]Node, 0)
	//treeB := make([]Node, 0)
	count := 1
	ans := 0
	for scanner.Scan() {
		pairA := scanner.Text()
		scanner.Scan()
		pairB := scanner.Text()
		scanner.Scan()
		treeA := Parse(pairA[1 : len(pairA)-1])
		treeB := Parse(pairB[1 : len(pairB)-1])
		//fmt.Println(treeA.ToString(), treeB.ToString())

		if compare(treeA, treeB) == 1 {
			fmt.Println(count)
			ans += count
		}
		count++
	}
	fmt.Println(ans)
}
