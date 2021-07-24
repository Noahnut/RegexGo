package regexgo

import (
	"testing"
)

func Test_Regex(t *testing.T) {
	reg := NewRegexGo("a*b")
	result := reg.CheckIsMatch("aaaab")

	if result == false {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("abc")

	if result {
		t.Error("regex is wrong")
	}

	reg = NewRegexGo("ab|b")
	result = reg.CheckIsMatch("ab")

	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("b")

	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("cd")

	if result {
		t.Error("regex is wrong")
	}

	reg = NewRegexGo("a?b")
	result = reg.CheckIsMatch("b")
	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("ab")
	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("cb")
	if result {
		t.Error("regex is wrong")
	}

	reg = NewRegexGo("a+b")

	result = reg.CheckIsMatch("aaaab")

	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("ab")

	if !result {
		t.Error("regex is wrong")
	}

	result = reg.CheckIsMatch("b")

	if result {
		t.Error("regex is wrong")
	}
}

func Test_RegexNFAbuild(t *testing.T) {
	compareStringOne := "abab"
	compareStringTwo := "abbb"

	regexGo := NewRegexGo(compareStringOne + "|" + compareStringTwo)
	iterator := regexGo.startState

	if iterator.stateType != split {
		t.Error("| first should be split type")
	}

	secondIterator := iterator.next
	firstIterator := iterator.nextTwo
	stringiter := 0
	for firstIterator.next != nil {
		if stringiter > len(compareStringOne) {
			t.Error("stringIter overflow")
		}

		if firstIterator.char != compareStringOne[stringiter] {
			t.Error("Character not match")
		}

		stringiter++
		firstIterator = firstIterator.next
	}

	stringiter = 0

	for secondIterator.next != nil {
		if stringiter > len(compareStringTwo) {
			t.Error("stringIter overflow")
		}

		if secondIterator.char != compareStringTwo[stringiter] {
			t.Error("Character not match")
		}

		stringiter++
		secondIterator = secondIterator.next
	}

}

func Test_RegexNFAbuildTwo(t *testing.T) {

	compareStringTwo := "cb"
	regexGo := NewRegexGo("a?b|" + compareStringTwo)

	iterator := regexGo.startState

	if iterator.stateType != split {
		t.Error("| first should be split type")
	}

	firstIterator := iterator.next
	twoIterator := iterator.nextTwo

	stringIterator := 0
	for firstIterator.stateType != match {
		if stringIterator > len(compareStringTwo) {
			t.Error("StringIter overflow")
		}

		if compareStringTwo[stringIterator] != firstIterator.char {
			t.Error("Character not match")
		}

		stringIterator++
		firstIterator = firstIterator.next
	}

	if twoIterator.stateType != split {
		t.Error("? first should be split type")
	}

	if twoIterator.nextTwo.char != 98 {
		t.Error("Character not match")
	}

	twoIterator = twoIterator.next

	if twoIterator.char != 97 {
		t.Error("Character not match")
	}

	if twoIterator.next.char != 98 {
		t.Error("Character not match")
	}

	twoIterator = twoIterator.next.next

	if twoIterator.stateType != match {
		t.Error("State not match")
	}
}
