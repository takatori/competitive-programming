func initDict() map[string][]string {

	return map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
}

var dict = initDict()

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	if len(digits) == 1 {
		return dict[digits]
	}

	result := []string{}

	follows := letterCombinations(digits[1:])

	if letters, ok := dict[string(digits[0])]; ok {
		for _, letter := range letters {
			for _, follow := range follows {
				result = append(result, letter+follow)
			}
		}
	}

	return result
}
