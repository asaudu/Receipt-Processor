package rewards

import (
	"fmt"
	"math"
	"strings"
	"time"

	"addyCodes.com/ReceiptProcessor/models"
)

func RewardPoints(receipt models.Receipt) int {
	points := 0

	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.Description))
		if descLength%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	if isBetweenTwoAndFour(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func isRoundDollarAmount(total string) bool {
	var totalAmount float64
	fmt.Sscanf(total, "%f", &totalAmount)
	return math.Mod(totalAmount, 1.0) == 0
}

func isMultipleOfQuarter(total string) bool {
	var totalAmount float64
	fmt.Sscanf(total, "%f", &totalAmount)
	return math.Mod(totalAmount, 0.25) == 0
}

func isOddDay(date string) bool {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	return t.Day()%2 != 0
}

func isBetweenTwoAndFour(timeOfPurchase time.Time) bool {
	hour := timeOfPurchase.Hour()
	return hour >= 14 && hour < 16
}
