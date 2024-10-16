package base

import (
	"math"
	"reflect"
	"testing"

	"github.com/blackwidow-sudo/govape/calculators"
)

func TestFullCalculator(t *testing.T) {
	t.Run("calculate recipe", func(t *testing.T) {
		inputs := Inputs{
			HaveNicotine: 48.0,
			WantNicotine: 3.0,
			WantAroma: 8.0,
			WantQuantity: 120.0,
			WantPG: 50.0,
			WantVG: 50.0,
		}

		got, err := inputs.Calculate()
		if err != nil {
			t.Errorf("didn't expect error but got: %s", err)
		}

		want := &Recipe{
			NicotineBase: 7.5,
			Aroma: 9.6,
			PG: 42.9,
			VG: 60.0,
			Quantity: 120.0,
		}

		assertEqual(t, got, want)
	})

	t.Run("calculate desired nicotine without providing nicotine base", func(t *testing.T) {
		inputs := Inputs{
			HaveNicotine: 0.0,
			WantNicotine: 3.0,
			WantAroma: 8.0,
			WantQuantity: 120.0,
			WantPG: 50.0,
			WantVG: 50.0,
		}

		_, err := inputs.Calculate()
		assertError(t, err)
	})
}

func FuzzFullCalculator(f *testing.F) {
    f.Add(0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
    f.Add(1.0, 6.0, 10.0, 100.0, 40.0, 60.0)
    f.Add(2.0, 9.0, 12.0, 80.0, 30.0, 70.0)

    f.Fuzz(func(t *testing.T, haveNicotine, wantNicotine, wantAroma, wantQuantity, wantPG, wantVG float64) {
        inputs := Inputs{
            HaveNicotine:	haveNicotine,
            WantNicotine:	wantNicotine,
            WantAroma:		wantAroma,
            WantQuantity:	wantQuantity,
            WantPG:			wantPG,
            WantVG:			wantVG,
        }

        result, err := inputs.Calculate()
        if err != nil {
			switch err {
			case calculators.ErrNoHaveNicotine:
				break
			default:
            	t.Errorf("didn't expect error but got: %s", err)
			}
        }

		assertNoFloatZeroDivision(t, result)
    })
}

func assertEqual(t testing.TB, got, want *Recipe) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}

func assertNoFloatZeroDivision(t testing.TB, got *Recipe) {
	t.Helper()

	val := reflect.ValueOf(got).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() != reflect.Float64 {
			continue
		}

		v := field.Float()
		if math.IsNaN(v) || math.IsInf(v, 0) {
			t.Error("possible unhandled division by zero found")
		}
	}
}