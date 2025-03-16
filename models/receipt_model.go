package models

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" binding:"required,datetime=15:04"`
	Total        string `json:"total" binding:"required,numeric"`
	Items        []item `json:"items" binding:"required"`
}

type item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required,numeric"`
}
