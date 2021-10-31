package lab2

import (
	"fmt"
	"regexp"
)

const num = "[0-9.]+"
const sym = "[\\+\\-\\*\\/]"
const brackets = "\\(.*\\)"

var symArg = regexp.MustCompile(fmt.Sprintf("^%s$", sym))
var numArg = regexp.MustCompile(fmt.Sprintf("%s|%s", brackets, num))
var anyArg = regexp.MustCompile(fmt.Sprintf("%s|%s|%s", brackets, num, sym))
var symNeedsBracket = regexp.MustCompile("[\\*\\/]")
var simpleNumber = regexp.MustCompile("^[0-9.]+$")

func slicePop(arr []string) ([]string, string) {
	l := len(arr)
	item := arr[l-1]
	return arr[:l-1], item
}

func PrefixToInfix(input string) (string, error) {
	res := []string{}
	res = anyArg.FindAllString(input, -1)

	for len(res) >= 3 {
		savedArgs := []string{}

		for (len(res) >= 4) && len(symArg.FindAllString(res[len(res)-1-2], -1)) == 0 {
			var lastElem string
			res, lastElem = slicePop(res)
			savedArgs = append(savedArgs, lastElem)
		}

		var curArg1 string
		res, curArg1 = slicePop(res)
		var curArg2 string
		res, curArg2 = slicePop(res)

		curArgs := []string{curArg1, curArg2}
		var curSym string
		res, curSym = slicePop(res)

		if len(symNeedsBracket.FindAllString(curSym, -1)) > 0 {
			for i := 0; i < len(curArgs); i++ {
				if len(simpleNumber.FindAllString(curArgs[i], -1)) == 0 {
					curArgs[i] = fmt.Sprintf("(%s)", curArgs[i])
				}
			}
		}

		var newArg = fmt.Sprintf("%s %s %s", curArgs[1], curSym, curArgs[0])
		res = append(res, newArg)

		savedArgsReversed := []string{}

		for i := len(savedArgs) - 1; i >= 0; i-- {
			savedArgsReversed = append(savedArgsReversed, savedArgs[i])
		}
		res = append(res, savedArgsReversed...)
	}

	return res[0], nil
}
