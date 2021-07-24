package regexgo

type stateType int

const (
	char  stateType = 0
	split stateType = 1
	match stateType = 2
	empty stateType = 3
)

type state struct {
	stateType     stateType
	char          byte
	next, nextTwo *state
}

type stackStruct struct {
	start *state
	out   *state
}

type regexgo struct {
	startState *state
	stateStack []*stackStruct
}

func NewRegexGo(pattern string) *regexgo {
	regexgo := regexgo{
		stateStack: make([]*stackStruct, 0),
	}

	regexgo.makeRegexNFA(pattern)

	return &regexgo
}

func (r *regexgo) CheckIsMatch(text string) bool {
	return r.matchTheNFAPattern(text)
}

func (r *regexgo) matchTheNFAPattern(text string) bool {
	stateIter := r.startState
	for i := 0; i < len(text); {
		switch stateIter.stateType {
		case split:
			nextChar := text[i]
			if stateIter.next.char == nextChar {
				stateIter = stateIter.next
			} else if stateIter.nextTwo.char == nextChar {
				stateIter = stateIter.nextTwo
			} else {
				return false
			}

		case char:
			if text[i] != stateIter.char {
				return false
			} else {
				stateIter = stateIter.next
			}
			i++
		case match:
			return false
		}
	}
	return true
}

func (r *regexgo) makeRegexNFA(pattern string) {
	newPattern := r.infix2Post(pattern)
	for i := range newPattern {
		newState := &state{}

		switch newPattern[i] {
		case '*':
			oldState := r.popFromStateStack()
			newState.stateType = split
			newState.next = oldState.start
			iter := oldState.start

			for iter.next.stateType != empty {
				iter = iter.next
			}

			iter.next = newState

			newState.nextTwo = r.makeEmptyState()
			r.pushToStateStack(&stackStruct{start: newState, out: newState.nextTwo})
		case '|':
			oneState := r.popFromStateStack()
			twoState := r.popFromStateStack()
			newState.stateType = split

			iter := oneState.start

			for iter.next.stateType != empty {
				iter = iter.next
			}

			r.appendTheState(&iter.next, twoState.out)
			newState.next = oneState.start
			newState.nextTwo = twoState.start
			r.pushToStateStack(&stackStruct{start: newState, out: twoState.out})
		case '?':
			oldState := r.popFromStateStack()
			newState.stateType = split
			newState.next = oldState.start
			r.appendTheState(&newState.nextTwo, oldState.start)
			r.pushToStateStack(&stackStruct{start: newState, out: newState.nextTwo})
		case '+':
			oldState := r.popFromStateStack()
			newState.stateType = split
			newState.next = oldState.start
			newState.nextTwo = r.makeEmptyState()
			r.patchTheState(oldState, newState)
			r.pushToStateStack(&stackStruct{start: oldState.start, out: newState})
		case '.':
			lastState := r.popFromStateStack()
			prevState := r.popFromStateStack()
			r.patchTheState(prevState, lastState.start)
			r.pushToStateStack(&stackStruct{start: prevState.start, out: lastState.start})
		default:
			newState.char = newPattern[i]
			newState.stateType = char
			newState.next = r.makeEmptyState()
			newState.nextTwo = r.makeEmptyState()
			r.pushToStateStack(&stackStruct{start: newState, out: newState.next})
		}
	}

	matchState := state{stateType: match}
	last := r.popFromStateStack()
	if (*last).out.stateType == empty {
		*last.out = matchState
	} else {
		*last.out.next = matchState
	}

	r.startState = last.start
}

func (r *regexgo) popFromStateStack() *stackStruct {
	topState := r.stateStack[len(r.stateStack)-1]
	r.stateStack = r.stateStack[:len(r.stateStack)-1]
	return topState
}

func (r *regexgo) pushToStateStack(s *stackStruct) {
	r.stateStack = append(r.stateStack, s)
}

func (r *regexgo) appendTheState(stateA **state, stateB *state) {

	iter := stateB
	for iter.stateType != empty {
		iter = iter.next
	}

	*stateA = iter
}

func (r *regexgo) makeEmptyState() *state {
	e := state{
		stateType: empty,
	}

	return &e
}

func (r *regexgo) patchTheState(prevState *stackStruct, lastState *state) {
	if prevState.out.stateType == empty {
		*prevState.out = *lastState
	} else {
		if prevState.out.next.stateType != empty {
			*prevState.out.nextTwo = *lastState
		} else {
			*prevState.out.next = *lastState
		}
	}
}
