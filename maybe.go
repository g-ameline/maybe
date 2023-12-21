package maybe

import (
	"fmt"
)

const red = "\033[31m"
const yellow = "\033[33m"
const green = "\033[32m"
const reset = "\033[0m"

type Maybe[a any] struct {
	Value a
	Error error
}

func Mayhaps[a any](val a, err error) Maybe[a] {
	return Maybe[a]{
		Value: val,
		Error: err,
	}
}

func Bind_i_o_e[in, out any](previous Maybe[in], some_function func(in) (out, error)) Maybe[out] {
	if previous.Error != nil {
		return Maybe[out]{Error: previous.Error}
	}
	return Mayhaps(some_function(previous.Value))
}
func Bind_x_o_e[in, out any](previous Maybe[in], some_function func() (out, error)) Maybe[out] {
	if previous.Error != nil {
		return Maybe[out]{Error: previous.Error}
	}
	return Mayhaps(some_function())
}
func Bind_i_o_x[in, out any](previous Maybe[in], some_function func() out) Maybe[out] {
	if previous.Error != nil {
		return Maybe[out]{Error: previous.Error}
	}
	return Maybe[out]{Value: some_function()}
}
func Bind_i_x_e[in any](previous Maybe[in], some_function func() error) Maybe[in] {
	if previous.Error != nil {
		return Maybe[in]{Error: previous.Error}
	}
	return Maybe[in]{Value: *new(in), Error: some_function()}
}
func Bind_x_x_e[in any](previous Maybe[in], some_function func() error) Maybe[in] {
	if previous.Error != nil {
		return Maybe[in]{Error: previous.Error}
	}
	return Maybe[in]{Value: previous.Value, Error: nil}
}
func Bind_x_o_x[in, out any](previous Maybe[in], some_function func() out) Maybe[out] {
	if previous.Error != nil {
		return Maybe[out]{Error: previous.Error}
	}
	return Maybe[out]{Value: some_function()}
}
func Bind_x_x_x[in any](previous Maybe[in], some_function func()) Maybe[in] {
	if previous.Error != nil {
		return Maybe[in]{Error: previous.Error}
	}
	some_function()
	return previous
}
func Convey[a, b any](previous Maybe[a], something any) Maybe[b] {
	if previous.Error != nil {
		return Maybe[b]{Error: previous.Error}
	}
	switch something.(type) {
	case error:
		return Mayhaps(something.(b), something.(error))
	case b:
		return Mayhaps(something.(b), nil)
	case func(a) (b, error):
		f := something.(func(a) (b, error))
		return Mayhaps(f(previous.Value))
	case func() (b, error):
		f := something.(func() (b, error))
		return Mayhaps(f())
	case func(a) b:
		f := something.(func(a) b)
		return Mayhaps(f(previous.Value), previous.Error)
	case func() b:
		f := something.(func() b)
		return Mayhaps(f(), previous.Error)
	case func() error:
		f := something.(func() error)
		return Mayhaps(*new(b), f())
	case func():
		something.(func())()
		return Mayhaps(*new(b), nil)
	}
	fmt.Printf("Underlying Type: %T\n", something)
	return Maybe[b]{Error: fmt.Errorf("badly fail at func type assertion when ligating")}
}
func Relinquish[a any](m Maybe[a]) (a, error) {
	if m.Error == nil {
		return m.Value, m.Error
	}
	return *new(a), m.Error
}
func (m Maybe[a]) Relinquish() (a, error) {
	return Relinquish(m)
}

func panic_red(things ...any) {
	fmt.Println(red, things, reset)
	panic("")
}
func panic_green(things ...any) {
	fmt.Println(green, things, reset)
	panic("")
}
func print_red(things ...any) {
	fmt.Println(red, things, reset)
}
func print_yellow(things ...any) {
	fmt.Println(yellow, things, reset)
}
func print_green(things ...any) {
	fmt.Println(green, things, reset)
}

func (m Maybe[a]) Panic(message string) Maybe[a] {
	Panic(m.Error)
	return m
}

func (m Maybe[a]) Ascertain() a {
	if m.Error != nil {
		panic_red(m.Error)
	}
	return m.Value
}

func (m Maybe[a]) Print(message ...string) Maybe[a] {
	if m.Error != nil {
		print_red(message, m.Error)
	}
	print_green(message, m.Value)
	return m
}

func (m Maybe[a]) Is_error() bool {
	if m.Error != nil {
		return true
	}
	return false
}

func (m Maybe[a]) Replace_error(message string) Maybe[a] {
	if m.Error != nil {
		return Maybe[a]{Error: fmt.Errorf(message)}
	}
	return m
}

func Panic(err error, message ...string) {
	if err != nil {
		panic_red(err, message)
	}
}
func Warn(err error, message ...string) {
	if err != nil {
		print_yellow(err, message)
	}
}
