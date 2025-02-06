package fill

import "github.com/blackwidow-sudo/govape/calculators"

type (
	Inputs struct {
		HaveVpg      float64 // in ml
		HaveNicotine float64 // in mg/ml
		WantAroma    float64 // in percentage
		WantNicotine float64 // in mg/ml
	}

	Recipe struct {
		Quantity     float64
		Vpg          float64
		NicotineBase float64
		Aroma        float64
	}
)

func (i *Inputs) Calculate() (*Recipe, error) {
	var r Recipe

	if i.HaveNicotine <= 0 {
		return &r, calculators.ErrNoHaveNicotine
	}

	nicPercentage := i.WantNicotine / i.HaveNicotine
	fillPercentage := nicPercentage + (i.WantAroma / 100)
	havePercentage := 1 - fillPercentage

	r.Quantity = (i.HaveVpg / havePercentage)
	r.Aroma = r.Quantity * (i.WantAroma / 100)
	r.NicotineBase = (i.WantNicotine * r.Quantity) / i.HaveNicotine
	r.Vpg = r.Quantity - r.Aroma - r.NicotineBase

	return &r, nil
}
