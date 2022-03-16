package main

import "fmt"

type Item struct {
	name       string
	price      float64
	quantity   int
	types      string
	salesTax   float64
	finalPrice float64
}

func calculateSurcharge(amount float64) float64 {
	if amount <= 100 {
		return 5
	} else if amount <= 200 {
		return 10
	} else {
		return (5 * amount) / 100
	}
}

//takes in an item and calculates MRP or Final Prize
func (item Item) calculateSalesTax() float64 {

	mrp := 0.0

	//repeatedly calculated result
	intermediateResult := (item.price * 12.5) / 100

	if item.types == "Raw" {
		return intermediateResult
	} else if item.types == "Manufactured" {
		return (intermediateResult) + (2 * (intermediateResult + item.price) / 100)
	} else {
		return ((item.price * 10) / 100) + calculateSurcharge(mrp)
	}
}
func main() {

	//declarations and initialisations
	items := []Item{}
	itemTypeChoice := []string{"Raw", "Manufactured", "Imported"}
	condition := true

	//default values for the first item
	itemName := "name"
	itemPrice := 100.0
	itemQuantity := 1
	itemType := 0
	stop := "y"

	for condition {

		//Taking input parameters for each item
		fmt.Print("Item name : ")
		fmt.Scanln(&itemName)
		fmt.Print("Item price : ")
		fmt.Scanln(&itemPrice)
		fmt.Print("Item Quantity : ")
		fmt.Scanln(&itemQuantity)
		fmt.Print("Item Type | 1 - Raw, 2 - Manufactured, 3 - Imported | : ")
		fmt.Scanln(&itemType)

		//create Item instance with the collected data
		item := Item{itemName, itemPrice, itemQuantity, itemTypeChoice[itemType-1], 0.0, 0.0}

		//Sales Liability Tax per item
		item.salesTax = item.calculateSalesTax()
		//Final Price calculated
		item.finalPrice = item.salesTax * float64(item.quantity)
		//Update the items array
		items = append(items, item)

		//Display details of the item
		fmt.Printf("Item : %s | Price : %g | Quantity : %d | Type : %s \n", item.name, item.price, item.quantity, item.types)

		//Functionality #2
		fmt.Println("Got some more items ? y/n")
		fmt.Scanln(&stop)
		if stop == "n" {
			break
		}
	}
	fmt.Println("********* Your order ************")

	//Logging the details to the user
	for i := 0; i < len(items); i++ {
		fmt.Printf("Item : %d | Name : %s | Sales Tax/item : %.4f | Final Prize : %.4f", i+1, items[i].name, items[i].salesTax, items[i].finalPrice)
	}
}

//TODO
// Type of product -- refactor
