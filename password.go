package letsmock

import (
	"strings"
)

type PasswordMock struct {
}

var defaultSpecialChar = []rune{',', '.', '/', '#', '$', '%', '^', '~', '@', '!', '&', '*', '(', ')', '{', '}', '[', ']', '-', '+', '=', ':', ';', '?', '<', '>', '|'}

func (m *PasswordMock) MockNumberPassword(n int) string {
	textMock := TextMock{}

	return textMock.MockNumberStr(n)
}

func (m *PasswordMock) MockAlphaPassword(n int) string  {
	textMock := TextMock{}

	return textMock.randomMockEnglish(n)
}

func (m *PasswordMock) MockCapitalPassword(n int) string  {
	return strings.ToUpper(m.MockAlphaPassword(n))
}

func (m *PasswordMock) MockAlphaPasswordWithCapital(n, nCapital int) string  {
	if nCapital > n {
		nCapital = n
	}

	if nCapital == 0 {
		return m.MockAlphaPassword(n)
	} else if nCapital == n {
		return m.MockCapitalPassword(n)
	} else {
		alphaPassword := m.MockAlphaPassword(n-nCapital)
		capitalPassword := m.MockCapitalPassword(nCapital)

		return shuffleStr(alphaPassword+capitalPassword)
	}
}

func (m *PasswordMock) MockAlphaNumberPassword(n, nAlpha int) string {
	if nAlpha >= n {
		nAlpha = n - 1
	}

	alphaPassword := m.MockAlphaPassword(nAlpha)
	numberPassword := m.MockNumberPassword(n - nAlpha)

	password := alphaPassword + numberPassword

	return shuffleStr(password)
}

func (m *PasswordMock) MockNumberSpecialCharPassword(n, nSpecial int, specialChar []rune) string  {
	if specialChar == nil {
		specialChar = defaultSpecialChar
	}

	if nSpecial >= n {
		nSpecial = n - 1
	}

	numberPwd := m.MockNumberPassword(n-nSpecial)
	spechaPwd := m.mockSpecialCharPassword(nSpecial, specialChar)

	return shuffleStr(numberPwd+spechaPwd)
}

func (m *PasswordMock) MockAlphaSpecialCharPassword(n, nSpecial int, specialChar []rune) string  {
	if specialChar == nil {
		specialChar = defaultSpecialChar
	}

	if nSpecial >= n {
		nSpecial = n - 1
	} else if nSpecial == 0 {
		return m.MockAlphaPassword(n)
	}

	alphaPwd := m.MockAlphaPassword(n - nSpecial)
	spechaPwd := m.mockSpecialCharPassword(nSpecial, specialChar)

	return shuffleStr(alphaPwd+spechaPwd)
}

func (m *PasswordMock) MockNumberAlphaSpecialCharPassword(n, nSpecial, nCapital, nAlpha int, specialChar []rune) string  {
	if specialChar == nil {
		specialChar = defaultSpecialChar
	}

	specialPwd := m.mockSpecialCharPassword(nSpecial, specialChar)
	capitalPwd := m.MockCapitalPassword(nCapital)
	alphaPwd   := m.MockAlphaPassword(nAlpha)
	numberPwd  := ""
	if nSpecial+nCapital+nAlpha < n {
		numberPwd = m.MockNumberPassword(n-nSpecial-nCapital-nAlpha)
	}

	password := []rune(shuffleStr(specialPwd+capitalPwd+alphaPwd+numberPwd))
	return string(password[0:n-1])
}

func (m *PasswordMock) mockSpecialCharPassword(n int, specialChar []rune) string  {

	password := ""
	for i := 0; i < n; i++ {
		password += string(specialChar[mockIntn(len(specialChar))])
	}

	return password
}