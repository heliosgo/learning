#![allow(unused)]
/*
503. Next Greater Element II
Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.

Example 1:

Input: nums = [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number.
The second 1's next greater number needs to search circularly, which is also 2.
*/

struct Solution {}

impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut res = vec![-1; n];
        let mut stack = Vec::new();
        for i in (0..2 * n).rev() {
            while !stack.is_empty() && *stack.last().unwrap() <= nums[i % n] {
                stack.pop();
            }
            if !stack.is_empty() {
                res[i % n] = *stack.last().unwrap();
            }
            stack.push(nums[i % n]);
        }

        res
    }
}
