package decorator

type normalPizza struct {
}

func (p *normalPizza) getPrice() int {
	return 20
}
