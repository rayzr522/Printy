package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	// Input The text passed to the program
	Input string
	// Parsed The parsed text, transformed by the program
	Parsed string
)

var (
	// FlagNewline Whether the program should print a newline or not
	FlagNewline bool
	// FlagReset If true, send a reset ANSI code to clear formatting
	FlagReset bool
)

func init() {
	flag.BoolVar(&FlagNewline, "n", false, "Disable newline")
	flag.BoolVar(&FlagReset, "r", false, "Send a reset ANSI code (clear formatting)")

	flag.Parse()

	Input = strings.Join(flag.Args(), " ")
}

func main() {
	parse()
	if FlagNewline {
		fmt.Print(Parsed)
	} else {
		fmt.Println(Parsed)
	}
}

func parse() {
	if FlagReset {
		Parsed = makeColor('r')
	} else {
		Parsed = parseColors(Input)
	}
}

func makeColor(b byte) string {
	var code string

	if b >= '0' && b <= '9' {
		code = "3" + string(b)
	} else if b == 'b' {
		code = "1"
	} else {
		code = "0"
	}

	return "\u001b[" + code + "m"
}

func isColorChar(b byte) bool {
	return (b >= '0' && b <= '9') || b == 'r' || b == 'b'
}

func parseColors(input string) string {
	output := ""

	i := 0
	for i <= len(input)-3 {
		a := input[i]
		b := input[i+1]
		c := input[i+2]

		if a == '#' && isColorChar(b) && c == '#' {
			output += makeColor(b)
			i += 3
		} else {
			output += string(a)
			i++
		}
	}

	output += input[i:]

	return output
}
