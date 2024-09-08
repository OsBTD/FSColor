package ascii

import (
	"log"
	"os"
	"strings"
)

var (
	filename   string
	align      string
	Banner     string
	input      string
	inputsplit []string
	args       []string
)

func ArgsManagement() (string, string, string, string, []string) {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 4 {
		log.Fatal("Usage: go run . [OUTPUT] [ALIGN] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> --align=center something standard")
	}

	// Usage instructions if no arguments or more than 2 arguments are provided
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner := "standard.txt"
	filename := "--output=result.txt"
	align := "--align=left"
	input := ""

	if len(args) == 1 {
		args = append(args, filename, Banner, align)
		input = args[0]
		filename = args[1]
		Banner = args[2]
		align = args[3]

	} else if len(args) == 2 && strings.HasPrefix(args[0], "--output=") {
		args = append(args, Banner, align)
		filename = args[0]
		input = args[1]
		Banner = args[2]
		align = args[3]

	} else if len(args) == 2 && (args[1] == "standard.txt" || args[1] == "shadow.txt" || args[1] == "thinkertoy.txt" || args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		args = append(args, filename, align)
		input = args[0]
		Banner = args[1]
		filename = args[2]
		align = args[3]

	} else if len(args) == 2 && strings.HasPrefix(args[0], "--align=") {
		args = append(args, Banner, filename)
		align = args[0]
		input = args[1]
		Banner = args[2]
		filename = args[3]
	} else if len(args) == 3 && strings.HasPrefix(args[0], "--align=") && (args[2] == "standard.txt" || args[2] == "shadow.txt" || args[2] == "thinkertoy.txt" || args[2] == "standard" || args[2] == "shadow" || args[2] == "thinkertoy") {
		args = append(args, filename)
		align = args[0]
		input = args[1]
		Banner = args[2]
		filename = args[3]
	} else if len(args) == 3 && strings.HasPrefix(args[0], "--output=") && (args[2] == "standard.txt" || args[2] == "shadow.txt" || args[2] == "thinkertoy.txt" || args[2] == "standard" || args[2] == "shadow" || args[2] == "thinkertoy") {
		args = append(args, align)
		filename = args[0]
		input = args[1]
		Banner = args[2]
		align = args[3]

	} else if len(args) == 3 && strings.HasPrefix(args[0], "--output=") && strings.HasPrefix(args[1], "--align=") {
		args = append(args, Banner)
		filename = args[0]
		align = args[1]
		input = args[2]
		Banner = args[3]
	} else if len(args) == 3 && strings.HasPrefix(args[0], "--align=") && strings.HasPrefix(args[1], "-output=") {
		args = append(args, Banner)
		align = args[0]
		filename = args[1]
		input = args[2]
		Banner = args[3]
	} else if len(args) == 4 && strings.HasPrefix(args[0], "--align=") {
		align = args[0]
		filename = args[1]
		input = args[2]
		Banner = args[3]
	} else if len(args) == 4 && strings.HasPrefix(args[0], "--output=") {
		filename = args[0]
		align = args[1]
		input = args[2]
		Banner = args[3]

	} else {
		log.Fatal("Usage: go run . [OUTPUT] [ALIGN] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> --align=center something standard\nyou can start with either the name of the output file or the alignment\nmake sure to respect the for of the alignment and the output file name\nif you don't input any of these parameters a default value will be assigned")
	}
	if !strings.HasSuffix(Banner, ".txt") {
		Banner = Banner + ".txt"
	}
	if Banner != "standard.txt" && Banner != "shadow.txt" && Banner != "thinkertoy.txt" {
		log.Fatal("Usage: this style is unavailabe\nplease chose one of the available styles\n1 : standard.txt\n2 : shadow.txt\n3 : thinkertoy.txt")
	}
	if !strings.HasPrefix(filename, "--output=") {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	inputsplit = strings.Split(input, "\\n")
	// we check if the input is valid and only contains printable ascii characters
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}

	filename = strings.TrimPrefix(filename, "--output=")

	return filename, align, Banner, input, inputsplit
}
