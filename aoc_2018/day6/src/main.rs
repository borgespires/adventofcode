mod coordinate;

use coordinate::Coordinate;
use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::str::FromStr;

/*
 * Find the smaller rectangle that bounds the all set of points.
 * Returns top-left and bottom-right corners of the bounding box.
*/
fn bounding_rectangle(coordinates: &Vec<Coordinate>) -> (Coordinate, Coordinate) {
    (
        Coordinate {
            x: coordinates.iter().min_by_key(|&c| c.x).unwrap().x,
            y: coordinates.iter().min_by_key(|&c| c.y).unwrap().y,
        },
        Coordinate {
            x: coordinates.iter().max_by_key(|&c| c.x).unwrap().x,
            y: coordinates.iter().max_by_key(|&c| c.y).unwrap().y,
        },
    )
}

fn nearest_coordinate<'a>(
    coordinates: &'a Vec<Coordinate>,
    coordinate: &Coordinate,
) -> Option<&'a Coordinate> {
    fn manhattan_distance(coordinate_a: &Coordinate, coordinate_b: &Coordinate) -> i32 {
        (coordinate_a.x - coordinate_b.x).abs() + (coordinate_a.y - coordinate_b.y).abs()
    }

    let nearest: Vec<(&Coordinate, i32)> = coordinates
        .iter()
        .map(|c| (c, manhattan_distance(c, coordinate)))
        .collect::<Vec<(&Coordinate, i32)>>();

    let first = nearest.iter().min_by_key(|(c, d)| d).unwrap();
    let others: Vec<&(&Coordinate, i32)> = nearest.iter().filter(|(c, d)| *d == first.1).collect();

    if others.len() > 1 {
        None
    } else {
        Some(first.0)
    }
}

fn calculate_areas<'a>(
    coordinates: &'a Vec<Coordinate>,
    boundary: &(Coordinate, Coordinate),
) -> HashMap<&'a Coordinate, Vec<Coordinate>> {
    let mut areas = HashMap::new();
    let (top_left, bottom_right) = boundary;

    for x in top_left.x..bottom_right.x + 1 {
        for y in top_left.y..bottom_right.y + 1 {
            let contested = Coordinate::new(x, y);
            match nearest_coordinate(&coordinates, &contested) {
                Some(coordinate) => areas
                    .entry(coordinate)
                    .or_insert_with(Vec::new)
                    .push(contested),
                None => (),
            }
        }
    }

    areas
}

fn near_area(coordinates: &Vec<Coordinate>, boundary: &(Coordinate, Coordinate)) -> i32 {
    let (top_left, bottom_right) = boundary;

    fn manhattan_distance(coordinate_a: &Coordinate, coordinate_b: &Coordinate) -> i32 {
        (coordinate_a.x - coordinate_b.x).abs() + (coordinate_a.y - coordinate_b.y).abs()
    }

    fn sum_all(coordinates: &Vec<Coordinate>, coordinate: &Coordinate) -> i32 {
        coordinates
            .iter()
            .map(|c| manhattan_distance(c, coordinate))
            .sum()
    }

    let mut counter = 0;

    for x in top_left.x..bottom_right.x + 1 {
        for y in top_left.y..bottom_right.y + 1 {
            let contested = Coordinate::new(x, y);
            if sum_all(&coordinates, &contested) < 10000 {
                counter += 1;
            }
        }
    }

    counter
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

    let boundary = bounding_rectangle(&coordinates);

    let areas = calculate_areas(&coordinates, &boundary);

    let res: HashMap<&Coordinate, usize> = areas
        .into_iter()
        .filter_map(|(key, value)| {
            if value
                .iter()
                .any(|v| v.y == boundary.0.y || v.y == boundary.1.y)
            {
                None
            } else if value
                .iter()
                .any(|v| v.x == boundary.0.x || v.x == boundary.1.x)
            {
                None
            } else if value
                .iter()
                .any(|v| v.y == boundary.0.y || v.y == boundary.1.y)
            {
                None
            } else {
                Some((key.to_owned(), value.len()))
            }
        }).collect();

    let m = res.values().max().unwrap();
    let s = near_area(&coordinates, &boundary);


    assert!(*m as i32 == 4976);
    assert!(s == 46462);

    println!("{:?}", coordinates);
    println!("{:?}", boundary);
    // println!("{:?}", areas);
    println!("{:?}", res);
    println!("{:?}", m);

    println!("{:?}", s);
}
