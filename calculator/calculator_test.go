package calculator

import (
	"fmt"
	"testing"
)

type MockProduct struct {
	unitPrice float64
}

func (z *MockProduct) GetUnitPrice() (price float64) {
	return z.unitPrice
}

func Test_GetDiscount(t *testing.T) {
	mp := &MockProduct{
		unitPrice: 12.34,
	}

	quantity := 10
	expected := 18.51

	actual := GetDiscount(mp, quantity)
	fmt.Printf("GetDiscount: Expected %f and got %f", expected, actual)

}