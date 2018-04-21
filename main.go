package main

import (
	"fmt"
	"reflect"
)

// STRUCTS
type Movie struct {
	Name   string
	Rating float32
}

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

// FUNCTIONS
func firstStruct() {
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)
}

func structInitialization() {
	var n Movie
	fmt.Printf("%+v\n", n)
}

func newStruct() {
	mov := new(Movie)
	mov.Name = "Pulp Fiction"
	mov.Rating = 9.99
	fmt.Println(mov)
}

func shortHand() {
	m := Movie{"Reservoir Dogs", 8.3}
	fmt.Println(m)
}

func preferredCreation() {
	m := Movie{
		Name:   "Star Wars",
		Rating: 9.99,
	}
	fmt.Println(m)
}

func createSuperHero() {
	e := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain Drive",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", e)
}

// this creates a struct with default values
func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

func structEquality() {
	a := NewAlarm("day")
	b := NewAlarm("day")

	if a == b {
		fmt.Println("a and b are the same")
	}
}

func structReflection() {
	a := NewAlarm("day")
	fmt.Println(reflect.TypeOf(a))
}

// MAIN
func main() {
	firstStruct()
	structInitialization()
	newStruct()
	shortHand()
	preferredCreation()
	createSuperHero()
	structEquality()
	structReflection()
}
