package maybe

import "fmt"

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

func Must(err error, message ...string) {
	if err != nil {
		panic_red(err, message)
	}
}
func Warn(err error, message ...string) {
	if err != nil {
		print_yellow(err, message)
	}
}
func Rename_error(err error, message string) error {
	if err != nil {
		return fmt.Errorf(message)
	}
	return err
}
func ok_to_err(ok_or_not bool, message string) error {
	if ok_or_not {
		return fmt.Errorf(message)
	}
	return error(nil)
}
func nok_to_err(ok_or_not bool, message string) error {
	if !ok_or_not {
		return fmt.Errorf(message)
	}
	return error(nil)
}
