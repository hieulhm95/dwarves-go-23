package main

import (
	"day-02/utils"
	"flag"
)

var isIntValue bool
var isStringValue bool
var isFloatValue bool
var isMixValue bool

func init() {
	flag.BoolVar(&isIntValue, "int", false, "An integer value")
	flag.BoolVar(&isStringValue, "string", false, "An string value")
	flag.BoolVar(&isFloatValue, "float", false, "An float value")
	flag.BoolVar(&isMixValue, "mix", false, "An mix value")
	flag.Parse()
}

func main() {
	args := flag.Args()
	var dataType string
	switch {
	case isIntValue:
		dataType = "int"
	case isFloatValue:
		dataType = "float"
	case isStringValue:
		dataType = "string"
	}
	utils.Sort(args, dataType)
}
