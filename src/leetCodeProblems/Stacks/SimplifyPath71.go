package main

/*
- Leetcode - https://leetcode.com/problems/simplify-path/description/
- Time - O(n)
- Space - O(n)
*/
type stackDir []string

func (s *stackDir) pop() string {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *stackDir) push(r string) {
	*s = append(*s, r)
}

func (s stackDir) isEmpty() bool {
	return len(s) == 0
}

func parseDir(dir string, s *stackDir) {
	switch dir {
	case ".": // do nothing
	case "..":
		if !s.isEmpty() {
			s.pop()
		}
	default:
		s.push(dir)
	}
}

func simplifyPath(path string) string {
	s := new(stackDir)
	start, i := 0, 1

	for i < len(path) {

		if string(path[i]) == "/" {

			if i-start > 1 {
				parseDir(path[start+1:i], s)
			}

			start = i
		}

		i++
	}

	if i-start > 1 {
		parseDir(path[start+1:i], s)
	}

	if s.isEmpty() {
		return "/"
	}

	rs := new(stackDir)

	for !s.isEmpty() {
		rs.push(s.pop())
	}

	out := ""

	for !rs.isEmpty() {
		out += "/" + rs.pop()
	}

	return out
}
