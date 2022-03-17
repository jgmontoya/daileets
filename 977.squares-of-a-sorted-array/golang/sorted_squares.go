package sorted_squares

func SortedSquares(nums []int) []int {
	negatives := []int{}
	result := []int{}

	index := 0
	for index < len(nums) && nums[index] < 0 {
		negatives = append(negatives, nums[index]*nums[index])
		index++
	}

	negatives_index := index - 1

	for len(result) < len(nums) {
		if index < len(nums) && negatives_index >= 0 {
			result, index, negatives_index = AppendSmallerSquare(nums, index, negatives, negatives_index, result)
		} else if negatives_index < 0 {
			result, index = AppendPositivesSquare(nums, index, result)
		} else {
			result, negatives_index = AppendNegativesSquare(negatives, negatives_index, result)
		}
	}

	return result
}

func AppendSmallerSquare(nums []int, index int, negatives []int, negatives_index int, result []int) ([]int, int, int) {
	non_negative_element := nums[index] * nums[index]
	negative_element := negatives[negatives_index]
	if non_negative_element <= negative_element {
		result = append(result, non_negative_element)
		index++
	} else {
		result = append(result, negative_element)
		negatives_index--
	}

	return result, index, negatives_index
}

func AppendPositivesSquare(nums []int, index int, result []int) ([]int, int) {
	result = append(result, nums[index]*nums[index])
	index++
	return result, index
}

func AppendNegativesSquare(negatives []int, negatives_index int, result []int) ([]int, int) {
	result = append(result, negatives[negatives_index])
	negatives_index--
	return result, negatives_index
}
