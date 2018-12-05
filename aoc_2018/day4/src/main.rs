extern crate regex;
#[macro_use]
extern crate lazy_static;

mod datetime;
mod event;

use event::Event;
use std::cmp::Ordering;
use std::fs::File;
use std::io::prelude::*;
use std::str::FromStr;

fn event_datetime(ev1: &Event, ev2: &Event) -> Ordering {
    return ev1.datetime.cmp(&ev2.datetime);
}

// fn sleep_records(events: &Vec<Event>) -> HashMap<u32, &[u32]> {
//     let mut records: HashMap<u32, [u32]> = HashMap::new();

//     records.insert(1, [0; 60]);

//     return records;
// }

/*
    iterate over events
    maintain state of guard
    each guard as a fixed array of minutes
    check diff between Asleep - wakeUp events
    increment each minute
**/

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let mut events: Vec<Event> = contents
        .lines()
        .map(|s| Event::from_str(s))
        .map(|r| r.unwrap())
        .collect();

    events.sort_by(event_datetime);

    for e in events.iter() {
        println!("{:?}", e);
    }
}
