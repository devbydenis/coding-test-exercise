class TreeNode {
  constructor(val, left, right) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

/**
 * @param {TreeNode} root
 * @return {number}
 */

function isSameTree(p, q) {
  if (p == null && q == null) {
    return true
  }

  if (p == null || q == null) {
    return false
  }

  if (p.val !== q.val) {
    return false
  }

  return isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}


const p = new TreeNode(10)
p.left = new TreeNode(2)
p.right = new TreeNode(5)

const q = new TreeNode(10)
q.left = new TreeNode(2)
q.right = new TreeNode(5)

console.time()
console.log(isSameTree(p, q))
console.timeEnd()