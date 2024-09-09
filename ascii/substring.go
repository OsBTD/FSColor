package ascii

import (
	"os"
	"strings"
)

func ColorSubstring() {
	args := os.Args[1:]
	filename, align, _, input, inputSplit := ArgsManagement() // we read the ASCII art characters from the chosen banner file
	replaceMap := Populate()

	result := PrintArt(filename, align, input, inputSplit, replaceMap)

	Fields := strings.Fields(result)
	for _, word := range Fields {
	}
}
