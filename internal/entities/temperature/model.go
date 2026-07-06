package temperature

import "github.com/stanleyHayes/gh-building-blocks/pkg/helpers"

type Unit string

const UnitFahrenheit Unit = "Fahrenheit"
const UnitCelsius Unit = "Celsius"

type Temperature struct {
	Value float64
	Unit  Unit
}

func New(value float64, unit Unit) *Temperature {
	return &Temperature{Value: value, Unit: unit}
}

func (t *Temperature) ToFahrenheit() *Temperature {
	if t.Unit == UnitFahrenheit {
		return t
	}
	return &Temperature{Value: helpers.CelsiusToFahrenheit(t.Value), Unit: UnitFahrenheit}
}
func (t *Temperature) ToCelsius() *Temperature {
	if t.Unit == UnitCelsius {
		return t
	}
	return &Temperature{Value: helpers.FahrenheitToCelsius(t.Value), Unit: UnitCelsius}
}
