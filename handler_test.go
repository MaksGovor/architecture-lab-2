package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockComputeHandler struct {
	Input, Output string
	Error         error
}

func (m *MockComputeHandler) Compute() {
	m.Output, m.Error = PrefixToInfix(m.Input)
}

func TestCompute(t *testing.T) {
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testMock := &MockComputeHandler{Input: test.input}
			testMock.Compute()
			assert.Equal(test.expected, testMock.Output, test.name)

			if test.hasError {
				assert.NotNil(testMock.Error, test.name)
			} else {
				assert.Nil(testMock.Error, test.name)
			}
		})
	}

}
