package simpleDecimal

import (
	"testing"
	"simpleDecimal/decimal"
	"fmt"
)

func TestAdd(t *testing.T) {
	a := decimal.New(11, 3)
	b := decimal.New(2, 2)
	fmt.Printf("%+v\n", a.Add(b))

	a = decimal.New(11, 2)
	fmt.Printf("%+v\n", a.Add(b))

	a = decimal.New(11, -1)
	fmt.Printf("%+v\n", a.Add(b))
}

func TestSub(t *testing.T) {
	a := decimal.New(11, 3)
	b := decimal.New(2, 2)
	fmt.Printf("%+v\n", a.Sub(b))

	a = decimal.New(11, 2)
	fmt.Printf("%+v\n", a.Sub(b))

	a = decimal.New(11, -1)
	fmt.Printf("%+v\n", a.Sub(b))
}

func TestMul(t *testing.T) {
	a := decimal.New(1, 2)
	b := decimal.New(-2, -3)
	fmt.Printf("%+v", a.Mul(b))
}


func TestDiv(t *testing.T) {
	a := decimal.New(3141592653, -9)
	b := decimal.New(2, 0)
	fmt.Printf("%+v", a.Div(b))
}
