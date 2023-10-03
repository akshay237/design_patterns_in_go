package main

import (
	"fmt"
	"log"
	"strings"
)

func describePizza(pizza PizzSvc) {
	fmt.Printf("the pizza %s has toppings %s. It costs Rs. %.2f\n", pizza.GetName(), strings.Join(pizza.GetToppings(), ", "), pizza.GetPrice())
}

func describeGarlicBread(garlicBread GarlicBreadSvc) {
	fmt.Printf("the garlic bread, %s costs Rs. %.2f\n", garlicBread.GetName(), garlicBread.GetPrice())
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Dominos or PizzaHut? (D/P)")
	var pizzaType string
	// //pizzaType, _ := reader.ReadString('\n')
	// //pizzaType = strings.Split(pizzaType, "\n")[0]
	_, err := fmt.Scanln(&pizzaType)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Pizza Type: ", pizzaType)

	// pizzaType := "D"

	pizzaFactory, _ := getFactory(pizzaType)

	pizza := pizzaFactory.CreatePizza()
	garlicBread := pizzaFactory.CreateBread()

	describePizza(pizza)
	describeGarlicBread(garlicBread)
}
