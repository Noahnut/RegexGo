package regexgo

type stackValue struct {
	val         byte
	numberOfDot int
	numberOfAlt int
}

// over two word have the concatenation operator
// And In the brackets mean the another pattern
// so need to record the dot number and alternative operation number push to stack
// for after end the  brackets

func (r *regexgo) infix2Post(pattern string) string {

	stack := make([]stackValue, 0)
	newPattern := make([]byte, 0)

	numberOfDot := 0
	numberOfAlt := 0

	for i := range pattern {
		switch pattern[i] {
		case '(':
			if numberOfDot > 1 {
				numberOfDot--
				newPattern = append(newPattern, '.')
			}
			s := stackValue{
				val:         pattern[i],
				numberOfDot: numberOfDot,
				numberOfAlt: numberOfAlt,
			}
			stack = append(stack, s)
			numberOfAlt = 0
			numberOfDot = 0
		case '+', '*', '?':
			newPattern = append(newPattern, pattern[i])
		case ')':
			if len(stack) == 0 {
				return ""
			}
			numberOfDot--
			for numberOfDot > 0 {
				numberOfDot--
				newPattern = append(newPattern, '.')
			}

			for numberOfAlt > 1 {
				numberOfAlt--
				newPattern = append(newPattern, '|')
			}
			numberOfAlt = stack[len(stack)-1].numberOfAlt
			numberOfDot = stack[len(stack)-1].numberOfDot
			stack = stack[0 : len(stack)-1]
			numberOfDot++
		case '|':
			numberOfDot--
			for numberOfDot > 0 {
				newPattern = append(newPattern, '.')
				numberOfDot--
			}

			numberOfAlt++
		default:
			if numberOfDot > 1 {
				newPattern = append(newPattern, '.')
				numberOfDot--
			}
			newPattern = append(newPattern, pattern[i])
			numberOfDot++
		}
	}

	numberOfDot--
	for numberOfDot > 0 {
		newPattern = append(newPattern, '.')
		numberOfDot--
	}

	for numberOfAlt > 0 {
		numberOfAlt--
		newPattern = append(newPattern, '|')
	}

	return string(newPattern)
}
