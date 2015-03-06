package vat

import (
	"time"
)

// GetApplicableTaxAtDate returns taxRate for person depending on
// countryCode and if vatNumber is given.
// err returns nil, ErrVATnumberNotValid or ErrVATserviceUnreachable
func GetApplicableTaxAtDate(countryCode, vatNumber string, date time.Time) (taxRate float64, reverseCharge bool, err error) {
	if vatNumber != "" {
		// eu business
		valid, err := IsValidVAT(vatNumber)
		if err != nil {
			if err == ErrVATserviceUnreachable {
				return 0, false, ErrVATserviceUnreachable
			} else {
				// override err with ErrVATnumberNotValid
				return 0, false, ErrVATnumberNotValid
			}
		} else {
			if valid {
				return 0, true, nil
			} else {
				return 0, false, ErrVATnumberNotValid
			}
		}
	} else {
		// person from EU or rest of world?
		rate, ok := StandardRateAtDate(countryCode, date)
		if ok {
			// person from eu
			return rate, false, nil
		} else {
			// person from rest of world
			return 0, false, nil
		}
	}
}

// GetApplicableTax is a convenience func for GetApplicableTaxAtDate(...)
func GetApplicableTax(countryCode, vatNumber string) (taxRate float64, reverseCharge bool, err error) {
	return GetApplicableTaxAtDate(countryCode, vatNumber, time.Now())
}
