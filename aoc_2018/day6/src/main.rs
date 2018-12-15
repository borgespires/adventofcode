mod coordinate;

use coordinate::Coordinate;
use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::str::FromStr;

struct Grid {
    points_of_interest: Vec<Coordinate>,
}

impl Grid {
    fn new(coordinates: Vec<Coordinate>) -> Grid {
        Grid {
            points_of_interest: coordinates,
        }
    }

    /*
     * Find the smaller rectangle that bounds the all set of points.
     * Returns top-left and bottom-right corners of the bounding box.
     */
    fn boundary(&self) -> (Coordinate, Coordinate) {
        (
            Coordinate {
                x: self
                    .points_of_interest
                    .iter()
                    .min_by_key(|&c| c.x)
                    .unwrap()
                    .x,
                y: self
                    .points_of_interest
                    .iter()
                    .min_by_key(|&c| c.y)
                    .unwrap()
                    .y,
            },
            Coordinate {
                x: self
                    .points_of_interest
                    .iter()
                    .max_by_key(|&c| c.x)
                    .unwrap()
                    .x,
                y: self
                    .points_of_interest
                    .iter()
                    .max_by_key(|&c| c.y)
                    .unwrap()
                    .y,
            },
        )
    }

    fn areas(&self) -> HashMap<&Coordinate, usize> {
        let claimed_areas = self.claim_near_points();

        self.remove_infinite_areas(claimed_areas)
            .into_iter()
            .map(|(point, claimed_points)| (point, claimed_points.len()))
            .collect()
    }

    fn claim_near_points(&self) -> HashMap<&Coordinate, Vec<Coordinate>> {
        let mut claimed = HashMap::new();
        let (top_left, bottom_right) = self.boundary();

        for x in top_left.x..bottom_right.x + 1 {
            for y in top_left.y..bottom_right.y + 1 {
                let current = Coordinate::new(x, y);
                match self.nearest_point_of_interest(&current) {
                    Some(point) => claimed.entry(point).or_insert_with(Vec::new).push(current),
                    None => (),
                }
            }
        }

        claimed
    }

    fn nearest_point_of_interest(&self, coordinate: &Coordinate) -> Option<(&Coordinate)> {
        let distances_to_points: Vec<(&Coordinate, i32)> = self
            .points_of_interest
            .iter()
            .map(|point| (point, manhattan_distance(point, coordinate)))
            .collect::<Vec<(&Coordinate, i32)>>();

        let (nearest, min_distance) = distances_to_points.iter().min_by_key(|(_, d)| d).unwrap();
        let all_with_min_distance: Vec<&(&Coordinate, i32)> = distances_to_points
            .iter()
            .filter(|(_, d)| *d == *min_distance)
            .collect();

        if all_with_min_distance.len() > 1 {
            None
        } else {
            Some(nearest)
        }
    }

    /**
     * remove points that claimed area on the boundary
     */
    fn remove_infinite_areas<'a>(
        &self,
        claimed: HashMap<&'a Coordinate, Vec<Coordinate>>,
    ) -> HashMap<&'a Coordinate, Vec<Coordinate>> {
        let (top_left, bottom_right) = self.boundary();
        claimed
            .into_iter()
            .filter(|(_, claimed_points)| {
                !claimed_points.iter().any(|point| {
                    point.y == top_left.y
                        || point.y == bottom_right.y
                        || point.x == top_left.x
                        || point.x == bottom_right.x
                })
            }).collect()
    }

    fn safe_region(&self, max_distance_sum: i32) -> Vec<Coordinate> {
        let (top_left, bottom_right) = self.boundary();

        fn sum_all(points_of_interest: &Vec<Coordinate>, coordinate: &Coordinate) -> i32 {
            points_of_interest
                .iter()
                .map(|point| manhattan_distance(point, coordinate))
                .sum()
        }

        let mut safe_points = vec![];

        for x in top_left.x..bottom_right.x + 1 {
            for y in top_left.y..bottom_right.y + 1 {
                let current = Coordinate::new(x, y);
                if sum_all(&self.points_of_interest, &current) < max_distance_sum {
                    safe_points.push(current);
                }
            }
        }

        safe_points
    }
}

fn manhattan_distance(coordinate_a: &Coordinate, coordinate_b: &Coordinate) -> i32 {
    (coordinate_a.x - coordinate_b.x).abs() + (coordinate_a.y - coordinate_b.y).abs()
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let coordinates: Vec<Coordinate> = contents
        .lines()
        .map(|s| Coordinate::from_str(s).unwrap())
        .collect();

    let grid = Grid::new(coordinates);
    
    println!("{:?}", grid.areas().values().max().unwrap());
    println!("{:?}", grid.safe_region(10000).into_iter().count());
}
