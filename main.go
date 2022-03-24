package main

import (
	"fmt"

	"github.com/Siriayanur/Nuclei_Assignments/controller"
	"github.com/Siriayanur/Nuclei_Assignments/model"
)

func main() {

	items := []model.Item{}

	for {
		var itemName string
		var itemPrice string
		var itemQuantity string
		var itemType string
		stop := "y"

		fmt.Print("Item name : ")
		fmt.Scanln(&itemName)
		fmt.Print("Item price : ")
		fmt.Scanln(&itemPrice)
		fmt.Print("Item Quantity : ")
		fmt.Scanln(&itemQuantity)

		fmt.Print("Item Type | 1 - Raw, 2 - Manufactured, 3 - Imported | : ")
		fmt.Scanln(&itemType)

		item, itemErrors := model.GetItem(itemName, itemPrice, itemQuantity, itemType)
		if itemErrors != "" {
			item.FinalPrice = 0.0
			item.SalesTax = 0.0
			item.Price = 0.0
			fmt.Println("| Error : |\n\n", itemErrors)
		}
		//Calculation of Tax and Final Price
		controller.CalculateFinalPrice(&item)

		items = append(items, item)

		fmt.Println("Got some more items ? y/n")
		fmt.Scanln(&stop)
		if stop != "y" && stop != "yes" {
			break
		}
	}
	fmt.Println("| Your order |")

	for i := 0; i < len(items); i++ {
		fmt.Printf("Item : %d | Name : %s | Type : %s | Sales Tax/item : %.4f | Final Prize : %.4f\n", i+1, items[i].Name, items[i].GetItemType(), items[i].SalesTax, items[i].FinalPrice)
	}
}
