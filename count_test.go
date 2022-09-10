package dajareGo

import "testing"

func TestEditDistance(t *testing.T) {
	if get := editDistance("kitten", "sitting"); get != 3 {
		t.Errorf("failed: expects\"%d\" but output was \"%d\"", 3, get)
	}
}
func TestFuzzyCount(t *testing.T) {
	if get := fuzzyCount(NewSyllables("フトンガフットンダ"), NewSyllables("フトン"), 1); get != 2 {
		t.Errorf("failed: %s expects\"%d\" but output was \"%d\"", "布団が吹っ飛んだ", 2, get)
	}
	if get := fuzzyCount(NewSyllables("アルミカンノウエニアルミカン"), NewSyllables("アルミカン"), 1); get != 2 {
		t.Errorf("failed: %s expects\"%d\" but output was \"%d\"", "アルミ缶の上にあるミカン", 2, get)
	}
}
