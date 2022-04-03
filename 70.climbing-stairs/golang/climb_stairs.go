package climb_stairs

func climbStairs(n int) int {
	first := 0
	second := 1
	var solution int
	for num := 1; num <= n; num++ {
		solution = first + second
		first = second
		second = solution
	}

	return solution
}
