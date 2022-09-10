package dajareGo

import (
	"unicode/utf8"
)

// Calculate count of similar word with edit distance.
func fuzzyCount(s Syllables, substring Syllables, threshold int) int {
	// initialize
	table := make([][]int, len(substring)+1)
	for index := range table {
		table[index] = make([]int, len(s)+1)
	}
	for i := 1; i <= len(substring); i += 1 {
		table[i][0] = i
	}
	for j := 1; j <= len(s); j += 1 {
		table[0][j] = 0
	}

	// calculate
	for i := 1; i <= len(substring); i += 1 {
		for j := 1; j <= len(s); j += 1 {
			var replaceCost int
			if substring[i-1].removeDot() == s[j-1].removeDot() {
				replaceCost = 0
			} else {
				replaceCost = editDistance(string(substring[i-1].removeDot()), string(s[j-1].removeDot()))
			}
			replace := table[i-1][j-1] + replaceCost
			delete := table[i-1][j] + 2
			insert := table[i][j-1] + 2
			costs := [...]int{delete, insert, replace}
			minimum := -1
			for _, c := range costs {
				if (minimum >= 0 && minimum > c) || minimum < 0 {
					minimum = c
				}
			}
			table[i][j] = minimum
		}
	}

	// count
	count := 0
	for _, value := range table[len(substring)] {
		if value <= threshold && value < substring.length() {
			count += 1
		}
	}
	return count
}

// Calculate edit distance.
func editDistance(s string, t string) int {
	// initialize
	table := make([][]int, utf8.RuneCountInString(s)+1)
	for index := range table {
		table[index] = make([]int, utf8.RuneCountInString(t)+1)
	}
	for i := 1; i <= utf8.RuneCountInString(s); i += 1 {
		table[i][0] = i
	}
	for j := 1; j <= utf8.RuneCountInString(t); j += 1 {
		table[0][j] = j
	}

	// calculate
	for i := 1; i <= utf8.RuneCountInString(s); i += 1 {
		s_i := []rune(s)[i-1]
		for j := 1; j <= utf8.RuneCountInString(t); j += 1 {
			var replaceCost int
			substring_j := []rune(t)[j-1]
			if s_i == substring_j {
				replaceCost = 0
			} else {
				replaceCost = 1
			}
			replace := table[i-1][j-1] + replaceCost
			delete := table[i-1][j] + 1
			insert := table[i][j-1] + 1
			costs := [...]int{delete, insert, replace}
			minimum := -1
			for _, c := range costs {
				if (minimum >= 0 && minimum > c) || minimum < 0 {
					minimum = c
				}
			}
			table[i][j] = minimum
		}
	}
	return table[utf8.RuneCountInString(s)][utf8.RuneCountInString(t)]
}
