#![allow(unused)]

fn heap_sort(nums: &mut Vec<i32>) {
    let sz = nums.len();
    for i in (0..(sz - 1) / 2).rev() {
        down(nums, i, sz);
    }
    for i in (0..sz - 1).rev() {
        nums.swap(0, i);
        down(nums, 0, i);
    }
}

fn down(nums: &mut Vec<i32>, k: usize, sz: usize) {
    let mut t = k;
    if 2 * k < sz && nums[t] < nums[2 * k] {
        t = 2 * k;
    }
    if 2 * k + 1 < sz && nums[t] < nums[2 * k + 1] {
        t = 2 * k + 1;
    }
    if t != k {
        nums.swap(k, t);
        down(nums, t, sz);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_heap_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        heap_sort(&mut nums);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        heap_sort(&mut nums2);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
