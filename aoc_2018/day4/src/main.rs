extern crate regex;
#[macro_use]
extern crate lazy_static;

mod datetime;
mod event;
mod sleep_recorder;

use event::Event;
use sleep_recorder::SleepRecorder;
use std::cmp::Ordering;
use std::fs::File;
use std::io::prelude::*;
use std::str::FromStr;

fn sorted_events(contents: String) -> Vec<Event> {
    fn event_datetime(ev1: &Event, ev2: &Event) -> Ordering {
        return ev1.datetime.cmp(&ev2.datetime);
    }

    let mut events: Vec<Event> = contents
        .lines()
        .map(|s| Event::from_str(s).unwrap())
        .collect();

    events.sort_by(event_datetime);
    return events;
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let mut recorder = SleepRecorder::new();

    recorder.record(&sorted_events(contents));

    let (guard, _total) = recorder.sleepiest_guard();
    let (minute, _times) = recorder.sleepiest_minute(guard);
    println!("{}", guard * minute);

    let (guard, (minute, _times)) = recorder.overall_sleepiest_minute();
    println!("{}", guard * minute);
}
