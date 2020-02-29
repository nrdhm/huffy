package core

import (
	"bytes"
	"sort"
)

func tokenize(text string) chan Symbol {
	ch := make(chan Symbol)
	go func() {
		for _, r := range text {
			ch <- Symbol(r)
		}
		close(ch)
	}()
	return ch
}

// SymbolCount holds count of a symbol
type SymbolCount struct {
	symbol Symbol
	count  int
}

// CountSymbols counts occurences of symbols
func CountSymbols(text string) []SymbolCount {
	ps := map[Symbol]int{}
	for sym := range tokenize(text) {
		ps[sym]++
	}
	sc := []SymbolCount{}
	for sym, cnt := range ps {
		sc = append(sc, SymbolCount{symbol: sym, count: cnt})
	}
	sort.Slice(sc, func(i, j int) bool {
		if sc[i].count == sc[j].count {
			return sc[i].symbol < sc[j].symbol
		}
		return sc[i].count < sc[j].count
	})
	return sc
}

// textToCodeBits emits Huffman codes bit by bit
func textToCodeBits(text string, tree *HuffTree) chan uint8 {
	table := GetCompressTable(tree)
	ch := make(chan uint8)
	go func() {
		codes := []Code{}
		for sym := range tokenize(text) {
			code := table[sym]
			codes = append(codes, code)
		}
		// EOF
		codes = append(codes, table[""])
		for _, code := range codes {
			for _, bit := range code {
				if bit == '1' {
					ch <- 1
				} else {
					ch <- 0
				}
			}
		}
		close(ch)
	}()
	return ch
}

func merge(a, b map[Symbol]Code) map[Symbol]Code {
	for k, v := range b {
		a[k] = v
	}
	return a
}

// emitBytes packs given bits to bytes
func emitBytes(bits <-chan uint8) chan byte {
	ch := make(chan byte)
	go func() {
		var (
			b  byte
			bi uint8
		)
		for bit := range bits {
			if bi == 8 {
				ch <- b
				bi = 0
				b = 0
			}
			if bit == 1 {
				b |= (1 << bi)
			}
			bi++
		}
		if bi > 0 {
			ch <- b
		}
		close(ch)
	}()
	return ch
}

// emitBytes unpacks given bytes to bits
func emitBits(buf *bytes.Buffer) chan uint8 {
	ch := make(chan uint8)
	go func() {
		for {
			b, err := buf.ReadByte()
			if err != nil {
				break
			}
			for i := uint(0); i < 8; i++ {
				if b&(1<<i) > 0 {
					ch <- 1
				} else {
					ch <- 0
				}
			}
		}
		close(ch)
	}()
	return ch
}
