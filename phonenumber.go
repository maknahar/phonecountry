//package phonecountry provide information about a phone number
package phonecountry

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/maknahar/phonecountry/internal/content"
)

var phoneMap map[string]string
var reversePhoneMap map[string]string

func init() {

	phoneMap = make(map[string]string)

	// Decode JSON into our map
	err := json.Unmarshal([]byte(content.PhoneCode), &phoneMap)
	if err != nil {
		println(err)
	}

	reversePhoneMap = make(map[string]string)
	for k, v := range phoneMap {
		reversePhoneMap[strings.ToUpper(v)] = k
	}
}

//GetCountryNameFromPhone returns country common name for given phone number
//Example +919445454528 - India
func GetCountryNameFromPhone(phoneNumber string) (countryName string, err error) {
	//strip + from starting
	phoneNumber = strings.TrimSpace(phoneNumber)
	phoneNumber = strings.TrimLeft(phoneNumber, "+")

	//check minimum length
	if len(phoneNumber) < 7 {
		return "", errors.New("Minimum length of a number should be atleast 7 digit")
	}

	//handle special cases
	switch phoneNumber[0:1] {
	case "1":
		c, ok := phoneMap[phoneNumber[0:4]]
		if ok {
			return c, nil
		}
		return "", fmt.Errorf("No country found with country code %s", phoneNumber[0:4])

	case "3":
		if phoneNumber[0:2] == "39" && phoneNumber[0:5] == "39066" {
			return phoneMap["39066"], nil
		}

	case "4":
		//Guernsey
		if phoneNumber[0:6] == "441481" || phoneNumber[0:6] == "447839" ||
			phoneNumber[0:6] == "447911" {
			return phoneMap["39066"], nil
		}
		//Isle of Man
		if phoneNumber[0:6] == "441624" || phoneNumber[0:6] == "447524" ||
			phoneNumber[0:6] == "447624" || phoneNumber[0:6] == "447924" {
			return phoneMap["447924"], nil
		}
		//Jersey
		if phoneNumber[0:6] == "441534" || phoneNumber[0:6] == "447509" ||
			phoneNumber[0:6] == "447797" || phoneNumber[0:6] == "447937" ||
			phoneNumber[0:6] == "447700" || phoneNumber[0:6] == "447829" {
			return phoneMap["447829"], nil
		}
	case "6":
		fmt.Println("here")
		if phoneNumber[0:4] == "6721" {
			return phoneMap["6721"], nil
		}
		if phoneNumber[0:4] == "6723" {
			return phoneMap["6723"], nil
		}
		//Christmas Island
		if phoneNumber[0:7] == "6189162" {
			return phoneMap["6189162"], nil
		}
		//Christmas Island
		if phoneNumber[0:7] == "6189164" {
			return phoneMap["6189164"], nil
		}
	case "7":
		if phoneNumber[0:2] == "76" || phoneNumber[0:2] == "77" {
			return phoneMap[phoneNumber[0:2]], nil
		}
		return phoneMap["7"], nil
	}
	c, ok := phoneMap[phoneNumber[0:2]]
	if ok {
		return c, nil
	}
	c, ok = phoneMap[phoneNumber[0:3]]
	if ok {
		return c, nil
	}
	return "", fmt.Errorf("No country code found for %s", phoneNumber)
}

//GetCountryISO2Code returns country iso2 code of given number
//For Example +919445454528 => IN
func GetCountryISO2Code(phoneNumber string) (countryCode string, err error) {
	countryName, err := GetCountryNameFromPhone(phoneNumber)
	if err != nil {
		return "", err
	}

	countryCode = content.CountryInfo[countryName].Cca2
	if countryCode == "" {
		return "", errors.New(fmt.Sprintf("country name %s not found", countryName))
	}
	return
}
