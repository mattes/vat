# Golang VAT number validation


[![Build Status](https://travis-ci.org/mattes/vat.svg?branch=master)](https://travis-ci.org/mattes/vat)
[![GoDoc](https://godoc.org/github.com/mattes/vat?status.svg)](https://godoc.org/github.com/mattes/vat)


Uses the official [VIES VAT number validation SOAP web service](http://ec.europa.eu/taxation_customs/vies/vatRequest.html?locale=en)
to validate european VAT numbers.

Unfortunately their service is super unreliable. WTF


## Install

```
go get -u github.com/mattes/vat
```


## Usage with Go

```go
import "github.com/mattes/vat"

response, err := vat.CheckVAT("IE6388047V")
if err != nil {
  // do sth with err
}
fmt.Println(response.Name, response.Valid)

// or ...
valid, err := vat.IsValidVAT("IE6388047V")

// increase timeout (default 10 seconds)
vat.Timeout = 10

// get VAT rates for EU countries
rate, ok := vat.StandardRate("DE")

// get applicable tax and if reverse charge is allowed,
// depending on VAT number and country
// (use at own risk!)
rate, reverseCharge, err := GetApplicableTax("DE", "")
rate, reverseCharge, err := GetApplicableTax("IE", "IE6388047V")
rate, reverseCharge, err := GetApplicableTax("CH", "")
```

## Usage via Console

There is a small cli included in this package. 

Install with ``go get -u github.com/mattes/vat/vat-check``

```
$ vat-check IE6388047V

Request date: 2015-03-06 00:00:00 +0100 CET
  VAT number: 6388047V
     Country: IE
        Name: GOOGLE IRELAND LIMITED
     Address: 3RD FLOOR ,GORDON HOUSE ,BARROW STREET ,DUBLIN 4

Success: The VAT number is valid!
```

Exit code is 0 if VAT number is valid.