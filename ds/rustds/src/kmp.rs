#![allow(unused)]

fn kmp(s: String, p: String) -> Vec<usize> {
    let (n, m) = (s.len(), p.len());
    let (ts, tp) = (format!(" {}", s), format!(" {}", p));
    let (bs, bp) = (ts.as_bytes(), tp.as_bytes());
    let mut ne: Vec<usize> = vec![0; m + 1];
    let mut j = 0;
    for i in 2..m + 1 {
        while j > 0 && bp[i] != bp[j + 1] {
            j = ne[j];
        }
        if bp[i] == bp[j + 1] {
            j += 1;
        }
        ne[i] = j;
    }

    let mut res: Vec<usize> = Vec::new();
    j = 0;
    for i in 1..n + 1 {
        while j > 0 && bs[i] != bp[j + 1] {
            j = ne[j];
        }
        if bs[i] == bp[j + 1] {
            j += 1;
        }
        if j == m {
            res.push(i - m);
            j = ne[j];
        }
    }

    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_kmp() {
        let s1 = String::from("helloworldworld");
        let p1 = String::from("world");
        assert_eq!(kmp(s1, p1), vec![5, 10]);
        let s2 = String::from("aaaaaaaa");
        let p2 = String::from("aaa");
        assert_eq!(kmp(s2, p2), vec![0, 1, 2, 3, 4, 5]);
        let s3 = String::from("aaaaaaaa");
        let p3 = String::from("a");
        assert_eq!(kmp(s3, p3), vec![0, 1, 2, 3, 4, 5, 6, 7]);
    }
}
