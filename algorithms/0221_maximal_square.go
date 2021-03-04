package main

func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	max := 0
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				tmp := matrixSize(matrix, i, j)
				area := tmp * tmp
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

func matrixSize(matrix [][]byte, i int, j int) int {
	local_max := 1

	for i+local_max < len(matrix) && j+local_max < len(matrix[0]) {

		for mi := 0; mi <= local_max; mi++ {
			if matrix[i+mi][j+local_max] == '0' {
				return local_max
			}
		}

		for nj := 0; nj < local_max; nj++ {
			if matrix[i+local_max][j+nj] == '0' {
				return local_max
			}
		}

		local_max++
	}

	return local_max
}
