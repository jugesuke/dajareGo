package dajareGo

import "unicode/utf8"

type (
	Syllable  string
	Syllables []Syllable
)

func NewSyllables(new string) Syllables {
	var s Syllables
	new = postNormalizer(new)
	for _, char := range new {
		isIndSyllables := false
		indSyllables := [...]rune{'ン', 'ッ', 'ー', 'ャ', 'ュ', 'ョ', 'ァ', 'ィ', 'ゥ', 'ェ', 'ォ', 'ヮ'}
		for _, indSyllable := range indSyllables {
			if char == indSyllable {
				isIndSyllables = true
			}
		}
		isNonSyllables := false
		nonSyllables := [...]rune{'、', '。'}
		for _, nonSyllable := range nonSyllables {
			if char == nonSyllable {
				isNonSyllables = true
			}
		}
		if isIndSyllables && len(s) > 0 {
			s[len(s)-1] += katakana2Romaji(Syllable(char))
		} else if !isNonSyllables {
			s = append(s, katakana2Romaji(Syllable(char)))
		}
	}
	return s
}

func (s Syllables) length() int {
	return len([]Syllable(s))
}

func (s Syllable) length() int {
	return utf8.RuneCountInString(string(s))
}

func (s *Syllable) removeDot() Syllable {
	var new Syllable
	for _, char := range *s {
		if char != '.' {
			new = Syllable(string(new) + string(char))
		}
	}
	return new
}

func (s1 Syllables) equals(s2 Syllables) bool {
	if s1.length() != s2.length() {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
