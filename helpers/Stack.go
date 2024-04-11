package helpers

type Stack []rune

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
