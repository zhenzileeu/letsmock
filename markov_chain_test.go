package letsmock

import (
	"testing"
	"fmt"
)

func TestMarkovChain_RandomNGenerate(t *testing.T) {
	words,err := readWordsFromFile("./res/first_name_en")
	if err != nil {
		t.Error("something goes wrong")
		return
	}

	c := markovChain{}
	c.Train(words, 2, 0.01)

	fmt.Println("random generate 4: ", c.RandomNGenerate(4))
	fmt.Println("random generate 4: ", c.RandomNGenerate(4))
	fmt.Println("random generate 4: ", c.RandomNGenerate(4))
}
