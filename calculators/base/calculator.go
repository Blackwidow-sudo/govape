package base

import (
	"github.com/blackwidow-sudo/govape/calculators"
)

type (
	Inputs struct {
		// The nicotine strength of the base you have in mg/ml
		HaveNicotine	float64

		// The amount of nicotine you want in the final mix in mg/ml
		WantNicotine	float64

		// The amount of aroma you want in the final mix in percentage
		WantAroma		float64

		// The total quantity of the liquid you want in milliliters
		WantQuantity	float64

		// The amount of PG you want in the final mix in percentage
		WantPG			float64

		// The amount of VG you want in the final mix in percentage
		WantVG			float64
	}

	Recipe struct {
		NicotineBase	float64
		Aroma			float64
		PG				float64
		VG				float64
		Quantity		float64
	}
)

func (i *Inputs) Calculate() (*Recipe, error) {
	var r Recipe
	
	if i.WantNicotine > 0 &&  i.HaveNicotine <= 0 {
		return &r, calculators.ErrNoHaveNicotine
	}

	if i.HaveNicotine > 0 {
		r.NicotineBase = i.WantQuantity * i.WantNicotine / i.HaveNicotine
	} else {
		r.NicotineBase = 0
	}

	r.Quantity = i.WantQuantity
	r.Aroma = i.WantQuantity * i.WantAroma / 100
	r.PG = i.WantQuantity * i.WantPG / 100 - r.NicotineBase - r.Aroma
	r.VG = i.WantQuantity * i.WantVG / 100

	return &r, nil
}
