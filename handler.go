package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	writter := bufio.NewWriter(ch.Output)

	for scanner.Scan() {
		exp := scanner.Text()
		res, err := PrefixToInfix(exp)
		if err != nil {
			return err
		}

		if _, err := writter.WriteString(res + "\n"); err != nil {
			return err
		}

		if err := writter.Flush(); err != nil {
			return err
		}
	}

	return nil
}
