package cmd

import (
	"bufio"
	"fmt"
	"os"

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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		rawText := scanner.Text()
		fmt.Println(undot.Undot(rawText))
	}
}
