package ascii

func Checknewline(inputsplit []string) bool {
	for _, line := range inputsplit {
		if len(line) != 0 {
			return false
		}
	}

	return true
}
