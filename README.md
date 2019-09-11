# phonecountry
phonecountry extract information about country from a phone number and vice versa.

__Phonecountry has the ability to distinguish the phone number between__ 

* USA, Canada and all caribian coutry. (All starting with +1)
* Russia and Kazakistan (Starting with +7)
* Italy and Vatican City (Starting with +39)
* UK, Guernsey, Isle of Man and Jersey (Starting with +44)
* And many more country with similar country code.

# Currently available APIs

1. __GetCountryNameFromPhone__ - GetCountryNameFromPhone returns country common name for given phone number.
Example +919445454528 => India

2. __GetCountryISO2Code__ - GetCountryISO2Code returns country iso2 code of given number. Example +919445454528 => IN

## Limitations

1. Input phone number need to be atleast 7 digit long.(Excluding + sign)
