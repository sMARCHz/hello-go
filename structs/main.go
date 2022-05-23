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

	// Initialization with zero memory allocation
	hero3 := Hero{}
	fmt.Println(&hero3)

	// Initialization with new function (return a pointer to a newly allocated zero value of that type)
	// Also allocate zero memory
	hero4 := new(Hero)
	fmt.Println(hero4)

	// Retrieve value
	fmt.Println(hero2.actorName)

	// Modify value
	hero2.actorName = "Hawkeye"
	fmt.Println(hero2.actorName)

	// Anonymous struct
	hero5 := struct{ name string }{name: "Bruce"}
	fmt.Println(hero5.name)

	// Struct is value type. hero7 is a copy of hero6 so, if hero7 is changed, hero6 remains the same
	hero6 := Hero{number: 4, actorName: "Starlord"}
	hero7 := hero6
	hero7.actorName = "Spiderman"
	fmt.Println(hero6.actorName, hero7.actorName)

	// Same as arrays, we can use & to point hero8 to hero6
	hero8 := &hero6
	hero8.actorName = "Antman"
	fmt.Println(hero6.actorName, hero8.actorName) // actually (*hero8).actorName but compiler help us to simplify syntax

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
	// Example  `json:"name"` -> When we encode it to json, the field's name will be 'name' instead of origin name
	type Animal2 struct {
		name   string `required max:"100"`
		origin string
	}
	animal2 := reflect.TypeOf(Animal2{})
	field, _ := animal2.FieldByName("name")
	fmt.Println(field.Tag)

	// Empty interfaces is a type which you can put value of any types into it (See more in the interfaces section)
	type Empty struct {
		value interface{}
	}
	empty := Empty{value: 10}
	empty2 := Empty{value: "abc"}
	empty3 := Empty{value: Empty{true}}
	fmt.Printf("Empty interfaces: %v, %v, %v\n", empty, empty2, empty3)

	// Method is function of specific struct
	tesla := car{name: "Tesla", speedKPH: 100}
	tesla.accelerate()
	tesla.show()
	fmt.Println(tesla.KPHToMPH())
}

type car struct {
	name     string
	speedKPH float32
}

// Method (Function) is value type so, c is a copy of the struct not the real
// Structure: func (receiver) name(argument) returnType {}
func (c car) show() {
	fmt.Printf("Name=%v SpeedKPH=%v\n", c.name, c.speedKPH)
}

// Method with fixed returnValue
func (c car) KPHToMPH() (mph float32) {
	mph = 0.621 * c.speedKPH
	return // always return mph because we declare it as returnValue
}

// Method's receiver can be both pointer receiver and value receiver
func (c *car) accelerate() {
	c.speedKPH += 20
}
