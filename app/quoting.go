package main

func filterArgs(input string) ([]string, bool) {
	if len(input) == 0 {
		return nil, true
	}
	var res = []string{}
	curr := ""
	inSingleQuote := false
	isDoubleQuote := false
	for _, runeValue := range input {
		if runeValue == '\'' && !isDoubleQuote {
			inSingleQuote = !inSingleQuote
			continue
		}
		if runeValue == '"' && !inSingleQuote {
			isDoubleQuote = !isDoubleQuote
			continue
		}

		if runeValue == ' ' && !inSingleQuote && !isDoubleQuote {
			if len(curr) > 0 {
				res = append(res, curr)
				curr = ""
			}
		} else {
			curr += string(runeValue)
		}
	}
	if len(curr) > 0 {
		res = append(res, curr)
	}
	return res, false
}

func filterFields(input string) ([]string, bool) {
	if len(input) == 0 {
		return nil, true
	}
	var res = []string{}

	// seperate command
	for index, runeValue := range input {
		// fmt.Printf("Index: %d, Rune: %c, Unicode Point: %#U\n", index, runeValue, runeValue)
		if runeValue == ' ' {
			command := input[:index]
			res = append(res, command)
			input = input[index:]
			break
		}
		if index == len(input)-1 {
			command := input[:index+1]
			res = append(res, command)
			return res, false
		}
	}

	args, _ := filterArgs(input)
	res = append(res, args...)

	return res, false
}
