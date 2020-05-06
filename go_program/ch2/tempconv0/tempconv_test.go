// Author: magician
// File:   tempconv_test.go
// Date:   2020/5/6
package tempconv

import (
	"testing"
)

func TestExampleOne(t *testing.T) {
	{
		t.Logf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		t.Logf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	}
	/*
		t.Logf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
	*/
}

func TestExampleTwo(t *testing.T) {
	c := FToC(212.0)
	t.Log(c.String()) // "100°C"
	t.Logf("%v\n", c) // "100°C"; no need to call String explicitly
	t.Logf("%s\n", c) // "100°C"
	t.Log(c)          // "100°C"
	t.Logf("%g\n", c) // "100"; does not call String
}
