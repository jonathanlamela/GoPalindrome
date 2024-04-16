package main

import (
	"fmt"
	"strings"
)

type StrPalindrome string

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func (parola StrPalindrome) IsPalindrom() bool {
	return strings.EqualFold(string(parola), reverse(string(parola)))
}

func main() {

	s1 := StrPalindrome("bab")
	fmt.Println(s1.IsPalindrom())

	s2 := StrPalindrome("Ajeje")
	fmt.Println(s2.IsPalindrom())
}
