package main

type Leon struct {
	Animal
}

func newLeon() IAnimal {
	return &Leon{
		Animal: Animal{
			name:   "Leon",
			race:   "africano",
			age:    "12",
			danger: 10,
		},
	}
}
