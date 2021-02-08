package arraysandslices

func Sum(numbers [5]int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumSlice(numbers []int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}
