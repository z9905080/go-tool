package gotool

import (
	"log"
	"testing"
)

func TestPermutations(t *testing.T) {
	allCombinationList := Permutations(1,2,3)
	log.Println(allCombinationList)
}