#![allow(unused)]

fn bubble_sort(nums: &mut Vec<i32>) {
    let n = nums.len();
    for i in 0..n {
        for j in 0..n - 1 {
            if nums[j] > nums[j + 1] {
                nums.swap(j, j + 1);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_bubble_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        bubble_sort(&mut nums);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        bubble_sort(&mut nums2);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
