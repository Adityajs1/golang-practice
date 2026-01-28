package main

import "fmt"

type product struct {
	Name  string
	Price float64
}

func addToCart(cart map[string]int, productName string, quantity int) {
	cart[productName] = cart[productName] + quantity
}

func removeFromCart(cart map[string]int, productName string) error {
	lengthBefore := len(cart)
	delete(cart, productName)
	
	if len(cart) == lengthBefore {
		return fmt.Errorf("product '%s' not found in cart", productName)
	}
	return nil
}

func calculateTotal(cart map[string]int, products map[string]product) float64 {
	total := 0.0
	for productName, quantity := range cart {
		product := products[productName]
		subtotal := product.Price * float64(quantity)
		total += subtotal
	}
	return total
}
func printCart(cart map[string]int) {
	fmt.Println("Shopping Cart:")
	for productName, quantity := range cart {
		fmt.Printf("- %s x %d\n", productName, quantity)
	}
	fmt.Println()
}

func Product() {
	// Create available products
	products := map[string]product{
		"Apple":  {Name: "Apple", Price: 1.50},
		"Banana": {Name: "Banana", Price: 0.75},
		"Orange": {Name: "Orange", Price: 2.00},
		"Grape":  {Name: "Grape", Price: 3.50},
	}

	// Create empty cart
	cart := make(map[string]int)

	// Add items to cart
	addToCart(cart, "Apple", 3)
	addToCart(cart, "Banana", 5)
	addToCart(cart, "Orange", 2)
	addToCart(cart, "Grape", 1)

	// Print cart
	printCart(cart)

	// Remove an item
	err := removeFromCart(cart, "Grape")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item 'Grape' removed from cart")
	}

	// Try removing non-existent item
	err = removeFromCart(cart, "Mango")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	printCart(cart)

	// Calculate total
	total := calculateTotal(cart, products)
	fmt.Printf("Total: $%.2f\n", total)
}
