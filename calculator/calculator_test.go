// +build unit_test

package calculator

import (
	"fmt"
	"testing"
)

type MockDiscountAgent struct {
	Params []*domain.DiscountParams
}

func (z *MockDiscountAgent) GetDiscountParams(productID int) (params []*domain.DiscountParams) {
	return z.Params
}

func Test_New(t *testing.T) {
	mda := &MockDiscountAgent{}

	fmt.Println("yyyayayyayayayyayayya")

	actual := New(mda)
	if actual == nil {
		t.Errorf("crap")
	}
}
