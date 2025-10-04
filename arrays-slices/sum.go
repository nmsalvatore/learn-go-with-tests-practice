package main

func Sum(numbers [5]int) (result int) {
	for i := range 5 {
		result += numbers[i]
	}
	return
}
