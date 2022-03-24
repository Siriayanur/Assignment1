package utils

const (
	ITEM_NAME        string  = "Item"
	ITEM_QUANTITY    int     = 1
	ITEM_PRICE       float64 = 0.0
	SALES_TAX        float64 = 0.0
	FINAL_PRICE      float64 = 0.0
	BASE_TAX                 = 0.125
	IMPORT_DUTY              = 0.01
	SURCHARGE_LEVEL1         = 5.0
	SURCHARGE_LEVEL2         = 10.0
	SURCHARGE_LEVEL3         = 0.05
)

const (
	RAW = iota
	IMPORTED
	MANUFACTURED
)
