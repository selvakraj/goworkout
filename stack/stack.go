package stack

type Data struct {
	Name string
}

type Stack struct {
	data []*Data
}

func Make(cap int) *Stack {
	return &Stack{
		data: make([]*Data, 0, cap),
	}
}

func (s *Stack) Count() int {
	return len(s.data)
}

func (s *Stack) Push(data *Data) {
	s.data = append(s.data, data)
}
