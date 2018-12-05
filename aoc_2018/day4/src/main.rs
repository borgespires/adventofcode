extern crate regex;
#[macro_use]
extern crate lazy_static;

mod datetime;
mod event;

use event::Action;
use event::Event;
use std::cmp::Ordering;
use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::str::FromStr;

fn event_datetime(ev1: &Event, ev2: &Event) -> Ordering {
    return ev1.datetime.cmp(&ev2.datetime);
}

#[derive(Debug)]
struct SleepCollector {
    records: HashMap<u32, Vec<u32>>,
}

impl SleepCollector {
    fn new() -> SleepCollector {
        SleepCollector {
            records: HashMap::new(),
        }
    }

    fn begin_shift(&self, guard: &u32) {
        println!("{} just started his shift", guard);
        // records.insert(guard, vec![0; 60]); if not exist
        // change current guard
    }

    fn fall_asleep(&self, minute: u32) {
        println!("{} fell asleep at {}", "guard", minute);
    }

    fn wake_up(&self, minute: u32) {
        println!("{} woke up at {}", "guard", minute);
    }

    fn sleep_records(&self, events: &Vec<Event>) -> HashMap<u32, Vec<u32>> {
        for e in events.iter() {
            match &e.action {
                Action::ShiftBegin { guard } => self.begin_shift(guard),
                Action::Asleep => self.fall_asleep(e.datetime.minute),
                Action::WakeUp => self.wake_up(e.datetime.minute),
                _ => println!("Ain't special"),
            }
        }

        return self.records.clone();
    }
}

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

    let collector = SleepCollector::new();
    let records = collector.sleep_records(&events);

    for (guard, sleep) in records.iter() {
        let minutes_asleep = sleep.iter().sum::<u32>();
        let max_minute = sleep.iter().enumerate().max().unwrap();
        println!("{}: \"{:?}\" \"{:?}\"", guard, minutes_asleep, max_minute);
    }
}
