// country project phonenumber.go
package phonecountry

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/relops/prefixes"
)

var iso2Map map[string]string
var reverseiso2Map map[string]string
var phoneMap map[string]string
var reversePhoneMap map[string]string

func init() {

	iso2Map = make(map[string]string)

	// Decode JSON into our map
	err := json.Unmarshal([]byte(phoneCode), &iso2Map)
	if err != nil {
		println(err)
	}

	reverseiso2Map = make(map[string]string)
	for k, v := range iso2Map {
		reverseiso2Map[strings.ToUpper(v)] = k
	}

	iso2Map = make(map[string]string)

	// Decode JSON into our map
	err = json.Unmarshal([]byte(isoCode), &iso2Map)
	if err != nil {
		println(err)
	}

	reverseiso2Map = make(map[string]string)
	for k, v := range iso2Map {
		reverseiso2Map[strings.ToUpper(v)] = k
	}
}

func GetCountryName(phoneNumber string) (countryName string, err error) {
	//strip + from starting
	phoneNumber = strings.TrimSpace(phoneNumber)
	phoneNumber = strings.TrimLeft(phoneNumber, "+")

	country, err := prefixes.Lookup(phoneNumber)
	if err != nil {
		return "", err
	}

	return country.Name, nil
}

func GetCountryISO2Code(phoneNumber string) (countryCode string, err error) {
	countryName, err := GetCountryName(phoneNumber)
	if err != nil {
		return "", err
	}

	countryCode = reverseiso2Map[strings.ToUpper(countryName)]
	if countryCode == "" {
		return "", errors.New(fmt.Sprintf("country name not found"))
	}
	return
}
