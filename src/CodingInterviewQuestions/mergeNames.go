package main

/*
Implement the uniqueNames function. When passed two slices of names, it will return a slice containing the names that appear in either or both slices. The returned slice should have no duplicates.

For example, calling uniqueNames([]string{"Ava", "Emma", "Olivia"}, []string{"Olivia", "Sophia", "Emma"}) should return a slice containing Ava, Emma, Olivia, and Sophia in any order.

Reference - https://www.testdome.com/tests/golang-online-test/123
*/
func uniqueNames(a, b []string) []string {
	var result []string

	m := make(map[string]bool)

	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}

	for k, _ := range m {
		result = append(result, k)
	}

	return result
}

// func main() {
// 	// should print Ava, Emma, Olivia, Sophia
// 	fmt.Println(uniqueNames(
// 		[]string{"Ava", "Emma", "Olivia"},
// 		[]string{"Olivia", "Sophia", "Emma"}))
// }
