// https://go.dev/tour/moretypes/11

package main

import "fmt"

type Product struct {
	Name     string
	Quantity int
}

func main() {
	products := []Product{
		{"Apple", 1},
		{"Banana", 2},
		{"Milk", 3},
	}

	fmt.Printf("Длина: %d, Емкость: %d\n", len(products), cap(products))

	products = append(products, Product{"Meat", 4})
	fmt.Printf("Длина: %d, Емкость: %d\n", len(products), cap(products))

	products = AddProduct(products, Product{"Ice Cream", 5})
	fmt.Printf("Длина: %d, Емкость: %d\n", len(products), cap(products))

}

func AddProduct(products []Product, newProduct Product) []Product {
	products = append(products, newProduct)
	return products
}

// функция AddProduct так-то избыточна, но её попросил написать ChatGPT, когда я попросил у него задание для закрепление материала
