package dajareGo

import (
	"regexp"
	"strings"
)

// Japanese Normalizer (referred to neologdn)
func preNormalizer(s string) string {
	reps := [...]rep{
		// replace Zenkaku Alphabet and numbers
		{"０", "0"},
		{"１", "1"},
		{"２", "2"},
		{"３", "3"},
		{"４", "4"},
		{"５", "5"},
		{"６", "6"},
		{"７", "7"},
		{"８", "8"},
		{"９", "9"},
		{"ａ", "a"},
		{"ｂ", "b"},
		{"ｃ", "c"},
		{"ｄ", "d"},
		{"ｅ", "e"},
		{"ｆ", "f"},
		{"ｇ", "g"},
		{"ｈ", "h"},
		{"ｉ", "i"},
		{"ｊ", "j"},
		{"ｋ", "k"},
		{"ｌ", "l"},
		{"ｍ", "m"},
		{"ｎ", "n"},
		{"ｏ", "o"},
		{"ｐ", "p"},
		{"ｑ", "q"},
		{"ｒ", "r"},
		{"ｓ", "s"},
		{"ｔ", "t"},
		{"ｕ", "u"},
		{"ｖ", "v"},
		{"ｗ", "w"},
		{"ｘ", "x"},
		{"ｙ", "y"},
		{"ｚ", "z"},
		{"Ａ", "A"},
		{"Ｂ", "B"},
		{"Ｃ", "C"},
		{"Ｄ", "D"},
		{"Ｅ", "E"},
		{"Ｆ", "F"},
		{"Ｇ", "G"},
		{"Ｈ", "H"},
		{"Ｉ", "I"},
		{"Ｊ", "J"},
		{"Ｋ", "K"},
		{"Ｌ", "L"},
		{"Ｍ", "M"},
		{"Ｎ", "N"},
		{"Ｏ", "O"},
		{"Ｐ", "P"},
		{"Ｑ", "Q"},
		{"Ｒ", "R"},
		{"Ｓ", "S"},
		{"Ｔ", "T"},
		{"Ｕ", "U"},
		{"Ｖ", "V"},
		{"Ｗ", "W"},
		{"Ｘ", "X"},
		{"Ｙ", "Y"},
		{"Ｚ", "Z"},

		// replace Zenkaku Katakana
		{"ｱ", "ア"},
		{"ｲ", "イ"},
		{"ｳ", "ウ"},
		{"ｴ", "エ"},
		{"ｵ", "オ"},
		{"ｶ", "カ"},
		{"ｷ", "キ"},
		{"ｸ", "ク"},
		{"ｹ", "ケ"},
		{"ｺ", "コ"},
		{"ｻ", "サ"},
		{"ｼ", "シ"},
		{"ｽ", "ス"},
		{"ｾ", "セ"},
		{"ｿ", "ソ"},
		{"ﾀ", "タ"},
		{"ﾁ", "チ"},
		{"ﾂ", "ツ"},
		{"ﾃ", "テ"},
		{"ﾄ", "ト"},
		{"ﾅ", "ナ"},
		{"ﾆ", "ニ"},
		{"ﾇ", "ヌ"},
		{"ﾈ", "ネ"},
		{"ﾉ", "ノ"},
		{"ﾊ", "ハ"},
		{"ﾋ", "ヒ"},
		{"ﾌ", "フ"},
		{"ﾍ", "ヘ"},
		{"ﾎ", "ホ"},
		{"ﾏ", "マ"},
		{"ﾐ", "ミ"},
		{"ﾑ", "ム"},
		{"ﾒ", "メ"},
		{"ﾓ", "モ"},
		{"ﾔ", "ヤ"},
		{"ﾕ", "ユ"},
		{"ﾖ", "ヨ"},
		{"ﾗ", "ラ"},
		{"ﾘ", "リ"},
		{"ﾙ", "ル"},
		{"ﾚ", "レ"},
		{"ﾛ", "ロ"},
		{"ﾜ", "ワ"},
		{"ｦ", "ヲ"},
		{"ﾝ", "ン"},
		{"ｧ", "ァ"},
		{"ｨ", "ィ"},
		{"ｩ", "ゥ"},
		{"ｪ", "ェ"},
		{"ｫ", "ォ"},
		{"ｯ", "ッ"},
		{"ｬ", "ャ"},
		{"ｭ", "ュ"},
		{"ｮ", "ョ"},

		// replace characters which like Hyphen or Minus sign
		{"\u02D7", "-"},
		{"\u058A", "-"},
		{"\u2010", "-"},
		{"\u2011", "-"},
		{"\u2012", "-"},
		{"\u2013", "-"},
		{"\u2043", "-"},
		{"\u207B", "-"},
		{"\u208B", "-"},
		{"\u2212", "-"},

		// replace characters that like Japanese Prolonged Sound Mark "ー"
		{"\u2014", "ー"},
		{"\u2015", "ー"},
		{"\u2500", "ー"},
		{"\u2501", "ー"},
		{"\uFE63", "ー"},
		{"\uFF0D", "ー"},
		{"\uFF70", "ー"},

		// replace or delete characters that like tilde
		{"~", "ー"},
		{"∼", "ー"},
		{"∾", "ー"},
		{"〰", "ー"},
		{"～", "ー"},

		// replace Zenkaku Symbols below to Hankaku
		// ！”＃＄％＆’（）＊＋，−．／：；＜＞？＠［￥］＾＿｀｛｜｝
		{"！", "!"},
		{"”", "\""},
		{"＃", "#"},
		{"＄", "$"},
		{"％", "%"},
		{"＆", "&"},
		{"’", "'"},
		{"（", "("},
		{"）", ")"},
		{"＊", "*"},
		{"＋", "+"},
		{"，", ","},
		{"−", "-"},
		{"．", "."},
		{"／", "/"},
		{"：", ":"},
		{"；", ";"},
		{"＜", "<"},
		{"＝", "="},
		{"＞", ">"},
		{"？", "?"},
		{"＠", "@"},
		{"［", "["},
		{"¥", "\""},
		{"］", "]"},
		{"＾", "^"},
		{"＿", "_"},
		{"‘", "`"},
		{"｛", "{"},
		{"｜", "|"},
		{"｝", "}"},

		// replace Hankaku Symbols below to Zenkaku
		// ｡､･=｢｣
		{"｡", "。"},
		{"､", "、"},
		{"･", "・"},
		{"゛", "ﾞ"},
		{"゜", "ﾟ"},
		{"｢", "「"},
		{"｣", "」"},
		{"ｰ", "ー"},

		// replace Zenkaku Space (　) to Hankaku Space ( )
		{"　", " "},
	}

	// apply
	for _, r := range reps {
		s = strings.ReplaceAll(s, r.old, r.new)
	}

	regReps := [...]regRep{
		// replace characters that like two or more occurrences of Japanese Prolonged Sound Mark "ー"
		{regexp.MustCompile(`ー+`), "ー"},

		// replace characters that like two or more occurrences of Hankaku Space
		{regexp.MustCompile(` +`), " "},

		// Delete heading Hankaku Space and tailing Hankaku Space
		{regexp.MustCompile(`^ `), ""},
		{regexp.MustCompile(` $`), ""},

		// Delete Hankaku Space among Hiragana, Katakana, and Kanji
		{regexp.MustCompile(`([^a-zA-Z0-9]) ([^a-zA-Z0-9])`), "$1$2"},
		{regexp.MustCompile(`([^a-zA-Z0-9]) ([^a-zA-Z0-9])`), "$1$2"},

		// Delete Hankaku Space between Hiragana, Katakana, and Kanji, and Alphabet
		{regexp.MustCompile(`([^a-zA-Z0-9]) ([a-zA-Z0-9])`), "$1$2"},
		{regexp.MustCompile(`([^a-zA-Z0-9]) ([a-zA-Z0-9])`), "$1$2"},
		{regexp.MustCompile(`([a-zA-Z0-9]) ([^a-zA-Z0-9])`), "$1$2"},
		{regexp.MustCompile(`([a-zA-Z0-9]) ([^a-zA-Z0-9])`), "$1$2"},
	}

	// apply
	for _, regRep := range regReps {
		s = regRep.regexp.ReplaceAllString(s, regRep.new)
	}

	return s
}

