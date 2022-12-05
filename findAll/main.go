package findall

func findAll(sumOfDigits, numberOfDigits int) []int {
	start := 0
	end := 0
	for i := 1; i < numberOfDigits; i++ {
		start *= 10
		end = start * 10
		println(start, end)
	}
	howMany := 0
	lowestValue := 0
	highestValue := 0
	res := []int{howMany, lowestValue, highestValue}

	println()
	return res
}
