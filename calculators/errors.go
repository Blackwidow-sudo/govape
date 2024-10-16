package calculators

import "errors"

var (
	ErrNoHaveNicotine = errors.New("cannot calculate desired nicotine without used nicotine")
	ErrNoHaveAroma = errors.New("cannot calculate recipe without used aroma")
	ErrNoWantAroma = errors.New("cannot calculate recipe without wanted aroma")
)