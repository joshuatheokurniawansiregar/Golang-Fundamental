package pkg

import "fmt"

//Go's structs are typed collections of fields. They're useful
//for grouping data together to from records
type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{name: name}
	p.age = age
	return &p
}

func CallStruct() {
	bobby := NewPerson("Bobby", 30)
	var johnson *Person = new(Person)
	johnson.name = "Johnson"
	johnson.age = 45
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 20})
	fmt.Println(bobby.name, bobby.age)
	fmt.Println(johnson)
	p := Person{name: "Sakurajima Mai", age: 19}
	fmt.Println("name: ", p.name, ", age: ", p.age)

	var dog = struct {
		name   string
		isGood bool
	}{}
	dog.name = "brontosaurus"
	dog.isGood = false
	fmt.Println(dog)
}
