package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    result := 0
    count := 0
    //helper func
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }
        //visit left
        dfs(node.Left) // if left nil increment count 
        count++

        if count == k {
            result = node.Val
            return
        }
        dfs(node.Right) //check right if count < k
    }
    dfs(root)
    return result
}