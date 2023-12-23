package maybe

import (
	"fmt"
	"testing"
)

func Test_Somehow(t *testing.T) {
	fmt.Println("\nif ... do")
	fmt.Println("should be default + error")
	number, err := If_nil_do[int](error(nil), func() (int, error) { return 2, fmt.Errorf("some_error") })
	fmt.Println("value :", number, "error :", err)
	fmt.Println("should be default + input error")
	number, err = If_nil_do[int](fmt.Errorf("input error"), func() (int, error) { return 2, fmt.Errorf("output error") })
	fmt.Println("value :", number, "error :", err)
	fmt.Println("should be value 2 + nil")
	number, err = If_nil_do[int](error(nil), func() (int, error) { return 2, error(nil) })
	fmt.Println("value :", number, "error :", err)

	fmt.Println("\nif ... must")
	// fmt.Println("should panic")
	// number, err = If_nil_must[int](error(nil), func() (int, error) { return 2, fmt.Errorf("some_error") })
	// fmt.Println("value :", number, "error :", err)
	fmt.Println("should not panic, and return default + input error")
	number, err = If_nil_must[int](fmt.Errorf("input error"), func() (int, error) { return 2, fmt.Errorf("output error") })
	fmt.Println("value :", number, "error :", err)
	fmt.Println("should be value 2 + nil")
	number, err = If_nil_must[int](error(nil), func() (int, error) { return 2, error(nil) })
	fmt.Println("value :", number, "error :", err)
}
