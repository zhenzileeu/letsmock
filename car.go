package letsmock

import (
	"encoding/json"
	"strings"
)

type CarMock struct {
}

var cnCarPlateNoPrefixCfgFile = "./res/data/plate_no_prefix_cn.json"
var defaultCarBrandsFile = "./res/car_brands"

func (m *CarMock) FeedResFile(plateNoPrefixResFile, brandResFile string)  {
	cnCarPlateNoPrefixCfgFile = plateNoPrefixResFile
	defaultCarBrandsFile = brandResFile
}

func (m *CarMock) MockCarBrands() string  {
	brands,err := readWordsFromFile(defaultCarBrandsFile)
	if err != nil {
		return ""
	}

	markovChain := markovChain{}
	markovChain.Train(brands, 3, 0.01)

	return strings.ToUpper(markovChain.RandomNGenerate(mockIntmn(3, 10)))
}

func (m *CarMock) MockCnPlateNo() string {
	plateNoPrefix,err := m.getCnPlateNoPrefixFromFile(cnCarPlateNoPrefixCfgFile)
	if err != nil {
		plateNoPrefix = map[string][]string{
			"京": {"A", "B", "C", "D", "E", "F", "G"},
		}
	}

	// random select prefix car plate no
	plateNo := ""
	for key,value := range plateNoPrefix {
		plateNo = plateNo + key + value[mockIntn(len(value))]
		break
	}

	// random select number 3 - 7, user password mocker
	pwdMocker := PasswordMock{}
	number37 := ""
	for ; ;  {
		number37 = pwdMocker.MockAlphaNumberPassword(5, mockIntn(5))
		if strings.ContainsAny(number37, "io") {
			continue
		}
		break
	}
	plateNo = plateNo + strings.ToUpper(number37)

	return plateNo
}

func (m *CarMock) MockCnXinPlateNo() string  {
	plateNoPrefix,err := m.getCnPlateNoPrefixFromFile(cnCarPlateNoPrefixCfgFile)
	if err != nil {
		plateNoPrefix = map[string][]string{
			"京": {"A", "B", "C", "D", "E", "F", "G"},
		}
	}

	// random select prefix car plate no
	plateNo := ""
	for key,value := range plateNoPrefix {
		plateNo = plateNo + key + value[mockIntn(len(value))]
		break
	}

	// random select car power type, D for pure electrical and F for mix electrical
	if mockIntn(10) >= 5 {
		plateNo = plateNo + "F"
	} else {
		plateNo = plateNo + "D"
	}

	// random select number 4 - 8, user password mocker
	pwdMocker := PasswordMock{}
	number48 := ""
	for ; ;  {
		number48 = pwdMocker.MockAlphaNumberPassword(5, mockIntn(5))
		if strings.ContainsAny(number48, "io") {
			continue
		}
		break
	}
	plateNo = plateNo + strings.ToUpper(number48)

	return plateNo
}

func (m *CarMock) getCnPlateNoPrefixFromFile(filename string) (map[string][]string, error) {
	rawCfgContent,err := readFile(filename)
	if err != nil {
		return nil, err
	}

	prefixCfg := make(map[string]interface{})
	err = json.Unmarshal(rawCfgContent, &prefixCfg)
	if err != nil {
		return nil,err
	}

	plateNoPrefixCfg := make(map[string][]string)
	for k,v := range prefixCfg {
		switch v.(type) {
		case []interface{}:
			prefix := make([]string, len(v.([]interface{})))
			for i,value := range v.([]interface{}) {
				switch value.(type) {
				case string:
					prefix[i] = value.(string)
				}
			}
			plateNoPrefixCfg[k] = prefix
		}
	}

	return plateNoPrefixCfg,nil
}

