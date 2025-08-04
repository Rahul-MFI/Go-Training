package snippets

import (
	"fmt"
)

type Vehicle interface {
	Drive() string
	Rest() string
}

type Car struct {
	Model string
	Sound string
}

type Bus struct {
	Model      string
	Passengers int
	Sound      string
}

func (b Bus) Drive() string {
	return "Busss!"
}

func (b Bus) Rest() string {
	return "Bzzzzz!"
}

func (c Car) Drive() string {
	return "Carrrr!"
}

func (c Car) Rest() string {
	return "Vroom!"
}

func checkVehicle(v Vehicle) {
	if car, ok := v.(Car); ok {
		fmt.Printf("Car model : %s\n", car.Model)
	} else if bus, ok := v.(Bus); ok {
		fmt.Printf("Bus model: %s\n", bus.Model)
	} else {
		fmt.Println("Unknown type")
	}
}

func Seven() {
	v := Car{Model: "Audi", Sound: "Vroom!"}
	b := Bus{Model: "Volvo", Passengers: 50, Sound: "Bzzzzz!"}
	fmt.Println(v)
	fmt.Println(b)
	checkVehicle(v)
	checkVehicle(b)
}
