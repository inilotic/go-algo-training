package fibonaccis

// Fibonacci, formula: F(N) = F(N - 1) + F(N - 2), for N > 1; examples F(0) = 0; F(1) = 1; F(2) = 1
func Fibonacci(number int) int {
	// borders start
	if number <= 0 {
		return 0
	}

	if number == 1 {
		return 1
	}
	// borders end

	return Fibonacci(number-1) + Fibonacci(number-2)
}
