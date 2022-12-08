package findall

func findAll(sumOfDigits, numberOfDigits int) []int {
	howMany := 0
	lowestValue := 0
	highestValue := 0
	start := 1
	end := 0
	// finds start and end range of possible numbers
	for i := 1; i < numberOfDigits; i++ {
		start *= 10
		end = start * 10
	}
	// finds lowest and highest values and how many there are
	for i := start; i < end; i++ {
		digits := digitize(i)
		summedDigits := addSlice(digits)
		digitsAreInIncreasingOrder := increasingOrderCheck(digits)
		switch {
		case digitsAreInIncreasingOrder:
			switch {
			case summedDigits == sumOfDigits && lowestValue == 0:
				lowestValue = i
				highestValue = i
				howMany++
			case summedDigits == sumOfDigits:
				highestValue = i
				howMany++
			}
		}
	}
	switch {
	case howMany == 0:
		println("no matches found")
		return []int{}
	default:
		return []int{howMany, lowestValue, highestValue}
	}
}

// range fill slice with digits
func digitize(x int) []int {
	size := howManyDigits(x)
	digit := 0
	digits := make([]int, size)
	for i := range digits {
		if x < 1 {
			break
		}
		digit = x % 10
		x /= 10
		digits[len(digits)-i-1] = digit
	}
	return digits
}

// gets the number of digits in a number.
func howManyDigits(x int) int {
	digits := 0
	for x >= 1 {
		x /= 10
		digits++
	}
	return digits
}

// addSlice adds the digits in a slice
func addSlice(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// checks if the digits in a slice are in increasing order
func increasingOrderCheck(x []int) bool {
	for i := 0; i < len(x)-1; i++ {
		if x[i] > x[i+1] {
			return false
		}
	}
	return true
}
