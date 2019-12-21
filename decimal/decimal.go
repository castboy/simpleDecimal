package decimal

type Decimal struct {
	value int64
	exp   int
}

var (
	zeroDecimal = Decimal{}
	fracPrec    = 18
	pow10       = [...]int64{
		0,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,

		100000000000,
		1000000000000,
		10000000000000,
		100000000000000,
		1000000000000000,
		10000000000000000,
		100000000000000000,
		1000000000000000000,
	}
)

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(value int64, exp int) Decimal {
	return Decimal{
		value: value,
		exp:   exp,
	}
}

func add(d, d2 *Decimal) Decimal {
	if d.exp == d2.exp {
		return Decimal{
			value: d.value + d2.value,
			exp:   d2.exp,
		}
	}

	if d.exp > d2.exp {
		return Decimal{
			value: d2.value + d.value*pow10[d.exp-d2.exp],
			exp:   d2.exp,
		}
	}

	return Decimal{
		value: d.value + d2.value*pow10[d2.exp-d.exp],
		exp:   d.exp,
	}
}

func (d Decimal) Add(d2 Decimal) Decimal {
	return add(&d, &d2)
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	d2.value = -d2.value
	return add(&d, &d2)
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal{
		value: d.value * d2.value,
		exp:   d.exp + d2.exp,
	}
}

func div(d, d2 Decimal, ds *[]Decimal) {
	if d.value > d2.value {
		nd := Decimal{
			value: d.value / d2.value,
			exp:   d.exp - d2.exp,
		}
		*ds = append(*ds, nd)
		if rest := d.value % d2.value; rest != 0 {
			nd := Decimal{
				value: rest,
				exp:   d.exp,
			}
			div(nd, d2, ds)
		}
		return
	} else if d.value == d2.value {
		*ds = append(*ds, Decimal{
			value: 1,
			exp:   d.exp - d2.exp,
		})
		return
	} else {
		n := 0
		for {
			d.value *= 10
			n--
			if d.value > d2.value {
				div(Decimal{
					value: d.value,
					exp:   d.exp + n,
				}, d2, ds)
				return
			}
		}
	}
}

func (d Decimal) Div(d2 Decimal) Decimal {
	ds := []Decimal{}
	div(d, d2, &ds)
	zd := zeroDecimal
	for i := range ds {
		zd = zd.Add(ds[i])
	}
	return zd
}
