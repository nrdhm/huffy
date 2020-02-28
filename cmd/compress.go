package cmd

import (
	"io/ioutil"

	"github.com/nrdhm/huff/core"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress data from stdin to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(fileIn)
		if err != nil {
			panic(err)
		}
		huf, err := core.Compress(string(data))
		if err != nil {
			panic(err)
		}
		fileOut.WriteString(huf)
	},
}

var deCompressCmd = &cobra.Command{
	Use:   "decompress",
	Short: "decompress data from stdin to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(fileIn)
		if err != nil {
			panic(err)
		}
		text, err := core.Decompress(string(data))
		if err != nil {
			panic(err)
		}
		fileOut.WriteString(text)
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	rootCmd.AddCommand(deCompressCmd)
}
