package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filename        string
	align           string
	Banner          string
	input           string
	inputsplit      []string
	args            []string
	color           string
	substring       string
	inputsubstr     string
	substringexists bool
)

func ArgsManagement() (string, string, string, string, string, string, []string, string, bool) {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 6 {
		log.Fatal("Usage: go run . [alignment] [output] [color] [substring] [STRING] [BANNER]\n\nEX: go run . --align=center --output=<fileName.txt> --color=red substring something standard")
	}

	// Usage instructions if no arguments or more than 2 arguments are provided
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner := "standard.txt"
	filename := "--output=result.txt"
	align := "--align=left"
	input := ""
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
			substringexists = true

		}
	} else if alignexists && outputexists && colorexists && !bannerexists {
		args = append(args, Banner)
		if len(args) == 5 {
			input = args[len(args)-2]
			Banner = args[len(args)-1]
		} else if len(args) == 6 {
			input = args[len(args)-2]
			Banner = args[len(args)-1]
			substring = args[index+1]
			substringexists = true

		}
	} else if alignexists && outputexists && bannerexists && !colorexists {
		input = args[len(args)-2]
	} else if alignexists && colorexists && bannerexists && !outputexists {
		args = append(args, filename)
		if len(args) == 5 {
			filename = args[len(args)-1]
			input = args[len(args)-3]
			Banner = args[len(args)-2]

		} else if len(args) == 6 {
			filename = args[len(args)-1]
			input = args[len(args)-3]
			substring = args[index+1]
			Banner = args[len(args)-2]
			substringexists = true

		}

	} else if outputexists && colorexists && bannerexists && !alignexists {
		args = append(args, align)
		if len(args) == 5 {
			align = args[len(args)-1]
			input = args[len(args)-3]
			Banner = args[len(args)-2]

		} else if len(args) == 6 {
			align = args[len(args)-1]
			input = args[len(args)-3]
			substring = args[index+1]
			Banner = args[len(args)-2]
			substringexists = true

		}
	} else if outputexists && colorexists && !bannerexists && !alignexists {
		args = append(args, align, Banner)
		if len(args) == 5 {
			align = args[len(args)-2]
			input = args[len(args)-3]
			Banner = args[len(args)-1]

		} else if len(args) == 6 {
			align = args[len(args)-2]
			input = args[len(args)-3]
			Banner = args[len(args)-1]
			substring = args[index+1]
			substringexists = true

		}

	} else if outputexists && bannerexists && !colorexists && !alignexists {
		args = append(args, align)
		align = args[len(args)-1]
		input = args[len(args)-3]
		Banner = args[len(args)-2]

	} else if alignexists && bannerexists && !colorexists && !outputexists {
		args = append(args, filename)
		filename = args[len(args)-1]
		input = args[len(args)-3]
		Banner = args[len(args)-2]

	} else if alignexists && colorexists && !bannerexists && !outputexists {
		args = append(args, filename, Banner)
		if len(args) == 5 {
			filename = args[len(args)-2]
			input = args[len(args)-3]
			Banner = args[len(args)-1]

		} else if len(args) == 6 {
			filename = args[len(args)-2]
			input = args[len(args)-3]
			Banner = args[len(args)-1]
			substring = args[index+1]
			substringexists = true

		}

	} else if colorexists && bannerexists && !alignexists && !outputexists {
		args = append(args, filename, align)
		if len(args) == 5 {
			filename = args[len(args)-2]
			input = args[len(args)-4]
			Banner = args[len(args)-3]
			align = args[len(args)-1]

		} else if len(args) == 6 {
			filename = args[len(args)-2]
			input = args[len(args)-4]
			align = args[len(args)-1]
			Banner = args[len(args)-3]
			substring = args[index+1]
			substringexists = true
			fmt.Println(alignexists, bannerexists, outputexists, "color:", colorexists)

		}

	} else if outputexists && alignexists && !colorexists && !bannerexists {
		args = append(args, Banner)
		input = args[len(args)-3]
		Banner = args[len(args)-1]

	} else if outputexists && !colorexists && !bannerexists && !alignexists {
		args = append(args, align, Banner)
		align = args[len(args)-2]
		input = args[len(args)-3]
		Banner = args[len(args)-1]
	} else if alignexists && !colorexists && !bannerexists && !outputexists {
		args = append(args, filename, Banner, color)
		align = args[0]
		input = args[1]
		filename = args[2]
		Banner = args[3]
		color = args[4]
	} else if bannerexists && !colorexists && !alignexists && !outputexists {
		args = append(args, align, filename)
		align = args[len(args)-2]
		filename = args[len(args)-1]
		Banner = args[len(args)-3]
		input = args[len(args)-4]

	} else if colorexists && !alignexists && !bannerexists && !outputexists {
		args = append(args, align, filename, Banner)
		if len(args) == 5 {
			filename = args[len(args)-2]
			align = args[len(args)-3]
			Banner = args[len(args)-1]
			input = args[len(args)-4]
		} else if len(args) == 6 {
			filename = args[len(args)-2]
			align = args[len(args)-3]
			Banner = args[len(args)-1]
			input = args[len(args)-4]
			substring = args[index+1]
			substringexists = true

			fmt.Println(alignexists, bannerexists, outputexists, "color:", colorexists)

		}

	} else if !colorexists && !alignexists && !bannerexists && !outputexists {
		args = append(args, align, filename, Banner, color)
		if len(args) == 5 {
			input = args[0]
			align = args[1]
			filename = args[2]
			Banner = args[3]
			color = args[4]

		} else {
			log.Fatal("Usage: go run . [alignment] [output] [color] [substring] [STRING] [BANNER]\n\nEX: go run . --align=center --output=<fileName.txt> --color=red substring something standard")
		}
	} else {
		log.Fatal("Usage: go run . [alignment] [output] [color] [substring] [STRING] [BANNER]\n\nEX: go run . --align=center --output=<fileName.txt> --color=red substring something standard")
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

	// var inputsubstr string
	// j := 0
	// k := 0
	// var substr bool
	// for i := 0; i < len(inputsplit); i++ {
	// 	if inputsplit[i] != "" {
	// 		for k < len(inputsplit[i]) && j < len(substring) {
	// 			if inputsplit[i][k] == substring[j] {
	// 				inputsubstr += string(substring[j])
	// 				j++
	// 				k++
	// 			} else {
	// 				j = 0
	// 				k++
	// 				continue
	// 			}
	// 		}

	// 	} else {
	// 		continue
	// 	}

	// 	if inputsubstr == substring {
	// 		substr = true
	// 	} else {
	// 		substr = false
	// 		inputsubstr = ""
	// 	}
	// }
	// for _, line := range inputsplit {
	// 	for k := 0; k+len(inputsubstr) < len(line); k += len(inputsubstr) {
	// 		if substringexists && line[k:len(inputsubstr)] == inputsubstr {
	// 			line = "é" + line[k:len(inputsubstr)] + "é"
	// 			continue
	// 		} else {
	// 			line = line[k:len(inputsubstr)]
	// 			continue
	// 		}

	// 	}
	// 	fmt.Println("line is :", line)
	// }

	filename = strings.TrimPrefix(filename, "--output=")
	color = strings.TrimPrefix(color, "--color=")
	fmt.Println("color in args is :", color, "substring in args is :", substring, "inputsubstr in args is :", inputsubstr, colorexists, len(args))
	return filename, align, Banner, color, substring, input, inputsplit, inputsubstr, substringexists
}
