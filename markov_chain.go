package letsmock

import (
	"unicode/utf8"
	"strings"
)

type markovChain struct {
	alphabetProbChain map[string][]float32
	probChainOrder  int
	alphabets	[]rune
}

func (chain *markovChain) Train(data []string, order int, prior float32) {
	// construct the alphabets
	chain.buildAlphabets(data)
	chain.probChainOrder = order

	// 使用占位符
	holdplaceStr := chain.holdPlaceStr(chain.probChainOrder)

	markovObservation := make(map[string][]string)
	for i := 0; i < len(data); i++ {
		word := make([]rune, 0)
		word = append(word, holdplaceStr...)
		word = append(word, []rune(data[i])...)
		word = append(word, holdplaceStr...)
		for k := 0; k < len(word)-order; k++ {
			key := string(word[k:k+order])

			if _, ok := markovObservation[key]; !ok {
				markovObservation[key] = []string{}
			}
			markovObservation[key] = append(markovObservation[key], string(word[k+order]))
		}
	}

	// construct markov chain
	chain.alphabetProbChain = make(map[string][]float32)
	for key,value := range markovObservation {
		if _, ok := chain.alphabetProbChain[key]; !ok {
			chain.alphabetProbChain[key] = make([]float32, 0, len(chain.alphabets))
		}

		for _,alpha := range chain.alphabets {
			chain.alphabetProbChain[key] = append(chain.alphabetProbChain[key], prior+float32(chain.countAlphabet(string(alpha), value)))
		}
	}
	return
}

func (chain *markovChain) buildAlphabets(data []string) {
	alphabetsMap := make(map[rune]int, 0)
	for i := 0; i < len(data); i++ {
		word := []rune(data[i])
		for k := 0; k < len(word); k++ {
			if _, ok := alphabetsMap[word[k]]; !ok {
				alphabetsMap[word[k]] = 0
			}
			alphabetsMap[word[k]] += 1
		}
	}

	chain.alphabets = make([]rune, 0, len(alphabetsMap))
	for alphabet := range alphabetsMap {
		chain.alphabets = append(chain.alphabets, alphabet)
	}

}

func (chain *markovChain) countAlphabet(alpha string, value []string) int  {
	cnt := 0
	for _,v := range value {
		if v == alpha {
			cnt ++
		}
	}

	return cnt
}

func (chain *markovChain) generate(alphabets []rune, markovProb []float32) string {
	if len(markovProb) != len(alphabets) {
		return ""
	}

	cumsumProb := make([]float32, len(markovProb))
	cumsum := float32(0.0)
	for k,prob := range markovProb {
		cumsum += prob
		cumsumProb[k] = cumsum
	}

	randNum := mockFloat32() * cumsum
	for k, prob := range cumsumProb {
		if randNum < prob {
			return string(alphabets[k])
		}
	}

	return ""
}

func (chain *markovChain) RandomNGenerate(n int) string  {

	name := string(chain.holdPlaceStr(chain.probChainOrder))
	nameLen := utf8.RuneCountInString(name)
	for ; nameLen < n+chain.probChainOrder;  {
		key := string([]rune(name)[nameLen-chain.probChainOrder:nameLen])

		if _,ok := chain.alphabetProbChain[key]; !ok {
			break
		}

		nextLetter := chain.generate(chain.alphabets, chain.alphabetProbChain[key])
		if nextLetter=="" {
			break
		}

		name = name + nextLetter
		nameLen = nameLen+1
	}

	return chain.removeHoldPlaceRune(name)
}

func (chain *markovChain) holdPlaceStr(n int) []rune {
	holdPlaceStr := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		holdPlaceStr = append(holdPlaceStr, chain.holdPlaceRune())
	}
	return holdPlaceStr
}

func (chain *markovChain) holdPlaceRune() rune {
	return '#'
}

func (chain *markovChain) removeHoldPlaceRune(str string) string  {
	return strings.Trim(str, string(chain.holdPlaceRune()))
}