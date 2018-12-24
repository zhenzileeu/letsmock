package letsmock

import (
	"testing"
	"fmt"
)

var numberMock = NumberMock{}

func TestNumberMock_MockPositive(t *testing.T) {
	fmt.Println("positive number: ", numberMock.MockPositive())
}

func TestNumberMock_MockNegative(t *testing.T) {
	fmt.Println("negative number: ", numberMock.MockNegative())
}

func TestNumberMock_MockNumber(t *testing.T) {
	fmt.Println("mock number: ", numberMock.MockNumber())
}

func TestNumberMock_MockIntn(t *testing.T) {
	fmt.Println("mock number 10: ", numberMock.MockIntn(10))
}

func TestNumberMock_MockIntmn(t *testing.T) {
	fmt.Println("mock number 10 ~ 100: ", numberMock.MockIntmn(10, 100))
	fmt.Println("mock number 100 ~ 10: ", numberMock.MockIntmn(100, 10))
	fmt.Println("mock number 100 ~ 100: ", numberMock.MockIntmn(100, 100))
}

func TestNumberMock_MockNorm(t *testing.T) {
	fmt.Println("mock norm distribution number: ", numberMock.MockNorm())
}

func TestNumberMock_MockFloat(t *testing.T) {
	fmt.Println("mock float number: ", numberMock.MockFloat())
}

func TestNumberMock_MockExp(t *testing.T) {
	fmt.Println("mock exp distribution number: ", numberMock.MockExp())
}

func TestNumberMock_MockFloatmn(t *testing.T) {
	fmt.Println("mock float number 100~1000: ", numberMock.MockFloatmn(100, 1000))
}









