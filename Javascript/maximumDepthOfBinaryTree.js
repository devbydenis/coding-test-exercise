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

const root = new TreeNode(10)
root.left = new TreeNode(2)
root.right = new TreeNode(5)

function maximumDepthOfBinaryTree(root) {
  if (!root) {
    return 0;
  }

  const leftDepth = maximumDepthOfBinaryTree(root.left);
  const rightDepth = maximumDepthOfBinaryTree(root.right);

  return Math.max(leftDepth, rightDepth) + 1;
}


console.time()
console.log(maximumDepthOfBinaryTree(root))
console.timeEnd()