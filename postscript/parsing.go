package postscript

import (
	"io/ioutil" // for ioutil.ReadFile
	"strconv"   // for strconv.Atoi
	"strings"   // for strings.Fields

	// where all out postscript objects are defined:
	"github.com/aznashwan/go-cg/postscript/objects"
)

// atois take a slice of string representations of numbers and applies atoi on
// each, returning the resulting list
// if any error occurs for a conversion; the function will promptly return it
func atois(strings []string) ([]uint, error) {
	uints := []uint{}

	for _, str := range strings {
		// attempt conversion to int
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		// cast to uint and append to result
		uints = append(uints, uint(n))
	}

	return uints, nil
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
			// fetch uints preceding "Line"
			uints, err := atois(items[i-4 : i])
			if err != nil {
				return nil, err
			}

			// create new line from fetched values
			line := objects.NewLine(
				objects.NewPoint(uints[0], uints[1]),
				objects.NewPoint(uints[2], uints[3]),
			)

			// add the newly created line to the list of returned lines
			lines = append(lines, line)
		}
	}

	return lines, nil
}
