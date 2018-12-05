use std::cmp::Ordering;

#[derive(Debug)]
pub struct DateTime {
    pub year: u32,
    pub month: u32,
    pub day: u32,
    pub hour: u32,
    pub minute: u32,
}

impl DateTime {
    pub fn new(year: u32, month: u32, day: u32, hour: u32, minute: u32) -> DateTime {
        DateTime {
            year,
            month,
            day,
            hour,
            minute,
        }
    }
}

impl Ord for DateTime {
    fn cmp(&self, other: &Self) -> Ordering {
        (self.year, self.month, self.day, self.hour, self.minute).cmp(&(
            other.year,
            other.month,
            other.day,
            other.hour,
            other.minute,
        ))
    }
}

impl PartialOrd for DateTime {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Eq for DateTime {}

impl PartialEq for DateTime {
    fn eq(&self, other: &Self) -> bool {
        (self.year, self.month, self.day, self.hour, self.minute)
            == (other.year, other.month, other.day, other.hour, other.minute)
    }
}
