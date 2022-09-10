package dajareGo

import "regexp"

type (
	rep struct {
		old string
		new string
	}
	regRep struct {
		regexp *regexp.Regexp
		new    string
	}
)
