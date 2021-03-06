package cmd

import (
	"fmt"
	"log"

	"github.com/buger/goterm"
	"github.com/latiif/ara/pkg/arabic"
	"github.com/spf13/cobra"
)

var (
	rtlFlag    bool   = true
	inputFile  string = ""
	outputFile string = ""
)

var rootCmd = &cobra.Command{
	Use:   "ara",
	Short: "Render Arabic text in terminal",
	Long:  "Correctly displays Arabic text in terminals",
	Run: func(cmd *cobra.Command, args []string) {
		arartl()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func arartl() {

	scanner := getScanner()
	writer := getWriter()

	for scanner.Scan() {
		rawText := scanner.Text()
		table := arabic.GenerateTashkeelTable(rawText)

		fmt.Fprintln(writer,
			guardedRTL(
				arabic.ApplyTashkeel(table,
					arabic.ReversePreservingNonArabic(
						arabic.ToGlyph(
							arabic.Smooth(
								arabic.RemoveTashkeel(rawText)))))))

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&rtlFlag, "adjust-right", "a", false, "Adjust output text to be rtl (useful when in shell, less so if you want to pipe into a file)")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Apply command on the contents of the input file.")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Write output to specified file.")
}

func guardedRTL(str string) string {
	if rtlFlag {
		ttywidth := int(goterm.Width())
		if ttywidth <= 0 {
			log.Fatalf("tty width is (%d), unable to adjust right. (You might be piping into a file).", ttywidth)
		}
		return arabic.MakeRTL(ttywidth, str)
	}
	return str
}
