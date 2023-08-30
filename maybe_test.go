package maybe

import "testing"
import "fmt"

// perhaps[0] is always the first significant error of the calacul chain

// func New_mayhaps() Mayhaps { return map[int]any{0: nil} }

func dummy_int_error(ok int) (int, error) {
	var value int
	var error error
	switch ok {
	case 0:
		error = fmt.Errorf("some dummy error")
	default:
		value = 2
		error = nil
	}
	return value, error
}
func Test_basic(t *testing.T) {
	ok_mayhaps := Maybe_that(dummy_int_error(99))
	fmt.Println("with ok value", ok_mayhaps)
	nok_mayhaps := Maybe_that(dummy_int_error(0))
	fmt.Println("with error", nok_mayhaps)
	fmt.Println("function last")
	extracted := Acertain(ok_mayhaps)
	fmt.Println(extracted)
	// extracted_bis := Acertain(nok_mayhaps)
	// fmt.Println(extracted_bis)
	fmt.Println("should be false", Is_error(ok_mayhaps))
	fmt.Println("should be true", Is_error(nok_mayhaps))
	fmt.Println("binding")
	new_maybe := Convey[int, int](ok_mayhaps, dummy_int_error)
	fmt.Println(new_maybe.Value)
	fmt.Println("yyyyooouu")
	func() (int, error) {
		return Relinquish(new_maybe)
	}()
	// Convey(nok_mayhaps, dummy_int_error)
	// fmt.Println(nok_mayhaps)
	// Last[int](ok_mayhaps)

	// fmt,Println(Last(nok_mayhaps))
}
