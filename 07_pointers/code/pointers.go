// package main

// func main() {
// 	var name string
// 	var namePointer *string

// 	name = "Marilyn Monroe"
// 	namePointer = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name **:", nameValue)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n *string) {
// 	*n = strings.ToUpper(*n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(&name)
// 	fmt.Println(name)
// }

// // ******************************************************

// type Cordinates struct {
// 	X, Y float64
// }

// var c = Cordinates{X: 1, Y: 2}

// func main() {
// 	coordinatesMemoryAddress := &c
// 	coordinatesMemoryAddress.X = 3
// 	fmt.Println(coordinatesMemoryAddress)
// }
