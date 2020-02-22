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

var rtlFlag bool = true

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

	for scanner.Scan() {
		rawText := scanner.Text()
		table := arabic.GenerateTashkeelTable(rawText)

		fmt.Println(
			guardedRTL(
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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&rtlFlag, "adjust-right", "a", false, "Adjust output text to be rtl (useful when in shell, less so if you want to pipe into a file)")
}

func guardedRTL(str string) string {
	if rtlFlag {
		ttywidth := int(goterm.Width())
		return arabic.MakeRTL(ttywidth, str)
	}
	return str
}
