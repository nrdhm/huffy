package core

import (
	"container/heap"
	"fmt"
	"strings"
)

// Symbol is some part of the text to compress
type Symbol string

// Code is a short string mapped to a Symbol
type Code string

// HuffTree is a Huffman tree used to construct compression codes
type HuffTree struct {
	Zero *HuffTree
	One  *HuffTree
	Sym  Symbol
	Cnt  int
	EOF  bool
}

// NewTree creates a new Huffman tree from a given table of symbols occurences
func NewTree(symbolsCnt []SymbolCount) *HuffTree {
	if len(symbolsCnt) == 0 {
		return nil
	}
	trees := treeHeap{&HuffTree{EOF: true, Cnt: 1}}
	for _, sc := range symbolsCnt {
		tree := &HuffTree{Zero: nil, One: nil, Sym: sc.Symbol, Cnt: sc.Count}
		trees = append(trees, tree)
	}
	heap.Init(&trees)
	for len(trees) > 1 {
		first := heap.Pop(&trees).(*HuffTree)
		second := heap.Pop(&trees).(*HuffTree)
		united := &HuffTree{Zero: second, One: first, Cnt: first.Cnt + second.Cnt}
		heap.Push(&trees, united)
	}
	return trees[0]
}

// GetSym returns symbol the node is holding
func (ht *HuffTree) GetSym() string {
	if ht.EOF {
		return "<EOF>"
	}
	return string(ht.Sym)
}

func (ht *HuffTree) String() string {
	if ht == nil {
		return ""
	}
	zero, one := "", ""
	if ht.Zero != nil {
		zero = fmt.Sprintf("zero:%s", ht.Zero)
	}
	if ht.One != nil {
		one = fmt.Sprintf("one:%s", ht.One)
	}
	eof := ""
	if ht.EOF {
		eof = " EOF "
	}
	return fmt.Sprintf("<%q %s %s %s>", ht.Sym, eof, zero, one)
}

func (ht *HuffTree) equal(other *HuffTree) bool {
	// if both are nil, they are the same
	if ht == nil && other == nil {
		return true
	}
	// otherwise, if only one of them is nil, they differ
	if ht == nil || other == nil {
		return false
	}
	if ht.Sym == other.Sym && ht.EOF == other.EOF {
		return ht.Zero.equal(other.Zero) && ht.One.equal(other.One)
	}
	return false
}

type treeHeap []*HuffTree

func (h treeHeap) Len() int           { return len(h) }
func (h treeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h treeHeap) Less(i, j int) bool { return h[i].Cnt < h[j].Cnt }

func (h *treeHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffTree))
}

func (h *treeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h treeHeap) String() string {
	fs := make([]string, len(h))
	for i, tree := range h {
		fs[i] = fmt.Sprintf("{%q, %v}", tree.Sym, tree.Cnt)
	}
	return strings.Join(fs, " | ")
}

func getTable(tree *HuffTree, codePrefix string) map[Symbol]Code {
	m := map[Symbol]Code{}
	if tree.Sym != "" || tree.EOF {
		m[tree.Sym] = Code(codePrefix)
	}
	if tree.Zero != nil {
		zero := getTable(tree.Zero, codePrefix+"0")
		m = merge(m, zero)
	}
	if tree.One != nil {
		one := getTable(tree.One, codePrefix+"1")
		m = merge(m, one)
	}
	return m
}

// GetCompressTable returns table with Huffman codes used to compress some text
func GetCompressTable(tree *HuffTree) map[Symbol]Code {
	return getTable(tree, "")
}
