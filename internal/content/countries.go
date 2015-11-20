package content

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

type NativeName struct {
	//official name translation
	Official string `json:"official"`
	//common name translation
	Common string `json:"common"`
}

type Countries struct {
	Name struct {
		//common name in english
		Common string `json:"common"`

		//official name in english
		Official string `json:"official"`

		//native - list of all native names
		//   key: three-letter ISO 639-3 language code
		//   value: name object
		Native map[string]NativeName `json:"native"`
	} `json:"name"`

	//country code top-level domain (tld)
	Tld []string `json:"tld"`

	//code ISO 3166-1 alpha-2 (cca2)
	Cca2 string `json:"cca2"`

	//code ISO 3166-1 numeric (ccn3)
	Ccn3 string `json:"ccn3"`

	//code ISO 3166-1 alpha-3 (cca3)
	Cca3 string `json:"cca3"`

	//code International Olympic Committee (cioc)
	Cioc string `json:"cioc"`

	//ISO 4217 currency code(s) (currency)
	Currency []string `json:"currency"`

	//calling code(s) (callingCode)
	Callingcode []string `json:"callingCode"`

	//capital city (capital)
	Capital string `json:"capital"`

	//alternative spellings (altSpellings)
	Altspellings []string `json:"altSpellings"`

	//region
	Region string `json:"region"`

	//subregion
	Subregion string `json:"subregion"`

	//list of official languages (languages)
	//   key: three-letter ISO 639-3 language code
	//   value: name of the language in english

	Languages map[string]string `json:"languages"`

	//list of name translations (translations)
	//    key: three-letter ISO 639-3 language code
	//    value: name object
	//        key: official - official name translation
	//        key: common - common name translation
	Translations struct {
		Deu struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"deu"`
		Fra struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fra"`
		Hrv struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"hrv"`
		Ita struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"ita"`
		Jpn struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"jpn"`
		Nld struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"nld"`
		Por struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"por"`
		Rus struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"rus"`
		Spa struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"spa"`
		Fin struct {
			Official string `json:"official"`
			Common   string `json:"common"`
		} `json:"fin"`
	} `json:"translations"`

	//   latitude and longitude (latlng)
	Latlng []float64 `json:"latlng"`

	//name of residents (demonym)
	Demonym string `json:"demonym"`

	//landlocked status (landlocked)
	Landlocked bool `json:"landlocked"`

	//land borders (borders)
	Borders []interface{} `json:"borders"`

	//land area in km² (area)
	Area float64 `json:"area"`
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

var CountryInfo map[string]Countries
var PhoneInfo map[string]Countries

func init() {
	var countryData []Countries
	_, filename, _, _ := runtime.Caller(1)
	f, err := ioutil.ReadFile(path.Join(path.Dir(filename), "countries.json"))
	handle(err)
	err = json.Unmarshal(f, &countryData)
	handle(err)
	CountryInfo = make(map[string]Countries)
	for _, v := range countryData {
		CountryInfo[v.Name.Common] = v
	}

	PhoneInfo = make(map[string]Countries)
	for _, v1 := range countryData {
		for _, v2 := range v1.Callingcode {
			PhoneInfo[v2] = v1
			//fmt.Printf("\"%s\":\"%s\",\n", v2, PhoneInfo[v2].Name.Common)
		}
	}

}
