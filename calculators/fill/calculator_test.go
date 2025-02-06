package fill

import (
	"math"
	"reflect"
	"testing"

	"github.com/blackwidow-sudo/govape/calculators"
)

func TestFillCalculator(t *testing.T) {
	t.Run("calculate recipe", func(t *testing.T) {
		ingredients := Inputs{
			HaveVpg:      85.75,
			HaveNicotine: 48.0,
			WantNicotine: 3.0,
			WantAroma:    8.0,
		}

		got, err := ingredients.Calculate()
		if err != nil {
			t.Errorf("didn't expect error but got: %s", err)
		}

		want := &Recipe{
			Quantity:     100.0,
			Vpg:          85.75,
			NicotineBase: 6.25,
			Aroma:        8.00,
		}

		assertTolerantEqual(t, got, want, 0.001)
	})

	t.Run("calculate desired nicotine without providing nicotine base", func(t *testing.T) {
		ingredients := Inputs{
			HaveVpg:      85.75,
			HaveNicotine: 0.0,
			WantNicotine: 3.0,
			WantAroma:    8.0,
		}

		_, err := ingredients.Calculate()
		assertError(t, err)
	})
}

func FuzzFillCalculator(f *testing.F) {
	f.Add(50.0, 60.0, 48.0, 3.0)
	f.Add(40.0, 50.0, 36.0, 6.0)
	f.Add(30.0, 40.0, 24.0, 9.0)
	f.Add(0.0, 0.0, 0.0, 0.0)

	f.Fuzz(func(t *testing.T, haveQuantity, haveNicotine, wantAroma, wantNicotine float64) {
		inputs := Inputs{
			HaveVpg:      haveQuantity,
			HaveNicotine: haveNicotine,
			WantAroma:    wantAroma,
			WantNicotine: wantNicotine,
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

func assertTolerantEqual(t testing.TB, got, want *Recipe, tolerance float64) {
	t.Helper()

	val := reflect.ValueOf(got).Elem()
	wantVal := reflect.ValueOf(want).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		wantField := wantVal.Field(i)

		if field.Kind() != reflect.Float64 || wantField.Kind() != reflect.Float64 {
			continue
		}

		v := field.Float()
		w := wantField.Float()
		if math.Abs(v-w) > tolerance {
			t.Errorf("got %+v, want %+v", got, want)
		}
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
