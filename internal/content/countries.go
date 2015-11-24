package content

import "encoding/json"

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

func init() {
	var countryData []Countries
	err := json.Unmarshal([]byte(bigData), &countryData)
	handle(err)
	CountryInfo = make(map[string]Countries)
	for _, v := range countryData {
		CountryInfo[v.Name.Common] = v
	}
}

const bigData = `[
    {
        "name": {
            "common": "Aruba",
            "official": "Aruba",
            "native": {
                "nld": {
                    "official": "Aruba",
                    "common": "Aruba"
                },
                "pap": {
                    "official": "Aruba",
                    "common": "Aruba"
                }
            }
        },
        "tld": [
            ".aw"
        ],
        "cca2": "AW",
        "ccn3": "533",
        "cca3": "ABW",
        "cioc": "ARU",
        "currency": [
            "AWG"
        ],
        "callingCode": [
            "297"
        ],
        "capital": "Oranjestad",
        "altSpellings": [
            "AW"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "nld": "Dutch",
            "pap": "Papiamento"
        },
        "translations": {
            "deu": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "fra": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "hrv": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "ita": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "jpn": {
                "official": "アルバ",
                "common": "アルバ"
            },
            "nld": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "por": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "rus": {
                "official": "Аруба",
                "common": "Аруба"
            },
            "spa": {
                "official": "Aruba",
                "common": "Aruba"
            },
            "fin": {
                "official": "Aruba",
                "common": "Aruba"
            }
        },
        "latlng": [
            12.5,
            -69.96666666
        ],
        "demonym": "Aruban",
        "landlocked": false,
        "borders": [],
        "area": 180
    },
    {
        "name": {
            "common": "Afghanistan",
            "official": "Islamic Republic of Afghanistan",
            "native": {
                "prs": {
                    "official": "جمهوری اسلامی افغانستان",
                    "common": "افغانستان"
                },
                "pus": {
                    "official": "د افغانستان اسلامي جمهوریت",
                    "common": "افغانستان"
                },
                "tuk": {
                    "official": "Owganystan Yslam Respublikasy",
                    "common": "Owganystan"
                }
            }
        },
        "tld": [
            ".af"
        ],
        "cca2": "AF",
        "ccn3": "004",
        "cca3": "AFG",
        "cioc": "AFG",
        "currency": [
            "AFN"
        ],
        "callingCode": [
            "93"
        ],
        "capital": "Kabul",
        "altSpellings": [
            "AF",
            "Afġānistān"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "prs": "Dari",
            "pus": "Pashto",
            "tuk": "Turkmen"
        },
        "translations": {
            "cym": {
                "official": "Islamic Republic of Afghanistan",
                "common": "Affganistan"
            },
            "deu": {
                "official": "Islamische Republik Afghanistan",
                "common": "Afghanistan"
            },
            "fra": {
                "official": "République islamique d'Afghanistan",
                "common": "Afghanistan"
            },
            "hrv": {
                "official": "Islamska Republika Afganistan",
                "common": "Afganistan"
            },
            "ita": {
                "official": "Repubblica islamica dell'Afghanistan",
                "common": "Afghanistan"
            },
            "jpn": {
                "official": "アフガニスタン·イスラム共和国",
                "common": "アフガニスタン"
            },
            "nld": {
                "official": "Islamitische Republiek Afghanistan",
                "common": "Afghanistan"
            },
            "por": {
                "official": "República Islâmica do Afeganistão",
                "common": "Afeganistão"
            },
            "rus": {
                "official": "Исламская Республика Афганистан",
                "common": "Афганистан"
            },
            "spa": {
                "official": "República Islámica de Afganistán",
                "common": "Afganistán"
            },
            "fin": {
                "official": "Afganistanin islamilainen tasavalta",
                "common": "Afganistan"
            }
        },
        "latlng": [
            33,
            65
        ],
        "demonym": "Afghan",
        "landlocked": true,
        "borders": [
            "IRN",
            "PAK",
            "TKM",
            "UZB",
            "TJK",
            "CHN"
        ],
        "area": 652230
    },
    {
        "name": {
            "common": "Angola",
            "official": "Republic of Angola",
            "native": {
                "por": {
                    "official": "República de Angola",
                    "common": "Angola"
                }
            }
        },
        "tld": [
            ".ao"
        ],
        "cca2": "AO",
        "ccn3": "024",
        "cca3": "AGO",
        "cioc": "ANG",
        "currency": [
            "AOA"
        ],
        "callingCode": [
            "244"
        ],
        "capital": "Luanda",
        "altSpellings": [
            "AO",
            "República de Angola",
            "ʁɛpublika de an'ɡɔla"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "cym": {
                "official": "Republic of Angola",
                "common": "Angola"
            },
            "deu": {
                "official": "Republik Angola",
                "common": "Angola"
            },
            "fra": {
                "official": "République d'Angola",
                "common": "Angola"
            },
            "hrv": {
                "official": "Republika Angola",
                "common": "Angola"
            },
            "ita": {
                "official": "Repubblica dell'Angola",
                "common": "Angola"
            },
            "jpn": {
                "official": "アンゴラ共和国",
                "common": "アンゴラ"
            },
            "nld": {
                "official": "Republiek Angola",
                "common": "Angola"
            },
            "por": {
                "official": "República de Angola",
                "common": "Angola"
            },
            "rus": {
                "official": "Республика Ангола",
                "common": "Ангола"
            },
            "spa": {
                "official": "República de Angola",
                "common": "Angola"
            },
            "fin": {
                "official": "Angolan tasavalta",
                "common": "Angola"
            }
        },
        "latlng": [
            -12.5,
            18.5
        ],
        "demonym": "Angolan",
        "landlocked": false,
        "borders": [
            "COG",
            "COD",
            "ZMB",
            "NAM"
        ],
        "area": 1246700
    },
    {
        "name": {
            "common": "Anguilla",
            "official": "Anguilla",
            "native": {
                "eng": {
                    "official": "Anguilla",
                    "common": "Anguilla"
                }
            }
        },
        "tld": [
            ".ai"
        ],
        "cca2": "AI",
        "ccn3": "660",
        "cca3": "AIA",
        "cioc": "",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1264"
        ],
        "capital": "The Valley",
        "altSpellings": [
            "AI"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Anguilla",
                "common": "Anguilla"
            },
            "fra": {
                "official": "Anguilla",
                "common": "Anguilla"
            },
            "hrv": {
                "official": "Anguilla",
                "common": "Angvila"
            },
            "ita": {
                "official": "Anguilla",
                "common": "Anguilla"
            },
            "jpn": {
                "official": "アングィラ",
                "common": "アンギラ"
            },
            "nld": {
                "official": "Anguilla",
                "common": "Anguilla"
            },
            "por": {
                "official": "Anguilla",
                "common": "Anguilla"
            },
            "rus": {
                "official": "Ангилья",
                "common": "Ангилья"
            },
            "spa": {
                "official": "Anguila",
                "common": "Anguilla"
            },
            "fin": {
                "official": "Anguilla",
                "common": "Anguilla"
            }
        },
        "latlng": [
            18.25,
            -63.16666666
        ],
        "demonym": "Anguillian",
        "landlocked": false,
        "borders": [],
        "area": 91
    },
    {
        "name": {
            "common": "Åland Islands",
            "official": "Åland Islands",
            "native": {
                "swe": {
                    "official": "Landskapet Åland",
                    "common": "Åland"
                }
            }
        },
        "tld": [
            ".ax"
        ],
        "cca2": "AX",
        "ccn3": "248",
        "cca3": "ALA",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "358"
        ],
        "capital": "Mariehamn",
        "altSpellings": [
            "AX",
            "Aaland",
            "Aland",
            "Ahvenanmaa"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "swe": "Swedish"
        },
        "translations": {
            "deu": {
                "official": "Åland-Inseln",
                "common": "Åland"
            },
            "fra": {
                "official": "Ahvenanmaa",
                "common": "Ahvenanmaa"
            },
            "hrv": {
                "official": "Aland Islands",
                "common": "Ålandski otoci"
            },
            "ita": {
                "official": "Isole Åland",
                "common": "Isole Aland"
            },
            "jpn": {
                "official": "オーランド諸島",
                "common": "オーランド諸島"
            },
            "nld": {
                "official": "Åland eilanden",
                "common": "Ålandeilanden"
            },
            "por": {
                "official": "Ilhas Åland",
                "common": "Alândia"
            },
            "rus": {
                "official": "Аландские острова",
                "common": "Аландские острова"
            },
            "spa": {
                "official": "Islas Åland",
                "common": "Alandia"
            },
            "fin": {
                "official": "Ahvenanmaan maakunta",
                "common": "Ahvenanmaa"
            }
        },
        "latlng": [
            60.116667,
            19.9
        ],
        "demonym": "Ålandish",
        "landlocked": false,
        "borders": [],
        "area": 1580
    },
    {
        "name": {
            "common": "Albania",
            "official": "Republic of Albania",
            "native": {
                "sqi": {
                    "official": "Republika e Shqipërisë",
                    "common": "Shqipëria"
                }
            }
        },
        "tld": [
            ".al"
        ],
        "cca2": "AL",
        "ccn3": "008",
        "cca3": "ALB",
        "cioc": "ALB",
        "currency": [
            "ALL"
        ],
        "callingCode": [
            "355"
        ],
        "capital": "Tirana",
        "altSpellings": [
            "AL",
            "Shqipëri",
            "Shqipëria",
            "Shqipnia"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "sqi": "Albanian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Albania",
                "common": "Albania"
            },
            "deu": {
                "official": "Republik Albanien",
                "common": "Albanien"
            },
            "fra": {
                "official": "République d'Albanie",
                "common": "Albanie"
            },
            "hrv": {
                "official": "Republika Albanija",
                "common": "Albanija"
            },
            "ita": {
                "official": "Repubblica d'Albania",
                "common": "Albania"
            },
            "jpn": {
                "official": "アルバニア共和国",
                "common": "アルバニア"
            },
            "nld": {
                "official": "Republiek Albanië",
                "common": "Albanië"
            },
            "por": {
                "official": "República da Albânia",
                "common": "Albânia"
            },
            "rus": {
                "official": "Республика Албания",
                "common": "Албания"
            },
            "spa": {
                "official": "República de Albania",
                "common": "Albania"
            },
            "fin": {
                "official": "Albanian tasavalta",
                "common": "Albania"
            }
        },
        "latlng": [
            41,
            20
        ],
        "demonym": "Albanian",
        "landlocked": false,
        "borders": [
            "MNE",
            "GRC",
            "MKD",
            "KOS"
        ],
        "area": 28748
    },
    {
        "name": {
            "common": "Andorra",
            "official": "Principality of Andorra",
            "native": {
                "cat": {
                    "official": "Principat d'Andorra",
                    "common": "Andorra"
                }
            }
        },
        "tld": [
            ".ad"
        ],
        "cca2": "AD",
        "ccn3": "020",
        "cca3": "AND",
        "cioc": "AND",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "376"
        ],
        "capital": "Andorra la Vella",
        "altSpellings": [
            "AD",
            "Principality of Andorra",
            "Principat d'Andorra"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "cat": "Catalan"
        },
        "translations": {
            "cym": {
                "official": "Principality of Andorra",
                "common": "Andorra"
            },
            "deu": {
                "official": "Fürstentum Andorra",
                "common": "Andorra"
            },
            "fra": {
                "official": "Principauté d'Andorre",
                "common": "Andorre"
            },
            "hrv": {
                "official": "Kneževina Andora",
                "common": "Andora"
            },
            "ita": {
                "official": "Principato di Andorra",
                "common": "Andorra"
            },
            "jpn": {
                "official": "アンドラ公国",
                "common": "アンドラ"
            },
            "nld": {
                "official": "Prinsdom Andorra",
                "common": "Andorra"
            },
            "por": {
                "official": "Principado de Andorra",
                "common": "Andorra"
            },
            "rus": {
                "official": "Княжество Андорра",
                "common": "Андорра"
            },
            "spa": {
                "official": "Principado de Andorra",
                "common": "Andorra"
            },
            "fin": {
                "official": "Andorran ruhtinaskunta",
                "common": "Andorra"
            }
        },
        "latlng": [
            42.5,
            1.5
        ],
        "demonym": "Andorran",
        "landlocked": true,
        "borders": [
            "FRA",
            "ESP"
        ],
        "area": 468
    },
    {
        "name": {
            "common": "United Arab Emirates",
            "official": "United Arab Emirates",
            "native": {
                "ara": {
                    "official": "الإمارات العربية المتحدة",
                    "common": "دولة الإمارات العربية المتحدة"
                }
            }
        },
        "tld": [
            ".ae",
            "امارات."
        ],
        "cca2": "AE",
        "ccn3": "784",
        "cca3": "ARE",
        "cioc": "UAE",
        "currency": [
            "AED"
        ],
        "callingCode": [
            "971"
        ],
        "capital": "Abu Dhabi",
        "altSpellings": [
            "AE",
            "UAE",
            "Emirates"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Vereinigte Arabische Emirate",
                "common": "Vereinigte Arabische Emirate"
            },
            "fra": {
                "official": "Émirats arabes unis",
                "common": "Émirats arabes unis"
            },
            "hrv": {
                "official": "Ujedinjeni Arapski Emirati",
                "common": "Ujedinjeni Arapski Emirati"
            },
            "ita": {
                "official": "Emirati Arabi Uniti",
                "common": "Emirati Arabi Uniti"
            },
            "jpn": {
                "official": "アラブ首長国連邦",
                "common": "アラブ首長国連邦"
            },
            "nld": {
                "official": "Verenigde Arabische Emiraten",
                "common": "Verenigde Arabische Emiraten"
            },
            "por": {
                "official": "Emirados Árabes Unidos",
                "common": "Emirados Árabes Unidos"
            },
            "rus": {
                "official": "Объединенные Арабские Эмираты",
                "common": "Объединённые Арабские Эмираты"
            },
            "spa": {
                "official": "Emiratos Árabes Unidos",
                "common": "Emiratos Árabes Unidos"
            },
            "fin": {
                "official": "Yhdistyneet arabiemiirikunnat",
                "common": "Arabiemiraatit"
            }
        },
        "latlng": [
            24,
            54
        ],
        "demonym": "Emirati",
        "landlocked": false,
        "borders": [
            "OMN",
            "SAU"
        ],
        "area": 83600
    },
    {
        "name": {
            "common": "Argentina",
            "official": "Argentine Republic",
            "native": {
                "grn": {
                    "official": "Argentine Republic",
                    "common": "Argentina"
                },
                "spa": {
                    "official": "República Argentina",
                    "common": "Argentina"
                }
            }
        },
        "tld": [
            ".ar"
        ],
        "cca2": "AR",
        "ccn3": "032",
        "cca3": "ARG",
        "cioc": "ARG",
        "currency": [
            "ARS"
        ],
        "callingCode": [
            "54"
        ],
        "capital": "Buenos Aires",
        "altSpellings": [
            "AR",
            "Argentine Republic",
            "República Argentina"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "grn": "Guaraní",
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Argentine Republic",
                "common": "Ariannin"
            },
            "deu": {
                "official": "Argentinische Republik",
                "common": "Argentinien"
            },
            "fra": {
                "official": "République argentine",
                "common": "Argentine"
            },
            "hrv": {
                "official": "Argentinski Republika",
                "common": "Argentina"
            },
            "ita": {
                "official": "Repubblica Argentina",
                "common": "Argentina"
            },
            "jpn": {
                "official": "アルゼンチン共和国",
                "common": "アルゼンチン"
            },
            "nld": {
                "official": "Argentijnse Republiek",
                "common": "Argentinië"
            },
            "por": {
                "official": "República Argentina",
                "common": "Argentina"
            },
            "rus": {
                "official": "Аргентинская Республика",
                "common": "Аргентина"
            },
            "spa": {
                "official": "República Argentina",
                "common": "Argentina"
            },
            "fin": {
                "official": "Argentiinan tasavalta",
                "common": "Argentiina"
            }
        },
        "latlng": [
            -34,
            -64
        ],
        "demonym": "Argentinean",
        "landlocked": false,
        "borders": [
            "BOL",
            "BRA",
            "CHL",
            "PRY",
            "URY"
        ],
        "area": 2780400
    },
    {
        "name": {
            "common": "Armenia",
            "official": "Republic of Armenia",
            "native": {
                "hye": {
                    "official": "Հայաստանի Հանրապետություն",
                    "common": "Հայաստան"
                },
                "rus": {
                    "official": "Республика Армения",
                    "common": "Армения"
                }
            }
        },
        "tld": [
            ".am"
        ],
        "cca2": "AM",
        "ccn3": "051",
        "cca3": "ARM",
        "cioc": "ARM",
        "currency": [
            "AMD"
        ],
        "callingCode": [
            "374"
        ],
        "capital": "Yerevan",
        "altSpellings": [
            "AM",
            "Hayastan",
            "Republic of Armenia",
            "Հայաստանի Հանրապետություն"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "hye": "Armenian",
            "rus": "Russian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Armenia",
                "common": "Armenia"
            },
            "deu": {
                "official": "Republik Armenien",
                "common": "Armenien"
            },
            "fra": {
                "official": "République d'Arménie",
                "common": "Arménie"
            },
            "hrv": {
                "official": "Republika Armenija",
                "common": "Armenija"
            },
            "ita": {
                "official": "Repubblica di Armenia",
                "common": "Armenia"
            },
            "jpn": {
                "official": "アルメニア共和国",
                "common": "アルメニア"
            },
            "nld": {
                "official": "Republiek Armenië",
                "common": "Armenië"
            },
            "por": {
                "official": "República da Arménia",
                "common": "Arménia"
            },
            "rus": {
                "official": "Республика Армения",
                "common": "Армения"
            },
            "spa": {
                "official": "República de Armenia",
                "common": "Armenia"
            },
            "fin": {
                "official": "Armenian tasavalta",
                "common": "Armenia"
            }
        },
        "latlng": [
            40,
            45
        ],
        "demonym": "Armenian",
        "landlocked": true,
        "borders": [
            "AZE",
            "GEO",
            "IRN",
            "TUR"
        ],
        "area": 29743
    },
    {
        "name": {
            "common": "American Samoa",
            "official": "American Samoa",
            "native": {
                "eng": {
                    "official": "American Samoa",
                    "common": "American Samoa"
                },
                "smo": {
                    "official": "Sāmoa Amelika",
                    "common": "Sāmoa Amelika"
                }
            }
        },
        "tld": [
            ".as"
        ],
        "cca2": "AS",
        "ccn3": "016",
        "cca3": "ASM",
        "cioc": "ASA",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1684"
        ],
        "capital": "Pago Pago",
        "altSpellings": [
            "AS",
            "Amerika Sāmoa",
            "Amelika Sāmoa",
            "Sāmoa Amelika"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "smo": "Samoan"
        },
        "translations": {
            "deu": {
                "official": "Amerikanisch-Samoa",
                "common": "Amerikanisch-Samoa"
            },
            "fra": {
                "official": "Samoa américaines",
                "common": "Samoa américaines"
            },
            "hrv": {
                "official": "američka Samoa",
                "common": "Američka Samoa"
            },
            "ita": {
                "official": "Samoa americane",
                "common": "Samoa Americane"
            },
            "jpn": {
                "official": "米サモア",
                "common": "アメリカ領サモア"
            },
            "nld": {
                "official": "Amerikaans Samoa",
                "common": "Amerikaans Samoa"
            },
            "por": {
                "official": "Samoa americana",
                "common": "Samoa Americana"
            },
            "rus": {
                "official": "американское Самоа",
                "common": "Американское Самоа"
            },
            "spa": {
                "official": "Samoa Americana",
                "common": "Samoa Americana"
            },
            "fin": {
                "official": "Amerikan Samoa",
                "common": "Amerikan Samoa"
            }
        },
        "latlng": [
            -14.33333333,
            -170
        ],
        "demonym": "American Samoan",
        "landlocked": false,
        "borders": [],
        "area": 199
    },
    {
        "name": {
            "common": "Antarctica",
            "official": "Antarctica",
            "native": {}
        },
        "tld": [
            ".aq"
        ],
        "cca2": "AQ",
        "ccn3": "010",
        "cca3": "ATA",
        "cioc": "",
        "currency": [],
        "callingCode": [],
        "capital": "",
        "altSpellings": [
            "AQ"
        ],
        "region": "",
        "subregion": "",
        "languages": {},
        "translations": {
            "cym": {
                "official": "Antarctica",
                "common": "Antarctica"
            },
            "deu": {
                "official": "Antarktika",
                "common": "Antarktis"
            },
            "fra": {
                "official": "Antarctique",
                "common": "Antarctique"
            },
            "hrv": {
                "official": "Antarktika",
                "common": "Antarktika"
            },
            "ita": {
                "official": "Antartide",
                "common": "Antartide"
            },
            "jpn": {
                "official": "南極大陸",
                "common": "南極"
            },
            "nld": {
                "official": "Antarctica",
                "common": "Antarctica"
            },
            "por": {
                "official": "Antártica",
                "common": "Antártida"
            },
            "rus": {
                "official": "Антарктида",
                "common": "Антарктида"
            },
            "spa": {
                "official": "Antártida",
                "common": "Antártida"
            },
            "fin": {
                "official": "Etelämanner",
                "common": "Etelämanner"
            }
        },
        "latlng": [
            -90,
            0
        ],
        "demonym": "Antarctican",
        "landlocked": false,
        "borders": [],
        "area": 14000000
    },
    {
        "name": {
            "common": "French Southern and Antarctic Lands",
            "official": "Territory of the French Southern and Antarctic Lands",
            "native": {
                "fra": {
                    "official": "Territoire des Terres australes et antarctiques françaises",
                    "common": "Terres australes et antarctiques françaises"
                }
            }
        },
        "tld": [
            ".tf"
        ],
        "cca2": "TF",
        "ccn3": "260",
        "cca3": "ATF",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [],
        "capital": "Port-aux-Français",
        "altSpellings": [
            "TF",
            "French Southern Territories"
        ],
        "region": "",
        "subregion": "",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Gebiet der Französisch Süd-und Antarktisgebiete",
                "common": "Französische Süd-und Antarktisgebiete"
            },
            "fra": {
                "official": "Territoire des Terres australes et antarctiques françaises",
                "common": "Terres australes et antarctiques françaises"
            },
            "hrv": {
                "official": "Teritoriju Francuski južni i antarktički teritoriji",
                "common": "Francuski južni i antarktički teritoriji"
            },
            "ita": {
                "official": "Territorio della australi e antartiche francesi Terre",
                "common": "Territori Francesi del Sud"
            },
            "jpn": {
                "official": "フランス領南方·南極地域の領土",
                "common": "フランス領南方・南極地域"
            },
            "nld": {
                "official": "Grondgebied van de Franse Zuidelijke en Antarctische gebieden",
                "common": "Franse Gebieden in de zuidelijke Indische Oceaan"
            },
            "por": {
                "official": "Território do Sul e Antártica Francesa",
                "common": "Terras Austrais e Antárticas Francesas"
            },
            "rus": {
                "official": "Территория Французские Южные и Антарктические земли",
                "common": "Французские Южные и Антарктические территории"
            },
            "spa": {
                "official": "Territorio del Francés Tierras australes y antárticas",
                "common": "Tierras Australes y Antárticas Francesas"
            },
            "fin": {
                "official": "Ranskan eteläiset ja antarktiset alueet",
                "common": "Ranskan eteläiset ja antarktiset alueet"
            }
        },
        "latlng": [
            -49.25,
            69.167
        ],
        "demonym": "French",
        "landlocked": false,
        "borders": [],
        "area": 7747
    },
    {
        "name": {
            "common": "Antigua and Barbuda",
            "official": "Antigua and Barbuda",
            "native": {
                "eng": {
                    "official": "Antigua and Barbuda",
                    "common": "Antigua and Barbuda"
                }
            }
        },
        "tld": [
            ".ag"
        ],
        "cca2": "AG",
        "ccn3": "028",
        "cca3": "ATG",
        "cioc": "ANT",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1268"
        ],
        "capital": "Saint John's",
        "altSpellings": [
            "AG"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Antigua and Barbuda",
                "common": "Antigwa a Barbiwda"
            },
            "deu": {
                "official": "Antigua und Barbuda",
                "common": "Antigua und Barbuda"
            },
            "fra": {
                "official": "Antigua -et-Barbuda",
                "common": "Antigua-et-Barbuda"
            },
            "hrv": {
                "official": "Antigva i Barbuda",
                "common": "Antigva i Barbuda"
            },
            "ita": {
                "official": "Antigua e Barbuda",
                "common": "Antigua e Barbuda"
            },
            "jpn": {
                "official": "アンチグアバーブーダ",
                "common": "アンティグア・バーブーダ"
            },
            "nld": {
                "official": "Antigua en Barbuda",
                "common": "Antigua en Barbuda"
            },
            "por": {
                "official": "Antigua e Barbuda",
                "common": "Antígua e Barbuda"
            },
            "rus": {
                "official": "Антигуа и Барбуда",
                "common": "Антигуа и Барбуда"
            },
            "spa": {
                "official": "Antigua y Barbuda",
                "common": "Antigua y Barbuda"
            },
            "fin": {
                "official": "Antigua ja Barbuda",
                "common": "Antigua ja Barbuda"
            }
        },
        "latlng": [
            17.05,
            -61.8
        ],
        "demonym": "Antiguan, Barbudan",
        "landlocked": false,
        "borders": [],
        "area": 442
    },
    {
        "name": {
            "common": "Australia",
            "official": "Commonwealth of Australia",
            "native": {
                "eng": {
                    "official": "Commonwealth of Australia",
                    "common": "Australia"
                }
            }
        },
        "tld": [
            ".au"
        ],
        "cca2": "AU",
        "ccn3": "036",
        "cca3": "AUS",
        "cioc": "AUS",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "61"
        ],
        "capital": "Canberra",
        "altSpellings": [
            "AU"
        ],
        "region": "Oceania",
        "subregion": "Australia and New Zealand",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Commonwealth of Australia",
                "common": "Awstralia"
            },
            "deu": {
                "official": "Commonwealth Australien",
                "common": "Australien"
            },
            "fra": {
                "official": "Australie",
                "common": "Australie"
            },
            "hrv": {
                "official": "Commonwealth of Australia",
                "common": "Australija"
            },
            "ita": {
                "official": "Commonwealth dell'Australia",
                "common": "Australia"
            },
            "jpn": {
                "official": "オーストラリア連邦",
                "common": "オーストラリア"
            },
            "nld": {
                "official": "Gemenebest van Australië",
                "common": "Australië"
            },
            "por": {
                "official": "Comunidade da Austrália",
                "common": "Austrália"
            },
            "rus": {
                "official": "Содружество Австралии",
                "common": "Австралия"
            },
            "spa": {
                "official": "Mancomunidad de Australia",
                "common": "Australia"
            },
            "fin": {
                "official": "Australian liittovaltio",
                "common": "Australia"
            }
        },
        "latlng": [
            -27,
            133
        ],
        "demonym": "Australian",
        "landlocked": false,
        "borders": [],
        "area": 7692024
    },
    {
        "name": {
            "common": "Austria",
            "official": "Republic of Austria",
            "native": {
                "bar": {
                    "official": "Republik Österreich",
                    "common": "Österreich"
                }
            }
        },
        "tld": [
            ".at"
        ],
        "cca2": "AT",
        "ccn3": "040",
        "cca3": "AUT",
        "cioc": "AUT",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "43"
        ],
        "capital": "Vienna",
        "altSpellings": [
            "AT",
            "Osterreich",
            "Oesterreich"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "bar": "Austro-Bavarian German"
        },
        "translations": {
            "cym": {
                "official": "Republic of Austria",
                "common": "Awstria"
            },
            "deu": {
                "official": "Republik Österreich",
                "common": "Österreich"
            },
            "fra": {
                "official": "République d'Autriche",
                "common": "Autriche"
            },
            "hrv": {
                "official": "Republika Austrija",
                "common": "Austrija"
            },
            "ita": {
                "official": "Repubblica d'Austria",
                "common": "Austria"
            },
            "jpn": {
                "official": "オーストリア共和国",
                "common": "オーストリア"
            },
            "nld": {
                "official": "Republiek Oostenrijk",
                "common": "Oostenrijk"
            },
            "por": {
                "official": "República da Áustria",
                "common": "Áustria"
            },
            "rus": {
                "official": "Австрийская Республика",
                "common": "Австрия"
            },
            "spa": {
                "official": "República de Austria",
                "common": "Austria"
            },
            "fin": {
                "official": "Itävallan tasavalta",
                "common": "Itävalta"
            }
        },
        "latlng": [
            47.33333333,
            13.33333333
        ],
        "demonym": "Austrian",
        "landlocked": true,
        "borders": [
            "CZE",
            "DEU",
            "HUN",
            "ITA",
            "LIE",
            "SVK",
            "SVN",
            "CHE"
        ],
        "area": 83871
    },
    {
        "name": {
            "common": "Azerbaijan",
            "official": "Republic of Azerbaijan",
            "native": {
                "aze": {
                    "official": "Azərbaycan Respublikası",
                    "common": "Azərbaycan"
                },
                "rus": {
                    "official": "Азербайджанская Республика",
                    "common": "Азербайджан"
                }
            }
        },
        "tld": [
            ".az"
        ],
        "cca2": "AZ",
        "ccn3": "031",
        "cca3": "AZE",
        "cioc": "AZE",
        "currency": [
            "AZN"
        ],
        "callingCode": [
            "994"
        ],
        "capital": "Baku",
        "altSpellings": [
            "AZ",
            "Republic of Azerbaijan",
            "Azərbaycan Respublikası"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "aze": "Azerbaijani",
            "rus": "Russian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Azerbaijan",
                "common": "Aserbaijan"
            },
            "deu": {
                "official": "Republik Aserbaidschan",
                "common": "Aserbaidschan"
            },
            "fra": {
                "official": "République d'Azerbaïdjan",
                "common": "Azerbaïdjan"
            },
            "hrv": {
                "official": "Republika Azerbajdžan",
                "common": "Azerbajdžan"
            },
            "ita": {
                "official": "Repubblica dell'Azerbaigian",
                "common": "Azerbaijan"
            },
            "jpn": {
                "official": "アゼルバイジャン共和国",
                "common": "アゼルバイジャン"
            },
            "nld": {
                "official": "Republiek Azerbeidzjan",
                "common": "Azerbeidzjan"
            },
            "por": {
                "official": "República do Azerbaijão",
                "common": "Azerbeijão"
            },
            "rus": {
                "official": "Азербайджанская Республика",
                "common": "Азербайджан"
            },
            "spa": {
                "official": "República de Azerbaiyán",
                "common": "Azerbaiyán"
            },
            "fin": {
                "official": "Azerbaidzanin tasavalta",
                "common": "Azerbaidzan"
            }
        },
        "latlng": [
            40.5,
            47.5
        ],
        "demonym": "Azerbaijani",
        "landlocked": true,
        "borders": [
            "ARM",
            "GEO",
            "IRN",
            "RUS",
            "TUR"
        ],
        "area": 86600
    },
    {
        "name": {
            "common": "Burundi",
            "official": "Republic of Burundi",
            "native": {
                "fra": {
                    "official": "République du Burundi",
                    "common": "Burundi"
                },
                "run": {
                    "official": "Republika y'Uburundi ",
                    "common": "Uburundi"
                }
            }
        },
        "tld": [
            ".bi"
        ],
        "cca2": "BI",
        "ccn3": "108",
        "cca3": "BDI",
        "cioc": "BDI",
        "currency": [
            "BIF"
        ],
        "callingCode": [
            "257"
        ],
        "capital": "Bujumbura",
        "altSpellings": [
            "BI",
            "Republic of Burundi",
            "Republika y'Uburundi",
            "République du Burundi"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "fra": "French",
            "run": "Kirundi"
        },
        "translations": {
            "cym": {
                "official": "Republic of Burundi",
                "common": "Bwrwndi"
            },
            "deu": {
                "official": "Republik Burundi",
                "common": "Burundi"
            },
            "fra": {
                "official": "République du Burundi",
                "common": "Burundi"
            },
            "hrv": {
                "official": "Burundi",
                "common": "Burundi"
            },
            "ita": {
                "official": "Repubblica del Burundi",
                "common": "Burundi"
            },
            "jpn": {
                "official": "ブルンジ共和国",
                "common": "ブルンジ"
            },
            "nld": {
                "official": "Republiek Burundi",
                "common": "Burundi"
            },
            "por": {
                "official": "República do Burundi",
                "common": "Burundi"
            },
            "rus": {
                "official": "Республика Бурунди",
                "common": "Бурунди"
            },
            "spa": {
                "official": "República de Burundi",
                "common": "Burundi"
            },
            "fin": {
                "official": "Burundin tasavalta",
                "common": "Burundi"
            }
        },
        "latlng": [
            -3.5,
            30
        ],
        "demonym": "Burundian",
        "landlocked": true,
        "borders": [
            "COD",
            "RWA",
            "TZA"
        ],
        "area": 27834
    },
    {
        "name": {
            "common": "Belgium",
            "official": "Kingdom of Belgium",
            "native": {
                "deu": {
                    "official": "Königreich Belgien",
                    "common": "Belgien"
                },
                "fra": {
                    "official": "Royaume de Belgique",
                    "common": "Belgique"
                },
                "nld": {
                    "official": "Koninkrijk België",
                    "common": "België"
                }
            }
        },
        "tld": [
            ".be"
        ],
        "cca2": "BE",
        "ccn3": "056",
        "cca3": "BEL",
        "cioc": "BEL",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "32"
        ],
        "capital": "Brussels",
        "altSpellings": [
            "BE",
            "België",
            "Belgie",
            "Belgien",
            "Belgique",
            "Kingdom of Belgium",
            "Koninkrijk België",
            "Royaume de Belgique",
            "Königreich Belgien"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "deu": "German",
            "fra": "French",
            "nld": "Dutch"
        },
        "translations": {
            "cym": {
                "official": "Kingdom of Belgium",
                "common": "Gwlad Belg"
            },
            "deu": {
                "official": "Königreich Belgien",
                "common": "Belgien"
            },
            "fra": {
                "official": "Royaume de Belgique",
                "common": "Belgique"
            },
            "hrv": {
                "official": "Kraljevina Belgija",
                "common": "Belgija"
            },
            "ita": {
                "official": "Regno del Belgio",
                "common": "Belgio"
            },
            "jpn": {
                "official": "ベルギー王国",
                "common": "ベルギー"
            },
            "nld": {
                "official": "Koninkrijk België",
                "common": "België"
            },
            "por": {
                "official": "Reino da Bélgica",
                "common": "Bélgica"
            },
            "rus": {
                "official": "Королевство Бельгия",
                "common": "Бельгия"
            },
            "spa": {
                "official": "Reino de Bélgica",
                "common": "Bélgica"
            },
            "fin": {
                "official": "Belgian kuningaskunta",
                "common": "Belgia"
            }
        },
        "latlng": [
            50.83333333,
            4
        ],
        "demonym": "Belgian",
        "landlocked": false,
        "borders": [
            "FRA",
            "DEU",
            "LUX",
            "NLD"
        ],
        "area": 30528
    },
    {
        "name": {
            "common": "Benin",
            "official": "Republic of Benin",
            "native": {
                "fra": {
                    "official": "République du Bénin",
                    "common": "Bénin"
                }
            }
        },
        "tld": [
            ".bj"
        ],
        "cca2": "BJ",
        "ccn3": "204",
        "cca3": "BEN",
        "cioc": "BEN",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "229"
        ],
        "capital": "Porto-Novo",
        "altSpellings": [
            "BJ",
            "Republic of Benin",
            "République du Bénin"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Republic of Benin",
                "common": "Benin"
            },
            "deu": {
                "official": "Republik Benin",
                "common": "Benin"
            },
            "fra": {
                "official": "République du Bénin",
                "common": "Bénin"
            },
            "hrv": {
                "official": "Republika Benin",
                "common": "Benin"
            },
            "ita": {
                "official": "Repubblica del Benin",
                "common": "Benin"
            },
            "jpn": {
                "official": "ベナン共和国",
                "common": "ベナン"
            },
            "nld": {
                "official": "Republiek Benin",
                "common": "Benin"
            },
            "por": {
                "official": "República do Benin",
                "common": "Benin"
            },
            "rus": {
                "official": "Республика Бенин",
                "common": "Бенин"
            },
            "spa": {
                "official": "República de Benin",
                "common": "Benín"
            },
            "fin": {
                "official": "Beninin tasavalta",
                "common": "Benin"
            }
        },
        "latlng": [
            9.5,
            2.25
        ],
        "demonym": "Beninese",
        "landlocked": false,
        "borders": [
            "BFA",
            "NER",
            "NGA",
            "TGO"
        ],
        "area": 112622
    },
    {
        "name": {
            "common": "Burkina Faso",
            "official": "Burkina Faso",
            "native": {
                "fra": {
                    "official": "République du Burkina",
                    "common": "Burkina Faso"
                }
            }
        },
        "tld": [
            ".bf"
        ],
        "cca2": "BF",
        "ccn3": "854",
        "cca3": "BFA",
        "cioc": "BUR",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "226"
        ],
        "capital": "Ouagadougou",
        "altSpellings": [
            "BF"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "deu": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "fra": {
                "official": "République du Burkina",
                "common": "Burkina Faso"
            },
            "hrv": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "ita": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "jpn": {
                "official": "ブルキナファソ",
                "common": "ブルキナファソ"
            },
            "nld": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "por": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "rus": {
                "official": "Буркина -Фасо",
                "common": "Буркина-Фасо"
            },
            "spa": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            },
            "fin": {
                "official": "Burkina Faso",
                "common": "Burkina Faso"
            }
        },
        "latlng": [
            13,
            -2
        ],
        "demonym": "Burkinabe",
        "landlocked": true,
        "borders": [
            "BEN",
            "CIV",
            "GHA",
            "MLI",
            "NER",
            "TGO"
        ],
        "area": 272967
    },
    {
        "name": {
            "common": "Bangladesh",
            "official": "People's Republic of Bangladesh",
            "native": {
                "ben": {
                    "official": "বাংলাদেশ গণপ্রজাতন্ত্রী",
                    "common": "বাংলাদেশ"
                }
            }
        },
        "tld": [
            ".bd"
        ],
        "cca2": "BD",
        "ccn3": "050",
        "cca3": "BGD",
        "cioc": "BAN",
        "currency": [
            "BDT"
        ],
        "callingCode": [
            "880"
        ],
        "capital": "Dhaka",
        "altSpellings": [
            "BD",
            "People's Republic of Bangladesh",
            "Gônôprôjatôntri Bangladesh"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "ben": "Bengali"
        },
        "translations": {
            "cym": {
                "official": "People's Republic of Bangladesh",
                "common": "Bangladesh"
            },
            "deu": {
                "official": "Volksrepublik Bangladesch",
                "common": "Bangladesch"
            },
            "fra": {
                "official": "La République populaire du Bangladesh",
                "common": "Bangladesh"
            },
            "hrv": {
                "official": "Narodna Republika Bangladeš",
                "common": "Bangladeš"
            },
            "ita": {
                "official": "Repubblica popolare del Bangladesh",
                "common": "Bangladesh"
            },
            "jpn": {
                "official": "バングラデシュ人民共和国",
                "common": "バングラデシュ"
            },
            "nld": {
                "official": "Volksrepubliek Bangladesh",
                "common": "Bangladesh"
            },
            "por": {
                "official": "República Popular do Bangladesh",
                "common": "Bangladesh"
            },
            "rus": {
                "official": "Народная Республика Бангладеш",
                "common": "Бангладеш"
            },
            "spa": {
                "official": "República Popular de Bangladesh",
                "common": "Bangladesh"
            },
            "fin": {
                "official": "Bangladeshin kansantasavalta",
                "common": "Bangladesh"
            }
        },
        "latlng": [
            24,
            90
        ],
        "demonym": "Bangladeshi",
        "landlocked": false,
        "borders": [
            "MMR",
            "IND"
        ],
        "area": 147570
    },
    {
        "name": {
            "common": "Bulgaria",
            "official": "Republic of Bulgaria",
            "native": {
                "bul": {
                    "official": "Република България",
                    "common": "България"
                }
            }
        },
        "tld": [
            ".bg"
        ],
        "cca2": "BG",
        "ccn3": "100",
        "cca3": "BGR",
        "cioc": "BUL",
        "currency": [
            "BGN"
        ],
        "callingCode": [
            "359"
        ],
        "capital": "Sofia",
        "altSpellings": [
            "BG",
            "Republic of Bulgaria",
            "Република България"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "bul": "Bulgarian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Bulgaria",
                "common": "Bwlgaria"
            },
            "deu": {
                "official": "Republik Bulgarien",
                "common": "Bulgarien"
            },
            "fra": {
                "official": "République de Bulgarie",
                "common": "Bulgarie"
            },
            "hrv": {
                "official": "Republika Bugarska",
                "common": "Bugarska"
            },
            "ita": {
                "official": "Repubblica di Bulgaria",
                "common": "Bulgaria"
            },
            "jpn": {
                "official": "ブルガリア共和国",
                "common": "ブルガリア"
            },
            "nld": {
                "official": "Republiek Bulgarije",
                "common": "Bulgarije"
            },
            "por": {
                "official": "República da Bulgária",
                "common": "Bulgária"
            },
            "rus": {
                "official": "Республика Болгария",
                "common": "Болгария"
            },
            "spa": {
                "official": "República de Bulgaria",
                "common": "Bulgaria"
            },
            "fin": {
                "official": "Bulgarian tasavalta",
                "common": "Bulgaria"
            }
        },
        "latlng": [
            43,
            25
        ],
        "demonym": "Bulgarian",
        "landlocked": false,
        "borders": [
            "GRC",
            "MKD",
            "ROU",
            "SRB",
            "TUR"
        ],
        "area": 110879
    },
    {
        "name": {
            "common": "Bahrain",
            "official": "Kingdom of Bahrain",
            "native": {
                "ara": {
                    "official": "مملكة البحرين",
                    "common": "‏البحرين"
                }
            }
        },
        "tld": [
            ".bh"
        ],
        "cca2": "BH",
        "ccn3": "048",
        "cca3": "BHR",
        "cioc": "BRN",
        "currency": [
            "BHD"
        ],
        "callingCode": [
            "973"
        ],
        "capital": "Manama",
        "altSpellings": [
            "BH",
            "Kingdom of Bahrain",
            "Mamlakat al-Baḥrayn"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "cym": {
                "official": "Kingdom of Bahrain",
                "common": "Bahrain"
            },
            "deu": {
                "official": "Königreich Bahrain",
                "common": "Bahrain"
            },
            "fra": {
                "official": "Royaume de Bahreïn",
                "common": "Bahreïn"
            },
            "hrv": {
                "official": "Kraljevina Bahrein",
                "common": "Bahrein"
            },
            "ita": {
                "official": "Regno del Bahrain",
                "common": "Bahrein"
            },
            "jpn": {
                "official": "バーレーン王国",
                "common": "バーレーン"
            },
            "nld": {
                "official": "Koninkrijk Bahrein",
                "common": "Bahrein"
            },
            "por": {
                "official": "Reino do Bahrein",
                "common": "Bahrein"
            },
            "rus": {
                "official": "Королевство Бахрейн",
                "common": "Бахрейн"
            },
            "spa": {
                "official": "Reino de Bahrein",
                "common": "Bahrein"
            },
            "fin": {
                "official": "Bahrainin kuningaskunta",
                "common": "Bahrain"
            }
        },
        "latlng": [
            26,
            50.55
        ],
        "demonym": "Bahraini",
        "landlocked": false,
        "borders": [],
        "area": 765
    },
    {
        "name": {
            "common": "Bahamas",
            "official": "Commonwealth of the Bahamas",
            "native": {
                "eng": {
                    "official": "Commonwealth of the Bahamas",
                    "common": "Bahamas"
                }
            }
        },
        "tld": [
            ".bs"
        ],
        "cca2": "BS",
        "ccn3": "044",
        "cca3": "BHS",
        "cioc": "BAH",
        "currency": [
            "BSD"
        ],
        "callingCode": [
            "1242"
        ],
        "capital": "Nassau",
        "altSpellings": [
            "BS",
            "Commonwealth of the Bahamas"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Commonwealth of the Bahamas",
                "common": "Bahamas"
            },
            "deu": {
                "official": "Commonwealth der Bahamas",
                "common": "Bahamas"
            },
            "fra": {
                "official": "Commonwealth des Bahamas",
                "common": "Bahamas"
            },
            "hrv": {
                "official": "Zajednica Bahama",
                "common": "Bahami"
            },
            "ita": {
                "official": "Commonwealth delle Bahamas",
                "common": "Bahamas"
            },
            "jpn": {
                "official": "バハマ",
                "common": "バハマ"
            },
            "nld": {
                "official": "Gemenebest van de Bahama's",
                "common": "Bahama’s"
            },
            "por": {
                "official": "Comunidade das Bahamas",
                "common": "Bahamas"
            },
            "rus": {
                "official": "Содружество Багамских Островов",
                "common": "Багамские Острова"
            },
            "spa": {
                "official": "Commonwealth de las Bahamas",
                "common": "Bahamas"
            },
            "fin": {
                "official": "Bahaman liittovaltio",
                "common": "Bahamasaaret"
            }
        },
        "latlng": [
            24.25,
            -76
        ],
        "demonym": "Bahamian",
        "landlocked": false,
        "borders": [],
        "area": 13943
    },
    {
        "name": {
            "common": "Bosnia and Herzegovina",
            "official": "Bosnia and Herzegovina",
            "native": {
                "bos": {
                    "official": "Bosna i Hercegovina",
                    "common": "Bosna i Hercegovina"
                },
                "hrv": {
                    "official": "Bosna i Hercegovina",
                    "common": "Bosna i Hercegovina"
                },
                "srp": {
                    "official": "Боснa и Херцеговина",
                    "common": "Боснa и Херцеговина"
                }
            }
        },
        "tld": [
            ".ba"
        ],
        "cca2": "BA",
        "ccn3": "070",
        "cca3": "BIH",
        "cioc": "BIH",
        "currency": [
            "BAM"
        ],
        "callingCode": [
            "387"
        ],
        "capital": "Sarajevo",
        "altSpellings": [
            "BA",
            "Bosnia-Herzegovina",
            "Босна и Херцеговина"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "bos": "Bosnian",
            "hrv": "Croatian",
            "srp": "Serbian"
        },
        "translations": {
            "cym": {
                "official": "Bosnia and Herzegovina",
                "common": "Bosnia a Hercegovina"
            },
            "deu": {
                "official": "Bosnien und Herzegowina",
                "common": "Bosnien und Herzegowina"
            },
            "fra": {
                "official": "Bosnie-et-Herzégovine",
                "common": "Bosnie-Herzégovine"
            },
            "hrv": {
                "official": "Bosna i Hercegovina",
                "common": "Bosna i Hercegovina"
            },
            "ita": {
                "official": "Bosnia-Erzegovina",
                "common": "Bosnia ed Erzegovina"
            },
            "jpn": {
                "official": "ボスニア·ヘルツェゴビナ",
                "common": "ボスニア・ヘルツェゴビナ"
            },
            "nld": {
                "official": "Bosnië-Herzegovina",
                "common": "Bosnië en Herzegovina"
            },
            "por": {
                "official": "Bósnia e Herzegovina",
                "common": "Bósnia e Herzegovina"
            },
            "rus": {
                "official": "Босния и Герцеговина",
                "common": "Босния и Герцеговина"
            },
            "spa": {
                "official": "Bosnia y Herzegovina",
                "common": "Bosnia y Herzegovina"
            },
            "fin": {
                "official": "Bosnia ja Hertsegovina",
                "common": "Bosnia ja Hertsegovina"
            }
        },
        "latlng": [
            44,
            18
        ],
        "demonym": "Bosnian, Herzegovinian",
        "landlocked": false,
        "borders": [
            "HRV",
            "MNE",
            "SRB"
        ],
        "area": 51209
    },
    {
        "name": {
            "common": "Saint Barthélemy",
            "official": "Collectivity of Saint Barthélemy",
            "native": {
                "fra": {
                    "official": "Collectivité de Saint-Barthélemy",
                    "common": "Saint-Barthélemy"
                }
            }
        },
        "tld": [
            ".bl"
        ],
        "cca2": "BL",
        "ccn3": "652",
        "cca3": "BLM",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "590"
        ],
        "capital": "Gustavia",
        "altSpellings": [
            "BL",
            "St. Barthelemy",
            "Collectivity of Saint Barthélemy",
            "Collectivité de Saint-Barthélemy"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Gebietskörperschaft Saint -Barthélemy",
                "common": "Saint-Barthélemy"
            },
            "fra": {
                "official": "Collectivité de Saint-Barthélemy",
                "common": "Saint-Barthélemy"
            },
            "hrv": {
                "official": "Kolektivnost sv Barthélemy",
                "common": "Saint Barthélemy"
            },
            "ita": {
                "official": "Collettività di Saint Barthélemy",
                "common": "Antille Francesi"
            },
            "jpn": {
                "official": "サン·バルテルミー島の集合体",
                "common": "サン・バルテルミー"
            },
            "nld": {
                "official": "Gemeenschap Saint Barthélemy",
                "common": "Saint Barthélemy"
            },
            "por": {
                "official": "Coletividade de Saint Barthélemy",
                "common": "São Bartolomeu"
            },
            "rus": {
                "official": "Коллективность Санкт -Бартельми",
                "common": "Сен-Бартелеми"
            },
            "spa": {
                "official": "Colectividad de San Barthélemy",
                "common": "San Bartolomé"
            },
            "fin": {
                "official": "Saint-Barthélemyn yhteisö",
                "common": "Saint-Barthélemy"
            }
        },
        "latlng": [
            18.5,
            -63.41666666
        ],
        "demonym": "Saint Barthélemy Islander",
        "landlocked": false,
        "borders": [],
        "area": 21
    },
    {
        "name": {
            "common": "Belarus",
            "official": "Republic of Belarus",
            "native": {
                "bel": {
                    "official": "Рэспубліка Беларусь",
                    "common": "Белару́сь"
                },
                "rus": {
                    "official": "Республика Беларусь",
                    "common": "Белоруссия"
                }
            }
        },
        "tld": [
            ".by"
        ],
        "cca2": "BY",
        "ccn3": "112",
        "cca3": "BLR",
        "cioc": "BLR",
        "currency": [
            "BYR"
        ],
        "callingCode": [
            "375"
        ],
        "capital": "Minsk",
        "altSpellings": [
            "BY",
            "Bielaruś",
            "Republic of Belarus",
            "Белоруссия",
            "Республика Беларусь",
            "Belorussiya",
            "Respublika Belarus’"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "bel": "Belarusian",
            "rus": "Russian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Belarus",
                "common": "Belarws"
            },
            "deu": {
                "official": "Republik Belarus",
                "common": "Weißrussland"
            },
            "fra": {
                "official": "République de Biélorussie",
                "common": "Biélorussie"
            },
            "hrv": {
                "official": "Republika Bjelorusija",
                "common": "Bjelorusija"
            },
            "ita": {
                "official": "Repubblica di Belarus",
                "common": "Bielorussia"
            },
            "jpn": {
                "official": "ベラルーシ共和国",
                "common": "ベラルーシ"
            },
            "nld": {
                "official": "Republiek Belarus",
                "common": "Wit-Rusland"
            },
            "por": {
                "official": "República da Bielorrússia",
                "common": "Bielorússia"
            },
            "rus": {
                "official": "Республика Беларусь",
                "common": "Белоруссия"
            },
            "spa": {
                "official": "República de Belarús",
                "common": "Bielorrusia"
            },
            "fin": {
                "official": "Valko-Venäjän tasavalta",
                "common": "Valko-Venäjä"
            }
        },
        "latlng": [
            53,
            28
        ],
        "demonym": "Belarusian",
        "landlocked": true,
        "borders": [
            "LVA",
            "LTU",
            "POL",
            "RUS",
            "UKR"
        ],
        "area": 207600
    },
    {
        "name": {
            "common": "Belize",
            "official": "Belize",
            "native": {
                "bjz": {
                    "official": "Belize",
                    "common": "Belize"
                },
                "eng": {
                    "official": "Belize",
                    "common": "Belize"
                },
                "spa": {
                    "official": "Belice",
                    "common": "Belice"
                }
            }
        },
        "tld": [
            ".bz"
        ],
        "cca2": "BZ",
        "ccn3": "084",
        "cca3": "BLZ",
        "cioc": "BIZ",
        "currency": [
            "BZD"
        ],
        "callingCode": [
            "501"
        ],
        "capital": "Belmopan",
        "altSpellings": [
            "BZ"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "bjz": "Belizean Creole",
            "eng": "English",
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Belize",
                "common": "Belize"
            },
            "deu": {
                "official": "Belize",
                "common": "Belize"
            },
            "fra": {
                "official": "Belize",
                "common": "Belize"
            },
            "hrv": {
                "official": "Belize",
                "common": "Belize"
            },
            "ita": {
                "official": "Belize",
                "common": "Belize"
            },
            "jpn": {
                "official": "ベリーズ",
                "common": "ベリーズ"
            },
            "nld": {
                "official": "Belize",
                "common": "Belize"
            },
            "por": {
                "official": "Belize",
                "common": "Belize"
            },
            "rus": {
                "official": "Белиз",
                "common": "Белиз"
            },
            "spa": {
                "official": "Belice",
                "common": "Belice"
            },
            "fin": {
                "official": "Belize",
                "common": "Belize"
            }
        },
        "latlng": [
            17.25,
            -88.75
        ],
        "demonym": "Belizean",
        "landlocked": false,
        "borders": [
            "GTM",
            "MEX"
        ],
        "area": 22966
    },
    {
        "name": {
            "common": "Bermuda",
            "official": "Bermuda",
            "native": {
                "eng": {
                    "official": "Bermuda",
                    "common": "Bermuda"
                }
            }
        },
        "tld": [
            ".bm"
        ],
        "cca2": "BM",
        "ccn3": "060",
        "cca3": "BMU",
        "cioc": "BER",
        "currency": [
            "BMD"
        ],
        "callingCode": [
            "1441"
        ],
        "capital": "Hamilton",
        "altSpellings": [
            "BM",
            "The Islands of Bermuda",
            "The Bermudas",
            "Somers Isles"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Bermuda",
                "common": "Bermiwda"
            },
            "deu": {
                "official": "Bermuda",
                "common": "Bermuda"
            },
            "fra": {
                "official": "Bermudes",
                "common": "Bermudes"
            },
            "hrv": {
                "official": "Bermuda",
                "common": "Bermudi"
            },
            "ita": {
                "official": "Bermuda",
                "common": "Bermuda"
            },
            "jpn": {
                "official": "バミューダ",
                "common": "バミューダ"
            },
            "nld": {
                "official": "Bermuda",
                "common": "Bermuda"
            },
            "por": {
                "official": "Bermudas",
                "common": "Bermudas"
            },
            "rus": {
                "official": "Бермудские острова",
                "common": "Бермудские Острова"
            },
            "spa": {
                "official": "Bermuda",
                "common": "Bermudas"
            },
            "fin": {
                "official": "Bermuda",
                "common": "Bermuda"
            }
        },
        "latlng": [
            32.33333333,
            -64.75
        ],
        "demonym": "Bermudian",
        "landlocked": false,
        "borders": [],
        "area": 54
    },
    {
        "name": {
            "common": "Bolivia",
            "official": "Plurinational State of Bolivia",
            "native": {
                "aym": {
                    "official": "Wuliwya Suyu",
                    "common": "Wuliwya"
                },
                "grn": {
                    "official": "Tetã Volívia",
                    "common": "Volívia"
                },
                "que": {
                    "official": "Buliwya Mamallaqta",
                    "common": "Buliwya"
                },
                "spa": {
                    "official": "Estado Plurinacional de Bolivia",
                    "common": "Bolivia"
                }
            }
        },
        "tld": [
            ".bo"
        ],
        "cca2": "BO",
        "ccn3": "068",
        "cca3": "BOL",
        "cioc": "BOL",
        "currency": [
            "BOB",
            "BOV"
        ],
        "callingCode": [
            "591"
        ],
        "capital": "Sucre",
        "altSpellings": [
            "BO",
            "Buliwya",
            "Wuliwya",
            "Bolivia, Plurinational State of",
            "Plurinational State of Bolivia",
            "Estado Plurinacional de Bolivia",
            "Buliwya Mamallaqta",
            "Wuliwya Suyu",
            "Tetã Volívia"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "aym": "Aymara",
            "grn": "Guaraní",
            "que": "Quechua",
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Plurinational State of Bolivia",
                "common": "Bolifia"
            },
            "deu": {
                "official": "Multinationaler Staat von Bolivien",
                "common": "Bolivien"
            },
            "fra": {
                "official": "État plurinational de Bolivie",
                "common": "Bolivie"
            },
            "hrv": {
                "official": "Plurinational State of Bolivia",
                "common": "Bolivija"
            },
            "ita": {
                "official": "Stato Plurinazionale della Bolivia",
                "common": "Bolivia"
            },
            "jpn": {
                "official": "ボリビアの多民族国",
                "common": "ボリビア多民族国"
            },
            "nld": {
                "official": "Plurinationale Staat van Bolivia",
                "common": "Bolivia"
            },
            "por": {
                "official": "Estado Plurinacional da Bolívia",
                "common": "Bolívia"
            },
            "rus": {
                "official": "Многонациональное Государство Боливия",
                "common": "Боливия"
            },
            "spa": {
                "official": "Estado Plurinacional de Bolivia",
                "common": "Bolivia"
            },
            "fin": {
                "official": "Bolivian monikansainen valtio",
                "common": "Bolivia"
            }
        },
        "latlng": [
            -17,
            -65
        ],
        "demonym": "Bolivian",
        "landlocked": true,
        "borders": [
            "ARG",
            "BRA",
            "CHL",
            "PRY",
            "PER"
        ],
        "area": 1098581
    },
    {
        "name": {
            "common": "Brazil",
            "official": "Federative Republic of Brazil",
            "native": {
                "por": {
                    "official": "República Federativa do Brasil",
                    "common": "Brasil"
                }
            }
        },
        "tld": [
            ".br"
        ],
        "cca2": "BR",
        "ccn3": "076",
        "cca3": "BRA",
        "cioc": "BRA",
        "currency": [
            "BRL"
        ],
        "callingCode": [
            "55"
        ],
        "capital": "Brasília",
        "altSpellings": [
            "BR",
            "Brasil",
            "Federative Republic of Brazil",
            "República Federativa do Brasil"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "cym": {
                "official": "Federative Republic of Brazil",
                "common": "Brasil"
            },
            "deu": {
                "official": "Föderative Republik Brasilien",
                "common": "Brasilien"
            },
            "fra": {
                "official": "République fédérative du Brésil",
                "common": "Brésil"
            },
            "hrv": {
                "official": "Savezne Republike Brazil",
                "common": "Brazil"
            },
            "ita": {
                "official": "Repubblica federativa del Brasile",
                "common": "Brasile"
            },
            "jpn": {
                "official": "ブラジル連邦共和国",
                "common": "ブラジル"
            },
            "nld": {
                "official": "Federale Republiek Brazilië",
                "common": "Brazilië"
            },
            "por": {
                "official": "República Federativa do Brasil",
                "common": "Brasil"
            },
            "rus": {
                "official": "Федеративная Республика Бразилия",
                "common": "Бразилия"
            },
            "spa": {
                "official": "República Federativa del Brasil",
                "common": "Brasil"
            },
            "fin": {
                "official": "Brasilian liittotasavalta",
                "common": "Brasilia"
            }
        },
        "latlng": [
            -10,
            -55
        ],
        "demonym": "Brazilian",
        "landlocked": false,
        "borders": [
            "ARG",
            "BOL",
            "COL",
            "GUF",
            "GUY",
            "PRY",
            "PER",
            "SUR",
            "URY",
            "VEN"
        ],
        "area": 8515767
    },
    {
        "name": {
            "common": "Barbados",
            "official": "Barbados",
            "native": {
                "eng": {
                    "official": "Barbados",
                    "common": "Barbados"
                }
            }
        },
        "tld": [
            ".bb"
        ],
        "cca2": "BB",
        "ccn3": "052",
        "cca3": "BRB",
        "cioc": "BAR",
        "currency": [
            "BBD"
        ],
        "callingCode": [
            "1246"
        ],
        "capital": "Bridgetown",
        "altSpellings": [
            "BB"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "deu": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "fra": {
                "official": "Barbade",
                "common": "Barbade"
            },
            "hrv": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "ita": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "jpn": {
                "official": "バルバドス",
                "common": "バルバドス"
            },
            "nld": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "por": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "rus": {
                "official": "Барбадос",
                "common": "Барбадос"
            },
            "spa": {
                "official": "Barbados",
                "common": "Barbados"
            },
            "fin": {
                "official": "Barbados",
                "common": "Barbados"
            }
        },
        "latlng": [
            13.16666666,
            -59.53333333
        ],
        "demonym": "Barbadian",
        "landlocked": false,
        "borders": [],
        "area": 430
    },
    {
        "name": {
            "common": "Brunei",
            "official": "Nation of Brunei, Abode of Peace",
            "native": {
                "msa": {
                    "official": "Nation of Brunei, Abode Damai",
                    "common": "Negara Brunei Darussalam"
                }
            }
        },
        "tld": [
            ".bn"
        ],
        "cca2": "BN",
        "ccn3": "096",
        "cca3": "BRN",
        "cioc": "BRU",
        "currency": [
            "BND"
        ],
        "callingCode": [
            "673"
        ],
        "capital": "Bandar Seri Begawan",
        "altSpellings": [
            "BN",
            "Brunei Darussalam",
            "Nation of Brunei",
            "the Abode of Peace"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "msa": "Malay"
        },
        "translations": {
            "cym": {
                "official": "Nation of Brunei, Abode of Peace",
                "common": "Brunei"
            },
            "deu": {
                "official": "Nation von Brunei, Wohnung des Friedens",
                "common": "Brunei"
            },
            "fra": {
                "official": "État de Brunei Darussalam",
                "common": "Brunei"
            },
            "hrv": {
                "official": "Nacija od Bruneja, Kuću Mira",
                "common": "Brunej"
            },
            "ita": {
                "official": "Nazione di Brunei, Dimora della Pace",
                "common": "Brunei"
            },
            "jpn": {
                "official": "ブルネイ、平和の精舎の国家",
                "common": "ブルネイ・ダルサラーム"
            },
            "nld": {
                "official": "Natie van Brunei, de verblijfplaats van de Vrede",
                "common": "Brunei"
            },
            "por": {
                "official": "Nação do Brunei, Morada da Paz",
                "common": "Brunei"
            },
            "rus": {
                "official": "Нация Бруней, обитель мира",
                "common": "Бруней"
            },
            "spa": {
                "official": "Nación de Brunei, Morada de la Paz",
                "common": "Brunei"
            },
            "fin": {
                "official": "Brunei Darussalamin valtio",
                "common": "Brunei"
            }
        },
        "latlng": [
            4.5,
            114.66666666
        ],
        "demonym": "Bruneian",
        "landlocked": false,
        "borders": [
            "MYS"
        ],
        "area": 5765
    },
    {
        "name": {
            "common": "Bhutan",
            "official": "Kingdom of Bhutan",
            "native": {
                "dzo": {
                    "official": "འབྲུག་རྒྱལ་ཁབ་",
                    "common": "འབྲུག་ཡུལ་"
                }
            }
        },
        "tld": [
            ".bt"
        ],
        "cca2": "BT",
        "ccn3": "064",
        "cca3": "BTN",
        "cioc": "BHU",
        "currency": [
            "BTN",
            "INR"
        ],
        "callingCode": [
            "975"
        ],
        "capital": "Thimphu",
        "altSpellings": [
            "BT",
            "Kingdom of Bhutan"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "dzo": "Dzongkha"
        },
        "translations": {
            "cym": {
                "official": "Kingdom of Bhutan",
                "common": "Bhwtan"
            },
            "deu": {
                "official": "Königreich Bhutan",
                "common": "Bhutan"
            },
            "fra": {
                "official": "Royaume du Bhoutan",
                "common": "Bhoutan"
            },
            "hrv": {
                "official": "Kraljevina Butan",
                "common": "Butan"
            },
            "ita": {
                "official": "Regno del Bhutan",
                "common": "Bhutan"
            },
            "jpn": {
                "official": "ブータン王国",
                "common": "ブータン"
            },
            "nld": {
                "official": "Koninkrijk Bhutan",
                "common": "Bhutan"
            },
            "por": {
                "official": "Reino do Butão",
                "common": "Butão"
            },
            "rus": {
                "official": "Королевство Бутан",
                "common": "Бутан"
            },
            "spa": {
                "official": "Reino de Bután",
                "common": "Bután"
            },
            "fin": {
                "official": "Bhutanin kuningaskunta",
                "common": "Bhutan"
            }
        },
        "latlng": [
            27.5,
            90.5
        ],
        "demonym": "Bhutanese",
        "landlocked": true,
        "borders": [
            "CHN",
            "IND"
        ],
        "area": 38394
    },
    {
        "name": {
            "common": "Bouvet Island",
            "official": "Bouvet Island",
            "native": {
                "nor": {
                    "official": "Bouvetøya",
                    "common": "Bouvetøya"
                }
            }
        },
        "tld": [
            ".bv"
        ],
        "cca2": "BV",
        "ccn3": "074",
        "cca3": "BVT",
        "cioc": "",
        "currency": [
            "NOK"
        ],
        "callingCode": [],
        "capital": "",
        "altSpellings": [
            "BV",
            "Bouvetøya",
            "Bouvet-øya"
        ],
        "region": "",
        "subregion": "",
        "languages": {
            "nor": "Norwegian"
        },
        "translations": {
            "deu": {
                "official": "Bouvet-Insel",
                "common": "Bouvetinsel"
            },
            "fra": {
                "official": "Île Bouvet",
                "common": "Île Bouvet"
            },
            "hrv": {
                "official": "Bouvet Island",
                "common": "Otok Bouvet"
            },
            "ita": {
                "official": "Isola Bouvet",
                "common": "Isola Bouvet"
            },
            "jpn": {
                "official": "ブーヴェ島",
                "common": "ブーベ島"
            },
            "nld": {
                "official": "Bouvet Island",
                "common": "Bouveteiland"
            },
            "por": {
                "official": "Ilha Bouvet",
                "common": "Ilha Bouvet"
            },
            "rus": {
                "official": "Остров Буве",
                "common": "Остров Буве"
            },
            "spa": {
                "official": "Isla Bouvet",
                "common": "Isla Bouvet"
            },
            "fin": {
                "official": "Bouvet'nsaari",
                "common": "Bouvet'nsaari"
            }
        },
        "latlng": [
            -54.43333333,
            3.4
        ],
        "demonym": "",
        "landlocked": false,
        "borders": [],
        "area": 49
    },
    {
        "name": {
            "common": "Botswana",
            "official": "Republic of Botswana",
            "native": {
                "eng": {
                    "official": "Republic of Botswana",
                    "common": "Botswana"
                },
                "tsn": {
                    "official": "Lefatshe la Botswana",
                    "common": "Botswana"
                }
            }
        },
        "tld": [
            ".bw"
        ],
        "cca2": "BW",
        "ccn3": "072",
        "cca3": "BWA",
        "cioc": "BOT",
        "currency": [
            "BWP"
        ],
        "callingCode": [
            "267"
        ],
        "capital": "Gaborone",
        "altSpellings": [
            "BW",
            "Republic of Botswana",
            "Lefatshe la Botswana"
        ],
        "region": "Africa",
        "subregion": "Southern Africa",
        "languages": {
            "eng": "English",
            "tsn": "Tswana"
        },
        "translations": {
            "deu": {
                "official": "Republik Botsuana",
                "common": "Botswana"
            },
            "fra": {
                "official": "République du Botswana",
                "common": "Botswana"
            },
            "hrv": {
                "official": "Republika Bocvana",
                "common": "Bocvana"
            },
            "ita": {
                "official": "Repubblica del Botswana",
                "common": "Botswana"
            },
            "jpn": {
                "official": "ボツワナ共和国",
                "common": "ボツワナ"
            },
            "nld": {
                "official": "Republiek Botswana",
                "common": "Botswana"
            },
            "por": {
                "official": "República do Botswana",
                "common": "Botswana"
            },
            "rus": {
                "official": "Республика Ботсвана",
                "common": "Ботсвана"
            },
            "spa": {
                "official": "República de Botswana",
                "common": "Botswana"
            },
            "fin": {
                "official": "Botswanan tasavalta",
                "common": "Botswana"
            }
        },
        "latlng": [
            -22,
            24
        ],
        "demonym": "Motswana",
        "landlocked": true,
        "borders": [
            "NAM",
            "ZAF",
            "ZMB",
            "ZWE"
        ],
        "area": 582000
    },
    {
        "name": {
            "common": "Central African Republic",
            "official": "Central African Republic",
            "native": {
                "fra": {
                    "official": "République centrafricaine",
                    "common": "République centrafricaine"
                },
                "sag": {
                    "official": "Ködörösêse tî Bêafrîka",
                    "common": "Bêafrîka"
                }
            }
        },
        "tld": [
            ".cf"
        ],
        "cca2": "CF",
        "ccn3": "140",
        "cca3": "CAF",
        "cioc": "CAF",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "236"
        ],
        "capital": "Bangui",
        "altSpellings": [
            "CF",
            "Central African Republic",
            "République centrafricaine"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "fra": "French",
            "sag": "Sango"
        },
        "translations": {
            "cym": {
                "official": "Central African Republic",
                "common": "Gweriniaeth Canolbarth Affrica"
            },
            "deu": {
                "official": "Zentralafrikanische Republik",
                "common": "Zentralafrikanische Republik"
            },
            "fra": {
                "official": "République centrafricaine",
                "common": "République centrafricaine"
            },
            "hrv": {
                "official": "Centralna Afrička Republika",
                "common": "Srednjoafrička Republika"
            },
            "ita": {
                "official": "Repubblica Centrafricana",
                "common": "Repubblica Centrafricana"
            },
            "jpn": {
                "official": "中央アフリカ共和国",
                "common": "中央アフリカ共和国"
            },
            "nld": {
                "official": "Centraal-Afrikaanse Republiek",
                "common": "Centraal-Afrikaanse Republiek"
            },
            "por": {
                "official": "República Centro-Africano",
                "common": "República Centro-Africana"
            },
            "rus": {
                "official": "Центрально-Африканская Республика",
                "common": "Центральноафриканская Республика"
            },
            "spa": {
                "official": "República Centroafricana",
                "common": "República Centroafricana"
            },
            "fin": {
                "official": "Keski-Afrikan tasavalta",
                "common": "Keski-Afrikan tasavalta"
            }
        },
        "latlng": [
            7,
            21
        ],
        "demonym": "Central African",
        "landlocked": true,
        "borders": [
            "CMR",
            "TCD",
            "COD",
            "COG",
            "SSD",
            "SDN"
        ],
        "area": 622984
    },
    {
        "name": {
            "common": "Canada",
            "official": "Canada",
            "native": {
                "eng": {
                    "official": "Canada",
                    "common": "Canada"
                },
                "fra": {
                    "official": "Canada",
                    "common": "Canada"
                }
            }
        },
        "tld": [
            ".ca"
        ],
        "cca2": "CA",
        "ccn3": "124",
        "cca3": "CAN",
        "cioc": "CAN",
        "currency": [
            "CAD"
        ],
        "callingCode": [
            "1"
        ],
        "capital": "Ottawa",
        "altSpellings": [
            "CA"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "eng": "English",
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Canada",
                "common": "Canada"
            },
            "deu": {
                "official": "Kanada",
                "common": "Kanada"
            },
            "fra": {
                "official": "Canada",
                "common": "Canada"
            },
            "hrv": {
                "official": "Kanada",
                "common": "Kanada"
            },
            "ita": {
                "official": "Canada",
                "common": "Canada"
            },
            "jpn": {
                "official": "カナダ",
                "common": "カナダ"
            },
            "nld": {
                "official": "Canada",
                "common": "Canada"
            },
            "por": {
                "official": "Canadá",
                "common": "Canadá"
            },
            "rus": {
                "official": "Канада",
                "common": "Канада"
            },
            "spa": {
                "official": "Canadá",
                "common": "Canadá"
            },
            "fin": {
                "official": "Kanada",
                "common": "Kanada"
            }
        },
        "latlng": [
            60,
            -95
        ],
        "demonym": "Canadian",
        "landlocked": false,
        "borders": [
            "USA"
        ],
        "area": 9984670
    },
    {
        "name": {
            "common": "Cocos (Keeling) Islands",
            "official": "Territory of the Cocos (Keeling) Islands",
            "native": {
                "eng": {
                    "official": "Territory of the Cocos (Keeling) Islands",
                    "common": "Cocos (Keeling) Islands"
                }
            }
        },
        "tld": [
            ".cc"
        ],
        "cca2": "CC",
        "ccn3": "166",
        "cca3": "CCK",
        "cioc": "",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "61"
        ],
        "capital": "West Island",
        "altSpellings": [
            "CC",
            "Territory of the Cocos (Keeling) Islands",
            "Keeling Islands"
        ],
        "region": "Oceania",
        "subregion": "Australia and New Zealand",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Territory of the Cocos (Keeling) Islands",
                "common": "Ynysoedd Cocos"
            },
            "deu": {
                "official": "Gebiet der Cocos (Keeling) Islands",
                "common": "Kokosinseln"
            },
            "fra": {
                "official": "Territoire des îles Cocos (Keeling)",
                "common": "Îles Cocos"
            },
            "hrv": {
                "official": "Teritoriju Kokosovi (Keeling) Islands",
                "common": "Kokosovi Otoci"
            },
            "ita": {
                "official": "Territorio della (Keeling) Isole Cocos",
                "common": "Isole Cocos e Keeling"
            },
            "jpn": {
                "official": "ココス諸島の領土",
                "common": "ココス（キーリング）諸島"
            },
            "nld": {
                "official": "Grondgebied van de Eilanden Cocos (Keeling )",
                "common": "Cocoseilanden"
            },
            "por": {
                "official": "Território dos Cocos (Keeling)",
                "common": "Ilhas Cocos (Keeling)"
            },
            "rus": {
                "official": "Территория Кокосовые (Килинг) острова",
                "common": "Кокосовые острова"
            },
            "spa": {
                "official": "Territorio de los (Keeling) Islas Cocos",
                "common": "Islas Cocos o Islas Keeling"
            },
            "fin": {
                "official": "Kookossaaret",
                "common": "Kookossaaret"
            }
        },
        "latlng": [
            -12.5,
            96.83333333
        ],
        "demonym": "Cocos Islander",
        "landlocked": false,
        "borders": [],
        "area": 14
    },
    {
        "name": {
            "common": "Switzerland",
            "official": "Swiss Confederation",
            "native": {
                "fra": {
                    "official": "Confédération suisse",
                    "common": "Suisse"
                },
                "gsw": {
                    "official": "Schweizerische Eidgenossenschaft",
                    "common": "Schweiz"
                },
                "ita": {
                    "official": "Confederazione Svizzera",
                    "common": "Svizzera"
                },
                "roh": {
                    "official": "Confederaziun svizra",
                    "common": "Svizra"
                }
            }
        },
        "tld": [
            ".ch"
        ],
        "cca2": "CH",
        "ccn3": "756",
        "cca3": "CHE",
        "cioc": "SUI",
        "currency": [
            "CHE",
            "CHF",
            "CHW"
        ],
        "callingCode": [
            "41"
        ],
        "capital": "Bern",
        "altSpellings": [
            "CH",
            "Swiss Confederation",
            "Schweiz",
            "Suisse",
            "Svizzera",
            "Svizra"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "fra": "French",
            "gsw": "Swiss German",
            "ita": "Italian",
            "roh": "Romansh"
        },
        "translations": {
            "deu": {
                "official": "Schweizerische Eidgenossenschaft",
                "common": "Schweiz"
            },
            "fra": {
                "official": "Confédération suisse",
                "common": "Suisse"
            },
            "hrv": {
                "official": "švicarska Konfederacija",
                "common": "Švicarska"
            },
            "ita": {
                "official": "Confederazione svizzera",
                "common": "Svizzera"
            },
            "jpn": {
                "official": "スイス連邦",
                "common": "スイス"
            },
            "nld": {
                "official": "Zwitserse Confederatie",
                "common": "Zwitserland"
            },
            "por": {
                "official": "Confederação Suíça",
                "common": "Suíça"
            },
            "rus": {
                "official": "Швейцарская Конфедерация",
                "common": "Швейцария"
            },
            "spa": {
                "official": "Confederación Suiza",
                "common": "Suiza"
            },
            "fin": {
                "official": "Sveitsin valaliitto",
                "common": "Sveitsi"
            }
        },
        "latlng": [
            47,
            8
        ],
        "demonym": "Swiss",
        "landlocked": true,
        "borders": [
            "AUT",
            "FRA",
            "ITA",
            "LIE",
            "DEU"
        ],
        "area": 41284
    },
    {
        "name": {
            "common": "Chile",
            "official": "Republic of Chile",
            "native": {
                "spa": {
                    "official": "República de Chile",
                    "common": "Chile"
                }
            }
        },
        "tld": [
            ".cl"
        ],
        "cca2": "CL",
        "ccn3": "152",
        "cca3": "CHL",
        "cioc": "CHI",
        "currency": [
            "CLF",
            "CLP"
        ],
        "callingCode": [
            "56"
        ],
        "capital": "Santiago",
        "altSpellings": [
            "CL",
            "Republic of Chile",
            "República de Chile"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Chile",
                "common": "Chile"
            },
            "deu": {
                "official": "Republik Chile",
                "common": "Chile"
            },
            "fra": {
                "official": "République du Chili",
                "common": "Chili"
            },
            "hrv": {
                "official": "Republika Čile",
                "common": "Čile"
            },
            "ita": {
                "official": "Repubblica del Cile",
                "common": "Cile"
            },
            "jpn": {
                "official": "チリ共和国",
                "common": "チリ"
            },
            "nld": {
                "official": "Republiek Chili",
                "common": "Chili"
            },
            "por": {
                "official": "República do Chile",
                "common": "Chile"
            },
            "rus": {
                "official": "Республика Чили",
                "common": "Чили"
            },
            "spa": {
                "official": "República de Chile",
                "common": "Chile"
            },
            "fin": {
                "official": "Chilen tasavalta",
                "common": "Chile"
            }
        },
        "latlng": [
            -30,
            -71
        ],
        "demonym": "Chilean",
        "landlocked": false,
        "borders": [
            "ARG",
            "BOL",
            "PER"
        ],
        "area": 756102
    },
    {
        "name": {
            "common": "China",
            "official": "People's Republic of China",
            "native": {
                "cmn": {
                    "official": "中华人民共和国",
                    "common": "中国"
                }
            }
        },
        "tld": [
            ".cn",
            ".中国",
            ".中國",
            ".公司",
            ".网络"
        ],
        "cca2": "CN",
        "ccn3": "156",
        "cca3": "CHN",
        "cioc": "CHN",
        "currency": [
            "CNY"
        ],
        "callingCode": [
            "86"
        ],
        "capital": "Beijing",
        "altSpellings": [
            "CN",
            "Zhōngguó",
            "Zhongguo",
            "Zhonghua",
            "People's Republic of China",
            "中华人民共和国",
            "Zhōnghuá Rénmín Gònghéguó"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "cmn": "Mandarin"
        },
        "translations": {
            "cym": {
                "official": "People's Republic of China",
                "common": "Tsieina"
            },
            "deu": {
                "official": "Volksrepublik China",
                "common": "China"
            },
            "fra": {
                "official": "République populaire de Chine",
                "common": "Chine"
            },
            "hrv": {
                "official": "Narodna Republika Kina",
                "common": "Kina"
            },
            "ita": {
                "official": "Repubblica popolare cinese",
                "common": "Cina"
            },
            "jpn": {
                "official": "中華人民共和国",
                "common": "中国"
            },
            "nld": {
                "official": "Volksrepubliek China",
                "common": "China"
            },
            "por": {
                "official": "República Popular da China",
                "common": "China"
            },
            "rus": {
                "official": "Народная Республика Китай",
                "common": "Китай"
            },
            "spa": {
                "official": "República Popular de China",
                "common": "China"
            },
            "fin": {
                "official": "Kiinan kansantasavalta",
                "common": "Kiina"
            }
        },
        "latlng": [
            35,
            105
        ],
        "demonym": "Chinese",
        "landlocked": false,
        "borders": [
            "AFG",
            "BTN",
            "MMR",
            "HKG",
            "IND",
            "KAZ",
            "PRK",
            "KGZ",
            "LAO",
            "MAC",
            "MNG",
            "PAK",
            "RUS",
            "TJK",
            "VNM"
        ],
        "area": 9706961
    },
    {
        "name": {
            "common": "Ivory Coast",
            "official": "Republic of Côte d'Ivoire",
            "native": {
                "fra": {
                    "official": "République de Côte d'Ivoire",
                    "common": "Côte d'Ivoire"
                }
            }
        },
        "tld": [
            ".ci"
        ],
        "cca2": "CI",
        "ccn3": "384",
        "cca3": "CIV",
        "cioc": "CIV",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "225"
        ],
        "capital": "Yamoussoukro",
        "altSpellings": [
            "CI",
            "Côte d'Ivoire",
            "Ivory Coast",
            "Republic of Côte d'Ivoire",
            "République de Côte d'Ivoire"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Côte d'Ivoire",
                "common": "Elfenbeinküste"
            },
            "fra": {
                "official": "République de Côte d' Ivoire",
                "common": "Côte d'Ivoire"
            },
            "hrv": {
                "official": "Republika Côte d'Ivoire",
                "common": "Obala Bjelokosti"
            },
            "ita": {
                "official": "Repubblica della Costa d'Avorio",
                "common": "Costa d'Avorio"
            },
            "jpn": {
                "official": "コートジボワール共和国",
                "common": "コートジボワール"
            },
            "nld": {
                "official": "Republiek Ivoorkust",
                "common": "Ivoorkust"
            },
            "por": {
                "official": "República da Côte d'Ivoire",
                "common": "Costa do Marfim"
            },
            "rus": {
                "official": "Республика Кот-д'Ивуаре",
                "common": "Кот-д’Ивуар"
            },
            "spa": {
                "official": "República de Côte d'Ivoire",
                "common": "Costa de Marfil"
            },
            "fin": {
                "official": "Norsunluurannikon tasavalta",
                "common": "Norsunluurannikko"
            }
        },
        "latlng": [
            8,
            -5
        ],
        "demonym": "Ivorian",
        "landlocked": false,
        "borders": [
            "BFA",
            "GHA",
            "GIN",
            "LBR",
            "MLI"
        ],
        "area": 322463
    },
    {
        "name": {
            "common": "Cameroon",
            "official": "Republic of Cameroon",
            "native": {
                "eng": {
                    "official": "Republic of Cameroon",
                    "common": "Cameroon"
                },
                "fra": {
                    "official": "République du Cameroun",
                    "common": "Cameroun"
                }
            }
        },
        "tld": [
            ".cm"
        ],
        "cca2": "CM",
        "ccn3": "120",
        "cca3": "CMR",
        "cioc": "CMR",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "237"
        ],
        "capital": "Yaoundé",
        "altSpellings": [
            "CM",
            "Republic of Cameroon",
            "République du Cameroun"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "eng": "English",
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Republic of Cameroon",
                "common": "Camerŵn"
            },
            "deu": {
                "official": "Republik Kamerun",
                "common": "Kamerun"
            },
            "fra": {
                "official": "République du Cameroun",
                "common": "Cameroun"
            },
            "hrv": {
                "official": "Republika Kamerun",
                "common": "Kamerun"
            },
            "ita": {
                "official": "Repubblica del Camerun",
                "common": "Camerun"
            },
            "jpn": {
                "official": "カメルーン共和国",
                "common": "カメルーン"
            },
            "nld": {
                "official": "Republiek Kameroen",
                "common": "Kameroen"
            },
            "por": {
                "official": "República dos Camarões",
                "common": "Camarões"
            },
            "rus": {
                "official": "Республика Камерун",
                "common": "Камерун"
            },
            "spa": {
                "official": "República de Camerún",
                "common": "Camerún"
            },
            "fin": {
                "official": "Kamerunin tasavalta",
                "common": "Kamerun"
            }
        },
        "latlng": [
            6,
            12
        ],
        "demonym": "Cameroonian",
        "landlocked": false,
        "borders": [
            "CAF",
            "TCD",
            "COG",
            "GNQ",
            "GAB",
            "NGA"
        ],
        "area": 475442
    },
    {
        "name": {
            "common": "DR Congo",
            "official": "Democratic Republic of the Congo",
            "native": {
                "fra": {
                    "official": "République démocratique du Congo",
                    "common": "RD Congo"
                },
                "kon": {
                    "official": "Repubilika ya Kongo Demokratiki",
                    "common": "Repubilika ya Kongo Demokratiki"
                },
                "lin": {
                    "official": "Republiki ya Kongó Demokratiki",
                    "common": "Republiki ya Kongó Demokratiki"
                },
                "lua": {
                    "official": "Ditunga dia Kongu wa Mungalaata",
                    "common": "Ditunga dia Kongu wa Mungalaata"
                },
                "swa": {
                    "official": "Jamhuri ya Kidemokrasia ya Kongo",
                    "common": "Jamhuri ya Kidemokrasia ya Kongo"
                }
            }
        },
        "tld": [
            ".cd"
        ],
        "cca2": "CD",
        "ccn3": "180",
        "cca3": "COD",
        "cioc": "COD",
        "currency": [
            "CDF"
        ],
        "callingCode": [
            "243"
        ],
        "capital": "Kinshasa",
        "altSpellings": [
            "CD",
            "DR Congo",
            "Congo-Kinshasa",
            "Congo, the Democratic Republic of the",
            "DRC"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "fra": "French",
            "kon": "Kikongo",
            "lin": "Lingala",
            "lua": "Tshiluba",
            "swa": "Swahili"
        },
        "translations": {
            "cym": {
                "official": "Democratic Republic of the Congo",
                "common": "Gweriniaeth Ddemocrataidd Congo"
            },
            "deu": {
                "official": "Demokratische Republik Kongo",
                "common": "Kongo (Dem. Rep.)"
            },
            "fra": {
                "official": "République démocratique du Congo",
                "common": "Congo (Rép. dém.)"
            },
            "hrv": {
                "official": "Demokratska Republika Kongo",
                "common": "Kongo, Demokratska Republika"
            },
            "ita": {
                "official": "Repubblica Democratica del Congo",
                "common": "Congo (Rep. Dem.)"
            },
            "jpn": {
                "official": "コンゴ民主共和国",
                "common": "コンゴ民主共和国"
            },
            "nld": {
                "official": "Democratische Republiek Congo",
                "common": "Congo (DRC)"
            },
            "por": {
                "official": "República Democrática do Congo",
                "common": "República Democrática do Congo"
            },
            "rus": {
                "official": "Демократическая Республика Конго",
                "common": "Демократическая Республика Конго"
            },
            "spa": {
                "official": "República Democrática del Congo",
                "common": "Congo (Rep. Dem.)"
            },
            "fin": {
                "official": "Kongon demokraattinen tasavalta",
                "common": "Kongon demokraattinen tasavalta"
            }
        },
        "latlng": [
            0,
            25
        ],
        "demonym": "Congolese",
        "landlocked": false,
        "borders": [
            "AGO",
            "BDI",
            "CAF",
            "COG",
            "RWA",
            "SSD",
            "TZA",
            "UGA",
            "ZMB"
        ],
        "area": 2344858
    },
    {
        "name": {
            "common": "Republic of the Congo",
            "official": "Republic of the Congo",
            "native": {
                "fra": {
                    "official": "République du Congo",
                    "common": "République du Congo"
                },
                "kon": {
                    "official": "Repubilika ya Kongo",
                    "common": "Repubilika ya Kongo"
                },
                "lin": {
                    "official": "Republíki ya Kongó",
                    "common": "Republíki ya Kongó"
                }
            }
        },
        "tld": [
            ".cg"
        ],
        "cca2": "CG",
        "ccn3": "178",
        "cca3": "COG",
        "cioc": "CGO",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "242"
        ],
        "capital": "Brazzaville",
        "altSpellings": [
            "CG",
            "Congo",
            "Congo-Brazzaville"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "fra": "French",
            "kon": "Kikongo",
            "lin": "Lingala"
        },
        "translations": {
            "cym": {
                "official": "Republic of the Congo",
                "common": "Gweriniaeth y Congo"
            },
            "deu": {
                "official": "Republik Kongo",
                "common": "Kongo"
            },
            "fra": {
                "official": "République du Congo",
                "common": "Congo"
            },
            "hrv": {
                "official": "Republika Kongo",
                "common": "Kongo"
            },
            "ita": {
                "official": "Repubblica del Congo",
                "common": "Congo"
            },
            "jpn": {
                "official": "コンゴ共和国",
                "common": "コンゴ共和国"
            },
            "nld": {
                "official": "Republiek Congo",
                "common": "Congo"
            },
            "por": {
                "official": "República do Congo",
                "common": "Congo"
            },
            "rus": {
                "official": "Республика Конго",
                "common": "Республика Конго"
            },
            "spa": {
                "official": "República del Congo",
                "common": "Congo"
            },
            "fin": {
                "official": "Kongon tasavalta",
                "common": "Kongo-Brazzaville"
            }
        },
        "latlng": [
            -1,
            15
        ],
        "demonym": "Congolese",
        "landlocked": false,
        "borders": [
            "AGO",
            "CMR",
            "CAF",
            "COD",
            "GAB"
        ],
        "area": 342000
    },
    {
        "name": {
            "common": "Cook Islands",
            "official": "Cook Islands",
            "native": {
                "eng": {
                    "official": "Cook Islands",
                    "common": "Cook Islands"
                },
                "rar": {
                    "official": "Kūki 'Āirani",
                    "common": "Kūki 'Āirani"
                }
            }
        },
        "tld": [
            ".ck"
        ],
        "cca2": "CK",
        "ccn3": "184",
        "cca3": "COK",
        "cioc": "COK",
        "currency": [
            "NZD"
        ],
        "callingCode": [
            "682"
        ],
        "capital": "Avarua",
        "altSpellings": [
            "CK",
            "Kūki 'Āirani"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "rar": "Cook Islands Māori"
        },
        "translations": {
            "cym": {
                "official": "Cook Islands",
                "common": "Ynysoedd Cook"
            },
            "deu": {
                "official": "Cook-Inseln",
                "common": "Cookinseln"
            },
            "fra": {
                "official": "Îles Cook",
                "common": "Îles Cook"
            },
            "hrv": {
                "official": "Cook Islands",
                "common": "Cookovo Otočje"
            },
            "ita": {
                "official": "Isole Cook",
                "common": "Isole Cook"
            },
            "jpn": {
                "official": "クック諸島",
                "common": "クック諸島"
            },
            "nld": {
                "official": "Cook eilanden",
                "common": "Cookeilanden"
            },
            "por": {
                "official": "Ilhas Cook",
                "common": "Ilhas Cook"
            },
            "rus": {
                "official": "острова Кука",
                "common": "Острова Кука"
            },
            "spa": {
                "official": "Islas Cook",
                "common": "Islas Cook"
            },
            "fin": {
                "official": "Cookinsaaret",
                "common": "Cookinsaaret"
            }
        },
        "latlng": [
            -21.23333333,
            -159.76666666
        ],
        "demonym": "Cook Islander",
        "landlocked": false,
        "borders": [],
        "area": 236
    },
    {
        "name": {
            "common": "Colombia",
            "official": "Republic of Colombia",
            "native": {
                "spa": {
                    "official": "República de Colombia",
                    "common": "Colombia"
                }
            }
        },
        "tld": [
            ".co"
        ],
        "cca2": "CO",
        "ccn3": "170",
        "cca3": "COL",
        "cioc": "COL",
        "currency": [
            "COP"
        ],
        "callingCode": [
            "57"
        ],
        "capital": "Bogotá",
        "altSpellings": [
            "CO",
            "Republic of Colombia",
            "República de Colombia"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Colombia",
                "common": "Colombia"
            },
            "deu": {
                "official": "Republik Kolumbien",
                "common": "Kolumbien"
            },
            "fra": {
                "official": "République de Colombie",
                "common": "Colombie"
            },
            "hrv": {
                "official": "Republika Kolumbija",
                "common": "Kolumbija"
            },
            "ita": {
                "official": "Repubblica di Colombia",
                "common": "Colombia"
            },
            "jpn": {
                "official": "コロンビア共和国",
                "common": "コロンビア"
            },
            "nld": {
                "official": "Republiek Colombia",
                "common": "Colombia"
            },
            "por": {
                "official": "República da Colômbia",
                "common": "Colômbia"
            },
            "rus": {
                "official": "Республика Колумбия",
                "common": "Колумбия"
            },
            "spa": {
                "official": "República de Colombia",
                "common": "Colombia"
            },
            "fin": {
                "official": "Kolumbian tasavalta",
                "common": "Kolumbia"
            }
        },
        "latlng": [
            4,
            -72
        ],
        "demonym": "Colombian",
        "landlocked": false,
        "borders": [
            "BRA",
            "ECU",
            "PAN",
            "PER",
            "VEN"
        ],
        "area": 1141748
    },
    {
        "name": {
            "common": "Comoros",
            "official": "Union of the Comoros",
            "native": {
                "ara": {
                    "official": "الاتحاد القمري",
                    "common": "القمر‎"
                },
                "fra": {
                    "official": "Union des Comores",
                    "common": "Comores"
                },
                "zdj": {
                    "official": "Udzima wa Komori",
                    "common": "Komori"
                }
            }
        },
        "tld": [
            ".km"
        ],
        "cca2": "KM",
        "ccn3": "174",
        "cca3": "COM",
        "cioc": "COM",
        "currency": [
            "KMF"
        ],
        "callingCode": [
            "269"
        ],
        "capital": "Moroni",
        "altSpellings": [
            "KM",
            "Union of the Comoros",
            "Union des Comores",
            "Udzima wa Komori",
            "al-Ittiḥād al-Qumurī"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "ara": "Arabic",
            "fra": "French",
            "zdj": "Comorian"
        },
        "translations": {
            "cym": {
                "official": "Union of the Comoros",
                "common": "Comoros"
            },
            "deu": {
                "official": "Union der Komoren",
                "common": "Union der Komoren"
            },
            "fra": {
                "official": "Union des Comores",
                "common": "Comores"
            },
            "hrv": {
                "official": "Savez Komori",
                "common": "Komori"
            },
            "ita": {
                "official": "Unione delle Comore",
                "common": "Comore"
            },
            "jpn": {
                "official": "コモロ連合",
                "common": "コモロ"
            },
            "nld": {
                "official": "Unie van de Comoren",
                "common": "Comoren"
            },
            "por": {
                "official": "União das Comores",
                "common": "Comores"
            },
            "rus": {
                "official": "Союз Коморских Островов",
                "common": "Коморы"
            },
            "spa": {
                "official": "Unión de las Comoras",
                "common": "Comoras"
            },
            "fin": {
                "official": "Komorien liitto",
                "common": "Komorit"
            }
        },
        "latlng": [
            -12.16666666,
            44.25
        ],
        "demonym": "Comoran",
        "landlocked": false,
        "borders": [],
        "area": 1862
    },
    {
        "name": {
            "common": "Cape Verde",
            "official": "Republic of Cabo Verde",
            "native": {
                "por": {
                    "official": "República de Cabo Verde",
                    "common": "Cabo Verde"
                }
            }
        },
        "tld": [
            ".cv"
        ],
        "cca2": "CV",
        "ccn3": "132",
        "cca3": "CPV",
        "cioc": "CPV",
        "currency": [
            "CVE"
        ],
        "callingCode": [
            "238"
        ],
        "capital": "Praia",
        "altSpellings": [
            "CV",
            "Republic of Cabo Verde",
            "República de Cabo Verde"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "cym": {
                "official": "Republic of Cabo Verde",
                "common": "Cape Verde"
            },
            "deu": {
                "official": "Republik Cabo Verde",
                "common": "Kap Verde"
            },
            "fra": {
                "official": "République du Cap-Vert",
                "common": "Îles du Cap-Vert"
            },
            "hrv": {
                "official": "Republika Cabo Verde",
                "common": "Zelenortska Republika"
            },
            "ita": {
                "official": "Repubblica di Capo Verde",
                "common": "Capo Verde"
            },
            "jpn": {
                "official": "カーボベルデ共和国",
                "common": "カーボベルデ"
            },
            "nld": {
                "official": "Republiek van Cabo Verde",
                "common": "Kaapverdië"
            },
            "por": {
                "official": "República de Cabo Verde",
                "common": "Cabo Verde"
            },
            "rus": {
                "official": "Республика Кабо -Верде",
                "common": "Кабо-Верде"
            },
            "spa": {
                "official": "República de Cabo Verde",
                "common": "Cabo Verde"
            },
            "fin": {
                "official": "Kap Verden tasavalta",
                "common": "Kap Verde"
            }
        },
        "latlng": [
            16,
            -24
        ],
        "demonym": "Cape Verdian",
        "landlocked": false,
        "borders": [],
        "area": 4033
    },
    {
        "name": {
            "common": "Costa Rica",
            "official": "Republic of Costa Rica",
            "native": {
                "spa": {
                    "official": "República de Costa Rica",
                    "common": "Costa Rica"
                }
            }
        },
        "tld": [
            ".cr"
        ],
        "cca2": "CR",
        "ccn3": "188",
        "cca3": "CRI",
        "cioc": "CRC",
        "currency": [
            "CRC"
        ],
        "callingCode": [
            "506"
        ],
        "capital": "San José",
        "altSpellings": [
            "CR",
            "Republic of Costa Rica",
            "República de Costa Rica"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Costa Rica",
                "common": "Costa Rica"
            },
            "deu": {
                "official": "Republik Costa Rica",
                "common": "Costa Rica"
            },
            "fra": {
                "official": "République du Costa Rica",
                "common": "Costa Rica"
            },
            "hrv": {
                "official": "Republika Kostarika",
                "common": "Kostarika"
            },
            "ita": {
                "official": "Repubblica di Costa Rica",
                "common": "Costa Rica"
            },
            "jpn": {
                "official": "コスタリカ共和国",
                "common": "コスタリカ"
            },
            "nld": {
                "official": "Republiek Costa Rica",
                "common": "Costa Rica"
            },
            "por": {
                "official": "República da Costa Rica",
                "common": "Costa Rica"
            },
            "rus": {
                "official": "Республика Коста-Рика",
                "common": "Коста-Рика"
            },
            "spa": {
                "official": "República de Costa Rica",
                "common": "Costa Rica"
            },
            "fin": {
                "official": "Costa Rican tasavalta",
                "common": "Costa Rica"
            }
        },
        "latlng": [
            10,
            -84
        ],
        "demonym": "Costa Rican",
        "landlocked": false,
        "borders": [
            "NIC",
            "PAN"
        ],
        "area": 51100
    },
    {
        "name": {
            "common": "Cuba",
            "official": "Republic of Cuba",
            "native": {
                "spa": {
                    "official": "República de Cuba",
                    "common": "Cuba"
                }
            }
        },
        "tld": [
            ".cu"
        ],
        "cca2": "CU",
        "ccn3": "192",
        "cca3": "CUB",
        "cioc": "CUB",
        "currency": [
            "CUC",
            "CUP"
        ],
        "callingCode": [
            "53"
        ],
        "capital": "Havana",
        "altSpellings": [
            "CU",
            "Republic of Cuba",
            "República de Cuba"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Cuba",
                "common": "Ciwba"
            },
            "deu": {
                "official": "Republik Kuba",
                "common": "Kuba"
            },
            "fra": {
                "official": "République de Cuba",
                "common": "Cuba"
            },
            "hrv": {
                "official": "Republika Kuba",
                "common": "Kuba"
            },
            "ita": {
                "official": "Repubblica di Cuba",
                "common": "Cuba"
            },
            "jpn": {
                "official": "キューバ共和国",
                "common": "キューバ"
            },
            "nld": {
                "official": "Republiek Cuba",
                "common": "Cuba"
            },
            "por": {
                "official": "República de Cuba",
                "common": "Cuba"
            },
            "rus": {
                "official": "Республика Куба",
                "common": "Куба"
            },
            "spa": {
                "official": "República de Cuba",
                "common": "Cuba"
            },
            "fin": {
                "official": "Kuuban tasavalta",
                "common": "Kuuba"
            }
        },
        "latlng": [
            21.5,
            -80
        ],
        "demonym": "Cuban",
        "landlocked": false,
        "borders": [],
        "area": 109884
    },
    {
        "name": {
            "common": "Curaçao",
            "official": "Country of Curaçao",
            "native": {
                "eng": {
                    "official": "Country of Curaçao",
                    "common": "Curaçao"
                },
                "nld": {
                    "official": "Land Curaçao",
                    "common": "Curaçao"
                },
                "pap": {
                    "official": "Pais Kòrsou",
                    "common": "Pais Kòrsou"
                }
            }
        },
        "tld": [
            ".cw"
        ],
        "cca2": "CW",
        "ccn3": "531",
        "cca3": "CUW",
        "cioc": "",
        "currency": [
            "ANG"
        ],
        "callingCode": [
            "5999"
        ],
        "capital": "Willemstad",
        "altSpellings": [
            "CW",
            "Curacao",
            "Kòrsou",
            "Country of Curaçao",
            "Land Curaçao",
            "Pais Kòrsou"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English",
            "nld": "Dutch",
            "pap": "Papiamento"
        },
        "translations": {
            "nld": {
                "official": "Land Curaçao",
                "common": "Curaçao"
            },
            "por": {
                "official": "País de Curaçao",
                "common": "ilha da Curação"
            },
            "rus": {
                "official": "Страна Кюрасао",
                "common": "Кюрасао"
            },
            "spa": {
                "official": "País de Curazao",
                "common": "Curazao"
            },
            "fin": {
                "official": "Curaçao",
                "common": "Curaçao"
            }
        },
        "latlng": [
            12.116667,
            -68.933333
        ],
        "demonym": "Dutch",
        "landlocked": false,
        "borders": [],
        "area": 444
    },
    {
        "name": {
            "common": "Christmas Island",
            "official": "Territory of Christmas Island",
            "native": {
                "eng": {
                    "official": "Territory of Christmas Island",
                    "common": "Christmas Island"
                }
            }
        },
        "tld": [
            ".cx"
        ],
        "cca2": "CX",
        "ccn3": "162",
        "cca3": "CXR",
        "cioc": "",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "61"
        ],
        "capital": "Flying Fish Cove",
        "altSpellings": [
            "CX",
            "Territory of Christmas Island"
        ],
        "region": "Oceania",
        "subregion": "Australia and New Zealand",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Territory of Christmas Island",
                "common": "Ynys y Nadolig"
            },
            "deu": {
                "official": "Gebiet der Weihnachtsinsel",
                "common": "Weihnachtsinsel"
            },
            "fra": {
                "official": "Territoire de l'île Christmas",
                "common": "Île Christmas"
            },
            "hrv": {
                "official": "Teritorij Božićni otok",
                "common": "Božićni otok"
            },
            "ita": {
                "official": "Territorio di Christmas Island",
                "common": "Isola di Natale"
            },
            "jpn": {
                "official": "クリスマス島の領土",
                "common": "クリスマス島"
            },
            "nld": {
                "official": "Grondgebied van Christmas Island",
                "common": "Christmaseiland"
            },
            "por": {
                "official": "Território da Ilha Christmas",
                "common": "Ilha do Natal"
            },
            "rus": {
                "official": "Территория острова Рождества",
                "common": "Остров Рождества"
            },
            "spa": {
                "official": "Territorio de la Isla de Navidad",
                "common": "Isla de Navidad"
            },
            "fin": {
                "official": "Joulusaaren alue",
                "common": "Joulusaari"
            }
        },
        "latlng": [
            -10.5,
            105.66666666
        ],
        "demonym": "Christmas Island",
        "landlocked": false,
        "borders": [],
        "area": 135
    },
    {
        "name": {
            "common": "Cayman Islands",
            "official": "Cayman Islands",
            "native": {
                "eng": {
                    "official": "Cayman Islands",
                    "common": "Cayman Islands"
                }
            }
        },
        "tld": [
            ".ky"
        ],
        "cca2": "KY",
        "ccn3": "136",
        "cca3": "CYM",
        "cioc": "CAY",
        "currency": [
            "KYD"
        ],
        "callingCode": [
            "1345"
        ],
        "capital": "George Town",
        "altSpellings": [
            "KY"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Cayman Islands",
                "common": "Ynysoedd_Cayman"
            },
            "deu": {
                "official": "Cayman-Inseln",
                "common": "Kaimaninseln"
            },
            "fra": {
                "official": "Îles Caïmans",
                "common": "Îles Caïmans"
            },
            "hrv": {
                "official": "Kajmanski otoci",
                "common": "Kajmanski otoci"
            },
            "ita": {
                "official": "Isole Cayman",
                "common": "Isole Cayman"
            },
            "jpn": {
                "official": "ケイマン諸島",
                "common": "ケイマン諸島"
            },
            "nld": {
                "official": "Caymaneilanden",
                "common": "Caymaneilanden"
            },
            "por": {
                "official": "Ilhas Cayman",
                "common": "Ilhas Caimão"
            },
            "rus": {
                "official": "Каймановы острова",
                "common": "Каймановы острова"
            },
            "spa": {
                "official": "Islas Caimán",
                "common": "Islas Caimán"
            },
            "fin": {
                "official": "Caymansaaret",
                "common": "Caymansaaret"
            }
        },
        "latlng": [
            19.5,
            -80.5
        ],
        "demonym": "Caymanian",
        "landlocked": false,
        "borders": [],
        "area": 264
    },
    {
        "name": {
            "common": "Cyprus",
            "official": "Republic of Cyprus",
            "native": {
                "ell": {
                    "official": "Δημοκρατία της Κύπρος",
                    "common": "Κύπρος"
                },
                "tur": {
                    "official": "Kıbrıs Cumhuriyeti",
                    "common": "Kıbrıs"
                }
            }
        },
        "tld": [
            ".cy"
        ],
        "cca2": "CY",
        "ccn3": "196",
        "cca3": "CYP",
        "cioc": "CYP",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "357"
        ],
        "capital": "Nicosia",
        "altSpellings": [
            "CY",
            "Kýpros",
            "Kıbrıs",
            "Republic of Cyprus",
            "Κυπριακή Δημοκρατία",
            "Kıbrıs Cumhuriyeti"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "ell": "Greek",
            "tur": "Turkish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Cyprus",
                "common": "Cyprus"
            },
            "deu": {
                "official": "Republik Zypern",
                "common": "Zypern"
            },
            "fra": {
                "official": "République de Chypre",
                "common": "Chypre"
            },
            "hrv": {
                "official": "Republika Cipar",
                "common": "Cipar"
            },
            "ita": {
                "official": "Repubblica di Cipro",
                "common": "Cipro"
            },
            "jpn": {
                "official": "キプロス共和国",
                "common": "キプロス"
            },
            "nld": {
                "official": "Republiek Cyprus",
                "common": "Cyprus"
            },
            "por": {
                "official": "República de Chipre",
                "common": "Chipre"
            },
            "rus": {
                "official": "Республика Кипр",
                "common": "Кипр"
            },
            "spa": {
                "official": "República de Chipre",
                "common": "Chipre"
            },
            "fin": {
                "official": "Kyproksen tasavalta",
                "common": "Kypros"
            }
        },
        "latlng": [
            35,
            33
        ],
        "demonym": "Cypriot",
        "landlocked": false,
        "borders": [
            "GBR"
        ],
        "area": 9251
    },
    {
        "name": {
            "common": "Czech Republic",
            "official": "Czech Republic",
            "native": {
                "ces": {
                    "official": "česká republika",
                    "common": "Česká republika"
                },
                "slk": {
                    "official": "Česká republika",
                    "common": "Česká republika"
                }
            }
        },
        "tld": [
            ".cz"
        ],
        "cca2": "CZ",
        "ccn3": "203",
        "cca3": "CZE",
        "cioc": "CZE",
        "currency": [
            "CZK"
        ],
        "callingCode": [
            "420"
        ],
        "capital": "Prague",
        "altSpellings": [
            "CZ",
            "Česká republika",
            "Česko"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "ces": "Czech",
            "slk": "Slovak"
        },
        "translations": {
            "cym": {
                "official": "Czech Republic",
                "common": "Y Weriniaeth Tsiec"
            },
            "deu": {
                "official": "Tschechische Republik",
                "common": "Tschechische Republik"
            },
            "fra": {
                "official": "République tchèque",
                "common": "République tchèque"
            },
            "hrv": {
                "official": "Češka",
                "common": "Češka"
            },
            "ita": {
                "official": "Repubblica Ceca",
                "common": "Repubblica Ceca"
            },
            "jpn": {
                "official": "チェコ共和国",
                "common": "チェコ"
            },
            "nld": {
                "official": "Tsjechische Republiek",
                "common": "Tsjechië"
            },
            "por": {
                "official": "República Checa",
                "common": "República Checa"
            },
            "rus": {
                "official": "Чешская Республика",
                "common": "Чехия"
            },
            "spa": {
                "official": "República Checa",
                "common": "República Checa"
            },
            "fin": {
                "official": "Tšekin tasavalta",
                "common": "Tšekki"
            }
        },
        "latlng": [
            49.75,
            15.5
        ],
        "demonym": "Czech",
        "landlocked": true,
        "borders": [
            "AUT",
            "DEU",
            "POL",
            "SVK"
        ],
        "area": 78865
    },
    {
        "name": {
            "common": "Germany",
            "official": "Federal Republic of Germany",
            "native": {
                "deu": {
                    "official": "Bundesrepublik Deutschland",
                    "common": "Deutschland"
                }
            }
        },
        "tld": [
            ".de"
        ],
        "cca2": "DE",
        "ccn3": "276",
        "cca3": "DEU",
        "cioc": "GER",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "49"
        ],
        "capital": "Berlin",
        "altSpellings": [
            "DE",
            "Federal Republic of Germany",
            "Bundesrepublik Deutschland"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "deu": "German"
        },
        "translations": {
            "deu": {
                "official": "Bundesrepublik Deutschland",
                "common": "Deutschland"
            },
            "fra": {
                "official": "République fédérale d'Allemagne",
                "common": "Allemagne"
            },
            "hrv": {
                "official": "Njemačka Federativna Republika",
                "common": "Njemačka"
            },
            "ita": {
                "official": "Repubblica federale di Germania",
                "common": "Germania"
            },
            "jpn": {
                "official": "ドイツ連邦共和国",
                "common": "ドイツ"
            },
            "nld": {
                "official": "Bondsrepubliek Duitsland",
                "common": "Duitsland"
            },
            "por": {
                "official": "República Federal da Alemanha",
                "common": "Alemanha"
            },
            "rus": {
                "official": "Федеративная Республика Германия",
                "common": "Германия"
            },
            "spa": {
                "official": "República Federal de Alemania",
                "common": "Alemania"
            },
            "fin": {
                "official": "Saksan liittotasavalta",
                "common": "Saksa"
            }
        },
        "latlng": [
            51,
            9
        ],
        "demonym": "German",
        "landlocked": false,
        "borders": [
            "AUT",
            "BEL",
            "CZE",
            "DNK",
            "FRA",
            "LUX",
            "NLD",
            "POL",
            "CHE"
        ],
        "area": 357114
    },
    {
        "name": {
            "common": "Djibouti",
            "official": "Republic of Djibouti",
            "native": {
                "ara": {
                    "official": "جمهورية جيبوتي",
                    "common": "جيبوتي‎"
                },
                "fra": {
                    "official": "République de Djibouti",
                    "common": "Djibouti"
                }
            }
        },
        "tld": [
            ".dj"
        ],
        "cca2": "DJ",
        "ccn3": "262",
        "cca3": "DJI",
        "cioc": "DJI",
        "currency": [
            "DJF"
        ],
        "callingCode": [
            "253"
        ],
        "capital": "Djibouti",
        "altSpellings": [
            "DJ",
            "Jabuuti",
            "Gabuuti",
            "Republic of Djibouti",
            "République de Djibouti",
            "Gabuutih Ummuuno",
            "Jamhuuriyadda Jabuuti"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "ara": "Arabic",
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Republic of Djibouti",
                "common": "Djibouti"
            },
            "deu": {
                "official": "Republik Dschibuti",
                "common": "Dschibuti"
            },
            "fra": {
                "official": "République de Djibouti",
                "common": "Djibouti"
            },
            "hrv": {
                "official": "Republika Džibuti",
                "common": "Džibuti"
            },
            "ita": {
                "official": "Repubblica di Gibuti",
                "common": "Gibuti"
            },
            "jpn": {
                "official": "ジブチ共和国",
                "common": "ジブチ"
            },
            "nld": {
                "official": "Republiek Djibouti",
                "common": "Djibouti"
            },
            "por": {
                "official": "República do Djibouti",
                "common": "Djibouti"
            },
            "rus": {
                "official": "Республика Джибути",
                "common": "Джибути"
            },
            "spa": {
                "official": "República de Djibouti",
                "common": "Djibouti"
            },
            "fin": {
                "official": "Dijiboutin tasavalta",
                "common": "Dijibouti"
            }
        },
        "latlng": [
            11.5,
            43
        ],
        "demonym": "Djibouti",
        "landlocked": false,
        "borders": [
            "ERI",
            "ETH",
            "SOM"
        ],
        "area": 23200
    },
    {
        "name": {
            "common": "Dominica",
            "official": "Commonwealth of Dominica",
            "native": {
                "eng": {
                    "official": "Commonwealth of Dominica",
                    "common": "Dominica"
                }
            }
        },
        "tld": [
            ".dm"
        ],
        "cca2": "DM",
        "ccn3": "212",
        "cca3": "DMA",
        "cioc": "DMA",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1767"
        ],
        "capital": "Roseau",
        "altSpellings": [
            "DM",
            "Dominique",
            "Wai‘tu kubuli",
            "Commonwealth of Dominica"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "Commonwealth of Dominica",
                "common": "Dominica"
            },
            "deu": {
                "official": "Commonwealth von Dominica",
                "common": "Dominica"
            },
            "fra": {
                "official": "Commonwealth de la Dominique",
                "common": "Dominique"
            },
            "hrv": {
                "official": "Zajednica Dominika",
                "common": "Dominika"
            },
            "ita": {
                "official": "Commonwealth di Dominica",
                "common": "Dominica"
            },
            "jpn": {
                "official": "ドミニカ国",
                "common": "ドミニカ国"
            },
            "nld": {
                "official": "Gemenebest Dominica",
                "common": "Dominica"
            },
            "por": {
                "official": "Comunidade da Dominica",
                "common": "Dominica"
            },
            "rus": {
                "official": "Содружество Доминики",
                "common": "Доминика"
            },
            "spa": {
                "official": "Mancomunidad de Dominica",
                "common": "Dominica"
            },
            "fin": {
                "official": "Dominican liittovaltio",
                "common": "Dominica"
            }
        },
        "latlng": [
            15.41666666,
            -61.33333333
        ],
        "demonym": "Dominican",
        "landlocked": false,
        "borders": [],
        "area": 751
    },
    {
        "name": {
            "common": "Denmark",
            "official": "Kingdom of Denmark",
            "native": {
                "dan": {
                    "official": "Kongeriget Danmark",
                    "common": "Danmark"
                }
            }
        },
        "tld": [
            ".dk"
        ],
        "cca2": "DK",
        "ccn3": "208",
        "cca3": "DNK",
        "cioc": "DEN",
        "currency": [
            "DKK"
        ],
        "callingCode": [
            "45"
        ],
        "capital": "Copenhagen",
        "altSpellings": [
            "DK",
            "Danmark",
            "Kingdom of Denmark",
            "Kongeriget Danmark"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "dan": "Danish"
        },
        "translations": {
            "cym": {
                "official": "Kingdom of Denmark",
                "common": "Denmarc"
            },
            "deu": {
                "official": "Königreich Dänemark",
                "common": "Dänemark"
            },
            "fra": {
                "official": "Royaume du Danemark",
                "common": "Danemark"
            },
            "hrv": {
                "official": "Kraljevina Danska",
                "common": "Danska"
            },
            "ita": {
                "official": "Regno di Danimarca",
                "common": "Danimarca"
            },
            "jpn": {
                "official": "デンマーク王国",
                "common": "デンマーク"
            },
            "nld": {
                "official": "Koninkrijk Denemarken",
                "common": "Denemarken"
            },
            "por": {
                "official": "Reino da Dinamarca",
                "common": "Dinamarca"
            },
            "rus": {
                "official": "Королевство Дания",
                "common": "Дания"
            },
            "spa": {
                "official": "Reino de Dinamarca",
                "common": "Dinamarca"
            },
            "fin": {
                "official": "Tanskan kuningaskunta",
                "common": "Tanska"
            }
        },
        "latlng": [
            56,
            10
        ],
        "demonym": "Danish",
        "landlocked": false,
        "borders": [
            "DEU"
        ],
        "area": 43094
    },
    {
        "name": {
            "common": "Dominican Republic",
            "official": "Dominican Republic",
            "native": {
                "spa": {
                    "official": "República Dominicana",
                    "common": "República Dominicana"
                }
            }
        },
        "tld": [
            ".do"
        ],
        "cca2": "DO",
        "ccn3": "214",
        "cca3": "DOM",
        "cioc": "DOM",
        "currency": [
            "DOP"
        ],
        "callingCode": [
            "1809",
            "1829",
            "1849"
        ],
        "capital": "Santo Domingo",
        "altSpellings": [
            "DO"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Dominican Republic",
                "common": "Gweriniaeth_Dominica"
            },
            "deu": {
                "official": "Dominikanische Republik",
                "common": "Dominikanische Republik"
            },
            "fra": {
                "official": "République Dominicaine",
                "common": "République dominicaine"
            },
            "hrv": {
                "official": "Dominikanska Republika",
                "common": "Dominikanska Republika"
            },
            "ita": {
                "official": "Repubblica Dominicana",
                "common": "Repubblica Dominicana"
            },
            "jpn": {
                "official": "ドミニカ共和国",
                "common": "ドミニカ共和国"
            },
            "nld": {
                "official": "Dominicaanse Republiek",
                "common": "Dominicaanse Republiek"
            },
            "por": {
                "official": "República Dominicana",
                "common": "República Dominicana"
            },
            "rus": {
                "official": "Доминиканская Республика",
                "common": "Доминиканская Республика"
            },
            "spa": {
                "official": "República Dominicana",
                "common": "República Dominicana"
            },
            "fin": {
                "official": "Dominikaaninen tasavalta",
                "common": "Dominikaaninen tasavalta"
            }
        },
        "latlng": [
            19,
            -70.66666666
        ],
        "demonym": "Dominican",
        "landlocked": false,
        "borders": [
            "HTI"
        ],
        "area": 48671
    },
    {
        "name": {
            "common": "Algeria",
            "official": "People's Democratic Republic of Algeria",
            "native": {
                "ara": {
                    "official": "الجمهورية الديمقراطية الشعبية الجزائرية",
                    "common": "الجزائر"
                }
            }
        },
        "tld": [
            ".dz",
            "الجزائر."
        ],
        "cca2": "DZ",
        "ccn3": "012",
        "cca3": "DZA",
        "cioc": "ALG",
        "currency": [
            "DZD"
        ],
        "callingCode": [
            "213"
        ],
        "capital": "Algiers",
        "altSpellings": [
            "DZ",
            "Dzayer",
            "Algérie"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "cym": {
                "official": "People's Democratic Republic of Algeria",
                "common": "Algeria"
            },
            "deu": {
                "official": "Demokratische Volksrepublik Algerien",
                "common": "Algerien"
            },
            "fra": {
                "official": "République démocratique et populaire d'Algérie",
                "common": "Algérie"
            },
            "hrv": {
                "official": "Narodna Demokratska Republika Alžir",
                "common": "Alžir"
            },
            "ita": {
                "official": "Repubblica popolare democratica di Algeria",
                "common": "Algeria"
            },
            "jpn": {
                "official": "アルジェリア人民民主共和国",
                "common": "アルジェリア"
            },
            "nld": {
                "official": "Democratische Volksrepubliek Algerije",
                "common": "Algerije"
            },
            "por": {
                "official": "República Argelina Democrática e Popular",
                "common": "Argélia"
            },
            "rus": {
                "official": "Народно-Демократическая Республика Алжир",
                "common": "Алжир"
            },
            "spa": {
                "official": "República Argelina Democrática y Popular",
                "common": "Argelia"
            },
            "fin": {
                "official": "Algerian demokraattinen kansantasavalta",
                "common": "Algeria"
            }
        },
        "latlng": [
            28,
            3
        ],
        "demonym": "Algerian",
        "landlocked": false,
        "borders": [
            "TUN",
            "LBY",
            "NER",
            "ESH",
            "MRT",
            "MLI",
            "MAR"
        ],
        "area": 2381741
    },
    {
        "name": {
            "common": "Ecuador",
            "official": "Republic of Ecuador",
            "native": {
                "spa": {
                    "official": "República del Ecuador",
                    "common": "Ecuador"
                }
            }
        },
        "tld": [
            ".ec"
        ],
        "cca2": "EC",
        "ccn3": "218",
        "cca3": "ECU",
        "cioc": "ECU",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "593"
        ],
        "capital": "Quito",
        "altSpellings": [
            "EC",
            "Republic of Ecuador",
            "República del Ecuador"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Ecuador",
                "common": "Ecwador"
            },
            "deu": {
                "official": "Republik Ecuador",
                "common": "Ecuador"
            },
            "fra": {
                "official": "République de l'Équateur",
                "common": "Équateur"
            },
            "hrv": {
                "official": "Republika Ekvador",
                "common": "Ekvador"
            },
            "ita": {
                "official": "Repubblica dell'Ecuador",
                "common": "Ecuador"
            },
            "jpn": {
                "official": "エクアドル共和国",
                "common": "エクアドル"
            },
            "nld": {
                "official": "Republiek Ecuador",
                "common": "Ecuador"
            },
            "por": {
                "official": "República do Equador",
                "common": "Equador"
            },
            "rus": {
                "official": "Республика Эквадор",
                "common": "Эквадор"
            },
            "spa": {
                "official": "República del Ecuador",
                "common": "Ecuador"
            },
            "fin": {
                "official": "Ecuadorin tasavalta",
                "common": "Ecuador"
            }
        },
        "latlng": [
            -2,
            -77.5
        ],
        "demonym": "Ecuadorean",
        "landlocked": false,
        "borders": [
            "COL",
            "PER"
        ],
        "area": 276841
    },
    {
        "name": {
            "common": "Egypt",
            "official": "Arab Republic of Egypt",
            "native": {
                "ara": {
                    "official": "جمهورية مصر العربية",
                    "common": "مصر"
                }
            }
        },
        "tld": [
            ".eg",
            ".مصر"
        ],
        "cca2": "EG",
        "ccn3": "818",
        "cca3": "EGY",
        "cioc": "EGY",
        "currency": [
            "EGP"
        ],
        "callingCode": [
            "20"
        ],
        "capital": "Cairo",
        "altSpellings": [
            "EG",
            "Arab Republic of Egypt"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "cym": {
                "official": "Arab Republic of Egypt",
                "common": "Yr Aifft"
            },
            "deu": {
                "official": "Arabische Republik Ägypten",
                "common": "Ägypten"
            },
            "fra": {
                "official": "République arabe d'Égypte",
                "common": "Égypte"
            },
            "hrv": {
                "official": "Arapska Republika Egipat",
                "common": "Egipat"
            },
            "ita": {
                "official": "Repubblica araba d'Egitto",
                "common": "Egitto"
            },
            "jpn": {
                "official": "エジプト·アラブ共和国",
                "common": "エジプト"
            },
            "nld": {
                "official": "Arabische Republiek Egypte",
                "common": "Egypte"
            },
            "por": {
                "official": "República Árabe do Egipto",
                "common": "Egito"
            },
            "rus": {
                "official": "Арабская Республика Египет",
                "common": "Египет"
            },
            "spa": {
                "official": "República Árabe de Egipto",
                "common": "Egipto"
            },
            "fin": {
                "official": "Egyptin arabitasavalta",
                "common": "Egypti"
            }
        },
        "latlng": [
            27,
            30
        ],
        "demonym": "Egyptian",
        "landlocked": false,
        "borders": [
            "ISR",
            "LBY",
            "SDN"
        ],
        "area": 1002450
    },
    {
        "name": {
            "common": "Eritrea",
            "official": "State of Eritrea",
            "native": {
                "ara": {
                    "official": "دولة إرتريا",
                    "common": "إرتريا‎"
                },
                "eng": {
                    "official": "State of Eritrea",
                    "common": "Eritrea"
                },
                "tir": {
                    "official": "ሃገረ ኤርትራ",
                    "common": "ኤርትራ"
                }
            }
        },
        "tld": [
            ".er"
        ],
        "cca2": "ER",
        "ccn3": "232",
        "cca3": "ERI",
        "cioc": "ERI",
        "currency": [
            "ERN"
        ],
        "callingCode": [
            "291"
        ],
        "capital": "Asmara",
        "altSpellings": [
            "ER",
            "State of Eritrea",
            "ሃገረ ኤርትራ",
            "Dawlat Iritriyá",
            "ʾErtrā",
            "Iritriyā"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "ara": "Arabic",
            "eng": "English",
            "tir": "Tigrinya"
        },
        "translations": {
            "cym": {
                "official": "State of Eritrea",
                "common": "Eritrea"
            },
            "deu": {
                "official": "Staat Eritrea",
                "common": "Eritrea"
            },
            "fra": {
                "official": "État d'Érythrée",
                "common": "Érythrée"
            },
            "hrv": {
                "official": "Država Eritreji",
                "common": "Eritreja"
            },
            "ita": {
                "official": "Stato di Eritrea",
                "common": "Eritrea"
            },
            "jpn": {
                "official": "エリトリア国",
                "common": "エリトリア"
            },
            "nld": {
                "official": "Staat Eritrea",
                "common": "Eritrea"
            },
            "por": {
                "official": "Estado da Eritreia",
                "common": "Eritreia"
            },
            "rus": {
                "official": "Государство Эритрея",
                "common": "Эритрея"
            },
            "spa": {
                "official": "Estado de Eritrea",
                "common": "Eritrea"
            },
            "fin": {
                "official": "Eritrean valtio",
                "common": "Eritrea"
            }
        },
        "latlng": [
            15,
            39
        ],
        "demonym": "Eritrean",
        "landlocked": false,
        "borders": [
            "DJI",
            "ETH",
            "SDN"
        ],
        "area": 117600
    },
    {
        "name": {
            "common": "Western Sahara",
            "official": "Sahrawi Arab Democratic Republic",
            "native": {
                "ber": {
                    "official": "Sahrawi Arab Democratic Republic",
                    "common": "Western Sahara"
                },
                "mey": {
                    "official": "الجمهورية العربية الصحراوية الديمقراطية",
                    "common": "الصحراء الغربية"
                },
                "spa": {
                    "official": "República Árabe Saharaui Democrática",
                    "common": "Sahara Occidental"
                }
            }
        },
        "tld": [
            ".eh"
        ],
        "cca2": "EH",
        "ccn3": "732",
        "cca3": "ESH",
        "cioc": "",
        "currency": [
            "MAD",
            "DZD",
            "MRO"
        ],
        "callingCode": [
            "212"
        ],
        "capital": "El Aaiún",
        "altSpellings": [
            "EH",
            "Taneẓroft Tutrimt"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ber": "Berber",
            "mey": "Hassaniya",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Demokratische Arabische Republik Sahara",
                "common": "Westsahara"
            },
            "fra": {
                "official": "République arabe sahraouie démocratique",
                "common": "Sahara Occidental"
            },
            "hrv": {
                "official": "Sahrawi Arab Demokratska Republika",
                "common": "Zapadna Sahara"
            },
            "ita": {
                "official": "Repubblica Araba Saharawi Democratica",
                "common": "Sahara Occidentale"
            },
            "jpn": {
                "official": "サハラアラブ民主共和国",
                "common": "西サハラ"
            },
            "nld": {
                "official": "Sahrawi Arabische Democratische Republiek",
                "common": "Westelijke Sahara"
            },
            "por": {
                "official": "República Árabe Saharaui Democrática",
                "common": "Saara Ocidental"
            },
            "rus": {
                "official": "Sahrawi Арабская Демократическая Республика",
                "common": "Западная Сахара"
            },
            "spa": {
                "official": "República Árabe Saharaui Democrática",
                "common": "Sahara Occidental"
            },
            "fin": {
                "official": "Länsi-Sahara",
                "common": "Länsi-Sahara"
            }
        },
        "latlng": [
            24.5,
            -13
        ],
        "demonym": "Sahrawi",
        "landlocked": false,
        "borders": [
            "DZA",
            "MRT",
            "MAR"
        ],
        "area": 266000
    },
    {
        "name": {
            "common": "Spain",
            "official": "Kingdom of Spain",
            "native": {
                "cat": {
                    "official": "Regne d'Espanya",
                    "common": "Espanya"
                },
                "eus": {
                    "official": "Espainiako Erresuma",
                    "common": "Espainia"
                },
                "glg": {
                    "official": "Reino de España",
                    "common": ""
                },
                "oci": {
                    "official": "Reialme d'Espanha",
                    "common": "Espanha"
                },
                "spa": {
                    "official": "Reino de España",
                    "common": "España"
                }
            }
        },
        "tld": [
            ".es"
        ],
        "cca2": "ES",
        "ccn3": "724",
        "cca3": "ESP",
        "cioc": "ESP",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "34"
        ],
        "capital": "Madrid",
        "altSpellings": [
            "ES",
            "Kingdom of Spain",
            "Reino de España"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "cat": "Catalan",
            "eus": "Basque",
            "glg": "Galician",
            "oci": "Occitan",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Königreich Spanien",
                "common": "Spanien"
            },
            "fra": {
                "official": "Royaume d'Espagne",
                "common": "Espagne"
            },
            "hrv": {
                "official": "Kraljevina Španjolska",
                "common": "Španjolska"
            },
            "ita": {
                "official": "Regno di Spagna",
                "common": "Spagna"
            },
            "jpn": {
                "official": "スペイン王国",
                "common": "スペイン"
            },
            "nld": {
                "official": "Koninkrijk Spanje",
                "common": "Spanje"
            },
            "por": {
                "official": "Reino de Espanha",
                "common": "Espanha"
            },
            "rus": {
                "official": "Королевство Испания",
                "common": "Испания"
            },
            "spa": {
                "official": "Reino de España",
                "common": "España"
            },
            "fin": {
                "official": "Espanjan kuningaskunta",
                "common": "Espanja"
            }
        },
        "latlng": [
            40,
            -4
        ],
        "demonym": "Spanish",
        "landlocked": false,
        "borders": [
            "AND",
            "FRA",
            "GIB",
            "PRT",
            "MAR"
        ],
        "area": 505992
    },
    {
        "name": {
            "common": "Estonia",
            "official": "Republic of Estonia",
            "native": {
                "est": {
                    "official": "Eesti Vabariik",
                    "common": "Eesti"
                }
            }
        },
        "tld": [
            ".ee"
        ],
        "cca2": "EE",
        "ccn3": "233",
        "cca3": "EST",
        "cioc": "EST",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "372"
        ],
        "capital": "Tallinn",
        "altSpellings": [
            "EE",
            "Eesti",
            "Republic of Estonia",
            "Eesti Vabariik"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "est": "Estonian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Estonia",
                "common": "Estonia"
            },
            "deu": {
                "official": "Republik Estland",
                "common": "Estland"
            },
            "fra": {
                "official": "République d'Estonie",
                "common": "Estonie"
            },
            "hrv": {
                "official": "Republika Estonija",
                "common": "Estonija"
            },
            "ita": {
                "official": "Repubblica di Estonia",
                "common": "Estonia"
            },
            "jpn": {
                "official": "エストニア共和国",
                "common": "エストニア"
            },
            "nld": {
                "official": "Republiek Estland",
                "common": "Estland"
            },
            "por": {
                "official": "República da Estónia",
                "common": "Estónia"
            },
            "rus": {
                "official": "Эстонская Республика",
                "common": "Эстония"
            },
            "spa": {
                "official": "República de Estonia",
                "common": "Estonia"
            },
            "fin": {
                "official": "Viron tasavalta",
                "common": "Viro"
            }
        },
        "latlng": [
            59,
            26
        ],
        "demonym": "Estonian",
        "landlocked": false,
        "borders": [
            "LVA",
            "RUS"
        ],
        "area": 45227
    },
    {
        "name": {
            "common": "Ethiopia",
            "official": "Federal Democratic Republic of Ethiopia",
            "native": {
                "amh": {
                    "official": "የኢትዮጵያ ፌዴራላዊ ዲሞክራሲያዊ ሪፐብሊክ",
                    "common": "ኢትዮጵያ"
                }
            }
        },
        "tld": [
            ".et"
        ],
        "cca2": "ET",
        "ccn3": "231",
        "cca3": "ETH",
        "cioc": "ETH",
        "currency": [
            "ETB"
        ],
        "callingCode": [
            "251"
        ],
        "capital": "Addis Ababa",
        "altSpellings": [
            "ET",
            "ʾĪtyōṗṗyā",
            "Federal Democratic Republic of Ethiopia",
            "የኢትዮጵያ ፌዴራላዊ ዲሞክራሲያዊ ሪፐብሊክ"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "amh": "Amharic"
        },
        "translations": {
            "cym": {
                "official": "Federal Democratic Republic of Ethiopia",
                "common": "Ethiopia"
            },
            "deu": {
                "official": "Demokratische Bundesrepublik Äthiopien",
                "common": "Äthiopien"
            },
            "fra": {
                "official": "République fédérale démocratique d'Éthiopie",
                "common": "Éthiopie"
            },
            "hrv": {
                "official": "Savezna Demokratska Republika Etiopija",
                "common": "Etiopija"
            },
            "ita": {
                "official": "Repubblica federale democratica di Etiopia",
                "common": "Etiopia"
            },
            "jpn": {
                "official": "エチオピア連邦民主共和国",
                "common": "エチオピア"
            },
            "nld": {
                "official": "Federale Democratische Republiek Ethiopië",
                "common": "Ethiopië"
            },
            "por": {
                "official": "República Federal Democrática da Etiópia",
                "common": "Etiópia"
            },
            "rus": {
                "official": "Федеративная Демократическая Республика Эфиопия",
                "common": "Эфиопия"
            },
            "spa": {
                "official": "República Democrática Federal de Etiopía",
                "common": "Etiopía"
            },
            "fin": {
                "official": "Etiopian demokraattinen liittotasavalta",
                "common": "Etiopia"
            }
        },
        "latlng": [
            8,
            38
        ],
        "demonym": "Ethiopian",
        "landlocked": true,
        "borders": [
            "DJI",
            "ERI",
            "KEN",
            "SOM",
            "SSD",
            "SDN"
        ],
        "area": 1104300
    },
    {
        "name": {
            "common": "Finland",
            "official": "Republic of Finland",
            "native": {
                "fin": {
                    "official": "Suomen tasavalta",
                    "common": "Suomi"
                },
                "swe": {
                    "official": "Republiken Finland",
                    "common": "Finland"
                }
            }
        },
        "tld": [
            ".fi"
        ],
        "cca2": "FI",
        "ccn3": "246",
        "cca3": "FIN",
        "cioc": "FIN",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "358"
        ],
        "capital": "Helsinki",
        "altSpellings": [
            "FI",
            "Suomi",
            "Republic of Finland",
            "Suomen tasavalta",
            "Republiken Finland"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "translations": {
            "deu": {
                "official": "Republik Finnland",
                "common": "Finnland"
            },
            "fra": {
                "official": "République de Finlande",
                "common": "Finlande"
            },
            "hrv": {
                "official": "Republika Finska",
                "common": "Finska"
            },
            "ita": {
                "official": "Repubblica di Finlandia",
                "common": "Finlandia"
            },
            "jpn": {
                "official": "フィンランド共和国",
                "common": "フィンランド"
            },
            "nld": {
                "official": "Republiek Finland",
                "common": "Finland"
            },
            "por": {
                "official": "República da Finlândia",
                "common": "Finlândia"
            },
            "rus": {
                "official": "Финляндская Республика",
                "common": "Финляндия"
            },
            "spa": {
                "official": "República de Finlandia",
                "common": "Finlandia"
            },
            "fin": {
                "official": "Suomen tasavalta",
                "common": "Suomi"
            }
        },
        "latlng": [
            64,
            26
        ],
        "demonym": "Finnish",
        "landlocked": false,
        "borders": [
            "NOR",
            "SWE",
            "RUS"
        ],
        "area": 338424
    },
    {
        "name": {
            "common": "Fiji",
            "official": "Republic of Fiji",
            "native": {
                "eng": {
                    "official": "Republic of Fiji",
                    "common": "Fiji"
                },
                "fij": {
                    "official": "Matanitu Tugalala o Viti",
                    "common": "Viti"
                },
                "hif": {
                    "official": "रिपब्लिक ऑफ फीजी",
                    "common": "फिजी"
                }
            }
        },
        "tld": [
            ".fj"
        ],
        "cca2": "FJ",
        "ccn3": "242",
        "cca3": "FJI",
        "cioc": "FIJ",
        "currency": [
            "FJD"
        ],
        "callingCode": [
            "679"
        ],
        "capital": "Suva",
        "altSpellings": [
            "FJ",
            "Viti",
            "Republic of Fiji",
            "Matanitu ko Viti",
            "Fijī Gaṇarājya"
        ],
        "region": "Oceania",
        "subregion": "Melanesia",
        "languages": {
            "eng": "English",
            "fij": "Fijian",
            "hif": "Fiji Hindi"
        },
        "translations": {
            "deu": {
                "official": "Republik Fidschi",
                "common": "Fidschi"
            },
            "fra": {
                "official": "République des Fidji",
                "common": "Fidji"
            },
            "hrv": {
                "official": "Republika Fidži",
                "common": "Fiđi"
            },
            "ita": {
                "official": "Repubblica di Figi",
                "common": "Figi"
            },
            "jpn": {
                "official": "フィジー共和国",
                "common": "フィジー"
            },
            "nld": {
                "official": "Republiek Fiji",
                "common": "Fiji"
            },
            "por": {
                "official": "República de Fiji",
                "common": "Fiji"
            },
            "rus": {
                "official": "Республика Фиджи",
                "common": "Фиджи"
            },
            "spa": {
                "official": "República de Fiji",
                "common": "Fiyi"
            },
            "fin": {
                "official": "Fidžin tasavalta",
                "common": "Fidži"
            }
        },
        "latlng": [
            -18,
            175
        ],
        "demonym": "Fijian",
        "landlocked": false,
        "borders": [],
        "area": 18272
    },
    {
        "name": {
            "common": "Falkland Islands",
            "official": "Falkland Islands",
            "native": {
                "eng": {
                    "official": "Falkland Islands",
                    "common": "Falkland Islands"
                }
            }
        },
        "tld": [
            ".fk"
        ],
        "cca2": "FK",
        "ccn3": "238",
        "cca3": "FLK",
        "cioc": "",
        "currency": [
            "FKP"
        ],
        "callingCode": [
            "500"
        ],
        "capital": "Stanley",
        "altSpellings": [
            "FK",
            "Islas Malvinas",
            "Falkland Islands (Malvinas)"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Falkland-Inseln",
                "common": "Falklandinseln"
            },
            "fra": {
                "official": "Îles Malouines",
                "common": "Îles Malouines"
            },
            "hrv": {
                "official": "Falklandski otoci",
                "common": "Falklandski Otoci"
            },
            "ita": {
                "official": "Isole Falkland",
                "common": "Isole Falkland o Isole Malvine"
            },
            "jpn": {
                "official": "フォークランド",
                "common": "フォークランド（マルビナス）諸島"
            },
            "nld": {
                "official": "Falkland eilanden",
                "common": "Falklandeilanden"
            },
            "por": {
                "official": "Ilhas Malvinas",
                "common": "Ilhas Malvinas"
            },
            "rus": {
                "official": "Фолклендские острова",
                "common": "Фолклендские острова"
            },
            "spa": {
                "official": "islas Malvinas",
                "common": "Islas Malvinas"
            },
            "fin": {
                "official": "Falkandinsaaret",
                "common": "Falkandinsaaret"
            }
        },
        "latlng": [
            -51.75,
            -59
        ],
        "demonym": "Falkland Islander",
        "landlocked": false,
        "borders": [],
        "area": 12173
    },
    {
        "name": {
            "common": "France",
            "official": "French Republic",
            "native": {
                "fra": {
                    "official": "République française",
                    "common": "France"
                }
            }
        },
        "tld": [
            ".fr"
        ],
        "cca2": "FR",
        "ccn3": "250",
        "cca3": "FRA",
        "cioc": "FRA",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "33"
        ],
        "capital": "Paris",
        "altSpellings": [
            "FR",
            "French Republic",
            "République française"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Französische Republik",
                "common": "Frankreich"
            },
            "fra": {
                "official": "République française",
                "common": "France"
            },
            "hrv": {
                "official": "Francuska Republika",
                "common": "Francuska"
            },
            "ita": {
                "official": "Repubblica francese",
                "common": "Francia"
            },
            "jpn": {
                "official": "フランス共和国",
                "common": "フランス"
            },
            "nld": {
                "official": "Franse Republiek",
                "common": "Frankrijk"
            },
            "por": {
                "official": "República Francesa",
                "common": "França"
            },
            "rus": {
                "official": "Французская Республика",
                "common": "Франция"
            },
            "spa": {
                "official": "República francés",
                "common": "Francia"
            },
            "fin": {
                "official": "Ranskan tasavalta",
                "common": "Ranska"
            }
        },
        "latlng": [
            46,
            2
        ],
        "demonym": "French",
        "landlocked": false,
        "borders": [
            "AND",
            "BEL",
            "DEU",
            "ITA",
            "LUX",
            "MCO",
            "ESP",
            "CHE"
        ],
        "area": 551695
    },
    {
        "name": {
            "common": "Faroe Islands",
            "official": "Faroe Islands",
            "native": {
                "dan": {
                    "official": "Færøerne",
                    "common": "Færøerne"
                },
                "fao": {
                    "official": "Føroyar",
                    "common": "Føroyar"
                }
            }
        },
        "tld": [
            ".fo"
        ],
        "cca2": "FO",
        "ccn3": "234",
        "cca3": "FRO",
        "cioc": "",
        "currency": [
            "DKK"
        ],
        "callingCode": [
            "298"
        ],
        "capital": "Tórshavn",
        "altSpellings": [
            "FO",
            "Føroyar",
            "Færøerne"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "dan": "Danish",
            "fao": "Faroese"
        },
        "translations": {
            "deu": {
                "official": "Färöer",
                "common": "Färöer-Inseln"
            },
            "fra": {
                "official": "Îles Féroé",
                "common": "Îles Féroé"
            },
            "hrv": {
                "official": "Farski Otoci",
                "common": "Farski Otoci"
            },
            "ita": {
                "official": "Isole Faroe",
                "common": "Isole Far Oer"
            },
            "jpn": {
                "official": "フェロー諸島",
                "common": "フェロー諸島"
            },
            "nld": {
                "official": "Faeröer",
                "common": "Faeröer"
            },
            "por": {
                "official": "Ilhas Faroe",
                "common": "Ilhas Faroé"
            },
            "rus": {
                "official": "Фарерские острова",
                "common": "Фарерские острова"
            },
            "spa": {
                "official": "Islas Feroe",
                "common": "Islas Faroe"
            },
            "fin": {
                "official": "Färsaaret",
                "common": "Färsaaret"
            }
        },
        "latlng": [
            62,
            -7
        ],
        "demonym": "Faroese",
        "landlocked": false,
        "borders": [],
        "area": 1393
    },
    {
        "name": {
            "common": "Micronesia",
            "official": "Federated States of Micronesia",
            "native": {
                "eng": {
                    "official": "Federated States of Micronesia",
                    "common": "Micronesia"
                }
            }
        },
        "tld": [
            ".fm"
        ],
        "cca2": "FM",
        "ccn3": "583",
        "cca3": "FSM",
        "cioc": "FSM",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "691"
        ],
        "capital": "Palikir",
        "altSpellings": [
            "FM",
            "Federated States of Micronesia",
            "Micronesia, Federated States of"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Föderierte Staaten von Mikronesien",
                "common": "Mikronesien"
            },
            "fra": {
                "official": "États fédérés de Micronésie",
                "common": "Micronésie"
            },
            "hrv": {
                "official": "Savezne Države Mikronezije",
                "common": "Mikronezija"
            },
            "ita": {
                "official": "Stati federati di Micronesia",
                "common": "Micronesia"
            },
            "jpn": {
                "official": "ミクロネシア連邦",
                "common": "ミクロネシア連邦"
            },
            "nld": {
                "official": "Federale Staten van Micronesia",
                "common": "Micronesië"
            },
            "por": {
                "official": "Estados Federados da Micronésia",
                "common": "Micronésia"
            },
            "rus": {
                "official": "Федеративные Штаты Микронезии",
                "common": "Федеративные Штаты Микронезии"
            },
            "spa": {
                "official": "Estados Federados de Micronesia",
                "common": "Micronesia"
            },
            "fin": {
                "official": "Mikronesian liittovaltio",
                "common": "Mikronesia"
            }
        },
        "latlng": [
            6.91666666,
            158.25
        ],
        "demonym": "Micronesian",
        "landlocked": false,
        "borders": [],
        "area": 702
    },
    {
        "name": {
            "common": "Gabon",
            "official": "Gabonese Republic",
            "native": {
                "fra": {
                    "official": "République gabonaise",
                    "common": "Gabon"
                }
            }
        },
        "tld": [
            ".ga"
        ],
        "cca2": "GA",
        "ccn3": "266",
        "cca3": "GAB",
        "cioc": "GAB",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "241"
        ],
        "capital": "Libreville",
        "altSpellings": [
            "GA",
            "Gabonese Republic",
            "République Gabonaise"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Gabunische Republik",
                "common": "Gabun"
            },
            "fra": {
                "official": "République gabonaise",
                "common": "Gabon"
            },
            "hrv": {
                "official": "Gabon Republika",
                "common": "Gabon"
            },
            "ita": {
                "official": "Repubblica gabonese",
                "common": "Gabon"
            },
            "jpn": {
                "official": "ガボン共和国",
                "common": "ガボン"
            },
            "nld": {
                "official": "Republiek Gabon",
                "common": "Gabon"
            },
            "por": {
                "official": "República do Gabão",
                "common": "Gabão"
            },
            "rus": {
                "official": "Габона Республика",
                "common": "Габон"
            },
            "spa": {
                "official": "República de Gabón",
                "common": "Gabón"
            },
            "fin": {
                "official": "Gabonin tasavalta",
                "common": "Gabon"
            }
        },
        "latlng": [
            -1,
            11.75
        ],
        "demonym": "Gabonese",
        "landlocked": false,
        "borders": [
            "CMR",
            "COG",
            "GNQ"
        ],
        "area": 267668
    },
    {
        "name": {
            "common": "United Kingdom",
            "official": "United Kingdom of Great Britain and Northern Ireland",
            "native": {
                "eng": {
                    "official": "United Kingdom of Great Britain and Northern Ireland",
                    "common": "United Kingdom"
                }
            }
        },
        "tld": [
            ".uk"
        ],
        "cca2": "GB",
        "ccn3": "826",
        "cca3": "GBR",
        "cioc": "GBR",
        "currency": [
            "GBP"
        ],
        "callingCode": [
            "44"
        ],
        "capital": "London",
        "altSpellings": [
            "GB",
            "UK",
            "Great Britain"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Vereinigtes Königreich Großbritannien und Nordirland",
                "common": "Vereinigtes Königreich"
            },
            "fra": {
                "official": "Royaume-Uni de Grande-Bretagne et d'Irlande du Nord",
                "common": "Royaume-Uni"
            },
            "hrv": {
                "official": "Ujedinjeno Kraljevstvo Velike Britanije i Sjeverne Irske",
                "common": "Ujedinjeno Kraljevstvo"
            },
            "ita": {
                "official": "Regno Unito di Gran Bretagna e Irlanda del Nord",
                "common": "Regno Unito"
            },
            "jpn": {
                "official": "グレート·ブリテンおよび北アイルランド連合王国",
                "common": "イギリス"
            },
            "nld": {
                "official": "Verenigd Koninkrijk van Groot-Brittannië en Noord-Ierland",
                "common": "Verenigd Koninkrijk"
            },
            "por": {
                "official": "Reino Unido da Grã-Bretanha e Irlanda do Norte",
                "common": "Reino Unido"
            },
            "rus": {
                "official": "Соединенное Королевство Великобритании и Северной Ирландии",
                "common": "Великобритания"
            },
            "spa": {
                "official": "Reino Unido de Gran Bretaña e Irlanda del Norte",
                "common": "Reino Unido"
            },
            "fin": {
                "official": "Ison-Britannian ja Pohjois-Irlannin yhdistynyt kuningaskunta",
                "common": "Yhdistynyt kuningaskunta"
            }
        },
        "latlng": [
            54,
            -2
        ],
        "demonym": "British",
        "landlocked": false,
        "borders": [
            "IRL"
        ],
        "area": 242900
    },
    {
        "name": {
            "common": "Georgia",
            "official": "Georgia",
            "native": {
                "kat": {
                    "official": "საქართველო",
                    "common": "საქართველო"
                }
            }
        },
        "tld": [
            ".ge"
        ],
        "cca2": "GE",
        "ccn3": "268",
        "cca3": "GEO",
        "cioc": "GEO",
        "currency": [
            "GEL"
        ],
        "callingCode": [
            "995"
        ],
        "capital": "Tbilisi",
        "altSpellings": [
            "GE",
            "Sakartvelo"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "kat": "Georgian"
        },
        "translations": {
            "deu": {
                "official": "Georgia",
                "common": "Georgien"
            },
            "fra": {
                "official": "République de Géorgie",
                "common": "Géorgie"
            },
            "hrv": {
                "official": "Gruzija",
                "common": "Gruzija"
            },
            "ita": {
                "official": "Georgia",
                "common": "Georgia"
            },
            "jpn": {
                "official": "グルジア",
                "common": "グルジア"
            },
            "nld": {
                "official": "Georgia",
                "common": "Georgië"
            },
            "por": {
                "official": "Georgia",
                "common": "Geórgia"
            },
            "rus": {
                "official": "Грузия",
                "common": "Грузия"
            },
            "spa": {
                "official": "Georgia",
                "common": "Georgia"
            },
            "fin": {
                "official": "Georgia",
                "common": "Georgia"
            }
        },
        "latlng": [
            42,
            43.5
        ],
        "demonym": "Georgian",
        "landlocked": false,
        "borders": [
            "ARM",
            "AZE",
            "RUS",
            "TUR"
        ],
        "area": 69700
    },
    {
        "name": {
            "common": "Guernsey",
            "official": "Bailiwick of Guernsey",
            "native": {
                "eng": {
                    "official": "Bailiwick of Guernsey",
                    "common": "Guernsey"
                },
                "fra": {
                    "official": "Bailliage de Guernesey",
                    "common": "Guernesey"
                },
                "nfr": {
                    "official": "Dgèrnésiais",
                    "common": "Dgèrnésiais"
                }
            }
        },
        "tld": [
            ".gg"
        ],
        "cca2": "GG",
        "ccn3": "831",
        "cca3": "GGY",
        "cioc": "",
        "currency": [
            "GBP"
        ],
        "callingCode": [
            "44"
        ],
        "capital": "St. Peter Port",
        "altSpellings": [
            "GG",
            "Bailiwick of Guernsey",
            "Bailliage de Guernesey"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "eng": "English",
            "fra": "French",
            "nfr": "Guernésiais"
        },
        "translations": {
            "deu": {
                "official": "Guernsey",
                "common": "Guernsey"
            },
            "fra": {
                "official": "Bailliage de Guernesey",
                "common": "Guernesey"
            },
            "hrv": {
                "official": "Struka Guernsey",
                "common": "Guernsey"
            },
            "ita": {
                "official": "Baliato di Guernsey",
                "common": "Guernsey"
            },
            "jpn": {
                "official": "ガーンジーの得意分野",
                "common": "ガーンジー"
            },
            "nld": {
                "official": "Baljuwschap Guernsey",
                "common": "Guernsey"
            },
            "por": {
                "official": "Bailiado de Guernsey",
                "common": "Guernsey"
            },
            "rus": {
                "official": "Коронное владение Гернси",
                "common": "Гернси"
            },
            "spa": {
                "official": "Bailía de Guernsey",
                "common": "Guernsey"
            },
            "fin": {
                "official": "Guernsey",
                "common": "Guernsey"
            }
        },
        "latlng": [
            49.46666666,
            -2.58333333
        ],
        "demonym": "Channel Islander",
        "landlocked": false,
        "borders": [],
        "area": 78
    },
    {
        "name": {
            "common": "Ghana",
            "official": "Republic of Ghana",
            "native": {
                "eng": {
                    "official": "Republic of Ghana",
                    "common": "Ghana"
                }
            }
        },
        "tld": [
            ".gh"
        ],
        "cca2": "GH",
        "ccn3": "288",
        "cca3": "GHA",
        "cioc": "GHA",
        "currency": [
            "GHS"
        ],
        "callingCode": [
            "233"
        ],
        "capital": "Accra",
        "altSpellings": [
            "GH"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Ghana",
                "common": "Ghana"
            },
            "fra": {
                "official": "République du Ghana",
                "common": "Ghana"
            },
            "hrv": {
                "official": "Republika Gana",
                "common": "Gana"
            },
            "ita": {
                "official": "Repubblica del Ghana",
                "common": "Ghana"
            },
            "jpn": {
                "official": "ガーナ共和国",
                "common": "ガーナ"
            },
            "nld": {
                "official": "Republiek Ghana",
                "common": "Ghana"
            },
            "por": {
                "official": "República do Gana",
                "common": "Gana"
            },
            "rus": {
                "official": "Республика Гана",
                "common": "Гана"
            },
            "spa": {
                "official": "República de Ghana",
                "common": "Ghana"
            },
            "fin": {
                "official": "Ghanan tasavalta",
                "common": "Ghana"
            }
        },
        "latlng": [
            8,
            -2
        ],
        "demonym": "Ghanaian",
        "landlocked": false,
        "borders": [
            "BFA",
            "CIV",
            "TGO"
        ],
        "area": 238533
    },
    {
        "name": {
            "common": "Gibraltar",
            "official": "Gibraltar",
            "native": {
                "eng": {
                    "official": "Gibraltar",
                    "common": "Gibraltar"
                }
            }
        },
        "tld": [
            ".gi"
        ],
        "cca2": "GI",
        "ccn3": "292",
        "cca3": "GIB",
        "cioc": "",
        "currency": [
            "GIP"
        ],
        "callingCode": [
            "350"
        ],
        "capital": "Gibraltar",
        "altSpellings": [
            "GI"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "fra": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "hrv": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "ita": {
                "official": "Gibilterra",
                "common": "Gibilterra"
            },
            "jpn": {
                "official": "ジブラルタル",
                "common": "ジブラルタル"
            },
            "nld": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "por": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "rus": {
                "official": "Гибралтар",
                "common": "Гибралтар"
            },
            "spa": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            },
            "fin": {
                "official": "Gibraltar",
                "common": "Gibraltar"
            }
        },
        "latlng": [
            36.13333333,
            -5.35
        ],
        "demonym": "Gibraltar",
        "landlocked": false,
        "borders": [
            "ESP"
        ],
        "area": 6
    },
    {
        "name": {
            "common": "Guinea",
            "official": "Republic of Guinea",
            "native": {
                "fra": {
                    "official": "République de Guinée",
                    "common": "Guinée"
                }
            }
        },
        "tld": [
            ".gn"
        ],
        "cca2": "GN",
        "ccn3": "324",
        "cca3": "GIN",
        "cioc": "GUI",
        "currency": [
            "GNF"
        ],
        "callingCode": [
            "224"
        ],
        "capital": "Conakry",
        "altSpellings": [
            "GN",
            "Republic of Guinea",
            "République de Guinée"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Guinea",
                "common": "Guinea"
            },
            "fra": {
                "official": "République de Guinée",
                "common": "Guinée"
            },
            "hrv": {
                "official": "Republika Gvineja",
                "common": "Gvineja"
            },
            "ita": {
                "official": "Repubblica di Guinea",
                "common": "Guinea"
            },
            "jpn": {
                "official": "ギニア共和国",
                "common": "ギニア"
            },
            "nld": {
                "official": "Republiek Guinee",
                "common": "Guinee"
            },
            "por": {
                "official": "República da Guiné",
                "common": "Guiné"
            },
            "rus": {
                "official": "Республика Гвинея",
                "common": "Гвинея"
            },
            "spa": {
                "official": "República de Guinea",
                "common": "Guinea"
            },
            "fin": {
                "official": "Guinean tasavalta",
                "common": "Guinea"
            }
        },
        "latlng": [
            11,
            -10
        ],
        "demonym": "Guinean",
        "landlocked": false,
        "borders": [
            "CIV",
            "GNB",
            "LBR",
            "MLI",
            "SEN",
            "SLE"
        ],
        "area": 245857
    },
    {
        "name": {
            "common": "Guadeloupe",
            "official": "Guadeloupe",
            "native": {
                "fra": {
                    "official": "Guadeloupe",
                    "common": "Guadeloupe"
                }
            }
        },
        "tld": [
            ".gp"
        ],
        "cca2": "GP",
        "ccn3": "312",
        "cca3": "GLP",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "590"
        ],
        "capital": "Basse-Terre",
        "altSpellings": [
            "GP",
            "Gwadloup"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Guadeloupe",
                "common": "Guadeloupe"
            },
            "fra": {
                "official": "Guadeloupe",
                "common": "Guadeloupe"
            },
            "hrv": {
                "official": "Gvadalupa",
                "common": "Gvadalupa"
            },
            "ita": {
                "official": "Guadeloupe",
                "common": "Guadeloupa"
            },
            "jpn": {
                "official": "グアドループ島",
                "common": "グアドループ"
            },
            "nld": {
                "official": "Guadeloupe",
                "common": "Guadeloupe"
            },
            "por": {
                "official": "Guadalupe",
                "common": "Guadalupe"
            },
            "rus": {
                "official": "Гваделупа",
                "common": "Гваделупа"
            },
            "spa": {
                "official": "Guadalupe",
                "common": "Guadalupe"
            },
            "fin": {
                "official": "Guadeloupen departmentti",
                "common": "Guadeloupe"
            }
        },
        "latlng": [
            16.25,
            -61.583333
        ],
        "demonym": "Guadeloupian",
        "landlocked": false,
        "borders": [],
        "area": 1628
    },
    {
        "name": {
            "common": "Gambia",
            "official": "Republic of the Gambia",
            "native": {
                "eng": {
                    "official": "Republic of the Gambia",
                    "common": "Gambia"
                }
            }
        },
        "tld": [
            ".gm"
        ],
        "cca2": "GM",
        "ccn3": "270",
        "cca3": "GMB",
        "cioc": "GAM",
        "currency": [
            "GMD"
        ],
        "callingCode": [
            "220"
        ],
        "capital": "Banjul",
        "altSpellings": [
            "GM",
            "Republic of the Gambia"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Gambia",
                "common": "Gambia"
            },
            "fra": {
                "official": "République de Gambie",
                "common": "Gambie"
            },
            "hrv": {
                "official": "Republika Gambija",
                "common": "Gambija"
            },
            "ita": {
                "official": "Repubblica del Gambia",
                "common": "Gambia"
            },
            "jpn": {
                "official": "ガンビア共和国",
                "common": "ガンビア"
            },
            "nld": {
                "official": "Republiek Gambia",
                "common": "Gambia"
            },
            "por": {
                "official": "República da Gâmbia",
                "common": "Gâmbia"
            },
            "rus": {
                "official": "Республика Гамбия",
                "common": "Гамбия"
            },
            "spa": {
                "official": "República de Gambia",
                "common": "Gambia"
            },
            "fin": {
                "official": "Gambian tasavalta",
                "common": "Gambia"
            }
        },
        "latlng": [
            13.46666666,
            -16.56666666
        ],
        "demonym": "Gambian",
        "landlocked": false,
        "borders": [
            "SEN"
        ],
        "area": 10689
    },
    {
        "name": {
            "common": "Guinea-Bissau",
            "official": "Republic of Guinea-Bissau",
            "native": {
                "por": {
                    "official": "República da Guiné-Bissau",
                    "common": "Guiné-Bissau"
                }
            }
        },
        "tld": [
            ".gw"
        ],
        "cca2": "GW",
        "ccn3": "624",
        "cca3": "GNB",
        "cioc": "GBS",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "245"
        ],
        "capital": "Bissau",
        "altSpellings": [
            "GW",
            "Republic of Guinea-Bissau",
            "República da Guiné-Bissau"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "deu": {
                "official": "Republik Guinea-Bissau",
                "common": "Guinea-Bissau"
            },
            "fra": {
                "official": "République de Guinée-Bissau",
                "common": "Guinée-Bissau"
            },
            "hrv": {
                "official": "Republika Gvineja Bisau",
                "common": "Gvineja Bisau"
            },
            "ita": {
                "official": "Repubblica di Guinea-Bissau",
                "common": "Guinea-Bissau"
            },
            "jpn": {
                "official": "ギニアビサウ共和国",
                "common": "ギニアビサウ"
            },
            "nld": {
                "official": "Republiek Guinee-Bissau",
                "common": "Guinee-Bissau"
            },
            "por": {
                "official": "República da Guiné-Bissau",
                "common": "Guiné-Bissau"
            },
            "rus": {
                "official": "Республика Гвинея -Бисау",
                "common": "Гвинея-Бисау"
            },
            "spa": {
                "official": "República de Guinea-Bissau",
                "common": "Guinea-Bisáu"
            },
            "fin": {
                "official": "Guinea-Bissaun tasavalta",
                "common": "Guinea-Bissau"
            }
        },
        "latlng": [
            12,
            -15
        ],
        "demonym": "Guinea-Bissauan",
        "landlocked": false,
        "borders": [
            "GIN",
            "SEN"
        ],
        "area": 36125
    },
    {
        "name": {
            "common": "Equatorial Guinea",
            "official": "Republic of Equatorial Guinea",
            "native": {
                "fra": {
                    "official": "République de la Guinée Équatoriale",
                    "common": "Guinée équatoriale"
                },
                "por": {
                    "official": "República da Guiné Equatorial",
                    "common": "Guiné Equatorial"
                },
                "spa": {
                    "official": "República de Guinea Ecuatorial",
                    "common": "Guinea Ecuatorial"
                }
            }
        },
        "tld": [
            ".gq"
        ],
        "cca2": "GQ",
        "ccn3": "226",
        "cca3": "GNQ",
        "cioc": "GEQ",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "240"
        ],
        "capital": "Malabo",
        "altSpellings": [
            "GQ",
            "Republic of Equatorial Guinea",
            "República de Guinea Ecuatorial",
            "République de Guinée équatoriale",
            "República da Guiné Equatorial"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "fra": "French",
            "por": "Portuguese",
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of Equatorial Guinea",
                "common": "Gini Gyhydeddol"
            },
            "deu": {
                "official": "Republik Äquatorialguinea",
                "common": "Äquatorialguinea"
            },
            "fra": {
                "official": "République de Guinée équatoriale",
                "common": "Guinée équatoriale"
            },
            "hrv": {
                "official": "Republika Ekvatorska Gvineja",
                "common": "Ekvatorijalna Gvineja"
            },
            "ita": {
                "official": "Repubblica della Guinea Equatoriale",
                "common": "Guinea Equatoriale"
            },
            "jpn": {
                "official": "赤道ギニア共和国",
                "common": "赤道ギニア"
            },
            "nld": {
                "official": "Republiek Equatoriaal-Guinea",
                "common": "Equatoriaal-Guinea"
            },
            "por": {
                "official": "República da Guiné Equatorial",
                "common": "Guiné Equatorial"
            },
            "rus": {
                "official": "Республика Экваториальная Гвинея",
                "common": "Экваториальная Гвинея"
            },
            "spa": {
                "official": "República de Guinea Ecuatorial",
                "common": "Guinea Ecuatorial"
            },
            "fin": {
                "official": "Päiväntasaajan Guinean tasavalta",
                "common": "Päiväntasaajan Guinea"
            }
        },
        "latlng": [
            2,
            10
        ],
        "demonym": "Equatorial Guinean",
        "landlocked": false,
        "borders": [
            "CMR",
            "GAB"
        ],
        "area": 28051
    },
    {
        "name": {
            "common": "Greece",
            "official": "Hellenic Republic",
            "native": {
                "ell": {
                    "official": "Ελληνική Δημοκρατία",
                    "common": "Ελλάδα"
                }
            }
        },
        "tld": [
            ".gr"
        ],
        "cca2": "GR",
        "ccn3": "300",
        "cca3": "GRC",
        "cioc": "GRE",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "30"
        ],
        "capital": "Athens",
        "altSpellings": [
            "GR",
            "Elláda",
            "Hellenic Republic",
            "Ελληνική Δημοκρατία"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "ell": "Greek"
        },
        "translations": {
            "deu": {
                "official": "Hellenische Republik",
                "common": "Griechenland"
            },
            "fra": {
                "official": "République hellénique",
                "common": "Grèce"
            },
            "hrv": {
                "official": "Helenska Republika",
                "common": "Grčka"
            },
            "ita": {
                "official": "Repubblica ellenica",
                "common": "Grecia"
            },
            "jpn": {
                "official": "ギリシャ共和国",
                "common": "ギリシャ"
            },
            "nld": {
                "official": "Helleense Republiek",
                "common": "Griekenland"
            },
            "por": {
                "official": "República Helénica",
                "common": "Grécia"
            },
            "rus": {
                "official": "Греческая Республика",
                "common": "Греция"
            },
            "spa": {
                "official": "República Helénica",
                "common": "Grecia"
            },
            "fin": {
                "official": "Helleenien tasavalta",
                "common": "Kreikka"
            }
        },
        "latlng": [
            39,
            22
        ],
        "demonym": "Greek",
        "landlocked": false,
        "borders": [
            "ALB",
            "BGR",
            "TUR",
            "MKD"
        ],
        "area": 131990
    },
    {
        "name": {
            "common": "Grenada",
            "official": "Grenada",
            "native": {
                "eng": {
                    "official": "Grenada",
                    "common": "Grenada"
                }
            }
        },
        "tld": [
            ".gd"
        ],
        "cca2": "GD",
        "ccn3": "308",
        "cca3": "GRD",
        "cioc": "GRN",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1473"
        ],
        "capital": "St. George's",
        "altSpellings": [
            "GD"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Grenada",
                "common": "Grenada"
            },
            "fra": {
                "official": "Grenade",
                "common": "Grenade"
            },
            "hrv": {
                "official": "Grenada",
                "common": "Grenada"
            },
            "ita": {
                "official": "Grenada",
                "common": "Grenada"
            },
            "jpn": {
                "official": "グレナダ",
                "common": "グレナダ"
            },
            "nld": {
                "official": "Grenada",
                "common": "Grenada"
            },
            "por": {
                "official": "Grenada",
                "common": "Granada"
            },
            "rus": {
                "official": "Гренада",
                "common": "Гренада"
            },
            "spa": {
                "official": "Granada",
                "common": "Grenada"
            },
            "fin": {
                "official": "Grenada",
                "common": "Grenada"
            }
        },
        "latlng": [
            12.11666666,
            -61.66666666
        ],
        "demonym": "Grenadian",
        "landlocked": false,
        "borders": [],
        "area": 344
    },
    {
        "name": {
            "common": "Greenland",
            "official": "Greenland",
            "native": {
                "kal": {
                    "official": "Kalaallit Nunaat",
                    "common": "Kalaallit Nunaat"
                }
            }
        },
        "tld": [
            ".gl"
        ],
        "cca2": "GL",
        "ccn3": "304",
        "cca3": "GRL",
        "cioc": "",
        "currency": [
            "DKK"
        ],
        "callingCode": [
            "299"
        ],
        "capital": "Nuuk",
        "altSpellings": [
            "GL",
            "Grønland"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "kal": "Greenlandic"
        },
        "translations": {
            "deu": {
                "official": "Grönland",
                "common": "Grönland"
            },
            "fra": {
                "official": "Groenland",
                "common": "Groenland"
            },
            "hrv": {
                "official": "Grenland",
                "common": "Grenland"
            },
            "ita": {
                "official": "Groenlandia",
                "common": "Groenlandia"
            },
            "jpn": {
                "official": "グリーンランド",
                "common": "グリーンランド"
            },
            "nld": {
                "official": "Groenland",
                "common": "Groenland"
            },
            "por": {
                "official": "Groenlândia",
                "common": "Gronelândia"
            },
            "rus": {
                "official": "Гренландия",
                "common": "Гренландия"
            },
            "spa": {
                "official": "Groenlandia",
                "common": "Groenlandia"
            },
            "fin": {
                "official": "Groönlanti",
                "common": "Groönlanti"
            }
        },
        "latlng": [
            72,
            -40
        ],
        "demonym": "Greenlandic",
        "landlocked": false,
        "borders": [],
        "area": 2166086
    },
    {
        "name": {
            "common": "Guatemala",
            "official": "Republic of Guatemala",
            "native": {
                "spa": {
                    "official": "República de Guatemala",
                    "common": "Guatemala"
                }
            }
        },
        "tld": [
            ".gt"
        ],
        "cca2": "GT",
        "ccn3": "320",
        "cca3": "GTM",
        "cioc": "GUA",
        "currency": [
            "GTQ"
        ],
        "callingCode": [
            "502"
        ],
        "capital": "Guatemala City",
        "altSpellings": [
            "GT"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Guatemala",
                "common": "Guatemala"
            },
            "fra": {
                "official": "République du Guatemala",
                "common": "Guatemala"
            },
            "hrv": {
                "official": "Republika Gvatemala",
                "common": "Gvatemala"
            },
            "ita": {
                "official": "Repubblica del Guatemala",
                "common": "Guatemala"
            },
            "jpn": {
                "official": "グアテマラ共和国",
                "common": "グアテマラ"
            },
            "nld": {
                "official": "Republiek Guatemala",
                "common": "Guatemala"
            },
            "por": {
                "official": "República da Guatemala",
                "common": "Guatemala"
            },
            "rus": {
                "official": "Республика Гватемала",
                "common": "Гватемала"
            },
            "spa": {
                "official": "República de Guatemala",
                "common": "Guatemala"
            },
            "fin": {
                "official": "Guatemalan tasavalta",
                "common": "Guatemala"
            }
        },
        "latlng": [
            15.5,
            -90.25
        ],
        "demonym": "Guatemalan",
        "landlocked": false,
        "borders": [
            "BLZ",
            "SLV",
            "HND",
            "MEX"
        ],
        "area": 108889
    },
    {
        "name": {
            "common": "French Guiana",
            "official": "Guiana",
            "native": {
                "fra": {
                    "official": "Guyanes",
                    "common": "Guyane française"
                }
            }
        },
        "tld": [
            ".gf"
        ],
        "cca2": "GF",
        "ccn3": "254",
        "cca3": "GUF",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "594"
        ],
        "capital": "Cayenne",
        "altSpellings": [
            "GF",
            "Guiana",
            "Guyane"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Guayana",
                "common": "Französisch Guyana"
            },
            "fra": {
                "official": "Guyane",
                "common": "Guyane"
            },
            "hrv": {
                "official": "Gijana",
                "common": "Francuska Gvajana"
            },
            "ita": {
                "official": "Guiana",
                "common": "Guyana francese"
            },
            "jpn": {
                "official": "ギアナ",
                "common": "フランス領ギアナ"
            },
            "nld": {
                "official": "Guyana",
                "common": "Frans-Guyana"
            },
            "por": {
                "official": "Guiana",
                "common": "Guiana Francesa"
            },
            "rus": {
                "official": "Гвиана",
                "common": "Французская Гвиана"
            },
            "spa": {
                "official": "Guayana",
                "common": "Guayana Francesa"
            },
            "fin": {
                "official": "Ranskan Guayana",
                "common": "Ranskan Guayana"
            }
        },
        "latlng": [
            4,
            -53
        ],
        "demonym": "",
        "landlocked": false,
        "borders": [
            "BRA",
            "SUR"
        ],
        "area": 83534
    },
    {
        "name": {
            "common": "Guam",
            "official": "Guam",
            "native": {
                "cha": {
                    "official": "Guåhån",
                    "common": "Guåhån"
                },
                "eng": {
                    "official": "Guam",
                    "common": "Guam"
                },
                "spa": {
                    "official": "Guam",
                    "common": "Guam"
                }
            }
        },
        "tld": [
            ".gu"
        ],
        "cca2": "GU",
        "ccn3": "316",
        "cca3": "GUM",
        "cioc": "GUM",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1671"
        ],
        "capital": "Hagåtña",
        "altSpellings": [
            "GU",
            "Guåhån"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "cha": "Chamorro",
            "eng": "English",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Guam",
                "common": "Guam"
            },
            "fra": {
                "official": "Guam",
                "common": "Guam"
            },
            "hrv": {
                "official": "Guam",
                "common": "Guam"
            },
            "ita": {
                "official": "Guam",
                "common": "Guam"
            },
            "jpn": {
                "official": "グアム",
                "common": "グアム"
            },
            "nld": {
                "official": "Guam",
                "common": "Guam"
            },
            "por": {
                "official": "Guam",
                "common": "Guam"
            },
            "rus": {
                "official": "Гуам",
                "common": "Гуам"
            },
            "spa": {
                "official": "Guam",
                "common": "Guam"
            },
            "fin": {
                "official": "Guam",
                "common": "Guam"
            }
        },
        "latlng": [
            13.46666666,
            144.78333333
        ],
        "demonym": "Guamanian",
        "landlocked": false,
        "borders": [],
        "area": 549
    },
    {
        "name": {
            "common": "Guyana",
            "official": "Co-operative Republic of Guyana",
            "native": {
                "eng": {
                    "official": "Co-operative Republic of Guyana",
                    "common": "Guyana"
                }
            }
        },
        "tld": [
            ".gy"
        ],
        "cca2": "GY",
        "ccn3": "328",
        "cca3": "GUY",
        "cioc": "GUY",
        "currency": [
            "GYD"
        ],
        "callingCode": [
            "592"
        ],
        "capital": "Georgetown",
        "altSpellings": [
            "GY",
            "Co-operative Republic of Guyana"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Kooperative Republik Guyana",
                "common": "Guyana"
            },
            "fra": {
                "official": "République coopérative de Guyana",
                "common": "Guyana"
            },
            "hrv": {
                "official": "Zadruga Republika Gvajana",
                "common": "Gvajana"
            },
            "ita": {
                "official": "Co -operative Republic of Guyana",
                "common": "Guyana"
            },
            "jpn": {
                "official": "ガイアナの協同共和国",
                "common": "ガイアナ"
            },
            "nld": {
                "official": "Coöperatieve Republiek Guyana",
                "common": "Guyana"
            },
            "por": {
                "official": "Co -operative República da Guiana",
                "common": "Guiana"
            },
            "rus": {
                "official": "Кооперативная Республика Гайана",
                "common": "Гайана"
            },
            "spa": {
                "official": "República Cooperativa de Guyana",
                "common": "Guyana"
            },
            "fin": {
                "official": "Guayanan osuustoiminnallinen tasavalta",
                "common": "Guayana"
            }
        },
        "latlng": [
            5,
            -59
        ],
        "demonym": "Guyanese",
        "landlocked": false,
        "borders": [
            "BRA",
            "SUR",
            "VEN"
        ],
        "area": 214969
    },
    {
        "name": {
            "common": "Hong Kong",
            "official": "Hong Kong Special Administrative Region of the People's Republic of China",
            "native": {
                "eng": {
                    "official": "Hong Kong Special Administrative Region of the People's Republic of China",
                    "common": "Hong Kong"
                },
                "zho": {
                    "official": "香港中国特别行政区的人民共和国",
                    "common": "香港"
                }
            }
        },
        "tld": [
            ".hk",
            ".香港"
        ],
        "cca2": "HK",
        "ccn3": "344",
        "cca3": "HKG",
        "cioc": "HKG",
        "currency": [
            "HKD"
        ],
        "callingCode": [
            "852"
        ],
        "capital": "City of Victoria",
        "altSpellings": [
            "HK"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "eng": "English",
            "zho": "Chinese"
        },
        "translations": {
            "deu": {
                "official": "Sonderverwaltungszone der Volksrepublik China",
                "common": "Hongkong"
            },
            "fra": {
                "official": "Région administrative spéciale de Hong Kong de la République populaire de Chine",
                "common": "Hong Kong"
            },
            "hrv": {
                "official": "Hong Kong Posebnog upravnog područjaNarodne Republike Kine",
                "common": "Hong Kong"
            },
            "ita": {
                "official": "Hong Kong Regione amministrativa speciale della Repubblica Popolare Cinese",
                "common": "Hong Kong"
            },
            "jpn": {
                "official": "中華人民共和国香港特別行政区",
                "common": "香港"
            },
            "nld": {
                "official": "Hong Kong Speciale Administratieve Regio van de Volksrepubliek China",
                "common": "Hongkong"
            },
            "por": {
                "official": "Hong Kong Região Administrativa Especial da República Popular da China",
                "common": "Hong Kong"
            },
            "rus": {
                "official": "Hong Kong Специальный административный район Китайской Народной Республики Китая",
                "common": "Гонконг"
            },
            "spa": {
                "official": "Hong Kong Región Administrativa Especial de la República Popular China",
                "common": "Hong Kong"
            },
            "fin": {
                "official": "Hong Kongin erityishallintoalue",
                "common": "Hongkong"
            }
        },
        "latlng": [
            22.267,
            114.188
        ],
        "demonym": "Hong Konger",
        "landlocked": false,
        "borders": [
            "CHN"
        ],
        "area": 1104
    },
    {
        "name": {
            "common": "Heard Island and McDonald Islands",
            "official": "Heard Island and McDonald Islands",
            "native": {
                "eng": {
                    "official": "Heard Island and McDonald Islands",
                    "common": "Heard Island and McDonald Islands"
                }
            }
        },
        "tld": [
            ".hm",
            ".aq"
        ],
        "cca2": "HM",
        "ccn3": "334",
        "cca3": "HMD",
        "cioc": "",
        "currency": [
            "AUD"
        ],
        "callingCode": [],
        "capital": "",
        "altSpellings": [
            "HM",
            "Heard Island and McDonald Mcdonald Islands"
        ],
        "region": "",
        "subregion": "",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Heard und McDonaldinseln",
                "common": "Heard und die McDonaldinseln"
            },
            "fra": {
                "official": "Des îles Heard et McDonald",
                "common": "Îles Heard-et-MacDonald"
            },
            "hrv": {
                "official": "Otok Heard i otočje McDonald",
                "common": "Otok Heard i otočje McDonald"
            },
            "ita": {
                "official": "Isole Heard e McDonald",
                "common": "Isole Heard e McDonald"
            },
            "jpn": {
                "official": "ハード島とマクドナルド諸島",
                "common": "ハード島とマクドナルド諸島"
            },
            "nld": {
                "official": "Heard en McDonaldeilanden",
                "common": "Heard-en McDonaldeilanden"
            },
            "por": {
                "official": "Ilha Heard e Ilhas McDonald",
                "common": "Ilha Heard e Ilhas McDonald"
            },
            "rus": {
                "official": "Остров Херд и острова Макдональд",
                "common": "Остров Херд и острова Макдональд"
            },
            "spa": {
                "official": "Islas Heard y McDonald",
                "common": "Islas Heard y McDonald"
            },
            "fin": {
                "official": "Heard ja McDonaldinsaaret",
                "common": "Heard ja McDonaldinsaaret"
            }
        },
        "latlng": [
            -53.1,
            72.51666666
        ],
        "demonym": "Heard and McDonald Islander",
        "landlocked": false,
        "borders": [],
        "area": 412
    },
    {
        "name": {
            "common": "Honduras",
            "official": "Republic of Honduras",
            "native": {
                "spa": {
                    "official": "República de Honduras",
                    "common": "Honduras"
                }
            }
        },
        "tld": [
            ".hn"
        ],
        "cca2": "HN",
        "ccn3": "340",
        "cca3": "HND",
        "cioc": "HON",
        "currency": [
            "HNL"
        ],
        "callingCode": [
            "504"
        ],
        "capital": "Tegucigalpa",
        "altSpellings": [
            "HN",
            "Republic of Honduras",
            "República de Honduras"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Honduras",
                "common": "Honduras"
            },
            "fra": {
                "official": "République du Honduras",
                "common": "Honduras"
            },
            "hrv": {
                "official": "Republika Honduras",
                "common": "Honduras"
            },
            "ita": {
                "official": "Repubblica di Honduras",
                "common": "Honduras"
            },
            "jpn": {
                "official": "ホンジュラス共和国",
                "common": "ホンジュラス"
            },
            "nld": {
                "official": "Republiek Honduras",
                "common": "Honduras"
            },
            "por": {
                "official": "República de Honduras",
                "common": "Honduras"
            },
            "rus": {
                "official": "Республика Гондурас",
                "common": "Гондурас"
            },
            "spa": {
                "official": "República de Honduras",
                "common": "Honduras"
            },
            "fin": {
                "official": "Hondurasin tasavalta",
                "common": "Honduras"
            }
        },
        "latlng": [
            15,
            -86.5
        ],
        "demonym": "Honduran",
        "landlocked": false,
        "borders": [
            "GTM",
            "SLV",
            "NIC"
        ],
        "area": 112492
    },
    {
        "name": {
            "common": "Croatia",
            "official": "Republic of Croatia",
            "native": {
                "hrv": {
                    "official": "Republika Hrvatska",
                    "common": "Hrvatska"
                }
            }
        },
        "tld": [
            ".hr"
        ],
        "cca2": "HR",
        "ccn3": "191",
        "cca3": "HRV",
        "cioc": "CRO",
        "currency": [
            "HRK"
        ],
        "callingCode": [
            "385"
        ],
        "capital": "Zagreb",
        "altSpellings": [
            "HR",
            "Hrvatska",
            "Republic of Croatia",
            "Republika Hrvatska"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "hrv": "Croatian"
        },
        "translations": {
            "cym": {
                "official": "Republic of Croatia",
                "common": "Croatia"
            },
            "deu": {
                "official": "Republik Kroatien",
                "common": "Kroatien"
            },
            "fra": {
                "official": "République de Croatie",
                "common": "Croatie"
            },
            "hrv": {
                "official": "Republika Hrvatska",
                "common": "Hrvatska"
            },
            "ita": {
                "official": "Repubblica di Croazia",
                "common": "Croazia"
            },
            "jpn": {
                "official": "クロアチア共和国",
                "common": "クロアチア"
            },
            "nld": {
                "official": "Republiek Kroatië",
                "common": "Kroatië"
            },
            "por": {
                "official": "República da Croácia",
                "common": "Croácia"
            },
            "rus": {
                "official": "Республика Хорватия",
                "common": "Хорватия"
            },
            "spa": {
                "official": "República de Croacia",
                "common": "Croacia"
            },
            "fin": {
                "official": "Kroatian tasavalta",
                "common": "Kroatia"
            }
        },
        "latlng": [
            45.16666666,
            15.5
        ],
        "demonym": "Croatian",
        "landlocked": false,
        "borders": [
            "BIH",
            "HUN",
            "MNE",
            "SRB",
            "SVN"
        ],
        "area": 56594
    },
    {
        "name": {
            "common": "Haiti",
            "official": "Republic of Haiti",
            "native": {
                "fra": {
                    "official": "République d'Haïti",
                    "common": "Haïti"
                },
                "hat": {
                    "official": "Repiblik Ayiti",
                    "common": "Ayiti"
                }
            }
        },
        "tld": [
            ".ht"
        ],
        "cca2": "HT",
        "ccn3": "332",
        "cca3": "HTI",
        "cioc": "HAI",
        "currency": [
            "HTG",
            "USD"
        ],
        "callingCode": [
            "509"
        ],
        "capital": "Port-au-Prince",
        "altSpellings": [
            "HT",
            "Republic of Haiti",
            "République d'Haïti",
            "Repiblik Ayiti"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "fra": "French",
            "hat": "Haitian Creole"
        },
        "translations": {
            "deu": {
                "official": "Republik Haiti",
                "common": "Haiti"
            },
            "fra": {
                "official": "République d'Haïti",
                "common": "Haïti"
            },
            "hrv": {
                "official": "Republika Haiti",
                "common": "Haiti"
            },
            "ita": {
                "official": "Repubblica di Haiti",
                "common": "Haiti"
            },
            "jpn": {
                "official": "ハイチ共和国",
                "common": "ハイチ"
            },
            "nld": {
                "official": "Republiek Haïti",
                "common": "Haïti"
            },
            "por": {
                "official": "República do Haiti",
                "common": "Haiti"
            },
            "rus": {
                "official": "Республика Гаити",
                "common": "Гаити"
            },
            "spa": {
                "official": "República de Haití",
                "common": "Haiti"
            },
            "fin": {
                "official": "Haitin tasavalta",
                "common": "Haiti"
            }
        },
        "latlng": [
            19,
            -72.41666666
        ],
        "demonym": "Haitian",
        "landlocked": false,
        "borders": [
            "DOM"
        ],
        "area": 27750
    },
    {
        "name": {
            "common": "Hungary",
            "official": "Hungary",
            "native": {
                "hun": {
                    "official": "Magyarország",
                    "common": "Magyarország"
                }
            }
        },
        "tld": [
            ".hu"
        ],
        "cca2": "HU",
        "ccn3": "348",
        "cca3": "HUN",
        "cioc": "HUN",
        "currency": [
            "HUF"
        ],
        "callingCode": [
            "36"
        ],
        "capital": "Budapest",
        "altSpellings": [
            "HU"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "hun": "Hungarian"
        },
        "translations": {
            "deu": {
                "official": "Ungarn",
                "common": "Ungarn"
            },
            "fra": {
                "official": "Hongrie",
                "common": "Hongrie"
            },
            "hrv": {
                "official": "Madžarska",
                "common": "Mađarska"
            },
            "ita": {
                "official": "Ungheria",
                "common": "Ungheria"
            },
            "jpn": {
                "official": "ハンガリー",
                "common": "ハンガリー"
            },
            "nld": {
                "official": "Hongarije",
                "common": "Hongarije"
            },
            "por": {
                "official": "Hungria",
                "common": "Hungria"
            },
            "rus": {
                "official": "Венгрия",
                "common": "Венгрия"
            },
            "spa": {
                "official": "Hungría",
                "common": "Hungría"
            },
            "fin": {
                "official": "Unkari",
                "common": "Unkari"
            }
        },
        "latlng": [
            47,
            20
        ],
        "demonym": "Hungarian",
        "landlocked": true,
        "borders": [
            "AUT",
            "HRV",
            "ROU",
            "SRB",
            "SVK",
            "SVN",
            "UKR"
        ],
        "area": 93028
    },
    {
        "name": {
            "common": "Indonesia",
            "official": "Republic of Indonesia",
            "native": {
                "ind": {
                    "official": "Republik Indonesia",
                    "common": "Indonesia"
                }
            }
        },
        "tld": [
            ".id"
        ],
        "cca2": "ID",
        "ccn3": "360",
        "cca3": "IDN",
        "cioc": "INA",
        "currency": [
            "IDR"
        ],
        "callingCode": [
            "62"
        ],
        "capital": "Jakarta",
        "altSpellings": [
            "ID",
            "Republic of Indonesia",
            "Republik Indonesia"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "ind": "Indonesian"
        },
        "translations": {
            "deu": {
                "official": "Republik Indonesien",
                "common": "Indonesien"
            },
            "fra": {
                "official": "République d'Indonésie",
                "common": "Indonésie"
            },
            "hrv": {
                "official": "Republika Indonezija",
                "common": "Indonezija"
            },
            "ita": {
                "official": "Repubblica di Indonesia",
                "common": "Indonesia"
            },
            "jpn": {
                "official": "インドネシア共和国",
                "common": "インドネシア"
            },
            "nld": {
                "official": "Republiek Indonesië",
                "common": "Indonesië"
            },
            "por": {
                "official": "República da Indonésia",
                "common": "Indonésia"
            },
            "rus": {
                "official": "Республика Индонезия",
                "common": "Индонезия"
            },
            "spa": {
                "official": "República de Indonesia",
                "common": "Indonesia"
            },
            "fin": {
                "official": "Indonesian tasavalta",
                "common": "Indonesia"
            }
        },
        "latlng": [
            -5,
            120
        ],
        "demonym": "Indonesian",
        "landlocked": false,
        "borders": [
            "TLS",
            "MYS",
            "PNG"
        ],
        "area": 1904569
    },
    {
        "name": {
            "common": "Isle of Man",
            "official": "Isle of Man",
            "native": {
                "eng": {
                    "official": "Isle of Man",
                    "common": "Isle of Man"
                },
                "glv": {
                    "official": "Ellan Vannin or Mannin",
                    "common": "Mannin"
                }
            }
        },
        "tld": [
            ".im"
        ],
        "cca2": "IM",
        "ccn3": "833",
        "cca3": "IMN",
        "cioc": "",
        "currency": [
            "GBP"
        ],
        "callingCode": [
            "44"
        ],
        "capital": "Douglas",
        "altSpellings": [
            "IM",
            "Ellan Vannin",
            "Mann",
            "Mannin"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "eng": "English",
            "glv": "Manx"
        },
        "translations": {
            "deu": {
                "official": "Isle of Man",
                "common": "Insel Man"
            },
            "fra": {
                "official": "Isle of Man",
                "common": "Île de Man"
            },
            "hrv": {
                "official": "Mana ostrvo",
                "common": "Otok Man"
            },
            "ita": {
                "official": "Isola di Man",
                "common": "Isola di Man"
            },
            "jpn": {
                "official": "マン島",
                "common": "マン島"
            },
            "nld": {
                "official": "Isle of Man",
                "common": "Isle of Man"
            },
            "por": {
                "official": "Isle of Man",
                "common": "Ilha de Man"
            },
            "rus": {
                "official": "Остров Мэн",
                "common": "Остров Мэн"
            },
            "spa": {
                "official": "Isla de Man",
                "common": "Isla de Man"
            },
            "fin": {
                "official": "Mansaari",
                "common": "Mansaari"
            }
        },
        "latlng": [
            54.25,
            -4.5
        ],
        "demonym": "Manx",
        "landlocked": false,
        "borders": [],
        "area": 572
    },
    {
        "name": {
            "common": "India",
            "official": "Republic of India",
            "native": {
                "eng": {
                    "official": "Republic of India",
                    "common": "India"
                },
                "hin": {
                    "official": "भारत गणराज्य",
                    "common": "भारत"
                },
                "tam": {
                    "official": "இந்தியக் குடியரசு",
                    "common": "இந்தியா"
                }
            }
        },
        "tld": [
            ".in"
        ],
        "cca2": "IN",
        "ccn3": "356",
        "cca3": "IND",
        "cioc": "IND",
        "currency": [
            "INR"
        ],
        "callingCode": [
            "91"
        ],
        "capital": "New Delhi",
        "altSpellings": [
            "IN",
            "Bhārat",
            "Republic of India",
            "Bharat Ganrajya",
            "இந்தியா"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "eng": "English",
            "hin": "Hindi",
            "tam": "Tamil"
        },
        "translations": {
            "deu": {
                "official": "Republik Indien",
                "common": "Indien"
            },
            "fra": {
                "official": "République de l'Inde",
                "common": "Inde"
            },
            "hrv": {
                "official": "Republika Indija",
                "common": "Indija"
            },
            "ita": {
                "official": "Repubblica dell'India",
                "common": "India"
            },
            "jpn": {
                "official": "インド共和国",
                "common": "インド"
            },
            "nld": {
                "official": "Republiek India",
                "common": "India"
            },
            "por": {
                "official": "República da Índia",
                "common": "Índia"
            },
            "rus": {
                "official": "Республика Индия",
                "common": "Индия"
            },
            "spa": {
                "official": "República de la India",
                "common": "India"
            },
            "fin": {
                "official": "Intian tasavalta",
                "common": "Intia"
            }
        },
        "latlng": [
            20,
            77
        ],
        "demonym": "Indian",
        "landlocked": false,
        "borders": [
            "AFG",
            "BGD",
            "BTN",
            "MMR",
            "CHN",
            "NPL",
            "PAK",
            "LKA"
        ],
        "area": 3287590
    },
    {
        "name": {
            "common": "British Indian Ocean Territory",
            "official": "British Indian Ocean Territory",
            "native": {
                "eng": {
                    "official": "British Indian Ocean Territory",
                    "common": "British Indian Ocean Territory"
                }
            }
        },
        "tld": [
            ".io"
        ],
        "cca2": "IO",
        "ccn3": "086",
        "cca3": "IOT",
        "cioc": "",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "246"
        ],
        "capital": "Diego Garcia",
        "altSpellings": [
            "IO"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "cym": {
                "official": "British Indian Ocean Territory",
                "common": "Tiriogaeth Brydeinig Cefnfor India"
            },
            "deu": {
                "official": "Britisch-Indischer Ozean",
                "common": "Britisches Territorium im Indischen Ozean"
            },
            "fra": {
                "official": "Territoire britannique de l' océan Indien",
                "common": "Territoire britannique de l'océan Indien"
            },
            "hrv": {
                "official": "British Indian Ocean Territory",
                "common": "Britanski Indijskooceanski teritorij"
            },
            "ita": {
                "official": "Territorio britannico dell'Oceano Indiano",
                "common": "Territorio britannico dell'oceano indiano"
            },
            "jpn": {
                "official": "イギリス領インド洋地域",
                "common": "イギリス領インド洋地域"
            },
            "nld": {
                "official": "Brits Indische Oceaan Territorium",
                "common": "Britse Gebieden in de Indische Oceaan"
            },
            "por": {
                "official": "British Indian Ocean Territory",
                "common": "Território Britânico do Oceano Índico"
            },
            "rus": {
                "official": "Британская территория Индийского океана",
                "common": "Британская территория в Индийском океане"
            },
            "spa": {
                "official": "Territorio Británico del Océano Índico",
                "common": "Territorio Británico del Océano Índico"
            },
            "fin": {
                "official": "Brittiläinen Intian valtameren alue",
                "common": "Brittiläinen Intian valtameren alue"
            }
        },
        "latlng": [
            -6,
            71.5
        ],
        "demonym": "Indian",
        "landlocked": false,
        "borders": [],
        "area": 60
    },
    {
        "name": {
            "common": "Ireland",
            "official": "Republic of Ireland",
            "native": {
                "eng": {
                    "official": "Republic of Ireland",
                    "common": "Ireland"
                },
                "gle": {
                    "official": "Poblacht na hÉireann",
                    "common": "Éire"
                }
            }
        },
        "tld": [
            ".ie"
        ],
        "cca2": "IE",
        "ccn3": "372",
        "cca3": "IRL",
        "cioc": "IRL",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "353"
        ],
        "capital": "Dublin",
        "altSpellings": [
            "IE",
            "Éire",
            "Republic of Ireland",
            "Poblacht na hÉireann"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "eng": "English",
            "gle": "Irish"
        },
        "translations": {
            "deu": {
                "official": "Republik Irland",
                "common": "Irland"
            },
            "fra": {
                "official": "République d'Irlande",
                "common": "Irlande"
            },
            "hrv": {
                "official": "Republika Irska",
                "common": "Irska"
            },
            "ita": {
                "official": "Repubblica d'Irlanda",
                "common": "Irlanda"
            },
            "jpn": {
                "official": "アイルランド共和国",
                "common": "アイルランド"
            },
            "nld": {
                "official": "Republic of Ireland",
                "common": "Ierland"
            },
            "por": {
                "official": "República da Irlanda",
                "common": "Irlanda"
            },
            "rus": {
                "official": "Ирландия",
                "common": "Ирландия"
            },
            "spa": {
                "official": "República de Irlanda",
                "common": "Irlanda"
            },
            "fin": {
                "official": "Irlannin tasavalta",
                "common": "Irlanti"
            }
        },
        "latlng": [
            53,
            -8
        ],
        "demonym": "Irish",
        "landlocked": false,
        "borders": [
            "GBR"
        ],
        "area": 70273
    },
    {
        "name": {
            "common": "Iran",
            "official": "Islamic Republic of Iran",
            "native": {
                "fas": {
                    "official": "جمهوری اسلامی ایران",
                    "common": "ایران"
                }
            }
        },
        "tld": [
            ".ir",
            "ایران."
        ],
        "cca2": "IR",
        "ccn3": "364",
        "cca3": "IRN",
        "cioc": "IRI",
        "currency": [
            "IRR"
        ],
        "callingCode": [
            "98"
        ],
        "capital": "Tehran",
        "altSpellings": [
            "IR",
            "Islamic Republic of Iran",
            "Iran, Islamic Republic of",
            "Jomhuri-ye Eslāmi-ye Irān"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "fas": "Persian"
        },
        "translations": {
            "deu": {
                "official": "Islamische Republik Iran",
                "common": "Iran"
            },
            "fra": {
                "official": "République islamique d'Iran",
                "common": "Iran"
            },
            "hrv": {
                "official": "Islamska Republika Iran",
                "common": "Iran"
            },
            "jpn": {
                "official": "イラン·イスラム共和国",
                "common": "イラン・イスラム共和国"
            },
            "nld": {
                "official": "Islamitische Republiek Iran",
                "common": "Iran"
            },
            "por": {
                "official": "República Islâmica do Irã",
                "common": "Irão"
            },
            "rus": {
                "official": "Исламская Республика Иран",
                "common": "Иран"
            },
            "spa": {
                "official": "República Islámica de Irán",
                "common": "Iran"
            },
            "fin": {
                "official": "Iranin islamilainen tasavalta",
                "common": "Iran"
            }
        },
        "latlng": [
            32,
            53
        ],
        "demonym": "Iranian",
        "landlocked": false,
        "borders": [
            "AFG",
            "ARM",
            "AZE",
            "IRQ",
            "PAK",
            "TUR",
            "TKM"
        ],
        "area": 1648195
    },
    {
        "name": {
            "common": "Iraq",
            "official": "Republic of Iraq",
            "native": {
                "ara": {
                    "official": "جمهورية العراق",
                    "common": "العراق"
                },
                "arc": {
                    "official": "ܩܘܼܛܢܵܐ ܐܝܼܪܲܩ",
                    "common": "ܩܘܼܛܢܵܐ"
                },
                "ckb": {
                    "official": "کۆماری عێراق",
                    "common": "کۆماری"
                }
            }
        },
        "tld": [
            ".iq"
        ],
        "cca2": "IQ",
        "ccn3": "368",
        "cca3": "IRQ",
        "cioc": "IRQ",
        "currency": [
            "IQD"
        ],
        "callingCode": [
            "964"
        ],
        "capital": "Baghdad",
        "altSpellings": [
            "IQ",
            "Republic of Iraq",
            "Jumhūriyyat al-‘Irāq"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic",
            "arc": "Aramaic",
            "ckb": "Sorani"
        },
        "translations": {
            "deu": {
                "official": "Republik Irak",
                "common": "Irak"
            },
            "fra": {
                "official": "République d'Irak",
                "common": "Irak"
            },
            "hrv": {
                "official": "Republika Irak",
                "common": "Irak"
            },
            "ita": {
                "official": "Repubblica dell'Iraq",
                "common": "Iraq"
            },
            "jpn": {
                "official": "イラク共和国",
                "common": "イラク"
            },
            "nld": {
                "official": "Republiek Irak",
                "common": "Irak"
            },
            "por": {
                "official": "República do Iraque",
                "common": "Iraque"
            },
            "rus": {
                "official": "Республика Ирак",
                "common": "Ирак"
            },
            "spa": {
                "official": "República de Irak",
                "common": "Irak"
            },
            "fin": {
                "official": "Irakin tasavalta",
                "common": "Irak"
            }
        },
        "latlng": [
            33,
            44
        ],
        "demonym": "Iraqi",
        "landlocked": false,
        "borders": [
            "IRN",
            "JOR",
            "KWT",
            "SAU",
            "SYR",
            "TUR"
        ],
        "area": 438317
    },
    {
        "name": {
            "common": "Iceland",
            "official": "Iceland",
            "native": {
                "isl": {
                    "official": "Ísland",
                    "common": "Ísland"
                }
            }
        },
        "tld": [
            ".is"
        ],
        "cca2": "IS",
        "ccn3": "352",
        "cca3": "ISL",
        "cioc": "ISL",
        "currency": [
            "ISK"
        ],
        "callingCode": [
            "354"
        ],
        "capital": "Reykjavik",
        "altSpellings": [
            "IS",
            "Island",
            "Republic of Iceland",
            "Lýðveldið Ísland"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "isl": "Icelandic"
        },
        "translations": {
            "deu": {
                "official": "Island",
                "common": "Island"
            },
            "fra": {
                "official": "République d'Islande",
                "common": "Islande"
            },
            "hrv": {
                "official": "Island",
                "common": "Island"
            },
            "ita": {
                "official": "Islanda",
                "common": "Islanda"
            },
            "jpn": {
                "official": "アイスランド",
                "common": "アイスランド"
            },
            "nld": {
                "official": "IJsland",
                "common": "IJsland"
            },
            "por": {
                "official": "Islândia",
                "common": "Islândia"
            },
            "rus": {
                "official": "Исландия",
                "common": "Исландия"
            },
            "spa": {
                "official": "Islandia",
                "common": "Islandia"
            },
            "fin": {
                "official": "Islanti",
                "common": "Islanti"
            }
        },
        "latlng": [
            65,
            -18
        ],
        "demonym": "Icelander",
        "landlocked": false,
        "borders": [],
        "area": 103000
    },
    {
        "name": {
            "common": "Israel",
            "official": "State of Israel",
            "native": {
                "ara": {
                    "official": "دولة إسرائيل",
                    "common": "إسرائيل"
                },
                "heb": {
                    "official": "מדינת ישראל",
                    "common": "ישראל"
                }
            }
        },
        "tld": [
            ".il"
        ],
        "cca2": "IL",
        "ccn3": "376",
        "cca3": "ISR",
        "cioc": "ISR",
        "currency": [
            "ILS"
        ],
        "callingCode": [
            "972"
        ],
        "capital": "Jerusalem",
        "altSpellings": [
            "IL",
            "State of Israel",
            "Medīnat Yisrā'el"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic",
            "heb": "Hebrew"
        },
        "translations": {
            "deu": {
                "official": "Staat Israel",
                "common": "Israel"
            },
            "fra": {
                "official": "État d'Israël",
                "common": "Israël"
            },
            "hrv": {
                "official": "Država Izrael",
                "common": "Izrael"
            },
            "ita": {
                "official": "Stato di Israele",
                "common": "Israele"
            },
            "jpn": {
                "official": "イスラエル国",
                "common": "イスラエル"
            },
            "nld": {
                "official": "Staat Israël",
                "common": "Israël"
            },
            "por": {
                "official": "Estado de Israel",
                "common": "Israel"
            },
            "rus": {
                "official": "Государство Израиль",
                "common": "Израиль"
            },
            "spa": {
                "official": "Estado de Israel",
                "common": "Israel"
            },
            "fin": {
                "official": "Israelin valtio",
                "common": "Israel"
            }
        },
        "latlng": [
            31.47,
            35.13
        ],
        "demonym": "Israeli",
        "landlocked": false,
        "borders": [
            "EGY",
            "JOR",
            "LBN",
            "SYR"
        ],
        "area": 20770
    },
    {
        "name": {
            "common": "Italy",
            "official": "Italian Republic",
            "native": {
                "bar": {
                    "official": "Italienische Republik",
                    "common": "Italien"
                },
                "ita": {
                    "official": "Repubblica italiana",
                    "common": "Italia"
                },
                "srd": {
                    "official": "Repubbricanu Italia",
                    "common": "Italia"
                }
            }
        },
        "tld": [
            ".it"
        ],
        "cca2": "IT",
        "ccn3": "380",
        "cca3": "ITA",
        "cioc": "ITA",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "39"
        ],
        "capital": "Rome",
        "altSpellings": [
            "IT",
            "Italian Republic",
            "Repubblica italiana"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "bar": "Austro-Bavarian German",
            "ita": "Italian",
            "srd": "Sardinian"
        },
        "translations": {
            "deu": {
                "official": "Italienische Republik",
                "common": "Italien"
            },
            "fra": {
                "official": "République italienne",
                "common": "Italie"
            },
            "hrv": {
                "official": "talijanska Republika",
                "common": "Italija"
            },
            "ita": {
                "official": "Repubblica italiana",
                "common": "Italia"
            },
            "jpn": {
                "official": "イタリア共和国",
                "common": "イタリア"
            },
            "nld": {
                "official": "Italiaanse Republiek",
                "common": "Italië"
            },
            "por": {
                "official": "República Italiana",
                "common": "Itália"
            },
            "rus": {
                "official": "итальянская Республика",
                "common": "Италия"
            },
            "spa": {
                "official": "República Italiana",
                "common": "Italia"
            },
            "fin": {
                "official": "Italian tasavalta",
                "common": "Italia"
            }
        },
        "latlng": [
            42.83333333,
            12.83333333
        ],
        "demonym": "Italian",
        "landlocked": false,
        "borders": [
            "AUT",
            "FRA",
            "SMR",
            "SVN",
            "CHE",
            "VAT"
        ],
        "area": 301336
    },
    {
        "name": {
            "common": "Jamaica",
            "official": "Jamaica",
            "native": {
                "eng": {
                    "official": "Jamaica",
                    "common": "Jamaica"
                },
                "jam": {
                    "official": "Jamaica",
                    "common": "Jamaica"
                }
            }
        },
        "tld": [
            ".jm"
        ],
        "cca2": "JM",
        "ccn3": "388",
        "cca3": "JAM",
        "cioc": "JAM",
        "currency": [
            "JMD"
        ],
        "callingCode": [
            "1876"
        ],
        "capital": "Kingston",
        "altSpellings": [
            "JM"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English",
            "jam": "Jamaican Patois"
        },
        "translations": {
            "deu": {
                "official": "Jamaika",
                "common": "Jamaika"
            },
            "fra": {
                "official": "Jamaïque",
                "common": "Jamaïque"
            },
            "hrv": {
                "official": "Jamajka",
                "common": "Jamajka"
            },
            "ita": {
                "official": "Giamaica",
                "common": "Giamaica"
            },
            "jpn": {
                "official": "ジャマイカ",
                "common": "ジャマイカ"
            },
            "nld": {
                "official": "Jamaica",
                "common": "Jamaica"
            },
            "por": {
                "official": "Jamaica",
                "common": "Jamaica"
            },
            "rus": {
                "official": "Ямайка",
                "common": "Ямайка"
            },
            "spa": {
                "official": "Jamaica",
                "common": "Jamaica"
            },
            "fin": {
                "official": "Jamaika",
                "common": "Jamaika"
            }
        },
        "latlng": [
            18.25,
            -77.5
        ],
        "demonym": "Jamaican",
        "landlocked": false,
        "borders": [],
        "area": 10991
    },
    {
        "name": {
            "common": "Jersey",
            "official": "Bailiwick of Jersey",
            "native": {
                "eng": {
                    "official": "Bailiwick of Jersey",
                    "common": "Jersey"
                },
                "fra": {
                    "official": "Bailliage de Jersey",
                    "common": "Jersey"
                },
                "nrf": {
                    "official": "Bailliage dé Jèrri",
                    "common": "Jèrri"
                }
            }
        },
        "tld": [
            ".je"
        ],
        "cca2": "JE",
        "ccn3": "832",
        "cca3": "JEY",
        "cioc": "",
        "currency": [
            "GBP"
        ],
        "callingCode": [
            "44"
        ],
        "capital": "Saint Helier",
        "altSpellings": [
            "JE",
            "Bailiwick of Jersey",
            "Bailliage de Jersey",
            "Bailliage dé Jèrri"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "eng": "English",
            "fra": "French",
            "nrf": "Jèrriais"
        },
        "translations": {
            "deu": {
                "official": "Vogtei Jersey",
                "common": "Jersey"
            },
            "fra": {
                "official": "Bailliage de Jersey",
                "common": "Jersey"
            },
            "hrv": {
                "official": "Struka od Jersey",
                "common": "Jersey"
            },
            "ita": {
                "official": "Baliato di Jersey",
                "common": "Isola di Jersey"
            },
            "jpn": {
                "official": "ジャージの得意分野",
                "common": "ジャージー"
            },
            "nld": {
                "official": "Baljuwschap Jersey",
                "common": "Jersey"
            },
            "por": {
                "official": "Bailiado de Jersey",
                "common": "Jersey"
            },
            "rus": {
                "official": "Коронное владение Джерси",
                "common": "Джерси"
            },
            "spa": {
                "official": "Bailía de Jersey",
                "common": "Jersey"
            },
            "fin": {
                "official": "Jersey",
                "common": "Jersey"
            }
        },
        "latlng": [
            49.25,
            -2.16666666
        ],
        "demonym": "Channel Islander",
        "landlocked": false,
        "borders": [],
        "area": 116
    },
    {
        "name": {
            "common": "Jordan",
            "official": "Hashemite Kingdom of Jordan",
            "native": {
                "ara": {
                    "official": "المملكة الأردنية الهاشمية",
                    "common": "الأردن"
                }
            }
        },
        "tld": [
            ".jo",
            "الاردن."
        ],
        "cca2": "JO",
        "ccn3": "400",
        "cca3": "JOR",
        "cioc": "JOR",
        "currency": [
            "JOD"
        ],
        "callingCode": [
            "962"
        ],
        "capital": "Amman",
        "altSpellings": [
            "JO",
            "Hashemite Kingdom of Jordan",
            "al-Mamlakah al-Urdunīyah al-Hāshimīyah"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Haschemitisches Königreich Jordanien",
                "common": "Jordanien"
            },
            "fra": {
                "official": "Royaume hachémite de Jordanie",
                "common": "Jordanie"
            },
            "hrv": {
                "official": "Hašemitske Kraljevine Jordan",
                "common": "Jordan"
            },
            "ita": {
                "official": "Regno hascemita di Giordania",
                "common": "Giordania"
            },
            "jpn": {
                "official": "ヨルダン·ハシミテ王国",
                "common": "ヨルダン"
            },
            "nld": {
                "official": "Hasjemitisch Koninkrijk Jordanië",
                "common": "Jordanië"
            },
            "por": {
                "official": "Reino Hachemita da Jordânia",
                "common": "Jordânia"
            },
            "rus": {
                "official": "Иорданского Хашимитского Королевства",
                "common": "Иордания"
            },
            "spa": {
                "official": "Reino Hachemita de Jordania",
                "common": "Jordania"
            },
            "fin": {
                "official": "Jordanian hašemiittinen kunigaskunta",
                "common": "Jordania"
            }
        },
        "latlng": [
            31,
            36
        ],
        "demonym": "Jordanian",
        "landlocked": false,
        "borders": [
            "IRQ",
            "ISR",
            "SAU",
            "SYR"
        ],
        "area": 89342
    },
    {
        "name": {
            "common": "Japan",
            "official": "Japan",
            "native": {
                "jpn": {
                    "official": "日本",
                    "common": "日本"
                }
            }
        },
        "tld": [
            ".jp",
            ".みんな"
        ],
        "cca2": "JP",
        "ccn3": "392",
        "cca3": "JPN",
        "cioc": "JPN",
        "currency": [
            "JPY"
        ],
        "callingCode": [
            "81"
        ],
        "capital": "Tokyo",
        "altSpellings": [
            "JP",
            "Nippon",
            "Nihon"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "jpn": "Japanese"
        },
        "translations": {
            "deu": {
                "official": "Japan",
                "common": "Japan"
            },
            "fra": {
                "official": "Japon",
                "common": "Japon"
            },
            "hrv": {
                "official": "Japan",
                "common": "Japan"
            },
            "ita": {
                "official": "Giappone",
                "common": "Giappone"
            },
            "jpn": {
                "official": "日本",
                "common": "日本"
            },
            "nld": {
                "official": "Japan",
                "common": "Japan"
            },
            "por": {
                "official": "Japão",
                "common": "Japão"
            },
            "rus": {
                "official": "Япония",
                "common": "Япония"
            },
            "spa": {
                "official": "Japón",
                "common": "Japón"
            },
            "fin": {
                "official": "Japani",
                "common": "Japani"
            }
        },
        "latlng": [
            36,
            138
        ],
        "demonym": "Japanese",
        "landlocked": false,
        "borders": [],
        "area": 377930
    },
    {
        "name": {
            "common": "Kazakhstan",
            "official": "Republic of Kazakhstan",
            "native": {
                "kaz": {
                    "official": "Қазақстан Республикасы",
                    "common": "Қазақстан"
                },
                "rus": {
                    "official": "Республика Казахстан",
                    "common": "Казахстан"
                }
            }
        },
        "tld": [
            ".kz",
            ".қаз"
        ],
        "cca2": "KZ",
        "ccn3": "398",
        "cca3": "KAZ",
        "cioc": "KAZ",
        "currency": [
            "KZT"
        ],
        "callingCode": [
            "76",
            "77"
        ],
        "capital": "Astana",
        "altSpellings": [
            "KZ",
            "Qazaqstan",
            "Казахстан",
            "Republic of Kazakhstan",
            "Қазақстан Республикасы",
            "Qazaqstan Respublïkası",
            "Республика Казахстан",
            "Respublika Kazakhstan"
        ],
        "region": "Asia",
        "subregion": "Central Asia",
        "languages": {
            "kaz": "Kazakh",
            "rus": "Russian"
        },
        "translations": {
            "deu": {
                "official": "Republik Kasachstan",
                "common": "Kasachstan"
            },
            "fra": {
                "official": "République du Kazakhstan",
                "common": "Kazakhstan"
            },
            "hrv": {
                "official": "Republika Kazahstan",
                "common": "Kazahstan"
            },
            "ita": {
                "official": "Repubblica del Kazakhstan",
                "common": "Kazakistan"
            },
            "jpn": {
                "official": "カザフスタン共和国",
                "common": "カザフスタン"
            },
            "nld": {
                "official": "Republiek Kazachstan",
                "common": "Kazachstan"
            },
            "por": {
                "official": "República do Cazaquistão",
                "common": "Cazaquistão"
            },
            "rus": {
                "official": "Республика Казахстан",
                "common": "Казахстан"
            },
            "spa": {
                "official": "República de Kazajstán",
                "common": "Kazajistán"
            },
            "fin": {
                "official": "Kazakstanin tasavalta",
                "common": "Kazakstan"
            }
        },
        "latlng": [
            48,
            68
        ],
        "demonym": "Kazakhstani",
        "landlocked": true,
        "borders": [
            "CHN",
            "KGZ",
            "RUS",
            "TKM",
            "UZB"
        ],
        "area": 2724900
    },
    {
        "name": {
            "common": "Kenya",
            "official": "Republic of Kenya",
            "native": {
                "eng": {
                    "official": "Republic of Kenya",
                    "common": "Kenya"
                },
                "swa": {
                    "official": "Republic of Kenya",
                    "common": "Kenya"
                }
            }
        },
        "tld": [
            ".ke"
        ],
        "cca2": "KE",
        "ccn3": "404",
        "cca3": "KEN",
        "cioc": "KEN",
        "currency": [
            "KES"
        ],
        "callingCode": [
            "254"
        ],
        "capital": "Nairobi",
        "altSpellings": [
            "KE",
            "Republic of Kenya",
            "Jamhuri ya Kenya"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "swa": "Swahili"
        },
        "translations": {
            "deu": {
                "official": "Republik Kenia",
                "common": "Kenia"
            },
            "fra": {
                "official": "République du Kenya",
                "common": "Kenya"
            },
            "hrv": {
                "official": "Republika Kenija",
                "common": "Kenija"
            },
            "ita": {
                "official": "Repubblica del Kenya",
                "common": "Kenya"
            },
            "jpn": {
                "official": "ケニア共和国",
                "common": "ケニア"
            },
            "nld": {
                "official": "Republiek Kenia",
                "common": "Kenia"
            },
            "por": {
                "official": "República do Quénia",
                "common": "Quénia"
            },
            "rus": {
                "official": "Республика Кения",
                "common": "Кения"
            },
            "spa": {
                "official": "República de Kenya",
                "common": "Kenia"
            },
            "fin": {
                "official": "Kenian tasavalta",
                "common": "Kenia"
            }
        },
        "latlng": [
            1,
            38
        ],
        "demonym": "Kenyan",
        "landlocked": false,
        "borders": [
            "ETH",
            "SOM",
            "SSD",
            "TZA",
            "UGA"
        ],
        "area": 580367
    },
    {
        "name": {
            "common": "Kyrgyzstan",
            "official": "Kyrgyz Republic",
            "native": {
                "kir": {
                    "official": "Кыргыз Республикасы",
                    "common": "Кыргызстан"
                },
                "rus": {
                    "official": "Кыргызская Республика",
                    "common": "Киргизия"
                }
            }
        },
        "tld": [
            ".kg"
        ],
        "cca2": "KG",
        "ccn3": "417",
        "cca3": "KGZ",
        "cioc": "KGZ",
        "currency": [
            "KGS"
        ],
        "callingCode": [
            "996"
        ],
        "capital": "Bishkek",
        "altSpellings": [
            "KG",
            "Киргизия",
            "Kyrgyz Republic",
            "Кыргыз Республикасы",
            "Kyrgyz Respublikasy"
        ],
        "region": "Asia",
        "subregion": "Central Asia",
        "languages": {
            "kir": "Kyrgyz",
            "rus": "Russian"
        },
        "translations": {
            "deu": {
                "official": "Kirgisische Republik",
                "common": "Kirgisistan"
            },
            "fra": {
                "official": "République kirghize",
                "common": "Kirghizistan"
            },
            "hrv": {
                "official": "Kirgistanu",
                "common": "Kirgistan"
            },
            "ita": {
                "official": "Kirghizistan",
                "common": "Kirghizistan"
            },
            "jpn": {
                "official": "キルギス共和国",
                "common": "キルギス"
            },
            "nld": {
                "official": "Kirgizische Republiek",
                "common": "Kirgizië"
            },
            "por": {
                "official": "República do Quirguistão",
                "common": "Quirguistão"
            },
            "rus": {
                "official": "Кыргызская Республика",
                "common": "Киргизия"
            },
            "spa": {
                "official": "República Kirguisa",
                "common": "Kirguizistán"
            },
            "fin": {
                "official": "Kirgisian tasavalta",
                "common": "Kirgisia"
            }
        },
        "latlng": [
            41,
            75
        ],
        "demonym": "Kirghiz",
        "landlocked": true,
        "borders": [
            "CHN",
            "KAZ",
            "TJK",
            "UZB"
        ],
        "area": 199951
    },
    {
        "name": {
            "common": "Cambodia",
            "official": "Kingdom of Cambodia",
            "native": {
                "khm": {
                    "official": "ព្រះរាជាណាចក្រកម្ពុជា",
                    "common": "Kâmpŭchéa"
                }
            }
        },
        "tld": [
            ".kh"
        ],
        "cca2": "KH",
        "ccn3": "116",
        "cca3": "KHM",
        "cioc": "CAM",
        "currency": [
            "KHR"
        ],
        "callingCode": [
            "855"
        ],
        "capital": "Phnom Penh",
        "altSpellings": [
            "KH",
            "Kingdom of Cambodia"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "khm": "Khmer"
        },
        "translations": {
            "cym": {
                "official": "Kingdom of Cambodia",
                "common": "Cambodia"
            },
            "deu": {
                "official": "Königreich Kambodscha",
                "common": "Kambodscha"
            },
            "fra": {
                "official": "Royaume du Cambodge",
                "common": "Cambodge"
            },
            "hrv": {
                "official": "Kraljevina Kambodža",
                "common": "Kambodža"
            },
            "ita": {
                "official": "Regno di Cambogia",
                "common": "Cambogia"
            },
            "jpn": {
                "official": "カンボジア王国",
                "common": "カンボジア"
            },
            "nld": {
                "official": "Koninkrijk Cambodja",
                "common": "Cambodja"
            },
            "por": {
                "official": "Reino do Camboja",
                "common": "Camboja"
            },
            "rus": {
                "official": "Королевство Камбоджа",
                "common": "Камбоджа"
            },
            "spa": {
                "official": "Reino de Camboya",
                "common": "Camboya"
            },
            "fin": {
                "official": "Kambodžan kuningaskunta",
                "common": "Kambodža"
            }
        },
        "latlng": [
            13,
            105
        ],
        "demonym": "Cambodian",
        "landlocked": false,
        "borders": [
            "LAO",
            "THA",
            "VNM"
        ],
        "area": 181035
    },
    {
        "name": {
            "common": "Kiribati",
            "official": "Independent and Sovereign Republic of Kiribati",
            "native": {
                "eng": {
                    "official": "Independent and Sovereign Republic of Kiribati",
                    "common": "Kiribati"
                },
                "gil": {
                    "official": "Ribaberiki Kiribati",
                    "common": "Kiribati"
                }
            }
        },
        "tld": [
            ".ki"
        ],
        "cca2": "KI",
        "ccn3": "296",
        "cca3": "KIR",
        "cioc": "KIR",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "686"
        ],
        "capital": "South Tarawa",
        "altSpellings": [
            "KI",
            "Republic of Kiribati",
            "Ribaberiki Kiribati"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "eng": "English",
            "gil": "Gilbertese"
        },
        "translations": {
            "deu": {
                "official": "Unabhängige und souveräne Republik Kiribati",
                "common": "Kiribati"
            },
            "fra": {
                "official": "République de Kiribati",
                "common": "Kiribati"
            },
            "hrv": {
                "official": "Samostalne i suverene Republike Kiribati",
                "common": "Kiribati"
            },
            "ita": {
                "official": "Repubblica indipendente e sovrano di Kiribati",
                "common": "Kiribati"
            },
            "jpn": {
                "official": "キリバスの独立と主権共和国",
                "common": "キリバス"
            },
            "nld": {
                "official": "Onafhankelijke en soevereine republiek Kiribati",
                "common": "Kiribati"
            },
            "por": {
                "official": "Independente e soberano República de Kiribati",
                "common": "Kiribati"
            },
            "rus": {
                "official": "Независимой и суверенной Республики Кирибати",
                "common": "Кирибати"
            },
            "spa": {
                "official": "República Independiente y Soberano de Kiribati",
                "common": "Kiribati"
            },
            "fin": {
                "official": "Kiribatin tasavalta",
                "common": "Kiribati"
            }
        },
        "latlng": [
            1.41666666,
            173
        ],
        "demonym": "I-Kiribati",
        "landlocked": false,
        "borders": [],
        "area": 811
    },
    {
        "name": {
            "common": "Saint Kitts and Nevis",
            "official": "Federation of Saint Christopher and Nevisa",
            "native": {
                "eng": {
                    "official": "Federation of Saint Christopher and Nevisa",
                    "common": "Saint Kitts and Nevis"
                }
            }
        },
        "tld": [
            ".kn"
        ],
        "cca2": "KN",
        "ccn3": "659",
        "cca3": "KNA",
        "cioc": "SKN",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1869"
        ],
        "capital": "Basseterre",
        "altSpellings": [
            "KN",
            "Federation of Saint Christopher and Nevis"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Föderation von Saint Kitts und Nevisa",
                "common": "Saint Christopher und Nevis"
            },
            "fra": {
                "official": "Fédération de Saint -Christophe-et Nevisa",
                "common": "Saint-Christophe-et-Niévès"
            },
            "hrv": {
                "official": "Federacija Sv.Kristofora i Nevisa",
                "common": "Sveti Kristof i Nevis"
            },
            "ita": {
                "official": "Federazione di Saint Christopher e Nevisa",
                "common": "Saint Kitts e Nevis"
            },
            "jpn": {
                "official": "セントクリストファーNevisa連盟",
                "common": "セントクリストファー・ネイビス"
            },
            "nld": {
                "official": "Federatie van Saint Kitts en Nevisa",
                "common": "Saint Kitts en Nevis"
            },
            "por": {
                "official": "Federação de São Cristóvão e Nevisa",
                "common": "São Cristóvão e Nevis"
            },
            "rus": {
                "official": "Федерация Сент-Кристофер и Nevisa",
                "common": "Сент-Китс и Невис"
            },
            "spa": {
                "official": "Federación de San Cristóbal y Nevisa",
                "common": "San Cristóbal y Nieves"
            },
            "fin": {
                "official": "Saint Christopherin ja Nevisin federaatio",
                "common": "Saint Kitts ja Nevis"
            }
        },
        "latlng": [
            17.33333333,
            -62.75
        ],
        "demonym": "Kittitian or Nevisian",
        "landlocked": false,
        "borders": [],
        "area": 261
    },
    {
        "name": {
            "common": "South Korea",
            "official": "Republic of Korea",
            "native": {
                "kor": {
                    "official": "한국",
                    "common": "대한민국"
                }
            }
        },
        "tld": [
            ".kr",
            ".한국"
        ],
        "cca2": "KR",
        "ccn3": "410",
        "cca3": "KOR",
        "cioc": "KOR",
        "currency": [
            "KRW"
        ],
        "callingCode": [
            "82"
        ],
        "capital": "Seoul",
        "altSpellings": [
            "KR",
            "Korea, Republic of",
            "Republic of Korea"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "kor": "Korean"
        },
        "translations": {
            "deu": {
                "official": "Republik Korea",
                "common": "Südkorea"
            },
            "fra": {
                "official": "République de Corée",
                "common": "Corée du Sud"
            },
            "hrv": {
                "official": "Republika Koreja",
                "common": "Južna Koreja"
            },
            "ita": {
                "official": "Repubblica di Corea",
                "common": "Corea del Sud"
            },
            "jpn": {
                "official": "大韓民国",
                "common": "大韓民国"
            },
            "nld": {
                "official": "Republiek Korea",
                "common": "Zuid-Korea"
            },
            "por": {
                "official": "República da Coreia",
                "common": "Coreia do Sul"
            },
            "rus": {
                "official": "Республика Корея",
                "common": "Южная Корея"
            },
            "spa": {
                "official": "República de Corea",
                "common": "Corea del Sur"
            },
            "fin": {
                "official": "Korean tasavalta",
                "common": "Etelä-Korea"
            }
        },
        "latlng": [
            37,
            127.5
        ],
        "demonym": "South Korean",
        "landlocked": false,
        "borders": [
            "PRK"
        ],
        "area": 100210
    },
    {
        "name": {
            "common": "Kosovo",
            "official": "Republic of Kosovo",
            "native": {
                "sqi": {
                    "official": "Republika e Kosovës",
                    "common": "Kosova"
                },
                "srp": {
                    "official": "Република Косово",
                    "common": "Косово"
                }
            }
        },
        "tld": [],
        "cca2": "XK",
        "ccn3": "",
        "cca3": "KOS",
        "cioc": "KOS",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "383"
        ],
        "capital": "Pristina",
        "altSpellings": [
            "XK",
            "Република Косово"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "sqi": "Albanian",
            "srp": "Serbian"
        },
        "translations": {
            "hrv": {
                "official": "Republika Kosovo",
                "common": "Kosovo"
            },
            "por": {
                "official": "República do Kosovo",
                "common": "Kosovo"
            },
            "rus": {
                "official": "Республика Косово",
                "common": "Республика Косово"
            },
            "spa": {
                "official": "República de Kosovo",
                "common": "Kosovo"
            },
            "fin": {
                "official": "Kosovon tasavalta",
                "common": "Kosovo"
            }
        },
        "latlng": [
            42.666667,
            21.166667
        ],
        "demonym": "Kosovar",
        "landlocked": true,
        "borders": [
            "ALB",
            "MKD",
            "MNE",
            "SRB"
        ],
        "area": 10908
    },
    {
        "name": {
            "common": "Kuwait",
            "official": "State of Kuwait",
            "native": {
                "ara": {
                    "official": "دولة الكويت",
                    "common": "الكويت"
                }
            }
        },
        "tld": [
            ".kw"
        ],
        "cca2": "KW",
        "ccn3": "414",
        "cca3": "KWT",
        "cioc": "KUW",
        "currency": [
            "KWD"
        ],
        "callingCode": [
            "965"
        ],
        "capital": "Kuwait City",
        "altSpellings": [
            "KW",
            "State of Kuwait",
            "Dawlat al-Kuwait"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Staat Kuwait",
                "common": "Kuwait"
            },
            "fra": {
                "official": "État du Koweït",
                "common": "Koweït"
            },
            "hrv": {
                "official": "Država Kuvajt",
                "common": "Kuvajt"
            },
            "ita": {
                "official": "Stato del Kuwait",
                "common": "Kuwait"
            },
            "jpn": {
                "official": "クウェート国",
                "common": "クウェート"
            },
            "nld": {
                "official": "Staat Koeweit",
                "common": "Koeweit"
            },
            "por": {
                "official": "Estado do Kuwait",
                "common": "Kuwait"
            },
            "rus": {
                "official": "Государство Кувейт",
                "common": "Кувейт"
            },
            "spa": {
                "official": "Estado de Kuwait",
                "common": "Kuwait"
            },
            "fin": {
                "official": "Kuwaitin valtio",
                "common": "Kuwait"
            }
        },
        "latlng": [
            29.5,
            45.75
        ],
        "demonym": "Kuwaiti",
        "landlocked": false,
        "borders": [
            "IRQ",
            "SAU"
        ],
        "area": 17818
    },
    {
        "name": {
            "common": "Laos",
            "official": "Lao People's Democratic Republic",
            "native": {
                "lao": {
                    "official": "ສາທາລະນະ ຊາທິປະໄຕ ຄົນລາວ ຂອງ",
                    "common": "ສປປລາວ"
                }
            }
        },
        "tld": [
            ".la"
        ],
        "cca2": "LA",
        "ccn3": "418",
        "cca3": "LAO",
        "cioc": "LAO",
        "currency": [
            "LAK"
        ],
        "callingCode": [
            "856"
        ],
        "capital": "Vientiane",
        "altSpellings": [
            "LA",
            "Lao",
            "Lao People's Democratic Republic",
            "Sathalanalat Paxathipatai Paxaxon Lao"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "lao": "Lao"
        },
        "translations": {
            "deu": {
                "official": "Laos, Demokratische Volksrepublik",
                "common": "Laos"
            },
            "fra": {
                "official": "République démocratique populaire lao",
                "common": "Laos"
            },
            "hrv": {
                "official": "Narodna Demokratska Republika",
                "common": "Laos"
            },
            "ita": {
                "official": "Repubblica democratica popolare del Laos",
                "common": "Laos"
            },
            "jpn": {
                "official": "ラオス人民民主共和国",
                "common": "ラオス人民民主共和国"
            },
            "nld": {
                "official": "Lao Democratische Volksrepubliek",
                "common": "Laos"
            },
            "por": {
                "official": "Laos, República Democrática",
                "common": "Laos"
            },
            "rus": {
                "official": "Лаосская Народно-Демократическая Республика",
                "common": "Лаос"
            },
            "spa": {
                "official": "República Democrática Popular Lao",
                "common": "Laos"
            },
            "fin": {
                "official": "Laosin demokraattinen kansantasavalta",
                "common": "Laos"
            }
        },
        "latlng": [
            18,
            105
        ],
        "demonym": "Laotian",
        "landlocked": true,
        "borders": [
            "MMR",
            "KHM",
            "CHN",
            "THA",
            "VNM"
        ],
        "area": 236800
    },
    {
        "name": {
            "common": "Lebanon",
            "official": "Lebanese Republic",
            "native": {
                "ara": {
                    "official": "الجمهورية اللبنانية",
                    "common": "لبنان"
                },
                "fra": {
                    "official": "République libanaise",
                    "common": "Liban"
                }
            }
        },
        "tld": [
            ".lb"
        ],
        "cca2": "LB",
        "ccn3": "422",
        "cca3": "LBN",
        "cioc": "LIB",
        "currency": [
            "LBP"
        ],
        "callingCode": [
            "961"
        ],
        "capital": "Beirut",
        "altSpellings": [
            "LB",
            "Lebanese Republic",
            "Al-Jumhūrīyah Al-Libnānīyah"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic",
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Libanesische Republik",
                "common": "Libanon"
            },
            "fra": {
                "official": "République libanaise",
                "common": "Liban"
            },
            "hrv": {
                "official": "Libanonska Republika",
                "common": "Libanon"
            },
            "ita": {
                "official": "Repubblica libanese",
                "common": "Libano"
            },
            "jpn": {
                "official": "レバノン共和国",
                "common": "レバノン"
            },
            "nld": {
                "official": "Libanese Republiek",
                "common": "Libanon"
            },
            "por": {
                "official": "República Libanesa",
                "common": "Líbano"
            },
            "rus": {
                "official": "Ливанская Республика",
                "common": "Ливан"
            },
            "spa": {
                "official": "República Libanesa",
                "common": "Líbano"
            },
            "fin": {
                "official": "Libanonin tasavalta",
                "common": "Libanon"
            }
        },
        "latlng": [
            33.83333333,
            35.83333333
        ],
        "demonym": "Lebanese",
        "landlocked": false,
        "borders": [
            "ISR",
            "SYR"
        ],
        "area": 10452
    },
    {
        "name": {
            "common": "Liberia",
            "official": "Republic of Liberia",
            "native": {
                "eng": {
                    "official": "Republic of Liberia",
                    "common": "Liberia"
                }
            }
        },
        "tld": [
            ".lr"
        ],
        "cca2": "LR",
        "ccn3": "430",
        "cca3": "LBR",
        "cioc": "LBR",
        "currency": [
            "LRD"
        ],
        "callingCode": [
            "231"
        ],
        "capital": "Monrovia",
        "altSpellings": [
            "LR",
            "Republic of Liberia"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Liberia",
                "common": "Liberia"
            },
            "fra": {
                "official": "République du Libéria",
                "common": "Liberia"
            },
            "hrv": {
                "official": "Republika Liberija",
                "common": "Liberija"
            },
            "ita": {
                "official": "Repubblica di Liberia",
                "common": "Liberia"
            },
            "jpn": {
                "official": "リベリア共和国",
                "common": "リベリア"
            },
            "nld": {
                "official": "Republiek Liberia",
                "common": "Liberia"
            },
            "por": {
                "official": "República da Libéria",
                "common": "Libéria"
            },
            "rus": {
                "official": "Республика Либерия",
                "common": "Либерия"
            },
            "spa": {
                "official": "República de Liberia",
                "common": "Liberia"
            },
            "fin": {
                "official": "Liberian tasavalta",
                "common": "Liberia"
            }
        },
        "latlng": [
            6.5,
            -9.5
        ],
        "demonym": "Liberian",
        "landlocked": false,
        "borders": [
            "GIN",
            "CIV",
            "SLE"
        ],
        "area": 111369
    },
    {
        "name": {
            "common": "Libya",
            "official": "State of Libya",
            "native": {
                "ara": {
                    "official": "الدولة ليبيا",
                    "common": "‏ليبيا"
                }
            }
        },
        "tld": [
            ".ly"
        ],
        "cca2": "LY",
        "ccn3": "434",
        "cca3": "LBY",
        "cioc": "LBA",
        "currency": [
            "LYD"
        ],
        "callingCode": [
            "218"
        ],
        "capital": "Tripoli",
        "altSpellings": [
            "LY",
            "State of Libya",
            "Dawlat Libya"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Staat Libyen",
                "common": "Libyen"
            },
            "fra": {
                "official": "Grande République arabe libyenne populaire et socialiste",
                "common": "Libye"
            },
            "hrv": {
                "official": "Država Libiji",
                "common": "Libija"
            },
            "ita": {
                "official": "Stato della Libia",
                "common": "Libia"
            },
            "jpn": {
                "official": "リビアの国家",
                "common": "リビア"
            },
            "nld": {
                "official": "Staat van Libië",
                "common": "Libië"
            },
            "por": {
                "official": "Estado da Líbia",
                "common": "Líbia"
            },
            "rus": {
                "official": "Государство Ливии",
                "common": "Ливия"
            },
            "spa": {
                "official": "Estado de Libia",
                "common": "Libia"
            },
            "fin": {
                "official": "Libyan valtio",
                "common": "Libya"
            }
        },
        "latlng": [
            25,
            17
        ],
        "demonym": "Libyan",
        "landlocked": false,
        "borders": [
            "DZA",
            "TCD",
            "EGY",
            "NER",
            "SDN",
            "TUN"
        ],
        "area": 1759540
    },
    {
        "name": {
            "common": "Saint Lucia",
            "official": "Saint Lucia",
            "native": {
                "eng": {
                    "official": "Saint Lucia",
                    "common": "Saint Lucia"
                }
            }
        },
        "tld": [
            ".lc"
        ],
        "cca2": "LC",
        "ccn3": "662",
        "cca3": "LCA",
        "cioc": "LCA",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1758"
        ],
        "capital": "Castries",
        "altSpellings": [
            "LC"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "St. Lucia",
                "common": "Saint Lucia"
            },
            "fra": {
                "official": "Sainte-Lucie",
                "common": "Sainte-Lucie"
            },
            "hrv": {
                "official": "Sveta Lucija",
                "common": "Sveta Lucija"
            },
            "ita": {
                "official": "Santa Lucia",
                "common": "Santa Lucia"
            },
            "jpn": {
                "official": "セントルシア",
                "common": "セントルシア"
            },
            "nld": {
                "official": "Saint Lucia",
                "common": "Saint Lucia"
            },
            "por": {
                "official": "Santa Lúcia",
                "common": "Santa Lúcia"
            },
            "rus": {
                "official": "Сент-Люсия",
                "common": "Сент-Люсия"
            },
            "spa": {
                "official": "Santa Lucía",
                "common": "Santa Lucía"
            },
            "fin": {
                "official": "Saint Lucia",
                "common": "Saint Lucia"
            }
        },
        "latlng": [
            13.88333333,
            -60.96666666
        ],
        "demonym": "Saint Lucian",
        "landlocked": false,
        "borders": [],
        "area": 616
    },
    {
        "name": {
            "common": "Liechtenstein",
            "official": "Principality of Liechtenstein",
            "native": {
                "deu": {
                    "official": "Fürstentum Liechtenstein",
                    "common": "Liechtenstein"
                }
            }
        },
        "tld": [
            ".li"
        ],
        "cca2": "LI",
        "ccn3": "438",
        "cca3": "LIE",
        "cioc": "LIE",
        "currency": [
            "CHF"
        ],
        "callingCode": [
            "423"
        ],
        "capital": "Vaduz",
        "altSpellings": [
            "LI",
            "Principality of Liechtenstein",
            "Fürstentum Liechtenstein"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "deu": "German"
        },
        "translations": {
            "deu": {
                "official": "Fürstentum Liechtenstein",
                "common": "Liechtenstein"
            },
            "fra": {
                "official": "Principauté du Liechtenstein",
                "common": "Liechtenstein"
            },
            "hrv": {
                "official": "Kneževina Lihtenštajn",
                "common": "Lihtenštajn"
            },
            "ita": {
                "official": "Principato del Liechtenstein",
                "common": "Liechtenstein"
            },
            "jpn": {
                "official": "リヒテンシュタイン公国",
                "common": "リヒテンシュタイン"
            },
            "nld": {
                "official": "Vorstendom Liechtenstein",
                "common": "Liechtenstein"
            },
            "por": {
                "official": "Principado de Liechtenstein",
                "common": "Liechtenstein"
            },
            "rus": {
                "official": "Княжество Лихтенштейн",
                "common": "Лихтенштейн"
            },
            "spa": {
                "official": "Principado de Liechtenstein",
                "common": "Liechtenstein"
            },
            "fin": {
                "official": "Liechensteinin ruhtinaskunta",
                "common": "Liechenstein"
            }
        },
        "latlng": [
            47.26666666,
            9.53333333
        ],
        "demonym": "Liechtensteiner",
        "landlocked": true,
        "borders": [
            "AUT",
            "CHE"
        ],
        "area": 160
    },
    {
        "name": {
            "common": "Sri Lanka",
            "official": "Democratic Socialist Republic of Sri Lanka",
            "native": {
                "sin": {
                    "official": "ශ්‍රී ලංකා ප්‍රජාතාන්ත්‍රික සමාජවාදී ජනරජය",
                    "common": "ශ්‍රී ලංකාව"
                },
                "tam": {
                    "official": "இலங்கை சனநாயக சோசலிசக் குடியரசு",
                    "common": "இலங்கை"
                }
            }
        },
        "tld": [
            ".lk",
            ".இலங்கை",
            ".ලංකා"
        ],
        "cca2": "LK",
        "ccn3": "144",
        "cca3": "LKA",
        "cioc": "SRI",
        "currency": [
            "LKR"
        ],
        "callingCode": [
            "94"
        ],
        "capital": "Colombo",
        "altSpellings": [
            "LK",
            "ilaṅkai",
            "Democratic Socialist Republic of Sri Lanka"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "sin": "Sinhala",
            "tam": "Tamil"
        },
        "translations": {
            "deu": {
                "official": "Demokratische Sozialistische Republik Sri Lanka",
                "common": "Sri Lanka"
            },
            "fra": {
                "official": "République démocratique socialiste du Sri Lanka",
                "common": "Sri Lanka"
            },
            "hrv": {
                "official": "Demokratska Socijalističke Republike Šri Lanke",
                "common": "Šri Lanka"
            },
            "ita": {
                "official": "Repubblica democratica socialista dello Sri Lanka",
                "common": "Sri Lanka"
            },
            "jpn": {
                "official": "スリランカ民主社会主義共和国",
                "common": "スリランカ"
            },
            "nld": {
                "official": "Democratische Socialistische Republiek Sri Lanka",
                "common": "Sri Lanka"
            },
            "por": {
                "official": "República Democrática Socialista do Sri Lanka",
                "common": "Sri Lanka"
            },
            "rus": {
                "official": "Демократическая Социалистическая Республика Шри-Ланка",
                "common": "Шри-Ланка"
            },
            "spa": {
                "official": "República Democrática Socialista de Sri Lanka",
                "common": "Sri Lanka"
            },
            "fin": {
                "official": "Sri Lankan demokraattinen sosialistinen tasavalta",
                "common": "Sri Lanka"
            }
        },
        "latlng": [
            7,
            81
        ],
        "demonym": "Sri Lankan",
        "landlocked": false,
        "borders": [
            "IND"
        ],
        "area": 65610
    },
    {
        "name": {
            "common": "Lesotho",
            "official": "Kingdom of Lesotho",
            "native": {
                "eng": {
                    "official": "Kingdom of Lesotho",
                    "common": "Lesotho"
                },
                "sot": {
                    "official": "Kingdom of Lesotho",
                    "common": "Lesotho"
                }
            }
        },
        "tld": [
            ".ls"
        ],
        "cca2": "LS",
        "ccn3": "426",
        "cca3": "LSO",
        "cioc": "LES",
        "currency": [
            "LSL",
            "ZAR"
        ],
        "callingCode": [
            "266"
        ],
        "capital": "Maseru",
        "altSpellings": [
            "LS",
            "Kingdom of Lesotho",
            "Muso oa Lesotho"
        ],
        "region": "Africa",
        "subregion": "Southern Africa",
        "languages": {
            "eng": "English",
            "sot": "Sotho"
        },
        "translations": {
            "deu": {
                "official": "Königreich Lesotho",
                "common": "Lesotho"
            },
            "fra": {
                "official": "Royaume du Lesotho",
                "common": "Lesotho"
            },
            "hrv": {
                "official": "Kraljevina Lesoto",
                "common": "Lesoto"
            },
            "ita": {
                "official": "Regno del Lesotho",
                "common": "Lesotho"
            },
            "jpn": {
                "official": "レソト王国",
                "common": "レソト"
            },
            "nld": {
                "official": "Koninkrijk Lesotho",
                "common": "Lesotho"
            },
            "por": {
                "official": "Reino do Lesoto",
                "common": "Lesoto"
            },
            "rus": {
                "official": "Королевство Лесото",
                "common": "Лесото"
            },
            "spa": {
                "official": "Reino de Lesotho",
                "common": "Lesotho"
            },
            "fin": {
                "official": "Lesothon kuningaskunta",
                "common": "Lesotho"
            }
        },
        "latlng": [
            -29.5,
            28.5
        ],
        "demonym": "Mosotho",
        "landlocked": true,
        "borders": [
            "ZAF"
        ],
        "area": 30355
    },
    {
        "name": {
            "common": "Lithuania",
            "official": "Republic of Lithuania",
            "native": {
                "lit": {
                    "official": "Lietuvos Respublikos",
                    "common": "Lietuva"
                }
            }
        },
        "tld": [
            ".lt"
        ],
        "cca2": "LT",
        "ccn3": "440",
        "cca3": "LTU",
        "cioc": "LTU",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "370"
        ],
        "capital": "Vilnius",
        "altSpellings": [
            "LT",
            "Republic of Lithuania",
            "Lietuvos Respublika"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "lit": "Lithuanian"
        },
        "translations": {
            "deu": {
                "official": "Republik Litauen",
                "common": "Litauen"
            },
            "fra": {
                "official": "République de Lituanie",
                "common": "Lituanie"
            },
            "hrv": {
                "official": "Republika Litva",
                "common": "Litva"
            },
            "ita": {
                "official": "Repubblica di Lituania",
                "common": "Lituania"
            },
            "jpn": {
                "official": "リトアニア共和国",
                "common": "リトアニア"
            },
            "nld": {
                "official": "Republiek Litouwen",
                "common": "Litouwen"
            },
            "por": {
                "official": "República da Lituânia",
                "common": "Lituânia"
            },
            "rus": {
                "official": "Литовская Республика",
                "common": "Литва"
            },
            "spa": {
                "official": "República de Lituania",
                "common": "Lituania"
            },
            "fin": {
                "official": "Liettuan tasavalta",
                "common": "Liettua"
            }
        },
        "latlng": [
            56,
            24
        ],
        "demonym": "Lithuanian",
        "landlocked": false,
        "borders": [
            "BLR",
            "LVA",
            "POL",
            "RUS"
        ],
        "area": 65300
    },
    {
        "name": {
            "common": "Luxembourg",
            "official": "Grand Duchy of Luxembourg",
            "native": {
                "deu": {
                    "official": "Großherzogtum Luxemburg",
                    "common": "Luxemburg"
                },
                "fra": {
                    "official": "Grand-Duché de Luxembourg",
                    "common": "Luxembourg"
                },
                "ltz": {
                    "official": "Groussherzogtum Lëtzebuerg",
                    "common": "Lëtzebuerg"
                }
            }
        },
        "tld": [
            ".lu"
        ],
        "cca2": "LU",
        "ccn3": "442",
        "cca3": "LUX",
        "cioc": "LUX",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "352"
        ],
        "capital": "Luxembourg",
        "altSpellings": [
            "LU",
            "Grand Duchy of Luxembourg",
            "Grand-Duché de Luxembourg",
            "Großherzogtum Luxemburg",
            "Groussherzogtum Lëtzebuerg"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "deu": "German",
            "fra": "French",
            "ltz": "Luxembourgish"
        },
        "translations": {
            "deu": {
                "official": "Großherzogtum Luxemburg,",
                "common": "Luxemburg"
            },
            "fra": {
                "official": "Grand-Duché de Luxembourg",
                "common": "Luxembourg"
            },
            "hrv": {
                "official": "Veliko Vojvodstvo Luksemburg",
                "common": "Luksemburg"
            },
            "ita": {
                "official": "Granducato di Lussemburgo",
                "common": "Lussemburgo"
            },
            "jpn": {
                "official": "ルクセンブルク大公国",
                "common": "ルクセンブルク"
            },
            "nld": {
                "official": "Groothertogdom Luxemburg",
                "common": "Luxemburg"
            },
            "por": {
                "official": "Grão-Ducado do Luxemburgo",
                "common": "Luxemburgo"
            },
            "rus": {
                "official": "Великое Герцогство Люксембург",
                "common": "Люксембург"
            },
            "spa": {
                "official": "Gran Ducado de Luxemburgo",
                "common": "Luxemburgo"
            },
            "fin": {
                "official": "Luxemburgin suurherttuakunta",
                "common": "Luxemburg"
            }
        },
        "latlng": [
            49.75,
            6.16666666
        ],
        "demonym": "Luxembourger",
        "landlocked": true,
        "borders": [
            "BEL",
            "FRA",
            "DEU"
        ],
        "area": 2586
    },
    {
        "name": {
            "common": "Latvia",
            "official": "Republic of Latvia",
            "native": {
                "lav": {
                    "official": "Latvijas Republikas",
                    "common": "Latvija"
                }
            }
        },
        "tld": [
            ".lv"
        ],
        "cca2": "LV",
        "ccn3": "428",
        "cca3": "LVA",
        "cioc": "LAT",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "371"
        ],
        "capital": "Riga",
        "altSpellings": [
            "LV",
            "Republic of Latvia",
            "Latvijas Republika"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "lav": "Latvian"
        },
        "translations": {
            "deu": {
                "official": "Republik Lettland",
                "common": "Lettland"
            },
            "fra": {
                "official": "République de Lettonie",
                "common": "Lettonie"
            },
            "hrv": {
                "official": "Republika Latvija",
                "common": "Latvija"
            },
            "ita": {
                "official": "Repubblica di Lettonia",
                "common": "Lettonia"
            },
            "jpn": {
                "official": "ラトビア共和国",
                "common": "ラトビア"
            },
            "nld": {
                "official": "Republiek Letland",
                "common": "Letland"
            },
            "por": {
                "official": "República da Letónia",
                "common": "Letónia"
            },
            "rus": {
                "official": "Латвийская Республика",
                "common": "Латвия"
            },
            "spa": {
                "official": "República de Letonia",
                "common": "Letonia"
            },
            "fin": {
                "official": "Latvian tasavalta",
                "common": "Latvia"
            }
        },
        "latlng": [
            57,
            25
        ],
        "demonym": "Latvian",
        "landlocked": false,
        "borders": [
            "BLR",
            "EST",
            "LTU",
            "RUS"
        ],
        "area": 64559
    },
    {
        "name": {
            "common": "Macau",
            "official": "Macao Special Administrative Region of the People's Republic of China",
            "native": {
                "por": {
                    "official": "Região Administrativa Especial de Macau da República Popular da China",
                    "common": "Macau"
                },
                "zho": {
                    "official": "澳门特别行政区中国人民共和国",
                    "common": "澳門"
                }
            }
        },
        "tld": [
            ".mo"
        ],
        "cca2": "MO",
        "ccn3": "446",
        "cca3": "MAC",
        "cioc": "",
        "currency": [
            "MOP"
        ],
        "callingCode": [
            "853"
        ],
        "capital": "",
        "altSpellings": [
            "MO",
            "澳门",
            "Macao",
            "Macao Special Administrative Region of the People's Republic of China",
            "中華人民共和國澳門特別行政區",
            "Região Administrativa Especial de Macau da República Popular da China"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "por": "Portuguese",
            "zho": "Chinese"
        },
        "translations": {
            "deu": {
                "official": "Sonderverwaltungsregion Macau der Volksrepublik China",
                "common": "Macao"
            },
            "fra": {
                "official": "Région administrative spéciale de Macao de la République populaire de Chine",
                "common": "Macao"
            },
            "hrv": {
                "official": "Makao Posebnog upravnog područjaNarodne Republike Kine",
                "common": "Makao"
            },
            "ita": {
                "official": "Macao Regione amministrativa speciale della Repubblica Popolare Cinese",
                "common": "Macao"
            },
            "jpn": {
                "official": "中華人民共和国マカオ特別行政区",
                "common": "マカオ"
            },
            "nld": {
                "official": "Speciale Administratieve Regio Macau van de Volksrepubliek China",
                "common": "Macao"
            },
            "por": {
                "official": "Macau Região Administrativa Especial da República Popular da China",
                "common": "Macau"
            },
            "rus": {
                "official": "Специальный административный район Макао Китайской Народной Республики Китай",
                "common": "Макао"
            },
            "spa": {
                "official": "Macao, Región Administrativa Especial de la República Popular China",
                "common": "Macao"
            },
            "fin": {
                "official": "Macaon Kiinan kansantasavallan erityishallintoalue",
                "common": "Macao"
            }
        },
        "latlng": [
            22.16666666,
            113.55
        ],
        "demonym": "Chinese",
        "landlocked": false,
        "borders": [
            "CHN"
        ],
        "area": 30
    },
    {
        "name": {
            "common": "Saint Martin",
            "official": "Saint Martin",
            "native": {
                "fra": {
                    "official": "Saint-Martin",
                    "common": "Saint-Martin"
                }
            }
        },
        "tld": [
            ".fr",
            ".gp"
        ],
        "cca2": "MF",
        "ccn3": "663",
        "cca3": "MAF",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "590"
        ],
        "capital": "Marigot",
        "altSpellings": [
            "MF",
            "Collectivity of Saint Martin",
            "Collectivité de Saint-Martin",
            "Saint Martin (French part)"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "St. Martin",
                "common": "Saint Martin"
            },
            "fra": {
                "official": "Saint-Martin",
                "common": "Saint-Martin"
            },
            "hrv": {
                "official": "Saint Martin",
                "common": "Sveti Martin"
            },
            "ita": {
                "official": "saint Martin",
                "common": "Saint Martin"
            },
            "jpn": {
                "official": "サンマルタン島",
                "common": "サン・マルタン（フランス領）"
            },
            "nld": {
                "official": "Saint Martin",
                "common": "Saint-Martin"
            },
            "por": {
                "official": "saint Martin",
                "common": "São Martinho"
            },
            "rus": {
                "official": "Сен-Мартен",
                "common": "Сен-Мартен"
            },
            "spa": {
                "official": "Saint Martin",
                "common": "Saint Martin"
            },
            "fin": {
                "official": "Saint-Martin",
                "common": "Saint-Martin"
            }
        },
        "latlng": [
            18.08333333,
            -63.95
        ],
        "demonym": "Saint Martin Islander",
        "landlocked": false,
        "borders": [
            "SXM"
        ],
        "area": 53
    },
    {
        "name": {
            "common": "Morocco",
            "official": "Kingdom of Morocco",
            "native": {
                "ara": {
                    "official": "المملكة المغربية",
                    "common": "المغرب"
                },
                "ber": {
                    "official": "ⵜⴰⴳⵍⴷⵉⵜ ⵏ ⵍⵎⵖⵔⵉⴱ",
                    "common": "ⵍⵎⴰⵖⵔⵉⴱ"
                }
            }
        },
        "tld": [
            ".ma",
            "المغرب."
        ],
        "cca2": "MA",
        "ccn3": "504",
        "cca3": "MAR",
        "cioc": "MAR",
        "currency": [
            "MAD"
        ],
        "callingCode": [
            "212"
        ],
        "capital": "Rabat",
        "altSpellings": [
            "MA",
            "Kingdom of Morocco",
            "Al-Mamlakah al-Maġribiyah"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic",
            "ber": "Berber"
        },
        "translations": {
            "deu": {
                "official": "Königreich Marokko",
                "common": "Marokko"
            },
            "fra": {
                "official": "Royaume du Maroc",
                "common": "Maroc"
            },
            "hrv": {
                "official": "Kraljevina Maroko",
                "common": "Maroko"
            },
            "ita": {
                "official": "Regno del Marocco",
                "common": "Marocco"
            },
            "jpn": {
                "official": "モロッコ王国",
                "common": "モロッコ"
            },
            "nld": {
                "official": "Koninkrijk Marokko",
                "common": "Marokko"
            },
            "por": {
                "official": "Reino de Marrocos",
                "common": "Marrocos"
            },
            "rus": {
                "official": "Королевство Марокко",
                "common": "Марокко"
            },
            "spa": {
                "official": "Reino de Marruecos",
                "common": "Marruecos"
            },
            "fin": {
                "official": "Marokon kuningaskunta",
                "common": "Marokko"
            }
        },
        "latlng": [
            32,
            -5
        ],
        "demonym": "Moroccan",
        "landlocked": false,
        "borders": [
            "DZA",
            "ESH",
            "ESP"
        ],
        "area": 446550
    },
    {
        "name": {
            "common": "Monaco",
            "official": "Principality of Monaco",
            "native": {
                "fra": {
                    "official": "Principauté de Monaco",
                    "common": "Monaco"
                }
            }
        },
        "tld": [
            ".mc"
        ],
        "cca2": "MC",
        "ccn3": "492",
        "cca3": "MCO",
        "cioc": "MON",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "377"
        ],
        "capital": "Monaco",
        "altSpellings": [
            "MC",
            "Principality of Monaco",
            "Principauté de Monaco"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Fürstentum Monaco",
                "common": "Monaco"
            },
            "fra": {
                "official": "Principauté de Monaco",
                "common": "Monaco"
            },
            "hrv": {
                "official": "Kneževina Monako",
                "common": "Monako"
            },
            "ita": {
                "official": "Principato di Monaco",
                "common": "Principato di Monaco"
            },
            "jpn": {
                "official": "モナコ公国",
                "common": "モナコ"
            },
            "nld": {
                "official": "Vorstendom Monaco",
                "common": "Monaco"
            },
            "por": {
                "official": "Principado do Mónaco",
                "common": "Mónaco"
            },
            "rus": {
                "official": "Княжество Монако",
                "common": "Монако"
            },
            "spa": {
                "official": "Principado de Mónaco",
                "common": "Mónaco"
            },
            "fin": {
                "official": "Monacon ruhtinaskunta",
                "common": "Monaco"
            }
        },
        "latlng": [
            43.73333333,
            7.4
        ],
        "demonym": "Monegasque",
        "landlocked": false,
        "borders": [
            "FRA"
        ],
        "area": 2.02
    },
    {
        "name": {
            "common": "Moldova",
            "official": "Republic of Moldova",
            "native": {
                "ron": {
                    "official": "Republica Moldova",
                    "common": "Moldova"
                }
            }
        },
        "tld": [
            ".md"
        ],
        "cca2": "MD",
        "ccn3": "498",
        "cca3": "MDA",
        "cioc": "MDA",
        "currency": [
            "MDL"
        ],
        "callingCode": [
            "373"
        ],
        "capital": "Chișinău",
        "altSpellings": [
            "MD",
            "Moldova, Republic of",
            "Republic of Moldova",
            "Republica Moldova"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "ron": "Moldavian"
        },
        "translations": {
            "deu": {
                "official": "Republik Moldau",
                "common": "Moldawie"
            },
            "fra": {
                "official": "République de Moldavie",
                "common": "Moldavie"
            },
            "hrv": {
                "official": "Moldavija",
                "common": "Moldova"
            },
            "ita": {
                "official": "Repubblica di Moldova",
                "common": "Moldavia"
            },
            "jpn": {
                "official": "モルドバ共和国",
                "common": "モルドバ共和国"
            },
            "nld": {
                "official": "Republiek Moldavië",
                "common": "Moldavië"
            },
            "por": {
                "official": "República da Moldávia",
                "common": "Moldávia"
            },
            "rus": {
                "official": "Молдова",
                "common": "Молдавия"
            },
            "spa": {
                "official": "República de Moldova",
                "common": "Moldavia"
            },
            "fin": {
                "official": "Moldovan tasavalta",
                "common": "Moldova"
            }
        },
        "latlng": [
            47,
            29
        ],
        "demonym": "Moldovan",
        "landlocked": true,
        "borders": [
            "ROU",
            "UKR"
        ],
        "area": 33846
    },
    {
        "name": {
            "common": "Madagascar",
            "official": "Republic of Madagascar",
            "native": {
                "fra": {
                    "official": "République de Madagascar",
                    "common": "Madagascar"
                },
                "mlg": {
                    "official": "Repoblikan'i Madagasikara",
                    "common": "Madagasikara"
                }
            }
        },
        "tld": [
            ".mg"
        ],
        "cca2": "MG",
        "ccn3": "450",
        "cca3": "MDG",
        "cioc": "MAD",
        "currency": [
            "MGA"
        ],
        "callingCode": [
            "261"
        ],
        "capital": "Antananarivo",
        "altSpellings": [
            "MG",
            "Republic of Madagascar",
            "Repoblikan'i Madagasikara",
            "République de Madagascar"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "fra": "French",
            "mlg": "Malagasy"
        },
        "translations": {
            "deu": {
                "official": "Republik Madagaskar",
                "common": "Madagaskar"
            },
            "fra": {
                "official": "République de Madagascar",
                "common": "Madagascar"
            },
            "hrv": {
                "official": "Republika Madagaskar",
                "common": "Madagaskar"
            },
            "ita": {
                "official": "Repubblica del Madagascar",
                "common": "Madagascar"
            },
            "jpn": {
                "official": "マダガスカル共和国",
                "common": "マダガスカル"
            },
            "nld": {
                "official": "Republiek Madagaskar",
                "common": "Madagaskar"
            },
            "por": {
                "official": "República de Madagáscar",
                "common": "Madagáscar"
            },
            "rus": {
                "official": "Республика Мадагаскар",
                "common": "Мадагаскар"
            },
            "spa": {
                "official": "República de Madagascar",
                "common": "Madagascar"
            },
            "fin": {
                "official": "Madagaskarin tasavalta",
                "common": "Madagaskar"
            }
        },
        "latlng": [
            -20,
            47
        ],
        "demonym": "Malagasy",
        "landlocked": false,
        "borders": [],
        "area": 587041
    },
    {
        "name": {
            "common": "Maldives",
            "official": "Republic of the Maldives",
            "native": {
                "div": {
                    "official": "ދިވެހިރާއްޖޭގެ ޖުމްހޫރިއްޔާ",
                    "common": "ދިވެހިރާއްޖޭގެ"
                }
            }
        },
        "tld": [
            ".mv"
        ],
        "cca2": "MV",
        "ccn3": "462",
        "cca3": "MDV",
        "cioc": "MDV",
        "currency": [
            "MVR"
        ],
        "callingCode": [
            "960"
        ],
        "capital": "Malé",
        "altSpellings": [
            "MV",
            "Maldive Islands",
            "Republic of the Maldives",
            "Dhivehi Raajjeyge Jumhooriyya"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "div": "Maldivian"
        },
        "translations": {
            "deu": {
                "official": "Republik Malediven",
                "common": "Malediven"
            },
            "fra": {
                "official": "République des Maldives",
                "common": "Maldives"
            },
            "hrv": {
                "official": "Republika Maldivi",
                "common": "Maldivi"
            },
            "ita": {
                "official": "Repubblica delle Maldive",
                "common": "Maldive"
            },
            "jpn": {
                "official": "モルディブ共和国",
                "common": "モルディブ"
            },
            "nld": {
                "official": "Republiek van de Malediven",
                "common": "Maldiven"
            },
            "por": {
                "official": "República das Maldivas",
                "common": "Maldivas"
            },
            "rus": {
                "official": "Республика Мальдивы",
                "common": "Мальдивы"
            },
            "spa": {
                "official": "República de las Maldivas",
                "common": "Maldivas"
            },
            "fin": {
                "official": "Malediivien tasavalta",
                "common": "Malediivit"
            }
        },
        "latlng": [
            3.25,
            73
        ],
        "demonym": "Maldivan",
        "landlocked": false,
        "borders": [],
        "area": 300
    },
    {
        "name": {
            "common": "Mexico",
            "official": "United Mexican States",
            "native": {
                "spa": {
                    "official": "Estados Unidos Mexicanos",
                    "common": "México"
                }
            }
        },
        "tld": [
            ".mx"
        ],
        "cca2": "MX",
        "ccn3": "484",
        "cca3": "MEX",
        "cioc": "MEX",
        "currency": [
            "MXN"
        ],
        "callingCode": [
            "52"
        ],
        "capital": "Mexico City",
        "altSpellings": [
            "MX",
            "Mexicanos",
            "United Mexican States",
            "Estados Unidos Mexicanos"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Vereinigte Mexikanische Staaten",
                "common": "Mexiko"
            },
            "fra": {
                "official": "États-Unis du Mexique",
                "common": "Mexique"
            },
            "hrv": {
                "official": "Sjedinjene Meksičke Države",
                "common": "Meksiko"
            },
            "ita": {
                "official": "Stati Uniti del Messico",
                "common": "Messico"
            },
            "jpn": {
                "official": "メキシコ合衆国",
                "common": "メキシコ"
            },
            "nld": {
                "official": "Verenigde Mexicaanse Staten",
                "common": "Mexico"
            },
            "por": {
                "official": "Estados Unidos Mexicanos",
                "common": "México"
            },
            "rus": {
                "official": "Мексиканских Соединенных Штатов",
                "common": "Мексика"
            },
            "spa": {
                "official": "Estados Unidos Mexicanos",
                "common": "México"
            },
            "fin": {
                "official": "Meksikon yhdysvallat",
                "common": "Meksiko"
            }
        },
        "latlng": [
            23,
            -102
        ],
        "demonym": "Mexican",
        "landlocked": false,
        "borders": [
            "BLZ",
            "GTM",
            "USA"
        ],
        "area": 1964375
    },
    {
        "name": {
            "common": "Marshall Islands",
            "official": "Republic of the Marshall Islands",
            "native": {
                "eng": {
                    "official": "Republic of the Marshall Islands",
                    "common": "Marshall Islands"
                },
                "mah": {
                    "official": "Republic of the Marshall Islands",
                    "common": "M̧ajeļ"
                }
            }
        },
        "tld": [
            ".mh"
        ],
        "cca2": "MH",
        "ccn3": "584",
        "cca3": "MHL",
        "cioc": "MHL",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "692"
        ],
        "capital": "Majuro",
        "altSpellings": [
            "MH",
            "Republic of the Marshall Islands",
            "Aolepān Aorōkin M̧ajeļ"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "eng": "English",
            "mah": "Marshallese"
        },
        "translations": {
            "deu": {
                "official": "Republik der Marshall-Inseln",
                "common": "Marshallinseln"
            },
            "fra": {
                "official": "République des Îles Marshall",
                "common": "Îles Marshall"
            },
            "hrv": {
                "official": "Republika Maršalovi Otoci",
                "common": "Maršalovi Otoci"
            },
            "ita": {
                "official": "Repubblica delle Isole Marshall",
                "common": "Isole Marshall"
            },
            "jpn": {
                "official": "マーシャル諸島共和国",
                "common": "マーシャル諸島"
            },
            "nld": {
                "official": "Republiek van de Marshall-eilanden",
                "common": "Marshalleilanden"
            },
            "por": {
                "official": "República das Ilhas Marshall",
                "common": "Ilhas Marshall"
            },
            "rus": {
                "official": "Республика Маршалловы острова",
                "common": "Маршалловы Острова"
            },
            "spa": {
                "official": "República de las Islas Marshall",
                "common": "Islas Marshall"
            },
            "fin": {
                "official": "Marshallinsaarten tasavalta",
                "common": "Marshallinsaaret"
            }
        },
        "latlng": [
            9,
            168
        ],
        "demonym": "Marshallese",
        "landlocked": false,
        "borders": [],
        "area": 181
    },
    {
        "name": {
            "common": "Macedonia",
            "official": "Republic of Macedonia",
            "native": {
                "mkd": {
                    "official": "Република Македонија",
                    "common": "Македонија"
                }
            }
        },
        "tld": [
            ".mk"
        ],
        "cca2": "MK",
        "ccn3": "807",
        "cca3": "MKD",
        "cioc": "MKD",
        "currency": [
            "MKD"
        ],
        "callingCode": [
            "389"
        ],
        "capital": "Skopje",
        "altSpellings": [
            "MK",
            "Macedonia, the Former Yugoslav Republic of",
            "Republic of Macedonia",
            "Република Македонија"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "mkd": "Macedonian"
        },
        "translations": {
            "deu": {
                "official": "Republik Mazedonien",
                "common": "Mazedonien"
            },
            "fra": {
                "official": "République de Macédoine",
                "common": "Macédoine"
            },
            "hrv": {
                "official": "Republika Makedonija",
                "common": "Makedonija"
            },
            "ita": {
                "official": "Repubblica di Macedonia",
                "common": "Macedonia"
            },
            "jpn": {
                "official": "マケドニア共和国",
                "common": "マケドニア旧ユーゴスラビア共和国"
            },
            "nld": {
                "official": "Republic of Macedonia",
                "common": "Macedonië"
            },
            "por": {
                "official": "República da Macedónia",
                "common": "Macedónia"
            },
            "rus": {
                "official": "Республика Македония",
                "common": "Республика Македония"
            },
            "spa": {
                "official": "República de Macedonia",
                "common": "Macedonia"
            },
            "fin": {
                "official": "Makedonian tasavalta",
                "common": "Makedonia"
            }
        },
        "latlng": [
            41.83333333,
            22
        ],
        "demonym": "Macedonian",
        "landlocked": true,
        "borders": [
            "ALB",
            "BGR",
            "GRC",
            "KOS",
            "SRB"
        ],
        "area": 25713
    },
    {
        "name": {
            "common": "Mali",
            "official": "Republic of Mali",
            "native": {
                "fra": {
                    "official": "République du Mali",
                    "common": "Mali"
                }
            }
        },
        "tld": [
            ".ml"
        ],
        "cca2": "ML",
        "ccn3": "466",
        "cca3": "MLI",
        "cioc": "MLI",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "223"
        ],
        "capital": "Bamako",
        "altSpellings": [
            "ML",
            "Republic of Mali",
            "République du Mali"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Mali",
                "common": "Mali"
            },
            "fra": {
                "official": "République du Mali",
                "common": "Mali"
            },
            "hrv": {
                "official": "Republika Mali",
                "common": "Mali"
            },
            "ita": {
                "official": "Repubblica del Mali",
                "common": "Mali"
            },
            "jpn": {
                "official": "マリ共和国",
                "common": "マリ"
            },
            "nld": {
                "official": "Republiek Mali",
                "common": "Mali"
            },
            "por": {
                "official": "República do Mali",
                "common": "Mali"
            },
            "rus": {
                "official": "Республика Мали",
                "common": "Мали"
            },
            "spa": {
                "official": "República de Malí",
                "common": "Mali"
            },
            "fin": {
                "official": "Malin tasavalta",
                "common": "Mali"
            }
        },
        "latlng": [
            17,
            -4
        ],
        "demonym": "Malian",
        "landlocked": true,
        "borders": [
            "DZA",
            "BFA",
            "GIN",
            "CIV",
            "MRT",
            "NER",
            "SEN"
        ],
        "area": 1240192
    },
    {
        "name": {
            "common": "Malta",
            "official": "Republic of Malta",
            "native": {
                "eng": {
                    "official": "Republic of Malta",
                    "common": "Malta"
                },
                "mlt": {
                    "official": "Repubblika ta ' Malta",
                    "common": "Malta"
                }
            }
        },
        "tld": [
            ".mt"
        ],
        "cca2": "MT",
        "ccn3": "470",
        "cca3": "MLT",
        "cioc": "MLT",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "356"
        ],
        "capital": "Valletta",
        "altSpellings": [
            "MT",
            "Republic of Malta",
            "Repubblika ta' Malta"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "eng": "English",
            "mlt": "Maltese"
        },
        "translations": {
            "deu": {
                "official": "Republik Malta",
                "common": "Malta"
            },
            "fra": {
                "official": "République de Malte",
                "common": "Malte"
            },
            "hrv": {
                "official": "Republika Malta",
                "common": "Malta"
            },
            "ita": {
                "official": "Repubblica di Malta",
                "common": "Malta"
            },
            "jpn": {
                "official": "マルタ共和国",
                "common": "マルタ"
            },
            "nld": {
                "official": "Republiek Malta",
                "common": "Malta"
            },
            "por": {
                "official": "República de Malta",
                "common": "Malta"
            },
            "rus": {
                "official": "Республика Мальта",
                "common": "Мальта"
            },
            "spa": {
                "official": "República de Malta",
                "common": "Malta"
            },
            "fin": {
                "official": "Maltan tasavalta",
                "common": "Malta"
            }
        },
        "latlng": [
            35.83333333,
            14.58333333
        ],
        "demonym": "Maltese",
        "landlocked": false,
        "borders": [],
        "area": 316
    },
    {
        "name": {
            "common": "Myanmar",
            "official": "Republic of the Union of Myanmar",
            "native": {
                "mya": {
                    "official": "ပ.ည.ထောင.စု သမ္မတ မ.န.မာနိုင.ငံတော.",
                    "common": "မ.န.မာ"
                }
            }
        },
        "tld": [
            ".mm"
        ],
        "cca2": "MM",
        "ccn3": "104",
        "cca3": "MMR",
        "cioc": "MYA",
        "currency": [
            "MMK"
        ],
        "callingCode": [
            "95"
        ],
        "capital": "Naypyidaw",
        "altSpellings": [
            "MM",
            "Burma",
            "Republic of the Union of Myanmar",
            "Pyidaunzu Thanmăda Myăma Nainngandaw"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "mya": "Burmese"
        },
        "translations": {
            "deu": {
                "official": "Republik der Union von Myanmar",
                "common": "Myanmar"
            },
            "fra": {
                "official": "République de l'Union du Myanmar",
                "common": "Birmanie"
            },
            "hrv": {
                "official": "Republika Unije Mijanmar",
                "common": "Mijanmar"
            },
            "ita": {
                "official": "Repubblica dell'Unione di Myanmar",
                "common": "Birmania"
            },
            "jpn": {
                "official": "ミャンマー連邦共和国",
                "common": "ミャンマー"
            },
            "nld": {
                "official": "Republiek van de Unie van Myanmar",
                "common": "Myanmar"
            },
            "por": {
                "official": "República da União de Myanmar",
                "common": "Myanmar"
            },
            "rus": {
                "official": "Республика Союза Мьянма",
                "common": "Мьянма"
            },
            "spa": {
                "official": "República de la Unión de Myanmar",
                "common": "Myanmar"
            },
            "fin": {
                "official": "Myanmarin liiton tasavalta",
                "common": "Myanmar"
            }
        },
        "latlng": [
            22,
            98
        ],
        "demonym": "Bamar",
        "landlocked": false,
        "borders": [
            "BGD",
            "CHN",
            "IND",
            "LAO",
            "THA"
        ],
        "area": 676578
    },
    {
        "name": {
            "common": "Montenegro",
            "official": "Montenegro",
            "native": {
                "srp": {
                    "official": "Црна Гора",
                    "common": "Црна Гора"
                }
            }
        },
        "tld": [
            ".me"
        ],
        "cca2": "ME",
        "ccn3": "499",
        "cca3": "MNE",
        "cioc": "MNE",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "382"
        ],
        "capital": "Podgorica",
        "altSpellings": [
            "ME",
            "Crna Gora"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "srp": "Montenegrin"
        },
        "translations": {
            "deu": {
                "official": "Montenegro",
                "common": "Montenegro"
            },
            "fra": {
                "official": "Monténégro",
                "common": "Monténégro"
            },
            "hrv": {
                "official": "Crna Gora",
                "common": "Crna Gora"
            },
            "ita": {
                "official": "Montenegro",
                "common": "Montenegro"
            },
            "jpn": {
                "official": "モンテネグロ",
                "common": "モンテネグロ"
            },
            "nld": {
                "official": "Montenegro",
                "common": "Montenegro"
            },
            "por": {
                "official": "Montenegro",
                "common": "Montenegro"
            },
            "rus": {
                "official": "Черногория",
                "common": "Черногория"
            },
            "spa": {
                "official": "Montenegro",
                "common": "Montenegro"
            },
            "fin": {
                "official": "Montenegro",
                "common": "Montenegro"
            }
        },
        "latlng": [
            42.5,
            19.3
        ],
        "demonym": "Montenegrin",
        "landlocked": false,
        "borders": [
            "ALB",
            "BIH",
            "HRV",
            "KOS",
            "SRB"
        ],
        "area": 13812
    },
    {
        "name": {
            "common": "Mongolia",
            "official": "Mongolia",
            "native": {
                "mon": {
                    "official": "Монгол улс",
                    "common": "Монгол улс"
                }
            }
        },
        "tld": [
            ".mn"
        ],
        "cca2": "MN",
        "ccn3": "496",
        "cca3": "MNG",
        "cioc": "MGL",
        "currency": [
            "MNT"
        ],
        "callingCode": [
            "976"
        ],
        "capital": "Ulan Bator",
        "altSpellings": [
            "MN"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "mon": "Mongolian"
        },
        "translations": {
            "deu": {
                "official": "Mongolei",
                "common": "Mongolei"
            },
            "fra": {
                "official": "Mongolie",
                "common": "Mongolie"
            },
            "hrv": {
                "official": "Mongolija",
                "common": "Mongolija"
            },
            "ita": {
                "official": "Mongolia",
                "common": "Mongolia"
            },
            "jpn": {
                "official": "モンゴル",
                "common": "モンゴル"
            },
            "nld": {
                "official": "Mongolië",
                "common": "Mongolië"
            },
            "por": {
                "official": "Mongólia",
                "common": "Mongólia"
            },
            "rus": {
                "official": "Монголия",
                "common": "Монголия"
            },
            "spa": {
                "official": "Mongolia",
                "common": "Mongolia"
            },
            "fin": {
                "official": "Mongolian tasavalta",
                "common": "Mongolia"
            }
        },
        "latlng": [
            46,
            105
        ],
        "demonym": "Mongolian",
        "landlocked": true,
        "borders": [
            "CHN",
            "RUS"
        ],
        "area": 1564110
    },
    {
        "name": {
            "common": "Northern Mariana Islands",
            "official": "Commonwealth of the Northern Mariana Islands",
            "native": {
                "cal": {
                    "official": "Commonwealth of the Northern Mariana Islands",
                    "common": "Northern Mariana Islands"
                },
                "cha": {
                    "official": "Sankattan Siha Na Islas Mariånas",
                    "common": "Na Islas Mariånas"
                },
                "eng": {
                    "official": "Commonwealth of the Northern Mariana Islands",
                    "common": "Northern Mariana Islands"
                }
            }
        },
        "tld": [
            ".mp"
        ],
        "cca2": "MP",
        "ccn3": "580",
        "cca3": "MNP",
        "cioc": "",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1670"
        ],
        "capital": "Saipan",
        "altSpellings": [
            "MP",
            "Commonwealth of the Northern Mariana Islands",
            "Sankattan Siha Na Islas Mariånas"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "cal": "Carolinian",
            "cha": "Chamorro",
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Commonwealth der Nördlichen Marianen",
                "common": "Nördliche Marianen"
            },
            "fra": {
                "official": "Commonwealth des îles Mariannes du Nord",
                "common": "Îles Mariannes du Nord"
            },
            "hrv": {
                "official": "Zajednica je Sjeverni Marijanski otoci",
                "common": "Sjevernomarijanski otoci"
            },
            "ita": {
                "official": "Commonwealth delle Isole Marianne Settentrionali",
                "common": "Isole Marianne Settentrionali"
            },
            "jpn": {
                "official": "北マリアナ諸島",
                "common": "北マリアナ諸島"
            },
            "nld": {
                "official": "Commonwealth van de Noordelijke Marianen",
                "common": "Noordelijke Marianeneilanden"
            },
            "por": {
                "official": "Comunidade das Ilhas Marianas do Norte",
                "common": "Marianas Setentrionais"
            },
            "rus": {
                "official": "Содружество Северных Марианских островов",
                "common": "Северные Марианские острова"
            },
            "spa": {
                "official": "Mancomunidad de las Islas Marianas del Norte",
                "common": "Islas Marianas del Norte"
            },
            "fin": {
                "official": "Pohjois-Mariaanit",
                "common": "Pohjois-Mariaanit"
            }
        },
        "latlng": [
            15.2,
            145.75
        ],
        "demonym": "American",
        "landlocked": false,
        "borders": [],
        "area": 464
    },
    {
        "name": {
            "common": "Mozambique",
            "official": "Republic of Mozambique",
            "native": {
                "por": {
                    "official": "República de Moçambique",
                    "common": "Moçambique"
                }
            }
        },
        "tld": [
            ".mz"
        ],
        "cca2": "MZ",
        "ccn3": "508",
        "cca3": "MOZ",
        "cioc": "MOZ",
        "currency": [
            "MZN"
        ],
        "callingCode": [
            "258"
        ],
        "capital": "Maputo",
        "altSpellings": [
            "MZ",
            "Republic of Mozambique",
            "República de Moçambique"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "deu": {
                "official": "Republik Mosambik",
                "common": "Mosambik"
            },
            "fra": {
                "official": "République du Mozambique",
                "common": "Mozambique"
            },
            "hrv": {
                "official": "Republika Mozambiku",
                "common": "Mozambik"
            },
            "ita": {
                "official": "Repubblica del Mozambico",
                "common": "Mozambico"
            },
            "jpn": {
                "official": "モザンビーク共和国",
                "common": "モザンビーク"
            },
            "nld": {
                "official": "Republiek Mozambique",
                "common": "Mozambique"
            },
            "por": {
                "official": "República de Moçambique",
                "common": "Moçambique"
            },
            "rus": {
                "official": "Республика Мозамбик",
                "common": "Мозамбик"
            },
            "spa": {
                "official": "República de Mozambique",
                "common": "Mozambique"
            },
            "fin": {
                "official": "Mosambikin tasavalta",
                "common": "Mosambik"
            }
        },
        "latlng": [
            -18.25,
            35
        ],
        "demonym": "Mozambican",
        "landlocked": false,
        "borders": [
            "MWI",
            "ZAF",
            "SWZ",
            "TZA",
            "ZMB",
            "ZWE"
        ],
        "area": 801590
    },
    {
        "name": {
            "common": "Mauritania",
            "official": "Islamic Republic of Mauritania",
            "native": {
                "ara": {
                    "official": "الجمهورية الإسلامية الموريتانية",
                    "common": "موريتانيا"
                }
            }
        },
        "tld": [
            ".mr"
        ],
        "cca2": "MR",
        "ccn3": "478",
        "cca3": "MRT",
        "cioc": "MTN",
        "currency": [
            "MRO"
        ],
        "callingCode": [
            "222"
        ],
        "capital": "Nouakchott",
        "altSpellings": [
            "MR",
            "Islamic Republic of Mauritania",
            "al-Jumhūriyyah al-ʾIslāmiyyah al-Mūrītāniyyah"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Islamische Republik Mauretanien",
                "common": "Mauretanien"
            },
            "fra": {
                "official": "République islamique de Mauritanie",
                "common": "Mauritanie"
            },
            "hrv": {
                "official": "Islamska Republika Mauritanija",
                "common": "Mauritanija"
            },
            "ita": {
                "official": "Repubblica islamica di Mauritania",
                "common": "Mauritania"
            },
            "jpn": {
                "official": "モーリタニア·イスラム共和国",
                "common": "モーリタニア"
            },
            "nld": {
                "official": "Islamitische Republiek Mauritanië",
                "common": "Mauritanië"
            },
            "por": {
                "official": "República Islâmica da Mauritânia",
                "common": "Mauritânia"
            },
            "rus": {
                "official": "Исламская Республика Мавритания",
                "common": "Мавритания"
            },
            "spa": {
                "official": "República Islámica de Mauritania",
                "common": "Mauritania"
            },
            "fin": {
                "official": "Mauritanian islamilainen tasavalta",
                "common": "Mauritania"
            }
        },
        "latlng": [
            20,
            -12
        ],
        "demonym": "Mauritanian",
        "landlocked": false,
        "borders": [
            "DZA",
            "MLI",
            "SEN",
            "ESH"
        ],
        "area": 1030700
    },
    {
        "name": {
            "common": "Montserrat",
            "official": "Montserrat",
            "native": {
                "eng": {
                    "official": "Montserrat",
                    "common": "Montserrat"
                }
            }
        },
        "tld": [
            ".ms"
        ],
        "cca2": "MS",
        "ccn3": "500",
        "cca3": "MSR",
        "cioc": "",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1664"
        ],
        "capital": "Plymouth",
        "altSpellings": [
            "MS"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "fra": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "hrv": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "ita": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "jpn": {
                "official": "モントセラト",
                "common": "モントセラト"
            },
            "nld": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "por": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "rus": {
                "official": "Монтсеррат",
                "common": "Монтсеррат"
            },
            "spa": {
                "official": "Montserrat",
                "common": "Montserrat"
            },
            "fin": {
                "official": "Montserrat",
                "common": "Montserrat"
            }
        },
        "latlng": [
            16.75,
            -62.2
        ],
        "demonym": "Montserratian",
        "landlocked": false,
        "borders": [],
        "area": 102
    },
    {
        "name": {
            "common": "Martinique",
            "official": "Martinique",
            "native": {
                "fra": {
                    "official": "Martinique",
                    "common": "Martinique"
                }
            }
        },
        "tld": [
            ".mq"
        ],
        "cca2": "MQ",
        "ccn3": "474",
        "cca3": "MTQ",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "596"
        ],
        "capital": "Fort-de-France",
        "altSpellings": [
            "MQ"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Martinique",
                "common": "Martinique"
            },
            "fra": {
                "official": "Martinique",
                "common": "Martinique"
            },
            "hrv": {
                "official": "Martinique",
                "common": "Martinique"
            },
            "ita": {
                "official": "Martinique",
                "common": "Martinica"
            },
            "jpn": {
                "official": "マルティニーク島",
                "common": "マルティニーク"
            },
            "nld": {
                "official": "Martinique",
                "common": "Martinique"
            },
            "por": {
                "official": "Martinique",
                "common": "Martinica"
            },
            "rus": {
                "official": "Мартиника",
                "common": "Мартиника"
            },
            "spa": {
                "official": "Martinica",
                "common": "Martinica"
            },
            "fin": {
                "official": "Martinique",
                "common": "Martinique"
            }
        },
        "latlng": [
            14.666667,
            -61
        ],
        "demonym": "French",
        "landlocked": false,
        "borders": [],
        "area": 1128
    },
    {
        "name": {
            "common": "Mauritius",
            "official": "Republic of Mauritius",
            "native": {
                "eng": {
                    "official": "Republic of Mauritius",
                    "common": "Mauritius"
                },
                "fra": {
                    "official": "République de Maurice",
                    "common": "Maurice"
                },
                "mfe": {
                    "official": "Republik Moris",
                    "common": "Moris"
                }
            }
        },
        "tld": [
            ".mu"
        ],
        "cca2": "MU",
        "ccn3": "480",
        "cca3": "MUS",
        "cioc": "MRI",
        "currency": [
            "MUR"
        ],
        "callingCode": [
            "230"
        ],
        "capital": "Port Louis",
        "altSpellings": [
            "MU",
            "Republic of Mauritius",
            "République de Maurice"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "fra": "French",
            "mfe": "Mauritian Creole"
        },
        "translations": {
            "deu": {
                "official": "Republik Mauritius",
                "common": "Mauritius"
            },
            "fra": {
                "official": "République de Maurice",
                "common": "Île Maurice"
            },
            "hrv": {
                "official": "Republika Mauricijus",
                "common": "Mauricijus"
            },
            "ita": {
                "official": "Repubblica di Mauritius",
                "common": "Mauritius"
            },
            "jpn": {
                "official": "モーリシャス共和国",
                "common": "モーリシャス"
            },
            "nld": {
                "official": "Republiek Mauritius",
                "common": "Mauritius"
            },
            "por": {
                "official": "República das Maurícias",
                "common": "Maurício"
            },
            "rus": {
                "official": "Республика Маврикий",
                "common": "Маврикий"
            },
            "spa": {
                "official": "República de Mauricio",
                "common": "Mauricio"
            },
            "fin": {
                "official": "Mauritiuksen tasavalta",
                "common": "Mauritius"
            }
        },
        "latlng": [
            -20.28333333,
            57.55
        ],
        "demonym": "Mauritian",
        "landlocked": false,
        "borders": [],
        "area": 2040
    },
    {
        "name": {
            "common": "Malawi",
            "official": "Republic of Malawi",
            "native": {
                "eng": {
                    "official": "Republic of Malawi",
                    "common": "Malawi"
                },
                "nya": {
                    "official": "Chalo cha Malawi, Dziko la Malaŵi",
                    "common": "Malaŵi"
                }
            }
        },
        "tld": [
            ".mw"
        ],
        "cca2": "MW",
        "ccn3": "454",
        "cca3": "MWI",
        "cioc": "MAW",
        "currency": [
            "MWK"
        ],
        "callingCode": [
            "265"
        ],
        "capital": "Lilongwe",
        "altSpellings": [
            "MW",
            "Republic of Malawi"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "nya": "Chewa"
        },
        "translations": {
            "deu": {
                "official": "Republik Malawi",
                "common": "Malawi"
            },
            "fra": {
                "official": "République du Malawi",
                "common": "Malawi"
            },
            "hrv": {
                "official": "Republika Malavi",
                "common": "Malavi"
            },
            "ita": {
                "official": "Repubblica del Malawi",
                "common": "Malawi"
            },
            "jpn": {
                "official": "マラウイ共和国",
                "common": "マラウイ"
            },
            "nld": {
                "official": "Republiek Malawi",
                "common": "Malawi"
            },
            "por": {
                "official": "República do Malawi",
                "common": "Malawi"
            },
            "rus": {
                "official": "Республика Малави",
                "common": "Малави"
            },
            "spa": {
                "official": "República de Malawi",
                "common": "Malawi"
            },
            "fin": {
                "official": "Malawin tasavalta",
                "common": "Malawi"
            }
        },
        "latlng": [
            -13.5,
            34
        ],
        "demonym": "Malawian",
        "landlocked": true,
        "borders": [
            "MOZ",
            "TZA",
            "ZMB"
        ],
        "area": 118484
    },
    {
        "name": {
            "common": "Malaysia",
            "official": "Malaysia",
            "native": {
                "eng": {
                    "official": "Malaysia",
                    "common": "Malaysia"
                },
                "msa": {
                    "official": "مليسيا",
                    "common": "مليسيا"
                }
            }
        },
        "tld": [
            ".my"
        ],
        "cca2": "MY",
        "ccn3": "458",
        "cca3": "MYS",
        "cioc": "MAS",
        "currency": [
            "MYR"
        ],
        "callingCode": [
            "60"
        ],
        "capital": "Kuala Lumpur",
        "altSpellings": [
            "MY"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "eng": "English",
            "msa": "Malay"
        },
        "translations": {
            "deu": {
                "official": "Malaysia",
                "common": "Malaysia"
            },
            "fra": {
                "official": "Fédération de Malaisie",
                "common": "Malaisie"
            },
            "hrv": {
                "official": "Malezija",
                "common": "Malezija"
            },
            "ita": {
                "official": "Malaysia",
                "common": "Malesia"
            },
            "jpn": {
                "official": "マレーシア",
                "common": "マレーシア"
            },
            "nld": {
                "official": "Maleisië",
                "common": "Maleisië"
            },
            "por": {
                "official": "Malásia",
                "common": "Malásia"
            },
            "rus": {
                "official": "Малайзия",
                "common": "Малайзия"
            },
            "spa": {
                "official": "Malasia",
                "common": "Malasia"
            },
            "fin": {
                "official": "Malesia",
                "common": "Malesia"
            }
        },
        "latlng": [
            2.5,
            112.5
        ],
        "demonym": "Malaysian",
        "landlocked": false,
        "borders": [
            "BRN",
            "IDN",
            "THA"
        ],
        "area": 330803
    },
    {
        "name": {
            "common": "Mayotte",
            "official": "Department of Mayotte",
            "native": {
                "fra": {
                    "official": "Département de Mayotte",
                    "common": "Mayotte"
                }
            }
        },
        "tld": [
            ".yt"
        ],
        "cca2": "YT",
        "ccn3": "175",
        "cca3": "MYT",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "262"
        ],
        "capital": "Mamoudzou",
        "altSpellings": [
            "YT",
            "Department of Mayotte",
            "Département de Mayotte"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Übersee-Département Mayotte",
                "common": "Mayotte"
            },
            "fra": {
                "official": "Département de Mayotte",
                "common": "Mayotte"
            },
            "hrv": {
                "official": "Odjel Mayotte",
                "common": "Mayotte"
            },
            "ita": {
                "official": "Dipartimento di Mayotte",
                "common": "Mayotte"
            },
            "jpn": {
                "official": "マヨット科",
                "common": "マヨット"
            },
            "nld": {
                "official": "Afdeling Mayotte",
                "common": "Mayotte"
            },
            "por": {
                "official": "Departamento de Mayotte",
                "common": "Mayotte"
            },
            "rus": {
                "official": "Департамент Майотта",
                "common": "Майотта"
            },
            "spa": {
                "official": "Departamento de Mayotte",
                "common": "Mayotte"
            },
            "fin": {
                "official": "Mayotte",
                "common": "Mayotte"
            }
        },
        "latlng": [
            -12.83333333,
            45.16666666
        ],
        "demonym": "Mahoran",
        "landlocked": false,
        "borders": [],
        "area": 374
    },
    {
        "name": {
            "common": "Namibia",
            "official": "Republic of Namibia",
            "native": {
                "afr": {
                    "official": "Republiek van Namibië",
                    "common": "Namibië"
                },
                "deu": {
                    "official": "Republik Namibia",
                    "common": "Namibia"
                },
                "eng": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "her": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "hgm": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "kwn": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "loz": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "ndo": {
                    "official": "Republic of Namibia",
                    "common": "Namibia"
                },
                "tsn": {
                    "official": "Lefatshe la Namibia",
                    "common": "Namibia"
                }
            }
        },
        "tld": [
            ".na"
        ],
        "cca2": "NA",
        "ccn3": "516",
        "cca3": "NAM",
        "cioc": "NAM",
        "currency": [
            "NAD",
            "ZAR"
        ],
        "callingCode": [
            "264"
        ],
        "capital": "Windhoek",
        "altSpellings": [
            "NA",
            "Namibië",
            "Republic of Namibia"
        ],
        "region": "Africa",
        "subregion": "Southern Africa",
        "languages": {
            "afr": "Afrikaans",
            "deu": "German",
            "eng": "English",
            "her": "Herero",
            "hgm": "Khoekhoe",
            "kwn": "Kwangali",
            "loz": "Lozi",
            "ndo": "Ndonga",
            "tsn": "Tswana"
        },
        "translations": {
            "deu": {
                "official": "Republik Namibia",
                "common": "Namibia"
            },
            "fra": {
                "official": "République de Namibie",
                "common": "Namibie"
            },
            "hrv": {
                "official": "Republika Namibija",
                "common": "Namibija"
            },
            "ita": {
                "official": "Repubblica di Namibia",
                "common": "Namibia"
            },
            "jpn": {
                "official": "ナミビア共和国",
                "common": "ナミビア"
            },
            "nld": {
                "official": "Republiek Namibië",
                "common": "Namibië"
            },
            "por": {
                "official": "República da Namíbia",
                "common": "Namíbia"
            },
            "rus": {
                "official": "Республика Намибия",
                "common": "Намибия"
            },
            "spa": {
                "official": "República de Namibia",
                "common": "Namibia"
            },
            "fin": {
                "official": "Namibian tasavalta",
                "common": "Namibia"
            }
        },
        "latlng": [
            -22,
            17
        ],
        "demonym": "Namibian",
        "landlocked": false,
        "borders": [
            "AGO",
            "BWA",
            "ZAF",
            "ZMB"
        ],
        "area": 825615
    },
    {
        "name": {
            "common": "New Caledonia",
            "official": "New Caledonia",
            "native": {
                "fra": {
                    "official": "Nouvelle-Calédonie",
                    "common": "Nouvelle-Calédonie"
                }
            }
        },
        "tld": [
            ".nc"
        ],
        "cca2": "NC",
        "ccn3": "540",
        "cca3": "NCL",
        "cioc": "",
        "currency": [
            "XPF"
        ],
        "callingCode": [
            "687"
        ],
        "capital": "Nouméa",
        "altSpellings": [
            "NC"
        ],
        "region": "Oceania",
        "subregion": "Melanesia",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Neukaledonien",
                "common": "Neukaledonien"
            },
            "fra": {
                "official": "Nouvelle-Calédonie",
                "common": "Nouvelle-Calédonie"
            },
            "hrv": {
                "official": "Nova Kaledonija",
                "common": "Nova Kaledonija"
            },
            "ita": {
                "official": "Nuova Caledonia",
                "common": "Nuova Caledonia"
            },
            "jpn": {
                "official": "ニューカレドニア",
                "common": "ニューカレドニア"
            },
            "nld": {
                "official": "nieuw -Caledonië",
                "common": "Nieuw-Caledonië"
            },
            "por": {
                "official": "New Caledonia",
                "common": "Nova Caledónia"
            },
            "rus": {
                "official": "Новая Каледония",
                "common": "Новая Каледония"
            },
            "spa": {
                "official": "nueva Caledonia",
                "common": "Nueva Caledonia"
            },
            "fin": {
                "official": "Uusi-Kaledonia",
                "common": "Uusi-Kaledonia"
            }
        },
        "latlng": [
            -21.5,
            165.5
        ],
        "demonym": "New Caledonian",
        "landlocked": false,
        "borders": [],
        "area": 18575
    },
    {
        "name": {
            "common": "Niger",
            "official": "Republic of Niger",
            "native": {
                "fra": {
                    "official": "République du Niger",
                    "common": "Niger"
                }
            }
        },
        "tld": [
            ".ne"
        ],
        "cca2": "NE",
        "ccn3": "562",
        "cca3": "NER",
        "cioc": "NIG",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "227"
        ],
        "capital": "Niamey",
        "altSpellings": [
            "NE",
            "Nijar"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Niger",
                "common": "Niger"
            },
            "fra": {
                "official": "République du Niger",
                "common": "Niger"
            },
            "hrv": {
                "official": "Republika Niger",
                "common": "Niger"
            },
            "ita": {
                "official": "Repubblica del Niger",
                "common": "Niger"
            },
            "jpn": {
                "official": "ニジェール共和国",
                "common": "ニジェール"
            },
            "nld": {
                "official": "Republiek Niger",
                "common": "Niger"
            },
            "por": {
                "official": "República do Níger",
                "common": "Níger"
            },
            "rus": {
                "official": "Республика Нигер",
                "common": "Нигер"
            },
            "spa": {
                "official": "República de Níger",
                "common": "Níger"
            },
            "fin": {
                "official": "Nigerin tasavalta",
                "common": "Niger"
            }
        },
        "latlng": [
            16,
            8
        ],
        "demonym": "Nigerien",
        "landlocked": true,
        "borders": [
            "DZA",
            "BEN",
            "BFA",
            "TCD",
            "LBY",
            "MLI",
            "NGA"
        ],
        "area": 1267000
    },
    {
        "name": {
            "common": "Norfolk Island",
            "official": "Territory of Norfolk Island",
            "native": {
                "eng": {
                    "official": "Territory of Norfolk Island",
                    "common": "Norfolk Island"
                },
                "pih": {
                    "official": "Teratri of Norf'k Ailen",
                    "common": "Norf'k Ailen"
                }
            }
        },
        "tld": [
            ".nf"
        ],
        "cca2": "NF",
        "ccn3": "574",
        "cca3": "NFK",
        "cioc": "",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "672"
        ],
        "capital": "Kingston",
        "altSpellings": [
            "NF",
            "Territory of Norfolk Island",
            "Teratri of Norf'k Ailen"
        ],
        "region": "Oceania",
        "subregion": "Australia and New Zealand",
        "languages": {
            "eng": "English",
            "pih": "Norfuk"
        },
        "translations": {
            "deu": {
                "official": "Gebiet der Norfolk-Insel",
                "common": "Norfolkinsel"
            },
            "fra": {
                "official": "Territoire de l'île Norfolk",
                "common": "Île Norfolk"
            },
            "hrv": {
                "official": "Teritorij Norfolk Island",
                "common": "Otok Norfolk"
            },
            "ita": {
                "official": "Territorio di Norfolk Island",
                "common": "Isola Norfolk"
            },
            "jpn": {
                "official": "ノーフォーク島の領土",
                "common": "ノーフォーク島"
            },
            "nld": {
                "official": "Grondgebied van Norfolk Island",
                "common": "Norfolkeiland"
            },
            "por": {
                "official": "Território da Ilha Norfolk",
                "common": "Ilha Norfolk"
            },
            "rus": {
                "official": "Территория острова Норфолк",
                "common": "Норфолк"
            },
            "spa": {
                "official": "Territorio de la Isla Norfolk",
                "common": "Isla de Norfolk"
            },
            "fin": {
                "official": "Norfolkinsaaren territorio",
                "common": "Norfolkinsaari"
            }
        },
        "latlng": [
            -29.03333333,
            167.95
        ],
        "demonym": "Norfolk Islander",
        "landlocked": false,
        "borders": [],
        "area": 36
    },
    {
        "name": {
            "common": "Nigeria",
            "official": "Federal Republic of Nigeria",
            "native": {
                "eng": {
                    "official": "Federal Republic of Nigeria",
                    "common": "Nigeria"
                }
            }
        },
        "tld": [
            ".ng"
        ],
        "cca2": "NG",
        "ccn3": "566",
        "cca3": "NGA",
        "cioc": "NGR",
        "currency": [
            "NGN"
        ],
        "callingCode": [
            "234"
        ],
        "capital": "Abuja",
        "altSpellings": [
            "NG",
            "Nijeriya",
            "Naíjíríà",
            "Federal Republic of Nigeria"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Bundesrepublik Nigeria",
                "common": "Nigeria"
            },
            "fra": {
                "official": "République fédérale du Nigeria",
                "common": "Nigéria"
            },
            "hrv": {
                "official": "Savezna Republika Nigerija",
                "common": "Nigerija"
            },
            "ita": {
                "official": "Repubblica federale di Nigeria",
                "common": "Nigeria"
            },
            "jpn": {
                "official": "ナイジェリア連邦共和国",
                "common": "ナイジェリア"
            },
            "nld": {
                "official": "Federale Republiek Nigeria",
                "common": "Nigeria"
            },
            "por": {
                "official": "República Federal da Nigéria",
                "common": "Nigéria"
            },
            "rus": {
                "official": "Федеративная Республика Нигерия",
                "common": "Нигерия"
            },
            "spa": {
                "official": "República Federal de Nigeria",
                "common": "Nigeria"
            },
            "fin": {
                "official": "Nigerian liittotasavalta",
                "common": "Nigeria"
            }
        },
        "latlng": [
            10,
            8
        ],
        "demonym": "Nigerian",
        "landlocked": false,
        "borders": [
            "BEN",
            "CMR",
            "TCD",
            "NER"
        ],
        "area": 923768
    },
    {
        "name": {
            "common": "Nicaragua",
            "official": "Republic of Nicaragua",
            "native": {
                "spa": {
                    "official": "República de Nicaragua",
                    "common": "Nicaragua"
                }
            }
        },
        "tld": [
            ".ni"
        ],
        "cca2": "NI",
        "ccn3": "558",
        "cca3": "NIC",
        "cioc": "NCA",
        "currency": [
            "NIO"
        ],
        "callingCode": [
            "505"
        ],
        "capital": "Managua",
        "altSpellings": [
            "NI",
            "Republic of Nicaragua",
            "República de Nicaragua"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Nicaragua",
                "common": "Nicaragua"
            },
            "fra": {
                "official": "République du Nicaragua",
                "common": "Nicaragua"
            },
            "hrv": {
                "official": "Republika Nikaragva",
                "common": "Nikaragva"
            },
            "ita": {
                "official": "Repubblica del Nicaragua",
                "common": "Nicaragua"
            },
            "jpn": {
                "official": "ニカラグア共和国",
                "common": "ニカラグア"
            },
            "nld": {
                "official": "Republiek Nicaragua",
                "common": "Nicaragua"
            },
            "por": {
                "official": "República da Nicarágua",
                "common": "Nicarágua"
            },
            "rus": {
                "official": "Республика Никарагуа",
                "common": "Никарагуа"
            },
            "spa": {
                "official": "República de Nicaragua",
                "common": "Nicaragua"
            },
            "fin": {
                "official": "Nicaraguan tasavalta",
                "common": "Nicaragua"
            }
        },
        "latlng": [
            13,
            -85
        ],
        "demonym": "Nicaraguan",
        "landlocked": false,
        "borders": [
            "CRI",
            "HND"
        ],
        "area": 130373
    },
    {
        "name": {
            "common": "Niue",
            "official": "Niue",
            "native": {
                "eng": {
                    "official": "Niue",
                    "common": "Niue"
                },
                "niu": {
                    "official": "Niuē",
                    "common": "Niuē"
                }
            }
        },
        "tld": [
            ".nu"
        ],
        "cca2": "NU",
        "ccn3": "570",
        "cca3": "NIU",
        "cioc": "",
        "currency": [
            "NZD"
        ],
        "callingCode": [
            "683"
        ],
        "capital": "Alofi",
        "altSpellings": [
            "NU"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "niu": "Niuean"
        },
        "translations": {
            "deu": {
                "official": "Niue",
                "common": "Niue"
            },
            "fra": {
                "official": "Niue",
                "common": "Niue"
            },
            "hrv": {
                "official": "Niue",
                "common": "Niue"
            },
            "ita": {
                "official": "Niue",
                "common": "Niue"
            },
            "jpn": {
                "official": "ニウエ",
                "common": "ニウエ"
            },
            "nld": {
                "official": "Niue",
                "common": "Niue"
            },
            "por": {
                "official": "Niue",
                "common": "Niue"
            },
            "rus": {
                "official": "Ниуэ",
                "common": "Ниуэ"
            },
            "spa": {
                "official": "Niue",
                "common": "Niue"
            },
            "fin": {
                "official": "Niue",
                "common": "Niue"
            }
        },
        "latlng": [
            -19.03333333,
            -169.86666666
        ],
        "demonym": "Niuean",
        "landlocked": false,
        "borders": [],
        "area": 260
    },
    {
        "name": {
            "common": "Netherlands",
            "official": "Netherlands",
            "native": {
                "nld": {
                    "official": "Nederland",
                    "common": "Nederland"
                }
            }
        },
        "tld": [
            ".nl"
        ],
        "cca2": "NL",
        "ccn3": "528",
        "cca3": "NLD",
        "cioc": "NED",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "31"
        ],
        "capital": "Amsterdam",
        "altSpellings": [
            "NL",
            "Holland",
            "Nederland"
        ],
        "region": "Europe",
        "subregion": "Western Europe",
        "languages": {
            "nld": "Dutch"
        },
        "translations": {
            "deu": {
                "official": "Niederlande",
                "common": "Niederlande"
            },
            "fra": {
                "official": "Pays-Bas",
                "common": "Pays-Bas"
            },
            "hrv": {
                "official": "Holandija",
                "common": "Nizozemska"
            },
            "ita": {
                "official": "Paesi Bassi",
                "common": "Paesi Bassi"
            },
            "jpn": {
                "official": "オランダ",
                "common": "オランダ"
            },
            "nld": {
                "official": "Nederland",
                "common": "Nederland"
            },
            "por": {
                "official": "Holanda",
                "common": "Holanda"
            },
            "rus": {
                "official": "Нидерланды",
                "common": "Нидерланды"
            },
            "spa": {
                "official": "Países Bajos",
                "common": "Países Bajos"
            },
            "fin": {
                "official": "Alankomaat",
                "common": "Alankomaat"
            }
        },
        "latlng": [
            52.5,
            5.75
        ],
        "demonym": "Dutch",
        "landlocked": false,
        "borders": [
            "BEL",
            "DEU"
        ],
        "area": 41850
    },
    {
        "name": {
            "common": "Norway",
            "official": "Kingdom of Norway",
            "native": {
                "nno": {
                    "official": "Kongeriket Noreg",
                    "common": "Noreg"
                },
                "nob": {
                    "official": "Kongeriket Norge",
                    "common": "Norge"
                },
                "smi": {
                    "official": "Norgga gonagasriika",
                    "common": "Norgga"
                }
            }
        },
        "tld": [
            ".no"
        ],
        "cca2": "NO",
        "ccn3": "578",
        "cca3": "NOR",
        "cioc": "NOR",
        "currency": [
            "NOK"
        ],
        "callingCode": [
            "47"
        ],
        "capital": "Oslo",
        "altSpellings": [
            "NO",
            "Norge",
            "Noreg",
            "Kingdom of Norway",
            "Kongeriket Norge",
            "Kongeriket Noreg"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "translations": {
            "deu": {
                "official": "Königreich Norwegen",
                "common": "Norwegen"
            },
            "fra": {
                "official": "Royaume de Norvège",
                "common": "Norvège"
            },
            "hrv": {
                "official": "Kraljevina Norveška",
                "common": "Norveška"
            },
            "ita": {
                "official": "Regno di Norvegia",
                "common": "Norvegia"
            },
            "jpn": {
                "official": "ノルウェー王国",
                "common": "ノルウェー"
            },
            "nld": {
                "official": "Koninkrijk Noorwegen",
                "common": "Noorwegen"
            },
            "por": {
                "official": "Reino da Noruega",
                "common": "Noruega"
            },
            "rus": {
                "official": "Королевство Норвегия",
                "common": "Норвегия"
            },
            "spa": {
                "official": "Reino de Noruega",
                "common": "Noruega"
            },
            "fin": {
                "official": "Norjan kuningaskunta",
                "common": "Norja"
            }
        },
        "latlng": [
            62,
            10
        ],
        "demonym": "Norwegian",
        "landlocked": false,
        "borders": [
            "FIN",
            "SWE",
            "RUS"
        ],
        "area": 323802
    },
    {
        "name": {
            "common": "Nepal",
            "official": "Federal Democratic Republic of Nepal",
            "native": {
                "nep": {
                    "official": "नेपाल संघीय लोकतान्त्रिक गणतन्त्र",
                    "common": "नपल"
                }
            }
        },
        "tld": [
            ".np"
        ],
        "cca2": "NP",
        "ccn3": "524",
        "cca3": "NPL",
        "cioc": "NEP",
        "currency": [
            "NPR"
        ],
        "callingCode": [
            "977"
        ],
        "capital": "Kathmandu",
        "altSpellings": [
            "NP",
            "Federal Democratic Republic of Nepal",
            "Loktāntrik Ganatantra Nepāl"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "nep": "Nepali"
        },
        "translations": {
            "deu": {
                "official": "Demokratischen Bundesrepublik Nepal",
                "common": "Népal"
            },
            "fra": {
                "official": "République du Népal",
                "common": "Népal"
            },
            "hrv": {
                "official": "Savezna Demokratska Republika Nepal",
                "common": "Nepal"
            },
            "ita": {
                "official": "Repubblica federale democratica del Nepal",
                "common": "Nepal"
            },
            "jpn": {
                "official": "ネパール連邦民主共和国",
                "common": "ネパール"
            },
            "nld": {
                "official": "Federale Democratische Republiek Nepal",
                "common": "Nepal"
            },
            "por": {
                "official": "República Democrática Federal do Nepal",
                "common": "Nepal"
            },
            "rus": {
                "official": "Федеративная Демократическая Республика Непал",
                "common": "Непал"
            },
            "spa": {
                "official": "República Democrática Federal de Nepal",
                "common": "Nepal"
            },
            "fin": {
                "official": "Nepalin demokraattinen liittotasavalta",
                "common": "Nepal"
            }
        },
        "latlng": [
            28,
            84
        ],
        "demonym": "Nepalese",
        "landlocked": true,
        "borders": [
            "CHN",
            "IND"
        ],
        "area": 147181
    },
    {
        "name": {
            "common": "Nauru",
            "official": "Republic of Nauru",
            "native": {
                "eng": {
                    "official": "Republic of Nauru",
                    "common": "Nauru"
                },
                "nau": {
                    "official": "Republic of Nauru",
                    "common": "Nauru"
                }
            }
        },
        "tld": [
            ".nr"
        ],
        "cca2": "NR",
        "ccn3": "520",
        "cca3": "NRU",
        "cioc": "NRU",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "674"
        ],
        "capital": "Yaren",
        "altSpellings": [
            "NR",
            "Naoero",
            "Pleasant Island",
            "Republic of Nauru",
            "Ripublik Naoero"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "eng": "English",
            "nau": "Nauru"
        },
        "translations": {
            "deu": {
                "official": "Republik Nauru",
                "common": "Nauru"
            },
            "fra": {
                "official": "République de Nauru",
                "common": "Nauru"
            },
            "hrv": {
                "official": "Republika Nauru",
                "common": "Nauru"
            },
            "ita": {
                "official": "Repubblica di Nauru",
                "common": "Nauru"
            },
            "jpn": {
                "official": "ナウル共和国",
                "common": "ナウル"
            },
            "nld": {
                "official": "Republiek Nauru",
                "common": "Nauru"
            },
            "por": {
                "official": "República de Nauru",
                "common": "Nauru"
            },
            "rus": {
                "official": "Республика Науру",
                "common": "Науру"
            },
            "spa": {
                "official": "República de Nauru",
                "common": "Nauru"
            },
            "fin": {
                "official": "Naurun tasavalta",
                "common": "Nauru"
            }
        },
        "latlng": [
            -0.53333333,
            166.91666666
        ],
        "demonym": "Nauruan",
        "landlocked": false,
        "borders": [],
        "area": 21
    },
    {
        "name": {
            "common": "New Zealand",
            "official": "New Zealand",
            "native": {
                "eng": {
                    "official": "New Zealand",
                    "common": "New Zealand"
                },
                "mri": {
                    "official": "Aotearoa",
                    "common": "Aotearoa"
                },
                "nzs": {
                    "official": "New Zealand",
                    "common": "New Zealand"
                }
            }
        },
        "tld": [
            ".nz"
        ],
        "cca2": "NZ",
        "ccn3": "554",
        "cca3": "NZL",
        "cioc": "NZL",
        "currency": [
            "NZD"
        ],
        "callingCode": [
            "64"
        ],
        "capital": "Wellington",
        "altSpellings": [
            "NZ",
            "Aotearoa"
        ],
        "region": "Oceania",
        "subregion": "Australia and New Zealand",
        "languages": {
            "eng": "English",
            "mri": "Māori",
            "nzs": "New Zealand Sign Language"
        },
        "translations": {
            "deu": {
                "official": "Neuseeland",
                "common": "Neuseeland"
            },
            "fra": {
                "official": "Nouvelle-Zélande",
                "common": "Nouvelle-Zélande"
            },
            "hrv": {
                "official": "Novi Zeland",
                "common": "Novi Zeland"
            },
            "ita": {
                "official": "Nuova Zelanda",
                "common": "Nuova Zelanda"
            },
            "jpn": {
                "official": "ニュージーランド",
                "common": "ニュージーランド"
            },
            "nld": {
                "official": "Nieuw Zeeland",
                "common": "Nieuw-Zeeland"
            },
            "por": {
                "official": "nova Zelândia",
                "common": "Nova Zelândia"
            },
            "rus": {
                "official": "Новая Зеландия",
                "common": "Новая Зеландия"
            },
            "spa": {
                "official": "nueva Zelanda",
                "common": "Nueva Zelanda"
            },
            "fin": {
                "official": "Uusi-Seelanti",
                "common": "Uusi-Seelanti"
            }
        },
        "latlng": [
            -41,
            174
        ],
        "demonym": "New Zealander",
        "landlocked": false,
        "borders": [],
        "area": 270467
    },
    {
        "name": {
            "common": "Oman",
            "official": "Sultanate of Oman",
            "native": {
                "ara": {
                    "official": "سلطنة عمان",
                    "common": "عمان"
                }
            }
        },
        "tld": [
            ".om"
        ],
        "cca2": "OM",
        "ccn3": "512",
        "cca3": "OMN",
        "cioc": "OMA",
        "currency": [
            "OMR"
        ],
        "callingCode": [
            "968"
        ],
        "capital": "Muscat",
        "altSpellings": [
            "OM",
            "Sultanate of Oman",
            "Salṭanat ʻUmān"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Sultanat Oman",
                "common": "Oman"
            },
            "fra": {
                "official": "Sultanat d'Oman",
                "common": "Oman"
            },
            "hrv": {
                "official": "Sultanat Oman",
                "common": "Oman"
            },
            "ita": {
                "official": "Sultanato dell'Oman",
                "common": "oman"
            },
            "jpn": {
                "official": "オマーン·スルタン国",
                "common": "オマーン"
            },
            "nld": {
                "official": "Sultanaat van Oman",
                "common": "Oman"
            },
            "por": {
                "official": "Sultanato de Omã",
                "common": "Omã"
            },
            "rus": {
                "official": "Султанат Оман",
                "common": "Оман"
            },
            "spa": {
                "official": "Sultanato de Omán",
                "common": "Omán"
            },
            "fin": {
                "official": "Omanin sulttaanikunta",
                "common": "Oman"
            }
        },
        "latlng": [
            21,
            57
        ],
        "demonym": "Omani",
        "landlocked": false,
        "borders": [
            "SAU",
            "ARE",
            "YEM"
        ],
        "area": 309500
    },
    {
        "name": {
            "common": "Pakistan",
            "official": "Islamic Republic of Pakistan",
            "native": {
                "eng": {
                    "official": "Islamic Republic of Pakistan",
                    "common": "Pakistan"
                },
                "urd": {
                    "official": "اسلامی جمہوریۂ پاكستان",
                    "common": "پاكستان"
                }
            }
        },
        "tld": [
            ".pk"
        ],
        "cca2": "PK",
        "ccn3": "586",
        "cca3": "PAK",
        "cioc": "PAK",
        "currency": [
            "PKR"
        ],
        "callingCode": [
            "92"
        ],
        "capital": "Islamabad",
        "altSpellings": [
            "PK",
            "Pākistān",
            "Islamic Republic of Pakistan",
            "Islāmī Jumhūriya'eh Pākistān"
        ],
        "region": "Asia",
        "subregion": "Southern Asia",
        "languages": {
            "eng": "English",
            "urd": "Urdu"
        },
        "translations": {
            "deu": {
                "official": "Islamische Republik Pakistan",
                "common": "Pakistan"
            },
            "fra": {
                "official": "République islamique du Pakistan",
                "common": "Pakistan"
            },
            "hrv": {
                "official": "Islamska Republika Pakistan",
                "common": "Pakistan"
            },
            "ita": {
                "official": "Repubblica islamica del Pakistan",
                "common": "Pakistan"
            },
            "jpn": {
                "official": "パキスタン",
                "common": "パキスタン"
            },
            "nld": {
                "official": "Islamitische Republiek Pakistan",
                "common": "Pakistan"
            },
            "por": {
                "official": "República Islâmica do Paquistão",
                "common": "Paquistão"
            },
            "rus": {
                "official": "Исламская Республика Пакистан",
                "common": "Пакистан"
            },
            "spa": {
                "official": "República Islámica de Pakistán",
                "common": "Pakistán"
            },
            "fin": {
                "official": "Pakistanin islamilainen tasavalta",
                "common": "Pakistan"
            }
        },
        "latlng": [
            30,
            70
        ],
        "demonym": "Pakistani",
        "landlocked": false,
        "borders": [
            "AFG",
            "CHN",
            "IND",
            "IRN"
        ],
        "area": 881912
    },
    {
        "name": {
            "common": "Panama",
            "official": "Republic of Panama",
            "native": {
                "spa": {
                    "official": "República de Panamá",
                    "common": "Panamá"
                }
            }
        },
        "tld": [
            ".pa"
        ],
        "cca2": "PA",
        "ccn3": "591",
        "cca3": "PAN",
        "cioc": "PAN",
        "currency": [
            "PAB",
            "USD"
        ],
        "callingCode": [
            "507"
        ],
        "capital": "Panama City",
        "altSpellings": [
            "PA",
            "Republic of Panama",
            "República de Panamá"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Panama",
                "common": "Panama"
            },
            "fra": {
                "official": "République du Panama",
                "common": "Panama"
            },
            "hrv": {
                "official": "Republika Panama",
                "common": "Panama"
            },
            "ita": {
                "official": "Repubblica di Panama",
                "common": "Panama"
            },
            "jpn": {
                "official": "パナマ共和国",
                "common": "パナマ"
            },
            "nld": {
                "official": "Republiek Panama",
                "common": "Panama"
            },
            "por": {
                "official": "República do Panamá",
                "common": "Panamá"
            },
            "rus": {
                "official": "Республика Панама",
                "common": "Панама"
            },
            "spa": {
                "official": "República de Panamá",
                "common": "Panamá"
            },
            "fin": {
                "official": "Panaman tasavalta",
                "common": "Panama"
            }
        },
        "latlng": [
            9,
            -80
        ],
        "demonym": "Panamanian",
        "landlocked": false,
        "borders": [
            "COL",
            "CRI"
        ],
        "area": 75417
    },
    {
        "name": {
            "common": "Pitcairn Islands",
            "official": "Pitcairn Group of Islands",
            "native": {
                "eng": {
                    "official": "Pitcairn Group of Islands",
                    "common": "Pitcairn Islands"
                }
            }
        },
        "tld": [
            ".pn"
        ],
        "cca2": "PN",
        "ccn3": "612",
        "cca3": "PCN",
        "cioc": "",
        "currency": [
            "NZD"
        ],
        "callingCode": [
            "64"
        ],
        "capital": "Adamstown",
        "altSpellings": [
            "PN",
            "Pitcairn",
            "Pitcairn Henderson Ducie and Oeno Islands"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Pitcairn Inselgruppe",
                "common": "Pitcairn"
            },
            "fra": {
                "official": "Groupe d'îles Pitcairn",
                "common": "Îles Pitcairn"
            },
            "hrv": {
                "official": "Pitcairn skupine otoka",
                "common": "Pitcairnovo otočje"
            },
            "ita": {
                "official": "Pitcairn gruppo di isole",
                "common": "Isole Pitcairn"
            },
            "jpn": {
                "official": "島のピトケアングループ",
                "common": "ピトケアン"
            },
            "nld": {
                "official": "Pitcairn groep eilanden",
                "common": "Pitcairneilanden"
            },
            "por": {
                "official": "Pitcairn grupo de ilhas",
                "common": "Ilhas Pitcairn"
            },
            "rus": {
                "official": "Питкэрн группа островов",
                "common": "Острова Питкэрн"
            },
            "spa": {
                "official": "Grupo de Islas Pitcairn",
                "common": "Islas Pitcairn"
            },
            "fin": {
                "official": "Pitcairn",
                "common": "Pitcairn"
            }
        },
        "latlng": [
            -25.06666666,
            -130.1
        ],
        "demonym": "Pitcairn Islander",
        "landlocked": false,
        "borders": [],
        "area": 47
    },
    {
        "name": {
            "common": "Peru",
            "official": "Republic of Peru",
            "native": {
                "aym": {
                    "official": "Piruw Suyu",
                    "common": "Piruw"
                },
                "que": {
                    "official": "Piruw Ripuwlika",
                    "common": "Piruw"
                },
                "spa": {
                    "official": "República del Perú",
                    "common": "Perú"
                }
            }
        },
        "tld": [
            ".pe"
        ],
        "cca2": "PE",
        "ccn3": "604",
        "cca3": "PER",
        "cioc": "PER",
        "currency": [
            "PEN"
        ],
        "callingCode": [
            "51"
        ],
        "capital": "Lima",
        "altSpellings": [
            "PE",
            "Republic of Peru",
            "República del Perú"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "aym": "Aymara",
            "que": "Quechua",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Peru",
                "common": "Peru"
            },
            "fra": {
                "official": "République du Pérou",
                "common": "Pérou"
            },
            "hrv": {
                "official": "Republika Peru",
                "common": "Peru"
            },
            "ita": {
                "official": "Repubblica del Perù",
                "common": "Perù"
            },
            "jpn": {
                "official": "ペルー共和国",
                "common": "ペルー"
            },
            "nld": {
                "official": "Republiek Peru",
                "common": "Peru"
            },
            "por": {
                "official": "República do Peru",
                "common": "Perú"
            },
            "rus": {
                "official": "Республика Перу",
                "common": "Перу"
            },
            "spa": {
                "official": "República de Perú",
                "common": "Perú"
            },
            "fin": {
                "official": "Perun tasavalta",
                "common": "Peru"
            }
        },
        "latlng": [
            -10,
            -76
        ],
        "demonym": "Peruvian",
        "landlocked": false,
        "borders": [
            "BOL",
            "BRA",
            "CHL",
            "COL",
            "ECU"
        ],
        "area": 1285216
    },
    {
        "name": {
            "common": "Philippines",
            "official": "Republic of the Philippines",
            "native": {
                "eng": {
                    "official": "Republic of the Philippines",
                    "common": "Philippines"
                },
                "fil": {
                    "official": "Republic of the Philippines",
                    "common": "Pilipinas"
                }
            }
        },
        "tld": [
            ".ph"
        ],
        "cca2": "PH",
        "ccn3": "608",
        "cca3": "PHL",
        "cioc": "PHI",
        "currency": [
            "PHP"
        ],
        "callingCode": [
            "63"
        ],
        "capital": "Manila",
        "altSpellings": [
            "PH",
            "Republic of the Philippines",
            "Repúblika ng Pilipinas"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "eng": "English",
            "fil": "Filipino"
        },
        "translations": {
            "deu": {
                "official": "Republik der Philippinen",
                "common": "Philippinen"
            },
            "fra": {
                "official": "République des Philippines",
                "common": "Philippines"
            },
            "hrv": {
                "official": "Republika Filipini",
                "common": "Filipini"
            },
            "ita": {
                "official": "Repubblica delle Filippine",
                "common": "Filippine"
            },
            "jpn": {
                "official": "フィリピン共和国",
                "common": "フィリピン"
            },
            "nld": {
                "official": "Republiek der Filipijnen",
                "common": "Filipijnen"
            },
            "por": {
                "official": "República das Filipinas",
                "common": "Filipinas"
            },
            "rus": {
                "official": "Республика Филиппины",
                "common": "Филиппины"
            },
            "spa": {
                "official": "República de las Filipinas",
                "common": "Filipinas"
            },
            "fin": {
                "official": "Filippiinien tasavalta",
                "common": "Filippiinit"
            }
        },
        "latlng": [
            13,
            122
        ],
        "demonym": "Filipino",
        "landlocked": false,
        "borders": [],
        "area": 342353
    },
    {
        "name": {
            "common": "Palau",
            "official": "Republic of Palau",
            "native": {
                "eng": {
                    "official": "Republic of Palau",
                    "common": "Palau"
                },
                "pau": {
                    "official": "Beluu er a Belau",
                    "common": "Belau"
                }
            }
        },
        "tld": [
            ".pw"
        ],
        "cca2": "PW",
        "ccn3": "585",
        "cca3": "PLW",
        "cioc": "PLW",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "680"
        ],
        "capital": "Ngerulmud",
        "altSpellings": [
            "PW",
            "Republic of Palau",
            "Beluu er a Belau"
        ],
        "region": "Oceania",
        "subregion": "Micronesia",
        "languages": {
            "eng": "English",
            "pau": "Palauan"
        },
        "translations": {
            "deu": {
                "official": "Palau",
                "common": "Palau"
            },
            "fra": {
                "official": "République des Palaos (Palau)",
                "common": "Palaos (Palau)"
            },
            "hrv": {
                "official": "Republika Palau",
                "common": "Palau"
            },
            "ita": {
                "official": "Repubblica di Palau",
                "common": "Palau"
            },
            "jpn": {
                "official": "パラオ共和国",
                "common": "パラオ"
            },
            "nld": {
                "official": "Republiek van Palau",
                "common": "Palau"
            },
            "por": {
                "official": "República de Palau",
                "common": "Palau"
            },
            "rus": {
                "official": "Республика Палау",
                "common": "Палау"
            },
            "spa": {
                "official": "República de Palau",
                "common": "Palau"
            },
            "fin": {
                "official": "Palaun tasavalta",
                "common": "Palau"
            }
        },
        "latlng": [
            7.5,
            134.5
        ],
        "demonym": "Palauan",
        "landlocked": false,
        "borders": [],
        "area": 459
    },
    {
        "name": {
            "common": "Papua New Guinea",
            "official": "Independent State of Papua New Guinea",
            "native": {
                "eng": {
                    "official": "Independent State of Papua New Guinea",
                    "common": "Papua New Guinea"
                },
                "hmo": {
                    "official": "Independen Stet bilong Papua Niugini",
                    "common": "Papua Niu Gini"
                },
                "tpi": {
                    "official": "Independen Stet bilong Papua Niugini",
                    "common": "Papua Niugini"
                }
            }
        },
        "tld": [
            ".pg"
        ],
        "cca2": "PG",
        "ccn3": "598",
        "cca3": "PNG",
        "cioc": "PNG",
        "currency": [
            "PGK"
        ],
        "callingCode": [
            "675"
        ],
        "capital": "Port Moresby",
        "altSpellings": [
            "PG",
            "Independent State of Papua New Guinea",
            "Independen Stet bilong Papua Niugini"
        ],
        "region": "Oceania",
        "subregion": "Melanesia",
        "languages": {
            "eng": "English",
            "hmo": "Hiri Motu",
            "tpi": "Tok Pisin"
        },
        "translations": {
            "deu": {
                "official": "Unabhängige Staat Papua-Neuguinea",
                "common": "Papua-Neuguinea"
            },
            "fra": {
                "official": "État indépendant de Papouasie-Nouvelle-Guinée",
                "common": "Papouasie-Nouvelle-Guinée"
            },
            "hrv": {
                "official": "Nezavisna Država Papui Novoj Gvineji",
                "common": "Papua Nova Gvineja"
            },
            "ita": {
                "official": "Stato indipendente di Papua Nuova Guinea",
                "common": "Papua Nuova Guinea"
            },
            "jpn": {
                "official": "パプアニューギニア独立国",
                "common": "パプアニューギニア"
            },
            "nld": {
                "official": "Onafhankelijke Staat Papoea -Nieuw-Guinea",
                "common": "Papoea-Nieuw-Guinea"
            },
            "por": {
                "official": "Estado Independente da Papua Nova Guiné",
                "common": "Papua Nova Guiné"
            },
            "rus": {
                "official": "Независимое Государство Папуа-Новой Гвинеи",
                "common": "Папуа — Новая Гвинея"
            },
            "spa": {
                "official": "Estado Independiente de Papúa Nueva Guinea",
                "common": "Papúa Nueva Guinea"
            },
            "fin": {
                "official": "Papua-Uuden-Guinean Itsenäinen valtio",
                "common": "Papua-Uusi-Guinea"
            }
        },
        "latlng": [
            -6,
            147
        ],
        "demonym": "Papua New Guinean",
        "landlocked": false,
        "borders": [
            "IDN"
        ],
        "area": 462840
    },
    {
        "name": {
            "common": "Poland",
            "official": "Republic of Poland",
            "native": {
                "pol": {
                    "official": "Rzeczpospolita Polska",
                    "common": "Polska"
                }
            }
        },
        "tld": [
            ".pl"
        ],
        "cca2": "PL",
        "ccn3": "616",
        "cca3": "POL",
        "cioc": "POL",
        "currency": [
            "PLN"
        ],
        "callingCode": [
            "48"
        ],
        "capital": "Warsaw",
        "altSpellings": [
            "PL",
            "Republic of Poland",
            "Rzeczpospolita Polska"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "pol": "Polish"
        },
        "translations": {
            "deu": {
                "official": "Republik Polen",
                "common": "Polen"
            },
            "fra": {
                "official": "République de Pologne",
                "common": "Pologne"
            },
            "hrv": {
                "official": "Republika Poljska",
                "common": "Poljska"
            },
            "ita": {
                "official": "Repubblica di Polonia",
                "common": "Polonia"
            },
            "jpn": {
                "official": "ポーランド共和国",
                "common": "ポーランド"
            },
            "nld": {
                "official": "Republiek Polen",
                "common": "Polen"
            },
            "por": {
                "official": "República da Polónia",
                "common": "Polónia"
            },
            "rus": {
                "official": "Республика Польша",
                "common": "Польша"
            },
            "spa": {
                "official": "República de Polonia",
                "common": "Polonia"
            },
            "fin": {
                "official": "Puolan tasavalta",
                "common": "Puola"
            }
        },
        "latlng": [
            52,
            20
        ],
        "demonym": "Polish",
        "landlocked": false,
        "borders": [
            "BLR",
            "CZE",
            "DEU",
            "LTU",
            "RUS",
            "SVK",
            "UKR"
        ],
        "area": 312679
    },
    {
        "name": {
            "common": "Puerto Rico",
            "official": "Commonwealth of Puerto Rico",
            "native": {
                "eng": {
                    "official": "Commonwealth of Puerto Rico",
                    "common": "Puerto Rico"
                },
                "spa": {
                    "official": "Estado Libre Asociado de Puerto Rico",
                    "common": "Puerto Rico"
                }
            }
        },
        "tld": [
            ".pr"
        ],
        "cca2": "PR",
        "ccn3": "630",
        "cca3": "PRI",
        "cioc": "PUR",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1787",
            "1939"
        ],
        "capital": "San Juan",
        "altSpellings": [
            "PR",
            "Commonwealth of Puerto Rico",
            "Estado Libre Asociado de Puerto Rico"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Commonwealth von Puerto Rico",
                "common": "Puerto Rico"
            },
            "fra": {
                "official": "Porto Rico",
                "common": "Porto Rico"
            },
            "hrv": {
                "official": "Zajednica Puerto Rico",
                "common": "Portoriko"
            },
            "ita": {
                "official": "Commonwealth di Porto Rico",
                "common": "Porto Rico"
            },
            "jpn": {
                "official": "プエルトリコのコモンウェルス",
                "common": "プエルトリコ"
            },
            "nld": {
                "official": "Gemenebest van Puerto Rico",
                "common": "Puerto Rico"
            },
            "por": {
                "official": "Commonwealth of Puerto Rico",
                "common": "Porto Rico"
            },
            "rus": {
                "official": "Содружество Пуэрто-Рико",
                "common": "Пуэрто-Рико"
            },
            "spa": {
                "official": "Asociado de Puerto Rico",
                "common": "Puerto Rico"
            },
            "fin": {
                "official": "Puerto Rico",
                "common": "Puerto Rico"
            }
        },
        "latlng": [
            18.25,
            -66.5
        ],
        "demonym": "Puerto Rican",
        "landlocked": false,
        "borders": [],
        "area": 8870
    },
    {
        "name": {
            "common": "North Korea",
            "official": "Democratic People's Republic of Korea",
            "native": {
                "kor": {
                    "official": "조선 민주주의 인민 공화국",
                    "common": "북한"
                }
            }
        },
        "tld": [
            ".kp"
        ],
        "cca2": "KP",
        "ccn3": "408",
        "cca3": "PRK",
        "cioc": "PRK",
        "currency": [
            "KPW"
        ],
        "callingCode": [
            "850"
        ],
        "capital": "Pyongyang",
        "altSpellings": [
            "KP",
            "Democratic People's Republic of Korea",
            "조선민주주의인민공화국",
            "Chosŏn Minjujuŭi Inmin Konghwaguk",
            "Korea, Democratic People's Republic of"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "kor": "Korean"
        },
        "translations": {
            "deu": {
                "official": "Demokratische Volksrepublik Korea",
                "common": "Nordkorea"
            },
            "fra": {
                "official": "République populaire démocratique de Corée",
                "common": "Corée du Nord"
            },
            "hrv": {
                "official": "Demokratska Narodna Republika Koreja",
                "common": "Sjeverna Koreja"
            },
            "ita": {
                "official": "Repubblica democratica popolare di Corea",
                "common": "Corea del Nord"
            },
            "jpn": {
                "official": "朝鮮民主主義人民共和国",
                "common": "朝鮮民主主義人民共和国"
            },
            "nld": {
                "official": "Democratische Volksrepubliek Korea",
                "common": "Noord-Korea"
            },
            "por": {
                "official": "República Popular Democrática da Coreia",
                "common": "Coreia do Norte"
            },
            "rus": {
                "official": "Корейская Народно-Демократическая Республика Корея",
                "common": "Северная Корея"
            },
            "spa": {
                "official": "República Popular Democrática de Corea",
                "common": "Corea del Norte"
            },
            "fin": {
                "official": "Korean demokraattinen kansantasavalta",
                "common": "Pohjois-Korea"
            }
        },
        "latlng": [
            40,
            127
        ],
        "demonym": "North Korean",
        "landlocked": false,
        "borders": [
            "CHN",
            "KOR",
            "RUS"
        ],
        "area": 120538
    },
    {
        "name": {
            "common": "Portugal",
            "official": "Portuguese Republic",
            "native": {
                "por": {
                    "official": "República português",
                    "common": "Portugal"
                }
            }
        },
        "tld": [
            ".pt"
        ],
        "cca2": "PT",
        "ccn3": "620",
        "cca3": "PRT",
        "cioc": "POR",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "351"
        ],
        "capital": "Lisbon",
        "altSpellings": [
            "PT",
            "Portuguesa",
            "Portuguese Republic",
            "República Portuguesa"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "deu": {
                "official": "Portugiesische Republik",
                "common": "Portugal"
            },
            "fra": {
                "official": "République portugaise",
                "common": "Portugal"
            },
            "hrv": {
                "official": "Portugalska Republika",
                "common": "Portugal"
            },
            "ita": {
                "official": "Repubblica portoghese",
                "common": "Portogallo"
            },
            "jpn": {
                "official": "ポルトガル共和国",
                "common": "ポルトガル"
            },
            "nld": {
                "official": "Portugese Republiek",
                "common": "Portugal"
            },
            "por": {
                "official": "República português",
                "common": "Portugal"
            },
            "rus": {
                "official": "Португальская Республика",
                "common": "Португалия"
            },
            "spa": {
                "official": "República Portuguesa",
                "common": "Portugal"
            },
            "fin": {
                "official": "Portugalin tasavalta",
                "common": "Portugali"
            }
        },
        "latlng": [
            39.5,
            -8
        ],
        "demonym": "Portuguese",
        "landlocked": false,
        "borders": [
            "ESP"
        ],
        "area": 92090
    },
    {
        "name": {
            "common": "Paraguay",
            "official": "Republic of Paraguay",
            "native": {
                "grn": {
                    "official": "Tetã Paraguái",
                    "common": "Paraguái"
                },
                "spa": {
                    "official": "República de Paraguay",
                    "common": "Paraguay"
                }
            }
        },
        "tld": [
            ".py"
        ],
        "cca2": "PY",
        "ccn3": "600",
        "cca3": "PRY",
        "cioc": "PAR",
        "currency": [
            "PYG"
        ],
        "callingCode": [
            "595"
        ],
        "capital": "Asunción",
        "altSpellings": [
            "PY",
            "Republic of Paraguay",
            "República del Paraguay",
            "Tetã Paraguái"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "grn": "Guaraní",
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Paraguay",
                "common": "Paraguay"
            },
            "fra": {
                "official": "République du Paraguay",
                "common": "Paraguay"
            },
            "hrv": {
                "official": "Republika Paragvaj",
                "common": "Paragvaj"
            },
            "ita": {
                "official": "Repubblica del Paraguay",
                "common": "Paraguay"
            },
            "jpn": {
                "official": "パラグアイ共和国",
                "common": "パラグアイ"
            },
            "nld": {
                "official": "Republiek Paraguay",
                "common": "Paraguay"
            },
            "por": {
                "official": "República do Paraguai",
                "common": "Paraguai"
            },
            "rus": {
                "official": "Республика Парагвай",
                "common": "Парагвай"
            },
            "spa": {
                "official": "República de Paraguay",
                "common": "Paraguay"
            },
            "fin": {
                "official": "Paraguayn tasavalta",
                "common": "Paraguay"
            }
        },
        "latlng": [
            -23,
            -58
        ],
        "demonym": "Paraguayan",
        "landlocked": true,
        "borders": [
            "ARG",
            "BOL",
            "BRA"
        ],
        "area": 406752
    },
    {
        "name": {
            "common": "Palestine",
            "official": "State of Palestine",
            "native": {
                "ara": {
                    "official": "دولة فلسطين",
                    "common": "فلسطين"
                }
            }
        },
        "tld": [
            ".ps",
            "فلسطين."
        ],
        "cca2": "PS",
        "ccn3": "275",
        "cca3": "PSE",
        "cioc": "PLE",
        "currency": [
            "ILS"
        ],
        "callingCode": [
            "970"
        ],
        "capital": "Ramallah",
        "altSpellings": [
            "PS",
            "Palestine, State of",
            "State of Palestine",
            "Dawlat Filasṭin"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Staat Palästina",
                "common": "Palästina"
            },
            "fra": {
                "official": "État de Palestine",
                "common": "Palestine"
            },
            "hrv": {
                "official": "State of Palestine",
                "common": "Palestina"
            },
            "ita": {
                "official": "Stato di Palestina",
                "common": "Palestina"
            },
            "jpn": {
                "official": "パレスチナ自治政府",
                "common": "パレスチナ"
            },
            "nld": {
                "official": "Staat Palestina",
                "common": "Palestijnse gebieden"
            },
            "por": {
                "official": "Estado da Palestina",
                "common": "Palestina"
            },
            "rus": {
                "official": "Государство Палестина",
                "common": "Палестина"
            },
            "spa": {
                "official": "Estado de Palestina",
                "common": "Palestina"
            },
            "fin": {
                "official": "Palestiinan valtio",
                "common": "Palestiina"
            }
        },
        "latlng": [
            31.9,
            35.2
        ],
        "demonym": "Palestinian",
        "landlocked": false,
        "borders": [
            "ISR",
            "EGY",
            "JOR"
        ],
        "area": 6220
    },
    {
        "name": {
            "common": "French Polynesia",
            "official": "French Polynesia",
            "native": {
                "fra": {
                    "official": "Polynésie française",
                    "common": "Polynésie française"
                }
            }
        },
        "tld": [
            ".pf"
        ],
        "cca2": "PF",
        "ccn3": "258",
        "cca3": "PYF",
        "cioc": "",
        "currency": [
            "XPF"
        ],
        "callingCode": [
            "689"
        ],
        "capital": "Papeetē",
        "altSpellings": [
            "PF",
            "Polynésie française",
            "French Polynesia",
            "Pōrīnetia Farāni"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Französisch-Polynesien",
                "common": "Französisch-Polynesien"
            },
            "fra": {
                "official": "Polynésie française",
                "common": "Polynésie française"
            },
            "hrv": {
                "official": "Francuska Polinezija",
                "common": "Francuska Polinezija"
            },
            "ita": {
                "official": "Polinesia Francese",
                "common": "Polinesia Francese"
            },
            "jpn": {
                "official": "フランス領ポリネシア",
                "common": "フランス領ポリネシア"
            },
            "nld": {
                "official": "Frans-Polynesië",
                "common": "Frans-Polynesië"
            },
            "por": {
                "official": "Polinésia Francesa",
                "common": "Polinésia Francesa"
            },
            "rus": {
                "official": "Французская Полинезия",
                "common": "Французская Полинезия"
            },
            "spa": {
                "official": "Polinesia francés",
                "common": "Polinesia Francesa"
            },
            "fin": {
                "official": "Ranskan Polynesia",
                "common": "Ranskan Polynesia"
            }
        },
        "latlng": [
            -15,
            -140
        ],
        "demonym": "French Polynesian",
        "landlocked": false,
        "borders": [],
        "area": 4167
    },
    {
        "name": {
            "common": "Qatar",
            "official": "State of Qatar",
            "native": {
                "ara": {
                    "official": "دولة قطر",
                    "common": "قطر"
                }
            }
        },
        "tld": [
            ".qa",
            "قطر."
        ],
        "cca2": "QA",
        "ccn3": "634",
        "cca3": "QAT",
        "cioc": "QAT",
        "currency": [
            "QAR"
        ],
        "callingCode": [
            "974"
        ],
        "capital": "Doha",
        "altSpellings": [
            "QA",
            "State of Qatar",
            "Dawlat Qaṭar"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Staat Katar",
                "common": "Katar"
            },
            "fra": {
                "official": "État du Qatar",
                "common": "Qatar"
            },
            "hrv": {
                "official": "Država Katar",
                "common": "Katar"
            },
            "ita": {
                "official": "Stato del Qatar",
                "common": "Qatar"
            },
            "jpn": {
                "official": "カタール国",
                "common": "カタール"
            },
            "nld": {
                "official": "Staat Qatar",
                "common": "Qatar"
            },
            "por": {
                "official": "Estado do Qatar",
                "common": "Catar"
            },
            "rus": {
                "official": "Государство Катар",
                "common": "Катар"
            },
            "spa": {
                "official": "Estado de Qatar",
                "common": "Catar"
            },
            "fin": {
                "official": "Qatarin valtio",
                "common": "Qatar"
            }
        },
        "latlng": [
            25.5,
            51.25
        ],
        "demonym": "Qatari",
        "landlocked": false,
        "borders": [
            "SAU"
        ],
        "area": 11586
    },
    {
        "name": {
            "common": "Réunion",
            "official": "Réunion Island",
            "native": {
                "fra": {
                    "official": "Ile de la Réunion",
                    "common": "La Réunion"
                }
            }
        },
        "tld": [
            ".re"
        ],
        "cca2": "RE",
        "ccn3": "638",
        "cca3": "REU",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "262"
        ],
        "capital": "Saint-Denis",
        "altSpellings": [
            "RE",
            "Reunion"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Réunion",
                "common": "Réunion"
            },
            "fra": {
                "official": "Ile de la Réunion",
                "common": "Réunion"
            },
            "hrv": {
                "official": "Réunion Island",
                "common": "Réunion"
            },
            "ita": {
                "official": "Réunion",
                "common": "Riunione"
            },
            "jpn": {
                "official": "レユニオン島",
                "common": "レユニオン"
            },
            "nld": {
                "official": "Réunion",
                "common": "Réunion"
            },
            "por": {
                "official": "Ilha da Reunião",
                "common": "Reunião"
            },
            "rus": {
                "official": "Реюньон",
                "common": "Реюньон"
            },
            "spa": {
                "official": "Isla de la Reunión",
                "common": "Reunión"
            },
            "fin": {
                "official": "Réunion",
                "common": "Réunion"
            }
        },
        "latlng": [
            -21.15,
            55.5
        ],
        "demonym": "French",
        "landlocked": false,
        "borders": [],
        "area": 2511
    },
    {
        "name": {
            "common": "Romania",
            "official": "Romania",
            "native": {
                "ron": {
                    "official": "România",
                    "common": "România"
                }
            }
        },
        "tld": [
            ".ro"
        ],
        "cca2": "RO",
        "ccn3": "642",
        "cca3": "ROU",
        "cioc": "ROU",
        "currency": [
            "RON"
        ],
        "callingCode": [
            "40"
        ],
        "capital": "Bucharest",
        "altSpellings": [
            "RO",
            "Rumania",
            "Roumania",
            "România"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "ron": "Romanian"
        },
        "translations": {
            "deu": {
                "official": "Rumänien",
                "common": "Rumänien"
            },
            "fra": {
                "official": "Roumanie",
                "common": "Roumanie"
            },
            "hrv": {
                "official": "Rumunija",
                "common": "Rumunjska"
            },
            "ita": {
                "official": "Romania",
                "common": "Romania"
            },
            "jpn": {
                "official": "ルーマニア",
                "common": "ルーマニア"
            },
            "nld": {
                "official": "Roemenië",
                "common": "Roemenië"
            },
            "por": {
                "official": "Romênia",
                "common": "Roménia"
            },
            "rus": {
                "official": "Румыния",
                "common": "Румыния"
            },
            "spa": {
                "official": "Rumania",
                "common": "Rumania"
            },
            "fin": {
                "official": "Romania",
                "common": "Romania"
            }
        },
        "latlng": [
            46,
            25
        ],
        "demonym": "Romanian",
        "landlocked": false,
        "borders": [
            "BGR",
            "HUN",
            "MDA",
            "SRB",
            "UKR"
        ],
        "area": 238391
    },
    {
        "name": {
            "common": "Russia",
            "official": "Russian Federation",
            "native": {
                "rus": {
                    "official": "Русская Федерация",
                    "common": "Россия"
                }
            }
        },
        "tld": [
            ".ru",
            ".su",
            ".рф"
        ],
        "cca2": "RU",
        "ccn3": "643",
        "cca3": "RUS",
        "cioc": "RUS",
        "currency": [
            "RUB"
        ],
        "callingCode": [
            "7"
        ],
        "capital": "Moscow",
        "altSpellings": [
            "RU",
            "Rossiya",
            "Russian Federation",
            "Российская Федерация",
            "Rossiyskaya Federatsiya"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "rus": "Russian"
        },
        "translations": {
            "deu": {
                "official": "Russische Föderation",
                "common": "Russland"
            },
            "fra": {
                "official": "Fédération de Russie",
                "common": "Russie"
            },
            "hrv": {
                "official": "Ruska Federacija",
                "common": "Rusija"
            },
            "ita": {
                "official": "Federazione russa",
                "common": "Russia"
            },
            "jpn": {
                "official": "ロシア連邦",
                "common": "ロシア連邦"
            },
            "nld": {
                "official": "Russische Federatie",
                "common": "Rusland"
            },
            "por": {
                "official": "Federação Russa",
                "common": "Rússia"
            },
            "rus": {
                "official": "Россия Федерация",
                "common": "Россия"
            },
            "spa": {
                "official": "Federación de Rusia",
                "common": "Rusia"
            },
            "fin": {
                "official": "Venäjän federaatio",
                "common": "Venäjä"
            }
        },
        "latlng": [
            60,
            100
        ],
        "demonym": "Russian",
        "landlocked": false,
        "borders": [
            "AZE",
            "BLR",
            "CHN",
            "EST",
            "FIN",
            "GEO",
            "KAZ",
            "PRK",
            "LVA",
            "LTU",
            "MNG",
            "NOR",
            "POL",
            "UKR"
        ],
        "area": 17098242
    },
    {
        "name": {
            "common": "Rwanda",
            "official": "Republic of Rwanda",
            "native": {
                "eng": {
                    "official": "Republic of Rwanda",
                    "common": "Rwanda"
                },
                "fra": {
                    "official": "République rwandaise",
                    "common": "Rwanda"
                },
                "kin": {
                    "official": "Repubulika y'u Rwanda",
                    "common": "Rwanda"
                }
            }
        },
        "tld": [
            ".rw"
        ],
        "cca2": "RW",
        "ccn3": "646",
        "cca3": "RWA",
        "cioc": "RWA",
        "currency": [
            "RWF"
        ],
        "callingCode": [
            "250"
        ],
        "capital": "Kigali",
        "altSpellings": [
            "RW",
            "Republic of Rwanda",
            "Repubulika y'u Rwanda",
            "République du Rwanda"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "fra": "French",
            "kin": "Kinyarwanda"
        },
        "translations": {
            "deu": {
                "official": "Republik Ruanda",
                "common": "Ruanda"
            },
            "fra": {
                "official": "République rwandaise",
                "common": "Rwanda"
            },
            "hrv": {
                "official": "Republika Ruandi",
                "common": "Ruanda"
            },
            "ita": {
                "official": "Repubblica del Ruanda",
                "common": "Ruanda"
            },
            "jpn": {
                "official": "ルワンダ共和国",
                "common": "ルワンダ"
            },
            "nld": {
                "official": "Republiek Rwanda",
                "common": "Rwanda"
            },
            "por": {
                "official": "República do Ruanda",
                "common": "Ruanda"
            },
            "rus": {
                "official": "Республика Руанда",
                "common": "Руанда"
            },
            "spa": {
                "official": "República de Rwanda",
                "common": "Ruanda"
            },
            "fin": {
                "official": "Ruandan tasavalta",
                "common": "Ruanda"
            }
        },
        "latlng": [
            -2,
            30
        ],
        "demonym": "Rwandan",
        "landlocked": true,
        "borders": [
            "BDI",
            "COD",
            "TZA",
            "UGA"
        ],
        "area": 26338
    },
    {
        "name": {
            "common": "Saudi Arabia",
            "official": "Kingdom of Saudi Arabia",
            "native": {
                "ara": {
                    "official": "المملكة العربية السعودية",
                    "common": "العربية السعودية"
                }
            }
        },
        "tld": [
            ".sa",
            ".السعودية"
        ],
        "cca2": "SA",
        "ccn3": "682",
        "cca3": "SAU",
        "cioc": "KSA",
        "currency": [
            "SAR"
        ],
        "callingCode": [
            "966"
        ],
        "capital": "Riyadh",
        "altSpellings": [
            "Saudi",
            "SA",
            "Kingdom of Saudi Arabia",
            "Al-Mamlakah al-‘Arabiyyah as-Su‘ūdiyyah"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Königreich Saudi-Arabien",
                "common": "Saudi-Arabien"
            },
            "fra": {
                "official": "Royaume d'Arabie Saoudite",
                "common": "Arabie Saoudite"
            },
            "hrv": {
                "official": "Kraljevina Saudijska Arabija",
                "common": "Saudijska Arabija"
            },
            "ita": {
                "official": "Arabia Saudita",
                "common": "Arabia Saudita"
            },
            "jpn": {
                "official": "サウジアラビア王国",
                "common": "サウジアラビア"
            },
            "nld": {
                "official": "Koninkrijk van Saoedi-Arabië",
                "common": "Saoedi-Arabië"
            },
            "por": {
                "official": "Reino da Arábia Saudita",
                "common": "Arábia Saudita"
            },
            "rus": {
                "official": "Королевство Саудовская Аравия",
                "common": "Саудовская Аравия"
            },
            "spa": {
                "official": "Reino de Arabia Saudita",
                "common": "Arabia Saudí"
            },
            "fin": {
                "official": "Saudi-Arabian kuningaskunta",
                "common": "Saudi-Arabia"
            }
        },
        "latlng": [
            25,
            45
        ],
        "demonym": "Saudi Arabian",
        "landlocked": false,
        "borders": [
            "IRQ",
            "JOR",
            "KWT",
            "OMN",
            "QAT",
            "ARE",
            "YEM"
        ],
        "area": 2149690
    },
    {
        "name": {
            "common": "Sudan",
            "official": "Republic of the Sudan",
            "native": {
                "ara": {
                    "official": "جمهورية السودان",
                    "common": "السودان"
                },
                "eng": {
                    "official": "Republic of the Sudan",
                    "common": "Sudan"
                }
            }
        },
        "tld": [
            ".sd"
        ],
        "cca2": "SD",
        "ccn3": "729",
        "cca3": "SDN",
        "cioc": "SUD",
        "currency": [
            "SDG"
        ],
        "callingCode": [
            "249"
        ],
        "capital": "Khartoum",
        "altSpellings": [
            "SD",
            "Republic of the Sudan",
            "Jumhūrīyat as-Sūdān"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic",
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Sudan",
                "common": "Sudan"
            },
            "fra": {
                "official": "République du Soudan",
                "common": "Soudan"
            },
            "hrv": {
                "official": "Republika Sudan",
                "common": "Sudan"
            },
            "ita": {
                "official": "Repubblica del Sudan",
                "common": "Sudan"
            },
            "jpn": {
                "official": "スーダン共和国",
                "common": "スーダン"
            },
            "nld": {
                "official": "Republiek Soedan",
                "common": "Soedan"
            },
            "por": {
                "official": "República do Sudão",
                "common": "Sudão"
            },
            "rus": {
                "official": "Республика Судан",
                "common": "Судан"
            },
            "spa": {
                "official": "República de Sudán",
                "common": "Sudán"
            },
            "fin": {
                "official": "Sudanin tasavalta",
                "common": "Sudan"
            }
        },
        "latlng": [
            15,
            30
        ],
        "demonym": "Sudanese",
        "landlocked": false,
        "borders": [
            "CAF",
            "TCD",
            "EGY",
            "ERI",
            "ETH",
            "LBY",
            "SSD"
        ],
        "area": 1886068
    },
    {
        "name": {
            "common": "Senegal",
            "official": "Republic of Senegal",
            "native": {
                "fra": {
                    "official": "République du Sénégal",
                    "common": "Sénégal"
                }
            }
        },
        "tld": [
            ".sn"
        ],
        "cca2": "SN",
        "ccn3": "686",
        "cca3": "SEN",
        "cioc": "SEN",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "221"
        ],
        "capital": "Dakar",
        "altSpellings": [
            "SN",
            "Republic of Senegal",
            "République du Sénégal"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Senegal",
                "common": "Senegal"
            },
            "fra": {
                "official": "République du Sénégal",
                "common": "Sénégal"
            },
            "hrv": {
                "official": "Republika Senegal",
                "common": "Senegal"
            },
            "ita": {
                "official": "Repubblica del Senegal",
                "common": "Senegal"
            },
            "jpn": {
                "official": "セネガル共和国",
                "common": "セネガル"
            },
            "nld": {
                "official": "Republiek Senegal",
                "common": "Senegal"
            },
            "por": {
                "official": "República do Senegal",
                "common": "Senegal"
            },
            "rus": {
                "official": "Республика Сенегал",
                "common": "Сенегал"
            },
            "spa": {
                "official": "República de Senegal",
                "common": "Senegal"
            },
            "fin": {
                "official": "Senegalin tasavalta",
                "common": "Senegal"
            }
        },
        "latlng": [
            14,
            -14
        ],
        "demonym": "Senegalese",
        "landlocked": false,
        "borders": [
            "GMB",
            "GIN",
            "GNB",
            "MLI",
            "MRT"
        ],
        "area": 196722
    },
    {
        "name": {
            "common": "Singapore",
            "official": "Republic of Singapore",
            "native": {
                "cmn": {
                    "official": "新加坡共和国",
                    "common": "新加坡"
                },
                "eng": {
                    "official": "Republic of Singapore",
                    "common": "Singapore"
                },
                "msa": {
                    "official": "Republik Singapura",
                    "common": "Singapura"
                },
                "tam": {
                    "official": "சிங்கப்பூர் குடியரசு",
                    "common": "சிங்கப்பூர்"
                }
            }
        },
        "tld": [
            ".sg",
            ".新加坡",
            ".சிங்கப்பூர்"
        ],
        "cca2": "SG",
        "ccn3": "702",
        "cca3": "SGP",
        "cioc": "SIN",
        "currency": [
            "SGD"
        ],
        "callingCode": [
            "65"
        ],
        "capital": "Singapore",
        "altSpellings": [
            "SG",
            "Singapura",
            "Republik Singapura",
            "新加坡共和国"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "cmn": "Mandarin",
            "eng": "English",
            "msa": "Malay",
            "tam": "Tamil"
        },
        "translations": {
            "deu": {
                "official": "Republik Singapur",
                "common": "Singapur"
            },
            "fra": {
                "official": "République de Singapour",
                "common": "Singapour"
            },
            "hrv": {
                "official": "Republika Singapur",
                "common": "Singapur"
            },
            "ita": {
                "official": "Repubblica di Singapore",
                "common": "Singapore"
            },
            "jpn": {
                "official": "シンガポール共和国",
                "common": "シンガポール"
            },
            "nld": {
                "official": "Republiek Singapore",
                "common": "Singapore"
            },
            "por": {
                "official": "República de Singapura",
                "common": "Singapura"
            },
            "rus": {
                "official": "Республика Сингапур",
                "common": "Сингапур"
            },
            "spa": {
                "official": "República de Singapur",
                "common": "Singapur"
            },
            "fin": {
                "official": "Singaporen tasavalta",
                "common": "Singapore"
            }
        },
        "latlng": [
            1.36666666,
            103.8
        ],
        "demonym": "Singaporean",
        "landlocked": false,
        "borders": [],
        "area": 710
    },
    {
        "name": {
            "common": "South Georgia",
            "official": "South Georgia and the South Sandwich Islands",
            "native": {
                "eng": {
                    "official": "South Georgia and the South Sandwich Islands",
                    "common": "South Georgia"
                }
            }
        },
        "tld": [
            ".gs"
        ],
        "cca2": "GS",
        "ccn3": "239",
        "cca3": "SGS",
        "cioc": "",
        "currency": [
            "GBP"
        ],
        "callingCode": [
            "500"
        ],
        "capital": "King Edward Point",
        "altSpellings": [
            "GS",
            "South Georgia and the South Sandwich Islands"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Südgeorgien und die Südlichen Sandwichinseln",
                "common": "Südgeorgien und die Südlichen Sandwichinseln"
            },
            "fra": {
                "official": "Géorgie du Sud et les îles Sandwich du Sud",
                "common": "Géorgie du Sud-et-les Îles Sandwich du Sud"
            },
            "hrv": {
                "official": "Južna Džordžija i Otoci Južni Sendvič",
                "common": "Južna Georgija i otočje Južni Sandwich"
            },
            "ita": {
                "official": "Georgia del Sud e isole Sandwich del Sud",
                "common": "Georgia del Sud e Isole Sandwich Meridionali"
            },
            "jpn": {
                "official": "サウスジョージア·サウスサンドウィッチ諸島",
                "common": "サウスジョージア・サウスサンドウィッチ諸島"
            },
            "nld": {
                "official": "Zuid-Georgië en de Zuidelijke Sandwich-eilanden",
                "common": "Zuid-Georgia en Zuidelijke Sandwicheilanden"
            },
            "por": {
                "official": "Geórgia do Sul e Sandwich do Sul",
                "common": "Ilhas Geórgia do Sul e Sandwich do Sul"
            },
            "rus": {
                "official": "Южная Георгия и Южные Сандвичевы острова",
                "common": "Южная Георгия и Южные Сандвичевы острова"
            },
            "spa": {
                "official": "Georgia del Sur y las Islas Sandwich del Sur",
                "common": "Islas Georgias del Sur y Sandwich del Sur"
            },
            "fin": {
                "official": "Etelä-Georgia ja Eteläiset Sandwichsaaret",
                "common": "Etelä-Georgia ja Eteläiset Sandwichsaaret"
            }
        },
        "latlng": [
            -54.5,
            -37
        ],
        "demonym": "South Georgian South Sandwich Islander",
        "landlocked": false,
        "borders": [],
        "area": 3903
    },
    {
        "name": {
            "common": "Svalbard and Jan Mayen",
            "official": "Svalbard og Jan Mayen",
            "native": {
                "nor": {
                    "official": "Svalbard og Jan Mayen",
                    "common": "Svalbard og Jan Mayen"
                }
            }
        },
        "tld": [
            ".sj"
        ],
        "cca2": "SJ",
        "ccn3": "744",
        "cca3": "SJM",
        "cioc": "",
        "currency": [
            "NOK"
        ],
        "callingCode": [
            "4779"
        ],
        "capital": "Longyearbyen",
        "altSpellings": [
            "SJ",
            "Svalbard and Jan Mayen Islands"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "nor": "Norwegian"
        },
        "translations": {
            "deu": {
                "official": "Inselgruppe Spitzbergen",
                "common": "Spitzbergen"
            },
            "fra": {
                "official": "Jan Mayen Svalbard",
                "common": "Svalbard et Jan Mayen"
            },
            "hrv": {
                "official": "Svalbard og Jan Mayen",
                "common": "Svalbard i Jan Mayen"
            },
            "ita": {
                "official": "Svalbard og Jan Mayen",
                "common": "Svalbard e Jan Mayen"
            },
            "jpn": {
                "official": "スバールバル諸島OGヤンマイエン",
                "common": "スヴァールバル諸島およびヤンマイエン島"
            },
            "nld": {
                "official": "Svalbard og Jan Mayen",
                "common": "Svalbard en Jan Mayen"
            },
            "por": {
                "official": "Svalbard og Jan Mayen",
                "common": "Ilhas Svalbard e Jan Mayen"
            },
            "rus": {
                "official": "Свальбарда ог Ян-Майен",
                "common": "Шпицберген и Ян-Майен"
            },
            "spa": {
                "official": "Svalbard og Jan Mayen",
                "common": "Islas Svalbard y Jan Mayen"
            },
            "fin": {
                "official": "Huippuvuoret",
                "common": "Huippuvuoret"
            }
        },
        "latlng": [
            78,
            20
        ],
        "demonym": "Norwegian",
        "landlocked": false,
        "borders": [],
        "area": -1
    },
    {
        "name": {
            "common": "Solomon Islands",
            "official": "Solomon Islands",
            "native": {
                "eng": {
                    "official": "Solomon Islands",
                    "common": "Solomon Islands"
                }
            }
        },
        "tld": [
            ".sb"
        ],
        "cca2": "SB",
        "ccn3": "090",
        "cca3": "SLB",
        "cioc": "SOL",
        "currency": [
            "SBD"
        ],
        "callingCode": [
            "677"
        ],
        "capital": "Honiara",
        "altSpellings": [
            "SB"
        ],
        "region": "Oceania",
        "subregion": "Melanesia",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Salomon-Inseln",
                "common": "Salomonen"
            },
            "fra": {
                "official": "Îles Salomon",
                "common": "Îles Salomon"
            },
            "hrv": {
                "official": "Solomonski Otoci",
                "common": "Solomonski Otoci"
            },
            "ita": {
                "official": "Isole Salomone",
                "common": "Isole Salomone"
            },
            "jpn": {
                "official": "ソロモン諸島",
                "common": "ソロモン諸島"
            },
            "nld": {
                "official": "Solomon eilanden",
                "common": "Salomonseilanden"
            },
            "por": {
                "official": "Ilhas Salomão",
                "common": "Ilhas Salomão"
            },
            "rus": {
                "official": "Соломоновы острова",
                "common": "Соломоновы Острова"
            },
            "spa": {
                "official": "islas Salomón",
                "common": "Islas Salomón"
            },
            "fin": {
                "official": "Salomonsaaret",
                "common": "Salomonsaaret"
            }
        },
        "latlng": [
            -8,
            159
        ],
        "demonym": "Solomon Islander",
        "landlocked": false,
        "borders": [],
        "area": 28896
    },
    {
        "name": {
            "common": "Sierra Leone",
            "official": "Republic of Sierra Leone",
            "native": {
                "eng": {
                    "official": "Republic of Sierra Leone",
                    "common": "Sierra Leone"
                }
            }
        },
        "tld": [
            ".sl"
        ],
        "cca2": "SL",
        "ccn3": "694",
        "cca3": "SLE",
        "cioc": "SLE",
        "currency": [
            "SLL"
        ],
        "callingCode": [
            "232"
        ],
        "capital": "Freetown",
        "altSpellings": [
            "SL",
            "Republic of Sierra Leone"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Sierra Leone",
                "common": "Sierra Leone"
            },
            "fra": {
                "official": "République de Sierra Leone",
                "common": "Sierra Leone"
            },
            "hrv": {
                "official": "Republika Sijera Leone",
                "common": "Sijera Leone"
            },
            "ita": {
                "official": "Repubblica della Sierra Leone",
                "common": "Sierra Leone"
            },
            "jpn": {
                "official": "シエラレオネ共和国",
                "common": "シエラレオネ"
            },
            "nld": {
                "official": "Republiek Sierra Leone",
                "common": "Sierra Leone"
            },
            "por": {
                "official": "República da Serra Leoa",
                "common": "Serra Leoa"
            },
            "rus": {
                "official": "Республика Сьерра-Леоне",
                "common": "Сьерра-Леоне"
            },
            "spa": {
                "official": "República de Sierra Leona",
                "common": "Sierra Leone"
            },
            "fin": {
                "official": "Sierra Leonen tasavalta",
                "common": "Sierra Leone"
            }
        },
        "latlng": [
            8.5,
            -11.5
        ],
        "demonym": "Sierra Leonean",
        "landlocked": false,
        "borders": [
            "GIN",
            "LBR"
        ],
        "area": 71740
    },
    {
        "name": {
            "common": "El Salvador",
            "official": "Republic of El Salvador",
            "native": {
                "spa": {
                    "official": "República de El Salvador",
                    "common": "El Salvador"
                }
            }
        },
        "tld": [
            ".sv"
        ],
        "cca2": "SV",
        "ccn3": "222",
        "cca3": "SLV",
        "cioc": "ESA",
        "currency": [
            "SVC",
            "USD"
        ],
        "callingCode": [
            "503"
        ],
        "capital": "San Salvador",
        "altSpellings": [
            "SV",
            "Republic of El Salvador",
            "República de El Salvador"
        ],
        "region": "Americas",
        "subregion": "Central America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "cym": {
                "official": "Republic of El Salvador",
                "common": "El Salvador"
            },
            "deu": {
                "official": "Republik El Salvador",
                "common": "El Salvador"
            },
            "fra": {
                "official": "République du Salvador",
                "common": "Salvador"
            },
            "hrv": {
                "official": "Republika El Salvador",
                "common": "Salvador"
            },
            "ita": {
                "official": "Repubblica di El Salvador",
                "common": "El Salvador"
            },
            "jpn": {
                "official": "エルサルバドル共和国",
                "common": "エルサルバドル"
            },
            "nld": {
                "official": "Republiek El Salvador",
                "common": "El Salvador"
            },
            "por": {
                "official": "República de El Salvador",
                "common": "El Salvador"
            },
            "rus": {
                "official": "Республика Эль-Сальвадор",
                "common": "Сальвадор"
            },
            "spa": {
                "official": "República de El Salvador",
                "common": "El Salvador"
            },
            "fin": {
                "official": "El Salvadorin tasavalta",
                "common": "El Salvador"
            }
        },
        "latlng": [
            13.83333333,
            -88.91666666
        ],
        "demonym": "Salvadoran",
        "landlocked": false,
        "borders": [
            "GTM",
            "HND"
        ],
        "area": 21041
    },
    {
        "name": {
            "common": "San Marino",
            "official": "Most Serene Republic of San Marino",
            "native": {
                "ita": {
                    "official": "Serenissima Repubblica di San Marino",
                    "common": "San Marino"
                }
            }
        },
        "tld": [
            ".sm"
        ],
        "cca2": "SM",
        "ccn3": "674",
        "cca3": "SMR",
        "cioc": "SMR",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "378"
        ],
        "capital": "City of San Marino",
        "altSpellings": [
            "SM",
            "Republic of San Marino",
            "Repubblica di San Marino"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "ita": "Italian"
        },
        "translations": {
            "deu": {
                "official": "Republik San Marino",
                "common": "San Marino"
            },
            "fra": {
                "official": "République de Saint-Marin",
                "common": "Saint-Marin"
            },
            "hrv": {
                "official": "Većina Serene Republika San Marino",
                "common": "San Marino"
            },
            "ita": {
                "official": "Serenissima Repubblica di San Marino",
                "common": "San Marino"
            },
            "jpn": {
                "official": "サンマリノのほとんどセリーヌ共和国",
                "common": "サンマリノ"
            },
            "nld": {
                "official": "Meest Serene Republiek San Marino",
                "common": "San Marino"
            },
            "por": {
                "official": "Sereníssima República de San Marino",
                "common": "San Marino"
            },
            "rus": {
                "official": "Большинство Serene Республика Сан-Марино",
                "common": "Сан-Марино"
            },
            "spa": {
                "official": "Serenísima República de San Marino",
                "common": "San Marino"
            },
            "fin": {
                "official": "San Marinon seesteinen tasavalta",
                "common": "San Marino"
            }
        },
        "latlng": [
            43.76666666,
            12.41666666
        ],
        "demonym": "Sammarinese",
        "landlocked": true,
        "borders": [
            "ITA"
        ],
        "area": 61
    },
    {
        "name": {
            "common": "Somalia",
            "official": "Federal Republic of Somalia",
            "native": {
                "ara": {
                    "official": "جمهورية الصومال‎‎",
                    "common": "الصومال‎‎"
                },
                "som": {
                    "official": "Jamhuuriyadda Federaalka Soomaaliya",
                    "common": "Soomaaliya"
                }
            }
        },
        "tld": [
            ".so"
        ],
        "cca2": "SO",
        "ccn3": "706",
        "cca3": "SOM",
        "cioc": "SOM",
        "currency": [
            "SOS"
        ],
        "callingCode": [
            "252"
        ],
        "capital": "Mogadishu",
        "altSpellings": [
            "SO",
            "aṣ-Ṣūmāl",
            "Federal Republic of Somalia",
            "Jamhuuriyadda Federaalka Soomaaliya",
            "Jumhūriyyat aṣ-Ṣūmāl al-Fiderāliyya"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "ara": "Arabic",
            "som": "Somali"
        },
        "translations": {
            "deu": {
                "official": "Bundesrepublik Somalia",
                "common": "Somalia"
            },
            "fra": {
                "official": "République fédérale de Somalie",
                "common": "Somalie"
            },
            "hrv": {
                "official": "Savezna Republika Somaliji",
                "common": "Somalija"
            },
            "ita": {
                "official": "Repubblica federale di Somalia",
                "common": "Somalia"
            },
            "jpn": {
                "official": "ソマリア連邦共和国",
                "common": "ソマリア"
            },
            "nld": {
                "official": "Federale Republiek Somalië",
                "common": "Somalië"
            },
            "por": {
                "official": "República Federal da Somália",
                "common": "Somália"
            },
            "rus": {
                "official": "Федеративная Республика Сомали",
                "common": "Сомали"
            },
            "spa": {
                "official": "República Federal de Somalia",
                "common": "Somalia"
            },
            "fin": {
                "official": "Somalian liittotasavalta",
                "common": "Somalia"
            }
        },
        "latlng": [
            10,
            49
        ],
        "demonym": "Somali",
        "landlocked": false,
        "borders": [
            "DJI",
            "ETH",
            "KEN"
        ],
        "area": 637657
    },
    {
        "name": {
            "common": "Saint Pierre and Miquelon",
            "official": "Saint Pierre and Miquelon",
            "native": {
                "fra": {
                    "official": "Collectivité territoriale de Saint-Pierre-et-Miquelon",
                    "common": "Saint-Pierre-et-Miquelon"
                }
            }
        },
        "tld": [
            ".pm"
        ],
        "cca2": "PM",
        "ccn3": "666",
        "cca3": "SPM",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "508"
        ],
        "capital": "Saint-Pierre",
        "altSpellings": [
            "PM",
            "Collectivité territoriale de Saint-Pierre-et-Miquelon"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "St. Pierre und Miquelon",
                "common": "Saint-Pierre und Miquelon"
            },
            "fra": {
                "official": "Saint-Pierre-et-Miquelon",
                "common": "Saint-Pierre-et-Miquelon"
            },
            "hrv": {
                "official": "Saint Pierre i Miquelon",
                "common": "Sveti Petar i Mikelon"
            },
            "ita": {
                "official": "Saint Pierre e Miquelon",
                "common": "Saint-Pierre e Miquelon"
            },
            "jpn": {
                "official": "サンピエール島·ミクロン島",
                "common": "サンピエール島・ミクロン島"
            },
            "nld": {
                "official": "Saint-Pierre en Miquelon",
                "common": "Saint Pierre en Miquelon"
            },
            "por": {
                "official": "Saint Pierre e Miquelon",
                "common": "Saint-Pierre e Miquelon"
            },
            "rus": {
                "official": "Сен-Пьер и Микелон",
                "common": "Сен-Пьер и Микелон"
            },
            "spa": {
                "official": "San Pedro y Miquelón",
                "common": "San Pedro y Miquelón"
            },
            "fin": {
                "official": "Saint-Pierre ja Miquelon",
                "common": "Saint-Pierre ja Miquelon"
            }
        },
        "latlng": [
            46.83333333,
            -56.33333333
        ],
        "demonym": "French",
        "landlocked": false,
        "borders": [],
        "area": 242
    },
    {
        "name": {
            "common": "Serbia",
            "official": "Republic of Serbia",
            "native": {
                "srp": {
                    "official": "Република Србија",
                    "common": "Србија"
                }
            }
        },
        "tld": [
            ".rs",
            ".срб"
        ],
        "cca2": "RS",
        "ccn3": "688",
        "cca3": "SRB",
        "cioc": "SRB",
        "currency": [
            "RSD"
        ],
        "callingCode": [
            "381"
        ],
        "capital": "Belgrade",
        "altSpellings": [
            "RS",
            "Srbija",
            "Republic of Serbia",
            "Република Србија",
            "Republika Srbija"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "srp": "Serbian"
        },
        "translations": {
            "deu": {
                "official": "Republik Serbien",
                "common": "Serbien"
            },
            "fra": {
                "official": "République de Serbie",
                "common": "Serbie"
            },
            "hrv": {
                "official": "Republika Srbija",
                "common": "Srbija"
            },
            "ita": {
                "official": "Repubblica di Serbia",
                "common": "Serbia"
            },
            "jpn": {
                "official": "セルビア共和国",
                "common": "セルビア"
            },
            "nld": {
                "official": "Republiek Servië",
                "common": "Servië"
            },
            "por": {
                "official": "República da Sérvia",
                "common": "Sérvia"
            },
            "rus": {
                "official": "Республика Сербия",
                "common": "Сербия"
            },
            "spa": {
                "official": "República de Serbia",
                "common": "Serbia"
            },
            "fin": {
                "official": "Serbian tasavalta",
                "common": "Serbia"
            }
        },
        "latlng": [
            44,
            21
        ],
        "demonym": "Serbian",
        "landlocked": true,
        "borders": [
            "BIH",
            "BGR",
            "HRV",
            "HUN",
            "KOS",
            "MKD",
            "MNE",
            "ROU"
        ],
        "area": 88361
    },
    {
        "name": {
            "common": "South Sudan",
            "official": "Republic of South Sudan",
            "native": {
                "eng": {
                    "official": "Republic of South Sudan",
                    "common": "South Sudan"
                }
            }
        },
        "tld": [
            ".ss"
        ],
        "cca2": "SS",
        "ccn3": "728",
        "cca3": "SSD",
        "cioc": "",
        "currency": [
            "SSP"
        ],
        "callingCode": [
            "211"
        ],
        "capital": "Juba",
        "altSpellings": [
            "SS"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Südsudan",
                "common": "Südsudan"
            },
            "fra": {
                "official": "République du Soudan du Sud",
                "common": "Soudan du Sud"
            },
            "hrv": {
                "official": "Republika Južni Sudan",
                "common": "Južni Sudan"
            },
            "ita": {
                "official": "Repubblica del Sudan del Sud",
                "common": "Sudan del sud"
            },
            "jpn": {
                "official": "南スーダン共和国",
                "common": "南スーダン"
            },
            "nld": {
                "official": "Republiek Zuid-Soedan",
                "common": "Zuid-Soedan"
            },
            "por": {
                "official": "República do Sudão do Sul",
                "common": "Sudão do Sul"
            },
            "rus": {
                "official": "Республика Южный Судан",
                "common": "Южный Судан"
            },
            "spa": {
                "official": "República de Sudán del Sur",
                "common": "Sudán del Sur"
            },
            "fin": {
                "official": "Etelä-Sudanin tasavalta",
                "common": "Etelä-Sudan"
            }
        },
        "latlng": [
            7,
            30
        ],
        "demonym": "South Sudanese",
        "landlocked": true,
        "borders": [
            "CAF",
            "COD",
            "ETH",
            "KEN",
            "SDN",
            "UGA"
        ],
        "area": 619745
    },
    {
        "name": {
            "common": "São Tomé and Príncipe",
            "official": "Democratic Republic of São Tomé and Príncipe",
            "native": {
                "por": {
                    "official": "República Democrática do São Tomé e Príncipe",
                    "common": "São Tomé e Príncipe"
                }
            }
        },
        "tld": [
            ".st"
        ],
        "cca2": "ST",
        "ccn3": "678",
        "cca3": "STP",
        "cioc": "STP",
        "currency": [
            "STD"
        ],
        "callingCode": [
            "239"
        ],
        "capital": "São Tomé",
        "altSpellings": [
            "ST",
            "Democratic Republic of São Tomé and Príncipe",
            "Sao Tome and Principe",
            "República Democrática de São Tomé e Príncipe"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "por": "Portuguese"
        },
        "translations": {
            "deu": {
                "official": "Demokratische Republik São Tomé und Príncipe",
                "common": "São Tomé und Príncipe"
            },
            "fra": {
                "official": "République démocratique de São Tomé et Príncipe",
                "common": "São Tomé et Príncipe"
            },
            "hrv": {
                "official": "Demokratska Republika São Tome i Principe",
                "common": "Sveti Toma i Princip"
            },
            "ita": {
                "official": "Repubblica democratica di São Tomé e Príncipe",
                "common": "São Tomé e Príncipe"
            },
            "jpn": {
                "official": "サントメ·プリンシペ民主共和国",
                "common": "サントメ・プリンシペ"
            },
            "nld": {
                "official": "Democratische Republiek Sao Tomé en Principe",
                "common": "Sao Tomé en Principe"
            },
            "por": {
                "official": "República Democrática de São Tomé e Príncipe",
                "common": "São Tomé e Príncipe"
            },
            "rus": {
                "official": "Демократическая Республика Сан-Томе и Принсипи",
                "common": "Сан-Томе и Принсипи"
            },
            "spa": {
                "official": "República Democrática de Santo Tomé y Príncipe",
                "common": "Santo Tomé y Príncipe"
            },
            "fin": {
                "official": "São Tomé ja Príncipen demokraattinen tasavalta",
                "common": "São Téme ja Príncipe"
            }
        },
        "latlng": [
            1,
            7
        ],
        "demonym": "Sao Tomean",
        "landlocked": false,
        "borders": [],
        "area": 964
    },
    {
        "name": {
            "common": "Suriname",
            "official": "Republic of Suriname",
            "native": {
                "nld": {
                    "official": "Republiek Suriname",
                    "common": "Suriname"
                }
            }
        },
        "tld": [
            ".sr"
        ],
        "cca2": "SR",
        "ccn3": "740",
        "cca3": "SUR",
        "cioc": "SUR",
        "currency": [
            "SRD"
        ],
        "callingCode": [
            "597"
        ],
        "capital": "Paramaribo",
        "altSpellings": [
            "SR",
            "Sarnam",
            "Sranangron",
            "Republic of Suriname",
            "Republiek Suriname"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "nld": "Dutch"
        },
        "translations": {
            "deu": {
                "official": "Republik Suriname",
                "common": "Suriname"
            },
            "fra": {
                "official": "République du Suriname",
                "common": "Surinam"
            },
            "hrv": {
                "official": "Republika Surinam",
                "common": "Surinam"
            },
            "ita": {
                "official": "Repubblica del Suriname",
                "common": "Suriname"
            },
            "jpn": {
                "official": "スリナム共和国",
                "common": "スリナム"
            },
            "nld": {
                "official": "Republiek Suriname",
                "common": "Suriname"
            },
            "por": {
                "official": "República do Suriname",
                "common": "Suriname"
            },
            "rus": {
                "official": "Республика Суринам",
                "common": "Суринам"
            },
            "spa": {
                "official": "República de Suriname",
                "common": "Surinam"
            },
            "fin": {
                "official": "Surinamen tasavalta",
                "common": "Suriname"
            }
        },
        "latlng": [
            4,
            -56
        ],
        "demonym": "Surinamer",
        "landlocked": false,
        "borders": [
            "BRA",
            "GUF",
            "GUY"
        ],
        "area": 163820
    },
    {
        "name": {
            "common": "Slovakia",
            "official": "Slovak Republic",
            "native": {
                "slk": {
                    "official": "Slovenská republika",
                    "common": "Slovensko"
                }
            }
        },
        "tld": [
            ".sk"
        ],
        "cca2": "SK",
        "ccn3": "703",
        "cca3": "SVK",
        "cioc": "SVK",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "421"
        ],
        "capital": "Bratislava",
        "altSpellings": [
            "SK",
            "Slovak Republic",
            "Slovenská republika"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "slk": "Slovak"
        },
        "translations": {
            "deu": {
                "official": "Slowakische Republik",
                "common": "Slowakei"
            },
            "fra": {
                "official": "République slovaque",
                "common": "Slovaquie"
            },
            "hrv": {
                "official": "slovačka",
                "common": "Slovačka"
            },
            "ita": {
                "official": "Repubblica slovacca",
                "common": "Slovacchia"
            },
            "jpn": {
                "official": "スロバキア共和国",
                "common": "スロバキア"
            },
            "nld": {
                "official": "Slowaakse Republiek",
                "common": "Slowakije"
            },
            "por": {
                "official": "República Eslovaca",
                "common": "Eslováquia"
            },
            "rus": {
                "official": "Словацкая Республика",
                "common": "Словакия"
            },
            "spa": {
                "official": "República Eslovaca",
                "common": "República Eslovaca"
            },
            "fin": {
                "official": "Slovakian tasavalta",
                "common": "Slovakia"
            }
        },
        "latlng": [
            48.66666666,
            19.5
        ],
        "demonym": "Slovak",
        "landlocked": true,
        "borders": [
            "AUT",
            "CZE",
            "HUN",
            "POL",
            "UKR"
        ],
        "area": 49037
    },
    {
        "name": {
            "common": "Slovenia",
            "official": "Republic of Slovenia",
            "native": {
                "slv": {
                    "official": "Republika Slovenija",
                    "common": "Slovenija"
                }
            }
        },
        "tld": [
            ".si"
        ],
        "cca2": "SI",
        "ccn3": "705",
        "cca3": "SVN",
        "cioc": "SLO",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "386"
        ],
        "capital": "Ljubljana",
        "altSpellings": [
            "SI",
            "Republic of Slovenia",
            "Republika Slovenija"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "slv": "Slovene"
        },
        "translations": {
            "deu": {
                "official": "Republik Slowenien",
                "common": "Slowenien"
            },
            "fra": {
                "official": "République de Slovénie",
                "common": "Slovénie"
            },
            "hrv": {
                "official": "Republika Slovenija",
                "common": "Slovenija"
            },
            "ita": {
                "official": "Repubblica di Slovenia",
                "common": "Slovenia"
            },
            "jpn": {
                "official": "スロベニア共和国",
                "common": "スロベニア"
            },
            "nld": {
                "official": "Republiek Slovenië",
                "common": "Slovenië"
            },
            "por": {
                "official": "República da Eslovénia",
                "common": "Eslovénia"
            },
            "rus": {
                "official": "Республика Словения",
                "common": "Словения"
            },
            "spa": {
                "official": "República de Eslovenia",
                "common": "Eslovenia"
            },
            "fin": {
                "official": "Slovenian tasavalta",
                "common": "Slovenia"
            }
        },
        "latlng": [
            46.11666666,
            14.81666666
        ],
        "demonym": "Slovene",
        "landlocked": false,
        "borders": [
            "AUT",
            "HRV",
            "ITA",
            "HUN"
        ],
        "area": 20273
    },
    {
        "name": {
            "common": "Sweden",
            "official": "Kingdom of Sweden",
            "native": {
                "swe": {
                    "official": "Konungariket Sverige",
                    "common": "Sverige"
                }
            }
        },
        "tld": [
            ".se"
        ],
        "cca2": "SE",
        "ccn3": "752",
        "cca3": "SWE",
        "cioc": "SWE",
        "currency": [
            "SEK"
        ],
        "callingCode": [
            "46"
        ],
        "capital": "Stockholm",
        "altSpellings": [
            "SE",
            "Kingdom of Sweden",
            "Konungariket Sverige"
        ],
        "region": "Europe",
        "subregion": "Northern Europe",
        "languages": {
            "swe": "Swedish"
        },
        "translations": {
            "deu": {
                "official": "Königreich Schweden",
                "common": "Schweden"
            },
            "fra": {
                "official": "Royaume de Suède",
                "common": "Suède"
            },
            "hrv": {
                "official": "Kraljevina Švedska",
                "common": "Švedska"
            },
            "ita": {
                "official": "Regno di Svezia",
                "common": "Svezia"
            },
            "jpn": {
                "official": "スウェーデン王国",
                "common": "スウェーデン"
            },
            "nld": {
                "official": "Koninkrijk Zweden",
                "common": "Zweden"
            },
            "por": {
                "official": "Reino da Suécia",
                "common": "Suécia"
            },
            "rus": {
                "official": "Королевство Швеция",
                "common": "Швеция"
            },
            "spa": {
                "official": "Reino de Suecia",
                "common": "Suecia"
            },
            "fin": {
                "official": "Ruotsin kuningaskunta",
                "common": "Ruotsi"
            }
        },
        "latlng": [
            62,
            15
        ],
        "demonym": "Swedish",
        "landlocked": false,
        "borders": [
            "FIN",
            "NOR"
        ],
        "area": 450295
    },
    {
        "name": {
            "common": "Swaziland",
            "official": "Kingdom of Swaziland",
            "native": {
                "eng": {
                    "official": "Kingdom of Swaziland",
                    "common": "Swaziland"
                },
                "ssw": {
                    "official": "Kingdom of Swaziland",
                    "common": "Swaziland"
                }
            }
        },
        "tld": [
            ".sz"
        ],
        "cca2": "SZ",
        "ccn3": "748",
        "cca3": "SWZ",
        "cioc": "SWZ",
        "currency": [
            "SZL"
        ],
        "callingCode": [
            "268"
        ],
        "capital": "Lobamba",
        "altSpellings": [
            "SZ",
            "weSwatini",
            "Swatini",
            "Ngwane",
            "Kingdom of Swaziland",
            "Umbuso waseSwatini"
        ],
        "region": "Africa",
        "subregion": "Southern Africa",
        "languages": {
            "eng": "English",
            "ssw": "Swazi"
        },
        "translations": {
            "deu": {
                "official": "Königreich Swasiland",
                "common": "Swasiland"
            },
            "fra": {
                "official": "Royaume du Swaziland",
                "common": "Swaziland"
            },
            "hrv": {
                "official": "Kraljevina Svazi",
                "common": "Svazi"
            },
            "ita": {
                "official": "Regno dello Swaziland",
                "common": "Swaziland"
            },
            "jpn": {
                "official": "スワジランド王国",
                "common": "スワジランド"
            },
            "nld": {
                "official": "Koninkrijk Swaziland",
                "common": "Swaziland"
            },
            "por": {
                "official": "Reino da Suazilândia",
                "common": "Suazilândia"
            },
            "rus": {
                "official": "Королевство Свазиленд",
                "common": "Свазиленд"
            },
            "spa": {
                "official": "Reino de Swazilandia",
                "common": "Suazilandia"
            },
            "fin": {
                "official": "Swazimaan kuningaskunta",
                "common": "Swazimaa"
            }
        },
        "latlng": [
            -26.5,
            31.5
        ],
        "demonym": "Swazi",
        "landlocked": true,
        "borders": [
            "MOZ",
            "ZAF"
        ],
        "area": 17364
    },
    {
        "name": {
            "common": "Sint Maarten",
            "official": "Sint Maarten",
            "native": {
                "eng": {
                    "official": "Sint Maarten",
                    "common": "Sint Maarten"
                },
                "fra": {
                    "official": "Saint-Martin",
                    "common": "Saint-Martin"
                },
                "nld": {
                    "official": "Sint Maarten",
                    "common": "Sint Maarten"
                }
            }
        },
        "tld": [
            ".sx"
        ],
        "cca2": "SX",
        "ccn3": "534",
        "cca3": "SXM",
        "cioc": "",
        "currency": [
            "ANG"
        ],
        "callingCode": [
            "1721"
        ],
        "capital": "Philipsburg",
        "altSpellings": [
            "SX",
            "Sint Maarten (Dutch part)"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English",
            "fra": "French",
            "nld": "Dutch"
        },
        "translations": {
            "deu": {
                "official": "Sint Maarten",
                "common": "Sint Maarten"
            },
            "fra": {
                "official": "Sint Maarten",
                "common": "Saint-Martin"
            },
            "ita": {
                "official": "Sint Maarten",
                "common": "Sint Maarten"
            },
            "jpn": {
                "official": "シントマールテン島",
                "common": "シント・マールテン"
            },
            "nld": {
                "official": "Sint Maarten",
                "common": "Sint Maarten"
            },
            "por": {
                "official": "Sint Maarten",
                "common": "São Martinho"
            },
            "rus": {
                "official": "Синт-Маартен",
                "common": "Синт-Мартен"
            },
            "spa": {
                "official": "Sint Maarten",
                "common": "Sint Maarten"
            },
            "fin": {
                "official": "Sint Maarten",
                "common": "Sint Maarten"
            }
        },
        "latlng": [
            18.033333,
            -63.05
        ],
        "demonym": "St. Maartener",
        "landlocked": false,
        "borders": [
            "MAF"
        ],
        "area": 34
    },
    {
        "name": {
            "common": "Seychelles",
            "official": "Republic of Seychelles",
            "native": {
                "crs": {
                    "official": "Repiblik Sesel",
                    "common": "Sesel"
                },
                "eng": {
                    "official": "Republic of Seychelles",
                    "common": "Seychelles"
                },
                "fra": {
                    "official": "République des Seychelles",
                    "common": "Seychelles"
                }
            }
        },
        "tld": [
            ".sc"
        ],
        "cca2": "SC",
        "ccn3": "690",
        "cca3": "SYC",
        "cioc": "SEY",
        "currency": [
            "SCR"
        ],
        "callingCode": [
            "248"
        ],
        "capital": "Victoria",
        "altSpellings": [
            "SC",
            "Republic of Seychelles",
            "Repiblik Sesel",
            "République des Seychelles"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "crs": "Seychellois Creole",
            "eng": "English",
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik der Seychellen",
                "common": "Seychellen"
            },
            "fra": {
                "official": "République des Seychelles",
                "common": "Seychelles"
            },
            "hrv": {
                "official": "Republika Sejšeli",
                "common": "Sejšeli"
            },
            "ita": {
                "official": "Repubblica delle Seychelles",
                "common": "Seychelles"
            },
            "jpn": {
                "official": "セイシェル共和国",
                "common": "セーシェル"
            },
            "nld": {
                "official": "Republiek der Seychellen",
                "common": "Seychellen"
            },
            "por": {
                "official": "República das Seychelles",
                "common": "Seicheles"
            },
            "rus": {
                "official": "Республика Сейшельские Острова",
                "common": "Сейшельские Острова"
            },
            "spa": {
                "official": "República de las Seychelles",
                "common": "Seychelles"
            },
            "fin": {
                "official": "Seychellien tasavalta",
                "common": "Seychellit"
            }
        },
        "latlng": [
            -4.58333333,
            55.66666666
        ],
        "demonym": "Seychellois",
        "landlocked": false,
        "borders": [],
        "area": 452
    },
    {
        "name": {
            "common": "Syria",
            "official": "Syrian Arab Republic",
            "native": {
                "ara": {
                    "official": "الجمهورية العربية السورية",
                    "common": "سوريا"
                }
            }
        },
        "tld": [
            ".sy",
            "سوريا."
        ],
        "cca2": "SY",
        "ccn3": "760",
        "cca3": "SYR",
        "cioc": "SYR",
        "currency": [
            "SYP"
        ],
        "callingCode": [
            "963"
        ],
        "capital": "Damascus",
        "altSpellings": [
            "SY",
            "Syrian Arab Republic",
            "Al-Jumhūrīyah Al-ʻArabīyah As-Sūrīyah"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Arabische Republik Syrien",
                "common": "Syrien"
            },
            "fra": {
                "official": "République arabe syrienne",
                "common": "Syrie"
            },
            "hrv": {
                "official": "Sirijska Arapska Republika",
                "common": "Sirija"
            },
            "ita": {
                "official": "Repubblica araba siriana",
                "common": "Siria"
            },
            "jpn": {
                "official": "シリアアラブ共和国",
                "common": "シリア・アラブ共和国"
            },
            "nld": {
                "official": "Syrische Arabische Republiek",
                "common": "Syrië"
            },
            "por": {
                "official": "República Árabe Síria",
                "common": "Síria"
            },
            "rus": {
                "official": "Сирийская Арабская Республика",
                "common": "Сирия"
            },
            "spa": {
                "official": "República Árabe Siria",
                "common": "Siria"
            },
            "fin": {
                "official": "Syyrian arabitasavalta",
                "common": "Syyria"
            }
        },
        "latlng": [
            35,
            38
        ],
        "demonym": "Syrian",
        "landlocked": false,
        "borders": [
            "IRQ",
            "ISR",
            "JOR",
            "LBN",
            "TUR"
        ],
        "area": 185180
    },
    {
        "name": {
            "common": "Turks and Caicos Islands",
            "official": "Turks and Caicos Islands",
            "native": {
                "eng": {
                    "official": "Turks and Caicos Islands",
                    "common": "Turks and Caicos Islands"
                }
            }
        },
        "tld": [
            ".tc"
        ],
        "cca2": "TC",
        "ccn3": "796",
        "cca3": "TCA",
        "cioc": "",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1649"
        ],
        "capital": "Cockburn Town",
        "altSpellings": [
            "TC"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Turks und Caicos Inseln",
                "common": "Turks-und Caicosinseln"
            },
            "fra": {
                "official": "Îles Turques et Caïques",
                "common": "Îles Turques-et-Caïques"
            },
            "hrv": {
                "official": "Otoci Turks i Caicos",
                "common": "Otoci Turks i Caicos"
            },
            "ita": {
                "official": "Turks e Caicos",
                "common": "Isole Turks e Caicos"
            },
            "jpn": {
                "official": "タークス·カイコス諸島",
                "common": "タークス・カイコス諸島"
            },
            "nld": {
                "official": "Turks-en Caicoseilanden",
                "common": "Turks-en Caicoseilanden"
            },
            "por": {
                "official": "Ilhas Turks e Caicos",
                "common": "Ilhas Turks e Caicos"
            },
            "rus": {
                "official": "Теркс и Кайкос острова",
                "common": "Теркс и Кайкос"
            },
            "spa": {
                "official": "Islas Turcas y Caicos",
                "common": "Islas Turks y Caicos"
            },
            "fin": {
                "official": "Turks-ja Caicossaaret",
                "common": "Turks-ja Caicossaaret"
            }
        },
        "latlng": [
            21.75,
            -71.58333333
        ],
        "demonym": "Turks and Caicos Islander",
        "landlocked": false,
        "borders": [],
        "area": 948
    },
    {
        "name": {
            "common": "Chad",
            "official": "Republic of Chad",
            "native": {
                "ara": {
                    "official": "جمهورية تشاد",
                    "common": "تشاد‎"
                },
                "fra": {
                    "official": "République du Tchad",
                    "common": "Tchad"
                }
            }
        },
        "tld": [
            ".td"
        ],
        "cca2": "TD",
        "ccn3": "148",
        "cca3": "TCD",
        "cioc": "CHA",
        "currency": [
            "XAF"
        ],
        "callingCode": [
            "235"
        ],
        "capital": "N'Djamena",
        "altSpellings": [
            "TD",
            "Tchad",
            "Republic of Chad",
            "République du Tchad"
        ],
        "region": "Africa",
        "subregion": "Middle Africa",
        "languages": {
            "ara": "Arabic",
            "fra": "French"
        },
        "translations": {
            "cym": {
                "official": "Republic of Chad",
                "common": "Tsiad"
            },
            "deu": {
                "official": "Republik Tschad",
                "common": "Tschad"
            },
            "fra": {
                "official": "République du Tchad",
                "common": "Tchad"
            },
            "hrv": {
                "official": "Čadu",
                "common": "Čad"
            },
            "ita": {
                "official": "Repubblica del Ciad",
                "common": "Ciad"
            },
            "jpn": {
                "official": "チャド共和国",
                "common": "チャド"
            },
            "nld": {
                "official": "Republiek Tsjaad",
                "common": "Tsjaad"
            },
            "por": {
                "official": "República do Chade",
                "common": "Chade"
            },
            "rus": {
                "official": "Республика Чад",
                "common": "Чад"
            },
            "spa": {
                "official": "República de Chad",
                "common": "Chad"
            },
            "fin": {
                "official": "Tšadin tasavalta",
                "common": "Tšad"
            }
        },
        "latlng": [
            15,
            19
        ],
        "demonym": "Chadian",
        "landlocked": true,
        "borders": [
            "CMR",
            "CAF",
            "LBY",
            "NER",
            "NGA",
            "SSD"
        ],
        "area": 1284000
    },
    {
        "name": {
            "common": "Togo",
            "official": "Togolese Republic",
            "native": {
                "fra": {
                    "official": "République togolaise",
                    "common": "Togo"
                }
            }
        },
        "tld": [
            ".tg"
        ],
        "cca2": "TG",
        "ccn3": "768",
        "cca3": "TGO",
        "cioc": "TOG",
        "currency": [
            "XOF"
        ],
        "callingCode": [
            "228"
        ],
        "capital": "Lomé",
        "altSpellings": [
            "TG",
            "Togolese",
            "Togolese Republic",
            "République Togolaise"
        ],
        "region": "Africa",
        "subregion": "Western Africa",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Republik Togo",
                "common": "Togo"
            },
            "fra": {
                "official": "République togolaise",
                "common": "Togo"
            },
            "hrv": {
                "official": "Togolese Republika",
                "common": "Togo"
            },
            "ita": {
                "official": "Repubblica del Togo",
                "common": "Togo"
            },
            "jpn": {
                "official": "トーゴ共和国",
                "common": "トーゴ"
            },
            "nld": {
                "official": "Republiek Togo",
                "common": "Togo"
            },
            "por": {
                "official": "República do Togo",
                "common": "Togo"
            },
            "rus": {
                "official": "Того Республика",
                "common": "Того"
            },
            "spa": {
                "official": "República de Togo",
                "common": "Togo"
            },
            "fin": {
                "official": "Togon tasavalta",
                "common": "Togo"
            }
        },
        "latlng": [
            8,
            1.16666666
        ],
        "demonym": "Togolese",
        "landlocked": false,
        "borders": [
            "BEN",
            "BFA",
            "GHA"
        ],
        "area": 56785
    },
    {
        "name": {
            "common": "Thailand",
            "official": "Kingdom of Thailand",
            "native": {
                "tha": {
                    "official": "ราชอาณาจักรไทย",
                    "common": "ประเทศไทย"
                }
            }
        },
        "tld": [
            ".th",
            ".ไทย"
        ],
        "cca2": "TH",
        "ccn3": "764",
        "cca3": "THA",
        "cioc": "THA",
        "currency": [
            "THB"
        ],
        "callingCode": [
            "66"
        ],
        "capital": "Bangkok",
        "altSpellings": [
            "TH",
            "Prathet",
            "Thai",
            "Kingdom of Thailand",
            "ราชอาณาจักรไทย",
            "Ratcha Anachak Thai"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "tha": "Thai"
        },
        "translations": {
            "deu": {
                "official": "Königreich Thailand",
                "common": "Thailand"
            },
            "fra": {
                "official": "Royaume de Thaïlande",
                "common": "Thaïlande"
            },
            "hrv": {
                "official": "Kraljevina Tajland",
                "common": "Tajland"
            },
            "ita": {
                "official": "Regno di Thailandia",
                "common": "Tailandia"
            },
            "jpn": {
                "official": "タイ王国",
                "common": "タイ"
            },
            "nld": {
                "official": "Koninkrijk Thailand",
                "common": "Thailand"
            },
            "por": {
                "official": "Reino da Tailândia",
                "common": "Tailândia"
            },
            "rus": {
                "official": "Королевство Таиланд",
                "common": "Таиланд"
            },
            "spa": {
                "official": "Reino de Tailandia",
                "common": "Tailandia"
            },
            "fin": {
                "official": "Thaimaan kuningaskunta",
                "common": "Thaimaa"
            }
        },
        "latlng": [
            15,
            100
        ],
        "demonym": "Thai",
        "landlocked": false,
        "borders": [
            "MMR",
            "KHM",
            "LAO",
            "MYS"
        ],
        "area": 513120
    },
    {
        "name": {
            "common": "Tajikistan",
            "official": "Republic of Tajikistan",
            "native": {
                "rus": {
                    "official": "Республика Таджикистан",
                    "common": "Таджикистан"
                },
                "tgk": {
                    "official": "Ҷумҳурии Тоҷикистон",
                    "common": "Тоҷикистон"
                }
            }
        },
        "tld": [
            ".tj"
        ],
        "cca2": "TJ",
        "ccn3": "762",
        "cca3": "TJK",
        "cioc": "TJK",
        "currency": [
            "TJS"
        ],
        "callingCode": [
            "992"
        ],
        "capital": "Dushanbe",
        "altSpellings": [
            "TJ",
            "Toçikiston",
            "Republic of Tajikistan",
            "Ҷумҳурии Тоҷикистон",
            "Çumhuriyi Toçikiston"
        ],
        "region": "Asia",
        "subregion": "Central Asia",
        "languages": {
            "rus": "Russian",
            "tgk": "Tajik"
        },
        "translations": {
            "deu": {
                "official": "Republik Tadschikistan",
                "common": "Tadschikistan"
            },
            "fra": {
                "official": "République du Tadjikistan",
                "common": "Tadjikistan"
            },
            "hrv": {
                "official": "Republika Tadžikistan",
                "common": "Tađikistan"
            },
            "ita": {
                "official": "Repubblica del Tajikistan",
                "common": "Tagikistan"
            },
            "jpn": {
                "official": "タジキスタン共和国",
                "common": "タジキスタン"
            },
            "nld": {
                "official": "Tadzjikistan",
                "common": "Tadzjikistan"
            },
            "por": {
                "official": "República do Tajiquistão",
                "common": "Tajiquistão"
            },
            "rus": {
                "official": "Республика Таджикистан",
                "common": "Таджикистан"
            },
            "spa": {
                "official": "República de Tayikistán",
                "common": "Tayikistán"
            },
            "fin": {
                "official": "Tadžikistanin tasavalta",
                "common": "Tadžikistan"
            }
        },
        "latlng": [
            39,
            71
        ],
        "demonym": "Tadzhik",
        "landlocked": true,
        "borders": [
            "AFG",
            "CHN",
            "KGZ",
            "UZB"
        ],
        "area": 143100
    },
    {
        "name": {
            "common": "Tokelau",
            "official": "Tokelau",
            "native": {
                "eng": {
                    "official": "Tokelau",
                    "common": "Tokelau"
                },
                "smo": {
                    "official": "Tokelau",
                    "common": "Tokelau"
                },
                "tkl": {
                    "official": "Tokelau",
                    "common": "Tokelau"
                }
            }
        },
        "tld": [
            ".tk"
        ],
        "cca2": "TK",
        "ccn3": "772",
        "cca3": "TKL",
        "cioc": "",
        "currency": [
            "NZD"
        ],
        "callingCode": [
            "690"
        ],
        "capital": "Fakaofo",
        "altSpellings": [
            "TK"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "smo": "Samoan",
            "tkl": "Tokelauan"
        },
        "translations": {
            "deu": {
                "official": "Tokelau",
                "common": "Tokelau"
            },
            "fra": {
                "official": "Îles Tokelau",
                "common": "Tokelau"
            },
            "hrv": {
                "official": "Tokelau",
                "common": "Tokelau"
            },
            "ita": {
                "official": "Tokelau",
                "common": "Isole Tokelau"
            },
            "jpn": {
                "official": "トケラウ諸島",
                "common": "トケラウ"
            },
            "nld": {
                "official": "Tokelau",
                "common": "Tokelau"
            },
            "por": {
                "official": "Tokelau",
                "common": "Tokelau"
            },
            "rus": {
                "official": "Токелау",
                "common": "Токелау"
            },
            "spa": {
                "official": "Tokelau",
                "common": "Islas Tokelau"
            },
            "fin": {
                "official": "Tokelau",
                "common": "Tokelau"
            }
        },
        "latlng": [
            -9,
            -172
        ],
        "demonym": "Tokelauan",
        "landlocked": false,
        "borders": [],
        "area": 12
    },
    {
        "name": {
            "common": "Turkmenistan",
            "official": "Turkmenistan",
            "native": {
                "rus": {
                    "official": "Туркменистан",
                    "common": "Туркмения"
                },
                "tuk": {
                    "official": "Türkmenistan",
                    "common": "Türkmenistan"
                }
            }
        },
        "tld": [
            ".tm"
        ],
        "cca2": "TM",
        "ccn3": "795",
        "cca3": "TKM",
        "cioc": "TKM",
        "currency": [
            "TMT"
        ],
        "callingCode": [
            "993"
        ],
        "capital": "Ashgabat",
        "altSpellings": [
            "TM"
        ],
        "region": "Asia",
        "subregion": "Central Asia",
        "languages": {
            "rus": "Russian",
            "tuk": "Turkmen"
        },
        "translations": {
            "deu": {
                "official": "Turkmenistan",
                "common": "Turkmenistan"
            },
            "fra": {
                "official": "Turkménistan",
                "common": "Turkménistan"
            },
            "hrv": {
                "official": "Turkmenistan",
                "common": "Turkmenistan"
            },
            "ita": {
                "official": "Turkmenistan",
                "common": "Turkmenistan"
            },
            "jpn": {
                "official": "トルクメニスタン",
                "common": "トルクメニスタン"
            },
            "nld": {
                "official": "Turkmenistan",
                "common": "Turkmenistan"
            },
            "por": {
                "official": "Turcomenistão",
                "common": "Turquemenistão"
            },
            "rus": {
                "official": "Туркменистан",
                "common": "Туркмения"
            },
            "spa": {
                "official": "Turkmenistán",
                "common": "Turkmenistán"
            },
            "fin": {
                "official": "Turkmenistan",
                "common": "Turkmenistan"
            }
        },
        "latlng": [
            40,
            60
        ],
        "demonym": "Turkmen",
        "landlocked": true,
        "borders": [
            "AFG",
            "IRN",
            "KAZ",
            "UZB"
        ],
        "area": 488100
    },
    {
        "name": {
            "common": "Timor-Leste",
            "official": "Democratic Republic of Timor-Leste",
            "native": {
                "por": {
                    "official": "República Democrática de Timor-Leste",
                    "common": "Timor-Leste"
                },
                "tet": {
                    "official": "Repúblika Demokrátika Timór-Leste",
                    "common": "Timór-Leste"
                }
            }
        },
        "tld": [
            ".tl"
        ],
        "cca2": "TL",
        "ccn3": "626",
        "cca3": "TLS",
        "cioc": "TLS",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "670"
        ],
        "capital": "Dili",
        "altSpellings": [
            "TL",
            "East Timor",
            "Democratic Republic of Timor-Leste",
            "República Democrática de Timor-Leste",
            "Repúblika Demokrátika Timór-Leste",
            "Timór Lorosa'e",
            "Timor Lorosae"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "por": "Portuguese",
            "tet": "Tetum"
        },
        "translations": {
            "deu": {
                "official": "Demokratische Republik Timor-Leste",
                "common": "Timor-Leste"
            },
            "fra": {
                "official": "République démocratique du Timor oriental",
                "common": "Timor oriental"
            },
            "hrv": {
                "official": "Demokratska Republika Timor-Leste",
                "common": "Istočni Timor"
            },
            "ita": {
                "official": "Repubblica Democratica di Timor Est",
                "common": "Timor Est"
            },
            "jpn": {
                "official": "東ティモール民主共和国",
                "common": "東ティモール"
            },
            "nld": {
                "official": "Democratische Republiek Oost-Timor",
                "common": "Oost-Timor"
            },
            "por": {
                "official": "República Democrática de Timor-Leste",
                "common": "Timor-Leste"
            },
            "rus": {
                "official": "Демократическая Республика Тимор -Лешти",
                "common": "Восточный Тимор"
            },
            "spa": {
                "official": "República Democrática de Timor-Leste",
                "common": "Timor Oriental"
            },
            "fin": {
                "official": "Itä-Timorin demokraattinen tasavalta",
                "common": "Itä-Timor"
            }
        },
        "latlng": [
            -8.83333333,
            125.91666666
        ],
        "demonym": "East Timorese",
        "landlocked": false,
        "borders": [
            "IDN"
        ],
        "area": 14874
    },
    {
        "name": {
            "common": "Tonga",
            "official": "Kingdom of Tonga",
            "native": {
                "eng": {
                    "official": "Kingdom of Tonga",
                    "common": "Tonga"
                },
                "ton": {
                    "official": "Kingdom of Tonga",
                    "common": "Tonga"
                }
            }
        },
        "tld": [
            ".to"
        ],
        "cca2": "TO",
        "ccn3": "776",
        "cca3": "TON",
        "cioc": "TGA",
        "currency": [
            "TOP"
        ],
        "callingCode": [
            "676"
        ],
        "capital": "Nuku'alofa",
        "altSpellings": [
            "TO"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "ton": "Tongan"
        },
        "translations": {
            "deu": {
                "official": "Königreich Tonga",
                "common": "Tonga"
            },
            "fra": {
                "official": "Royaume des Tonga",
                "common": "Tonga"
            },
            "hrv": {
                "official": "Kraljevina Tonga",
                "common": "Tonga"
            },
            "ita": {
                "official": "Regno di Tonga",
                "common": "Tonga"
            },
            "jpn": {
                "official": "トンガ王国",
                "common": "トンガ"
            },
            "nld": {
                "official": "Koninkrijk Tonga",
                "common": "Tonga"
            },
            "por": {
                "official": "Reino de Tonga",
                "common": "Tonga"
            },
            "rus": {
                "official": "Королевство Тонга",
                "common": "Тонга"
            },
            "spa": {
                "official": "Reino de Tonga",
                "common": "Tonga"
            },
            "fin": {
                "official": "Tongan kuningaskunta",
                "common": "Tonga"
            }
        },
        "latlng": [
            -20,
            -175
        ],
        "demonym": "Tongan",
        "landlocked": false,
        "borders": [],
        "area": 747
    },
    {
        "name": {
            "common": "Trinidad and Tobago",
            "official": "Republic of Trinidad and Tobago",
            "native": {
                "eng": {
                    "official": "Republic of Trinidad and Tobago",
                    "common": "Trinidad and Tobago"
                }
            }
        },
        "tld": [
            ".tt"
        ],
        "cca2": "TT",
        "ccn3": "780",
        "cca3": "TTO",
        "cioc": "TTO",
        "currency": [
            "TTD"
        ],
        "callingCode": [
            "1868"
        ],
        "capital": "Port of Spain",
        "altSpellings": [
            "TT",
            "Republic of Trinidad and Tobago"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Trinidad und Tobago",
                "common": "Trinidad und Tobago"
            },
            "fra": {
                "official": "République de Trinité-et-Tobago",
                "common": "Trinité-et-Tobago"
            },
            "hrv": {
                "official": "Republika Trinidad i Tobago",
                "common": "Trinidad i Tobago"
            },
            "ita": {
                "official": "Repubblica di Trinidad e Tobago",
                "common": "Trinidad e Tobago"
            },
            "jpn": {
                "official": "トリニダード·トバゴ共和国",
                "common": "トリニダード・トバゴ"
            },
            "nld": {
                "official": "Republiek Trinidad en Tobago",
                "common": "Trinidad en Tobago"
            },
            "por": {
                "official": "República de Trinidad e Tobago",
                "common": "Trinidade e Tobago"
            },
            "rus": {
                "official": "Республика Тринидад и Тобаго",
                "common": "Тринидад и Тобаго"
            },
            "spa": {
                "official": "República de Trinidad y Tobago",
                "common": "Trinidad y Tobago"
            },
            "fin": {
                "official": "Trinidadin ja Tobagon tasavalta",
                "common": "Trinidad ja Tobago"
            }
        },
        "latlng": [
            11,
            -61
        ],
        "demonym": "Trinidadian",
        "landlocked": false,
        "borders": [],
        "area": 5130
    },
    {
        "name": {
            "common": "Tunisia",
            "official": "Tunisian Republic",
            "native": {
                "ara": {
                    "official": "الجمهورية التونسية",
                    "common": "تونس"
                }
            }
        },
        "tld": [
            ".tn"
        ],
        "cca2": "TN",
        "ccn3": "788",
        "cca3": "TUN",
        "cioc": "TUN",
        "currency": [
            "TND"
        ],
        "callingCode": [
            "216"
        ],
        "capital": "Tunis",
        "altSpellings": [
            "TN",
            "Republic of Tunisia",
            "al-Jumhūriyyah at-Tūnisiyyah"
        ],
        "region": "Africa",
        "subregion": "Northern Africa",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Tunesische Republik",
                "common": "Tunesien"
            },
            "fra": {
                "official": "République tunisienne",
                "common": "Tunisie"
            },
            "hrv": {
                "official": "Tuniski Republika",
                "common": "Tunis"
            },
            "ita": {
                "official": "Repubblica tunisina",
                "common": "Tunisia"
            },
            "jpn": {
                "official": "チュニジア共和国",
                "common": "チュニジア"
            },
            "nld": {
                "official": "Republiek Tunesië",
                "common": "Tunesië"
            },
            "por": {
                "official": "República da Tunísia",
                "common": "Tunísia"
            },
            "rus": {
                "official": "Тунисской Республики",
                "common": "Тунис"
            },
            "spa": {
                "official": "República de Túnez",
                "common": "Túnez"
            },
            "fin": {
                "official": "Tunisian tasavalta",
                "common": "Tunisia"
            }
        },
        "latlng": [
            34,
            9
        ],
        "demonym": "Tunisian",
        "landlocked": false,
        "borders": [
            "DZA",
            "LBY"
        ],
        "area": 163610
    },
    {
        "name": {
            "common": "Turkey",
            "official": "Republic of Turkey",
            "native": {
                "tur": {
                    "official": "Türkiye Cumhuriyeti",
                    "common": "Türkiye"
                }
            }
        },
        "tld": [
            ".tr"
        ],
        "cca2": "TR",
        "ccn3": "792",
        "cca3": "TUR",
        "cioc": "TUR",
        "currency": [
            "TRY"
        ],
        "callingCode": [
            "90"
        ],
        "capital": "Ankara",
        "altSpellings": [
            "TR",
            "Turkiye",
            "Republic of Turkey",
            "Türkiye Cumhuriyeti"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "tur": "Turkish"
        },
        "translations": {
            "deu": {
                "official": "Republik Türkei",
                "common": "Türkei"
            },
            "fra": {
                "official": "République de Turquie",
                "common": "Turquie"
            },
            "hrv": {
                "official": "Republika Turska",
                "common": "Turska"
            },
            "ita": {
                "official": "Repubblica di Turchia",
                "common": "Turchia"
            },
            "jpn": {
                "official": "トルコ共和国",
                "common": "トルコ"
            },
            "nld": {
                "official": "Republiek Turkije",
                "common": "Turkije"
            },
            "por": {
                "official": "República da Turquia",
                "common": "Turquia"
            },
            "rus": {
                "official": "Республика Турции",
                "common": "Турция"
            },
            "spa": {
                "official": "República de Turquía",
                "common": "Turquía"
            },
            "fin": {
                "official": "Turkin tasavalta",
                "common": "Turkki"
            }
        },
        "latlng": [
            39,
            35
        ],
        "demonym": "Turkish",
        "landlocked": false,
        "borders": [
            "ARM",
            "AZE",
            "BGR",
            "GEO",
            "GRC",
            "IRN",
            "IRQ",
            "SYR"
        ],
        "area": 783562
    },
    {
        "name": {
            "common": "Tuvalu",
            "official": "Tuvalu",
            "native": {
                "eng": {
                    "official": "Tuvalu",
                    "common": "Tuvalu"
                },
                "tvl": {
                    "official": "Tuvalu",
                    "common": "Tuvalu"
                }
            }
        },
        "tld": [
            ".tv"
        ],
        "cca2": "TV",
        "ccn3": "798",
        "cca3": "TUV",
        "cioc": "TUV",
        "currency": [
            "AUD"
        ],
        "callingCode": [
            "688"
        ],
        "capital": "Funafuti",
        "altSpellings": [
            "TV"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "tvl": "Tuvaluan"
        },
        "translations": {
            "deu": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "fra": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "hrv": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "ita": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "jpn": {
                "official": "ツバル",
                "common": "ツバル"
            },
            "nld": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "por": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "rus": {
                "official": "Тувалу",
                "common": "Тувалу"
            },
            "spa": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            },
            "fin": {
                "official": "Tuvalu",
                "common": "Tuvalu"
            }
        },
        "latlng": [
            -8,
            178
        ],
        "demonym": "Tuvaluan",
        "landlocked": false,
        "borders": [],
        "area": 26
    },
    {
        "name": {
            "common": "Taiwan",
            "official": "Republic of China (Taiwan)",
            "native": {
                "cmn": {
                    "official": "中华民国",
                    "common": "臺灣"
                }
            }
        },
        "tld": [
            ".tw",
            ".台湾",
            ".台灣"
        ],
        "cca2": "TW",
        "ccn3": "158",
        "cca3": "TWN",
        "cioc": "TPE",
        "currency": [
            "TWD"
        ],
        "callingCode": [
            "886"
        ],
        "capital": "Taipei",
        "altSpellings": [
            "TW",
            "Táiwān",
            "Republic of China",
            "中華民國",
            "Zhōnghuá Mínguó",
            "Chinese Taipei for IOC",
            "Taiwan, Province of China"
        ],
        "region": "Asia",
        "subregion": "Eastern Asia",
        "languages": {
            "cmn": "Mandarin"
        },
        "translations": {
            "deu": {
                "official": "Republik China (Taiwan)",
                "common": "Taiwan"
            },
            "fra": {
                "official": "République de Chine (Taïwan)",
                "common": "Taïwan"
            },
            "hrv": {
                "official": "Republika Kina",
                "common": "Tajvan"
            },
            "ita": {
                "official": "Repubblica cinese (Taiwan)",
                "common": "Taiwan"
            },
            "jpn": {
                "official": "中華民国",
                "common": "台湾（台湾省/中華民国）"
            },
            "nld": {
                "official": "Republiek China (Taiwan)",
                "common": "Taiwan"
            },
            "por": {
                "official": "República da China",
                "common": "Ilha Formosa"
            },
            "rus": {
                "official": "Китайская Республика",
                "common": "Тайвань"
            },
            "spa": {
                "official": "República de China en Taiwán",
                "common": "Taiwán"
            },
            "fin": {
                "official": "Kiinan tasavalta",
                "common": "Taiwan"
            }
        },
        "latlng": [
            23.5,
            121
        ],
        "demonym": "Taiwanese",
        "landlocked": false,
        "borders": [],
        "area": 36193
    },
    {
        "name": {
            "common": "Tanzania",
            "official": "United Republic of Tanzania",
            "native": {
                "eng": {
                    "official": "United Republic of Tanzania",
                    "common": "Tanzania"
                },
                "swa": {
                    "official": "Jamhuri ya Muungano wa Tanzania",
                    "common": "Tanzania"
                }
            }
        },
        "tld": [
            ".tz"
        ],
        "cca2": "TZ",
        "ccn3": "834",
        "cca3": "TZA",
        "cioc": "TAN",
        "currency": [
            "TZS"
        ],
        "callingCode": [
            "255"
        ],
        "capital": "Dodoma",
        "altSpellings": [
            "TZ",
            "Tanzania, United Republic of",
            "United Republic of Tanzania",
            "Jamhuri ya Muungano wa Tanzania"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "swa": "Swahili"
        },
        "translations": {
            "deu": {
                "official": "Vereinigte Republik Tansania",
                "common": "Tansania"
            },
            "fra": {
                "official": "République -Unie de Tanzanie",
                "common": "Tanzanie"
            },
            "hrv": {
                "official": "Ujedinjena Republika Tanzanija",
                "common": "Tanzanija"
            },
            "ita": {
                "official": "Repubblica Unita di Tanzania",
                "common": "Tanzania"
            },
            "jpn": {
                "official": "タンザニア連合共和国",
                "common": "タンザニア"
            },
            "nld": {
                "official": "Verenigde Republiek Tanzania",
                "common": "Tanzania"
            },
            "por": {
                "official": "República Unida da Tanzânia",
                "common": "Tanzânia"
            },
            "rus": {
                "official": "Объединенная Республика Танзания",
                "common": "Танзания"
            },
            "spa": {
                "official": "República Unida de Tanzania",
                "common": "Tanzania"
            },
            "fin": {
                "official": "Tansanian yhdistynyt tasavalta",
                "common": "Tansania"
            }
        },
        "latlng": [
            -6,
            35
        ],
        "demonym": "Tanzanian",
        "landlocked": false,
        "borders": [
            "BDI",
            "COD",
            "KEN",
            "MWI",
            "MOZ",
            "RWA",
            "UGA",
            "ZMB"
        ],
        "area": 945087
    },
    {
        "name": {
            "common": "Uganda",
            "official": "Republic of Uganda",
            "native": {
                "eng": {
                    "official": "Republic of Uganda",
                    "common": "Uganda"
                },
                "swa": {
                    "official": "Republic of Uganda",
                    "common": "Uganda"
                }
            }
        },
        "tld": [
            ".ug"
        ],
        "cca2": "UG",
        "ccn3": "800",
        "cca3": "UGA",
        "cioc": "UGA",
        "currency": [
            "UGX"
        ],
        "callingCode": [
            "256"
        ],
        "capital": "Kampala",
        "altSpellings": [
            "UG",
            "Republic of Uganda",
            "Jamhuri ya Uganda"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English",
            "swa": "Swahili"
        },
        "translations": {
            "deu": {
                "official": "Republik Uganda",
                "common": "Uganda"
            },
            "fra": {
                "official": "République de l'Ouganda",
                "common": "Ouganda"
            },
            "hrv": {
                "official": "Republika Uganda",
                "common": "Uganda"
            },
            "ita": {
                "official": "Repubblica di Uganda",
                "common": "Uganda"
            },
            "jpn": {
                "official": "ウガンダ共和国",
                "common": "ウガンダ"
            },
            "nld": {
                "official": "Republiek Uganda",
                "common": "Oeganda"
            },
            "por": {
                "official": "República do Uganda",
                "common": "Uganda"
            },
            "rus": {
                "official": "Республика Уганда",
                "common": "Уганда"
            },
            "spa": {
                "official": "República de Uganda",
                "common": "Uganda"
            },
            "fin": {
                "official": "Ugandan tasavalta",
                "common": "Uganda"
            }
        },
        "latlng": [
            1,
            32
        ],
        "demonym": "Ugandan",
        "landlocked": true,
        "borders": [
            "COD",
            "KEN",
            "RWA",
            "SSD",
            "TZA"
        ],
        "area": 241550
    },
    {
        "name": {
            "common": "Ukraine",
            "official": "Ukraine",
            "native": {
                "rus": {
                    "official": "Украина",
                    "common": "Украина"
                },
                "ukr": {
                    "official": "Україна",
                    "common": "Україна"
                }
            }
        },
        "tld": [
            ".ua",
            ".укр"
        ],
        "cca2": "UA",
        "ccn3": "804",
        "cca3": "UKR",
        "cioc": "UKR",
        "currency": [
            "UAH"
        ],
        "callingCode": [
            "380"
        ],
        "capital": "Kiev",
        "altSpellings": [
            "UA",
            "Ukrayina"
        ],
        "region": "Europe",
        "subregion": "Eastern Europe",
        "languages": {
            "rus": "Russian",
            "ukr": "Ukrainian"
        },
        "translations": {
            "deu": {
                "official": "Ukraine",
                "common": "Ukraine"
            },
            "fra": {
                "official": "Ukraine",
                "common": "Ukraine"
            },
            "hrv": {
                "official": "Ukrajina",
                "common": "Ukrajina"
            },
            "ita": {
                "official": "Ucraina",
                "common": "Ucraina"
            },
            "jpn": {
                "official": "ウクライナ",
                "common": "ウクライナ"
            },
            "nld": {
                "official": "Oekraïne",
                "common": "Oekraïne"
            },
            "por": {
                "official": "Ucrânia",
                "common": "Ucrânia"
            },
            "rus": {
                "official": "Украина",
                "common": "Украина"
            },
            "spa": {
                "official": "Ucrania",
                "common": "Ucrania"
            },
            "fin": {
                "official": "Ukraina",
                "common": "Ukraina"
            }
        },
        "latlng": [
            49,
            32
        ],
        "demonym": "Ukrainian",
        "landlocked": false,
        "borders": [
            "BLR",
            "HUN",
            "MDA",
            "POL",
            "ROU",
            "RUS",
            "SVK"
        ],
        "area": 603500
    },
    {
        "name": {
            "common": "United States Minor Outlying Islands",
            "official": "United States Minor Outlying Islands",
            "native": {
                "eng": {
                    "official": "United States Minor Outlying Islands",
                    "common": "United States Minor Outlying Islands"
                }
            }
        },
        "tld": [
            ".us"
        ],
        "cca2": "UM",
        "ccn3": "581",
        "cca3": "UMI",
        "cioc": "",
        "currency": [
            "USD"
        ],
        "callingCode": [],
        "capital": "",
        "altSpellings": [
            "UM"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "USA, kleinere ausgelagerte Inseln",
                "common": "Kleinere Inselbesitzungen der Vereinigten Staaten"
            },
            "fra": {
                "official": "Îles mineures éloignées des États-Unis",
                "common": "Îles mineures éloignées des États-Unis"
            },
            "hrv": {
                "official": "Mali udaljeni otoci SAD-a",
                "common": "Mali udaljeni otoci SAD-a"
            },
            "ita": {
                "official": "Stati Uniti Isole Minori",
                "common": "Isole minori esterne degli Stati Uniti d'America"
            },
            "jpn": {
                "official": "アメリカ合衆国外諸島",
                "common": "合衆国領有小離島"
            },
            "nld": {
                "official": "Kleine afgelegen eilanden van de Verenigde Staten",
                "common": "Kleine afgelegen eilanden van de Verenigde Staten"
            },
            "por": {
                "official": "Estados Unidos Ilhas Menores Distantes",
                "common": "Ilhas Menores Distantes dos Estados Unidos"
            },
            "rus": {
                "official": "Внешние малые острова США",
                "common": "Внешние малые острова США"
            },
            "spa": {
                "official": "Estados Unidos Islas menores alejadas de",
                "common": "Islas Ultramarinas Menores de Estados Unidos"
            },
            "fin": {
                "official": "Yhdysvaltain asumattomat saaret",
                "common": "Yhdysvaltain asumattomat saaret"
            }
        },
        "latlng": [],
        "demonym": "American",
        "landlocked": false,
        "borders": [],
        "area": 34.2
    },
    {
        "name": {
            "common": "Uruguay",
            "official": "Oriental Republic of Uruguay",
            "native": {
                "spa": {
                    "official": "República Oriental del Uruguay",
                    "common": "Uruguay"
                }
            }
        },
        "tld": [
            ".uy"
        ],
        "cca2": "UY",
        "ccn3": "858",
        "cca3": "URY",
        "cioc": "URU",
        "currency": [
            "UYI",
            "UYU"
        ],
        "callingCode": [
            "598"
        ],
        "capital": "Montevideo",
        "altSpellings": [
            "UY",
            "Oriental Republic of Uruguay",
            "República Oriental del Uruguay"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Republik Östlich des Uruguay",
                "common": "Uruguay"
            },
            "fra": {
                "official": "République orientale de l'Uruguay",
                "common": "Uruguay"
            },
            "hrv": {
                "official": "Orijentalna Republika Urugvaj",
                "common": "Urugvaj"
            },
            "ita": {
                "official": "Repubblica Orientale dell'Uruguay",
                "common": "Uruguay"
            },
            "jpn": {
                "official": "ウルグアイ東方共和国",
                "common": "ウルグアイ"
            },
            "nld": {
                "official": "Oosterse Republiek Uruguay",
                "common": "Uruguay"
            },
            "por": {
                "official": "República Oriental do Uruguai",
                "common": "Uruguai"
            },
            "rus": {
                "official": "Восточной Республики Уругвай",
                "common": "Уругвай"
            },
            "spa": {
                "official": "República Oriental del Uruguay",
                "common": "Uruguay"
            },
            "fin": {
                "official": "Uruguayn itäinen tasavalta",
                "common": "Uruguay"
            }
        },
        "latlng": [
            -33,
            -56
        ],
        "demonym": "Uruguayan",
        "landlocked": false,
        "borders": [
            "ARG",
            "BRA"
        ],
        "area": 181034
    },
    {
        "name": {
            "common": "United States",
            "official": "United States of America",
            "native": {
                "eng": {
                    "official": "United States of America",
                    "common": "United States"
                }
            }
        },
        "tld": [
            ".us"
        ],
        "cca2": "US",
        "ccn3": "840",
        "cca3": "USA",
        "cioc": "USA",
        "currency": [
            "USD",
            "USN",
            "USS"
        ],
        "callingCode": [
            "1"
        ],
        "capital": "Washington D.C.",
        "altSpellings": [
            "US",
            "USA",
            "United States of America"
        ],
        "region": "Americas",
        "subregion": "Northern America",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Vereinigte Staaten von Amerika",
                "common": "Vereinigte Staaten von Amerika"
            },
            "fra": {
                "official": "Les états-unis d'Amérique",
                "common": "États-Unis"
            },
            "hrv": {
                "official": "Sjedinjene Države Amerike",
                "common": "Sjedinjene Američke Države"
            },
            "ita": {
                "official": "Stati Uniti d'America",
                "common": "Stati Uniti d'America"
            },
            "jpn": {
                "official": "アメリカ合衆国",
                "common": "アメリカ合衆国"
            },
            "nld": {
                "official": "Verenigde Staten van Amerika",
                "common": "Verenigde Staten"
            },
            "por": {
                "official": "Estados Unidos da América",
                "common": "Estados Unidos"
            },
            "rus": {
                "official": "Соединенные Штаты Америки",
                "common": "Соединённые Штаты Америки"
            },
            "spa": {
                "official": "Estados Unidos de América",
                "common": "Estados Unidos"
            },
            "fin": {
                "official": "Amerikan yhdysvallat",
                "common": "Yhdysvallat"
            }
        },
        "latlng": [
            38,
            -97
        ],
        "demonym": "American",
        "landlocked": false,
        "borders": [
            "CAN",
            "MEX"
        ],
        "area": 9372610
    },
    {
        "name": {
            "common": "Uzbekistan",
            "official": "Republic of Uzbekistan",
            "native": {
                "rus": {
                    "official": "Республика Узбекистан",
                    "common": "Узбекистан"
                },
                "uzb": {
                    "official": "O'zbekiston Respublikasi",
                    "common": "O‘zbekiston"
                }
            }
        },
        "tld": [
            ".uz"
        ],
        "cca2": "UZ",
        "ccn3": "860",
        "cca3": "UZB",
        "cioc": "UZB",
        "currency": [
            "UZS"
        ],
        "callingCode": [
            "998"
        ],
        "capital": "Tashkent",
        "altSpellings": [
            "UZ",
            "Republic of Uzbekistan",
            "O‘zbekiston Respublikasi",
            "Ўзбекистон Республикаси"
        ],
        "region": "Asia",
        "subregion": "Central Asia",
        "languages": {
            "rus": "Russian",
            "uzb": "Uzbek"
        },
        "translations": {
            "deu": {
                "official": "Republik Usbekistan",
                "common": "Usbekistan"
            },
            "fra": {
                "official": "République d'Ouzbékistan",
                "common": "Ouzbékistan"
            },
            "hrv": {
                "official": "Republika Uzbekistan",
                "common": "Uzbekistan"
            },
            "ita": {
                "official": "Repubblica di Uzbekistan",
                "common": "Uzbekistan"
            },
            "jpn": {
                "official": "ウズベキスタン共和国",
                "common": "ウズベキスタン"
            },
            "nld": {
                "official": "Republiek Oezbekistan",
                "common": "Oezbekistan"
            },
            "por": {
                "official": "República do Usbequistão",
                "common": "Uzbequistão"
            },
            "rus": {
                "official": "Республика Узбекистан",
                "common": "Узбекистан"
            },
            "spa": {
                "official": "República de Uzbekistán",
                "common": "Uzbekistán"
            },
            "fin": {
                "official": "Uzbekistanin tasavalta",
                "common": "Uzbekistan"
            }
        },
        "latlng": [
            41,
            64
        ],
        "demonym": "Uzbekistani",
        "landlocked": true,
        "borders": [
            "AFG",
            "KAZ",
            "KGZ",
            "TJK",
            "TKM"
        ],
        "area": 447400
    },
    {
        "name": {
            "common": "Vatican City",
            "official": "Vatican City State",
            "native": {
                "ita": {
                    "official": "Stato della Città del Vaticano",
                    "common": "Vaticano"
                },
                "lat": {
                    "official": "Status Civitatis Vaticanæ",
                    "common": "Vaticanæ"
                }
            }
        },
        "tld": [
            ".va"
        ],
        "cca2": "VA",
        "ccn3": "336",
        "cca3": "VAT",
        "cioc": "",
        "currency": [
            "EUR"
        ],
        "callingCode": [
            "3906698",
            "379"
        ],
        "capital": "Vatican City",
        "altSpellings": [
            "VA",
            "Holy See (Vatican City State)",
            "Vatican City State",
            "Stato della Città del Vaticano"
        ],
        "region": "Europe",
        "subregion": "Southern Europe",
        "languages": {
            "ita": "Italian",
            "lat": "Latin"
        },
        "translations": {
            "deu": {
                "official": "Staat Vatikanstadt",
                "common": "Vatikanstadt"
            },
            "fra": {
                "official": "Cité du Vatican",
                "common": "Cité du Vatican"
            },
            "hrv": {
                "official": "Vatikan",
                "common": "Vatikan"
            },
            "ita": {
                "official": "Città del Vaticano",
                "common": "Città del Vaticano"
            },
            "jpn": {
                "official": "バチカン市国の状態",
                "common": "バチカン市国"
            },
            "nld": {
                "official": "Vaticaanstad",
                "common": "Vaticaanstad"
            },
            "por": {
                "official": "Cidade do Vaticano",
                "common": "Cidade do Vaticano"
            },
            "rus": {
                "official": "Город-государство Ватикан",
                "common": "Ватикан"
            },
            "spa": {
                "official": "Ciudad del Vaticano",
                "common": "Ciudad del Vaticano"
            },
            "fin": {
                "official": "Vatikaanin kaupunkivaltio",
                "common": "Vatikaani"
            }
        },
        "latlng": [
            41.9,
            12.45
        ],
        "demonym": "Italian",
        "landlocked": true,
        "borders": [
            "ITA"
        ],
        "area": 0.44
    },
    {
        "name": {
            "common": "Saint Vincent and the Grenadines",
            "official": "Saint Vincent and the Grenadines",
            "native": {
                "eng": {
                    "official": "Saint Vincent and the Grenadines",
                    "common": "Saint Vincent and the Grenadines"
                }
            }
        },
        "tld": [
            ".vc"
        ],
        "cca2": "VC",
        "ccn3": "670",
        "cca3": "VCT",
        "cioc": "VIN",
        "currency": [
            "XCD"
        ],
        "callingCode": [
            "1784"
        ],
        "capital": "Kingstown",
        "altSpellings": [
            "VC"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "St. Vincent und die Grenadinen",
                "common": "Saint Vincent und die Grenadinen"
            },
            "fra": {
                "official": "Saint-Vincent-et-les Grenadines",
                "common": "Saint-Vincent-et-les-Grenadines"
            },
            "hrv": {
                "official": "Sveti Vincent i Grenadini",
                "common": "Sveti Vincent i Grenadini"
            },
            "ita": {
                "official": "Saint Vincent e Grenadine",
                "common": "Saint Vincent e Grenadine"
            },
            "jpn": {
                "official": "セントビンセントおよびグレナディーン諸島",
                "common": "セントビンセントおよびグレナディーン諸島"
            },
            "nld": {
                "official": "Saint Vincent en de Grenadines",
                "common": "Saint Vincent en de Grenadines"
            },
            "por": {
                "official": "São Vicente e Granadinas",
                "common": "São Vincente e Granadinas"
            },
            "rus": {
                "official": "Сент-Винсент и Гренадины",
                "common": "Сент-Винсент и Гренадины"
            },
            "spa": {
                "official": "San Vicente y las Granadinas",
                "common": "San Vicente y Granadinas"
            },
            "fin": {
                "official": "Saint Vincent ja Grenadiinit",
                "common": "Saint Vincent ja Grenadiinit"
            }
        },
        "latlng": [
            13.25,
            -61.2
        ],
        "demonym": "Saint Vincentian",
        "landlocked": false,
        "borders": [],
        "area": 389
    },
    {
        "name": {
            "common": "Venezuela",
            "official": "Bolivarian Republic of Venezuela",
            "native": {
                "spa": {
                    "official": "República Bolivariana de Venezuela",
                    "common": "Venezuela"
                }
            }
        },
        "tld": [
            ".ve"
        ],
        "cca2": "VE",
        "ccn3": "862",
        "cca3": "VEN",
        "cioc": "VEN",
        "currency": [
            "VEF"
        ],
        "callingCode": [
            "58"
        ],
        "capital": "Caracas",
        "altSpellings": [
            "VE",
            "Bolivarian Republic of Venezuela",
            "Venezuela, Bolivarian Republic of",
            "República Bolivariana de Venezuela"
        ],
        "region": "Americas",
        "subregion": "South America",
        "languages": {
            "spa": "Spanish"
        },
        "translations": {
            "deu": {
                "official": "Bolivarische Republik Venezuela",
                "common": "Venezuela"
            },
            "fra": {
                "official": "République bolivarienne du Venezuela",
                "common": "Venezuela"
            },
            "hrv": {
                "official": "BOLIVARIJANSKA Republika Venezuela",
                "common": "Venezuela"
            },
            "ita": {
                "official": "Repubblica Bolivariana del Venezuela",
                "common": "Venezuela"
            },
            "jpn": {
                "official": "ベネズエラ·ボリバル共和国",
                "common": "ベネズエラ・ボリバル共和国"
            },
            "nld": {
                "official": "Bolivariaanse Republiek Venezuela",
                "common": "Venezuela"
            },
            "por": {
                "official": "República Bolivariana da Venezuela",
                "common": "Venezuela"
            },
            "rus": {
                "official": "Боливарианская Республика Венесуэла",
                "common": "Венесуэла"
            },
            "spa": {
                "official": "República Bolivariana de Venezuela",
                "common": "Venezuela"
            },
            "fin": {
                "official": "Venezuelan bolivariaainen tasavalta",
                "common": "Venezuela"
            }
        },
        "latlng": [
            8,
            -66
        ],
        "demonym": "Venezuelan",
        "landlocked": false,
        "borders": [
            "BRA",
            "COL",
            "GUY"
        ],
        "area": 916445
    },
    {
        "name": {
            "common": "British Virgin Islands",
            "official": "Virgin Islands",
            "native": {
                "eng": {
                    "official": "Virgin Islands",
                    "common": "British Virgin Islands"
                }
            }
        },
        "tld": [
            ".vg"
        ],
        "cca2": "VG",
        "ccn3": "092",
        "cca3": "VGB",
        "cioc": "IVB",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1284"
        ],
        "capital": "Road Town",
        "altSpellings": [
            "VG",
            "Virgin Islands, British"
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Jungferninseln",
                "common": "Britische Jungferninseln"
            },
            "fra": {
                "official": "îles Vierges",
                "common": "Îles Vierges britanniques"
            },
            "hrv": {
                "official": "Djevičanski Otoci",
                "common": "Britanski Djevičanski Otoci"
            },
            "ita": {
                "official": "Isole Vergini",
                "common": "Isole Vergini Britanniche"
            },
            "jpn": {
                "official": "バージン諸島",
                "common": "イギリス領ヴァージン諸島"
            },
            "nld": {
                "official": "Maagdeneilanden",
                "common": "Britse Maagdeneilanden"
            },
            "por": {
                "official": "Ilhas Virgens",
                "common": "Ilhas Virgens"
            },
            "rus": {
                "official": "Виргинские острова",
                "common": "Британские Виргинские острова"
            },
            "spa": {
                "official": "Islas Vírgenes",
                "common": "Islas Vírgenes del Reino Unido"
            },
            "fin": {
                "official": "Brittiläiset Neitsytsaaret",
                "common": "Neitsytsaaret"
            }
        },
        "latlng": [
            18.431383,
            -64.62305
        ],
        "demonym": "Virgin Islander",
        "landlocked": false,
        "borders": [],
        "area": 151
    },
    {
        "name": {
            "common": "United States Virgin Islands",
            "official": "Virgin Islands of the United States",
            "native": {
                "eng": {
                    "official": "Virgin Islands of the United States",
                    "common": "United States Virgin Islands"
                }
            }
        },
        "tld": [
            ".vi"
        ],
        "cca2": "VI",
        "ccn3": "850",
        "cca3": "VIR",
        "cioc": "ISV",
        "currency": [
            "USD"
        ],
        "callingCode": [
            "1340"
        ],
        "capital": "Charlotte Amalie",
        "altSpellings": [
            "VI",
            "Virgin Islands, U.S."
        ],
        "region": "Americas",
        "subregion": "Caribbean",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Jungferninseln der Vereinigten Staaten",
                "common": "Amerikanische Jungferninseln"
            },
            "fra": {
                "official": "Îles Vierges des États-Unis",
                "common": "Îles Vierges des États-Unis"
            },
            "hrv": {
                "official": "Djevičanski Otoci SAD",
                "common": "Američki Djevičanski Otoci"
            },
            "ita": {
                "official": "Isole Vergini degli Stati Uniti",
                "common": "Isole Vergini americane"
            },
            "jpn": {
                "official": "米国のバージン諸島",
                "common": "アメリカ領ヴァージン諸島"
            },
            "nld": {
                "official": "Maagdeneilanden van de Verenigde Staten",
                "common": "Amerikaanse Maagdeneilanden"
            },
            "por": {
                "official": "Ilhas Virgens dos Estados Unidos",
                "common": "Ilhas Virgens dos Estados Unidos"
            },
            "rus": {
                "official": "Виргинские острова Соединенных Штатов",
                "common": "Виргинские Острова"
            },
            "spa": {
                "official": "Islas Vírgenes de los Estados Unidos",
                "common": "Islas Vírgenes de los Estados Unidos"
            },
            "fin": {
                "official": "Yhdysvaltain Neitsytsaaret",
                "common": "Neitsytsaaret"
            }
        },
        "latlng": [
            18.35,
            -64.933333
        ],
        "demonym": "Virgin Islander",
        "landlocked": false,
        "borders": [],
        "area": 347
    },
    {
        "name": {
            "common": "Vietnam",
            "official": "Socialist Republic of Vietnam",
            "native": {
                "vie": {
                    "official": "Cộng hòa xã hội chủ nghĩa Việt Nam",
                    "common": "Việt Nam"
                }
            }
        },
        "tld": [
            ".vn"
        ],
        "cca2": "VN",
        "ccn3": "704",
        "cca3": "VNM",
        "cioc": "VIE",
        "currency": [
            "VND"
        ],
        "callingCode": [
            "84"
        ],
        "capital": "Hanoi",
        "altSpellings": [
            "VN",
            "Socialist Republic of Vietnam",
            "Cộng hòa Xã hội chủ nghĩa Việt Nam",
            "Viet Nam"
        ],
        "region": "Asia",
        "subregion": "South-Eastern Asia",
        "languages": {
            "vie": "Vietnamese"
        },
        "translations": {
            "deu": {
                "official": "Sozialistische Republik Vietnam",
                "common": "Vietnam"
            },
            "fra": {
                "official": "République socialiste du Viêt Nam",
                "common": "Viêt Nam"
            },
            "hrv": {
                "official": "Socijalistička Republika Vijetnam",
                "common": "Vijetnam"
            },
            "ita": {
                "official": "Repubblica socialista del Vietnam",
                "common": "Vietnam"
            },
            "jpn": {
                "official": "ベトナム社会主義共和国",
                "common": "ベトナム"
            },
            "nld": {
                "official": "Socialistische Republiek Vietnam",
                "common": "Vietnam"
            },
            "por": {
                "official": "República Socialista do Vietname",
                "common": "Vietname"
            },
            "rus": {
                "official": "Социалистическая Республика Вьетнам",
                "common": "Вьетнам"
            },
            "spa": {
                "official": "República Socialista de Vietnam",
                "common": "Vietnam"
            },
            "fin": {
                "official": "Vietnamin sosialistinen tasavalta",
                "common": "Vietnam"
            }
        },
        "latlng": [
            16.16666666,
            107.83333333
        ],
        "demonym": "Vietnamese",
        "landlocked": false,
        "borders": [
            "KHM",
            "CHN",
            "LAO"
        ],
        "area": 331212
    },
    {
        "name": {
            "common": "Vanuatu",
            "official": "Republic of Vanuatu",
            "native": {
                "bis": {
                    "official": "Ripablik blong Vanuatu",
                    "common": "Vanuatu"
                },
                "eng": {
                    "official": "Republic of Vanuatu",
                    "common": "Vanuatu"
                },
                "fra": {
                    "official": "République de Vanuatu",
                    "common": "Vanuatu"
                }
            }
        },
        "tld": [
            ".vu"
        ],
        "cca2": "VU",
        "ccn3": "548",
        "cca3": "VUT",
        "cioc": "VAN",
        "currency": [
            "VUV"
        ],
        "callingCode": [
            "678"
        ],
        "capital": "Port Vila",
        "altSpellings": [
            "VU",
            "Republic of Vanuatu",
            "Ripablik blong Vanuatu",
            "République de Vanuatu"
        ],
        "region": "Oceania",
        "subregion": "Melanesia",
        "languages": {
            "bis": "Bislama",
            "eng": "English",
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Vanuatu",
                "common": "Vanuatu"
            },
            "fra": {
                "official": "République de Vanuatu",
                "common": "Vanuatu"
            },
            "hrv": {
                "official": "Republika Vanuatu",
                "common": "Vanuatu"
            },
            "ita": {
                "official": "Repubblica di Vanuatu",
                "common": "Vanuatu"
            },
            "jpn": {
                "official": "バヌアツ共和国",
                "common": "バヌアツ"
            },
            "nld": {
                "official": "Republiek Vanuatu",
                "common": "Vanuatu"
            },
            "por": {
                "official": "República de Vanuatu",
                "common": "Vanuatu"
            },
            "rus": {
                "official": "Республика Вануату",
                "common": "Вануату"
            },
            "spa": {
                "official": "República de Vanuatu",
                "common": "Vanuatu"
            },
            "fin": {
                "official": "Vanuatun tasavalta",
                "common": "Vanuatu"
            }
        },
        "latlng": [
            -16,
            167
        ],
        "demonym": "Ni-Vanuatu",
        "landlocked": false,
        "borders": [],
        "area": 12189
    },
    {
        "name": {
            "common": "Wallis and Futuna",
            "official": "Territory of the Wallis and Futuna Islands",
            "native": {
                "fra": {
                    "official": "Territoire des îles Wallis et Futuna",
                    "common": "Wallis et Futuna"
                }
            }
        },
        "tld": [
            ".wf"
        ],
        "cca2": "WF",
        "ccn3": "876",
        "cca3": "WLF",
        "cioc": "",
        "currency": [
            "XPF"
        ],
        "callingCode": [
            "681"
        ],
        "capital": "Mata-Utu",
        "altSpellings": [
            "WF",
            "Territory of the Wallis and Futuna Islands",
            "Territoire des îles Wallis et Futuna"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "fra": "French"
        },
        "translations": {
            "deu": {
                "official": "Gebiet der Wallis und Futuna",
                "common": "Wallis und Futuna"
            },
            "fra": {
                "official": "Territoire des îles Wallis et Futuna",
                "common": "Wallis-et-Futuna"
            },
            "hrv": {
                "official": "Teritoriju Wallis i Futuna",
                "common": "Wallis i Fortuna"
            },
            "ita": {
                "official": "Territorio delle Isole Wallis e Futuna",
                "common": "Wallis e Futuna"
            },
            "jpn": {
                "official": "ウォリス·フツナ諸島の領土",
                "common": "ウォリス・フツナ"
            },
            "nld": {
                "official": "Grondgebied van de Wallis en Futuna",
                "common": "Wallis en Futuna"
            },
            "por": {
                "official": "Território das Ilhas Wallis e Futuna",
                "common": "Wallis e Futuna"
            },
            "rus": {
                "official": "Территория Уоллис и Футуна острова",
                "common": "Уоллис и Футуна"
            },
            "spa": {
                "official": "Territorio de las Islas Wallis y Futuna",
                "common": "Wallis y Futuna"
            },
            "fin": {
                "official": "Wallisin ja Futunan yhteisö",
                "common": "Wallis ja Futuna"
            }
        },
        "latlng": [
            -13.3,
            -176.2
        ],
        "demonym": "Wallis and Futuna Islander",
        "landlocked": false,
        "borders": [],
        "area": 142
    },
    {
        "name": {
            "common": "Samoa",
            "official": "Independent State of Samoa",
            "native": {
                "eng": {
                    "official": "Independent State of Samoa",
                    "common": "Samoa"
                },
                "smo": {
                    "official": "Malo Saʻoloto Tutoʻatasi o Sāmoa",
                    "common": "Sāmoa"
                }
            }
        },
        "tld": [
            ".ws"
        ],
        "cca2": "WS",
        "ccn3": "882",
        "cca3": "WSM",
        "cioc": "SAM",
        "currency": [
            "WST"
        ],
        "callingCode": [
            "685"
        ],
        "capital": "Apia",
        "altSpellings": [
            "WS",
            "Independent State of Samoa",
            "Malo Saʻoloto Tutoʻatasi o Sāmoa"
        ],
        "region": "Oceania",
        "subregion": "Polynesia",
        "languages": {
            "eng": "English",
            "smo": "Samoan"
        },
        "translations": {
            "deu": {
                "official": "Unabhängige Staat Samoa",
                "common": "Samoa"
            },
            "fra": {
                "official": "Samoa",
                "common": "Samoa"
            },
            "hrv": {
                "official": "Nezavisna Država Samoa",
                "common": "Samoa"
            },
            "ita": {
                "official": "Stato indipendente di Samoa",
                "common": "Samoa"
            },
            "jpn": {
                "official": "サモア独立国",
                "common": "サモア"
            },
            "nld": {
                "official": "Onafhankelijke Staat Samoa",
                "common": "Samoa"
            },
            "por": {
                "official": "Estado Independente de Samoa",
                "common": "Samoa"
            },
            "rus": {
                "official": "Независимое Государство Самоа",
                "common": "Самоа"
            },
            "spa": {
                "official": "Estado Independiente de Samoa",
                "common": "Samoa"
            },
            "fin": {
                "official": "Samoan itsenäinen valtio",
                "common": "Samoa"
            }
        },
        "latlng": [
            -13.58333333,
            -172.33333333
        ],
        "demonym": "Samoan",
        "landlocked": false,
        "borders": [],
        "area": 2842
    },
    {
        "name": {
            "common": "Yemen",
            "official": "Republic of Yemen",
            "native": {
                "ara": {
                    "official": "الجمهورية اليمنية",
                    "common": "اليَمَن"
                }
            }
        },
        "tld": [
            ".ye"
        ],
        "cca2": "YE",
        "ccn3": "887",
        "cca3": "YEM",
        "cioc": "YEM",
        "currency": [
            "YER"
        ],
        "callingCode": [
            "967"
        ],
        "capital": "Sana'a",
        "altSpellings": [
            "YE",
            "Yemeni Republic",
            "al-Jumhūriyyah al-Yamaniyyah"
        ],
        "region": "Asia",
        "subregion": "Western Asia",
        "languages": {
            "ara": "Arabic"
        },
        "translations": {
            "deu": {
                "official": "Republik Jemen",
                "common": "Jemen"
            },
            "fra": {
                "official": "République du Yémen",
                "common": "Yémen"
            },
            "hrv": {
                "official": "Republika Jemen",
                "common": "Jemen"
            },
            "ita": {
                "official": "Repubblica dello Yemen",
                "common": "Yemen"
            },
            "jpn": {
                "official": "イエメン共和国",
                "common": "イエメン"
            },
            "nld": {
                "official": "Republiek Jemen",
                "common": "Jemen"
            },
            "por": {
                "official": "República do Iêmen",
                "common": "Iémen"
            },
            "rus": {
                "official": "Йеменская Республика",
                "common": "Йемен"
            },
            "spa": {
                "official": "República de Yemen",
                "common": "Yemen"
            },
            "fin": {
                "official": "Jemenin tasavalta",
                "common": "Jemen"
            }
        },
        "latlng": [
            15,
            48
        ],
        "demonym": "Yemeni",
        "landlocked": false,
        "borders": [
            "OMN",
            "SAU"
        ],
        "area": 527968
    },
    {
        "name": {
            "common": "South Africa",
            "official": "Republic of South Africa",
            "native": {
                "afr": {
                    "official": "Republiek van Suid-Afrika",
                    "common": "South Africa"
                },
                "eng": {
                    "official": "Republic of South Africa",
                    "common": "South Africa"
                },
                "nbl": {
                    "official": "IRiphabliki yeSewula Afrika",
                    "common": "Sewula Afrika"
                },
                "nso": {
                    "official": "Rephaboliki ya Afrika-Borwa ",
                    "common": "Afrika-Borwa"
                },
                "sot": {
                    "official": "Rephaboliki ya Afrika Borwa",
                    "common": "Afrika Borwa"
                },
                "ssw": {
                    "official": "IRiphabhulikhi yeNingizimu Afrika",
                    "common": "Ningizimu Afrika"
                },
                "tsn": {
                    "official": "Rephaboliki ya Aforika Borwa",
                    "common": "Aforika Borwa"
                },
                "tso": {
                    "official": "Riphabliki ra Afrika Dzonga",
                    "common": "Afrika Dzonga"
                },
                "ven": {
                    "official": "Riphabuḽiki ya Afurika Tshipembe",
                    "common": "Afurika Tshipembe"
                },
                "xho": {
                    "official": "IRiphabliki yaseMzantsi Afrika",
                    "common": "Mzantsi Afrika"
                },
                "zul": {
                    "official": "IRiphabliki yaseNingizimu Afrika",
                    "common": "Ningizimu Afrika"
                }
            }
        },
        "tld": [
            ".za"
        ],
        "cca2": "ZA",
        "ccn3": "710",
        "cca3": "ZAF",
        "cioc": "RSA",
        "currency": [
            "ZAR"
        ],
        "callingCode": [
            "27"
        ],
        "capital": "Pretoria",
        "altSpellings": [
            "ZA",
            "RSA",
            "Suid-Afrika",
            "Republic of South Africa"
        ],
        "region": "Africa",
        "subregion": "Southern Africa",
        "languages": {
            "afr": "Afrikaans",
            "eng": "English",
            "nbl": "Southern Ndebele",
            "nso": "Northern Sotho",
            "sot": "Southern Sotho",
            "ssw": "Swazi",
            "tsn": "Tswana",
            "tso": "Tsonga",
            "ven": "Venda",
            "xho": "Xhosa",
            "zul": "Zulu"
        },
        "translations": {
            "deu": {
                "official": "Republik Südafrika",
                "common": "Republik Südafrika"
            },
            "fra": {
                "official": "République d'Afrique du Sud",
                "common": "Afrique du Sud"
            },
            "hrv": {
                "official": "Južnoafrička Republika",
                "common": "Južnoafrička Republika"
            },
            "ita": {
                "official": "Repubblica del Sud Africa",
                "common": "Sud Africa"
            },
            "jpn": {
                "official": "南アフリカ共和国",
                "common": "南アフリカ"
            },
            "nld": {
                "official": "Zuid -Afrika",
                "common": "Zuid-Afrika"
            },
            "por": {
                "official": "República da África do Sul",
                "common": "África do Sul"
            },
            "rus": {
                "official": "Южно-Африканская Республика",
                "common": "Южно-Африканская Республика"
            },
            "spa": {
                "official": "República de Sudáfrica",
                "common": "República de Sudáfrica"
            },
            "fin": {
                "official": "Etelä-Afrikan tasavalta",
                "common": "Etelä-Afrikka"
            }
        },
        "latlng": [
            -29,
            24
        ],
        "demonym": "South African",
        "landlocked": false,
        "borders": [
            "BWA",
            "LSO",
            "MOZ",
            "NAM",
            "SWZ",
            "ZWE"
        ],
        "area": 1221037
    },
    {
        "name": {
            "common": "Zambia",
            "official": "Republic of Zambia",
            "native": {
                "eng": {
                    "official": "Republic of Zambia",
                    "common": "Zambia"
                }
            }
        },
        "tld": [
            ".zm"
        ],
        "cca2": "ZM",
        "ccn3": "894",
        "cca3": "ZMB",
        "cioc": "ZAM",
        "currency": [
            "ZMW"
        ],
        "callingCode": [
            "260"
        ],
        "capital": "Lusaka",
        "altSpellings": [
            "ZM",
            "Republic of Zambia"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "eng": "English"
        },
        "translations": {
            "deu": {
                "official": "Republik Sambia",
                "common": "Sambia"
            },
            "fra": {
                "official": "République de Zambie",
                "common": "Zambie"
            },
            "hrv": {
                "official": "Republika Zambija",
                "common": "Zambija"
            },
            "ita": {
                "official": "Repubblica di Zambia",
                "common": "Zambia"
            },
            "jpn": {
                "official": "ザンビア共和国",
                "common": "ザンビア"
            },
            "nld": {
                "official": "Republiek Zambia",
                "common": "Zambia"
            },
            "por": {
                "official": "República da Zâmbia",
                "common": "Zâmbia"
            },
            "rus": {
                "official": "Республика Замбия",
                "common": "Замбия"
            },
            "spa": {
                "official": "República de Zambia",
                "common": "Zambia"
            },
            "fin": {
                "official": "Sambian tasavalta",
                "common": "Sambia"
            }
        },
        "latlng": [
            -15,
            30
        ],
        "demonym": "Zambian",
        "landlocked": true,
        "borders": [
            "AGO",
            "BWA",
            "COD",
            "MWI",
            "MOZ",
            "NAM",
            "TZA",
            "ZWE"
        ],
        "area": 752612
    },
    {
        "name": {
            "common": "Zimbabwe",
            "official": "Republic of Zimbabwe",
            "native": {
                "bwg": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "eng": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "kck": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "khi": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "ndc": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "nde": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "nya": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "sna": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "sot": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "toi": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "tsn": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "tso": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "ven": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "xho": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                },
                "zib": {
                    "official": "Republic of Zimbabwe",
                    "common": "Zimbabwe"
                }
            }
        },
        "tld": [
            ".zw"
        ],
        "cca2": "ZW",
        "ccn3": "716",
        "cca3": "ZWE",
        "cioc": "ZIM",
        "currency": [
            "ZWL"
        ],
        "callingCode": [
            "263"
        ],
        "capital": "Harare",
        "altSpellings": [
            "ZW",
            "Republic of Zimbabwe"
        ],
        "region": "Africa",
        "subregion": "Eastern Africa",
        "languages": {
            "bwg": "Chibarwe",
            "eng": "English",
            "kck": "Kalanga",
            "khi": "Khoisan",
            "ndc": "Ndau",
            "nde": "Northern Ndebele",
            "nya": "Chewa",
            "sna": "Shona",
            "sot": "Sotho",
            "toi": "Tonga",
            "tsn": "Tswana",
            "tso": "Tsonga",
            "ven": "Venda",
            "xho": "Xhosa",
            "zib": "Zimbabwean Sign Language"
        },
        "translations": {
            "deu": {
                "official": "Republik Simbabwe",
                "common": "Simbabwe"
            },
            "fra": {
                "official": "République du Zimbabwe",
                "common": "Zimbabwe"
            },
            "hrv": {
                "official": "Republika Zimbabve",
                "common": "Zimbabve"
            },
            "ita": {
                "official": "Repubblica dello Zimbabwe",
                "common": "Zimbabwe"
            },
            "jpn": {
                "official": "ジンバブエ共和国",
                "common": "ジンバブエ"
            },
            "nld": {
                "official": "Republiek Zimbabwe",
                "common": "Zimbabwe"
            },
            "por": {
                "official": "República do Zimbabwe",
                "common": "Zimbabwe"
            },
            "rus": {
                "official": "Республика Зимбабве",
                "common": "Зимбабве"
            },
            "spa": {
                "official": "República de Zimbabue",
                "common": "Zimbabue"
            },
            "fin": {
                "official": "Zimbabwen tasavalta",
                "common": "Zimbabwe"
            }
        },
        "latlng": [
            -20,
            30
        ],
        "demonym": "Zimbabwean",
        "landlocked": true,
        "borders": [
            "BWA",
            "MOZ",
            "ZAF",
            "ZMB"
        ],
        "area": 390757
    }
]
`
