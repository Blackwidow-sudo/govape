package aroma

import (
	"github.com/blackwidow-sudo/govape/calculators"
)

type (
	Inputs struct {
		// Amount of aroma you have in milliliters
		HaveAroma float64

		// Percentage of aroma you want in the final mix
		WantAroma float64
	}

	Recipe struct {
		// Total quantity of the liquid
		Quantity	float64

		// How many milliliters of VG/PG to add
		VPG			float64

		// How many milliliters of aroma to add
		Aroma		float64
	}
)

func (i *Inputs) Calculate() (*Recipe, error) {
	var r Recipe

	if i.HaveAroma <= 0 {
		return &r, calculators.ErrNoHaveAroma
	}

	if i.WantAroma <= 0 {
		return &r, calculators.ErrNoWantAroma
	}

	r.Quantity = i.HaveAroma * 100 / i.WantAroma
	r.VPG = r.Quantity - i.HaveAroma
	r.Aroma = i.HaveAroma

	return &r, nil
}
