package combine

func combine(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	combinations := [][]int{}

	if k == 1 {
		for num := 1; num <= n; num++ {
			combinations = append(combinations, []int{num})
		}
		return combinations
	}

	smaller_combinations := combine(n-1, k-1)
	for _, combination := range smaller_combinations {
		for extra_num := combination[k-2] + 1; extra_num <= n; extra_num++ {
			new_combination := make([]int, k)
			copy(new_combination, combination)
			new_combination[k-1] = extra_num
			combinations = append(combinations, new_combination)
		}
	}

	return combinations
}
