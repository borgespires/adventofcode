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
    current_guard: u32,
    feel_asleep: u32,
}

impl SleepCollector {
    fn new() -> SleepCollector {
        SleepCollector {
            records: HashMap::new(),
            current_guard: 0,
            feel_asleep: 0,
        }
    }

    fn begin_shift(&mut self, guard: &u32) {
        self.records.entry(*guard).or_insert(vec![0; 60]);
        self.current_guard = *guard;
    }

    fn fall_asleep(&mut self, minute: u32) {
        self.feel_asleep = minute;
    }

    fn wake_up(&mut self, minute: u32) {
        let guard_record = self.records.get_mut(&self.current_guard).unwrap();
        for m in self.feel_asleep..minute {
            guard_record[m as usize] += 1;
        }
    }

    fn sleep_records(&mut self, events: &Vec<Event>) -> HashMap<u32, Vec<u32>> {
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

    let mut collector = SleepCollector::new();
    let records = collector.sleep_records(&events);

    fn sleepiest_minute(record: &Vec<u32>) -> (u32, u32) {
        return record
            .iter()
            .enumerate()
            .max_by_key(|&(_, times_asleep)| times_asleep)
            .map(|(minute, times_asleep)| (minute as u32, *times_asleep))
            .unwrap();
    }

    let sleepiest_guard = records
        .iter()
        .map(|(guard, sleep_record)| (guard, sleep_record.iter().sum::<u32>()))
        .max_by_key(|&(_, sleep_time)| sleep_time)
        .map(|(guard, _)| guard)
        .unwrap();

    let guard_record = records.get(sleepiest_guard).unwrap();

    let (guard, (minute, times)) = records
        .iter()
        .map(|(guard, sleep_record)| (guard, sleepiest_minute(sleep_record)))
        .max_by_key(|&(_, (min, time_slept))| time_slept)
        .unwrap();

    println!("{}", sleepiest_guard * sleepiest_minute(guard_record).0);
    println!("{}", guard * minute);
}