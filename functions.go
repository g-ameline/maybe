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
