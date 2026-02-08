package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root} //first root [3]

	for len(queue) > 0 { //level 1
		levelSize := len(queue) //how many node at this level (1)
		level := []int{}        //list temp ([3])

		for i := 0; i < levelSize; i++ { //loop each node at this level
			node := queue[0]                //get first idx (3)
			queue = queue[1:]               //slice first node ([3]->[])
			level = append(level, node.Val) //save each node value at this level

			if node.Left != nil { //if have left chile, store to back queue
				queue = append(queue, node.Left)
			}
			if node.Right != nil { //if have right chile, store to back queue
				queue = append(queue, node.Right)
			}
		}
		//save if one level succesed
		result = append(result, level)
	}
	return result // [[3],[9,20]]
}