package letsmock

import (
	"math/rand"
	"time"
	"unicode"
	"strings"
	"io/ioutil"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

// rand positive number
func mockInt() int {
	return rand.Int()
}

// float number in [0.0, 1.0)
func mockFloat32() float32  {
	return rand.Float32()
}

// standard normal distribution number in [-maxFloat64, +maxFloat64]
func mockNormStandard() float64  {
	return mockNorm(0.0, 1.0)
}

// normal distribution number in [-maxFloat64, +maxFloat64]
func mockNorm(mu, sigma float64) float64  {
	return sigma*rand.NormFloat64()+mu
}

// normal distribution number in (0.0, +maxFloat64]
func mockExp() float64  {
	return rand.ExpFloat64()
}

// random number between [0,n]
func mockIntn(n int) int  {
	return rand.Intn(n)
}

// rand number between [m, n]
func mockIntmn(m, n int) int  {
	if m == n {
		return m
	}

	if m > n {
		return mockIntn(m - n) + n
	}

	return mockIntn(n-m) + m
}

// random negative number
func mockNegative() int  {
	return -mockInt()
}

//
func capitalFirstCharacter(str string) string {
	runeStr := []rune(str)
	if len(runeStr) == 0 {
		return ""
	}

	runeStr[0] = unicode.ToUpper(runeStr[0])

	return string(runeStr)
}

func readDataFromFile(filename string) ([]string, error)  {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}

	words := strings.Split(string(content), ",")

	return words,nil
}

func shuffleStr(str string) string {
	runeStr := []rune(str)
	idxMap := make(map[int]int, len(runeStr))
	for idx := range runeStr {
		idxMap[idx] = idx
	}

	var shuffleStr = ""
	for idx := range idxMap {
		shuffleStr += string(runeStr[idx])
	}

	return shuffleStr
}
