package main

import (
	"strconv"
	"unicode"
)

type NodeKind int

const (
	ND_ADD NodeKind = iota
	ND_SUB
	ND_MUL
	ND_DIV
	ND_NUM
)

type Node struct {
	kind NodeKind // Node の種類
	lhs  *Node    // 左の子ノード
	rhs  *Node    // 右の子ノード
	val  int      // kind が ND_NUM の場合のみ値が入る
}

func newNode(kind NodeKind, lhs, rhs *Node) *Node {
	return &Node{
		kind: kind,
		lhs:  lhs,
		rhs:  rhs,
	}
}

func newNumberNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func is(src string, op byte) bool {
	if len(src) == 0 {
		return false
	}

	return src[0] == op
}

// expr = mul ("+" mul | "-" mul)*
// expr は mul と mul の足し算、引き算の0回以上の繰り返し
func expr(src string) (*Node, string) {
	var node *Node
	if len(src) == 0 {
		return nil, src
	}

	node, src = mul(src)

	for {
		if is(src, ' ') {
			src = src[1:] // consume(' ')
			continue
		}

		if is(src, '+') {
			src = src[1:] // consume('+')
			var m *Node
			m, src = mul(src)
			node = newNode(ND_ADD, node, m) // expr = mul "+" mul
			continue
		}

		if is(src, '-') {
			src = src[1:] // consume('-')
			var m *Node
			m, src = mul(src)
			node = newNode(ND_SUB, node, m) // expr = mul "-" mul
			continue
		}

		return node, src
	}
}

// mul = primary ("*" primary | "/" primary)*
// mul は primary と primary の掛け算、割り算の0回以上の繰り返し
func mul(src string) (*Node, string) {
	var node *Node
	if len(src) == 0 {
		return nil, src
	}

	node, src = primary(src)

	for {
		if is(src, ' ') {
			src = src[1:] // consume(' ')
			continue
		}

		if is(src, '*') {
			src = src[1:] // consume('*')
			var p *Node
			p, src = primary(src)
			node = newNode(ND_MUL, node, p) // mul = primary "*" primary
			continue
		}

		if is(src, '/') {
			src = src[1:] // consume('/')
			var p *Node
			p, src = primary(src)
			node = newNode(ND_DIV, node, p) // mul = primary "/" primary
			continue
		}

		return node, src
	}
}

func primary(src string) (*Node, string) {
	if is(src, ' ') {
		src = src[1:] // consume(' ')
		return primary(src)
	}

	var node *Node
	if len(src) == 0 {
		return nil, src
	}

	if is(src, '(') {
		src = src[1:] // consume('(')
		node, src = expr(src)
		if is(src, ')') {
			src = src[1:] // expect('(')
		}

		return node, src
	}

	i := getLastDigitIndex(src)
	v, err := strconv.Atoi(src[:i])
	if err != nil {
		panic(err)
	}

	return newNumberNode(v), src[i:]
}

func getLastDigitIndex(src string) int {
	i := 0
	for i < len(src) {
		if !unicode.IsDigit(rune(src[i])) {
			break
		}
		i++
	}

	return i
}

func Parse(src string) *Node {
	node, _ := expr(src)
	return node
}
