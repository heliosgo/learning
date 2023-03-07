#![allow(unused)]

use rand::Rng;

fn knuth_shuffle<T>(v: &mut Vec<T>) {
    let n = v.len();
    for i in 0..n {
        let k = rand::thread_rng().gen_range(i..n);
        v.swap(i, k);
    }
}
