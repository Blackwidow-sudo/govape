package aroma

import (
	"math"
	"reflect"
	"testing"

	"github.com/blackwidow-sudo/govape/calculators"
)

func TestAromaCalculator(t *testing.T) {
	t.Run("calculate recipe", func(t *testing.T) {
		inputs := Inputs{
			HaveAroma: 9.6,
			WantAroma: 8.0,
		}

		recipe, err := inputs.Calculate()
		if err != nil {
			t.Errorf("didn't expect error but got: %s", err)
		}

		want := &Recipe{
			Quantity: 120.0,
			VPG: 110.4,
			Aroma: 9.6,
		}

		assertEqual(t, recipe, want)
	})

	t.Run("calculate desired aroma without providing aroma base", func(t *testing.T) {
		inputs := Inputs{
			HaveAroma: 0.0,
			WantAroma: 8.0,
		}

		_, err := inputs.Calculate()
		assertError(t, err)
	})
}

func FuzzAromaCalculator(f *testing.F) {
	f.Add(0.0, 1.0)
	f.Add(1.0, 0.0)
	f.Add(0.0, 0.0)
	f.Add(1.0, 1.0)

    f.Fuzz(func(t *testing.T, haveAroma, wantAroma float64) {
        inputs := Inputs{
			HaveAroma: haveAroma,
			WantAroma: wantAroma,
        }

        result, err := inputs.Calculate()
        if err != nil {
			switch err {
			case calculators.ErrNoHaveAroma:
			case calculators.ErrNoWantAroma:
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