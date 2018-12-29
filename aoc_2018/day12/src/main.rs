use std::collections::HashMap;
use std::collections::VecDeque;
use std::fmt;
use std::fs::File;
use std::io::prelude::*;

struct PlantSimulator {
    state: Vec<bool>,
    rules: HashMap<Vec<bool>, bool>,
    min_pot: i64,
    generation: i64,
}

impl PlantSimulator {
    fn new(setup: &str, notes: Vec<&str>) -> PlantSimulator {
        let state = setup.chars().map(|c| c == '#').collect::<Vec<_>>();
        let mut rules = HashMap::<Vec<bool>, bool>::new();

        for note in notes.into_iter() {
            let from = note.split(" => ").nth(0).unwrap();
            let to = match note.split(" => ").nth(1).unwrap() {
                "." => false,
                "#" => true,
                _ => panic!("parsing error"),
            };

            rules.insert(from.chars().map(|c| c == '#').collect(), to);
        }

        PlantSimulator {
            state: state,
            rules: rules,
            min_pot: 0,
            generation: 0,
        }
    }

    fn next(&mut self) {
        fn pad_with_empty_pots(state: &Vec<bool>) -> Vec<bool> {
            let mut padded = vec![false, false, false, false];
            padded.extend(state);
            padded.extend(&vec![false, false, false, false]);
            padded
        }

        fn trim_empty_pots(mut state: VecDeque<bool>) -> (Vec<bool>, i64) {
            const POTS_ADDED_BY_PAD: i64 = 2;
            let mut trimmed_pots = 0;

            while let Some(false) = state.front() {
                state.pop_front();
                trimmed_pots += 1;
            }

            while let Some(false) = state.back() {
                state.pop_back();
            }

            (state.into_iter().collect(), trimmed_pots - POTS_ADDED_BY_PAD)
        }

        let next_gen = pad_with_empty_pots(&self.state)
            .windows(5)
            .into_iter()
            .map(|w| *self.rules.get(w).unwrap_or(&false))
            .collect::<VecDeque<_>>();

        let (trimmed, pot_diff) = trim_empty_pots(next_gen);

        self.state = trimmed;
        self.min_pot += pot_diff;
        self.generation += 1;
    }

    fn generation_value(&self) -> i64 {
        let max_pot = self.state.len() as i64 + self.min_pot;
        self.state
            .iter()
            .zip(self.min_pot..max_pot)
            .map(|(has_plant, pot_number)| pot_number * *has_plant as i64)
            .sum::<i64>()
    }

    fn value_after(&mut self, generations: i64) -> i64 {
        let mut last_value = self.generation_value();
        let mut last_delta = -1;

        for _ in 0..generations {
            self.next();
            let value = self.generation_value();
            let delta = value - last_value;

            if delta == last_delta {
                return last_value + (generations - self.generation + 1) * delta;
            }

            last_value = value;
            last_delta = delta;
        }

        last_value
    }
}

impl fmt::Debug for PlantSimulator {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let state_str: String = self
            .state
            .iter()
            .map(|has_plant| match has_plant {
                false => ".",
                true => "#",
            }).collect();

        write!(
            f,
            "{} |{}| => {}",
            self.generation,
            self.generation_value(),
            state_str
        )
    }
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let lines = contents.lines().collect::<Vec<_>>();

    let setup = lines[0].split(": ").nth(1).unwrap();
    let notes = lines[2..].to_vec();
    let mut simulator = PlantSimulator::new(setup, notes);

    println!("{:?}", simulator.value_after(20));
    println!("{:?}", simulator.value_after(50000000000));
}
