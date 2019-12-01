mod coordinate;

use coordinate::Coordinate;
use std::fs::File;
use std::io::prelude::*;

struct TrackSystemSimulator {
    carts: Vec<Coordinate>,
    
}

fn main() {
    println!("Hello, world!");

    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    println!("{}", contents);
}
