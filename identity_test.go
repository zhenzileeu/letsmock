package letsmock

import (
	"testing"
	"fmt"
)

var m = IdentityMock{}

func TestIdentityMock_MockIdCardNum(t *testing.T) {
	fmt.Println("mock id card num: ", m.MockIdCardNum())
}

func TestIdentityMock_MockCnMobilePhoneNumber(t *testing.T) {
	fmt.Println("mock china mobile phone number: ", m.MockCnMobilePhoneNumber())
}

func TestIdentityMock_MockCnUnicomPhoneNumber(t *testing.T) {
	fmt.Println("mock china unicom phone number: ", m.MockCnUnicomPhoneNumber())
}

func TestIdentityMock_MockCnTelecomPhoneNumber(t *testing.T) {
	fmt.Println("mock china telecom phone number: ", m.MockCnTelecomPhoneNumber())
}

func TestIdentityMock_MockCnVirtualOperatorPhoneNumber(t *testing.T) {
	fmt.Println("mock china virtual operator phone number: ", m.MockCnVirtualOperatorPhoneNumber())
}

func TestIdentityMock_MockCnPhoneNumber(t *testing.T) {
	fmt.Println("mock china phone number: ", m.MockCnPhoneNumber())
}

