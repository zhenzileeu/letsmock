package letsmock

import (
	"testing"
	"fmt"
)

var passwordMock = PasswordMock{}

func TestPasswordMock_MockNumberPassword(t *testing.T) {
	fmt.Println("mock number password [6]: ", passwordMock.MockNumberPassword(6))
}

func TestPasswordMock_MockAlphaPassword(t *testing.T) {
	fmt.Println("mock alpha password [6]: ", passwordMock.MockAlphaPassword(6))
}

func TestPasswordMock_MockCapitalPassword(t *testing.T) {
	fmt.Println("mock capital alpha password [6]: ", passwordMock.MockCapitalPassword(6))
}

func TestPasswordMock_MockAlphaPasswordWithCapital(t *testing.T) {
	fmt.Println("mock alpha password with capital [6, 1]: ", passwordMock.MockAlphaPasswordWithCapital(6, 1))
	fmt.Println("mock alpha password with capital [6, 3]: ", passwordMock.MockAlphaPasswordWithCapital(6, 3))
}

func TestPasswordMock_MockAlphaNumberPassword(t *testing.T) {
	fmt.Println("mock password composed of number and alpha [6, 1]: ", passwordMock.MockAlphaNumberPassword(6, 1))
	fmt.Println("mock password composed of number and alpha [6, 3]: ", passwordMock.MockAlphaNumberPassword(6, 3))
}

func TestPasswordMock_MockNumberSpecialCharPassword(t *testing.T) {
	fmt.Println("mock password composed of number and special character [6, 1]: ", passwordMock.MockNumberSpecialCharPassword(6, 1, nil))
	fmt.Println("mock password composed of number and special character [6, 3]: ", passwordMock.MockNumberSpecialCharPassword(6, 3, nil))
	fmt.Println("mock password composed of number and special character [6, 2]: ", passwordMock.MockNumberSpecialCharPassword(6, 2, []rune{'$', '#', '@', '*'}))
}

func TestPasswordMock_MockAlphaSpecialCharPassword(t *testing.T) {
	fmt.Println("mock password composed of alpha and special character [6, 1]: ", passwordMock.MockAlphaSpecialCharPassword(6, 1, nil))
	fmt.Println("mock password composed of alpha and special character [6, 3]: ", passwordMock.MockAlphaSpecialCharPassword(6, 3, nil))
	fmt.Println("mock password composed of alpha and special character [6, 2]: ", passwordMock.MockAlphaSpecialCharPassword(6, 2, []rune{'$', '#', '@', '*'}))
}

func TestPasswordMock_MockNumberAlphaSpecialCharPassword(t *testing.T) {
	fmt.Println("mock password composed of alpha, number, capital and special character [8, 3, 1, 1]: ", passwordMock.MockNumberAlphaSpecialCharPassword(8, 1, 1, 3, nil))
	fmt.Println("mock password composed of alpha, number, capital and special character [8, 2, 2, 2]: ", passwordMock.MockNumberAlphaSpecialCharPassword(8, 2, 2, 2, nil))
}
