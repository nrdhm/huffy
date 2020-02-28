package core

import (
	"container/heap"
	"fmt"
	"strings"
)

type Symbol string
type Code string

// HuffmanTree is a binary tree
type HuffmanTree struct {
	zeroBranch *HuffmanTree
	oneBranch  *HuffmanTree
	symbol     Symbol
	occurrence int
	eof        bool
}

func NewTree(letters map[Symbol]int) *HuffmanTree {
	if len(letters) == 0 {
		return nil
	}
	trees := TreeHeap{&HuffmanTree{eof: true}}
	for letter, occurrence := range letters {
		tree := &HuffmanTree{zeroBranch: nil, oneBranch: nil, symbol: letter, occurrence: occurrence}
		trees = append(trees, tree)
	}
	heap.Init(&trees)
	for len(trees) > 1 {
		first := heap.Pop(&trees).(*HuffmanTree)
		second := heap.Pop(&trees).(*HuffmanTree)
		united := &HuffmanTree{zeroBranch: second, oneBranch: first, occurrence: first.occurrence + second.occurrence}
		heap.Push(&trees, united)
	}
	return trees[0]
}

func (ht *HuffmanTree) String() string {
	if ht == nil {
		return ""
	}
	zero, one := "", ""
	if ht.zeroBranch != nil {
		zero = fmt.Sprintf("zero:%s", ht.zeroBranch)
	}
	if ht.oneBranch != nil {
		one = fmt.Sprintf("one:%s", ht.oneBranch)
	}
	eof := ""
	if ht.eof {
		eof = " EOF "
	}
	return fmt.Sprintf("<%q %s %s %s>", ht.symbol, eof, zero, one)
}

func (ht *HuffmanTree) equal(other *HuffmanTree) bool {
	// if both are nil, they are the same
	if ht == nil && other == nil {
		return true
	}
	// otherwise, if only one of them is nil, they differ
	if ht == nil || other == nil {
		return false
	}
	if ht.symbol == other.symbol && ht.eof == other.eof {
		return ht.zeroBranch.equal(other.zeroBranch) && ht.oneBranch.equal(other.oneBranch)
	}
	return false
}

func addCodeSymbol(tree *HuffmanTree, code Code, symbol Symbol) *HuffmanTree {
	// tree is either nil or pointer to a tree
	// code is either empty string or contains 0s and 1s
	// symbol is just a string
	if tree == nil {
		tree = &HuffmanTree{}
	}
	if len(code) == 0 {
		tree.symbol = symbol
		if symbol == "" {
			tree.eof = true
		}
		return tree
	}
	if code[0] == '0' {
		tree.zeroBranch = addCodeSymbol(tree.zeroBranch, code[1:], symbol)
	}
	if code[0] == '1' {
		tree.oneBranch = addCodeSymbol(tree.oneBranch, code[1:], symbol)
	}
	return tree
}

func FromCompressTable(table map[Symbol]Code) *HuffmanTree {
	var tree *HuffmanTree
	for symbol, code := range table {
		tree = addCodeSymbol(tree, code, symbol)
	}
	return tree
}

type TreeHeap []*HuffmanTree

func (h TreeHeap) Len() int           { return len(h) }
func (h TreeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h TreeHeap) Less(i, j int) bool { return h[i].occurrence < h[j].occurrence }

func (h *TreeHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanTree))
}

func (h *TreeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h TreeHeap) String() string {
	fs := make([]string, len(h))
	for i, tree := range h {
		fs[i] = fmt.Sprintf("{%q, %v}", tree.symbol, tree.occurrence)
	}
	return strings.Join(fs, " | ")
}

func getTable(tree *HuffmanTree, codePrefix string) map[Symbol]Code {
	m := map[Symbol]Code{}
	if tree.symbol != "" || tree.eof {
		m[tree.symbol] = Code(codePrefix)
	}
	if tree.zeroBranch != nil {
		zero := getTable(tree.zeroBranch, codePrefix+"0")
		m = merge(m, zero)
	}
	if tree.oneBranch != nil {
		one := getTable(tree.oneBranch, codePrefix+"1")
		m = merge(m, one)
	}
	return m
}

func GetCompressTable(tree *HuffmanTree) map[Symbol]Code {
	return getTable(tree, "")
}
