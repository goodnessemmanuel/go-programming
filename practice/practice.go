// You can edit this code!
// Click here and start typing.
package practice;

import "fmt"

func Practice() {
	user := Person{
		First: "John",
		Last:  "Doe",
		Age:   27,
	}
	user.speak()

	sq := Square{
		Length: 7.00,
		Breath: 7.00,
	}
	info(sq)

	cir := Circle{
		Radius: 3.5,
	}
	info(cir)

	callbackFunc := callBackDemoDoubleMultiplier()
	double2returnTheSquare := callbackFunc(2)
	fmt.Println("Result of callback func is: ", double2returnTheSquare)

	x := func(rand int) int {
		return int(rand / 2)
	}

	actOnAnotherFunc(x)
	//notice on underly that although incrementer, incrementer2 calls thesame function the enclosed variable within this function 
	//that is being called still maintain its scope within and did not pass whatever changes made to it to another new call. It is more like fresh call fresh beginning

	incrementer := demoVariableEnclosure()
	fmt.Println("first scope ", incrementer())
	fmt.Println("first scope", incrementer())
	fmt.Println("first scope", incrementer())
	fmt.Println("first scope", incrementer())

	incrementer2 := demoVariableEnclosure()
	fmt.Println("second scope", incrementer2())
	fmt.Println("second scope", incrementer2())
	fmt.Println("second scope", incrementer2())
	fmt.Println("second scope", incrementer2())

}

func info(any Shape) {
	area := any.area()
	shape := any.ShapeType()

	fmt.Println("This", shape, any, "is", area)
}

type Shape interface {
	ShapeType() string
	area() float64
}

type Square struct {
	Length float64
	Breath float64
}

func (s Square) area() float64 {
	return s.Length * s.Breath
}

func (s Square) ShapeType() string {
	return "Square"
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.142 * (c.Radius * c.Radius)
}

func (c Circle) ShapeType() string {
	return "Circle"
}

type Person struct {
	First string
	Last  string
	Age   int
}

func (u Person) speak() {
	fmt.Println("Name is", u.First, "Age is", u.Age)
}

func callBackDemoDoubleMultiplier() func(a int) int {
	return func(a int) int {
		a += a
		return a * a
	}
}

func actOnAnotherFunc(fn func(in int) int) {
	afterAction := fn(10)
	fmt.Println("**** count down", afterAction, "****")
	i := afterAction
	for i >= 0 {
		fmt.Println("Count: ", i)
		i--
	}
}

func demoVariableEnclosure() func() int {
	//this variable scope is enclosed within this function
	x := 0
	return func() int {
		x++
		return x
	}
}
