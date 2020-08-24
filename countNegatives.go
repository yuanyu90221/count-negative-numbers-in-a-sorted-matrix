package count_neg

func countNegatives(grid [][]int) int {
	count := 0
	lenY := len(grid)
	lenX := len(grid[0])
	for idxY := lenY - 1; idxY >= 0; idxY-- {
		for idxX := lenX - 1; idxX >= 0; idxX-- {
			if grid[idxY][idxX] < 0 {
				count += 1
			} else {
				continue
			}
		}
	}
	return count
}
