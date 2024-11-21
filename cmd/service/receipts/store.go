package receipts

import (
	"fmt"
	"strconv"
	"takehome/cmd/types"
)

var receipts = make(map[string]types.Receipt)

const maxReceipts = "10000000"

type ReceiptStore struct {
}

func NewStore() *ReceiptStore {
	return &ReceiptStore{}
}

func (rs *ReceiptStore) GetReceiptById(key string) (types.Receipt, error) {
	value, exists := receipts[key]

	if !exists {
		return types.Receipt{}, fmt.Errorf("receipt does not exist")
	}

	return value, nil
}

func (rs *ReceiptStore) CreateReceipt(receipt types.Receipt) (string, error) {
	id := strconv.Itoa(len(receipts))

	if id == maxReceipts {
		return "", fmt.Errorf("max number of receipts exceeded")
	}
	receipts[id] = receipt

	return id, nil
}
