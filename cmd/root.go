package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/buger/goterm"
	"github.com/latiif/ara/pkg/arabic"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ara",
	Short: "Render Arabic text in terminal",
	Long:  "Correctly displays Arabic text in terminals",
	Run: func(cmd *cobra.Command, args []string) {
		arartl()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func arartl() {
	scanner := bufio.NewScanner(os.Stdin)

	ttywidth := int(goterm.Width())

	for scanner.Scan() {

		rawText := scanner.Text()
		table := arabic.GenerateTashkeelTable(rawText)

		fmt.Println(
			arabic.MakeRTL(ttywidth,
				arabic.ApplyTashkeel(table,
					arabic.ReversePreservingNonArabic(
						arabic.ToGlyph(
							arabic.Smooth(
								arabic.RemoveTashkeel(rawText)))))))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
