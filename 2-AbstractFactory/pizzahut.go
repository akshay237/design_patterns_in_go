package main

type PizzaHutFactory struct{}

type PizzaHutPizza struct {
	pizza
}

type PizzaHutBread struct {
	garlicBread
}

func (p *PizzaHutFactory) CreatePizza() PizzSvc {
	return &PizzaHutPizza{
		pizza{
			name:     "Marghrita",
			price:    234.5,
			toppings: []string{"myo", "ayo", "tyo"},
		},
	}
}

func (p *PizzaHutFactory) CreateBread() GarlicBreadSvc {
	return &PizzaHutBread{
		garlicBread{
			name:  "bread1",
			price: 123.3,
		},
	}
}
