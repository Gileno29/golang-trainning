package main

import "fmt"

// to implements an interface in go the struct just need have the same methods (not have special world for this )

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorila := Gorilla{
		Name:          "Camila",
		Color:         "Red",
		NumberOfTeeth: 3,
	}

	PrintInfo(&gorila)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal Says", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "ulala"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
