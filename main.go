package main

import (
	"flag"
	"fmt"
	"strconv"
	"unicode"
)

func init() {
	fmt.Printf(".intel_syntax noprefix\n")
}

type Kind string

const (
	Reserved Kind = "Reserved"
	Number   Kind = "Number"
	EOF      Kind = "EOF"
)

type Token struct {
	Kind  Kind
	Value int
	Name  byte
	Next  *Token
}

func (k Kind) IsReserved() bool {
	return k == Reserved
}

func (k Kind) IsNumber() bool {
	return k == Number
}

func (k Kind) IsEOF() bool {
	return k == EOF
}

func NewToken(k Kind, v int, name byte, cur *Token) *Token {
	t := &Token{
		Kind:  k,
		Value: v,
		Name:  name,
		Next:  nil,
	}
	cur.Next = t
	return t
}

func Tokenize(src string) *Token {
	head := new(Token)
	cur := head
	i := 0
	for i < len(src) {
		if isSpace(rune(src[i])) {
			i++
			continue
		}

		if src[i] == '+' || src[i] == '-' {
			cur = NewToken(Reserved, -1, src[i], cur)
		}

		if isDigit(rune(src[i])) {
			end := getLastDigitIndex(src, i)
			v, err := strconv.Atoi(src[i:end])
			if err != nil {
				panic(err)
			}

			cur = NewToken(Number, v, src[i], cur)
			i = end
			continue
		}

		i++
	}

	NewToken(EOF, -1, byte(0), cur)

	return head.Next
}

func isSpace(o rune) bool {
	return o == ' ' || o == '\t'
}

func isDigit(o rune) bool {
	return unicode.IsDigit(o)
}

func getLastDigitIndex(src string, start int) int {
	for start < len(src) {
		if !isDigit(rune(src[start])) {
			return start
		}
		start++
	}
	return start
}

func consume(t *Token) *Token {
	return t.Next
}

func main() {
	flag.Parse()
	src := flag.Arg(0)
	t := Tokenize(src)

	fmt.Printf(".globl _main\n")
	fmt.Printf("_main:\n")
	fmt.Printf("  mov rax, %d\n", t.Value)
	t = consume(t)

	for !t.Kind.IsEOF() {
		if t.Kind.IsReserved() && t.Name == '+' {
			t = consume(t)
			fmt.Printf("  add rax, %d\n", t.Value)
			continue
		}

		if t.Kind.IsReserved() && t.Name == '-' {
			t = consume(t)
			fmt.Printf("  sub rax, %d\n", t.Value)
			continue
		}

		t = consume(t)
	}

	fmt.Printf("  ret\n")
}
