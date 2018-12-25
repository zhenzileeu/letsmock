package letsmock

import (
	"testing"
	"fmt"
)

var internetMock = InternetMock{}

func TestInternetMock_MockIpv4Addr(t *testing.T) {
	fmt.Println("mock ipv4 address: ", internetMock.MockIpv4Addr())
}

func TestInternetMock_MockIpv4IntAddr(t *testing.T) {
	fmt.Println("mock ipv4 integer address: ", internetMock.MockIpv4IntAddr())
}

func TestInternetMock_MockIpv4HexAddr(t *testing.T) {
	fmt.Println("mock ipv4 hex address: ", internetMock.MockIpv4HexAddr())
}

func TestInternetMock_MockIpv6Addr(t *testing.T) {
	fmt.Println("mock ipv6 address: ", internetMock.MockIpv6Addr())
}

func TestInternetMock_MockSiteName(t *testing.T) {
	fmt.Println("mock site name [4]: ", internetMock.MockSiteName(4))
	fmt.Println("mock site name [random]: ", internetMock.MockSiteName(mockIntmn(2, 10)))
}

func TestInternetMock_MockSiteDomainName(t *testing.T) {
	fmt.Println("mock site domain name [4]: ", internetMock.MockSiteDomainName(4))
	fmt.Println("mock site domain name [random]: ", internetMock.MockSiteDomainName(mockIntmn(2, 10)))
}

func TestInternetMock_MockEmailAddress(t *testing.T) {
	fmt.Println("mock email address [3]: ", internetMock.MockEmailAddress(3))
	fmt.Println("mock email address [5]: ", internetMock.MockEmailAddress(5))
	fmt.Println("mock email address [10]: ", internetMock.MockEmailAddress(10))
	fmt.Println("mock email address [random]: ", internetMock.MockEmailAddress(mockIntmn(1, 9)))
}
