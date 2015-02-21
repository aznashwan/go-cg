package xpm

import (
	"encoding/hex" // for hex.EncodeToString
	"fmt"          // for fmt.Sprintf
	"strings"      // for strings.ToUpper
)

// Color represents all the components an XPM color has
type Color struct {
	// the character combination the color will be encoded as
	chars string

	// the respective 8-bit degrees of each of the main colors
	red   byte
	green byte
	blue  byte
}

// encodeByte encodes the given byte to its respective 2 hexa characters and
// returns the corresponding 2-character string
func encodeByte(b byte) string {
	return strings.ToUpper(hex.EncodeToString([]byte{b}))
}

// Serialize returns the encoding of the current color in the classic XPM
// "CHAR c #0xR0xG0xB" format
func (c *Color) Serialize() string {
	return fmt.Sprintf("\"%s c #%s%s%s\"", c.chars, encodeByte(c.red), encodeByte(c.green), encodeByte(c.blue))
}
