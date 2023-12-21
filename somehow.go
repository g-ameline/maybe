package maybe

import (
	"fmt"
)

func Somehow_must[Out any](error_in error, somwhow func(error) bool, somewhat any) Out {
	if !somwhow(error_in) {
		panic_red(error_in)
	}
	switch somewhat.(type) {
	case Out:
		return somewhat.(Out)
	case func() Out:
		return somewhat.(func() Out)()
	case func() (Out, error):
		output, error_out := somewhat.(func() (Out, error))()
		Panic(error_out)
		return output
	case func():
		somewhat.(func())()
		return *new(Out)
	}
	fmt.Printf("Failed type assertion;\nUnderlying Type: %T\n", somewhat)
	panic(somewhat)
}
func Somehow_do[Out any](error_in error, somwhow func(error) bool, somewhat any) (Out, error) {
	if !somwhow(error_in) {
		return *new(Out), error_in
	}
	switch somewhat.(type) {
	case Out:
		return somewhat.(Out), error(nil)
	case func() Out:
		return somewhat.(func() Out)(), error(nil)
	case func() (Out, error):
		return somewhat.(func() (Out, error))()
	case func():
		somewhat.(func())()
		return *new(Out), error(nil)
	}
	fmt.Printf("Failed type assertion;\nUnderlying Type: %T\n", somewhat)
	panic(somewhat)
}
func Somehow_try[Out any](error_in error, somwhow func(error) bool, somewhat any) (Out, error) {
	if !somwhow(error_in) {
		return *new(Out), error_in
	}
	switch somewhat.(type) {
	case Out:
		return somewhat.(Out), error(nil)
	case func() Out:
		return somewhat.(func() Out)(), error(nil)
	case func() (Out, error):
		output, _ := somewhat.(func() (Out, error))()
		return output, error_in
	case func():
		somewhat.(func())()
		return *new(Out), error_in
	}
	fmt.Printf("Failed type assertion;\nUnderlying Type: %T\n", somewhat)
	panic(somewhat)
}
func If_nil_must[Out any](wrength error, something any) Out {
	return Somehow_must[Out](
		wrength,
		func(err error) bool { return err == nil },
		something,
	)
}
func If_error_must[Out any](wrength error, something any) Out {
	return Somehow_must[Out](
		wrength,
		func(err error) bool { return err != nil },
		something,
	)
}
func If_nil_do[Out any](wrength error, something any) (Out, error) {
	return Somehow_do[Out](
		wrength,
		func(err error) bool { return err == nil },
		something,
	)
}
func If_error_do[Out any](wrength error, something any) (Out, error) {
	return Somehow_do[Out](
		wrength,
		func(err error) bool { return err != nil },
		something,
	)
}
func If_nil_try[Out any](wrength error, something any) (Out, error) {
	return Somehow_try[Out](
		wrength,
		func(err error) bool { return err == nil },
		something,
	)
}
func If_error_try[Out any](wrength error, something any) (Out, error) {
	return Somehow_try[Out](
		wrength,
		func(err error) bool { return err != nil },
		something,
	)
}

// for map stuff
func If_ok_must[Out any](ok_or_nok bool, something any) Out {
	return Somehow_must[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
func If_nok_must[Out any](ok_or_nok bool, something any) Out {
	return Somehow_must[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
func If_ok_do[Out any](ok_or_nok bool, something any) (Out, error) {
	return Somehow_do[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
func If_nok_do[Out any](ok_or_nok bool, something any) (Out, error) {
	return Somehow_do[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
func If_ok_try[Out any](ok_or_nok bool, something any) (Out, error) {
	return Somehow_try[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
func If_nok_try[Out any](ok_or_nok bool, something any) (Out, error) {
	return Somehow_try[Out](
		nil,
		func(err error) bool { return ok_or_nok },
		something,
	)
}
