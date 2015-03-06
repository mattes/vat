package vat

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"text/template"
	"time"
)

type Response struct {
	CountryCode string
	VATnumber   string
	RequestDate time.Time
	Valid       bool
	Name        string
	Address     string
}

const serviceUrl = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"

const envelope = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:v1="http://schemas.conversesolutions.com/xsd/dmticta/v1">
<soapenv:Header/>
<soapenv:Body>
  <checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
    <countryCode>{{.countryCode}}</countryCode>
    <vatNumber>{{.vatNumber}}</vatNumber>
  </checkVat>
</soapenv:Body>
</soapenv:Envelope>
`

var Timeout = 10 // seconds

var (
	ErrVATnumberNotValid     = errors.New("VAT number is not valid.")
	ErrVATserviceUnreachable = errors.New("VAT number validation service is offline.")
)

// Check returns *Response for vat number
func Check(vatNumber string) (*Response, error) {
	vatNumber = sanitizeVatNumber(vatNumber)

	e, err := getEnvelope(vatNumber)
	if err != nil {
		return nil, err
	}
	eb := bytes.NewBufferString(e)
	client := http.Client{
		Timeout: time.Duration(time.Duration(Timeout) * time.Second),
	}
	res, err := client.Post(serviceUrl, "text/xml;charset=UTF-8", eb)
	if err != nil {
		return nil, ErrVATserviceUnreachable
	}
	defer res.Body.Close()
	xmlRes, err := ioutil.ReadAll(res.Body)

	if bytes.Contains(xmlRes, []byte("INVALID_INPUT")) {
		return nil, ErrVATnumberNotValid
	}

	var rd struct {
		XMLName xml.Name `xml:"Envelope"`
		Soap    struct {
			XMLName xml.Name `xml:"Body"`
			Soap    struct {
				XMLName     xml.Name `xml:"checkVatResponse"`
				CountryCode string   `xml:"countryCode"`
				VATnumber   string   `xml:"vatNumber"`
				RequestDate string   `xml:"requestDate"` // 2015-03-06+01:00
				Valid       bool     `xml:"valid"`
				Name        string   `xml:"name"`
				Address     string   `xml:"address"`
			}
		}
	}
	if err := xml.Unmarshal(xmlRes, &rd); err != nil {
		return nil, err
	}

	if rd.Soap.Soap.RequestDate == "" {
		return nil, errors.New("service returned invalid request date")
	}

	pDate, err := time.Parse("2006-01-02-07:00", rd.Soap.Soap.RequestDate)
	if err != nil {
		return nil, err
	}

	r := &Response{
		CountryCode: rd.Soap.Soap.CountryCode,
		VATnumber:   rd.Soap.Soap.VATnumber,
		RequestDate: pDate,
		Valid:       rd.Soap.Soap.Valid,
		Name:        rd.Soap.Soap.Name,
		Address:     rd.Soap.Soap.Address,
	}

	return r, nil
}

// IsValid returns true if vat number is correct
func IsValid(vatNumber string) (bool, error) {
	r, err := Check(vatNumber)
	if err != nil {
		return false, err
	}
	return r.Valid, nil
}

// sanitizeVatNumber removes white space
func sanitizeVatNumber(vatNumber string) string {
	vatNumber = strings.TrimSpace(vatNumber)
	return regexp.MustCompile(" ").ReplaceAllString(vatNumber, "")
}

// getEnvelope parses envelope template
func getEnvelope(vatNumber string) (string, error) {
	if len(vatNumber) < 3 {
		return "", errors.New("VAT number is too short.")
	}

	t, err := template.New("envelope").Parse(envelope)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	if err := t.Execute(&result, map[string]string{
		"countryCode": vatNumber[0:2],
		"vatNumber":   vatNumber[2:],
	}); err != nil {
		return "", err
	}
	return result.String(), nil
}
