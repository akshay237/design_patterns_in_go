package main

type subSvc interface {
	setBread()
	setCheese()
	setToppings()
	setSauces()
	getSub() sub
}
