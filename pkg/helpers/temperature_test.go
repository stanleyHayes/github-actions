package helpers_test

import (
	"testing"

	"github.com/stanleyHayes/gh-building-blocks/pkg/helpers"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	t.Run("FreezingPoint", func(t *testing.T) {
		celsius := 0.0
		expected := 32.0
		actual := helpers.CelsiusToFahrenheit(celsius)
		if expected != actual {
			t.Errorf("expected %.2f but got %.2f", expected, actual)
		}
	})
}
