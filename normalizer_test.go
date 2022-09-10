package dajareGo

import "testing"

type testcase struct {
	input  string
	expect string
}

func TestPreNormalizer(t *testing.T) {
	testcases := [...]testcase{
		// replace characters that like two or more occurrences of Japanese Prolonged Sound Mark "ー"
		{"ずーーーーーーしーーーーーーほっきーーー－－", "ずーしーほっきー"},

		// Delete heading Hankaku Space and tailing Hankaku Space
		{"  ほげほげ", "ほげほげ"},
		{"ほげほげ  ", "ほげほげ"},

		// Delete Hankaku Space among Hiragana, Katakana, and Kanji
		{"情報 表現  入門 の 教科書 を 買い ました", "情報表現入門の教科書を買いました"},
		{"Introduction to Information Expression", "Introduction to Information Expression"},

		// Delete Hankaku Space between Hiragana, Katakana, and Kanji, and Alphabet
		{"アルゴリズム Go", "アルゴリズムGo"},
		{"Algorithm Go", "Algorithm Go"},

		// general test
		{"公立 はこだて　未来  大学 Future　Ｕｎｉｖｅｒｓｉｔｙ Hakodate",
			"公立はこだて未来大学Future University Hakodate"},
		{"公立 はこだて　未来  大学-　Future*　Ｕｎｉｖｅｒｓｉｔｙ+　Hakodate",
			"公立はこだて未来大学-Future*University+Hakodate"},
	}
	for index, testcase := range testcases {
		if s := preNormalizer(testcase.input); s != testcase.expect {
			t.Errorf("failed: %d expects\"%s\" but output was \"%s\"", index, testcase.expect, s)
		}
	}
}
