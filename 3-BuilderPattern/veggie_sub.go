package main

type vegieDelightBuilder struct {
	sub
}

func (v *vegieDelightBuilder) setBread() {
	v.sub.bread = "brown bread"
}

func (v *vegieDelightBuilder) setCheese() {
	v.sub.hasCheese = true
}

func (v *vegieDelightBuilder) setToppings() {
	v.sub.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
}

func (v *vegieDelightBuilder) setSauces() {
	v.sub.hasSauces = []string{"black", "chilly green", "red chilly"}
}

func (v *vegieDelightBuilder) getSub() sub {
	return v.sub
}
