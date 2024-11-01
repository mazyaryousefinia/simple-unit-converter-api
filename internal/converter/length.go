package converter

import "fmt"

type LengthConverter struct {
	FromUnit string
	ToUnit   string
	Value    float64
}

var conversionRates = map[string]float64{
	"m":  1,
	"km": 0.001,
	"cm": 100,
}

func (lc *LengthConverter) Convert() (float64, error) {

	fromRate, fromExist := conversionRates[lc.FromUnit]
	toRate, toExist := conversionRates[lc.ToUnit]

	if !fromExist || !toExist {
		return 0, fmt.Errorf("convertion from %s to %s is not supported", lc.FromUnit, lc.ToUnit)
	}

	meter := lc.Value / fromRate

	return meter * toRate, nil
}
