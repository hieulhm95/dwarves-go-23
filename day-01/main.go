package main

import (
	"fmt"
	"os"
)

func swap(firstName, lastName, country, middleName string) string {
	var formattedStr string
	switch country {
	case "VN":
		if middleName != "" {
			formattedStr = lastName + " " + middleName + " " + firstName
		} else {
			formattedStr = lastName + " " + firstName
		}
	case "US":
		if middleName != "" {
			formattedStr = firstName + " " + lastName + " " + middleName
		} else {
			formattedStr = firstName + " " + lastName
		}
	}
	return formattedStr
}

type Name struct {
	FirstName  string
	LastName   string
	MiddleName string
	Country    string
}

func parseArgFromOs(osInput []string) interface{} {
	country := osInput[len(osInput)-1]
	if country != "VN" && country != "US" {
		panic("Please input region")
	}
	var middleName string
	firstName := osInput[1]
	lastName := osInput[2]
	if len(osInput) > 4 {
		middleName = osInput[3]
	}
	return Name{firstName, lastName, middleName, country}
}

func main() {
	osInput := os.Args
	if len(osInput) < 4 {
		panic("Please input at least First name, Last name, Middle name (option) and Country")
	}
	nameObj := parseArgFromOs(osInput)
	name, ok := nameObj.(Name)
	if !ok {
		fmt.Println("Failed to parse the arguments")
		return
	}
	// Get firstName, lastName, and middleName
	firstName := name.FirstName
	lastName := name.LastName
	middleName := name.MiddleName
	country := name.Country
	result := swap(firstName, lastName, country, middleName)
	fmt.Println(result)
}
