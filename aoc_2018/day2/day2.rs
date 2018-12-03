use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;

fn checksum(ids: &[&str]) -> i32 {
    fn box_value(id: &str) -> (i32, i32) {
        let mut count = HashMap::new();

        for c in id.chars() {
            *count.entry(c).or_insert(0) += 1;
        }

        let twos = count.values().any(|v| *v == 2) as i32;
        let threes = count.values().any(|v| *v == 3) as i32;

        return (twos, threes);
    }

    let (twos, threes) = ids
        .iter()
        .map(|id| box_value(id))
        .fold((0, 0), |(twos, threes), value| {
            (twos + value.0, threes + value.1)
        });

    return twos * threes;
}

fn common_box_letters(ids: &[&str]) -> Option<String> {
    fn diff(a: &str, b: &str) -> String {
        return a
            .chars()
            .zip(b.chars())
            .filter(|&(a, b)| a == b)
            .map(|(a, _)| a)
            .collect::<String>();
    }

    for i in 1..ids.len() {
        let (_, tail) = ids.split_at(i);
        let head = &ids[i - 1];

        if let Some(common) = tail
            .iter()
            .map(|s| diff(head, s))
            .find(|s| head.len() - s.len() == 1)
        {
            return Some(common);
        }
    }

    None
}

fn main() {
    let mut f = File::open("input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let ids: Vec<&str> = contents.lines().collect();

    println!("{}", checksum(&ids));
    println!("{}", common_box_letters(&ids).unwrap());
}
