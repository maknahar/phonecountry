package main

import (
	"fmt"
	"phonecountry"
)

func main() {
	country, err := phonecountry.GetCountryISO2Code("+949445454528")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Country: %s\n", country)
}
