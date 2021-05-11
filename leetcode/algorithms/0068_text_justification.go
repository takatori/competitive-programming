package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {

	var i int
	l := len(words)
	result := make([]string, 0)

	for i < l {

		var pack []string
		currentWidth := 0
		j := i

		for j < l && currentWidth+len(words[j]) <= maxWidth {
			pack = append(pack, words[j])
			currentWidth = currentWidth + len(words[j]) + 1 // スペース分を足す
			j++
		}

		wl := len(pack)
		var cl int
		for k := 0; k < wl; k++ {
			cl = cl + len(pack[k])
		}
		spaces := maxWidth - cl
		slots := wl - 1

		if j >= l || slots == 0 {
			line := pack[0]
			for p := 1; p < wl; p++ {
				line = line + " " + pack[p]
			}
			line = line + strings.Repeat(" ", spaces-wl+1)
			result = append(result, line)
		} else {

			spaceStr := strings.Repeat(" ", spaces/slots)
			moreSpacePosition := spaces % slots
			line := pack[0]

			for p := 1; p < wl; p++ {
				line = line + spaceStr
				if p <= moreSpacePosition {
					line = line + " "
				}
				line = line + pack[p]
			}

			result = append(result, line)
		}

		i = j
	}

	return result
}
