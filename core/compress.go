package core

import (
	"bytes"
	"encoding/gob"
)

// Compress encodes text with Huffman codes
func Compress(text string) (string, error) {
	if len(text) == 0 {
		return "", nil
	}
	var out bytes.Buffer
	ps := countSymbols(text)
	tree := NewTree(ps)
	err := writeHeader(tree, &out)
	if err != nil {
		return "", err
	}
	bits := textToCodeBits(text, tree)
	for b := range emitBytes(bits) {
		out.WriteByte(b)
	}
	return out.String(), nil
}

// Decompress decodes Huffman codes to clear text
func Decompress(compressed string) (string, error) {
	if len(compressed) == 0 {
		return "", nil
	}
	in := bytes.NewBufferString(compressed)
	root, err := readHeader(in)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	bits := emitBits(in)
	tree := root
	for bit := range bits {
		if bit == 1 {
			tree = tree.One
		} else {
			tree = tree.Zero
		}
		if tree == nil || tree.EOF {
			break
		}
		if tree.Sym != "" {
			out.WriteString(string(tree.Sym))
			tree = root
		}
	}
	return out.String(), nil
}

func writeHeader(tree *HuffTree, buf *bytes.Buffer) error {
	enc := gob.NewEncoder(buf)
	err := enc.Encode(tree)
	if err != nil {
		return err
	}
	return nil
}

func readHeader(b *bytes.Buffer) (*HuffTree, error) {
	dec := gob.NewDecoder(b)
	var tree *HuffTree
	err := dec.Decode(&tree)
	if err != nil {
		return nil, err
	}
	return tree, nil
}
