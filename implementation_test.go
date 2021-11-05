package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := PrefixToInfix("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
	if assert.Nil(t, err) {
		assert.Equal(t, "(15 / (7 - 1 + 1)) * 3 - 2 + 1 + 1", res)
	}
}

func TestPrefixToPostfix2(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "Base",
			input:    "- * / 15 - 7 + 1 1 3 + 2 + 1 1",
			expected: "(15 / (7 - 1 + 1)) * 3 - 2 + 1 + 1",
			hasError: false,
		},
		{
			name:     "Base With Error",
			input:    "- * / 15 - 7 + 1 1 3 + 2 + 1 ",
			expected: "",
			hasError: true,
		},
		{
			name:     "Two Operands",
			input:    "^ 17 19",
			expected: "17 ^ 19",
			hasError: false,
		},
		{
			name:     "Three Operands",
			input:    "+ 1 ^ 2 3",
			expected: "1 + 2 ^ 3",
			hasError: false,
		},
		{
			name:     "Seven Operands",
			input:    "+ / - ^ 20 2 34 / 15 - 4 1 56",
			expected: "(20 ^ 2 - 34) / (15 / (4 - 1)) + 56",
			hasError: false,
		},
		{
			name:     "Ten Operands",
			input:    "+ 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 3",
			expected: "3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3",
			hasError: false,
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
			hasError: true,
		},
		{
			name:     "Not Valid Symbols",
			input:    "$ 3 & a b . : ;",
			expected: "",
			hasError: true,
		},
		{
			name:     "Symbols",
			input:    "+ b c",
			expected: "",
			hasError: true,
		},
		{
			name:     "one symbol",
			input:    "b",
			expected: "",
			hasError: true,
		},
		{
			name:     "spacing",
			input:    " ",
			expected: "",
			hasError: true,
		},
		{
			name:     "one digit",
			input:    "3",
			expected: "3",
			hasError: false,
		},
		{
			name:     "no digits 1",
			input:    "+",
			expected: "",
			hasError: true,
		},
		{
			name:     "no digits 2",
			input:    "*",
			expected: "",
			hasError: true,
		},
		{
			name:     "no digits 4",
			input:    "/",
			expected: "",
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := PrefixToInfix(test.input)
			assert.Equal(test.expected, actual, test.name)

			if test.hasError {
				assert.NotNil(err, test.name)
			} else {
				assert.Nil(err, test.name)
			}
		})
	}

}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
