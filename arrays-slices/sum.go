package main

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(allNumbers ...[]int) (result []int) {
	for _, numbers := range allNumbers {
		result = append(result, Sum(numbers))
	}
	return
}
