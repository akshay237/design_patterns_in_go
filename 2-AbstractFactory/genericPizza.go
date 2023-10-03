package main

type pizza struct {
	name     string
	toppings []string
	price    float64
}

type PizzSvc interface {
	GetName() string
	GetToppings() []string
	GetPrice() float64
}

func (p *pizza) GetName() string {
	return p.name
}

func (p *pizza) GetPrice() float64 {
	return p.price
}

func (p *pizza) GetToppings() []string {
	return p.toppings
}
