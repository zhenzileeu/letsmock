package letsmock

import (
	"testing"
	"fmt"
)

var textMock = TextMock{}

func TestTextMock_MockEnglishWord(t *testing.T) {
	fmt.Println("english word: ", textMock.MockEnglishWord())
}

func TestTextMock_MockEnglishSentence(t *testing.T) {
	fmt.Println("english sentence: ", textMock.MockEnglishSentence(10))
}

func TestTextMock_MockChineseWord(t *testing.T) {
	fmt.Println("chinese word: ", textMock.MockChineseWord())
}

func TestTextMock_MockChineseSentence(t *testing.T) {
	fmt.Println("chinese sentence: ", textMock.MockChineseSentence(10))
}

func TestTextMock_MockUniversityCn(t *testing.T) {
	fmt.Println("chinese university: ", textMock.MockUniversityCn())
}
