#![allow(unused)]

fn quick_sort(nums: &mut Vec<i32>, left: usize, right: usize) {
    if left >= right {
        return;
    }
    let (mut i, mut j) = (left, right);
    let x = nums[i + j >> 1];
    while i < j {
        while i < j && nums[j] >= x {
            j -= 1;
        }
        while i < j && nums[i] < x {
            i += 1;
        }
        if i < j {
            nums.swap(i, j);
        }
    }
    quick_sort(nums, left, j);
    quick_sort(nums, j + 1, right);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_quick_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        let r = nums.len();
        quick_sort(&mut nums, 0, r - 1);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        let r2 = nums2.len();
        quick_sort(&mut nums2, 0, r2 - 1);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
