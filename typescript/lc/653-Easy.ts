/*
Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.
*/

function findTarget(root: TreeNode | null, k: number): boolean {
    const set = new Set<number>();
    return dfs(root, k, set);
};

function dfs(root: TreeNode | null, k: number, set: Set<number>): boolean {
    if (root === null) return false;
    if (set.has(k - root.val)) return true;
    set.add(root.val);
    return dfs(root.left, k, set) || dfs(root.right, k, set);
}