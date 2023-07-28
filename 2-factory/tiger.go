package main

type Tiger struct {
	Animal
}

func newTiger() IAnimal {
	return &Tiger{
		Animal: Animal{
			name:   "Tigre",
			race:   "bengala",
			age:    "23",
			danger: 9,
		},
	}
}
