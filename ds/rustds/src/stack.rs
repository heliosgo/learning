#![allow(unused)]

struct Stack<T> {
    data: Vec<T>,
}

impl<T> Stack<T> {
    fn new() -> Self {
        Stack { data: Vec::new() }
    }

    fn push(&mut self, item: T) {
        self.data.push(item)
    }

    fn pop(&mut self) -> Option<T> {
        self.data.pop()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_stack() {
        let mut stack = Stack::new();
        stack.push(1);
        stack.push(2);
        stack.push(4);
        assert_eq!(stack.pop(), Some(4));
        assert_eq!(stack.pop(), Some(2));
        stack.push(5);
        assert_eq!(stack.pop(), Some(5));
        assert_eq!(stack.pop(), Some(1));
        assert_eq!(stack.pop(), None);
    }
}
