package creditcard

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {

	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func (c CreditCard) OwnerName() string {
	return c.ownerName
}

func (c *CreditCard) setOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = value
	return nil
}

func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

func (c *CreditCard) SetCardNumber(value string) error {

	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid caredi card number format")
	}
	c.cardNumber = value
	return nil
}

func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) SetExpirationDate(month, year int) error {

	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}
	c.expirationMonth = month
	c.expirationYear = year
	return nil
}

func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("Invalid security code")
	}
	c.securityCode = value
	return nil
}

func (c CreditCard) AvailableCredit() float32 {
	return 5000.
}
