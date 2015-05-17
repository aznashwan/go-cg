package main

import (
	"flag" // for flag-handling related work
	"fmt"

	ps "../../postscript"
	"../../transformations/twod"
	"../../xpm"
)

var usage string = `
USAGE: cmd.exe -f /path/to/input.ps -w 200 -h 200 -o /path/to/output.xpm

-f:
	Path to the postscript-format input file containing line definitions.
	Mandatory argument.
-w:
	Width of the output XPM bitmap file.
	Mandatory. Must be greater than 0.
-h:
	Height of the output XPM bitmap file.
	Mandatory. Must be greater than 0.
-o:
	Path to the output XPM bitmap file.
	Default value is ./output.xpm
-t:
    Path to file defining 2d transformations.
`[1:]

// height command line argument
// usage: -h UINT
// mandatory
var height int

// width command line argument
// usage: -w UINT
// mandatory
var width int

// postscript input file command line argument
// usage: -f /path/to/file.ps
// mandatory
var input string

// transformations file command line argument
// usage: -t /path/to/file.tsf
// optional
var trans string

// xpm output file command line argument
// usage: -o /path/to/file.xpm
// default: ./output.xpm
var output string

// flaginit sets up all command line flag handling
func flaginit() {
	flag.IntVar(&width, "w", 0, "width of the resulting bitmap")
	flag.IntVar(&height, "h", 0, "height of the resulting bitmap")
	flag.StringVar(&input, "f", "", "postscript input file given for processing")
	flag.StringVar(&output, "o", "./output.xpm", "output file for resulting bitmap")
	flag.StringVar(&trans, "t", "", "transformations definition file")
	flag.Parse()
}

//			Assignment 2:
// Write a program which takes some command line aruments and parses a provided
// input file which *exclusively* contains postscript line definitions,
// generating an output file in XPM format with the resulting bitmap
func main() {
	var err error

	// initialize all command line flags
	flaginit()

	// check all arguments
	if height <= 0 || width <= 0 || input == "" {
		fmt.Println(usage)
		return
	}

	// create XPM struct to be worked on
	xpm := xpm.NewXPM(width, height, 1)

	// add our preffered color for all subsequent line designs
	// in this case, 100% blue balance
	xpm.AddColor(0, 0, 255, "b")

	// parse the input file
	lines, err := ps.ParseFile(input)
	if err != nil {
		fmt.Printf("Error parsing input file %s:\n%s\n", input, err)
		return
	}

	// apply transformations if transformations file was given:
	if trans != "" {
		ops, err := twod.ParseFile(trans)
		if err != nil {
			fmt.Printf("Error parsing transformations file %s: \n%s\n", trans, err)
		}
		lines = twod.ApplyTransformationsToLines(lines, ops)
	}

	// have each line draw itself to the XPM
	for i, line := range lines {
		if err := line.Draw(xpm, "b"); err != nil {
			fmt.Printf("Error drawing %d'th line:\n%s\n", i, err)
		}
	}

	// finally, write out out resulting XPM to the output file
	if err := xpm.WriteToFile(output); err != nil {
		fmt.Printf("Error writing output file %s:\n%s", output, err)
	}

}
