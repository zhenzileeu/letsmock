package letsmock

import (
	"testing"
	"fmt"
)

var carMock = CarMock{}

func TestCarMock_MockCnPlateNo(t *testing.T) {
	fmt.Println("mock car plate no: ", carMock.MockCnPlateNo())
}

func TestCarMock_MockCnXinPlateNo(t *testing.T) {
	fmt.Println("mock xin car plate no: ", carMock.MockCnXinPlateNo())
}

func TestCarMock_MockCarBrands(t *testing.T) {
	fmt.Println("mock car brands: ", carMock.MockCarBrands())
}
