package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// ----Map----

	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"Texas": 1234,
		"NY":    3210,
		"Ohio":  551,
	}

	statePopulation["Georgia"] = 1101123
	delete(statePopulation, "Georgia")

	_, ok := statePopulation["Georgia"] // check if present

	fmt.Println(statePopulation)
	fmt.Println(statePopulation["NY"])
	fmt.Println(ok)

	// ----Struct----

	aDoctor := Doctor{
		number:    3,
		actorName: "Josh Doe",
		companions: []string{
			"Ion Vasile",
			"Maria Ioana",
			"Petre",
		},
	}

	fmt.Println(aDoctor.actorName)

	// --- Embedding ---
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(b)
	fmt.Println(field.Tag)
}
