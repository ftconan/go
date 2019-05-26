package client

import "testing"
import "ch15/series"

func TestPackage(t *testing.T)  {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
}
