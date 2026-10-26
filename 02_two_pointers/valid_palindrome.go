// Pattern: Two Pointers — shrink from both ends toward the middle.
//
// Problem (LeetCode 125, Valid Palindrome): after lower-casing and removing
// non-alphanumeric characters, does the string read the same both ways?
//
// Idea: one pointer at each end; skip non-alphanumerics, compare lower-cased.
// Time O(n), space O(1).
package twopointers

import "unicode"

func IsPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		for lo < hi && !isAlnum(rune(s[lo])) {
			lo++
		}
		for lo < hi && !isAlnum(rune(s[hi])) {
			hi--
		}
		if unicode.ToLower(rune(s[lo])) != unicode.ToLower(rune(s[hi])) {
			return false
		}
		lo++
		hi--
	}
	return true
}

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
