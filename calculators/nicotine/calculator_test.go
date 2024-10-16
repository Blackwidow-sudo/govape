package nicotine

import (
	"math"
	"reflect"
	"testing"

	"github.com/blackwidow-sudo/govape/calculators"
)

func TestNicotineCalculat(t *testing.T) {
	t.Run("calculate recipe", func(t *testing.T) {
		ingredients := Inputs{
			HaveQuantity: 50.0,
			WantQuantity: 60.0,
			HaveNicotine: 48.0,
			WantNicotine: 3.0,
		}

		got, err := ingredients.Calculate()
		if err != nil {
			t.Errorf("didn't expect error but got: %s", err)
		}

		want := &Recipe{
			Quantity: 60.0,
			NicotineBase: 3.75,
			Rest: 6.25,
		}

		assertEqual(t, got, want)
	})

	t.Run("calculate desired nicotine without providing nicotine base", func(t *testing.T) {
		ingredients := Inputs{
			HaveQuantity: 50.0,
			WantQuantity: 60.0,
			HaveNicotine: 0.0,
			WantNicotine: 3.0,
		}

		_, err := ingredients.Calculate()
		assertError(t, err)
	})
}

func FuzzFullCalculator(f *testing.F) {
	f.Add(50.0, 60.0, 48.0, 3.0)
	f.Add(40.0, 50.0, 36.0, 6.0)
	f.Add(30.0, 40.0, 24.0, 9.0)
	f.Add(0.0, 0.0, 0.0, 0.0)

    f.Fuzz(func(t *testing.T, haveQuantity, wantQuantity, haveNicotine, wantNicotine float64) {
        inputs := Inputs{
            HaveQuantity:	haveQuantity,
            WantQuantity: 	wantQuantity,
            HaveNicotine:   haveNicotine,
            WantNicotine: 	wantNicotine,
        }

        results, err := inputs.Calculate()
        if err != nil {
			switch err {
			case calculators.ErrNoHaveNicotine:
				break
			default:
            	t.Errorf("didn't expect error but got: %s", err)
			}
        }

		assertNoFloatZeroDivision(t, results)
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