package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type writterMock struct {
	buffer    *bytes.Buffer
	writeCall bool
}

func (wm *writterMock) Write(bytes []byte) (int, error) {
	wm.writeCall = true
	return wm.buffer.Write(bytes)
}

func (wm *writterMock) String() string {
	return wm.buffer.String()
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
			expected: "(15 / (7 - 1 + 1)) * 3 - 2 + 1 + 1\n",
			hasError: false,
		},
		{
			name:     "Base With Error",
			input:    "- * / 15 - 7 + 1 1 3 + 2 + 1 ",
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
			inputMock := strings.NewReader(test.input)
			outputMock := &writterMock{
				buffer: new(bytes.Buffer),
			}

			handler := ComputeHandler{
				Input:  inputMock,
				Output: outputMock,
			}

			err := handler.Compute()

			if test.hasError {
				assert.NotNil(err)
			} else {
				assert.True(outputMock.writeCall)
				assert.Equal(test.expected, outputMock.String(), test.name)
			}
		})
	}
}
