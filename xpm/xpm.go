package xpm

import (
	"fmt"       // for general formatting and fmt.Errorf
	"io/ioutil" // for ioutil.WriteFile
)

// XPM is an aggregate of all the data required for an XPM image
type XPM struct {
	// bitmap dimesions
	width  uint
	height uint

	// number of charaters per pixel
	cpp uint

	// slice of all the colors
	colors []Color

	// data is a map between each row number and a slice of strings represing
	// the row's contents
	data map[uint][]string
}

// NewXPM returns a new XPM object
func NewXPM(width, height, cpp uint, colors []Color) *XPM {
	var i uint

	// create new XPM struct
	xpm := &XPM{
		width:  width,
		height: height,
		cpp:    cpp,
		colors: colors,
		data:   make(map[uint][]string),
	}

	// fully initialize the data map
	for i = 0; i < height; i++ {
		xpm.data[i] = make([]string, width)
	}

	return xpm
}

// validatePixel validates the inputs given to SetPixel
func (xpm *XPM) validatePixel(x, y uint, cc string) error {
	if x > xpm.width {
		return fmt.Errorf("Invalid x=%d", x)
	}
	if y > xpm.height {
		return fmt.Errorf("Invalid y=%d", y)
	}

	for _, color := range xpm.colors {
		if color.chars == cc {
			return nil
		}
	}

	return fmt.Errorf("Nonexistent color combination %s in this XPM", cc)
}

// SetPixel sets a pixel to the given row, column, and character combination
// with respect to how the data matrix is represented in memory, not with
// respect to the cartesian system
// Returns an error if any of the given coordinates is out of range or if
// the color character combination has not been defined
func (xpm *XPM) SetPixel(x, y uint, cc string) error {
	if err := xpm.validatePixel(x, y, cc); err != nil {
		return err
	}

	xpm.data[x][y] = cc
	return nil
}

// SetPixelCartesian sets a pixel at the given 0-ordered left-handed cartesian
// coordinates x and y and with the given color character combination
// Returns an error if any of the given coordinates is out of range or if
// the color character combination has not been defined
func (xpm *XPM) SetPixelCartesian(x, y uint, cc string) error {
	return xpm.SetPixel(y, x, cc)
}

// AddColor adds a color to the associated XPM structure with the given
// red, green and blue values, as well as the character combination
// If a color with the given character combination has already been
// defined in the XPM, the function will return an error
func (xpm *XPM) AddColor(r, g, b byte, cc string) error {
	for _, color := range xpm.colors {
		if color.chars == cc {
			return fmt.Errorf("Color %q already defined!", cc)
		}
	}

	xpm.colors = append(xpm.colors, Color{cc, r, g, b})
	return nil
}

// Serialize returns the serialization of the particular XPM instance in
// the XPM format, for it to be ready to be printed to a file
// It will return an error if the XPM does not feature any colors
func (xpm *XPM) Serialize() ([]byte, error) {
	if len(xpm.colors) == 0 {
		return nil, fmt.Errorf("No colors included into this XPM!")
	}

	// add the XPM3 header
	res := "/* XPM */\n"
	res = res + "static char* XPM[] = {\n"

	// add initial params line (width, height, colors, chars/pixel)
	res = res + fmt.Sprintf("\"%d %d %d %d\",\n", xpm.width, xpm.height, len(xpm.colors), xpm.cpp)

	// add each color
	for _, color := range xpm.colors {
		res = res + color.Serialize() + ",\n"
	}

	// add each row
	for _, row := range xpm.data {
		res = res + "\""
		for _, col := range row {
			res = res + col
		}
		res = res + "\",\n"
	}

	// remove trailing comma and close brace
	res = res[:len(res)-2] + "\n}\n"

	return []byte(res), nil
}

// WriteToFile serializes the current XPM instance and writes it to the file
// given as parameter
// If the file does not exist, it will be created with default 0644 permissions
// If the file exists, it will be truncated
func (xpm *XPM) WriteToFile(filename string) error {
	// first serialize the XPM struct
	contents, err := xpm.Serialize()
	if err != nil {
		return err
	}

	// then write it to the file
	if err := ioutil.WriteFile(filename, contents, 0644); err != nil {
		return err
	}

	return nil
}
