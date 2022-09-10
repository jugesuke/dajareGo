package dajareGo

import (
	"strings"

	ipaneologd "github.com/ikawaha/kagome-dict-ipa-neologd"
	"github.com/ikawaha/kagome-dict/dict"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

type (
	// Result has a result of IsDajare function.
	Result struct {
		// If it is Dajare, this field is True, else False.
		IsDajare bool

		// This field shows an index of all word which has a similar reading but a different meaning
		DajareWordIndex []int

		// a pair of sentence and its Syllables
		Sentence Token

		// Tokens of given sentence
		Tokens []Token
	}

	// Token is a pair of Surface and its Syllables
	Token struct {
		Surface   string
		Syllables Syllables
	}
)

var t *tokenizer.Tokenizer

// Set dictionary. You must do this first.
func Init() error {
	if _t, err := tokenizer.New(ipaneologd.Dict(), tokenizer.OmitBosEos()); err != nil {
		return err
	} else {
		t = _t
		return nil
	}
}

// Set dictionary you like. You can use a Kagome Dictionary.
// https://github.com/ikawaha/kagome#dictionaries
func SetCustomDictionary(dict *dict.Dict) error {
	if _t, err := tokenizer.New(dict, tokenizer.OmitBosEos()); err != nil {
		return err
	} else {
		t = _t
		IsDajare("")
		return nil
	}
}

// IsDajare checks if a sentence is Dajare.
func IsDajare(s string) Result {
	// normalize before Morphological analysis
	s = preNormalizer(s)
	// Morphological analysis
	tokens := t.Tokenize(s)
	var r Result
	var rTokens []Token
	for _, token := range tokens {
		surface := token.Surface
		r.Sentence.Surface += surface

		pron, ok := token.Pronunciation()
		if ok {
			syllables := NewSyllables(pron)
			newToken := Token{surface, syllables}
			rTokens = append(rTokens, newToken)
			r.Sentence.Syllables = append(r.Sentence.Syllables, syllables...)
		} else {
			suspectedPron := ""
			flag := true
			for _, value := range surface {
				if 0x3041 <= int(value) && int(value) <= 0x3094 {
					// if value is Hiragana
					suspectedPron += string(value + 0x0060)
				} else if 0x30A1 <= int(value) && int(value) <= 0x30FF {
					// if value is Katakana
					suspectedPron += string(value)
				} else {
					flag = false
				}
			}
			if flag {
				syllables := NewSyllables(suspectedPron)
				newToken := Token{surface, syllables}
				rTokens = append(rTokens, newToken)
				r.Sentence.Syllables = append(r.Sentence.Syllables, syllables...)
			}
		}
	}
	r.Tokens = rTokens

	// Analyze if the string is Dajare
	for i, token := range r.Tokens {
		// prepare for analyzing
		var target Token
		if token.pronLen() < 3 {
			// If token is too short...
			continue
		} else if token.pronLen() == 2 {
			if i == len(r.Tokens)-1 || i == 0 {
				continue
			}
			if r.Tokens[i-1].pronLen() >= 2 {
				target = token.union(r.Tokens[i-1])
			}
		} else {
			target = token
		}

		// analyze
		if r.analyze(target) {
			if r.IsDajare {
				r.DajareWordIndex = append(r.DajareWordIndex, i)
			} else {
				r.IsDajare = true
				r.DajareWordIndex = []int{i}
			}
		}
	}
	return r
}

// Check each token if a sentence has a similar reading but a different meaning.
func (r *Result) analyze(t Token) bool {
	surfaceCount := strings.Count(r.Sentence.Surface, t.Surface)
	pronCount := fuzzyCount(r.Sentence.Syllables, t.Syllables, 1)
	return surfaceCount < pronCount && 2 <= pronCount
}

// get pron length
func (t *Token) pronLen() int {
	var len int
	for _, syllable := range t.Syllables {
		len += syllable.length() - 1
	}
	return len
}

// union 2 tokens
func (t Token) union(new Token) Token {
	return Token{t.Surface + new.Surface, append(t.Syllables, new.Syllables...)}
}
