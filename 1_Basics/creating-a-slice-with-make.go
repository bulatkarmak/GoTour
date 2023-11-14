// https://go.dev/tour/moretypes/13

package main

type Book struct {
	Title  string
	Author string
	Year   int
}

// func main() {
// 	library := make([]Book, 2, 2)

// 	fmt.Println(len(library), cap(library))

// 	library[0] = Book{"\"It\"", "King", 1984}
// 	fmt.Println(len(library), cap(library))

// 	library[1] = Book{"\"Captain's daughter\"", "Pushkin", 1784}
// 	fmt.Println(len(library), cap(library))

// 	fmt.Println(library)
// }
