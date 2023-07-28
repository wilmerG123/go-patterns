package main

type Animal struct {
	name   string
	race   string
	age    string
	danger int
}

func (a *Animal) setName(name string) {
	a.name = name
}
func (a *Animal) getName() string {
	return a.name
}
func (a *Animal) setRace(race string) {
	a.race = race
}
func (a *Animal) setAge(age string) {
	a.age = age
}
func (a *Animal) setDanger(danger int) {
	a.danger = danger
}
func (a *Animal) getRace() string {
	return a.race
}
func (a *Animal) getAge() string {
	return a.age
}
func (a *Animal) getDanger() int {
	return a.danger
}
