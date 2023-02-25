#![allow(unused)]

fn binary_search1(nums: &Vec<i32>, target: i32) -> bool {
    let (mut l, mut r) = (0, nums.len() - 1);
    while l < r {
        let m = l + r >> 1;
        if nums[m] < target {
            l = m + 1;
        } else {
            r = m;
        }
    }

    nums[l] == target
}

fn binary_search2(nums: &Vec<i32>, target: i32) -> bool {
    let (mut l, mut r) = (0, nums.len() - 1);
    while l < r {
        let m = l + r + 1 >> 1;
        if nums[m] < target {
            r = m - 1;
        } else {
            l = m;
        }
    }

    nums[l] == target
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_binary_search() {
        let nums = vec![1, 3, 9, 11, 13, 19];
        assert_eq!(binary_search1(&nums, 13), true);
        assert_eq!(binary_search1(&nums, 19), true);
        assert_eq!(binary_search1(&nums, 14), false);

        let nums2 = vec![10, 4, 3, 2, 1, 0];
        assert_eq!(binary_search2(&nums2, 10), true);
        assert_eq!(binary_search2(&nums2, 0), true);
        assert_eq!(binary_search2(&nums2, 14), false);
    }
}
