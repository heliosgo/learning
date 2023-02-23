#![allow(unused)]

use rand::Rng;

fn reservoir_sampling(nums: Vec<i32>, k: usize) -> Vec<i32> {
    if nums.len() <= k {
        return nums;
    }

    let mut res: Vec<i32> = Vec::with_capacity(k);
    for i in 0..k {
        res[i] = nums[i];
    }
    for i in k..nums.len() {
        let m = rand::thread_rng().gen_range(0..i + 1);
        if m < k {
            res[m] = nums[i];
        }
    }

    res
}
