package cmd

import (
	"bufio"
	"log"
	"os"
)

func getScanner() *bufio.Scanner {
	if inputFile == "" {
		return bufio.NewScanner(os.Stdin)
	}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func getWriter() *bufio.Writer {
	if outputFile == "" {
		return bufio.NewWriter(os.Stdout)
	}
	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewWriter(file)
}
