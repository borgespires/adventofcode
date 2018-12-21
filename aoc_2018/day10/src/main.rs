use std::fs::File;
use std::io::prelude::*;

impl FromStr for Point {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let v = s.split(", ").collect::<Vec<_>>();
        Ok(Point {
            x: v[0].parse()?,
            y: v[1].parse()?,
        })
    }
}

fn main() {
    println!("Hello, world!");

    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");
    
    let points: Vec<Point> = contents
        .lines()
        .map(|s| Point::from_str(s).unwrap())
        .collect();

    // Get how far apart are the points
    // while this square doe not grow continue moving points
    // stop moving when the points stop converging
}
