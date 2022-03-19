package two_sum

func TwoSum(numbers []int, target int) []int {
	left_index := 0
	right_index := len(numbers) - 1

	sum := numbers[left_index] + numbers[right_index]

	for sum != target {
		if sum < target {
			left_index++
		} else {
			right_index--
		}
		sum = numbers[left_index] + numbers[right_index]
	}

	return []int{left_index + 1, right_index + 1}
}
