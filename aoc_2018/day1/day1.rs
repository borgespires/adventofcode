use std::collections::HashSet;
use std::fs::File;
use std::io::prelude::*;

fn resulting_freq(changes: &[i32]) -> i32 {
    return changes.iter().sum::<i32>();
}

fn first_repeating_freq(changes: &[i32]) -> i32 {
    let mut seen = HashSet::new();
    let mut freq = 0;
    seen.insert(0);

    // execute a block of code over and over again forever or until explicitly tell it to stop.
    loop {
        for c in changes {
            freq += c;
            if seen.contains(&freq) {
                return freq;
            }

            seen.insert(freq);
        }
    }
}

fn main() {
    let mut f = File::open("input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let changes: Vec<i32> = contents
        .split("\n")
        .map(|x| x.parse::<i32>().unwrap())
        .collect();

    println!("{}", resulting_freq(&changes));
    println!("{}", first_repeating_freq(&changes));
}
