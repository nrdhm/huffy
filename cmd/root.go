package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	fnameIn  string
	fnameOut string
	fileIn   *os.File = os.Stdin
	fileOut  *os.File = os.Stdout
	err      error
)

var rootCmd = &cobra.Command{
	Use:   "huffy",
	Short: "huffman compressor",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if fnameIn != "" && fnameIn != "-" {
			fileIn, err = os.Open(fnameIn)
			if err != nil {
				panic(err)
			}
		}
		if fnameOut != "" && fnameOut != "-" {
			fileOut, err = os.Create(fnameOut)
			if err != nil {
				panic(err)
			}
		}
	},
}

// Execute runs the program
func Execute() {
	rootCmd.Execute()
	rootCmd.PersistentFlags().StringVar(&fnameIn, "in", "-", "file to read from")
	rootCmd.PersistentFlags().StringVar(&fnameOut, "out", "-", "file to write to")
}
