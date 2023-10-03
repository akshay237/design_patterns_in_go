package main

type dominosFactory struct{}

type dominosPizza struct {
	pizza
}

type dominosGarlicBread struct {
	garlicBread
}

func (d *dominosFactory) CreatePizza() PizzSvc {
	return &dominosPizza{
		pizza{
			name:     "farm house",
			price:    269.0,
			toppings: []string{"t1", "t2"},
		},
	}
}

func (d *dominosFactory) CreateBread() GarlicBreadSvc {
	return &dominosGarlicBread{
		garlicBread{
			name:  "bread b2",
			price: 119.0,
		},
	}
}
