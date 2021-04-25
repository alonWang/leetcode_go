package questions

var result = make([]int, 0)

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	if root == nil {
		return result[:]
	}

	doInorderTraversal(root)
	return result[:]
}

func doInorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	doInorderTraversal(root.Left)
	result = append(result, root.Val)
	doInorderTraversal(root.Right)
}
