package main

import "fmt"

func main() {
	fmt.Println("Welcome to go!")
	//slices()
	//handsOnExecise()
	workWithMap()
}

/**
 * Arrays are the building block of slices. It is recommended
 * to use slices instead of array
 */
func array() {
	var items [3]int //notice the type here is [3]int. item can also be populated on declaration as in 'var items := [3]int{10, 20, 30}'
	items[0] = 10
	items[1] = 20
	items[2] = 30

	fmt.Println(items)
	fmt.Printf("%T\n", items)
}

func slices() {
	sampleSlice := []int{10, 20, 30, 40}
	//print all the slice elements
	fmt.Println(sampleSlice)
	//target ranges to be printed
	fmt.Printf("Print item at index 1 to 2\n%v\n", sampleSlice[1:3]) //outputs [20 30]

	for i := 0; i < len(sampleSlice); i++ {
		fmt.Println(i, sampleSlice[i])
	}

	fmt.Println("Using range for slice data")
	for i, v := range sampleSlice {
		fmt.Println(i, v)
	}
	//print only the slice values and not the index
	for _, v := range sampleSlice {
		fmt.Println(v)
	}
}

func handsOnExecise() {
	//splice of type int
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[0:5])  //can also be x[:5]
	fmt.Println(x[5:10]) //can also be x[5:]
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	//appending to a slice
	fmt.Println("Append 52 to:", x)
	x = append(x, 52)
	fmt.Println(x)

	fmt.Println("Append 53, 54, 55 to:\n", x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	//appending another slice to existing slice
	y := []int{56, 57, 58, 59, 60}
	fmt.Println("Append", y, "to", x)
	x = append(x, y...)
	fmt.Println(x)

	//deleting from a slice
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("deleting 45, 46, 47 from", x)
	x = append(x[0:3], x[6:]...) //remove 45, 46, 47
	fmt.Println(x)               //outputs [42 43 44 48 49 50 51]

	//creating a slice with the make keyword
	sliceWithMake := make([]string, 5, 5) //last two parameters is initial capacity, size
	states := []string{"Rivers", "Benue", "Bayelsa", "Abuja", "Lagos"}

	for i, v := range states {
		sliceWithMake[i] = v
	}
	fmt.Println("\n", sliceWithMake)

	//multi dimensional slice
	stringSlice1 := []string{"james", "bond", "gifted not twisted"}
	stringSlice2 := []string{"God's grace", "found", "me"}
	fmt.Println("printing multi dimensional slice")
	multiSlice := [][]string{stringSlice1, stringSlice2}
	for _, v := range multiSlice {
		fmt.Println(v)
	}

}

func workWithMap()  {
	maplNameToFavourites := map[string]string{
		"bond_james": "Shaken, not starred",
		"Martinis": `"Women", 'moneypenny_miss'`,
		"James Bond" : `"Literature", "Computer Science", 'Ice cream'`,
	}

	fmt.Println(maplNameToFavourites)

	maplNameToFavourites["oche"] = "God's grace lifted me"
	//range through map and print all valuee
	for key, val := range maplNameToFavourites{
		fmt.Printf("key => %s, value => %s\n", key, val)
	}

	fmt.Printf("%T\n", maplNameToFavourites) //outputs map type which is: map[string]string

	//declaring map to a slice of string. Notice also the use of backtiks in the string keys
	m := map[string][]string{
		"bond_james" : []string{"Shaken, not starred", "Martini", "Women"},
		`moneypenny_miss` : []string{"Shaken, not starred", "Martini", "Women"},
		`no_dr`: []string{"Being good", `Ice Cream`, "Sunsets"},
	} 
	//adding a record to an existing known map
	m["oche"] = []string{`Overcomer`, `Goodness`, `Dominion`, `Gratitude`}	

	//deleting a record from a map
	delete(m, `bond_james`)
	
	for k, v := range m{
		fmt.Printf("%s : %v \n", k, v);
		//range through each key records
		fmt.Println("Record for", k)
		for i, rec := range v{
			fmt.Println(i, rec)
		}
	}
	fmt.Printf("%T\n", m) //outputs map type which is: map[string][]string
}
