package power

func Power(number, exponent int) float64 {
	if number == 0 {
		return 0
	}

	if exponent == 0 {
		return 1.0
	}

	if exponent == 1 {
		return float64(number)
	}

	if exponent < 0 {
		return 1.0 / Power(number, -exponent)
	}

	result := Power(number, exponent/2)

	if exponent%2 == 0 {
		return result * result
	}

	return float64(number) * result * result
}
