package letsmock

import (
	"time"
	"strings"
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
)

type IdentityMock struct {
}

var cnAdministrativeDivisionCodeFile = "./res/administrative_division_code_cn"
var cnDefaultBankBINFile = "./res/data/bank_identify_number_cn.json"
var defaultAdministrativeDivisionCode = []string{"110101","东城区","110102","西城区","110105","朝阳区","110106","丰台区","110107","石景山区","110108","海淀区","110109","门头沟区","110111","房山区","110112","通州区","110113","顺义区","110114","昌平区","110115","大兴区","110116","怀柔区","110117","平谷区"}

func (m *IdentityMock) MockIdCardNum() string  {
	//
	content,err := readWordsFromFile(cnAdministrativeDivisionCodeFile)
	if err != nil {
		content = defaultAdministrativeDivisionCode
	}

	var cnAdministrativeDivisionCode = make(map[string]string)
	for i := 0; i < len(content); i+=2 {
		cnAdministrativeDivisionCode[content[i]] = content[i+1]
	}

	// random select administrative division code
	divisionCode := "110108"
	for key := range cnAdministrativeDivisionCode {
		if strings.HasSuffix(key, "0") {
			continue
		}
		divisionCode = key
		break
	}

	// random birthday
	dateMock := DateMock{}
	birthday := dateMock.MockHistoryDate()
	birthdayTime,err := time.Parse("2006-01-02", birthday)
	if err != nil {
		birthday = "19800808"
	} else {
		birthday = birthdayTime.Format("20060102")
	}

	// random sequence number
	textMock := TextMock{}
	numberSeq := textMock.MockNumberStr(3)

	// calculate the last one number
	randomIDNum := divisionCode + birthday + numberSeq
	numberCoef := []int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
	sum := 0
	for i,v := range numberCoef {
		sum += int(randomIDNum[i]-'0')*v
	}
	checkCode := []rune{'1','0','X','9','8','7','6','5','4','3','2'}
	randomIDNum += string(checkCode[sum%11])

	return randomIDNum
}

func (m *IdentityMock) MockCnPhoneNumber() string  {
	idx := mockIntn(4)

	switch idx {
	case 0:
		return m.MockCnMobilePhoneNumber()
	case 1:
		return m.MockCnUnicomPhoneNumber()
	case 2:
		return m.MockCnTelecomPhoneNumber()
	default:
		return m.MockCnVirtualOperatorPhoneNumber()
	}
}

func (m *IdentityMock) MockCnMobilePhoneNumber() string  {
	segment := []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188", "198"}

	randSeg := segment[mockIntn(len(segment))]
	textMock := TextMock{}

	return randSeg+textMock.MockNumberStr(8)
}

func (m *IdentityMock) MockCnUnicomPhoneNumber() string  {
	segment := []string{"130", "131", "132", "145", "155", "156", "175", "176", "185", "186", "166"}

	randSeg := segment[mockIntn(len(segment))]
	textMock := TextMock{}

	return randSeg+textMock.MockNumberStr(8)
}

func (m *IdentityMock) MockCnTelecomPhoneNumber() string  {
	segment := []string{"133", "153", "173", "177", "180", "181", "189", "199"}

	randSeg := segment[mockIntn(len(segment))]
	textMock := TextMock{}

	return randSeg+textMock.MockNumberStr(8)
}

func (m *IdentityMock) MockCnVirtualOperatorPhoneNumber() string  {
	segment := []string{"170"}

	randSeg := segment[mockIntn(len(segment))]
	textMock := TextMock{}

	return randSeg+textMock.MockNumberStr(8)
}

func (m *IdentityMock) MockBankNumberOfCMB() string  {
	bankBinNumber := make(map[string]int)
	bankBinNumber["621483"] = 16
	bankBinNumber["690755"] = 18
	bankBinNumber["622575"] = 16
	return m.mockBankNumberOfBank(cnDefaultBankBINFile, "cmb", bankBinNumber)
}

