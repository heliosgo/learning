#![allow(unused)]
/*
102. Binary Tree Level Order Traversal
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
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
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut res: Vec<Vec<i32>> = Vec::new();
        if let Some(root) = root.clone() {
            let mut cur: Vec<Rc<RefCell<TreeNode>>> = vec![root];
            while cur.len() > 0 {
                let mut next: Vec<Rc<RefCell<TreeNode>>> = Vec::new();
                let mut vals: Vec<i32> = Vec::new();
                for node in &cur {
                    vals.push(node.borrow().val);
                    if let Some(left) = node.borrow().left.clone() {
                        next.push(left);
                    }
                    if let Some(right) = node.borrow().right.clone() {
                        next.push(right);
                    }
                }
                res.push(vals);
                cur = next;
            }
        }

        res
    }
}
