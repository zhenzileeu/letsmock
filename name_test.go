package letsmock

import (
	"testing"
	"fmt"
)

var nameMock = NameMock{}

func TestNameMock_MockCnFamilyName(t *testing.T) {
	fmt.Println("mock CN family name: ", nameMock.MockCnFamilyName())
}

func TestNameMock_MockEnFamilyName(t *testing.T) {
	fmt.Println("mock EN family name: ", nameMock.MockEnFamilyName())
}

func TestNameMock_MockCnFirstName(t *testing.T) {
	fmt.Println("mock CN first name: ", nameMock.MockCnFirstName())
}

func TestNameMock_MockEnFirstName(t *testing.T) {
	fmt.Println("mock EN first name: ", nameMock.MockEnFirstName())
}

func TestNameMock_MockCnName(t *testing.T) {
	fmt.Println("mock CN name: ", nameMock.MockCnName())
}

func TestNameMock_MockEnName(t *testing.T) {
	fmt.Println("mock EN name: ", nameMock.MockEnName())
}






