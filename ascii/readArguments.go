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
	if len(args) == 0 || len(args) > 6 {
		log.Fatal("Usage: go run . [OUTPUT] [ALIGN] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> --align=center something standard")
	}

	// Usage instructions if no arguments or more than 2 arguments are provided
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner := "standard.txt"
	filename := "--output=result.txt"
	align := "--align=left"
	input := ""
	substring := ""
	color := ""
	var alignexists, outputexists, colorexists, bannerexists bool
	var index int
	for i := 0; i < len(args); i++ {
		if i != len(args)-1 && strings.HasPrefix(args[i], "--align=") {
			align = args[i]
			alignexists = true
		} else if i != len(args)-1 && strings.HasPrefix(args[i], "--output=") {
			filename = args[i]
			outputexists = true
		} else if i != len(args)-1 && strings.HasPrefix(args[i], "--color=") {
			color = args[i]
			colorexists = true
			index = i
		} else if args[len(args)-1] == "standard.txt" || args[len(args)-1] == "shadow.txt" || args[len(args)-1] == "thinkertoy.txt" || args[len(args)-1] == "standard" || args[len(args)-1] == "shadow" || args[len(args)-1] == "thinkertoy" {
			Banner = args[len(args)-1]
			bannerexists = true
		}
	}

	if alignexists && outputexists && colorexists && bannerexists {
		if len(args) == 5 {
			input = args[len(args)-2]
		} else if len(args) == 6 {
			input = args[len(args)-2]
			substring = args[index+1]
		}
	} else if alignexists && outputexists && colorexists && !bannerexists {
		if len(args) == 4 {
			input = args[len(args)-1]
		} else if len(args) == 5 {
			input = args[len(args)-1]
			substring = args[index+1]
		}
	} else if alignexists && outputexists && bannerexists && !colorexists {
	} else if alignexists && colorexists && bannerexists && !outputexists {
	} else if outputexists && colorexists && bannerexists && !alignexists {
	} else if outputexists && colorexists && !bannerexists && !alignexists {
	} else if outputexists && bannerexists && !colorexists && !alignexists {
	} else if alignexists && bannerexists && !colorexists && !outputexists {
	} else if alignexists && colorexists && !bannerexists && !outputexists {
	} else if colorexists && bannerexists && !alignexists && !outputexists {
	} else if outputexists && alignexists && !colorexists && !bannerexists {
	} else if outputexists && !colorexists && !bannerexists && !alignexists {
	} else if alignexists && !colorexists && !bannerexists && !outputexists {
	} else if bannerexists && !colorexists && !alignexists && !outputexists {
	} else if colorexists && !alignexists && !bannerexists && !outputexists {
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
