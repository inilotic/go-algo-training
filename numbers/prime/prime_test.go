package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	assert.False(t, IsPrimeNaive(-1))
	assert.False(t, IsPrimeNaive(0))
	assert.True(t, IsPrimeNaive(2))
	assert.True(t, IsPrimeNaive(3))
	assert.True(t, IsPrimeNaive(5))
	assert.False(t, IsPrimeNaive(6))
	assert.True(t, IsPrimeNaive(7))
	assert.False(t, IsPrimeNaive(9))
	assert.True(t, IsPrimeNaive(31))
	assert.True(t, IsPrimeNaive(67))
	assert.False(t, IsPrimeNaive(4))
	assert.False(t, IsPrimeNaive(194))
	assert.False(t, IsPrimeNaive(9999))

	assert.False(t, IsPrimeVer1(-1))
	assert.False(t, IsPrimeVer1(0))
	assert.True(t, IsPrimeVer1(2))
	assert.True(t, IsPrimeVer1(3))
	assert.True(t, IsPrimeVer1(5))
	assert.False(t, IsPrimeVer1(6))
	assert.True(t, IsPrimeVer1(7))
	assert.False(t, IsPrimeVer1(9))
	assert.True(t, IsPrimeVer1(31))
	assert.True(t, IsPrimeVer1(67))
	assert.False(t, IsPrimeVer1(4))
	assert.False(t, IsPrimeVer1(194))
	assert.False(t, IsPrimeVer1(9999))
}

// primes: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199
