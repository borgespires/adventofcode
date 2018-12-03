use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;

fn checksum(ids: &[&str]) -> i32 {
    fn score_box(id: &str) -> (bool, bool) {
        let mut count = HashMap::new();

        for c in id.chars() {
            *count.entry(c).or_insert(0) += 1;
        }

        let has_duplicates = count.values().any(|v| *v == 2);
        let has_triplicates = count.values().any(|v| *v == 3);

        return (has_duplicates, has_triplicates);
    }

    fn into_occurrence_count(
        (prev_dup, prev_trip): (i32, i32),
        (has_duplicates, has_triplicates): (bool, bool),
    ) -> (i32, i32) {
        let duplicates = if has_duplicates { prev_dup + 1 } else { prev_dup };
        let triplicates = if has_triplicates { prev_trip + 1 } else { prev_trip };
        return (duplicates, triplicates);
    }

    let (duplicates, triplicates) = ids
        .iter()
        .map(|id| score_box(id))
        .fold((0, 0), into_occurrence_count);

    return duplicates * triplicates;
}

fn common_box_letters(ids: &[&str]) -> Option<String> {
    fn exact_diff_of(a: &str, b: &str, n: usize) -> bool {
        return a.len() - b.len() == n;
    }
    fn equals((a, b): &(char, char)) -> bool {
        return a == b;
    }
    fn common(a: &str, b: &str) -> String {
        return a
            .chars()
            .zip(b.chars())
            .filter(equals)
            .map(|t| t.0)
            .collect::<String>();
    }

    let mut common_box_letters: Option<String> = None;
    let mut ids_to_search = ids.to_vec();

    while common_box_letters.is_none() && ids_to_search.len() > 0 {
        let current_id = ids_to_search.pop().unwrap();
        common_box_letters = ids_to_search
            .iter()
            .map(|other_id| common(current_id, other_id))
            .find(|common_letters| exact_diff_of(current_id, common_letters, 1));
    }

    return common_box_letters;
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