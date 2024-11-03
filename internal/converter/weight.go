package converter

import "fmt"

type WeightConverter struct {
	FromUnit string
	ToUnit   string
	Value    float64
}

var weightConversionRates = map[string]float64{
	"g":  1,
	"kg": 0.001,
	"lb": 0.00220462,
}

func (wc *WeightConverter) Convert() (float64, error) {

	fromRate, fromExists := weightConversionRates[wc.FromUnit]
	toRate, toExists := weightConversionRates[wc.ToUnit]

	if !fromExists || !toExists {
		return 0, fmt.Errorf("conversion from %s to %s not supported", wc.FromUnit, wc.ToUnit)
	}

	kilograms := wc.Value / fromRate
	return kilograms * toRate, nil

}
