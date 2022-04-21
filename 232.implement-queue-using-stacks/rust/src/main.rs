struct MyQueue {
    push_stack: Vec<i32>,
    pop_stack: Vec<i32>
}

impl MyQueue {

    fn new() -> Self {
        MyQueue {
            push_stack: Vec::with_capacity(100),
            pop_stack: Vec::with_capacity(100),
        }
    }

    fn push(&mut self, x: i32) {
        self.push_stack.push(x)
    }

    fn pop(&mut self) -> i32 {
        if !self.pop_stack.is_empty() {
            return self.pop_stack.pop().unwrap()
        }
        self.refill_pop_stack();
        return self.pop()
    }

    fn peek(&mut self) -> i32 {
        if !self.pop_stack.is_empty() {
            let element = self.pop_stack.last();
            match element {
                Some(number) => return *number,
                None => {}
            }
        }
        self.refill_pop_stack();
        return self.peek()
    }

    fn empty(&self) -> bool {
        self.push_stack.is_empty() && self.pop_stack.is_empty()
    }

    fn refill_pop_stack(&mut self) {
        while !self.push_stack.is_empty() {
            self.pop_stack.push(self.push_stack.pop().unwrap())
        }
    }
}

fn main() {
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let mut queue = MyQueue::new();
        queue.push(1);
        queue.push(2);
        assert_eq!(queue.peek(), 1);
        assert_eq!(queue.pop(), 1);
        assert_eq!(queue.empty(), false);
        assert_eq!(queue.pop(), 2);
        assert_eq!(queue.empty(), true);
    }
}
