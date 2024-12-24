package money

type Money struct {
	Amount   int
	Currency string
}

func New(a int, cur string) Money {
	return Money{
		Amount:   a,
		Currency: cur,
	}
}
