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
			name:     "Two-Three Operands",
			input:    "+ 1 ^ 2 3",
			expected: "1 + 2 ^ 3",
			hasError: false,
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
