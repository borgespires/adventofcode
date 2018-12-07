use event::Action;
use event::Event;
use std::collections::HashMap;

#[derive(Debug)]
pub struct SleepRecorder {
    records: HashMap<u32, Vec<u32>>,
    current_guard: u32,
    asleep_minute: u32,
}

impl SleepRecorder {
    pub fn new() -> SleepRecorder {
        SleepRecorder {
            records: HashMap::new(),
            current_guard: 0,
            asleep_minute: 0,
        }
    }

    fn begin_shift(&mut self, guard: &u32) {
        self.records.entry(*guard).or_insert(vec![0; 60]);
        self.current_guard = *guard;
    }

    fn fall_asleep(&mut self, minute: u32) {
        self.asleep_minute = minute;
    }

    fn wake_up(&mut self, minute: u32) {
        let guard_record = self.records.get_mut(&self.current_guard).unwrap();
        for m in self.asleep_minute..minute {
            guard_record[m as usize] += 1;
        }
    }

    pub fn record(&mut self, events: &Vec<Event>) -> HashMap<u32, Vec<u32>> {
        for e in events.iter() {
            match &e.action {
                Action::ShiftBegin { guard } => self.begin_shift(guard),
                Action::Asleep => self.fall_asleep(e.datetime.minute),
                Action::WakeUp => self.wake_up(e.datetime.minute),
                _ => (),
            }
        }

        return self.records.clone();
    }

    pub fn total_sleep_time(&self, guard: u32) -> u32 {
        return self.records.get(&guard).unwrap().iter().sum::<u32>();
    }

    pub fn sleepiest_guard(&self) -> (u32, u32) {
        fn time((_, time): &(u32, u32)) -> u32 {
            return *time;
        }

        return self
            .records
            .keys()
            .map(|&guard| (guard, self.total_sleep_time(guard)))
            .max_by_key(time)
            .unwrap();
    }

    pub fn sleepiest_minute(&self, guard: u32) -> (u32, u32) {
        fn times_found_asleep((_min, times): &(usize, &u32)) -> u32 {
            return **times;
        }

        fn to_u32((minute, times_asleep): (usize, &u32)) -> (u32, u32) {
            return (minute as u32, *times_asleep);
        }

        return self
            .records
            .get(&guard)
            .unwrap()
            .iter()
            .enumerate()
            .max_by_key(times_found_asleep)
            .map(to_u32)
            .unwrap();
    }

    pub fn overall_sleepiest_minute(&self) -> (u32, (u32, u32)) {
        fn times_found_asleep((_guard, (_min, times)): &(u32, (u32, u32))) -> u32 {
            return *times;
        }

        return self
            .records
            .keys()
            .map(|&guard| (guard, self.sleepiest_minute(guard)))
            .max_by_key(times_found_asleep)
            .unwrap();
    }
}
