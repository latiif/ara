package cmd

import (
	"fmt"

	"github.com/latiif/ara/pkg/undot"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(normCmd)
}

var normCmd = &cobra.Command{
	Use:   "undot",
	Short: "Removes the dots of Arabic letters",
	Long:  `Displays text in Rasm removing dots from dotted characters and tashkeels`,
	Run:   runUndotCmd,
}

func runUndotCmd(cmd *cobra.Command, args []string) {

	scanner := getScanner()
	writer := getWriter()

	for scanner.Scan() {
		rawText := scanner.Text()
		fmt.Fprintln(writer, undot.Undot(rawText))
		writer.Flush()
	}
}
