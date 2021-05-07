package prime

func IsPrimeNaive(number int) bool {
	// border
	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if (number % i) == 0 {
			return false
		}
	}

	return true
}

func IsPrimeVer1(number int) bool {
	// border
	if number <= 1 {
		return false
	}

	if number == 2 {
		return true
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i = i + 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}
