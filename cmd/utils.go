package cmd

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/nrdhm/huffy/core"
	"github.com/spf13/cobra"
)

var utilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "some utilities",
}

var printTreeCmd = &cobra.Command{
	Use:   "print-tree",
	Short: "print huffman tree generated for the text",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(fileIn)
		if err != nil {
			panic(err)
		}
		count := core.CountSymbols(string(data))
		tree := core.NewTree(count)
		printTree(tree, "")
		fmt.Println()
		fmt.Println(tree)
	},
}

var depth int

func printTree(tree *core.HuffTree, start string) {
	if tree == nil {
		return
	}
	depth++
	fmt.Printf("%s___%5q", start, tree.GetSym())
	printTree(tree.Zero, "")
	fmt.Printf("\n%s", strings.Repeat(" ", 8*depth-2))
	printTree(tree.One, " |")
	depth--
}

type symCode struct {
	sym  core.Symbol
	code core.Code
}

var printTableCmd = &cobra.Command{
	Use:   "print-table",
	Short: "print huffman codes table for the text",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(fileIn)
		if err != nil {
			panic(err)
		}
		count := core.CountSymbols(string(data))
		tree := core.NewTree(count)
		table := core.GetCompressTable(tree)
		pairs := []symCode{}
		for sym, code := range table {
			pairs = append(pairs, symCode{sym: sym, code: code})
		}
		sort.Slice(pairs, func(i, j int) bool {
			return len(pairs[i].code) < len(pairs[j].code)
		})
		for _, pair := range pairs {
			fmt.Printf("%v %q\n", pair.code, pair.sym)
		}
	},
}

func init() {
	utilsCmd.AddCommand(printTreeCmd)
	utilsCmd.AddCommand(printTableCmd)
	rootCmd.AddCommand(utilsCmd)
}
