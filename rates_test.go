package vat

import (
	"testing"
)

func TestStandardRate(t *testing.T) {
	rate, ok := StandardRate("DE")
	if !ok {
		t.Error("should give rate for DE")
	}
	if rate != 19 {
		t.Error("wrong rate for DE")
	}

	_, ok = StandardRate("de")
	if !ok {
		t.Error("should give rate for DE")
	}

	_, ok = StandardRate("ZZ")
	if ok {
		t.Error("don't give rate for unknown countries")
	}
}
