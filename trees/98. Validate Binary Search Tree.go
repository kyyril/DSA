package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

    var prev *int  // simpan nilai sebelumnya saat inorder

    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {

        if node == nil {
            return true
        }

        // cek subtree kiri
        if !inorder(node.Left) {
            return false
        }

        // bandingkan dengan nilai sebelumnya
        if prev != nil && node.Val <= *prev {
            return false
        }

        // update prev ke nilai sekarang
        val := node.Val
        prev = &val

        // cek subtree kanan
        return inorder(node.Right)
    }

    return inorder(root)
}
