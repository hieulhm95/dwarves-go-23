package utils

import (
	"fmt"
	"sort"
	"strconv"
)

type DataType struct {
	IntValue    []int
	FloatValue  []float64
	StringValue []string
}

func Sort(data []string, dataType string) (interface{}, error) {
	parsedData, err := parseDataValue(data, dataType)
	if err != nil {
		return nil, err
	}
	var tmp interface{}
	switch dataType {
	case "int":
		sort.Ints(parsedData.IntValue)
		tmp = parsedData.IntValue
	case "float":
		sort.Float64s(parsedData.FloatValue)
		tmp = parsedData.FloatValue
	case "string":
		sort.Strings(parsedData.StringValue)
		tmp = parsedData.StringValue
	}
	fmt.Println("Result: ", tmp)
	return tmp, nil
}

func parseDataValue(data []string, dataType string) (*DataType, error) {
	var ints []int
	var floats []float64
	var strings []string

	for _, val := range data {
		switch dataType {
		case "int":
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			ints = append(ints, num)
		case "float":
			fl, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
			floats = append(floats, fl)
		default:
			strings = append(strings, val)
		}
	}
	return &DataType{
		IntValue:    ints,
		FloatValue:  floats,
		StringValue: strings,
	}, nil
}
