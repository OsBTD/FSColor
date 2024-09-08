package ascii

import (
	"log"
	"os"
	"strings"
)

var Lines []string

func ReadText(string) []string {
	_, _, Banner, _, _ = ArgsManagement()
	// here we read the text file containing the ascii art characters
	content, err := os.ReadFile("ascii/Banner/" + Banner)
	if err != nil {
		log.Fatal("Error : couldn't read file ", err)
	}
	// we remove \r characters to keep all the files homogeneous no matter on which os they were created
	noreturn := strings.ReplaceAll(string(content), "\r", "")
	// then we split the content by newlines
	Lines = strings.Split(noreturn, "\n")
	return Lines
}
