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
