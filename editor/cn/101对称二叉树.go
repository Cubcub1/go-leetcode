// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1966 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// [1,2,2,3,4,4,3]
	node2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	node3 := &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	// node3 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	result := isSymmetric(node1)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var symmetry func(left *TreeNode, right *TreeNode) bool
	symmetry = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return symmetry(left.Left, right.Right) && symmetry(left.Right, right.Left)
	}
	return symmetry(root.Left, root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() != 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
