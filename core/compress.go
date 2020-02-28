package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Compress(text string) (string, error) {
	var (
		out bytes.Buffer
	)
	ps := CalcOccurrences(text)
	tree := NewTree(ps)
	table := GetCompressTable(tree)
	header, err := makeHeader(table)
	if err != nil {
		return "", err
	}
	out.WriteString(string(header))
	bits := textToCodeBits(text, table)
	for b := range emitBytes(bits) {
		out.WriteByte(b)
	}
	return out.String(), nil
}

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

func Decompress(compressed string) (string, error) {
	buf := bytes.NewBufferString(compressed)
	table, err := readHeader(buf)
	if err != nil {
		return "", err
	}
	root := FromCompressTable(table)
	tree := root
	var out bytes.Buffer
	bits := emitBits(buf)
	for bit := range bits {
		if bit == 1 {
			tree = tree.oneBranch
		} else if tree.zeroBranch != nil { // ignore trailing zeroes
			tree = tree.zeroBranch
		}
		if tree == nil || tree.eof {
			break
		}
		if tree.symbol != "" {
			out.WriteString(string(tree.symbol))
			tree = root
		}

	}
	return out.String(), nil
}

func makeHeader(table map[Symbol]Code) ([]byte, error) {
	var b bytes.Buffer
	_, err := b.WriteString("HUF|")
	if err != nil {
		return nil, err
	}
	enc := gob.NewEncoder(&b)
	err = enc.Encode(table)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func readHeader(b *bytes.Buffer) (map[Symbol]Code, error) {
	magic, err := b.ReadString('|')
	if err != nil {
		return nil, err
	}
	if magic != "HUF|" {
		return nil, fmt.Errorf("invalid magic: %v", magic)
	}
	dec := gob.NewDecoder(b)
	var m map[Symbol]Code
	err = dec.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
