package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Structure or struct is a user-defined type that allows to group/combine items of possibly different types into a single type
	// Similar to Object in OOP
	// Zero value = nil

	// Declare struct
	type Hero struct {
		number     int
		actorName  string
		companions []string
	}

	// Initialization (same order as struct)
	var hero1 Hero = Hero{1, "Tony", []string{"Steve", "Thor"}}
	fmt.Println(hero1)

	// Initialization with named fields
	hero2 := Hero{number: 2, actorName: "Wanda", companions: []string{"Vision"}}
	fmt.Println(hero2)

	// Retrieve value
	fmt.Println(hero2.actorName)

	// Modify value
	hero2.actorName = "Hawkeye"
	fmt.Println(hero2.actorName)

	// Anonymous struct
	hero3 := struct{ name string }{name: "Bruce"}
	fmt.Println(hero3.name)

	// Struct is value type. hero5 is a copy of hero4 so, if hero5 is changed, hero4 remains the same
	hero4 := Hero{number: 4, actorName: "Starlord"}
	hero5 := hero4
	hero5.actorName = "Spiderman"
	fmt.Println(hero4.actorName, hero5.actorName)
	// Same as arrays, we can use & to point hero6 to the same value as hero4
	hero6 := &hero4
	hero6.actorName = "Antman"
	fmt.Println(hero4.actorName, hero6.actorName) // actually (*hero6).actorName but compiler help us to simplify syntax

	// Struct has no inheritance relation but it has composition relation
	type Animal struct {
		name   string
		origin string
	}
	type Bird struct {
		Animal
		speedKPH float32
		canFly   bool
	}
	bird := Bird{}
	bird.name = "Emu"
	bird.origin = "Australia"
	bird.speedKPH = 48
	bird.canFly = false
	fmt.Println(bird)
	fmt.Println(bird.Animal) // get Animal struct

	// A tag for a field allows you to attach meta-information to the field which can be acquired using reflection.
	// Usually it is used to provide transformation info on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.
	// Example -> `json:"name"`
	type Animal2 struct {
		name   string `required max:"100"`
		origin string
	}
	animal2 := reflect.TypeOf(Animal2{})
	field, _ := animal2.FieldByName("name")
	fmt.Println(field.Tag)
}