package main

type Builder struct {
	builder subSvc
}

func (b *Builder) setBuilder(builder subSvc) {
	b.builder = builder
}

func (b *Builder) buildSub() sub {
	b.builder.setBread()
	b.builder.setCheese()
	b.builder.setToppings()
	b.builder.setSauces()

	return b.builder.getSub()
}
