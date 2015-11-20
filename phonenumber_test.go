//package phonecountry provide information about a phone number
package phonecountry

import (
	"testing"
)

type countryName struct {
	Phone       string
	CountryName string
	CountrCode  string
}

var tests = []countryName{
	{"+919445454528", "India", "IN"},
	{"+13102316474", "United States", "US"},
	{"+929445454528", "Pakistan", "PK"},
	{"+237691928984", "Cameroon", "CM"},
	{"+24639434384", "British Indian Ocean Territory", "IO"},
	{"+613658374537", "Australia", "AU"},
	{"+618916423423", "Christmas Island", "CX"},
	{"+124248753487", "Bahamas", "BS"},
}

func TestGetCountryNameFromPhone(t *testing.T) {
	for _, pair := range tests {
		v, err := GetCountryNameFromPhone(pair.Phone)
		if err != nil {
			t.Error(err)
		}
		if v != pair.CountryName {
			t.Error(
				"For", pair.Phone,
				"expected", pair.CountryName,
				"got", v,
			)
		}
	}
}

func TestGetCountryISO2Code(t *testing.T) {
	for _, pair := range tests {
		v, err := GetCountryISO2Code(pair.Phone)
		if err != nil {
			t.Error(err)
		}
		if v != pair.CountrCode {
			t.Error(
				"For", pair.Phone,
				"expected", pair.CountrCode,
				"got", v,
			)
		}
	}
}
