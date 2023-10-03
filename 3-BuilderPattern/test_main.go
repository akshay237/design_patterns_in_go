package main

import "testing"

func TestVeggieSub(t *testing.T) {
	t.Run("test the create veggie sub", func(t *testing.T) {
		veggie := &vegieDelightBuilder{}
		builder := &Builder{
			builder: veggie,
		}
		veggieDelightSub := builder.buildSub()
		describeSub(veggieDelightSub)
	})

}
