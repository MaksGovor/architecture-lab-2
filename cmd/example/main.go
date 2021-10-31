package main

import (
	lab2 "github.com/Scopics/architecture-lab-2"

	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file with expression to compute")
	outputFile      = flag.String("o", "", "Output file for result")
)

func main() {
	flag.Parse()

	const LOG_ERR_PREFIX = "\x1b[1;31mE: "
	errLogger := log.New(os.Stderr, LOG_ERR_PREFIX, log.LstdFlags)

	var (
		fromReader   io.Reader
		outputWriter io.Writer
		err          error
	)

	if *inputExpression != "" {
		fromReader = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		fromReader, err = os.Open(*inputFile)
		if err != nil {
			errLogger.Printf("Can't open file: %s. msg: %s", *inputFile, err)
		}
	} else {
		errLogger.Print("Expression not found")
	}

	if *outputFile != "" {
		outputWriter, err = os.Create(*outputFile)
		if err != nil {
			errLogger.Printf("Can't create file: %s. msg: %s", *outputFile, err)
		}
	}

	handler := &lab2.ComputeHandler{
		Input:  fromReader,
		Output: outputWriter,
	}

	err = handler.Compute()
	if err != nil {
		errLogger.Print(err)
	}
}
