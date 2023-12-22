package maybe

import (
	"fmt"
)

func Somehow[In any, Out any](state_in In, state_in_ok func(In) bool, error_reaction func(error) error, callback any) (Out, error) {
	if !state_in_ok(state_in) {
		return *new(Out), error_reaction(fmt.Errorf(fmt.Sprint(state_in)))
	}
	switch callback.(type) {
	case Out:
		return callback.(Out), nil
	case func() Out:
		return callback.(func() Out)(), nil
	case func() (Out, error):
		output, error_out := callback.(func() (Out, error))()
		if error_out != nil {
			error_out = error_reaction(error_out)
		}
		return output, error_out
	case func() error:
		error_out := callback.(func() error)()
		if error_out != nil {
			error_out = error_reaction(error_out)
		}
		return *new(Out), nil
	case func():
		callback.(func())()
		return *new(Out), nil
	}
	fmt.Printf("Failed type assertion;\nUnderlying Type: %T\n", callback)
	panic(callback)
}

// state_in_checks
func if_nil(err error) bool          { return err == nil }
func if_error(err error) bool        { return err != nil }
func if_ok(true_or_false bool) bool  { return true_or_false }
func if_nok(true_or_false bool) bool { return !true_or_false }

// error_reactions
func must(err error) error { panic_red(err); return err }
func do(err error) error   { return err }
func try(err error) error  { return nil }

// main usge when sticking to happy path
func If_nil_must[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_nil, must, callback)
}
func If_nil_do[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_nil, do, callback)
}
func If_nil_try[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_nil, try, callback)
}

// if need to deal with sad path
func If_error_must[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_error, must, callback)
}
func If_error_do[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_error, do, callback)
}
func If_error_try[In error, Out any](state_in In, callback any) (Out, error) {
	return Somehow[error, Out](state_in, if_error, try, callback)
}

// if working with boolean
func If_ok_must[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_ok, must, callback)
}
func If_ok_do[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_ok, do, callback)
}
func If_ok_try[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_ok, try, callback)
}
func If_nok_must[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_nok, must, callback)
}
func If_nok_do[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_nok, do, callback)
}
func If_nok_try[In bool, Out any](state_in bool, callback any) (Out, error) {
	return Somehow[bool, Out](state_in, if_nok, try, callback)
}
