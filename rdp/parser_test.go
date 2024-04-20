package main

import "testing"

func TestParse_1(t *testing.T) {
	src := "1+2*3"
	node := Parse(src)
	if node.kind != ND_ADD {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_ADD)
	}
	if node.lhs.val != 1 {
		t.Errorf("node.lhs.val = %d, want %d", node.lhs.val, 1)
	}
	if node.rhs.kind != ND_MUL {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_MUL)
	}
	if node.rhs.lhs.val != 2 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 2)
	}
	if node.rhs.rhs.val != 3 {
		t.Errorf("node.rhs.rhs.val = %d, want %d", node.rhs.rhs.val, 3)
	}
}

func TestParse_2(t *testing.T) {
	src := "1*(2+3)"
	node := Parse(src)

	if node.kind != ND_MUL {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_MUL)
	}
	if node.lhs.val != 1 {
		t.Errorf("node.lhs.val = %d, want %d", node.lhs.val, 1)
	}
	if node.rhs.kind != ND_ADD {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_ADD)
	}
	if node.rhs.lhs.val != 2 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 2)
	}
	if node.rhs.rhs.val != 3 {
		t.Errorf("node.rhs.rhs.val = %d, want %d", node.rhs.rhs.val, 3)
	}
}

func TestParse_3(t *testing.T) {
	src := "1+2*3+4"
	node := Parse(src)

	if node.kind != ND_ADD {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_ADD)
	}
	if node.lhs.kind != ND_ADD {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_ADD)
	}
	if node.lhs.lhs.val != 1 {
		t.Errorf("node.lhs.lhs.val = %d, want %d", node.lhs.lhs.val, 1)
	}
	if node.lhs.rhs.kind != ND_MUL {
		t.Errorf("node.lhs.rhs.kind = %d, want %d", node.lhs.rhs.kind, ND_MUL)
	}
	if node.lhs.rhs.lhs.val != 2 {
		t.Errorf("node.lhs.rhs.lhs.val = %d, want %d", node.lhs.rhs.lhs.val, 2)
	}
	if node.lhs.rhs.rhs.val != 3 {
		t.Errorf("node.lhs.rhs.rhs.val = %d, want %d", node.lhs.rhs.rhs.val, 3)
	}
	if node.rhs.val != 4 {
		t.Errorf("node.rhs.val = %d, want %d", node.rhs.val, 4)
	}
}

func TestParse_4(t *testing.T) {
	src := "2 + 3 * 4 - 5"
	node := Parse(src)
	if node.kind != ND_SUB {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_SUB)
	}
	if node.lhs.kind != ND_ADD {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_ADD)
	}
	if node.rhs.val != 5 {
		t.Errorf("node.rhs.val = %d, want %d", node.rhs.val, 5)
	}
	if node.lhs.lhs.val != 2 {
		t.Errorf("node.lhs.lhs.val = %d, want %d", node.lhs.lhs.val, 2)
	}
	if node.lhs.rhs.kind != ND_MUL {
		t.Errorf("node.lhs.rhs.kind = %d, want %d", node.lhs.rhs.kind, ND_MUL)
	}
	if node.lhs.rhs.lhs.val != 3 {
		t.Errorf("node.lhs.rhs.lhs.val = %d, want %d", node.lhs.rhs.lhs.val, 3)
	}
	if node.lhs.rhs.rhs.val != 4 {
		t.Errorf("node.lhs.rhs.rhs.val = %d, want %d", node.lhs.rhs.rhs.val, 4)
	}
}

func TestParse_5(t *testing.T) {
	src := "1 + (2 * (3 + 4))"
	node := Parse(src)
	if node.kind != ND_ADD {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_ADD)
	}
	if node.lhs.val != 1 {
		t.Errorf("node.lhs.val = %d, want %d", node.lhs.val, 1)
	}
	if node.rhs.kind != ND_MUL {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_MUL)
	}
	if node.rhs.lhs.val != 2 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 2)
	}
	if node.rhs.rhs.kind != ND_ADD {
		t.Errorf("node.rhs.rhs.kind = %d, want %d", node.rhs.rhs.kind, ND_ADD)
	}
	if node.rhs.rhs.lhs.val != 3 {
		t.Errorf("node.rhs.rhs.lhs.val = %d, want %d", node.rhs.rhs.lhs.val, 3)
	}
	if node.rhs.rhs.rhs.val != 4 {
		t.Errorf("node.rhs.rhs.rhs.val = %d, want %d", node.rhs.rhs.rhs.val, 4)
	}
}

