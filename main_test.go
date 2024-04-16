package main

import "testing"

func TestPalindrome(t *testing.T) {
	s1 := StrPalindrome("bab")

	if s1.IsPalindrom() != true {
		t.Fail()
	}

	s2 := StrPalindrome("Ajeje")
	if s2.IsPalindrom() == true {
		t.Fail()
	}
}
