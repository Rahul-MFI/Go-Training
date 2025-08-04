package snippets

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type Human struct {
	Name string
}

type Cow struct {
	Name string
}

func (h Human) Greet() string {
	return "Hello from human!"
}

func (c Cow) Greet() string {
	return "Moooooo!"
}

func Six() {
	h := Human{"Rocky"}
	fmt.Println(h.Greet())
	c := Cow{"Cow"}
	fmt.Println(c.Greet())
}
