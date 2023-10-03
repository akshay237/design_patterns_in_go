package main

type chickenDelightBuilder struct {
	sub
}

func (c *chickenDelightBuilder) setBread() {
	c.sub.bread = "brown bread"
}

func (c *chickenDelightBuilder) setCheese() {
	c.sub.hasCheese = false
}

func (c *chickenDelightBuilder) setToppings() {
	c.sub.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
}

func (c *chickenDelightBuilder) setSauces() {
	c.sub.hasSauces = []string{"black", "chilly green", "red chilly"}
}

func (c *chickenDelightBuilder) getSub() sub {
	return c.sub
}
