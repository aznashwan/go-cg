package postscript

import (
	"io/ioutil" // for ioutil.ReadFile
	"strconv"   // for strconv.Atoi
	"strings"   // for strings.Fields

	// where all out postscript objects are defined:
	"./objects"
)

// atois take a slice of string representations of numbers and applies atoi on
// each, returning the resulting list
// if any error occurs for a conversion; the function will promptly return it
func atois(strings []string) ([]int, error) {
	ints := []int{}

	for _, str := range strings {
		// attempt conversion to int
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		// cast to int and append to result
		ints = append(ints, int(n))
	}

	return ints, nil
}

// ParseFile parses a postscript file and returns a slice of all Line objects
// defined in the file
// NOTE: only accepts files of this form:
//
// example.ps
//
// %%%BEGIN
// x11 y11 x12 y12 Line
// x21 y21 x22 y22 Line
// ...
// xn1 yn1 xn2 yn2 Line
// %%%END
//
// Newlines and anything above and below the BEGIN and END tags is ignored
func ParseFile(filename string) ([]*objects.Line, error) {
	// read out the file's contents
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// a slice in which to store our parsed lines
	lines := []*objects.Line{}

	// convert raw bytes to string and split into fields
	items := strings.Fields(string(contents))

	// pass through all items and interpret all fields preceding "Line"
	for i, item := range items {
		if item == "Line" {
			// fetch ints preceding "Line"
			ints, err := atois(items[i-4 : i])
			if err != nil {
				return nil, err
			}

			// create new line from fetched values
			line := objects.NewLine(
				objects.NewPoint(ints[0], ints[1]),
				objects.NewPoint(ints[2], ints[3]),
			)

			// add the newly created line to the list of returned lines
			lines = append(lines, line)
		}
	}

	return lines, nil
}
