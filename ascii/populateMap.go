package ascii

var Replace map[rune]([]string)

func Populate() map[rune]([]string) {
	// this map will associate between normal characters with ascii art characters from the text file
	Replace = map[rune]([]string){}
	//we get the style the user chose in Banner, while Lines is the the whole text file split by newline characters
	_, _, Banner, _, _ = ArgsManagement()
	Lines = ReadText(Banner)
	//char will be the keys of the map and will go from 32 ' ' to 126 '~' so all ascii printable characters
	Char := 32
	//we iterate by 9 and populate the map with 8 lines representing the full ascii character from the file
	for i := 0; i < len(Lines); i += 9 {
		if i+9 <= len(Lines)-1 {
			Replace[rune(Char)] = Lines[i+1 : i+9]
		}
		if Char <= 126 {
			Char++
		}

	}
	//return the map fully populated with all printable characters as keys and the corresponding ascii art characters as values
	return Replace
}
