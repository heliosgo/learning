#![allow(unused)]

struct Node {
    val: i32,
    next: Option<Box<Node>>,
}

struct LinkedList {
    head: Option<Box<Node>>,
}

impl LinkedList {
    fn new() -> Self {
        LinkedList {
            head: Some(Box::new(Node { val: 0, next: None })),
        }
    }

    fn get(&self, index: i32) -> Option<i32> {
        if let Some(ref cur) = self.head {
            let mut cur = cur;
            let mut cur_index = 0;
            while cur_index < index {
                if let Some(ref next) = cur.next {
                    cur = next;
                    cur_index += 1;
                } else {
                    return None;
                }
            }

            return Some(cur.val);
        }
        None
    }

    fn add_at_head(&mut self, val: i32) {
        if let Some(ref mut head) = self.head {
            head.next = Some(Box::new(Node {
                val,
                next: head.next.take(),
            }))
        }
    }

    fn add_at_tail(&mut self, val: i32) {
        if let Some(ref mut cur) = self.head {
            let mut cur = cur;
            while let Some(ref mut next) = cur.next {
                cur = next;
            }
            cur.next = Some(Box::new(Node { val, next: None }));
        }
    }

    fn add_at_index(&mut self, index: i32, val: i32) {
        if let Some(ref mut cur) = self.head {
            let mut cur = cur;
            let mut cur_index = 1;
            while cur_index < index {
                if let Some(ref mut next) = cur.next {
                    cur = next;
                } else {
                    break;
                }
                cur_index += 1;
            }
            cur.next = Some(Box::new(Node {
                val,
                next: cur.next.take(),
            }))
        }
    }

    fn delete_at_index(&mut self, index: i32) {
        if let Some(ref mut cur) = self.head {
            let mut cur = cur;
            let mut cur_index = 1;
            while cur_index < index {
                if let Some(ref mut next) = cur.next {
                    cur = next;
                }
                cur_index += 1;
            }
            cur.next = cur.next.take().and_then(|a| a.next);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_linked_list() {
        let mut linked_list = LinkedList::new();
        linked_list.add_at_head(3);
        linked_list.add_at_head(2);
        linked_list.add_at_head(8);
        assert_eq!(linked_list.get(1), Some(8));
        assert_eq!(linked_list.get(2), Some(2));
        assert_eq!(linked_list.get(3), Some(3));
        linked_list.add_at_tail(1);
        linked_list.add_at_tail(9);
        assert_eq!(linked_list.get(4), Some(1));
        assert_eq!(linked_list.get(5), Some(9));
        linked_list.add_at_index(1, 11);
        linked_list.add_at_index(2, 33);
        assert_eq!(linked_list.get(1), Some(11));
        assert_eq!(linked_list.get(2), Some(33));
        assert_eq!(linked_list.get(3), Some(8));
        linked_list.delete_at_index(1);
        assert_eq!(linked_list.get(1), Some(33));
        linked_list.delete_at_index(2);
        assert_eq!(linked_list.get(2), Some(2));
    }
}
