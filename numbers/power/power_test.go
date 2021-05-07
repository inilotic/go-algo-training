package power

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPerfect(t *testing.T) {
	assert.Equal(t, 0.0, Power(0, 0))
	assert.Equal(t, 0.0, Power(0, 3))

	assert.Equal(t, 1.0/8, Power(2, -3))
	assert.Equal(t, 1.0/4, Power(2, -2))
	assert.Equal(t, 1.0/2, Power(2, -1))
	assert.Equal(t, 1.0, Power(2, 0))
	assert.Equal(t, 1.0, Power(3, 0))
	assert.Equal(t, 2.0, Power(2, 1))
	assert.Equal(t, 4.0, Power(2, 2))
	assert.Equal(t, 8.0, Power(2, 3))
	assert.Equal(t, 27.0, Power(3, 3))
	assert.Equal(t, 256.0, Power(4, 4))
}
