#![allow(unused)]

struct UnionFind {
    parent: Vec<usize>,
}

impl UnionFind {
    fn new(size: usize) -> Self {
        UnionFind {
            parent: (0..size).collect(),
        }
    }

    fn find(&mut self, x: usize) -> usize {
        if self.parent[x] != x {
            self.parent[x] = self.find(self.parent[x]);
        }

        self.parent[x]
    }

    fn union(&mut self, a: usize, b: usize) {
        let (root_x, root_y) = (self.find(a), self.find(b));
        if root_x == root_y {
            return;
        }
        self.parent[root_x] = root_y;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_union_find() {
        let mut uf = UnionFind::new(10);
        uf.union(1, 2);
        uf.union(2, 9);
        assert_eq!(uf.find(1), uf.find(9));
    }
}
