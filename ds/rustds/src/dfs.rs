/*
129. Sum Root to Leaf Numbers
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
*/
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}
//
impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;

struct Solution {}

impl Solution {
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut buf: Vec<i32> = Vec::new();
        Self::dfs(root.clone(), &mut buf, 0);

        buf.iter().sum()
    }
    pub fn dfs(node: Option<Rc<RefCell<TreeNode>>>, buf: &mut Vec<i32>, mut val: i32) {
        if let Some(n) = node {
            val = val * 10 + n.borrow().val;
            if n.borrow().left.is_none() && n.borrow().right.is_none() {
                buf.push(val);
                return;
            }

            Self::dfs(n.borrow().left.clone(), buf, val);
            Self::dfs(n.borrow().right.clone(), buf, val);
        }
    }
}
