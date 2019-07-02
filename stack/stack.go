package stack

import "strings"

type stack struct {
	data []string
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() string {
	return s.data[len(s.data)-1]
}

func (s *stack) pop() string {
	var x string
	x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return x
}

func (s *stack) push(a string) {
	s.data = append(s.data, a)
}

// SyntaxCheckSimple performs a simple check based on open and closed braces
// 0 = fine
// 1 = unclosed
// 2 = brace mismatch
func SyntaxCheckSimple(code string) int {
	s := stack{}

	for _, c := range strings.Split(code, "") {
		if c == "{" || c == "[" || c == "(" {
			s.push(c)
			continue
		}

		if c == "}" || c == "]" || c == ")" {
			if s.empty() {
				return 2
			}

			var expected string

			switch s.pop() {
			case "{":
				expected = "}"
			case "(":
				expected = ")"
			case "[":
				expected = "]"
			}

			if c != expected {
				return 2
			}
		}
	}

	if !s.empty() {
		return 1
	}

	return 0
}
