#![allow(unused)]

fn selection_sort(nums: &mut Vec<i32>) {
    let n = nums.len();
    for i in 0..n {
        let mut min_index = i;
        for j in i + 1..n {
            if nums[j] < nums[min_index] {
                min_index = j;
            }
        }
        nums.swap(min_index, i);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_selection_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        selection_sort(&mut nums);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        selection_sort(&mut nums2);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
