use std::cmp::max;
use std::collections::HashMap;
use std::io::stdin;

fn longest_unique_interval(seq: &[u32]) -> usize {
    let mut longest = 0;
    let mut left = 0;
    let mut map = HashMap::<u32, usize>::new();

    for (right, &val) in seq.iter().enumerate() {
        if let Some(&pos) = map.get(&val) {
            if pos >= left {
                left = pos + 1;
            }
        }
        map.insert(val, right);
        longest = max(longest, right - left + 1);
    }

    longest
}

fn main() {
    let mut lines = stdin().lines();
    lines.next();
    let seq: Vec<u32> = lines
        .next()
        .unwrap()
        .unwrap()
        .split_whitespace()
        .map(|c| c.parse().unwrap())
        .collect();

    println!("{}", longest_unique_interval(&seq));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_stupid() {
        assert_eq!(longest_unique_interval(&vec![1, 2]), 2);
    }

    #[test]
    fn test_example() {
        assert_eq!(longest_unique_interval(&vec![1, 2, 1, 3, 2, 7, 4, 2]), 5);
    }
}
