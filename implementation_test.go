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

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
