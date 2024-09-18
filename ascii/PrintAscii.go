package ascii

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func PrintArt() string {
	TerminalLength, AsciiLength, NoSpaces, countspace = CalculateLength()
	filename, align, Banner, _, _, input, inputsplit, _, substringexists = ArgsManagement()
	color := Color()

	var result string
	var before, middle, after string
	// we start by checking if the user input only contains literal newlines
	// if so we print newlines accordingly
	for idx, line := range inputsplit {
		if substringexists {
			index := strings.Index(line, substring)

			if index == -1 || !substringexists {
				fmt.Println("Error : the substring wasn't found in the text")
				before, middle, after = line, "", ""
			} else {
				before = line[:index]
				middle = substring
				after = line[index+len(substring):]
			}
			fmt.Println("before is : ", before, "middle is : ", middle, "after is : ", after)
		}

		// also if there's empty strings resulting from the spliting we print a newline
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			result += "\n"
			fmt.Println()
			continue
		} else if len(line) == 0 && !Checknewline(inputsplit) {
			result += "\n"
			fmt.Println()
		} else if len(line) != 0 && !Checknewline(inputsplit) {
			if align == "--align=left" {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(before); j++ {
						if before == "" {
							continue
						} else {
							inputrune := rune(before[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}
					for j := 0; j < len(middle); j++ {
						inputrune := rune(middle[j])
						result += Replace[inputrune][i]
						fmt.Print(color, Replace[inputrune][i], Reset)
					}
					for j := 0; j < len(after); j++ {
						if after == "" {
							continue
						} else {
							inputrune := rune(after[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}

					// after each line is printed we print a newline
					result += "\n"
					fmt.Println()

				}
			} else if align == "--align=right" {
				for i := 0; i < 8; i++ {
					for l := 0; l < (TerminalLength - AsciiLength); l++ {
						result += " "
						fmt.Print(" ")

					}

					for j := 0; j < len(before); j++ {
						if before == "" {
							continue
						} else {
							inputrune := rune(before[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}
					for j := 0; j < len(middle); j++ {
						inputrune := rune(middle[j])
						result += Replace[inputrune][i]
						fmt.Print(color, Replace[inputrune][i], Reset)
					}
					for j := 0; j < len(after); j++ {
						if after == "" {
							continue
						} else {
							inputrune := rune(after[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}

					// after each line is printed we print a newline
					result += "\n"
					fmt.Println()
				}
			} else if align == "--align=center" {
				for i := 0; i < 8; i++ {
					for l := 0; l < (TerminalLength/2 - AsciiLength/2); l++ {
						result += " "
						fmt.Print("S")
					}

					for j := 0; j < len(before); j++ {
						if before == "" {
							continue
						} else {
							inputrune := rune(before[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}
					for j := 0; j < len(middle); j++ {
						inputrune := rune(middle[j])
						result += Replace[inputrune][i]
						fmt.Print(color, Replace[inputrune][i], Reset)
					}
					for j := 0; j < len(after); j++ {
						if after == "" {
							continue
						} else {
							inputrune := rune(after[j])
							result += Replace[inputrune][i]
							fmt.Print(Reset, Replace[inputrune][i], Reset)

						}
					}

					// after each line is printed we print a newline
					result += "\n"
					fmt.Println()
				}
			} else if align == "--align=justify" {
				for i := 0; i < 8; i++ {
					for j := 0; j < len(before); j++ {
						if before == "" {
							continue
						} else {
							if before[j] != 32 {
								inputrune := rune(before[j])
								result += Replace[inputrune][i]
								fmt.Print(color, Replace[inputrune][i], Reset)

							} else {
								for l := 0; l < (TerminalLength-NoSpaces)/countspace; l++ {
									result += " "
									fmt.Print(" ")

								}
							}

						}
					}
					for j := 0; j < len(middle); j++ {
						if middle[j] != 32 {
							inputrune := rune(middle[j])
							result += Replace[inputrune][i]
							fmt.Print(color, Replace[inputrune][i], Reset)

						} else {
							for l := 0; l < (TerminalLength-NoSpaces)/countspace; l++ {
								result += " "
								fmt.Print(" ")

							}
						}
					}
					for j := 0; j < len(after); j++ {
						if after == "" {
							continue
						} else {
							if after[j] != 32 {
								inputrune := rune(after[j])
								result += Replace[inputrune][i]
								fmt.Print(color, Replace[inputrune][i], Reset)

							} else {
								for l := 0; l < (TerminalLength-NoSpaces)/countspace; l++ {
									result += " "
									fmt.Print(" ")

								}
							}

						}
					}

					// after each line is printed we print a newline
					result += "\n"
					fmt.Println()
				}
			}
		}
	}
	resfile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()

	_, err = io.WriteString(resfile, result)
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}

	return result
}
