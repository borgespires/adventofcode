extern crate lazy_static;
extern crate regex;

use lazy_static::lazy_static;
use regex::Regex;
use std::fs::File;
use std::io::prelude::*;
use std::num::ParseIntError;
use std::str::FromStr;

#[derive(Debug)]
struct Light {
    x: i64,
    y: i64,
    vx: i64,
    vy: i64,
}

impl FromStr for Light {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static! {
            static ref RE: Regex =
                Regex::new(r"^position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>$")
                    .unwrap();
        }

        let caps = RE.captures(s).unwrap();
        Ok(Light {
            x: caps[1].parse()?,
            y: caps[2].parse()?,
            vx: caps[3].parse()?,
            vy: caps[4].parse()?,
        })
    }
}

#[derive(Debug)]
struct SkySimulator {
    lights: Vec<Light>,
}

impl SkySimulator {
    fn new(lights: Vec<Light>) -> SkySimulator {
        SkySimulator { lights }
    }

    fn boundary(lights: &Vec<Light>) -> ((i64, i64), (i64, i64)) {
        (
            (
                lights.iter().min_by_key(|&c| c.x).unwrap().x,
                lights.iter().min_by_key(|&c| c.y).unwrap().y,
            ),
            (
                lights.iter().max_by_key(|&c| c.x).unwrap().x,
                lights.iter().max_by_key(|&c| c.y).unwrap().y,
            ),
        )
    }

    fn area(lights: &Vec<Light>) -> i64 {
        let ((min_x, min_y), (max_x, max_y)) = SkySimulator::boundary(&lights);
        (max_x - min_x).abs() * (max_y - min_y).abs()
    }

    fn tick(&mut self, seconds: i64) -> Vec<Light> {
        self.lights
            .iter()
            .map(|light| Light {
                x: light.x + seconds * light.vx,
                y: light.y + seconds * light.vy,
                vx: light.vx,
                vy: light.vy,
            }).collect()
    }

    fn print(lights: Vec<Light>) {
        let ((min_x, min_y), (max_x, max_y)) = SkySimulator::boundary(&lights);

        for y in min_y..max_y + 1 {
            for x in min_x..max_x + 1 {
                if lights.iter().any(|light| light.x == x && light.y == y) {
                    print!("#");
                } else {
                    print!(".");
                }
            }
            print!("\n");
        }
    }

    fn convergence_time(&mut self) -> i64 {
        let mut time = 0;
        let mut current_area = SkySimulator::area(&self.lights);

        loop {
            let next_position = self.tick(time + 1);
            let next_area = SkySimulator::area(&next_position);

            if next_area > current_area {
                return time;
            }

            current_area = next_area;
            time += 1;
        }
    }
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let lights: Vec<Light> = contents
        .lines()
        .map(|s| Light::from_str(s).unwrap())
        .collect();

    let mut simulator = SkySimulator::new(lights);
    let time = simulator.convergence_time();

    println!("{}", time);
    SkySimulator::print(simulator.tick(time));
}
