package receipts

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"takehome/cmd/types"
	"time"
	"unicode"
)

var ErrReceiptNotFound = errors.New("receipt not found")

type ReceiptService struct {
	store types.ReceiptStore
}

func NewService(rs types.ReceiptStore) *ReceiptService {
	return &ReceiptService{store: rs}
}

func computePointsAlphanumeric(s string) float64 {
	points := 0.0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points++
		}
	}
	return points
}

// If we take the remainder of total % 1, we know the number is round if we get 0
func computePointsIfRoundTotal(total float64) float64 {
	if math.Mod(total, 1) == 0 {
		return 50
	}

	return 0
}

// Since 1/.25 = 4, we can multiply the total by 4 and see if it's a whole number.
// DOES NOT ACCOUNT FOR OVERFLOW
func computePointsTotalQuarter(total float64) float64 {
	if math.Mod(total*4, 1) == 0 {
		return 25.0
	}

	return 0.0
}

func computePointsForItems(items []types.Item) float64 {
	points := 0.0

	for _, item := range items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += math.Ceil(price * 0.2)
		}
	}
	return points
}

func computePointsIfDayIsOdd(purchasedOn time.Time) float64 {
	day := purchasedOn.Day()
	if day%2 != 0 {
		return 6
	}

	return 0
}

func computePointsIfWithinThreshold(purchasedOn time.Time) float64 {
	hour := purchasedOn.Hour()
	minute := purchasedOn.Minute()

	if (hour > 14 || (hour == 14 && minute >= 0)) && (hour < 16) {
		return 10
	}

	return 0
}

func (rserv *ReceiptService) GetReceiptPointsById(id string) (float64, error) {

	// retrieve receipt
	receipt, err := rserv.store.GetReceiptById(id)
	if err != nil {
		return 0, ErrReceiptNotFound
	}

	points := 0.0
	// one point for each alphanumeric character
	points += computePointsAlphanumeric(receipt.Retailer)
	// 50 points if round number
	points += computePointsIfRoundTotal(receipt.Total)
	// 25 points if total is multiple of 0.25
	points += computePointsTotalQuarter(receipt.Total)
	// 5 points for every two items on receipt
	points += float64(5 * (len(receipt.Items) / 2))
	// if trimmed length of item is multiple of 3, multiply price by 0.2 and round up to integer
	points += computePointsForItems(receipt.Items)
	// 6 points if day is odd
	points += computePointsIfDayIsOdd(receipt.PurchasedOn)
	// 10 points if time is between 2PM and 4 PM
	points += computePointsIfWithinThreshold(receipt.PurchasedOn)

	return points, nil
}

func (rserv *ReceiptService) CreateReceipt(receipt types.Receipt) (string, error) {
	id, err := rserv.store.CreateReceipt(receipt)
	if err != nil {
		return "", err
	}

	return id, nil

}
