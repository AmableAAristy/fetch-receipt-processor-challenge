package services

import (
	"Fetch/models"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func PointCalc(receipt models.Receipt) int {
	point := 0

	//if the receipt is mostly returns then the receipts subtract points.
	var negativeAmount bool

	amount := strings.Split(receipt.Total, ".")

	cents := 0
	if len(amount) > 1 {
		cents, _ = strconv.Atoi(amount[1])
	}

	if amount[0][0:1] == "-" {
		negativeAmount = true
	}

	//50 points if the total is a round dollar amount with no cents.
	if cents == 0 && !negativeAmount {
		point += 50
	}
	//25 points if the total is a multiple of 0.25.
	//I am making the assumption that it is asking the total amount including the cents,
	//the above condition should always apply hence why it is not an else if
	if cents%25 == 0 && !negativeAmount {
		point += 25
	}

	//If receipt is mostly returns subtract points
	if cents == 0 && negativeAmount {
		point -= 50
	}
	if cents%25 == 0 && negativeAmount {
		point -= 25
	}

	//One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			point += 1
		}
	}

	//5 points for every two items on the receipt.
	numItems := len(receipt.Items)
	point += (numItems / 2) * 5

	//If the trimmed length of the item description is a multiple of 3,
	//multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)

		if len(trimmedDescription)%3 == 0 {

			value, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				print(err.Error() + " Price could not be parsed.")
			}

			valueWithCeiling := int(math.Ceil(value * 0.2))

			point += valueWithCeiling
		}

	}

	//6 points if the day in the purchase date is odd.
	lastDigitDate := receipt.PurchaseDate[9:10]
	lastDigitDateInt, err := strconv.Atoi(lastDigitDate)
	if err != nil {
		print(err.Error() + " Date in incorrect format")
	}
	if lastDigitDateInt%2 == 1 {
		point += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm.
	hourOfPurchase := receipt.PurchaseTime[0:2]
	if hourOfPurchase == "14" || hourOfPurchase == "15" {
		point += 10
	}
	return point
}
