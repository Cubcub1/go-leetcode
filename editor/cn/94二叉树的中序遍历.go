// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [1]
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1460 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	result := binaryTreeInorderTraversal()
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
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var traver func(node *TreeNode)
	traver = func(node *TreeNode) {
		if node == nil {
			return
		}
		traver(node.Left)
		ans = append(ans, node.Val)
		traver(node.Right)
	}
	traver(root)
	return ans
}

// 从跟节点压栈,一直迭代到最左节点,然后弹栈
// inorderTraversal2 迭代法
func inorderTraversal2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	stack := list.New()
	cur := root
	if cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			// 说明上一个cur是最左子节点,所以当前cur才是nil
			cur = stack.Remove(stack.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}

// TODO:统一迭代法
// leetcode submit region end(Prohibit modification and deletion)
