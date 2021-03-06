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

		item := model.GetItem(itemName, itemPrice, itemQuantity, itemType)
		if item.Error != "" {
			fmt.Println("Error : \n", item.Error)
			return
		}
		//Calculation of Tax and Final Price
		controller.CalculateFinalPrice(&item)

		items = append(items, item)

		fmt.Println("Got some more items ? y/n")
		fmt.Scanln(&stop)
		if stop == "n" {
			break
		}
	}
	fmt.Println("| Your order |")

	for i := 0; i < len(items); i++ {
		fmt.Printf("Item : %d | Name : %s | Type : %s | Sales Tax/item : %.4f | Final Prize : %.4f\n", i+1, items[i].Name, items[i].GetItemType(), items[i].SalesTax, items[i].FinalPrice)
	}
}
