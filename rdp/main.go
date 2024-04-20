package main

import "fmt"

func main() {
	src := "1+2*3+4"
	node := Parse(src)
	fmt.Println(node)
	fmt.Println(node.kind == ND_ADD)
	fmt.Println(node.lhs)
	fmt.Println(node.rhs)
	fmt.Println(node.rhs.kind == ND_MUL)
	fmt.Println(node.rhs.lhs)
	fmt.Println(node.rhs.rhs)
}