func TestParse_6(t *testing.T) {
	src := " 1 + 2 - 3 * 4 "
	node := Parse(src)
	if node.kind != ND_SUB {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_SUB)
	}
	if node.lhs.kind != ND_ADD {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_ADD)
	}
	if node.rhs.kind != ND_MUL {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_MUL)
	}
	if node.lhs.lhs.val != 1 {
		t.Errorf("node.lhs.lhs.val = %d, want %d", node.lhs.lhs.val, 1)
	}
	if node.lhs.rhs.val != 2 {
		t.Errorf("node.lhs.rhs.val = %d, want %d", node.lhs.rhs.val, 2)
	}
	if node.rhs.lhs.val != 3 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 3)
	}
	if node.rhs.rhs.val != 4 {
		t.Errorf("node.rhs.rhs.val = %d, want %d", node.rhs.rhs.val, 4)
	}
}

func TestParse_7(t *testing.T) {
	src := "(1 + 2) * (3 - 4)"
	node := Parse(src)
	if node.kind != ND_MUL {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_MUL)
	}
	if node.lhs.kind != ND_ADD {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_ADD)
	}
	if node.rhs.kind != ND_SUB {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_SUB)
	}
	if node.lhs.lhs.val != 1 {
		t.Errorf("node.lhs.lhs.val = %d, want %d", node.lhs.lhs.val, 1)
	}
	if node.lhs.rhs.val != 2 {
		t.Errorf("node.lhs.rhs.val = %d, want %d", node.lhs.rhs.val, 2)
	}
	if node.rhs.lhs.val != 3 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 3)
	}
	if node.rhs.rhs.val != 4 {
		t.Errorf("node.rhs.rhs.val = %d, want %d", node.rhs.rhs.val, 4)
	}
}

func TestParse_8(t *testing.T) {
	src := "((1 - 2) + 3) * 4"
	node := Parse(src)
	if node.kind != ND_MUL {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_MUL)
	}
	if node.lhs.kind != ND_ADD {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_ADD)
	}
	if node.rhs.val != 4 {
		t.Errorf("node.rhs.val = %d, want %d", node.rhs.val, 4)
	}
	if node.lhs.lhs.kind != ND_SUB {
		t.Errorf("node.lhs.lhs.kind = %d, want %d", node.lhs.lhs.kind, ND_SUB)
	}
	if node.lhs.rhs.val != 3 {
		t.Errorf("node.lhs.rhs.val = %d, want %d", node.lhs.rhs.val, 3)
	}
	if node.lhs.lhs.lhs.val != 1 {
		t.Errorf("node.lhs.lhs.lhs.val = %d, want %d", node.lhs.lhs.lhs.val, 1)
	}
	if node.lhs.lhs.rhs.val != 2 {
		t.Errorf("node.lhs.lhs.rhs.val = %d, want %d", node.lhs.lhs.rhs.val, 2)
	}
}

func TestParse_9(t *testing.T) {
	src := "5 - (6 - 3) + 2"
	node := Parse(src)
	if node.kind != ND_ADD {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_ADD)
	}
	if node.lhs.kind != ND_SUB {
		t.Errorf("node.lhs.kind = %d, want %d", node.lhs.kind, ND_SUB)
	}
	if node.rhs.val != 2 {
		t.Errorf("node.rhs.val = %d, want %d", node.rhs.val, 2)
	}
	if node.lhs.lhs.val != 5 {
		t.Errorf("node.lhs.lhs.val = %d, want %d", node.lhs.lhs.val, 5)
	}
	if node.lhs.rhs.kind != ND_SUB {
		t.Errorf("node.lhs.rhs.kind = %d, want %d", node.lhs.rhs.kind, ND_SUB)
	}
	if node.lhs.rhs.lhs.val != 6 {
		t.Errorf("node.lhs.rhs.lhs.val = %d, want %d", node.lhs.rhs.lhs.val, 6)
	}
	if node.lhs.rhs.rhs.val != 3 {
		t.Errorf("node.lhs.rhs.rhs.val = %d, want %d", node.lhs.rhs.rhs.val, 3)
	}
}

func TestParse_10(t *testing.T) {
	src := "3 * (4 + 5)"
	node := Parse(src)
	if node.kind != ND_MUL {
		t.Errorf("node.kind = %d, want %d", node.kind, ND_MUL)
	}
	if node.lhs.val != 3 {
		t.Errorf("node.lhs.val = %d, want %d", node.lhs.val, 3)
	}
	if node.rhs.kind != ND_ADD {
		t.Errorf("node.rhs.kind = %d, want %d", node.rhs.kind, ND_ADD)
	}
	if node.rhs.lhs.val != 4 {
		t.Errorf("node.rhs.lhs.val = %d, want %d", node.rhs.lhs.val, 4)
	}
	if node.rhs.rhs.val != 5 {
		t.Errorf("node.rhs.rhs.val = %d, want %d", node.rhs.rhs.val, 5)
	}
}
