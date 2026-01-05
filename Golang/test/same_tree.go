package test

func SameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // kalo dua-duanya nil berarti sama alias bener
		return true
	}

	if p == nil || q == nil { // kalo ada salah satu yang nil berarti salah
		return false
	}

	if p.Val != q.Val { // cek nilai dari tiap nodenya
		return false
	}

	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right) // rekursif lagi -> cek sampe kedalem-dalemnya lagi
}
