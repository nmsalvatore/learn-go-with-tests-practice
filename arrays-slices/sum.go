package main

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(allNumbers ...[]int) (sum []int) {
	for _, numbers := range allNumbers {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers))
		}
	}
	return
}

func SumAllTails(allNumbers ...[]int) (sum []int) {
	for _, numbers := range allNumbers {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return
}
