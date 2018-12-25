package letsmock

import (
	"time"
	"strings"
)

type IdentityMock struct {
}

var cnAdministrativeDivisionCodeFile = "./res/administrative_division_code_cn"
var defaultAdministrativeDivisionCode = []string{"110101","东城区","110102","西城区","110105","朝阳区","110106","丰台区","110107","石景山区","110108","海淀区","110109","门头沟区","110111","房山区","110112","通州区","110113","顺义区","110114","昌平区","110115","大兴区","110116","怀柔区","110117","平谷区"}

func (m *IdentityMock) MockIdCardNum() string  {
	//
	content,err := readDataFromFile(cnAdministrativeDivisionCodeFile)
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