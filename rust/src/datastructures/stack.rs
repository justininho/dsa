pub struct Stack<T> {
    head: Link<T>,
}

type Link<T> = Option<Box<Node<T>>>;

struct Node<T> {
    elem: T,
    next: Link<T>,
}

impl<T> Stack<T> {
    pub fn new() -> Self {
        Stack { head: None }
    }

    pub fn push(&mut self, elem: T) {
        let new_node = Box::new(Node {
            elem: elem,
            next: self.head.take(),
        });

        self.head = Some(new_node);
    }

    pub fn pop(&mut self) -> Option<T> {
        self.head.take().map(|node| {
            self.head = node.next;
            node.elem
        })
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_ref().map(|node| {
            &node.elem
        })
    }

    pub fn peek_mut(&mut self) -> Option<&mut T> {
        self.head.as_mut().map(|node| {
            &mut node.elem
        })
    }

    pub fn into_iter(self) -> IntoIter<T> {
        IntoIter(self)
    }

    pub fn iter(&self) -> Iter<'_, T> {
        Iter { next: self.head.as_deref() }
    }

    pub fn iter_mut(&mut self) -> IterMut<'_, T> {
        IterMut { next: self.head.as_deref_mut() }
    }
}

impl<T> Drop for Stack<T> {
    fn drop(&mut self) {
        let mut cur_link = self.head.take();
        while let Some(mut boxed_node) = cur_link {
            cur_link = boxed_node.next.take();
        }
    }
}

pub struct IntoIter<T>(Stack<T>);

impl<T> Iterator for IntoIter<T> {
    type Item = T;
    fn next(&mut self) -> Option<Self::Item> {
        // access fields of a tuple struct numerically
        self.0.pop()
    }
}

pub struct Iter<'a, T> {
    next: Option<&'a Node<T>>,
}

impl<'a, T> Iterator for Iter<'a, T> {
    type Item = &'a T;
    fn next(&mut self) -> Option<Self::Item> {
        self.next.map(|node| {
            self.next = node.next.as_deref();
            &node.elem
        })
    }
}

pub struct IterMut<'a, T> {
    next: Option<&'a mut Node<T>>,
}

impl<'a, T> Iterator for IterMut<'a, T> {
    type Item = &'a mut T;

    fn next(&mut self) -> Option<Self::Item> {
        self.next.take().map(|node| {
            self.next = node.next.as_deref_mut();
            &mut node.elem
        })
    }
}

#[cfg(test)]
mod test {
    use super::Stack;

    #[test]
    fn basics() {
        let mut Stack = Stack::new();

        // Check empty Stack behaves right
        assert_eq!(Stack.pop(), None);

        // Populate Stack
        Stack.push(1);
        Stack.push(2);
        Stack.push(3);

        // Check normal removal
        assert_eq!(Stack.pop(), Some(3));
        assert_eq!(Stack.pop(), Some(2));

        // Push some more just to make sure nothing's corrupted
        Stack.push(4);
        Stack.push(5);

        // Check normal removal
        assert_eq!(Stack.pop(), Some(5));
        assert_eq!(Stack.pop(), Some(4));

        // Check exhaustion
        assert_eq!(Stack.pop(), Some(1));
        assert_eq!(Stack.pop(), None);
    }

    #[test]
    fn peek() {
        let mut Stack = Stack::new();
        assert_eq!(Stack.peek(), None);
        assert_eq!(Stack.peek_mut(), None);
        Stack.push(1); Stack.push(2); Stack.push(3);

        assert_eq!(Stack.peek(), Some(&3));
        assert_eq!(Stack.peek_mut(), Some(&mut 3));

        Stack.peek_mut().map(|value| {
            *value = 42
        });

        assert_eq!(Stack.peek(), Some(&42));
        assert_eq!(Stack.pop(), Some(42));
    }

    #[test]
    fn into_iter() {
        let mut Stack = Stack::new();
        Stack.push(1); Stack.push(2); Stack.push(3);

        let mut iter = Stack.into_iter();
        assert_eq!(iter.next(), Some(3));
        assert_eq!(iter.next(), Some(2));
        assert_eq!(iter.next(), Some(1));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn iter() {
        let mut Stack = Stack::new();
        Stack.push(1); Stack.push(2); Stack.push(3);

        let mut iter = Stack.iter();
        assert_eq!(iter.next(), Some(&3));
        assert_eq!(iter.next(), Some(&2));
        assert_eq!(iter.next(), Some(&1));
    }

    #[test]
    fn iter_mut() {
        let mut Stack = Stack::new();
        Stack.push(1); Stack.push(2); Stack.push(3);

        let mut iter = Stack.iter_mut();
        assert_eq!(iter.next(), Some(&mut 3));
        assert_eq!(iter.next(), Some(&mut 2));
        assert_eq!(iter.next(), Some(&mut 1));
    }
}