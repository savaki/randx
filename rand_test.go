package randx_test

import (
	"regexp"
	"testing"

	"github.com/savaki/randx"
	"github.com/stretchr/testify/assert"
)

func TestAlphaN(t *testing.T) {
	for i := 0; i < 20; i++ {
		v := randx.AlphaN(i)
		assert.Len(t, v, i)
	}
}

func TestDigitN(t *testing.T) {
	for i := 0; i < 20; i++ {
		v := randx.DigitN(i)
		assert.Len(t, v, i)
	}
}

func TestEmail(t *testing.T) {
	re := regexp.MustCompile(`\S+.\S+@example.com`)
	email := randx.Email()
	assert.True(t, re.MatchString(email))

	email = randx.Email("a", "b")
	assert.Contains(t, email, "a.b")
	assert.Contains(t, email, "@example.com")
}

func TestIntN(t *testing.T) {
	max := 100
	for i := 1; i < 1000; i++ {
		v := randx.IntN(max)
		assert.True(t, v < max)
	}
}
