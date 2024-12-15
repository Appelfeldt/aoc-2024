package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	solver "github.com/appelfeldt/aoc-2024/04-ceres-search/internal/solver"
	"github.com/spf13/cobra"
)

var BuildVersion string

var rootCmd = &cobra.Command{
	Use:     "aoc-2024-04 <filepath>",
	Version: BuildVersion,
	Short:   "aoc-2024-04 - Counts word occurances",
	Long:    "aoc-2024-04 counts occurances of 'XMAS' and 'X-MAS'",
	Args:    cobra.MaximumNArgs(1),
	Run:     command,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred '%s'", err)
		os.Exit(1)
	}
}

func hasPipedInput() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func command(cmd *cobra.Command, args []string) {
	var input string
	var inputReader io.Reader = cmd.InOrStdin()

	if hasPipedInput() {
		var buffer bytes.Buffer
		io.Copy(&buffer, inputReader)
		input = buffer.String()
	} else {
		if len(args) > 0 && args[0] != "-" {
			file, err := os.Open(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "failure opening file: %v", err)
				os.Exit(1)
			}
			inputReader = file
		}

		var buffer bytes.Buffer
		io.Copy(&buffer, inputReader)

		input = buffer.String()
	}

	sum1, sum2 := solver.Calculate(input)
	fmt.Printf("XMAS count:\n%v\n", sum1)
	fmt.Printf("X-MAS count:\n%v\n", sum2)
}