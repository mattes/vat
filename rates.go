package vat

import (
	"sort"
	"strings"
	"time"
)

// reference:
// http://ec.europa.eu/taxation_customs/resources/documents/taxation/vat/how_vat_works/rates/vat_rates_en.pdf
// last updated: March 6, 2015

var standardRate = map[string]float64{
	"BE": 21,
	"BG": 20,
	"CZ": 21,
	"DK": 25,
	"DE": 19,
	"EE": 20,
	"EL": 23,
	"ES": 21,
	"FR": 20,
	"HR": 25,
	"IE": 23,
	"IT": 22,
	"CY": 19,
	"LV": 21,
	"LT": 21,
	"LU": 17,
	"HU": 27,
	"MT": 18,
	"NL": 21,
	"AT": 20,
	"PL": 23,
	"PT": 23,
	"RO": 24,
	"SI": 22,
	"SK": 20,
	"FI": 24,
	"SE": 25,
	"UK": 20,
}

// StandardRate returns VAT rate in EU country at time.Now()
func StandardRate(countryCode string) (rate float64, ok bool) {
	return StandardRateAtDate(countryCode, time.Now())
}

// StandardRateAtDate returns VAT rate in EU country at given date
func StandardRateAtDate(countryCode string, date time.Time) (rate float64, ok bool) {
	// TODO ignore date for now, but implement it should new rates apply
	countryCode = strings.ToUpper(countryCode)
	rate, ok = standardRate[countryCode]
	return
}

// Countries returns a list of all EU countries
func Countries() []string {
	keys := []string{}
	for k := range standardRate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
