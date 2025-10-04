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

func SumAllTails(allTails ...[]int) (sum []int) {
	for _, tails := range allTails {
		sum = append(sum, Sum(tails[1:]))
	}
	return
}
