package main

import "fmt"

type Flyer interface {
	Fly()
}

type Plane struct {
	Flyer
}

// super struct
type Animal struct {
	name string // ``
	free bool   // false
}

func (a Animal) Run() {
	fmt.Printf("%s is free\n", a.name)
	a.free = true
}

// child struct 1
type Tiger struct {
	Animal
	powsNum int
}

// child struct 2
type Flamingo struct {
	Animal
	wingsNum int
	name     string
}

func (a Flamingo) Fly() {
	fmt.Printf("Flamingo %s and %s flown\n", a.name, a.Animal.name)
}

// child struct 2
type Sparrow struct {
	Animal
	wingsNum int
	name     string
}

func (a Sparrow) Fly() {
	fmt.Printf("Sparrow %s and %s flown\n", a.name, a.Animal.name)
}

// super method in child method

func main() {
	// creating embedding
	tg := Tiger{Animal{name: `Sam`}, 4}
	fl := Flamingo{Animal{name: `Tim`}, 2, `Tim Junior`}
	sp := Sparrow{Animal{name: `Jim`}, 2, `Tim Junior`}
	// using super method
	tg.Animal.Run()
	fl.Animal.Run()
	// using child method
	tg.Run()
	fl.Run()
	// using common field
	if tg.free {
		fmt.Println(`People are scared`)
	}
	// shadowing of embedding fields

	var birds []Flyer
	birds = append(birds, fl)
	birds = append(birds, sp)

	birdsFlying(birds)
}

func birdsFlying(birds []Flyer) {
	for _, bird := range birds {
		bird.Fly()
	}
}
