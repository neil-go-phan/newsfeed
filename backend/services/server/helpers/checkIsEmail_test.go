package helpers
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsEmail(t *testing.T) {
	validEmail := "test@example.com"
	isValid := CheckIsEmail(validEmail)
	assert.True(t, isValid)

	invalidEmail := "testexample.com"
	isValid = CheckIsEmail(invalidEmail)
	assert.False(t, isValid)

	emailWithWhitespace := "test @example.com"
	isValid = CheckIsEmail(emailWithWhitespace)
	assert.False(t, isValid)
}