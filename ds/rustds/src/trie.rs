#![allow(unused)]

#[derive(Clone)]
struct Trie {
    children: Vec<Option<Trie>>,
    end: bool,
}

impl Trie {
    fn new() -> Self {
        Trie {
            children: vec![None; 26],
            end: false,
        }
    }

    fn insert(&mut self, word: String) {
        let mut p = self;
        for c in word.chars() {
            let idx = c as usize - 'a' as usize;
            if let None = p.children[idx] {
                p.children[idx] = Some(Trie::new());
            }
            p = p.children[idx].as_mut().unwrap();
        }
        p.end = true;
    }

    fn search(&mut self, word: String) -> bool {
        let mut p = self;
        for c in word.chars() {
            let idx = c as usize - 'a' as usize;
            if let None = p.children[idx] {
                return false;
            }
            p = p.children[idx].as_mut().unwrap();
        }
        p.end
    }

    fn start_with(&mut self, word: String) -> bool {
        let mut p = self;
        for c in word.chars() {
            let idx = c as usize - 'a' as usize;
            if let None = p.children[idx] {
                return false;
            }
            p = p.children[idx].as_mut().unwrap();
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_trie() {
        let mut trie = Trie::new();
        trie.insert("apple".to_string());
        assert_eq!(trie.search("apple".to_string()), true);
        assert_eq!(trie.start_with("app".to_string()), true);
        assert_eq!(trie.start_with("appe".to_string()), false);
        trie.insert("cat".to_string());
        assert_eq!(trie.search("cat".to_string()), true);
        assert_eq!(trie.start_with("ca".to_string()), true);
    }
}
