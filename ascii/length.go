package ascii

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	TerminalLength int
	AsciiLength    int
	NoSpaces       int
	countspace     int
)

func CalculateLength() (int, int, int, int) {
	// we calculate the length of the terminal
	cmd := exec.Command("stty", "size")
	// tput cols or tput lines
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	size := strings.Split(strings.TrimSpace(string(output)), " ")
	TerminalLength, err := strconv.Atoi(size[1])
	if err != nil {
		fmt.Println("Invalid terminal width")
	}

	// Now for the ascii characters
	// starting by getting all the values needed
	_, _, Banner, _, inputsplit = ArgsManagement()
	Lines = ReadText(Banner)
	Replace = Populate()

	// no space would be the lenght of the ascii characters minus the spaces
	// countspace will count the number of spaces
	var NoSpaces int = 0
	var countspace int

	// for the calculation we calculate the len of the value of the map in line 0, so the slices corresponding to characters from the input
	for _, line := range inputsplit {
		for j := 0; j < len(line); j++ {
			inputrune := rune(line[j])
			AsciiLength += len(Replace[inputrune][0])
			if inputrune == 32 {
				countspace++
			}

		}
		NoSpaces = AsciiLength - countspace*6

	}
	return TerminalLength, AsciiLength, NoSpaces, countspace
}
