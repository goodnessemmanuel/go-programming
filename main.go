package main

import "fmt"

func main() {
	fmt.Println("Welcome to go!")
	workingWithStruct()
}

/**
* A struct is a datastructure with allows you to 
* compose together values of different datatypes.
* It is a composite datatype
**/
type Person struct{
	firstName string;
	lastName string;
}
type SecretAgent struct{
	Person;
	licenseToKill bool;
}

func workingWithStruct()  {
	
	p1 := Person{
		firstName: "James",
		lastName: "Bond",
	}

	p2 := Person{
		firstName: "Goodness",
		lastName: "Emmanuel",
	}

	fmt.Println(p1)
	//print the type
	fmt.Printf("struct p1 type is: %T\n", p1)
	//using dot notation to access a struct
	fmt.Println(p2.firstName, p2.lastName)

	//liveraging the composition of a struct
	sa1 := SecretAgent{
		Person: p1,
		licenseToKill: true,
	}

	sa2 := SecretAgent{
		Person: p2,
		licenseToKill: false,
	}

	
	fmt.Println(sa2, sa1)
}

func New(SecretAgent SecretAgent) {
	panic("unimplemented")
}