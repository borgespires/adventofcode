extern crate regex;

use regex::Regex;
use std::fs::File;
use std::io::prelude::*;

struct Fabric {
    width: usize,
    _height: usize,
    vec: Vec<i32>,
}

impl Fabric {
    fn new(width: usize, height: usize) -> Fabric {
        Fabric {
            width: width,
            _height: height,
            vec: vec![0; width * height],
        }
    }

    fn claim_square(&mut self, (x, y): (usize, usize)) {
        self.vec[y * self.width + x] += 1;
    }

    fn get_square_claims(&self, (x, y): (usize, usize)) -> i32 {
        return self.vec[y * self.width + x];
    }

    fn apply(&mut self, claim: &Claim) {
        let (x, y) = claim.start;

        for y_iter in 0..claim.height {
            for x_iter in 0..claim.width {
                self.claim_square((x + x_iter, y + y_iter));
            }
        }
    }

    fn is_uncontested(&self, claim: &Claim) -> bool {
        let (x, y) = claim.start;

        for y_iter in 0..claim.height {
            for x_iter in 0..claim.width {
                if self.get_square_claims((x + x_iter, y + y_iter)) != 1 {
                    return false;
                }
            }
        }

        return true;
    }

    fn contested_area(&self) -> usize {
        return self
            .vec
            .iter()
            .filter(|&number_of_claims| *number_of_claims > 1)
            .count();
    }
}

struct Claim {
    id: usize,
    start: (usize, usize),
    width: usize,
    height: usize,
}

impl Claim {
    fn parse(s: &str) -> Claim {
        let re = Regex::new(r"#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)").unwrap();
        let cap = re.captures(&s).unwrap();

        fn to_usize(s: &str) -> usize {
            return s.parse::<usize>().unwrap();
        }

        Claim {
            id: to_usize(&cap[1]),
            start: (to_usize(&cap[2]), to_usize(&cap[3])),
            width: to_usize(&cap[4]),
            height: to_usize(&cap[5]),
        }
    }
}

fn find_uncontested_claim(fabric: &Fabric, claims: &Vec<Claim>) -> usize {
    return claims
        .iter()
        .find(|&claim| fabric.is_uncontested(claim))
        .unwrap()
        .id;
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let mut fabric: Fabric = Fabric::new(1000, 1000);
    let claims: Vec<Claim> = contents.lines().map(|s| Claim::parse(s)).collect();

    for claim in claims.iter() {
        fabric.apply(claim);
    }

    println!("{}", fabric.contested_area());
    println!("{}", find_uncontested_claim(&fabric, &claims));
}
