package main

/*
- LeetCode - https://leetcode.com/problems/evaluate-division/description
- Time - O(n)
- Space - O(n)
*/

// find parent of x
func dfs(x string, p map[string]string, w map[string]float64) string {

	if p[x] != x {
		originalP := p[x]
		p[x] = dfs(p[x], p, w) // Path compression
		w[x] *= w[originalP]
	}

	return p[x]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	// p holds the parent/root representative of each variable
	p := make(map[string]string)

	// w holds the relative weight of a variable to its root
	w := make(map[string]float64)

	for _, v := range equations {
		p[v[0]] = v[0]
		p[v[1]] = v[1]
		w[v[0]] = 1.0
		w[v[1]] = 1.0
	}

	for i, v := range equations {
		p0 := dfs(v[0], p, w)
		p1 := dfs(v[1], p, w)

		if p0 != p1 {
			p[p0] = p1
			w[p0] = (w[v[1]] * values[i]) / w[v[0]]
		}
	}

	out := make([]float64, len(queries))

	for i, v := range queries {

		// if query belongs to existing parents or not
		exist := true

		if _, ok := p[v[0]]; !ok {
			exist = false
		}

		if _, ok := p[v[1]]; !ok {
			exist = false
		}

		if exist && dfs(v[0], p, w) != dfs(v[1], p, w) {
			exist = false
		}

		if !exist {
			out[i] = -1.0
			continue
		}

		out[i] = w[v[0]] / w[v[1]]
	}
	return out
}
