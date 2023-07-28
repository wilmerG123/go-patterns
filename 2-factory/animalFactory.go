package main

import (
	"fmt"
)

func getAnimal(animalRace string) (IAnimal, error) {

	if animalRace == "bengala" {
		return newTiger(), nil
	}
	if animalRace == "africano" {
		return newLeon(), nil
	}

	return nil, fmt.Errorf("Animal no encontrado...")
}
