package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	solver "github.com/appelfeldt/aoc-2024/02-historian-hysteria/internal/solver"
	"github.com/spf13/cobra"
)

var BuildVersion string

var rootCmd = &cobra.Command{
	Use:     "aoc-2024-02 <filepath>",
	Version: BuildVersion,
	Short:   "aoc-2024-02 - Calculates safe reports.",
	Long:    "aoc-2024-02 calculates count of safe reports, with and without problem dampener.",
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

	safe_reports, safe_reports_dampener := solver.Calculate(input)
	fmt.Printf("Safe Reports:\n%v\n", safe_reports)
	fmt.Printf("Safe Reports with dampener:\n%v\n", safe_reports_dampener)
}
