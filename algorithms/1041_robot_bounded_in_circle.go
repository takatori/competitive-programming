package main

func isRobotBounded(instructions string) bool {

	directions := make([][]int, 4)

	directions[0] = []int{0, 1}
	directions[1] = []int{1, 0}
	directions[2] = []int{0, -1}
	directions[3] = []int{-1, 0}

	var x, y, idx int

	for _, i := range instructions {
		if i == 'L' {
			idx = (idx + 3) % 4
		} else if i == 'R' {
			idx = (idx + 1) % 4
		} else {
			x += directions[idx][0]
			y += directions[idx][1]
		}
	}

	return (x == 0 && y == 0) || idx != 0
}
