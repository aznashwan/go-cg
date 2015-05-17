package twod

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../../postscript/objects"
)

// atofs take a slice of string representations of numbers and applies atof on
// each, returning the resulting list
// if any error occurs for a conversion; the function will promptly return it
func atofs(strings []string) ([]float64, error) {
	floats := []float64{}

	for _, str := range strings {
		// attempt conversion to int
		n, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing float: %s", err)
		}

		//  append to result
		floats = append(floats, n)
	}

	return floats, nil
}

// ParseFile parses the given .tsf file and returns all the 2D
// operations defined there.
func ParseFile(file string) ([][]interface{}, error) {
	// read the file's contents:
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %s", err)
	}

	operations := [][]interface{}{}
	items := strings.Fields(string(contents))

	// pass through all items and interpret all fields:
	for i, item := range items {
		switch item {
		case "t":
			floats, err := atofs(items[i+1 : i+3])
			if err != nil {
				return nil, err
			}
			operations = append(operations, []interface{}{
				"t", int(floats[0]), int(floats[1]),
			})
		case "s":
			floats, err := atofs(items[i+1 : i+5])
			if err != nil {
				return nil, err
			}
			operations = append(operations, []interface{}{
				"s", int(floats[0]), int(floats[1]), floats[2], floats[3],
			})
		case "r":
			floats, err := atofs(items[i+1 : i+4])
			if err != nil {
				return nil, err
			}
			operations = append(operations, []interface{}{
				"r", int(floats[0]), int(floats[1]), int(floats[2]),
			})
		}
	}

	return operations, nil
}

// applyOperationToLine applies the given operation to the the given Line.
func applyOperationToLine(l *objects.Line, op []interface{}) *objects.Line {
	switch op[0].(string) {
	case "t":
		return TranslateLine(l, op[1].(int), op[2].(int))
	case "s":
		return ScaleLineAroundPoint(
			l,
			objects.NewPoint(op[1].(int), op[2].(int)),
			op[3].(float64),
			op[4].(float64),
		)
	case "r":
		return RotateLineAroundPoint(
			l,
			objects.NewPoint(op[1].(int), op[2].(int)),
			op[3].(int),
		)
	}

	return nil
}

// ApplyTransformationsToLines applies all the given transformations
// to the given line and returns the result:
func ApplyTransformationsToLines(lines []*objects.Line, ops [][]interface{}) []*objects.Line {
	res := []*objects.Line{}

	for _, line := range lines {
		var l = line
		for _, op := range ops {
			l = applyOperationToLine(l, op)
		}
		res = append(res, l)
	}

	return res
}
