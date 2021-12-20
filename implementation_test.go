package lab2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
		err      string
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
			err:      "wrong argument(s) in operation: 2 + +",
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
			err:      "empty input",
		},
		{
			name:     "Not Valid Symbols",
			input:    "$ 3 & a b . : ;",
			expected: "",
			hasError: true,
			err:      "wrong input: $ 3 & a b . : ;",
		},
		{
			name:     "Symbols",
			input:    "+ b c",
			expected: "",
			hasError: true,
			err:      "wrong input: + b c",
		},
		{
			name:     "One symbol",
			input:    "b",
			expected: "",
			hasError: true,
			err:      "wrong input: b",
		},
		{
			name:     "Spacing",
			input:    " ",
			expected: "",
			hasError: true,
			err:      "empty input",
		},
		{
			name:     "One digit",
			input:    "3",
			expected: "3",
			hasError: false,
		},
		{
			name:     "One operator",
			input:    "+",
			expected: "",
			hasError: true,
			err:      "wrong input: +",
		},
		{
			name:     "Only operators",
			input:    "+ - / * ^ *",
			expected: "",
			hasError: true,
			err:      "wrong argument(s) in operation: / - (^) * (*)",
		},
		{
			name:     "Only digits",
			input:    "1 2 3 4",
			expected: "",
			hasError: true,
			err:      "wrong operators in operation: 2 1 3",
		},
		{
			name:     "wrong argument amount",
			input:    "3 - 2 2",
			expected: "",
			hasError: true,
			err:      "wrong argument amount: 3 - 2 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := PrefixToInfix(test.input)
			assert.Equal(test.expected, actual, test.name)

			if test.hasError {
				assert.NotNil(err, test.name)
				assert.EqualError(err, test.err)
			} else {
				assert.Nil(err, test.name)
			}
		})
	}

}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
	fmt.Println(res)

	// Output:
	// (15 / (7 - 1 + 1)) * 3 - 2 + 1 + 1
}

var res string

func BenchmarkPrefixToInfix(b *testing.B) {
	str := "- * / 15 - 7 + 1 1 3 + 2 + 1 + 9 "
	for i := 1; i <= 100; i++ {
		testStr := strings.Repeat(str, i) + "1"
		b.Run(fmt.Sprintf("%d-operators", i*6), func(b *testing.B) {
			res, _ = PrefixToInfix(testStr)
		})
	}
}
