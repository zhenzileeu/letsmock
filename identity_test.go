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

func TestIdentityMock_MockBankNumberOfCMB(t *testing.T) {
	fmt.Println("mock cmb bank number: ", m.MockBankNumberOfCMB())
}

func TestIdentityMock_MockBankNumberOfPSBC(t *testing.T) {
	fmt.Println("mock psbc bank number: ", m.MockBankNumberOfPSBC())
}

func TestIdentityMock_MockBankNumberOfCCB(t *testing.T) {
	fmt.Println("mock ccb bank number: ", m.MockBankNumberOfCCB())
}

func TestIdentityMock_MockBankNumberOfBOC(t *testing.T) {
	fmt.Println("mock boc bank number: ", m.MockBankNumberOfBOC())
}

func TestIdentityMock_MockBankNumberOfICBC(t *testing.T) {
	fmt.Println("mock icbc bank number: ", m.MockBankNumberOfICBC())
}

func TestIdentityMock_MockBankNumber(t *testing.T) {
	fmt.Println("mock bank number: ", m.MockBankNumber())
}