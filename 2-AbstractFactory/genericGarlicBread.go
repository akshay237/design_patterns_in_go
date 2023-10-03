package main

type garlicBread struct {
	name  string
	price float64
}

type GarlicBreadSvc interface {
	GetName() string
	GetPrice() float64
}

func (g *garlicBread) GetName() string {
	return g.name
}

func (g *garlicBread) GetPrice() float64 {
	return g.price
}
