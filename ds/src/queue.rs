#![allow(unused)]

struct Queue<T> {
    data: Vec<T>,
}

impl<T> Queue<T> {
    fn new() -> Self {
        Queue { data: Vec::new() }
    }

    fn push(&mut self, item: T) {
        self.data.push(item)
    }

    fn pop(&mut self) -> Option<T> {
        if self.data.is_empty() {
            return None;
        }

        Some(self.data.remove(0))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_queue() {
        let mut queue = Queue::new();
        queue.push(1);
        queue.push(2);
        queue.push(4);
        assert_eq!(queue.pop(), Some(1));
        assert_eq!(queue.pop(), Some(2));
        queue.push(5);
        assert_eq!(queue.pop(), Some(4));
        assert_eq!(queue.pop(), Some(5));
        assert_eq!(queue.pop(), None);
    }
}
