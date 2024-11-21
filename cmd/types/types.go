package types

import "time"

type ReceiptStore interface {
	GetReceiptById(key string) (Receipt, error)
	CreateReceipt (receipt Receipt) (string, error)
}

type ReceiptService interface {
	GetReceiptPointsById(id string) (float64, error)
	CreateReceipt(receipt Receipt) (string, error)
}

type Receipt struct {
	Retailer    string   `json:"retailer"`
	PurchasedOn time.Time `json:"purchaseDate"`
	Items       []Item   `json:"items"`
	Total       float64   `json:"total"`
}
type RegisterReceiptPayload struct {
	Retailer     string `json:"retailer" validate:"omitempty"`
	PurchaseDate string `json:"purchaseDate" validate:"omitempty,date_format"`
	PurchaseTime string `json:"purchaseTime" validate:"omitempty,time_format"`
	Items        []Item `json:"items" validate:"omitempty,dive"`
	Total        string `json:"total" validate:"omitempty,float_format"`
}

type RegisterReceiptResponse struct {
	ID string `json:"id"`
}

type RetrievePointsResponse struct {
	Points string `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"omitempty"`
	Price            string `json:"price" validate:"omitempty,float_format"`
}