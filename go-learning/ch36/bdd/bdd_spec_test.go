package bdd

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSpec(t *testing.T)  {
	Convey("Give two even numbers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c % 2, ShouldEqual, 0)
			})
		})
	})
}
