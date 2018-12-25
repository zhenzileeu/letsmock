package letsmock

import (
	"math/rand"
	"unicode/utf8"
	"strings"
)

type TextMock struct {
	wordChain markovChain
}

var defaultCommonWordsEnFile = "./res/common_words_en"
var defaultCommonWordsCnFile = "./res/common_words_cn"
var defaultUniversityCnFile = "./res/university_cn"

func (m *TextMock) MockEnglishWord() string  {
	return m.mockEnglish(mockIntmn(1, 10))
}

func (m *TextMock) MockEnglishSentence(n int) string  {
	sentence := ""
	for i := 0; i < n; i++ {
		sentence = sentence + " " + m.MockEnglishWord()
	}

	return sentence
}

func (m *TextMock) MockChineseWord() string  {
	return m.mockChinese(mockIntmn(1, 5))
}

func (m *TextMock) MockChineseSentence(n int) string  {
	sentence := ""
	for i:=0; i<n; i++ {
		sentence = sentence + m.MockChineseWord()
	}

	return sentence
}

func (m *TextMock) MockNumberStr(n int) string  {
	numberStr := ""
	for i:=0; i<n; i++ {
		numberStr += string('0'+mockIntn(10))
	}

	return numberStr
}

func (m *TextMock) MockUniversityCn() string  {
	universities,err := readWordsFromFile(defaultUniversityCnFile)
	if err != nil {
		return ""
	}

	universitiesChain := markovChain{}
	universitiesChain.Train(universities, 2, 0.01)

	fakeUniversity := ""
	for ; utf8.RuneCountInString(fakeUniversity) < 2;  {
		 fakeUniversity = universitiesChain.RandomNGenerate(mockIntmn(2, 10))
	}

	if strings.HasSuffix(fakeUniversity, "学院") {
		return fakeUniversity
	} else if strings.HasSuffix(fakeUniversity, "大学") {
		return fakeUniversity
	} else if mockIntn(10) < 5 {
		return fakeUniversity + "学院"
	}
	return fakeUniversity + "大学"
}

func (m *TextMock) mockChinese(n int) string  {
	word,err := readWordsFromFile(defaultCommonWordsCnFile)
	if err != nil {
		return m.randomMockChinese(n)
	}

	m.wordChain.Train(word, 2, 0.01)

	fakeWord := m.wordChain.RandomNGenerate(n)
	if fakeWord == "" {
		return m.randomMockChinese(n)
	}

	return fakeWord
}

func (m *TextMock) mockEnglish(n int) string  {
	word,err := readWordsFromFile(defaultCommonWordsEnFile)
	if err != nil {
		return m.randomMockEnglish(n)
	}

	m.wordChain.Train(word, 2, 0.01)

	fakeWord := m.wordChain.RandomNGenerate(n)
	if fakeWord == "" {
		return m.randomMockEnglish(n)
	}

	return fakeWord
}

func (m *TextMock) mockPunctuation() string  {
	return ","
}

func (m *TextMock) randomMockChinese(n int) string {
	hf,he := 0x4e00,0x9fa5

	word := ""
	for i := 0; i < n; i++ {
		word = word + string(hf + mockIntn(he-hf))
	}

	return word
}

func (m *TextMock) randomMockEnglish(n int) string  {
	word := make([]rune, 0, n)
	for i:=0; i<n; i++  {
		word = append(word, rune('a'+rand.Intn(26)))
	}
	return string(word)
}