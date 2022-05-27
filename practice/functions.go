package practice


import (
	"fmt"
)

func Functions() {
	fmt.Println("Inside function.go!")

	//demonstrateDefer();
	
	nums := [] int { 1, 2, 3, 4, 5, 6, 8, 9, 10 }
	sumEvenNums(sum, nums...) //observe here that sum is a function
	
}

/**
* A struct is a datastructure with allows you to 
* compose together values of different datatypes.
* It is a composite datatype
**/
type SecretAgent struct{
	Person;
	licenseToKill bool;
}

func workingWithStruct()  {
	
	p1 := Person{
		First: "James",
		Last: "Bond",
	}

	p2 := Person{
		First: "Goodness",
		Last: "Emmanuel",
	}

	fmt.Println(p1)
	//print the type
	fmt.Printf("struct p1 type is: %T\n", p1)
	//using dot notation to access a struct
	fmt.Println(p2.First, p2.Last)

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

func workingWithFunction()  {

	fmt.Println(sum(1, 2, 3, 4, 5, 6)) //outputs 1+2+3+4+5+6 = 21
	
	//unfurling a slice 
	y := []int{7, 8, 9, 10, 11}
	
	//this for instance will not compile because some is expecting variety of separate ints
	//fmt.Println(sum(y)) 

	//but you can unfurl the splice of ints y as shown below to an accepted format
	fmt.Println(sum( y... ))// outputs 7+8+9+10+11 = 45
 	
	fmt.Println(sum()) // outputs 0
}

/**
* function sum accepts a variadic parameters
* It accepts 0 or values of int hence you can
* call sum with no parameter at all and it will return 0
* i.e sum() => 0 see driver code above
**/
func sum(x ...int) int{
	sum := 0;
	
	for _, v := range x {
		sum += v;
	}

	return sum;
	
}

func demonstrateDefer(){
	defer foo(); //here I am defering foo function call
	bar()
	soak()
}

func foo() {
	fmt.Println("I am foo, the first function that was called! but because of defer keyword attached to my name see where I landed")
}
func soak() {
	fmt.Println("====Inside soak function=======")
}

func bar() {
	fmt.Println("******Inside bar function******");
}

func demonstrateCallbackFunc(){

}

func sumEvenNums( x func(val ...int) int,  y ...int){
	sliceOfEvenNums := []int{};
	for _, num  := range y {
		if(num % 2 == 0)    {
			sliceOfEvenNums =	append(sliceOfEvenNums, num)
		}
	}
	fmt.Println(" Even numbers in ", y, "\n are: ", sliceOfEvenNums)
	evenSum   := x( sliceOfEvenNums... );
	
	fmt.Println("Even sum: ", evenSum);
}

