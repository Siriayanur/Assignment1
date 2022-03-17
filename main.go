package main

import (
	"fmt"

	controller "github.com/Siriayanur/Nuclei_Assignments/controller"

	itemData "github.com/Siriayanur/Nuclei_Assignments/model"
)

func main() {

	items := []itemData.Item{}
	condition := true

	// correctType := false

	for condition {
		var itemName string
		var itemPrice float64
		var itemQuantity int
		var itemType int
		stop := "y"

		fmt.Print("Item name : ")
		fmt.Scanln(&itemName)
		fmt.Print("Item price : ")
		fmt.Scanln(&itemPrice)
		fmt.Print("Item Quantity : ")
		fmt.Scanln(&itemQuantity)

		fmt.Print("Item Type | 0 - Raw, 1 - Manufactured, 2 - Imported | : ")
		fmt.Scanln(&itemType)
		correctType := controller.ValidateItemType(itemType)

		for !correctType {
			fmt.Println("Enter correct type : ")
			fmt.Scanln(&itemType)
			correctType = controller.ValidateItemType(itemType)
		}

		item := itemData.GetItem(itemName, itemPrice, itemQuantity, itemType)
		controller.CalculateFinalPrice(&item)
		items = append(items, item)

		fmt.Println("Got some more items ? y/n")
		fmt.Scanln(&stop)
		if stop == "n" {
			break
		}
	}
	fmt.Println("| Your order |")

	//Items details
	for i := 0; i < len(items); i++ {
		fmt.Printf("Item : %d | Name : %s | Sales Tax/item : %.4f | Final Prize : %.4f\n", i+1, items[i].Name, items[i].SalesTax, items[i].FinalPrice)
	}
}
