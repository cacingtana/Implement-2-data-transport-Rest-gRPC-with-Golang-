package main

import (
	"fmt"
	"reflect"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func findFirstStringInBracketRefactor(s string) string {
	i := strings.Index(s, "(")
	fmt.Println(reflect.TypeOf(i))
	if i < 0 {
		return ""
	}
	i++
	j := strings.Index(s[i:], ")")
	fmt.Println(reflect.TypeOf(j))
	if j < 0 {
		return ""
	}
	return s[i : i+j]
}

func main() {

	fmt.Println(findFirstStringInBracket("(stevrasilvanus)"))
	fmt.Println(findFirstStringInBracketRefactor("(stevrasilvanus)"))

}
