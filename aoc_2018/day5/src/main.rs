use std::fs::File;
use std::io::prelude::*;

const REACTED_UNIT: u8 = 0;

fn reacts(unit_a: u8, unit_b: u8) -> bool {
    return (unit_a as i8 - unit_b as i8).abs() == 32;
}

fn start_reaction(sample_polymer: &Vec<u8>) -> Vec<u8> {
    let mut polymer = sample_polymer.clone();

    fn previous_unit_index(polymer: &Vec<u8>, current: usize) -> usize {
        // reverse iterator .rev()
        for i in (0..current).rev() {
            if polymer[i] != REACTED_UNIT {
                return i;
            }
        }
        return current;
    }

    fn clean_reacted_units(polymer: &mut Vec<u8>) {
        polymer.retain(|&unit| unit != REACTED_UNIT);
    }

    for current in 1..polymer.len() {
        let previous = previous_unit_index(&polymer, current);

        if reacts(polymer[current], polymer[previous]) {
            polymer[previous] = REACTED_UNIT;
            polymer[current] = REACTED_UNIT;
        }
    }

    clean_reacted_units(&mut polymer);
    return polymer;
}

fn find_shortest_polymer(sample_polymer: &Vec<u8>) -> Vec<u8> {
    let mut shortest_polymer = sample_polymer.clone();
    let mut shortest = shortest_polymer.len();

    fn clean(polymer: &mut Vec<u8>, unit_type: u8) {
        let l_polarized = unit_type;
        let h_polarized = unit_type + 32;
        polymer.retain(|&u| u != l_polarized && u != h_polarized);
    }

    for unit_type in b'A'..=b'Z' {
        let mut polymer = sample_polymer.clone();
        clean(&mut polymer, unit_type);

        let result = start_reaction(&polymer);

        if result.len() < shortest {
            shortest = result.len();
            shortest_polymer = result;
        }
    }

    return shortest_polymer;
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let polymer = contents.as_bytes().to_vec();

    println!("{}", start_reaction(&polymer).iter().count());
    println!("{}", find_shortest_polymer(&polymer).iter().count());
}