func (m *IdentityMock) MockBankNumberOfPSBC() string  {
	bankBinNumber := make(map[string]int)
	bankBinNumber["621096"] = 19
	bankBinNumber["621095"] = 19
	bankBinNumber["622810"] = 16
	return m.mockBankNumberOfBank(cnDefaultBankBINFile, "psbc", bankBinNumber)
}

func (m *IdentityMock) MockBankNumberOfBOC() string  {
	bankBinNumber := make(map[string]int)
	bankBinNumber["621661"] = 19
	bankBinNumber["377677"] = 15
	bankBinNumber["622750"] = 16
	return m.mockBankNumberOfBank(cnDefaultBankBINFile, "boc", bankBinNumber)
}

func (m *IdentityMock) MockBankNumberOfCCB() string  {
	bankBinNumber := make(map[string]int)
	bankBinNumber["621081"] = 19
	bankBinNumber["5453242"] = 18
	bankBinNumber["531693"] = 16
	return m.mockBankNumberOfBank(cnDefaultBankBINFile, "ccb", bankBinNumber)
}

func (m *IdentityMock) MockBankNumberOfICBC() string  {
	bankBinNumber := make(map[string]int)
	bankBinNumber["620403"] = 18
	bankBinNumber["620200"] = 18
	bankBinNumber["427029"] = 16
	return m.mockBankNumberOfBank(cnDefaultBankBINFile, "icbc", bankBinNumber)
}

func (m *IdentityMock) MockBankNumber() string  {
	banks := []string{"cmb", "icbc", "boc", "psbc", "ccb"}

	defaultBinNumber := map[string]int{"620403":18,"621081":19,"622750":16,"621095":19,"690755":18}

	return m.mockBankNumberOfBank(cnDefaultBankBINFile, banks[mockIntn(len(banks))], defaultBinNumber)
}

func (m *IdentityMock) mockBankNumberOfBank(filename, bankName string, defaultBinNumber map[string]int) string  {
	bankBinNumber, err := m.getBankBINFromFile(cnDefaultBankBINFile, bankName)
	if err != nil || bankBinNumber==nil {
		bankBinNumber = defaultBinNumber
	}

	// random select one bin number
	binNumber := ""
	bankNumberLen := 0
	for key,value := range bankBinNumber {
		binNumber = key
		bankNumberLen = value
		break
	}

	return m.generateBankNumberUseLuhn(binNumber, bankNumberLen)
}

func (m *IdentityMock) getBankBINFromFile(filename, bankName string) (map[string]int, error)  {
	content,err := readFile(filename)
	if err != nil {
		return nil, err
	}

	configure := make(map[string]interface{})
	err = json.Unmarshal(content, &configure)
	if err != nil {
		fmt.Println("error load config")
		return nil, err
	}

	if _,ok := configure[bankName]; !ok {
		return nil, errors.New("not configure")
	}

	switch configure[bankName].(type) {
	case map[string]interface{}:
		binNumber := make(map[string]int)
		for key,value := range configure[bankName].(map[string]interface{}) {
			switch value.(type) {
			case int:
				binNumber[key] = value.(int)
			case float64:
				binNumber[key] = int(value.(float64))
			}
		}
		return binNumber,nil
	default:
		return nil,errors.New("type error")
	}
}

func (m *IdentityMock) generateBankNumberUseLuhn(binNum string, n int) string {
	userDefinedNumberLen := n - len(binNum) - 1

	if userDefinedNumberLen < 0 {
		return binNum
	}

	txtMocker := TextMock{}
	userDefinedNumber := txtMocker.MockNumberStr(userDefinedNumberLen)

	// Luhn algorithm
	bankNumber := binNum + userDefinedNumber
	sumNumber := 0
	for i,j := n - 2,0; i >= 0; i-- {
		if j == 0 {
			sumNumber = sumNumber + int(bankNumber[i] - '0')*2
			j = 1
		} else {
			sumNumber = sumNumber + int(bankNumber[i] - '0')
			j = 0
		}
	}

	// mod 10
	rem := sumNumber%10
	if rem == 0 {
		rem = 10
	}

	return bankNumber+string(10 - rem + '0')
}