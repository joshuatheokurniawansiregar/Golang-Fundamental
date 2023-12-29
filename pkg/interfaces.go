package pkg

import (
	"fmt"
	"math"
)

type Employee interface {
	PrintEmployee(id int64, name string) *Emp
	PrintSalary(basic int64, tax int64) int64
}
type Emp struct {
	id, basic, taxt int64
	name            string
}

type geometry interface {
	Area() float64
	Perim() float64
}

func (e Emp) PrintEmployee(id int64, name string) *Emp {
	var employee = Emp{id: 1, name: name}
	return &employee
}

func (e Emp) PrintSalary(basic int64, taxt int64) int64 {
	return (basic * taxt) / 100
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}

func (r rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func EmployeeFunc(emp Employee) (*Emp, int64) {
	var test_emp = emp.PrintEmployee(1, "Joshua")
	var test_emp_salary int64 = emp.PrintSalary(12, 2)
	return test_emp, test_emp_salary
}
func CallsInterface() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("Rectangle: ")
	Measure(r)
	fmt.Println("Circle: ")
	Measure(c)

	fmt.Println("Employee System")
	var newEmployee = new(Emp)
	fmt.Println(EmployeeFunc(newEmployee))
}

type HasName interface {
	GetName() string
}

type HasNameStruct struct {
	name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

func (hasName *HasNameStruct) GetName() string {
	return hasName.name
}

func CallsInterfaceProgrammerZamanNow() {
	var objHasName *HasNameStruct
	objHasName.name = "Joshua Theo Kurniawan"
	SayHello(objHasName)
}
