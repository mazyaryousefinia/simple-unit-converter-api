package converter

import "fmt"

type TemperatureConvert struct {
	FromUnit string
	ToUnit   string
	Value    float64
}

func (tc *TemperatureConvert) Convert() (float64, error) {

	switch tc.FromUnit {
	case "c":
		if tc.ToUnit == "f" {
			return (tc.Value * 9 / 5) + 32, nil
		} else if tc.ToUnit == "k" {
			return tc.Value + 273.15, nil
		}
	case "f":
		if tc.ToUnit == "c" {
			return (tc.Value - 32) * 5 / 9, nil
		} else if tc.ToUnit == "k" {
			return (tc.Value-32)*5/9 + 273.15, nil
		}
	case "k":
		if tc.ToUnit == "c" {
			return tc.Value - 273.15, nil
		} else if tc.ToUnit == "f" {
			return (tc.Value-273.15)*9/5 + 32, nil
		}
	}
	return 0, fmt.Errorf("convertion from %s to %s is not supported", tc.FromUnit, tc.ToUnit)
}
