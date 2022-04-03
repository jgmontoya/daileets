package letter_case_permutation

import (
	"strings"
	"unicode"
)

func letterCasePermutation(s string) []string {
	if len(s) == 1 {
		if isDigit(s) {
			return []string{s}
		}
		return []string{strings.ToLower(s), strings.ToUpper(s)}
	}

	permutations := []string{}
	subpermutations := letterCasePermutation(s[1:])
	first_letter_permutations := letterCasePermutation(s[:1])
	for _, subpermutation := range subpermutations {
		for _, first_letter_permutation := range first_letter_permutations {
			permutation := copyStringStartingWith(rune(first_letter_permutation[0]), subpermutation)
			permutations = append(permutations, permutation)
		}
	}

	return permutations
}

func isDigit(character string) bool {
	return !unicode.IsLetter(rune(character[0]))
}

func copyStringStartingWith(initial_char rune, s string) string {
	permutation_bytearray := make([]byte, len(s)+1)
	permutation_bytearray[0] = byte(initial_char)
	for index, string_rune := range s {
		permutation_bytearray[index+1] = byte(string_rune)
	}
	return string(permutation_bytearray)
}
