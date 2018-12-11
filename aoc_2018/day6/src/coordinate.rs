use std::hash::{Hash, Hasher};
use std::num::ParseIntError;
use std::str::FromStr;

#[derive(Debug)]
pub struct Coordinate {
    pub x: i32,
    pub y: i32,
}

impl Coordinate {
    pub fn new(x: i32, y: i32) -> Coordinate {
        Coordinate { x, y }
    }
}

impl Eq for Coordinate {}

impl PartialEq for Coordinate {
    fn eq(&self, other: &Self) -> bool {
        (self.x, self.y) == (other.x, other.y)
    }
}

impl FromStr for Coordinate {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let v = s.split(", ").collect::<Vec<_>>();
        Ok(Coordinate {
            x: v[0].parse()?,
            y: v[1].parse()?,
        })
    }
}

impl Hash for Coordinate {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.x.hash(state);
        self.y.hash(state);
    }
}
