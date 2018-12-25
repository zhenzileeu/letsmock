package letsmock

import (
	"strconv"
	"fmt"
	"github.com/engoengine/math"
	"strings"
)

type InternetMock struct {
}

func (m *InternetMock) MockIpv4Addr() string  {
	addr1 := strconv.Itoa(mockIntmn(1, 256))
	addr2 := strconv.Itoa(mockIntn(256))
	addr3 := strconv.Itoa(mockIntn(256))
	addr4 := strconv.Itoa(mockIntn(256))

	return addr1 + "." + addr2 + "." + addr3 + "." + addr4
}

func (m *InternetMock) MockIpv4IntAddr() uint32  {
	addr1 := uint32(mockIntmn(1, 256)) << 24
	addr2 := uint32(mockIntn(256)) << 16
	addr3 := uint32(mockIntn(256)) << 8
	addr4 := uint32(mockIntn(256))

	return addr1+addr2+addr3+addr4
}

func (m *InternetMock) MockIpv4HexAddr() string  {
	addr := m.MockIpv4IntAddr()

	return fmt.Sprintf("%X", addr)
}

func (m *InternetMock) MockIpv6Addr() string  {
	addr1 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr2 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr3 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr4 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr5 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr6 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr7 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))
	addr8 := fmt.Sprintf("%04x", uint16(mockIntn(math.MaxUint16+1)))

	return addr1+":"+addr2+":"+addr3+":"+addr4+":"+addr5+":"+addr6+":"+addr7+":"+addr8
}

func (m *InternetMock) MockSiteDomainName(n int) string  {
	siteName := m.MockSiteName(n)
	domain := m.mockTopDomain()

	return "www."+siteName+"."+domain
}

func (m *InternetMock) MockSiteName(n int) string {
	sites,err := readWordsFromFile("./res/website_name_list")
	if err != nil {
		txtMock := TextMock{}
		return txtMock.mockEnglish(n)
	}

	siteMarkovChain := markovChain{}
	siteMarkovChain.Train(sites, 1, 0.01)

	siteName := ""
	for ; ;  {
		siteName = siteMarkovChain.RandomNGenerate(n)
		if strings.HasSuffix(siteName, ".") {
			continue
		} else if strings.HasPrefix(siteName, ".") {
			continue
		} else if strings.HasPrefix(siteName, "-") {
			continue
		} else if strings.HasSuffix(siteName, "-") {
			continue
		}
		break
	}

	return siteName
}

func (m *InternetMock) MockEmailAddress(n int) string  {
	// use pwd mocker
	pwdMocker := PasswordMock{}
	emailAddr := ""
	for ; ;  {
		if n <= 3 {
			emailAddr = pwdMocker.MockAlphaPassword(n)
		} else if n <= 6 {
			emailAddr = pwdMocker.MockAlphaNumberPassword(n, mockIntn(n))
		} else {
			nSpecial := mockIntn(1)
			nCapital := mockIntn(2)
			nAlpha   := mockIntn(n-3)
			emailAddr = pwdMocker.MockNumberAlphaSpecialCharPassword(n, nSpecial, nCapital, nAlpha, []rune{'-', '_'})
		}

		if strings.HasPrefix(emailAddr, "-") {
			continue
		} else if strings.HasPrefix(emailAddr, "_") {
			continue
		} else if strings.HasSuffix(emailAddr, "-") {
			continue
		} else if strings.HasSuffix(emailAddr, "_") {
			continue
		}
		break
	}

	//
	siteName := m.MockSiteName(mockIntmn(2,10))
	domain := m.mockTopDomain()

	return emailAddr+"@"+siteName+"."+domain
}

func (m *InternetMock) mockTopDomain() string {
	domains,err := readWordsFromFile("./res/domain_top_list")
	if err != nil {
		domains = []string{"com", "org", "edu", "xyz"}
	}

	return domains[mockIntn(len(domains))]
}