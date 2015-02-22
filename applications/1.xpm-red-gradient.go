package main

import (
	"fmt"

	"github.com/aznashwan/go-cg/xpm"
)

//		Assignment number 1:
// Write a program that uses your custom XPM library to create the file of a
// 50x50 XPM containing a red vertical gradient (#000000 -> #FF0000)
func main() {
	var i, j uint
	var err error

	// create th new XPM object
	XPM := xpm.NewXPM(50, 50, 1, []xpm.Color{})

	// generate and add all the colors and set the appropriate pixel column
	for i = 0; i < 50; i++ {
		// i + 34 => the ASCII code of a character, staring from 34 because
		// there is a big enough chunk of non-space, non-backslash characters
		// in there for us to use for our 50 color encodings
		cc := string(i + 35)

		// (i / 49) = "percentage" of the red value as we go along
		// green and blue are set to 0 throughout
		err = XPM.AddColor((byte((float64(i) / 49) * 255)), 0, 0, cc)
		if err != nil {
			fmt.Println(err)
			return
		}

		// set all pixels on the i-th column to the same color
		for j = 0; j < 50; j++ {
			err = XPM.SetPixel(j, i, cc)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// write the file
	err = XPM.WriteToFile("red-gradient.xpm")
	if err != nil {
		fmt.Println(err)
	}
}
