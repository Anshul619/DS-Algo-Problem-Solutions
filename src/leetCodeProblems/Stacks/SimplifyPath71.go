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

func (s *stackDir) push(t string) {
	*s = append(*s, t)
}

func (s *stackDir) isEmpty() bool {
	return len(*s) == 0
}

func parseDir(dir string, st *stackDir) {
	switch dir {
	case ".", "":
		// do nothing
	case "..":
		if !st.isEmpty() {
			st.pop()
		}
	default:
		st.push(dir)
	}
}
func simplifyPath(path string) string {
	st := new(stackDir)

	start := -1

	for i, v := range path {
		if string(v) == "/" {
			if start != -1 {
				parseDir(path[start+1:i], st)
			}
			start = i
		}
	}

	if start != -1 && start < len(path)-1 {
		parseDir(path[start+1:], st)
	}

	if st.isEmpty() {
		return "/"
	}

	out := ""
	for !st.isEmpty() {
		out = "/" + st.pop() + out
	}

	return out
}
