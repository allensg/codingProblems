package helpers

type Stack []rune
type StringStack []string
type IntStack []int

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

// string
func (s *StringStack) Push(val string) {
	*s = append(*s, val)
}

func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s StringStack) IsEmpty() bool {
	return len(s) == 0
}

// int
func (s *IntStack) Push(val int) {
	*s = append(*s, val)
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

func (s IntStack) IsEmpty() bool {
	return len(s) == 0
}
