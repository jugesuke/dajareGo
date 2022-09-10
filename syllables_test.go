package dajareGo

import (
	"testing"
)

func TestSyllablesNew(t *testing.T) {
	// case 1
	// 布団が吹っ飛んだ
	wont := Syllables{
		"hu",
		"ton",
		"ga",
		"hut",
		"ton",
		"da",
	}
	got := NewSyllables("フトンガフットンダ")
	if !wont.equals(got) {
		t.Error("failed: expects", wont, "but output was", got)
	}

	// case 2
	// チャットツールの導入
	wont = Syllables{
		"tiyt",
		"to",
		"tu-",
		"ru",
		"no",
		"do-",
		"niy-",
	}
	got = NewSyllables("チャットツールノドーニュウ。")
	if !wont.equals(got) {
		t.Error("failed: expects", wont, "but output was", got)
	}
}

func TestSyllablesEquals(t *testing.T) {
	syllable1 := NewSyllables("イカ")
	syllable2 := NewSyllables("タコ")
	if syllable1.equals(syllable2) {
		t.Error("failed: イカ/タコ expects", false, "but output was", true)
	}
	syllable1 = NewSyllables("ジャガイモ")
	syllable2 = NewSyllables("トマト")
	if syllable1.equals(syllable2) {
		t.Error("failed: ジャガイモ/トマト expects", false, "but output was", true)
	}
	syllable1 = NewSyllables("ダンシャクイモ")
	syllable2 = NewSyllables("ダンシャクイモ")
	if !syllable1.equals(syllable2) {
		t.Error("failed: ダンシャクイモ/ダンシャクイモ expects", true, "but output was", false)
	}
}
