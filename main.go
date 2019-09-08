package main

import (
	"ara/goarabic"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		fmt.Println(goarabic.Reverse(goarabic.ToGlyph(goarabic.Smooth(goarabic.RemoveTashkeel(scanner.Text())))))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
