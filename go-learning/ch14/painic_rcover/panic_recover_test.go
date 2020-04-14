package painic_rcover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
