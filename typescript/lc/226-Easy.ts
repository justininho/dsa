// Given the root of a binary tree, invert the tree, and return its root.

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

function invertTreeBfs(root: TreeNode | null): TreeNode | null {
  var queue = [root];
  while (queue.length > 0) {
    var node = queue.shift();
    if (node) {
      var t = node.left;
      node.left = node.right;
      node.right = t;
      if (node.left) queue.push(node.left);
      if (node.right) queue.push(node.right);
    }
  }
  return root;
};

function invertTreeDfs(root: TreeNode | null): TreeNode | null {
  if (root === null) return root;

  const t = root.left;
  root.left = root.right;
  root.right = t;

  invertTreeDfs(root.left);
  invertTreeDfs(root.right);
  return root;
}