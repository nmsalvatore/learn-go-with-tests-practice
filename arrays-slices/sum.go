package main

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(allNumbers ...[]int) (sum []int) {
	for _, numbers := range allNumbers {
		sum = append(sum, Sum(numbers))
	}
	return
}

func SumAllTails(allNumbers ...[]int) (sum []int) {
	for _, numbers := range allNumbers {
		tail := numbers[1:]
		sum = append(sum, Sum(tail))
	}
	return
}
