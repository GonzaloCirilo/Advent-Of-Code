package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ans = int64(0)

func calculateDirSizes(node *Node) int64 {
	if node.Files == nil {
		return node.Size
	}
	aux := int64(0)
	for _, v := range node.Files {
		aux += calculateDirSizes(v)
	}
	node.Size = aux
	if node.Size <= 100000 {
		ans += (node.Size)
	}
	return node.Size
}

type Node struct {
	Parent *Node
	Files  map[string]*Node
	Size   int64
	Name   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	currentNode := &Node{Files: map[string]*Node{}}
	root := currentNode
	for scanner.Scan() {
		line := scanner.Text()
		var cmd, current string
		if line[0] == '$' {
			fmt.Sscanf(line, "$ %s", &cmd)
			if cmd == "cd" {
				current = line[5:]
				if current == "/" {
					continue
				}
				if current == ".." {
					current = currentNode.Parent.Name
					currentNode = currentNode.Parent
				} else {
					currentNode = currentNode.Files[current]
				}
			}
		} else {
			var tk1, tk2 string
			fmt.Sscanf(line, "%s %s", &tk1, &tk2)
			if tk1 == "dir" {
				currentNode.Files[tk2] = &Node{Name: tk2, Files: map[string]*Node{}, Parent: currentNode}
			} else {
				size, _ := strconv.Atoi(tk1)
				currentNode.Files[tk2] = &Node{Size: int64(size), Name: tk2, Parent: currentNode}
			}
		}

	}
	calculateDirSizes(root)
	fmt.Println(ans)
}
