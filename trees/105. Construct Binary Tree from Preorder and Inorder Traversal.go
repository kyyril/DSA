package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }//if not have num -> branch is finished

    rootVal := preorder[0]
    root := &TreeNode{Val: rootVal}

    idx:=0 //find root
    for inorder[idx] != rootVal {
        idx++
    }

    //pre[9] - in[9]
    root.Left = buildTree(preorder[1:idx+1], inorder[:idx])//left child
    //pre[20,15,7] - in[15,20,7]
    root.Right = buildTree(preorder[idx+1:], inorder[idx+1:]) //right child

    return root
}