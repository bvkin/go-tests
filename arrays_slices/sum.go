package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sumAll []int) {
	for _, numSlice := range numbersToSum {
		sumAll = append(sumAll, Sum(numSlice))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sumAllTails[]int) {
	for _, numSlice := range numbersToSum {
		if len(numSlice) == 0 {
			sumAllTails = append(sumAllTails, 0)
		} else {
			sumAllTails = append(sumAllTails, Sum(numSlice[1:]))
		}
	}
	return
}
