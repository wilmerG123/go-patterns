package main

type IAnimal interface {
	setName(name string)
	getName() string
	setRace(raza string)
	setAge(age string)
	setDanger(danger int)
	getRace() string
	getAge() string
	getDanger() int
}
