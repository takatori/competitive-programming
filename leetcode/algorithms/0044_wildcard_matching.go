func isMatch(s string, p string) bool {
    
	si := 0
	pi := 0
	match := 0
	starIdx := -1

	for si < len(s) {

		if pi < len(p) && (p[pi] == '?' || s[si] == p[pi]) {
			si++
			pi++
		} else if pi < len(p) && p[pi] == '*' {
			starIdx = pi
			match = si
			pi++
		} else if starIdx != -1 {
			pi = starIdx + 1
			match++
			si = match
		} else {
			return false
		}
	}

	for pi < len(p) && p[pi] == '*' {
		pi++
	}

	return pi == len(p)
}
