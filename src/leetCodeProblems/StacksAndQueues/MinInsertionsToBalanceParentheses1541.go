package main

func minInsertions(s string) int {
	st := new(stackParenthesis)
	out := 0

	for _, r := range s {
		if st.isEmpty() {
			st.push(r)
		} else {
			if string(r) == ")" && string(st.peek()) == ")" {
				pop := st.pop()
				if !st.isEmpty() && string(st.peek()) == "(" {
					st.pop()
				} else {
					st.push(pop)
				}
			} else {
				st.push(r)
			}
		}
	}

	for !st.isEmpty() {
		if string(st.peek()) == ")" {
			st.pop()

			if !st.isEmpty() {
				if string(st.peek()) == ")" {
					st.pop()

					if !st.isEmpty() {
						if string(st.peek()) != "(" {
							out++
						}
					} else {
						out++
					}
				} else {
					st.pop()
					out += 1
				}
			} else {
				out += 2
			}
		} else {
			st.pop()
			out += 2
		}
	}

	return out
}
