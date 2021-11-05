package lab2

import (
	"fmt"
	"regexp"
)

func slicePop(arr []string) ([]string, string) {
	l := len(arr)
	item := arr[l-1]
	return arr[:l-1], item
}

func reverseSlice(slice []string) []string {
	res := []string{}
	for i := len(slice) - 1; i >= 0; i-- {
		res = append(res, slice[i])
	}
	return res
}

const operators = `\+\-\*\/\^`
const num = "[0-9]+"
const brackets = `\(.*\)`

var symbol = fmt.Sprintf(`[%s]`, operators)
var symbolItem = regexp.MustCompile(fmt.Sprintf("^%s$", symbol))
var anyItem = regexp.MustCompile(fmt.Sprintf("%s|%s|%s", brackets, num, symbol))
var numItem = regexp.MustCompile(fmt.Sprintf("%s|%s", brackets, num))
var symbolNeedsBracket = regexp.MustCompile(`[\*\/\^]`)
var simpleNumber = regexp.MustCompile("^[0-9]+$")
var wrongChar = regexp.MustCompile(fmt.Sprintf(`[^\s0-9%s]`, operators))

func PrefixToInfix(input string) (string, error) {
	tempResult := anyItem.FindAllString(input, -1)

	if wrongChar.MatchString(input) {
		return "", fmt.Errorf("wrong symbols in input: %s", input)
	}

	if len(tempResult) == 0 {
		return "", fmt.Errorf("empty input")
	} else if len(tempResult) == 1 && !simpleNumber.MatchString(tempResult[0]) {
		return "", fmt.Errorf("wrong input: %s", tempResult[0])
	}

	for len(tempResult) >= 3 {
		savedItems := []string{}

		for len(tempResult) >= 4 && !symbolItem.MatchString(tempResult[len(tempResult)-1-2]) {
			var lastElem string
			tempResult, lastElem = slicePop(tempResult)
			savedItems = append(savedItems, lastElem)
		}

		curItems := []string{tempResult[len(tempResult)-1], tempResult[len(tempResult)-2]}
		tempResult = tempResult[:len(tempResult)-2]

		var curSymbol string
		tempResult, curSymbol = slicePop(tempResult)

		if symbolNeedsBracket.MatchString(curSymbol) {
			for i := 0; i < len(curItems); i++ {
				if !simpleNumber.MatchString(curItems[i]) {
					curItems[i] = fmt.Sprintf("(%s)", curItems[i])
				}
			}
		}

		for _, item := range curItems {
			if !numItem.MatchString(item) {
				return "", fmt.Errorf("operator, argument mismatch: %s", input)
			}
		}

		if !symbolItem.MatchString(curSymbol) {
			return "", fmt.Errorf("operator, argument mismatch: %s", input)
		}

		var newItem = fmt.Sprintf("%s %s %s", curItems[1], curSymbol, curItems[0])
		tempResult = append(append(tempResult, newItem), reverseSlice(savedItems)...)
	}

	if len(tempResult) != 1 {
		return "", fmt.Errorf("wrong argument amount: %s", input)
	}

	return tempResult[0], nil
}
