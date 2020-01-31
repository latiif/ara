package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/buger/goterm"
	ara "github.com/latiif/ara/goarabic"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ttywidth := int(goterm.Width())

	for scanner.Scan() {

		rawText := scanner.Text()
		table := ara.GenerateTashkeelTable(rawText)

		fmt.Println(
			ara.MakeRTL(ttywidth,
				ara.ApplyTashkeel(table,
					ara.ReversePreservingNonArabic(
						ara.ToGlyph(
							ara.Smooth(
								ara.RemoveTashkeel(rawText)))))))

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
