package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root
    for root != nil {
        if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right //if curr.val < p/q = LCA is right 
        }else if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
        }else {
            return curr
        }
    }
    return nil
}