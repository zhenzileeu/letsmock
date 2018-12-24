package letsmock

import "math/rand"

type TextMock struct {
	wordChain markovChain
}

var defaultCommonWordsEnFile = "./res/common_words_en"

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

func (m *TextMock) mockChinese(n int) string  {
	hf,he := 0x4e00,0x9fa5

	word := ""
	for i := 0; i < n; i++ {
		word = word + string(hf + mockIntn(he-hf))
	}

	return word
}

func (m *TextMock) mockEnglish(n int) string  {
	word,err := readDataFromFile(defaultCommonWordsEnFile)
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

func (m *TextMock) randomMockEnglish(n int) string  {
	word := make([]rune, 0, n)
	for i:=0; i<n; i++  {
		word = append(word, rune('a'+rand.Intn(26)))
	}
	return string(word)
}