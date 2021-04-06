package questions

var count int
var res int
var t int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	t = k
	count = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	count += 1
	if count == t {
		res = root.Val
		return
	}
	dfs(root.Right)

}
