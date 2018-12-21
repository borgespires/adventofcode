use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;

#[derive(Debug)]
struct Marble {
    value: usize,
    next: usize,
    prev: usize,
}

const UNREACHABLE_MARBLE: Marble = Marble {
    value: 0,
    prev: 0,
    next: 0,
};

struct Circle {
    marbles: Vec<Marble>,
    current: usize,
}

impl Circle {
    fn new() -> Circle {
        Circle {
            marbles: vec![Marble {
                value: 0,
                next: 0,
                prev: 0,
            }],
            current: 0,
        }
    }

    fn counter_clockwise(&self, delta: u32) -> usize {
        let mut marble = self.current;
        for _ in 0..delta {
            marble = self.marbles[marble].prev;
        }
        marble
    }

    fn insert(&mut self, marble: usize) {
        let clockwise_1 = self.marbles[self.current].next;
        let clockwise_2 = self.marbles[clockwise_1].next;

        self.marbles[clockwise_1].next = marble;
        self.marbles[clockwise_2].prev = marble;
        self.marbles.push(Marble {
            value: marble,
            prev: clockwise_1,
            next: clockwise_2,
        });

        self.current = marble;
    }

    fn remove(&mut self, marble: usize) {
        let next = self.marbles[marble].next;
        let prev = self.marbles[marble].prev;
        self.marbles[next].prev = prev;
        self.marbles[prev].next = next;

        self.marbles.push(UNREACHABLE_MARBLE);
        self.current = next;
    }
}

struct Game {
    score: HashMap<u32, u32>,
    circle: Circle,
}

impl Game {
    fn new() -> Game {
        Game {
            score: HashMap::new(),
            circle: Circle::new(),
        }
    }

    fn simulate(&mut self, n_players: u32, n_marbles: u32) -> HashMap<u32, u32> {
        for (player, marble) in (0..n_players).cycle().zip(1..n_marbles + 1) {
            self.turn(player, marble);
        }
        self.score.clone()
    }

    fn turn(&mut self, player: u32, marble: u32) {
        if marble % 23 == 0 {
            let to_remove = self.circle.counter_clockwise(7);
            *self.score.entry(player).or_insert(0) += marble + to_remove as u32;
            self.circle.remove(to_remove);
        } else {
            self.circle.insert(marble as usize);
        }
    }
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    // println!("{:?}", Game::new().simulate(10, 1618).values().max().unwrap()); // 8317
    // println!("{:?}", Game::new().simulate(13, 7999).values().max().unwrap()); // 14637
    // println!("{:?}", Game::new().simulate(17, 1104).values().max().unwrap()); // 2764
    // println!("{:?}", Game::new().simulate(21, 6111).values().max().unwrap()); // 54718
    // println!("{:?}", Game::new().simulate(30, 5807).values().max().unwrap()); // 37305

    println!(
        "{:?}",
        Game::new().simulate(428, 70825).values().max().unwrap()
    );
    println!(
        "{:?}",
        Game::new().simulate(428, 7082500).values().max().unwrap()
    );
}
