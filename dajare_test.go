package dajareGo

import (
	"bufio"
	"os"
	"testing"
	"unicode/utf8"

	"github.com/ikawaha/kagome-dict/ipa"
)

func benchmarkIsDajare(b *testing.B, pn bool, testCasePath string, lengthLimit bool) {
	// loading testCase
	f, err := os.Open(testCasePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// prepare for output
	out, err := os.Create("testdata/failedDajare.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// timer reset
	b.ResetTimer()

	if err := Init(); err != nil {
		panic(err)
	}

	// scan
	testCount := 0
	passCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		test := scanner.Text()
		if lengthLimit && utf8.RuneCountInString(test) >= 18 {
			continue
		}
		testCount += 1
		r := IsDajare(test)
		isAC := (r.IsDajare == pn)
		if isAC {
			passCount += 1
		} else {
			// fmt.Println(r.IsDajare, r)
			_, err := out.WriteString(test + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	b.StopTimer()
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	b.ReportMetric((float64(passCount) / float64(testCount)), "AC/cases")
}

func BenchmarkIsDajarePositive(b *testing.B) {
	benchmarkIsDajare(b, true, "testdata/positive/dajareList.txt", false)
}

func BenchmarkIsDajareNegative(b *testing.B) {
	benchmarkIsDajare(b, false, "testdata/negative/kokoro.txt", false)
}

func BenchmarkIsDajarePositiveLengthLimit(b *testing.B) {
	benchmarkIsDajare(b, true, "testdata/positive/dajareList.txt", true)
}

func BenchmarkIsDajareNegativeLengthLimit(b *testing.B) {
	benchmarkIsDajare(b, false, "testdata/negative/kokoro.txt", true)
}

func TestIsDajare(t *testing.T) {
	if err := Init(); err != nil {
		panic(err)
	}
	result := IsDajare("アルミ缶の上にあるミカン")
	if !result.IsDajare {
		t.Errorf("failed: アルミ缶の上にあるミカン is Dajare but returned false")
	}
	result = IsDajare("布団が吹っ飛んだ")
	if !result.IsDajare {
		t.Errorf("failed: 布団が吹っ飛んだ is Dajare but returned false")
	}
	result = IsDajare("人民の人民による人民のための政治")
	if result.IsDajare {
		t.Errorf("failed: 人民の人民による人民のための政治 is NOT Dajare but returned true")
	}
	result = IsDajare("グァテマラに行ってきた")
	if result.IsDajare {
		t.Errorf("failed: グァテマラに行ってきた is NOT Dajare but returned true")
	}
}
func TestCustomDictionary(t *testing.T) {
	dict := ipa.Dict()
	if err := SetCustomDictionary(dict); err != nil {
		t.Error("failed: an error occurred. the error is", err)
	}
}
