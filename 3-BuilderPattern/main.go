package main

import (
	"fmt"
	"strings"
)

func describeSub(sub sub) {
	fmt.Printf("bread: %s, cheese: %t, toppings: %s, sauces: %s\n", sub.bread, sub.hasCheese, strings.Join(sub.toppings, ", "), strings.Join(sub.hasSauces, ", "))
}

func main() {
	veggie := &vegieDelightBuilder{}
	builder := &Builder{
		builder: veggie,
	}
	veggieDelighSub := builder.buildSub()
	describeSub(veggieDelighSub)

	fmt.Println("------------ Build the Chicken Sub---------------------")
	chicken := &chickenDelightBuilder{}
	chickenBuilder := &Builder{
		builder: chicken,
	}
	chickenDelightSub := chickenBuilder.buildSub()
	describeSub(chickenDelightSub)
}
