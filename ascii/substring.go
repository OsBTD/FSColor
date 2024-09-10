package ascii

func ColorSubstring() {
	filename, align, _, color, substring, input, inputSplit := ArgsManagement() // we read the ASCII art characters from the chosen banner file
	replaceMap := Populate()

	result := PrintArt(filename, align, input, inputSplit, replaceMap)

	for _, word := range result {
	}
}
