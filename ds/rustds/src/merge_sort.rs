#![allow(unused)]

fn merge_sort(nums: &mut Vec<i32>, l: usize, r: usize) {
    if l >= r {
        return;
    }
    let m = l + r >> 1;
    merge_sort(nums, l, m);
    merge_sort(nums, m + 1, r);
    let (mut i, mut j, mut k) = (l, m + 1, 0);
    let mut tmp: Vec<i32> = vec![0; r - l + 1];
    while i <= m && j <= r {
        if nums[i] < nums[j] {
            tmp[k] = nums[i];
            i += 1;
        } else {
            tmp[k] = nums[j];
            j += 1;
        }
        k += 1;
    }
    while i <= m {
        tmp[k] = nums[i];
        i += 1;
        k += 1;
    }
    while j <= r {
        tmp[k] = nums[j];
        j += 1;
        k += 1;
    }
    for i in 0..k {
        nums[l + i] = tmp[i];
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_merge_sort() {
        let mut nums = vec![4, 5, 2, 1, 3, 6];
        let r = nums.len();
        merge_sort(&mut nums, 0, r - 1);
        assert_eq!(nums, vec![1, 2, 3, 4, 5, 6]);

        let mut nums2 = vec![1, 1, 1, 1, 1, 1];
        let r2 = nums2.len();
        merge_sort(&mut nums2, 0, r2 - 1);
        assert_eq!(nums2, vec![1, 1, 1, 1, 1, 1]);
    }
}
