#![allow(unused)]

fn insertion_sort(nums: &mut Vec<i32>) {
    let n = nums.len();
    for i in 1..n {
        let cur = nums[i];
        let mut j = i;
        while j > 0 && nums[j - 1] > cur {
            nums[j] = nums[j - 1];
            j -= 1;
        }
        nums[j] = cur;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_insertion_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        insertion_sort(&mut nums);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        insertion_sort(&mut nums2);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
