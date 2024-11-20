package models

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Receipt struct {
	ID           string    `json:"id"`
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime time.Time `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        string    `json:"total"`
}

type Item struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (r *Receipt) CalculateTotal() {
	var total float64
	for _, item := range r.Items {
		total += item.Price
	}

	r.Total = fmt.Sprintf("%.2f", total)
}

var (
	receiptStorage = make(map[string]Receipt)
	mu             sync.Mutex
)

func SaveReceipt(receipt Receipt) {
	mu.Lock()
	defer mu.Unlock()
	receiptStorage[receipt.ID] = receipt
	fmt.Printf("Receipt saved: %v\n", receipt)
}

func GetReceiptByID(id string) (Receipt, error) {
	mu.Lock()
	defer mu.Unlock()
	receipt, exists := receiptStorage[id]
	if !exists {
		return Receipt{}, errors.New("receipt not found")
	}
	fmt.Printf("Receipt lookup for ID %s: found = %v\n", id, exists)
	return receipt, nil
}
