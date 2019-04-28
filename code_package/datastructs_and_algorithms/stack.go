package datastructs_and_algorithms

type Stack struct {
	root *Node
	size int
}

type Node struct {
	next *Node
	val  interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Add(val interface{}) {
	oldRoot := s.root
	s.root = newNode(val)
	s.root.next = oldRoot
	s.size++
}

func (s *Stack) Pop() interface{} {
	oldRoot := s.root
	if oldRoot == nil {
		return nil
	}

	s.root = oldRoot.next
	s.size--
	return oldRoot.val
}

func (s *Stack) Read(n int) interface{} {

	if n < 0 || n >= s.size {
		return nil
	}

	var (
		r     = s.root
		i int = 0
	)
	for {
		if i == n {
			break
		}
		if r == nil {
			return nil
		}
		r = r.next
		i++
	}

	if r == nil {
		return nil
	}
	return r.val
}

func (s *Stack) SubN(n int) {
	if n < 0 || n >= s.size {
		return
	}
	var (
		r = s.root
	)
	for i := 0; i < n; i++ {
		if r == nil {
			break
		}
		r = r.next
	}
	s.root = r
	s.size -= n
}

func (s *Stack) GetSize() int {
	return s.size
}

func newNode(val interface{}) *Node {
	return &Node{
		val: val,
	}
}
