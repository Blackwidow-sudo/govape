package nicotine

import (
	"github.com/blackwidow-sudo/govape/calculators"
)

type (
	Inputs struct {
		HaveQuantity    float64
		WantQuantity	float64
		HaveNicotine    float64
		WantNicotine	float64
	}

	Recipe struct {
		Quantity     	float64
		NicotineBase 	float64
		Rest			float64
	}
)

func (i *Inputs) Calculate() (*Recipe, error) {
	var r Recipe

	if i.HaveNicotine <= 0 {
		return &r, calculators.ErrNoHaveNicotine
	}

	r.Quantity = i.WantQuantity
	r.NicotineBase = i.WantQuantity * i.WantNicotine / i.HaveNicotine
	r.Rest = r.Quantity - r.NicotineBase - i.HaveQuantity

	return &r, nil
}
