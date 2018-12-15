package calculator

import "testing"

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
	if Math.Abs(actual-expected) < 0.00000001 {
		t.Errorf("GetDiscount: Expected %f and got %f", expected, actual)
	}

}
