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

const num = "[0-9.]+"
const sym = "[\\+\\-\\*\\/\\^]"
const brackets = "\\(.*\\)"

var symbolItem = regexp.MustCompile(fmt.Sprintf("^%s$", sym))
var anyItem = regexp.MustCompile(fmt.Sprintf("%s|%s|%s", brackets, num, sym))
var symbolNeedsBracket = regexp.MustCompile(`[\*\/\^]`)
var simpleNumber = regexp.MustCompile("^[0-9.]+$")

func PrefixToInfix(input string) (string, error) {
	tempResult := anyItem.FindAllString(input, -1)

	if len(tempResult) < 3 {
		return "", fmt.Errorf("insufficient input items: %d", len(tempResult))
	}

	for len(tempResult) >= 3 {
		savedItems := []string{}

		for len(tempResult) >= 4 && !symbolItem.MatchString(tempResult[len(tempResult)-1-2]) {
			var lastElem string
			tempResult, lastElem = slicePop(tempResult)
			savedItems = append(savedItems, lastElem)
		}

		curItem1, curItem2 := tempResult[len(tempResult)-1], tempResult[len(tempResult)-2]
		tempResult = tempResult[:len(tempResult)-2]
		curItems := []string{curItem1, curItem2}

		var curSymbols string
		tempResult, curSymbols = slicePop(tempResult)

		if symbolNeedsBracket.MatchString(curSymbols) {
			for i := 0; i < len(curItems); i++ {
				if !simpleNumber.MatchString(curItems[i]) {
					curItems[i] = fmt.Sprintf("(%s)", curItems[i])
				}
			}
		}

		var newItem = fmt.Sprintf("%s %s %s", curItems[1], curSymbols, curItems[0])
		tempResult = append(tempResult, newItem)
		tempResult = append(tempResult, reverseSlice(savedItems)...)
	}

	if len(tempResult) != 1 {
		return "", fmt.Errorf("wrong input: %s", input)
	}

	return tempResult[0], nil
}
