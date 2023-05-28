pub fn binary_search<T>(value: T, vec: &Vec<T>) -> Option<usize>
where T: std::cmp::PartialOrd + Default {
    let mut min = 0;
    let mut max = vec.len();

    while min < max {
        let mid = min + (max - min) / 2;
        let mid_value = &vec[mid];
        if value < *mid_value { max = mid; }
        else if value > *mid_value { min = mid + 1; }
        else { return Some(mid); }
    } 
    None
}

// test
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_binary_search() {
        let vec = vec![1, 2, 3, 4, 5];
        assert_eq!(binary_search(0, &vec), None);
        assert_eq!(binary_search(1, &vec), Some(0));
        assert_eq!(binary_search(2, &vec), Some(1));
        assert_eq!(binary_search(3, &vec), Some(2));
        assert_eq!(binary_search(4, &vec), Some(3));
        assert_eq!(binary_search(5, &vec), Some(4));
        assert_eq!(binary_search(6, &vec), None);
    }
}
