package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(content interface{}) {
	*s = append(*s, content)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	content := (*s)[len(*s) -1]
	*s = (*s)[:len(*s) -1]
	return content
}