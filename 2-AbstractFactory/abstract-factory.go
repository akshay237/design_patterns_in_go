package main

import "errors"

var ErrInvalidBrand = errors.New("BRAND_NOT_EXIST")

type PizzaFactory interface {
	CreatePizza() PizzSvc
	CreateBread() GarlicBreadSvc
}

func getFactory(chain string) (PizzaFactory, error) {
	if chain == "P" {
		return &PizzaHutFactory{}, nil
	}
	if chain == "D" {
		return &dominosFactory{}, nil
	}
	return nil, ErrInvalidBrand
}
