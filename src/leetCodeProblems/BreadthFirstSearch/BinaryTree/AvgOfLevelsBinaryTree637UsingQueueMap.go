package main

/*
- Leetcode - https://leetcode.com/problems/average-of-levels-in-binary-tree
- Time - O(n)
- Space - O(n)
*/

type TreeNodeLevel struct {
	level int
	node  *TreeNode
}

type LevelSumCount struct {
	count int
	sum   int
}

type tnlQueue []TreeNodeLevel

func (q *tnlQueue) pop() TreeNodeLevel {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q *tnlQueue) push(t TreeNodeLevel) {
	*q = append(*q, t)
}

func (q tnlQueue) isEmpty() bool {
	return len(q) == 0
}

func averageOfLevelsUsingQueueAndMap(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	q := new(tnlQueue)
	m := make(map[int]LevelSumCount)

	q.push(TreeNodeLevel{0, root})

	for !q.isEmpty() {
		tnl := q.pop()

		if k, ok := m[tnl.level]; ok {
			k.count++
			k.sum += tnl.node.Val
			m[tnl.level] = k
		} else {
			m[tnl.level] = LevelSumCount{1, tnl.node.Val}
		}

		if tnl.node.Left != nil {
			q.push(TreeNodeLevel{tnl.level + 1, tnl.node.Left})
		}

		if tnl.node.Right != nil {
			q.push(TreeNodeLevel{tnl.level + 1, tnl.node.Right})
		}
	}

	out := make([]float64, len(m))

	for i, v := range m {
		out[i] = float64(v.sum) / float64(v.count)
	}
	return out
}
