package letsmock

type NameMock struct {
	textMarkovChain markovChain
}

var defaultEnFamilyNameWordsFile, defaultEnFirstNameWordsFile string
var defaultCnFamilyNameWordsFile, defaultCnFirstNameWordsFile string

func init()  {
	defaultEnFamilyNameWordsFile = "./res/family_name_en"
	defaultEnFirstNameWordsFile  = "./res/first_name_en"
	defaultCnFamilyNameWordsFile = "./res/family_name_cn"
	defaultCnFirstNameWordsFile  = "./res/first_name_cn"
}


func (m *NameMock) MockEnName() string  {
	return m.MockEnFirstName() + " " + m.MockEnFamilyName()
}

func (m *NameMock) MockEnFirstName() string {
	firstNameEn, err := readWordsFromFile(defaultEnFirstNameWordsFile)
	if err != nil {
		return ""
	}

	m.textMarkovChain.Train(firstNameEn, 3, 0.01)
	name := m.textMarkovChain.RandomNGenerate(mockIntmn(3, 8))
	return capitalFirstCharacter(name)
}

func (m *NameMock) MockEnFamilyName() string  {
	familyNameEn, err := readWordsFromFile(defaultEnFamilyNameWordsFile)
	if err != nil {
		return ""
	}

	m.textMarkovChain.Train(familyNameEn, 2, 0.01)
	name := m.textMarkovChain.RandomNGenerate(mockIntmn(2, 8))
	return capitalFirstCharacter(name)
}

func (m *NameMock) MockCnName() string  {
	return m.MockCnFamilyName()+m.MockCnFirstName()
}

func (m *NameMock) MockCnFirstName() string  {
	firstNameCn,err := readWordsFromFile(defaultCnFirstNameWordsFile)
	if err != nil {
		return ""
	}

	// 语义训练
	m.textMarkovChain.Train(firstNameCn, 1, 0.01)

	return m.textMarkovChain.RandomNGenerate(mockIntmn(1, 3))
}

func (m *NameMock) MockCnFamilyName() string  {
	familyNameCn,err := readWordsFromFile(defaultCnFamilyNameWordsFile)
	if err != nil {
		return ""
	}

	return familyNameCn[mockIntn(len(familyNameCn))]
}