// Normalize Pronunciation
func postNormalizer(pronunciation string) string {
	reps := [...]rep{
		// replace Zenkaku Katakana
		{"クヮ", "クァ"},
		{"グヮ", "グァ"},

		{"アア", "アー"},
		{"イイ", "イー"},
		{"ウウ", "ウー"},
		{"エエ", "エー"},
		{"オオ", "オー"},
		{"カア", "カー"},
		{"キイ", "キー"},
		{"クウ", "クー"},
		{"ケエ", "ケー"},
		{"コオ", "コー"},
		{"サア", "サー"},
		{"シイ", "シー"},
		{"スウ", "スー"},
		{"セエ", "セー"},
		{"ソオ", "ソー"},
		{"タア", "ター"},
		{"チイ", "チー"},
		{"ツウ", "ツー"},
		{"テエ", "テー"},
		{"トオ", "トー"},
		{"ナア", "ナー"},
		{"ニイ", "ニー"},
		{"ヌウ", "ヌー"},
		{"ネエ", "ネー"},
		{"ノオ", "ノー"},
		{"ハア", "ハー"},
		{"ヒイ", "ヒー"},
		{"フウ", "フー"},
		{"ヘエ", "ヘー"},
		{"ホオ", "ホー"},
		{"マア", "マー"},
		{"ミイ", "ミー"},
		{"ムウ", "ムー"},
		{"メエ", "メー"},
		{"モオ", "モー"},
		{"ヤア", "ヤー"},
		{"ユウ", "ユー"},
		{"ヨオ", "ヨー"},
		{"ラア", "ラー"},
		{"リイ", "リー"},
		{"ルウ", "ルー"},
		{"レエ", "レー"},
		{"ロオ", "ロー"},
		{"ワア", "ワー"},
		{"ヲオ", "ヲー"},
		{"ガア", "ガー"},
		{"ギイ", "ギー"},
		{"グウ", "グー"},
		{"ゲエ", "ゲー"},
		{"ゴオ", "ゴー"},
		{"ザア", "ザー"},
		{"ジイ", "ジー"},
		{"ズウ", "ズー"},
		{"ゼエ", "ゼー"},
		{"ゾオ", "ゾー"},
		{"ダア", "ダー"},
		{"ヂイ", "ジー"},
		{"ヅウ", "ズー"},
		{"デエ", "デー"},
		{"ドオ", "ドー"},
		{"バア", "バー"},
		{"ビイ", "ビー"},
		{"ブウ", "ブー"},
		{"ベエ", "ベー"},
		{"ボオ", "ボー"},
		{"パア", "パー"},
		{"ピイ", "ピー"},
		{"プウ", "プー"},
		{"ペエ", "ペー"},
		{"ポオ", "ポー"},

		{"アァ", "アー"},
		{"イィ", "イー"},
		{"ウゥ", "ウー"},
		{"エェ", "エー"},
		{"オォ", "オー"},
		{"カァ", "カー"},
		{"キィ", "キー"},
		{"クゥ", "クー"},
		{"ケェ", "ケー"},
		{"コォ", "コー"},
		{"サァ", "サー"},
		{"シィ", "シー"},
		{"スゥ", "スー"},
		{"セェ", "セー"},
		{"ソォ", "ソー"},
		{"タァ", "ター"},
		{"チィ", "チー"},
		{"ツゥ", "ツー"},
		{"テェ", "テー"},
		{"トォ", "トー"},
		{"ナァ", "ナー"},
		{"ニィ", "ニー"},
		{"ヌゥ", "ヌー"},
		{"ネェ", "ネー"},
		{"ノォ", "ノー"},
		{"ハァ", "ハー"},
		{"ヒィ", "ヒー"},
		{"フゥ", "フー"},
		{"ヘェ", "ヘー"},
		{"ホォ", "ホー"},
		{"マァ", "マー"},
		{"ミィ", "ミー"},
		{"ムゥ", "ムー"},
		{"メェ", "メー"},
		{"モォ", "モー"},
		{"ヤァ", "ヤー"},
		{"ユゥ", "ユー"},
		{"ヨォ", "ヨー"},
		{"ラァ", "ラー"},
		{"リィ", "リー"},
		{"ルゥ", "ルー"},
		{"レェ", "レー"},
		{"ロォ", "ロー"},
		{"ワァ", "ワー"},
		{"ヲォ", "ヲー"},
		{"ガァ", "ガー"},
		{"ギィ", "ギー"},
		{"グゥ", "グー"},
		{"ゲェ", "ゲー"},
		{"ゴォ", "ゴー"},
		{"ザァ", "ザー"},
		{"ジィ", "ジー"},
		{"ズゥ", "ズー"},
		{"ゼェ", "ゼー"},
		{"ゾォ", "ゾー"},
		{"ダァ", "ダー"},
		{"ヂィ", "ジー"},
		{"ヅゥ", "ズー"},
		{"デェ", "デー"},
		{"ドォ", "ドー"},
		{"バァ", "バー"},
		{"ビィ", "ビー"},
		{"ブゥ", "ブー"},
		{"ベェ", "ベー"},
		{"ボォ", "ボー"},
		{"パァ", "パー"},
		{"ピィ", "ピー"},
		{"プゥ", "プー"},
		{"ペェ", "ペー"},
		{"ポォ", "ポー"},

		{"ャア", "ャー"},
		{"ュウ", "ュー"},
		{"ョオ", "ョー"},
		{"ャァ", "ャー"},
		{"ュゥ", "ュー"},
		{"ョォ", "ョー"},
	}
	// apply
	for _, r := range reps {
		pronunciation = strings.ReplaceAll(pronunciation, r.old, r.new)
	}
	return pronunciation
}
